package main

//タイミングによってはたまに回らなくなる

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/schecthellraiser606/Ramsom_Go_Sample/explorer"
)

func main() {
	decryptKey := "" // Insert starting directory
	dir := "" // Insert starting directory

	key, err := hex.DecodeString(decryptKey)

	if err != nil {
		fmt.Println("Wrong key.")
	} else {

		files := explorer.MapFiles(dir)

		for _, v := range files {
			file, err := ioutil.ReadFile(v)

			if err != nil {
				continue
			}

			decrypted, err := Decrypt(file, key)

			if err != nil {
				continue
			}

			ioutil.WriteFile(v, decrypted, 0644)
		}

		fmt.Println("Files Decrypted.")
	}

	os.Exit(3)
}

func Decrypt(cypherText []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)

	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)

	if err != nil {
		return nil, err
	}

	plainText, err := gcm.Open(nil, cypherText[:gcm.NonceSize()], cypherText[gcm.NonceSize():], nil)

	if err != nil {
		return nil, err
	}

	return plainText, nil
}
