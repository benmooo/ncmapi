package necmcrypto_test

import (
	"crypto/aes"
	"testing"

	necmcrypto "github.com/benmooo/necm-api/crypto"
)

func TestEapi(t *testing.T) {
	pt := "this is a test"

	ct := necmcrypto.AesEncryptECB([]byte(pt), make([]byte, 16), []byte{}, aes.BlockSize)
	_pt, err := necmcrypto.AesDecryptECB(ct, make([]byte, 16), []byte{})
	if err != nil {
		t.Error(err)
	}

	if string(_pt) != pt {
		t.Fatal("decrypt and encryption do not match")
	}
}
