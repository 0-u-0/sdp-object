package main

import (
	"encoding/json"
	"fmt"
	"gortc.io/sdp"
)


type jsonMap = map[string]interface{}

func parse(sdpStr string ) string  {
	var (
		s   sdp.Session
		err error
	)
	if s, err = sdp.DecodeSession([]byte(sdpStr), s); err != nil {

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

func stringify()  {
	
}
