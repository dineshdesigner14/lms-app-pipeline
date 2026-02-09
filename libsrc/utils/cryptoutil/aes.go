package cryptoutil

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

func AESEncryptText(Key string, Data string, EData *string, ErrDesc *string) int {
	if len(Key) != 16 && len(Key) != 24 && len(Key) != 32 {
		*ErrDesc = fmt.Sprintf("Key length should be 16/24/32")
		return -1
	}
	plainText := []byte(Data)

	block, err := aes.NewCipher([]byte(Key))
	if err != nil {
		*ErrDesc = fmt.Sprintf("%s", err.Error())
		return -1
	}
	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		*ErrDesc = fmt.Sprintf("%s", err.Error())
		return -1
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	*EData = base64.RawStdEncoding.EncodeToString(cipherText)
	return 1
}

func AESDecryptText(Key string, EData string, Data *string, ErrDesc *string) int {
	if len(Key) != 16 && len(Key) != 24 && len(Key) != 32 {
		*ErrDesc = fmt.Sprintf("Key length should be 16/24/32")
		return -1
	}
	cipherText, err := base64.RawStdEncoding.DecodeString(EData)
	if err != nil {
		*ErrDesc = fmt.Sprintf("%s", err.Error())
		return -1
	}
	block, err := aes.NewCipher([]byte(Key))
	if err != nil {
		*ErrDesc = fmt.Sprintf("%s", err.Error())
		return -1
	}
	if len(cipherText) < aes.BlockSize {
		*ErrDesc = "Ciphertext block size is too short"
		return -1
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)
	*Data = string(cipherText)
	return 1
}
