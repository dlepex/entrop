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
	fmt.Println(opts.Password())
}

func Terminate(msg string, args ...interface{}) {
	log.Fatalf(msg, args...)
}
