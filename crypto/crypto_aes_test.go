package crypto

import (
	"encoding/hex"
	"testing"
)

func TestCryptoAesCFBEncrypt(t *testing.T) {
	key, _ := hex.DecodeString("6368616e676520746869732070617373776f726420746f206120736563726574")
	plaintext := []byte("exampleplaintextaaa")

	encrypt, err := CryptoAesCFBEncrypt(plaintext, key)
	t.Log(encrypt)
	t.Log(err)
	plain, err := CryptoAesCFBDecrypt(encrypt, key)
	t.Log(string(plain))
}
