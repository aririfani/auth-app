package generalhelper

import (
	"crypto/aes"
	"crypto/cipher"
	cryptoRand "crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func GenerateString(length int) string {
	return StringWithCharset(length, charset)
}

func EncryptString(message string) (string, error) {
	key := []byte("0123456789abcdef")
	byteMsg := []byte(message)
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		err = errors.New(fmt.Sprintf("could not create new cipher: %s", err))
		return "", err
	}

	cipherText := make([]byte, aes.BlockSize+len(byteMsg))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(cryptoRand.Reader, iv); err != nil {
		err = errors.New(fmt.Sprintf("could not encrypt: %s", err))
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], byteMsg)

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func DecryptString(message string) (string, error) {
	key := []byte("0123456789abcdef")
	cipherText, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		err = errors.New(fmt.Sprintf("could not base64 decode: %s", err))
		return "", err
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		err = errors.New(fmt.Sprintf("could not create new cipher: %s", err))
		return "", err
	}

	if len(cipherText) < aes.BlockSize {
		err = errors.New(fmt.Sprintf("could not create new cipher: %s", err))
		return "", err
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText), nil
}
