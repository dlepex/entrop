package main

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"

	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/pbkdf2"
)

type AlgArgs struct {
	Str    []byte
	ReqLen int
	Spec   AlgSpec
}
type AlgFunc = func(AlgArgs) []byte

var algFuncMap = map[string]AlgFunc{
	"pbs1": AlgPBKDF2(sha1.New),
	"pbs2": AlgPBKDF2(sha256.New),
	"pbs5": AlgPBKDF2(sha512.New),
	"ar":   AlgArgon2,
	"rh":   AlgRepeatedHash, // only for "short" pwd: length<=64
	// deprecated algs:
	"old":  AlgSimpleHash(md5.New),
	"old2": AlgSimpleHash(sha256.New),
	"old5": AlgSimpleHash(sha512.New),
}

func AlgSimpleHash(hfac func() hash.Hash) AlgFunc {
	return func(args AlgArgs) []byte {
		h := hfac()
		h.Write(args.Str)
		return h.Sum(nil)
	}
}

func AlgPBKDF2(hfac func() hash.Hash) AlgFunc {
	return func(args AlgArgs) []byte {
		return pbkdf2.Key(args.Str, algDefs.Salt, args.Spec.Param(0, algDefs.PBKDF2Rounds), args.ReqLen, hfac)
	}
}

func AlgArgon2(args AlgArgs) []byte {
	return argon2.Key(args.Str, algDefs.Salt, uint32(args.Spec.Param(0, algDefs.ArgonTime)),
		uint32(args.Spec.Param(1, algDefs.ArgonMem)), 1, uint32(args.ReqLen))
}

func AlgRepeatedHash(args AlgArgs) []byte {
	r := RepeatedHash(sha512.New(), args.Spec.Param(0, algDefs.RHRounds), algDefs.Salt, args.Str)
	if len(r) > args.ReqLen {
		return r[:args.ReqLen]
	}
	return r
}

func WordsMapperNone(w string, idx int) string     { return w }
func WordsMapperCounting(w string, idx int) string { return fmt.Sprintf("%d)%s", idx+1, w) }

func WordsToString(words []string, sep string, mapper func(string, int) string) string {
	stk := [256]byte{}
	buf := bytes.NewBuffer(stk[:0])
	buf.WriteString(mapper(words[0], 0))
	for i, w := range words[1:] {
		buf.WriteString(sep)
		buf.WriteString(mapper(w, i+1))
	}
	return buf.String()
}
