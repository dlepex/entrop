package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"golang.org/x/term"
)

var (
	verbose = true
)

const (
	DefaultSep     = " "
	DefaultLen     = 25
	DefaultCharset = "alnum"
	DefaultAlg     = "pbs5"
	OptInlineMode  = "i"
)

type Options struct {
	Alg        AlgSpec
	Sep        string
	Ver        uint // alg_defaults version
	CountWords bool
	Words      []string
	PwdLen     uint
	Charset    string
	Quality    bool
	wstr       string
}

type AlgSpec struct {
	Name   string
	Params []int
}

func (opts *Options) Parse(args []string) {
	fs := flag.NewFlagSet("entrop", flag.ExitOnError)
	var spec string
	fs.BoolVar(&verbose, "v", false, "verbose mode")
	isNotInteractive := fs.Bool(OptInlineMode, false, "inline mode, i.e. no hidden inputs")
	fs.StringVar(&spec, "a", DefaultAlg, "algorithm with optional params: e.g. ar:3:32768 or rsha:11111111")
	fs.UintVar(&opts.PwdLen, "l", DefaultLen, "password length")
	fs.UintVar(&opts.Ver, "V", 0, "alg settings ('defaults') version")
	fs.StringVar(&opts.Sep, "s", DefaultSep, "separator")
	fs.StringVar(&opts.Charset, "c", DefaultCharset, "charset, see charsets.go")

	isNotCountWords := fs.Bool("ncw", false, "no word numbering/counting")
	isNoQuality := fs.Bool("nq", false, "no quality check/retry")

	Check(fs.Parse(args))
	opts.CountWords = !*isNotCountWords
	opts.Quality = !*isNoQuality
	opts.Alg = AlgSpecFromStr(spec)

	if *isNotInteractive {
		opts.Words = fs.Args()
	} else {
		opts.readSecrets()
	}
}

// Init must be called after Parse()
func (opts *Options) Init() {
	if opts.Alg.Name == "old" {
		opts.Charset = "old"
		opts.CountWords = false
		opts.Sep = " "
	}
	if len(opts.Words) == 0 {
		Terminate("no words")
	}
	opts.wstr = opts.WordsToString()
	opts.Quality = opts.Quality && len(opts.wstr) >= 6 && CharsetSupportsQuality(opts.Charset)
	SetAlgDefaults(int(opts.Ver))
	if verbose {
		log.Printf("init opts: %+v", opts)
	}
}

func (opts *Options) readSecrets() {
	opts.Words = strings.Fields(readSecretInput("words"))
	opts.Sep = readSecretInput("separator")
	if opts.Sep == "" {
		opts.Sep = DefaultSep
	}

}

func readSecretInput(title string) string {
	fmt.Print(title + ": ")
	a, err := term.ReadPassword(0)
	if err != nil {
		Terminate("failed to read: %s, %s", title, err)
	}
	fmt.Println("")
	return string(a)
}

func (opts *Options) WordsToString() string {
	mapper := WordsMapperNone
	if opts.CountWords {
		mapper = WordsMapperCounting
	}
	return WordsToString(opts.Words, opts.Sep, mapper)
}

func (opts *Options) Password() string {
	wstr := opts.wstr
	for i := 2; ; i++ {
		pwd := opts.tryGenPassword()
		if !opts.Quality || PasswordQuality(pwd) >= 3 {
			return pwd
		}
		opts.wstr = wstr + "@" + strconv.Itoa(i)
		if verbose {
			log.Printf("quality retry: %+v", opts.Words)
		}
	}
}

func (opts *Options) tryGenPassword() string {
	alg, ok := algFuncMap[opts.Alg.Name]
	if !ok {
		Terminate("no such alg: %s", opts.Alg.Name)
	}
	wstr := opts.wstr
	if verbose {
		log.Println("wstr:[" + wstr + "]")
	}
	reqlen := int(opts.PwdLen)
	now := time.Now()
	raw := alg(AlgArgs{[]byte(wstr), reqlen, opts.Alg})
	if verbose {
		log.Printf("duration: %s", time.Since(now))
	}

	pwd := StringInCharset(raw, opts.Charset)
	if len(pwd) <= reqlen {
		return pwd
	}
	return pwd[:reqlen]
}

func Check(e error) {
	if e != nil {
		Terminate(e.Error())
	}
}

func AlgSpecFromStr(s string) (a AlgSpec) {
	w := strings.Split(s, ":")
	if len(w) == 0 {
		Terminate("bad alg spec")
	}
	a.Name = w[0]
	if len(w) == 1 {
		return
	}
	for _, v := range w[1:] {
		p, err := strconv.Atoi(v)
		Check(err)
		a.Params = append(a.Params, p)
	}
	return
}

func (spec *AlgSpec) Param(idx int, defvalue int) int {
	if idx < len(spec.Params) {
		return spec.Params[idx]
	}
	return defvalue
}

func CallEntrop(line string) string {
	opts := Options{}
	args := StringToArgs("-" + OptInlineMode + " " + line)
	opts.Parse(args)
	opts.Init()
	return opts.Password()
}

func StringToArgs(line string) []string {
	//TODO better implementation, supporting quotation (e.g. -s "   ")
	return strings.Fields(line)
}
