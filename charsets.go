package main

import (
	"encoding/base64"
	"strings"
)

const (
	CharsetPrintableAscii = "!\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~"
)

var charsetsMap = map[string]string{
	"alnum": "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz", // alphanumeric, default
	"pasc":  CharsetPrintableAscii,
	"goog":  CharsetPrintableAscii, // synonym for pasc
	"num":   "0123456789",
	// some others:
	"bin":   "01",
	"hex":   "0123456789ABCDEF",
	"al":    "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",
	"lower": "abcdefghijklmnopqrstuvwxyz",
	"upper": "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
	"words": "abcdefghijklmnopqrstuvwxyz ",
}

func StringInCharset(a []byte, charsetname string) string {
	if charsetname != "old" {
		cs, ok := charsetsMap[charsetname]
		if !ok {
			Terminate("no such charset: %s", charsetname)
		}
		repl := []byte(cs)
		mod := len(repl)
		b := make([]byte, len(a))
		for i, c := range a {
			b[i] = repl[int(c)%mod]
		}
		return string(b)
	} else {
		b0 := base64.RawURLEncoding.EncodeToString(a)
		repl := strings.NewReplacer("_", "", "-", "", "=", "")
		return repl.Replace(b0)
	}
}

var (
	catLower  = "abcdefghijklmnopqrstuvwxyz"
	catUpper  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	catDigits = "0123456789"
	catPunct  = "!\"#$%&'()*+,-.:;<=>?@[\\]^_`{|}~"
	catsAll   = []string{catLower, catUpper, catDigits, catPunct}
)

// PasswordQuality returns number of char. categories, it is expected to be >= 3 for good pwd.
func PasswordQuality(pwd string) int {
	cats := make(map[string]struct{})
	for _, rn := range pwd {
		for _, cat := range catsAll {
			if strings.ContainsRune(cat, rn) {
				cats[cat] = struct{}{}
			}
		}
	}
	return len(cats)
}

func CharsetSupportsQuality(charsetname string) bool {
	// not all charsets support quality (like number charsets etc)
	return PasswordQuality(charsetsMap[charsetname]) >= 3
}
