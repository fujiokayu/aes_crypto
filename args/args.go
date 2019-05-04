package args

import "flag"

// Args struct
type Args struct {
	Key string
	Msg string
}

// ParseArgs : constructor of struct "args"
func ParseArgs() *Args {

	arg1 := flag.String("key", "default", "byte key")
	arg2 := flag.String("msg", "default", "message to crypto")

	flag.Parse()

	newArgs := &Args{
		Key: *arg1,
		Msg: *arg2,
	}
	return newArgs
}
