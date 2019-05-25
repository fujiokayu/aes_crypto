package args

import (
	"flag"
	"fmt"
	"os"
)

// Args struct
type Args struct {
	EncryptMode bool
	DecryptMode bool
	Key         string
	Msg         string
}

// ParseArgs : constructor of struct "args"
func ParseArgs() *Args {

	arg1 := flag.Bool("e", false, "encrypt mode")
	arg2 := flag.Bool("d", false, "decrypt mode")
	arg3 := flag.String("key", "default", "byte key")
	arg4 := flag.String("msg", "default", "message to crypto")

	flag.Parse()

	newArgs := &Args{
		EncryptMode: *arg1,
		DecryptMode: *arg2,
		Key:         *arg3,
		Msg:         *arg4,
	}

	if newArgs.EncryptMode == newArgs.DecryptMode {
		fmt.Println("have to specify Encrypt Mode or Decrypt Mode.")
		os.Exit(1)
	}
	return newArgs
}
