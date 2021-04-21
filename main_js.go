// +build js

package main

import (
	"fmt"
	"syscall/js"
)

func jsGenPassword(this js.Value, args []js.Value) (res interface{}) {
	defer func() {
		if r := recover(); r != nil {
			res = fmt.Sprintf("error: %s", r)
		}
	}()
	res = CallEntrop(args[0].String())
	return
}

func main() {
	done := make(chan struct{}, 0)
	js.Global().Set("Entrop_GenPassword", js.FuncOf(jsGenPassword))
	js.Global().Set("Entrop_Version", js.ValueOf(EntropVersion()))
	<-done
}

func Terminate(msg string, args ...interface{}) {
	panic(fmt.Sprintf(msg, args...))
}
