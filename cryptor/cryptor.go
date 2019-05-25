package cryptor

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

// Cryptor struct
type Cryptor struct {
	Key []byte
	Msg string
}

// NewCryptor : constructor of struct "Cryptor"
func NewCryptor(k []byte, b string) *Cryptor {
	newCryptor := &Cryptor{
		Key: k,
		Msg: b,
	}
	return newCryptor
}

// copy from this article: https://qiita.com/ken5scal/items/538febf69f9ac98ab292#_reference-9c566e9ec8e63f468270
func padByPkcs7(data []byte) []byte {
	padSize := aes.BlockSize
	if len(data)%aes.BlockSize != 0 {
		padSize = aes.BlockSize - (len(data))%aes.BlockSize
	}

	pad := bytes.Repeat([]byte{byte(padSize)}, padSize)
	return append(data, pad...) // Dots represent it unpack Slice(pad) into individual bytes
}

func unPadByPkcs7(data []byte) []byte {
	padSize := int(data[len(data)-1])
	return data[:len(data)-padSize]
}

// EncryptByCBCMode: have to review
func EncryptByCBCMode(key []byte, plainText string) ([]byte, error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	paddedPlaintext := padByPkcs7([]byte(plainText))
	fmt.Printf("Padded Plain Text in byte format: %v\n", paddedPlaintext)
	cipherText := make([]byte, aes.BlockSize+len(paddedPlaintext)) // cipher text must be larger than plaintext
	iv := cipherText[:aes.BlockSize]                               // Unique iv is required
	_, err = rand.Read(iv)
	if err != nil {
		return nil, err
	}

	cbc := cipher.NewCBCEncrypter(block, iv)
	cbc.CryptBlocks(cipherText[aes.BlockSize:], paddedPlaintext)
	cipherTextBase64 := base64.StdEncoding.EncodeToString(cipherText)
	return []byte(cipherTextBase64), nil
}

// DecryptByCBCMode: have to review
func DecryptByCBCMode(key []byte, cipherTextBase64 []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	cipherText, _ := base64.StdEncoding.DecodeString(string(cipherTextBase64))

	if len(cipherText) < aes.BlockSize {
		panic("cipher text must be longer than blocksize")
	} else if len(cipherText)%aes.BlockSize != 0 {
		panic("cipher text must be multiple of blocksize(128bit)")
	}
	iv := cipherText[:aes.BlockSize] // assuming iv is stored in the first block of ciphertext
	cipherText = cipherText[aes.BlockSize:]
	plainText := make([]byte, len(cipherText))

	cbc := cipher.NewCBCDecrypter(block, iv)
	cbc.CryptBlocks(plainText, cipherText)
	return string(unPadByPkcs7(plainText)), nil
}
