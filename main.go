package main

import (
	"flag"
	"fmt"
)

type args struct {
	key  string
	text string
}

// constructor
func parseFlags() *args {
	var (
		a = flag.String("key", "default", "help message for \"key\" option")
		b = flag.String("text", "default", "help message for \"text\" option")
	)

	newArgs := &args{
		key:  *a,
		text: *b,
	}
	return newArgs
}

func main() {
	args := parseFlags()
	fmt.Println(args)
}
