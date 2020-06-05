package main

import (
	"encoding/json"
	"fmt"
	"gortc.io/sdp"
	"syscall/js"
)

type jsonMap = map[string]interface{}

func fibFunc(this js.Value, args []js.Value) interface{} {
	message := args[0].String()

	var (
		s   sdp.Session
		err error
	)
	if s, err = sdp.DecodeSession([]byte(message), s); err != nil {

	}

	d := sdp.NewDecoder(s)

	m := new(sdp.Message)
	if err = d.Decode(m); err != nil {
		fmt.Printf("err: %v", err)
	}
	//fmt.Println("Decoded session", m.Name)
	//fmt.Println("Info:", m.Info)
	//fmt.Println("Origin:", m.Origin)

	//jsonSdp := jsonMap{
	//	"abc":d
	//}
	//jsonSdp,err := json.Marshal(m)
	jsonSdp := jsonMap{
		"name":m.Name,
	}

	jsonSdpStr,err := json.Marshal(jsonSdp)

	return string(jsonSdpStr)

}

func main() {
	done := make(chan int, 0)
	js.Global().Set("fibFunc", js.FuncOf(fibFunc))
	<-done
}