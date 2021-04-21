package main

import _ "embed"

//go:embed embed/salt/salt_v0
var algSaltV0 []byte

//go:embed embed/salt/salt_v1
var algSaltV1 []byte

//go:embed embed/salt/salt_v2
var algSaltV2 []byte

//go:embed embed/salt/salt_v3
var algSaltV3 []byte

//go:embed embed/salt/salt_v4
var algSaltV4 []byte

type AlgDefaultsStruct struct {
	PBKDF2Rounds        int
	ArgonTime, ArgonMem int
	RHRounds            int
	Salt                []byte
}

var algDefaultsVersions = []AlgDefaultsStruct{
	{Salt: algSaltV0, PBKDF2Rounds: 431_998, ArgonTime: 9, ArgonMem: 70656, RHRounds: 558_231},
	{Salt: algSaltV1, PBKDF2Rounds: 439_557, ArgonTime: 10, ArgonMem: 65070, RHRounds: 591_438},
	{Salt: algSaltV2, PBKDF2Rounds: 437_672, ArgonTime: 11, ArgonMem: 71821, RHRounds: 623_976},
	{Salt: algSaltV3, PBKDF2Rounds: 438_130, ArgonTime: 9, ArgonMem: 67199, RHRounds: 613_105},
	{Salt: algSaltV4, PBKDF2Rounds: 451_961, ArgonTime: 10, ArgonMem: 72128, RHRounds: 615_711},
}

// current algorithm defaults
var algDefs = algDefaultsVersions[0]

func SetAlgDefaults(ver int) {
	if ver >= len(algDefaultsVersions) || ver < 0 {
		Terminate("no such version: %d", ver)
	}
	algDefs = algDefaultsVersions[ver]
}
