package args

import "flag"

// Args struct
type Args struct {
	key  string
	text string
}

// ParseArgs : constructor of struct "args"
func ParseArgs() *Args {

	arg1 := flag.String("key", "default", "byte key")
	arg2 := flag.String("text", "default", "text to crypto")

	flag.Parse()

	newArgs := &Args{
		key:  *arg1,
		text: *arg2,
	}
	return newArgs
}
