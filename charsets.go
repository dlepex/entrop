package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
)

var ErrNotCharsetSpec = errors.New("not a charset spec")

const (
	CharsetPrintableAscii = "!\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~"
	CharsetDigits         = "0123456789"
	CharsetUpper          = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	CharsetLower          = "abcdefghijklmnopqrstuvwxyz"
	CharsetAlphanum       = CharsetDigits + CharsetUpper + CharsetLower
	CharsetPunct          = "!\"#$%&'()*+,-.:;<=>?@[\\]^_`{|}~"
)

var (
	CharCatsAll = []string{CharsetLower, CharsetUpper, CharsetDigits, CharsetPunct}
)

type Charset string

var predefCharsets = map[string]Charset{
	"alnum": CharsetAlphanum, // alphanumeric, default
	"pasc":  CharsetPrintableAscii,
	"goog":  CharsetPrintableAscii,   // synonym for pasc
	"ora":   CharsetAlphanum + "#$_", // oracle password requirements
	// some others:
	"num":   CharsetDigits,
	"bin":   "01",
	"hex":   "0123456789ABCDEF",
	"al":    CharsetUpper + CharsetLower,
	"lower": CharsetLower,
	"upper": CharsetUpper,
	"words": "abcdefghijklmnopqrstuvwxyz ",
	"old":   CharsetAlphanum, // old charset is for backward compat. and is handled differently
}

// GetCharset resolves charset by name or from spec
func GetCharset(nameOrSpec string) Charset {
	if cs, ok := predefCharsets[nameOrSpec]; ok {
		return cs
	}
	cs, err := CharsetFromSpec(nameOrSpec)
	if err == ErrNotCharsetSpec {
		Terminate("no such charset: %s", nameOrSpec)
	} else if err != nil {
		Terminate("illegal charset spec: %s", err)
	}
	return Charset(cs)
}

// StringInCharset converts bytes into string in some charset
// charsetname - name of predefined charset or charset spec
// see CharsetFromSpec
func StringInCharset(a []byte, charsetname string) string {
	if charsetname != "old" {
		repl := []byte(GetCharset(charsetname))
		mod := len(repl)
		b := make([]byte, len(a))
		for i, c := range a {
			b[i] = repl[int(c)%mod]
		}
		return string(b)
	}
	// backward compatibility:
	b0 := base64.RawURLEncoding.EncodeToString(a)
	repl := strings.NewReplacer("_", "", "-", "", "=", "")
	return repl.Replace(b0)
}

// NumOfCharCats returns number of char. categories, it is expected to be >= 3 for good pwd.
func NumOfCharCats(str string) int {
	cats := make(map[string]struct{})
	for _, rn := range str {
		for _, cat := range CharCatsAll {
			if strings.ContainsRune(cat, rn) {
				cats[cat] = struct{}{}
			}
		}
	}
	return len(cats)
}

func CharsetQuality(charsetname string) int {
	return NumOfCharCats(string(GetCharset(charsetname)))
}

// CharsetFromSpec parses charset specification and produces charset
// Spec format: <categories>|<additional characters>
// B - upper letters, b - lower, 1 - digits, A - all alphanumeric (Bb1)
// Example: A|_@-  charset that contains alphanumeric letters and symbols: _ @ -
func CharsetFromSpec(spec string) (Charset, error) {
	if !strings.ContainsRune(spec, '|') {
		return "", ErrNotCharsetSpec
	}
	split := strings.SplitN(spec, "|", 2)
	catsPart := split[0]
	charsPart := ""
	if len(split) == 2 {
		charsPart = split[1]
	}
	b := strings.Builder{}

	for _, r := range catsPart {
		if !strings.ContainsRune("ABb1", r) {
			return "", fmt.Errorf("no such category: %c", r)
		}
	}

	// write cats part:
	if strings.ContainsRune(catsPart, 'A') {
		b.WriteString(CharsetAlphanum)
	} else {
		if strings.ContainsRune(catsPart, '1') {
			b.WriteString(CharsetDigits)
		}
		if strings.ContainsRune(catsPart, 'B') {
			b.WriteString(CharsetUpper)
		}
		if strings.ContainsRune(catsPart, 'b') {
			b.WriteString(CharsetLower)
		}
	}
	// write additional characters part:
	b.WriteString(charsPart)

	charset := b.String()
	if !EachRuneUnique(charset) {
		return "", fmt.Errorf("repeated characters")
	}
	return Charset(charset), nil
}

func EachRuneUnique(s string) bool {
	m := make(map[rune]struct{})
	for _, r := range s {
		m[r] = struct{}{}
	}
	return len(m) == len([]rune(s))
}
