package main

import (
	"strings"
	"testing"
)

func TestAlgorithms(t *testing.T) {
	// This test ensures alg. implementations are not modified
	cases := []struct{ line, pwd string }{
		{line: "-a pbs5:1234 -c alnum -l 10 -s _ hello world", pwd: "sWXA4jh6m5"},
		{line: "-a pbs2:1234 -c alnum -l 10 -s ^ hello world", pwd: "3ZSOR4xclq"},
		{line: "-a pbs2:1234 -c alnum -l 10 -s ^ -ncw hello world", pwd: "1TUL7DZ0kk"},
		{line: "-a rh:1234 -c goog -l 10 hello world", pwd: "RsuQ)bNCJG"},
		{line: "-a ar:7:123 -c num -l 4 -s ,, hello world", pwd: "7289"},
	}

	for _, tc := range cases {
		if p := CallEntrop(tc.line); p != tc.pwd {
			t.Errorf("line: %s, expects: %s", tc.line, p)
		}
	}
}

func TestWordsToString(t *testing.T) {
	words := []string{"hello", "world", "123"}
	expected := "1)hello,2)world,3)123"
	if WordsToString(words, ",", WordsMapperCounting) != expected {
		t.Error()
	}
	expected = "hello,world,123"
	if WordsToString(words, ",", WordsMapperNone) != expected {
		t.Error()
	}
}

func TestGCD(t *testing.T) {
	g := GCD(25, 40)
	t.Logf("gcd =%d", g)
	cop := FirstCoprime(15, 25)
	t.Logf("fcp =%d", cop)
	if g != 5 && cop != 26 {
		t.Error()
	}
}

func TestStringToArgs(t *testing.T) {
	line := "-a rh:1001 -c goog -l 50 -s \"?   $\"  hello world"
	args := StringToArgs(line)
	t.Logf("args: %+v", strings.Join(args, ","))
}
