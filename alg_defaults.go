package main

import _ "embed"

//go:embed embed/salt_v0
var algSaltV0 []byte

type AlgDefaultsStruct struct {
	PBKDF2Rounds        int
	ArgonTime, ArgonMem int
	RSHARounds          int
	Salt                []byte
}

var algDefaultsVersions = []AlgDefaultsStruct{
	{ // version#0
		PBKDF2Rounds: 431_998,
		ArgonTime:    9, ArgonMem: 70656,
		RSHARounds: 558_231,
		Salt:       algSaltV0,
	},
}

var algDefs = algDefaultsVersions[0]

func SetAlgDefaults(ver int) {
	if ver >= len(algDefaultsVersions) || ver < 0 {
		Terminate("no such version: %d", ver)
	}
	algDefs = algDefaultsVersions[ver]
}
