package main

import (
	"flag"
	"fmt"
)

type args struct {
	key  string
	text string
}

// constructor of struct "args"
func parseArgs() *args {

	arg1 := flag.String("key", "default", "byte key")
	arg2 := flag.String("text", "default", "text to crypto")

	flag.Parse()

	newArgs := &args{
		key:  *arg1,
		text: *arg2,
	}
	return newArgs
}

func main() {
	args := parseArgs()
	fmt.Println(args)
}
