package cryptor

// Cryptor struct
type Cryptor struct {
	Key []byte
	Msg []byte
}

// NewCryptor : constructor of struct "Cryptor"
func NewCryptor(k []byte, b []byte) *Cryptor {
	newCryptor := &Cryptor{
		Key: k,
		Msg: b,
	}
	return newCryptor
}
