package main

import (
	"embed"
	"strings"
)

//go:embed embed/**
var embedFS embed.FS

func EntropVersion() string {
	commit, err := embedFS.ReadFile("embed/gitcommit")
	if err != nil {
		return ""
	}
	tag, err := embedFS.ReadFile("embed/gittag")
	if err != nil {
		return ""
	}
	return trimEnd(string(tag)) + "-" + string(commit[:8])
}

func trimEnd(s string) string {
	return strings.Trim(s, " \t\n\r")
}
