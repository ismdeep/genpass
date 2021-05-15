package main

import (
	"fmt"
	"github.com/ismdeep/args"
	"strconv"
)

func main() {
	if args.Exists("--help") {
		fmt.Println(HelpMsg())
		return
	}

	var err error
	n := int64(16)
	if args.Exists("-n") {
		n, err = strconv.ParseInt(args.GetValue("-n"), 10, 64)
		if err != nil {
			n = int64(16)
		}
	}

	if args.Exists("--hex") {
		fmt.Println(RandHex(n))
		return
	}

	fmt.Println(GenPass(args.Exists("-d"), args.Exists("-a"), args.Exists("-A"), args.Exists("--without-fuzzy"), n))
	return
}
