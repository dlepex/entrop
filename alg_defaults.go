package main

import (
	"fmt"
)

type AlgDefaultsStruct struct {
	PBKDF2Rounds        int
	ArgonTime, ArgonMem int
	RHRounds            int
	Salt                []byte
}

var algDefaultsVersions = []AlgDefaultsStruct{
	{Salt: saltV(0), PBKDF2Rounds: 431_998, ArgonTime: 9, ArgonMem: 70656, RHRounds: 558_231},
	{Salt: saltV(1), PBKDF2Rounds: 439_557, ArgonTime: 10, ArgonMem: 65070, RHRounds: 591_438},
	{Salt: saltV(2), PBKDF2Rounds: 437_672, ArgonTime: 11, ArgonMem: 71821, RHRounds: 623_976},
	{Salt: saltV(3), PBKDF2Rounds: 438_130, ArgonTime: 9, ArgonMem: 67199, RHRounds: 613_105},
	{Salt: saltV(4), PBKDF2Rounds: 451_961, ArgonTime: 10, ArgonMem: 72128, RHRounds: 615_711},
	{Salt: saltV(5), PBKDF2Rounds: 462_877, ArgonTime: 11, ArgonMem: 64012, RHRounds: 621_285},
}

// current algorithm defaults
var algDefs = algDefaultsVersions[0]

func SetAlgDefaults(ver int) {
	if ver >= len(algDefaultsVersions) || ver < 0 {
		Terminate("no such version: %d", ver)
	}
	algDefs = algDefaultsVersions[ver]
}

func saltV(idx int) []byte {
	salt, err := embedFS.ReadFile(fmt.Sprintf("embed/salt/salt_v%d", idx))
	Check(err)
	return salt
}
