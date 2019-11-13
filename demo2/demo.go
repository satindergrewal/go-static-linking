package main

import "fmt"

// #cgo CFLAGS: -I/Users/satinder/Documents/GitHub/go-static-linking/include
// #cgo LDFLAGS: /Users/satinder/Documents/GitHub/go-static-linking/build/libgb.a
// #include <junk.h>
import "C"

//Run exported c code
func Run() {
	fmt.Printf("Invoking c library...\n")
	C.x(10)
	fmt.Printf("Done\n")
}

func main() {
	Run()
}
