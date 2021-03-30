package main

import "testing"

func TestGCD(t *testing.T) {
	g := GCD(25, 40)
	t.Logf("gcd =%d", g)
	t.Logf("fcp =%d", FirstCoprime(15, 25))
}

func TestRun(t *testing.T) {
	line := "-a rsha:1001 -c goog -l 50 -s *  hello world"
	pwd := CallEntrop(line)
	t.Logf("pwd = %s", pwd)
}
