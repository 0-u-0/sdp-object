package main

import "C"

//export c_parse
func c_parse(sdpStr *C.char) (int, *C.char) {
	result := parse(C.GoString(sdpStr))
	return len(result), C.CString(result)
}

func main() {

}
