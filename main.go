// +build !js

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)
	opts := Options{}
	opts.Parse(os.Args[1:])
	opts.Init()
	fmt.Println(opts.Password())
}

func Terminate(msg string, args ...interface{}) {
	if msg == "" {
		os.Exit(0)
	}
	log.Fatalf(msg, args...)
}
