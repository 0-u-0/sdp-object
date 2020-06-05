package main

import (
	"syscall/js"
)


func jsParse(this js.Value, args []js.Value) interface{} {
	message := args[0].String()
	return parse(message)
}

func main() {
	done := make(chan int, 0)
	js.Global().Set("sdpParse", js.FuncOf(jsParse))
	<-done
}