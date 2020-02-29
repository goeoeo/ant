package encryptreq

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

// Encrypt string to base64 crypto using AES
func AESEncrypt(key []byte, text string) (string, bool) {
	//填充
	plaintext := ZerosPadding([]byte(text), aes.BlockSize)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", false
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", false
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return base64.URLEncoding.EncodeToString(ciphertext), true
}

// Decrypt from base64 to decrypted string
func AESDecrypt(key []byte, cryptoText string) (string, bool) {
	ciphertext, _ := base64.URLEncoding.DecodeString(cryptoText)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", false
	}

	if len(ciphertext) < aes.BlockSize {
		return "", false
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(ciphertext, ciphertext)

	ciphertext = ZerosUnPadding(ciphertext)

	return fmt.Sprintf("%s", ciphertext), true
}

//去补码
func ZerosUnPadding(origData []byte) []byte {
	return bytes.TrimFunc(origData,
		func(r rune) bool {
			return r == rune(0)
		})
}

//补码
func ZerosPadding(origData []byte, blockSize int) []byte {
	//计算需要补几位数
	padding := blockSize - len(origData)%blockSize
	//在切片后面追加char数量的byte(char)
	padtext := bytes.Repeat([]byte{byte(0)}, padding)

	return append(origData, padtext...)
}
