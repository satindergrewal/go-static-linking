package main

import (
	"fmt"
)

// #cgo CFLAGS: -I/Users/satinder/repositories/jl777/libnspv/include -I/Users/satinder/repositories/jl777/libnspv/src/tools/cryptoconditions/include
// #cgo LDFLAGS: /Users/satinder/repositories/jl777/libnspv/.libs/libbtc.a /Users/satinder/repositories/jl777/libnspv/src/tools/cryptoconditions/.libs/libcryptoconditions.a /Users/satinder/repositories/jl777/libnspv/src/tools/nspv-nspv.o
// #include "libbtc-config.h"

// #include <btc/chainparams.h>
// #include <btc/ecc.h>
// #include <btc/net.h>
// #include <btc/netspv.h>
// #include <btc/protocol.h>
// #include <btc/random.h>
// #include <btc/serialize.h>
// #include <btc/sha2.h>
// #include <btc/base58.h>
// #include <btc/tool.h>
// #include <btc/tx.h>
// #include <btc/utils.h>
// #include <btc/wallet.h>

// #include <assert.h>
// #include <getopt.h>
// #include <inttypes.h>
// #include <stdbool.h>
// #include <stdio.h>
// #include <stdlib.h>
// #include <string.h>
// #include <unistd.h>

// #include <nSPV_defs.h>

// #include <cryptoconditions.h>

import "C"

//Run exported c code
func Run() {
	fmt.Printf("Invoking c library...\n")
	cs := C.CString("satinder")
	fmt.Println(cs)
	//C.btc_base58_encode_check(x, x.BitLen(), cs, 10)
	fmt.Printf("Done\n")
}

func main() {
	Run()
}
