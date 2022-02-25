package main

import (
	"testing"
)

func TestCharsets(t *testing.T) {
	for name, set := range charsetsMap {
		if !EachRuneUnique(string(set)) {
			t.Errorf("non unique char in charset: %s", name)
		}
	}

	expectedSize := map[string]int{
		"alnum": 26*2 + 10,
		"pasc":  94,
		"goog":  94,
		"num":   10,
		"lower": 26,
		"upper": 26,
		"al":    26 * 2,
		"bin":   2,
		"hex":   16,
		"words": 26 + 1,
	}
	for name, size := range expectedSize {
		cs := charsetsMap[name]
		if len(cs) != size {
			t.Errorf("bad charset size: %s", name)
		}
	}
}

func TestCharsetSpecs(t *testing.T) {
	cs, err := CharsetFromSpec("A|_-=")
	t.Logf("[%s]err:%s", string(cs), err)
}
