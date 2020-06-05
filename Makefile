

js:
	GOOS=js GOARCH=wasm go build -o main.wasm main.go

lib:
	go build -o libsdpobj.so -buildmode=c-shared csdp.go sdp.go

clibtest:
	 g++ test.cpp -std=c++11 -L. libsdpobj.so