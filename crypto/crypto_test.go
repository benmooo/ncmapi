package necmcrypto_test

import (
	"crypto/aes"
	"encoding/hex"
	"testing"

	necmcrypto "github.com/benmooo/ncmapi/crypto"
)

func TestAesEncryptCBC(t *testing.T) {
	expected := "c997554d6eb3120c6676dfce2991baf0"
	key := "0CoJUm6Qyw8W8jud"
	iv := "0102030405060708"
	pt := "plain text"
	ct := necmcrypto.AesEncryptCBC([]byte(pt), []byte(key), []byte(iv), aes.BlockSize)

	if hex.EncodeToString(ct) != expected {
		t.Fatal("aes cbc encryption test fail")
	}
}
func TestAesEncryptECB(t *testing.T) {
	pt := "plain text"
	ct := necmcrypto.AesEncryptECB([]byte(pt), make([]byte, 16), []byte{}, aes.BlockSize)
	_pt, err := necmcrypto.AesDecryptECB(ct, make([]byte, 16), []byte{})
	if err != nil {
		t.Error(err)
	}

	if string(_pt) != pt {
		t.Fatal("decrypt and encryption do not match")
	}
}

func TestEapi(t *testing.T) {
	ct := necmcrypto.EapiEncrypt([]byte("/url"), "plain text")
	t.Log(ct)

	// pt := "plain text"

	// ct := necmcrypto.AesEncryptECB([]byte(pt), make([]byte, 16), []byte{}, aes.BlockSize)
	// _pt, err := necmcrypto.AesDecryptECB(ct, make([]byte, 16), []byte{})
	// if err != nil {
	// 	t.Error(err)
	// }

	// if string(_pt) != pt {
	// 	t.Fatal("decrypt and encryption do not match")
	// }
}

func TestRsa(t *testing.T) {
	expected := "0346440ff4249e431e0aaf6633bdf4bb31ef63563d2103932b053cf3875711da78617966fe383d80ecbbbc81dc8edbeefa095fb47024fdd7b472e4b2b8e919c2fbb9961721d9d12c92c6f037641dd5c2cfb76903707df3364024bf3f6498ad9b1877f9196a08e1ce277039af0e4aecf804016d9b39d700034ce1b964e7574413"
	key := "-----BEGIN PUBLIC KEY-----\nMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDgtQn2JZ34ZC28NWYpAUd98iZ37BUrX/aKzmFbt7clFSs6sXqHauqKWqdtLkF2KexO40H1YTX8z2lSgBBOAxLsvaklV8k4cBFK9snQXE9/DDaFt6Rr7iVZMldczhC0JNgTz+SHXT6CBHuX3e9SdB1Ua44oncaTWz7OBGLbCiK45wIDAQAB\n-----END PUBLIC KEY-----"
	pt := "0CoJUm6Qyw8W8jud"

	ct := necmcrypto.RsaEncrypt([]byte(pt), []byte(key))

	if hex.EncodeToString(ct) != expected {
		t.Error("ras encryption error")
	}
}
