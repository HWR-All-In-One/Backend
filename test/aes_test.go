package test

import (
	"testing"

	aes "github.com/HWR-All-In-One/Backend/internal/pkg/aes"
)

var plaintexts = []string{"hwr-all-in-one-app", "1", "ä", "#somethingtest!_"}
var key = "2f924bc0b82243c60f9f1a5d2afdf7de941239362fc1af2fa9c18ac0ca1f9e2c"

func TestDecryption(t *testing.T) {
	ciphertexts := []string{"4f0dc470fc6c818f4af607b1f562f8982a30a89d661850fdd86fd796dcd5dede045b", "95fd202c6f6245ae49b979efbe822b0f02", "67677c1d34bd09a630dd02ce813e9d880588", "09eca0bac11f8c891245873c97df329d7f95c78bf09d52d75c7d2469b0447f94"}

	for i := 0; i < len(ciphertexts); i++ {
		dec, err := aes.Decrypt(key, ciphertexts[i])

		if err != nil {
			t.Error(err)
		}

		if dec != plaintexts[i] {
			t.Errorf("%s not equal to %s", dec, plaintexts[i])
		}
	}

}

func TestEncryption(t *testing.T) {
	for _, item := range plaintexts {
		enc, err := aes.Encrypt(key, item)
		if err != nil {
			t.Error(err)
		}
		t.Log(enc)
	}

}
