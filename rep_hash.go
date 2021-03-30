package main

import (
	"hash"
)

// partly inspired by 7z KDF but a bit more creative :)
func RepeatedHash(h hash.Hash, numrounds int, salt, input []byte) []byte {
	buf := make([]byte, h.Size())
	const ctrlen = 8
	ctr := [ctrlen]byte{}
	inplen := len(input)
	if inplen <= 1 {
		Terminate("repeated hash requires longer input")
	}
	mod := FirstCoprime(inplen, 25)
	k := 0
	for ; numrounds != 0; numrounds-- {
		h.Write(salt)
		h.Write(input)
		for i := 0; i < ctrlen; i++ {
			k++
			ctr[i]++
			if k%mod == 0 {
				ctr[i] += input[k%inplen] % 2
			}
			if ctr[i] == 0 {
				break
			}
		}
		h.Write(ctr[:])
		b := h.Sum(buf[:0])
		h.Reset()
		h.Write(b)
	}
	return buf
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func FirstCoprime(x, after int) int {
	a := after
	for GCD(a, x) != 1 {
		a++
	}
	return a
}
