package encrypt

import (
	"crypto/aes"
	"encoding/hex"
)

func AESEncrypt(key, plaintext string) (string, error) {
	c, err := aes.NewCipher([]byte(key))

	if err != nil {
		return "", err
	}

	out := make([]byte, len(plaintext))

	c.Encrypt(out, []byte(plaintext))

	return hex.EncodeToString(out), nil
}

func AESDecrypt(key, ciphertext string) (string, error) {
	ct, _ := hex.DecodeString(ciphertext)
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	pt := make([]byte, len(ct))
	c.Decrypt(pt, ct)

	return string(pt), nil
}
