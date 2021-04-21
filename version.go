package main

import (
	"embed"
	"strings"
)

//go:embed embed/**
var content embed.FS

func EntropVersion() string {
	commit, err := content.ReadFile("embed/gitcommit")
	if err != nil {
		return ""
	}
	tag, err := content.ReadFile("embed/gittag")
	if err != nil {
		return ""
	}
	return trimEnd(string(tag)) + "-" + string(commit[:8])
}

func trimEnd(s string) string {
	return strings.Trim(s, " \t\n\r")
}
