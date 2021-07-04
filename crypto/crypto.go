package necmcrypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"math/big"
	"strings"
)

const iv = "0102030405060708"
const presetKey = "0CoJUm6Qyw8W8jud"

// const linuxapiKey = "rFgB&h#%2?^eDg:Q"
const base62 = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const publicKey = "-----BEGIN PUBLIC KEY-----\nMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDgtQn2JZ34ZC28NWYpAUd98iZ37BUrX/aKzmFbt7clFSs6sXqHauqKWqdtLkF2KexO40H1YTX8z2lSgBBOAxLsvaklV8k4cBFK9snQXE9/DDaFt6Rr7iVZMldczhC0JNgTz+SHXT6CBHuX3e9SdB1Ua44oncaTWz7OBGLbCiK45wIDAQAB\n-----END PUBLIC KEY-----"

const eapiKey = "e82ckenh8dichen8"
const secretKeyLen = 16

func AesEncryptCBC(text []byte, key []byte, iv []byte, blockSize int) []byte {
	btext := PKCS5Padding(text, blockSize)
	ciphertext := make([]byte, len(btext))

	block, _ := aes.NewCipher(key)
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, btext)

	return ciphertext
}

func AesEncryptECB(pt, key, iv []byte, blockSize int) []byte {
	cipher, _ := aes.NewCipher(key)

	// blocks num
	bn := len(pt)/blockSize + 1
	_pt := make([]byte, bn*blockSize)
	copy(_pt, pt)

	// padding num
	padding := byte(bn*blockSize - len(pt))
	// fill with padding
	for i := 0; i < int(padding); i++ {
		_pt[len(_pt)-1-i] = padding
	}

	ciphertext := make([]byte, len(_pt))
	size := cipher.BlockSize()
	for bs, be := 0, size; bs < len(_pt); bs, be = bs+size, be+size {
		cipher.Encrypt(ciphertext[bs:be], _pt[bs:be])
	}
	return ciphertext
}
func AesDecryptECB(ct, key, iv []byte) ([]byte, error) {
	if len(ct)%aes.BlockSize != 0 {
		return nil, errors.New("aes block size division fail")
	}
	cipher, _ := aes.NewCipher(key)
	pt := make([]byte, len(ct))
	size := cipher.BlockSize()
	for bs, be := 0, size; bs < len(ct); bs, be = bs+size, be+size {
		cipher.Decrypt(pt[bs:be], ct[bs:be])
	}

	padding := pt[len(pt)-1]
	if padding > aes.BlockSize {
		return nil, errors.New("decryption error")
	}

	return pt[:len(pt)-int(padding)], nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := (blockSize - len(ciphertext)%blockSize)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func RsaEncrypt(text []byte, pk []byte) []byte {
	data := make([]byte, 128-len(text))
	btext := append(data, []byte(text)...)

	block, _ := pem.Decode(pk)
	pubInterface, _ := x509.ParsePKIXPublicKey(block.Bytes)

	pubKey := pubInterface.(*rsa.PublicKey)
	ciphertext := new(big.Int)
	e := big.NewInt(int64(pubKey.E))
	payload := new(big.Int).SetBytes(btext)
	ciphertext.Exp(payload, e, pubKey.N)

	return ciphertext.Bytes()
}

func WeapiEncrypt(text []byte) map[string]string {
	buf := make([]byte, secretKeyLen)
	rand.Read(buf)

	sk := make([]byte, secretKeyLen)
	reverseSK := make([]byte, secretKeyLen)

	for index, val := range buf {
		sk[index] = byte(base62[val%62])
		reverseSK[secretKeyLen-index-1] = byte(base62[val%62])
	}

	_params := base64.StdEncoding.EncodeToString(
		AesEncryptCBC(text, []byte(presetKey), []byte(iv), aes.BlockSize))
	params := base64.StdEncoding.EncodeToString(
		AesEncryptCBC([]byte(_params), sk, []byte(iv), aes.BlockSize))

	encSecKey := hex.EncodeToString(RsaEncrypt(reverseSK, []byte(publicKey)))

	log.Printf("params=%v", params)
	log.Printf("encSecKey=%v", encSecKey)

	return map[string]string{
		"params":    params,
		"encSecKey": encSecKey,
	}
}

func Md5Hex(message string) string {
	md := md5.New()
	md.Write([]byte(message))
	return hex.EncodeToString(md.Sum(nil))
}

func EapiEncrypt(url []byte, data interface{}) map[string]string {
	text, _ := json.Marshal(data)
	message := fmt.Sprintf("nobody%suse%smd5forencrypt", string(url), text)
	_hash := md5.Sum([]byte(message))
	digest := hex.EncodeToString(_hash[:])

	_data := fmt.Sprintf("%s-36cd479b6b5-%s-36cd479b6b5-%s", url, text, digest)
	return map[string]string{
		"params": strings.ToUpper(hex.EncodeToString(
			AesEncryptECB([]byte(_data), []byte(eapiKey), []byte{}, aes.BlockSize),
		)),
	}
}

func EapiDecrypt(ct []byte) ([]byte, error) {
	return AesDecryptECB(ct, []byte(eapiKey), []byte{})
}

type Crypto uint8

const (
	CtryptoWeapi Crypto = iota + 1
	CtryptoEapi
	CtryptoLinuxapi
)
