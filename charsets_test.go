package main

import (
	"testing"
)

func TestCharsets(t *testing.T) {
	for name, set := range predefCharsets {
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
		"old":   26*2 + 10,
	}
	for name, size := range expectedSize {
		cs := GetCharset(name)
		if len(cs) != size {
			t.Errorf("bad charset size: %s", name)
		}
	}
}

func TestCharsetSpecs(t *testing.T) {
	cs, err := CharsetFromSpec("A|_-=")
	t.Logf("[%s]err:%s", string(cs), err)
	cs = GetCharset("B|123")
	if len(cs) != 26+3 {
		t.Errorf("bad charset")
	}
	cs = GetCharset("1|()")
	if len(cs) != 10+2 {
		t.Errorf("bad charset")
	}
	cs = GetCharset("b1|")
	t.Logf("%s", cs)
	if len(cs) != 10+26 {
		t.Errorf("bad charset")
	}
	cs = GetCharset("|12345")
	t.Logf("%s", cs)
	if len(cs) != 5 {
		t.Errorf("bad charset")
	}
	if GetCharset("Bb1|+!$") != GetCharset("b1B|!$+") {
		t.Errorf("charset spec order should not matter")
	}
}
