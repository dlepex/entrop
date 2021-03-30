package main

import _ "embed"

//go:embed embed/salt_v1
var algSaltV1 []byte

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
		Salt:       algSaltV1,
	},
}

var algDefs = algDefaultsVersions[0]

func SetAlgDefaults(ver int) {
	if ver >= len(algDefaultsVersions) || ver < 0 {
		algDefs = algDefaultsVersions[0]
	}
	algDefs = algDefaultsVersions[ver]
}
