package main

import (
	"aes_crypto/args"
	"aes_crypto/cryptor"
	"fmt"
)

func main() {
	args := args.ParseArgs()
	cryptor := cryptor.NewCryptor([]byte(args.Key), []byte(args.Msg))

	fmt.Println(cryptor)
}
