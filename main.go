package main

import (
	"aes_crypto/args"
	"fmt"
)

func main() {
	args := args.ParseArgs()
	fmt.Println(args)
}
