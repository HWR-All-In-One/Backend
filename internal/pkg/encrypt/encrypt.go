package encrypt

import (
	"crypto/aes"
	"encoding/hex"
	"fmt"
)

const AES256 = 32

func getKey(key string) []byte {
	fmt.Println(key[:AES256])
	return []byte(key[:AES256])
}

func AESEncrypt(key, plaintext string) (string, error) {
	fmt.Println(len(plaintext))
	fmt.Println(key, plaintext)
	fmt.Println(len(getKey(key)))
	fmt.Println(getKey(key))
	c, err := aes.NewCipher(getKey(key))

	if err != nil {
		return "", err
	}

	out := make([]byte, len(plaintext))

	c.Encrypt(out, []byte(plaintext))

	return hex.EncodeToString(out), nil
}

func AESDecrypt(key, ciphertext string) (string, error) {
	ct, _ := hex.DecodeString(ciphertext)
	c, err := aes.NewCipher(getKey(key))
	if err != nil {
		return "", err
	}

	pt := make([]byte, len(ct))
	c.Decrypt(pt, ct)

	return string(pt), nil
}
