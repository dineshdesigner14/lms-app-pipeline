package cryptoutil

import (
	"crypto/cipher"
	"crypto/des"
	"encoding/hex"
	"fmt"
	"strings"
)

func DESEncryptText(Key string, Data string, EData *string, ErrDesc *string) int {
	var block cipher.Block
	var err error
	var lKey string

	if len(Key) != 16 && len(Key) != 32 && len(Key) != 48 {
		*ErrDesc = fmt.Sprintf("Key:%s length should be 16/32/48", Key)
		return -1
	}
	if !IsValidHexString(Key) {
		*ErrDesc = fmt.Sprintf("Key:%s should be hexadecimal", Key)
		return -1
	}
	if len(Key) == 32 {
		lKey = Key + Key[:16]
	} else {
		lKey = Key
	}

	lData := Data

	KeyBytes, _ := hex.DecodeString(lKey)
	if len(Key) == 16 {
		block, err = des.NewCipher(KeyBytes)
		if err != nil {
			*ErrDesc = fmt.Sprintf("%s", err.Error())
			return -1
		}
	} else if len(Key) == 32 {
		block, err = des.NewTripleDESCipher(KeyBytes)
		if err != nil {
			*ErrDesc = fmt.Sprintf("%s", err.Error())
			return -1
		}
	} else {
		block, err = des.NewTripleDESCipher(KeyBytes)
		if err != nil {
			*ErrDesc = fmt.Sprintf("%s", err.Error())
			return -1
		}
	}
	padding := des.BlockSize - len(lData)%des.BlockSize
	if padding < des.BlockSize {
		for i := 0; i < padding; i++ {
			lData += string(padding)
		}
	}
	iv := make([]byte, des.BlockSize)
	mode := cipher.NewCBCEncrypter(block, iv)
	ciphertext := make([]byte, len(lData))
	mode.CryptBlocks(ciphertext, []byte(lData))
	*EData = BcdToAscii(ciphertext, len(ciphertext)*2)
	return 1
}

func DESDecryptText(Key string, EData string, Data *string, ErrDesc *string) int {
	var block cipher.Block
	var err error
	var lKey string

	if len(Key) != 16 && len(Key) != 32 && len(Key) != 48 {
		*ErrDesc = fmt.Sprintf("Key:%s length should be 16/32/48", Key)
		return -1
	}
	if !IsValidHexString(Key) {
		*ErrDesc = fmt.Sprintf("Key:%s should be hexadecimal", Key)
		return -1
	}
	if len(Key) == 32 {
		lKey = Key + Key[:16]
	} else {
		lKey = Key
	}
	KeyBytes, _ := hex.DecodeString(lKey)
	if len(Key) == 16 {
		block, err = des.NewCipher(KeyBytes)
		if err != nil {
			*ErrDesc = fmt.Sprintf("%s", err.Error())
			return -1
		}
	} else if len(Key) == 32 {
		block, err = des.NewTripleDESCipher(KeyBytes)
		if err != nil {
			*ErrDesc = fmt.Sprintf("%s", err.Error())
			return -1
		}
	} else {
		block, err = des.NewTripleDESCipher(KeyBytes)
		if err != nil {
			*ErrDesc = fmt.Sprintf("%s", err.Error())
			return -1
		}
	}
	iv := make([]byte, des.BlockSize)
	mode := cipher.NewCBCDecrypter(block, iv)

	EDataBytes, _ := hex.DecodeString(EData)
	plaintext := make([]byte, len(EDataBytes))
	mode.CryptBlocks(plaintext, EDataBytes)

	padding := des.BlockSize - len(plaintext)%des.BlockSize
	if padding < des.BlockSize {
		padding := plaintext[len(plaintext)-1]
		plaintext = plaintext[:len(plaintext)-int(padding)]
	}
	*Data = string(plaintext)
	return 1
}

func desEncryptHexBlock(Key string, Data string, EData *string, ErrDesc *string) int {
	var block1, block2, block3 cipher.Block
	var err error

	if len(Key) != 16 && len(Key) != 32 && len(Key) != 48 {
		*ErrDesc = fmt.Sprintf("Key:%s length should be 16/32/48", Key)
		return -1
	}
	if !IsValidHexString(Key) {
		*ErrDesc = fmt.Sprintf("Key:%s should be hexadecimal", Key)
		return -1
	}
	KeyBytes, _ := hex.DecodeString(Key)
	if len(Key) == 16 {
		block1, err = des.NewCipher(KeyBytes)
		if err != nil {
			*ErrDesc = fmt.Sprintf("%s", err.Error())
			return -1
		}
	} else if len(Key) == 32 {
		block1, err = des.NewCipher(KeyBytes[:8])
		if err != nil {
			*ErrDesc = fmt.Sprintf("%s", err.Error())
			return -1
		}
		block2, err = des.NewCipher(KeyBytes[8:16])
		if err != nil {
			*ErrDesc = fmt.Sprintf("%s", err.Error())
			return -1
		}
		block3, err = des.NewCipher(KeyBytes[:8])
		if err != nil {
			*ErrDesc = fmt.Sprintf("%s", err.Error())
			return -1
		}
	} else {
		block1, err = des.NewCipher(KeyBytes[:8])
		if err != nil {
			*ErrDesc = fmt.Sprintf("%s", err.Error())
			return -1
		}
		block2, err = des.NewCipher(KeyBytes[8:16])
		if err != nil {
			*ErrDesc = fmt.Sprintf("%s", err.Error())
			return -1
		}
		block3, err = des.NewCipher(KeyBytes[16:24])
		if err != nil {
			*ErrDesc = fmt.Sprintf("%s", err.Error())
			return -1
		}
	}
	DataBytes, _ := hex.DecodeString(Data)
	ciphertext := make([]byte, len(DataBytes))

	encrypter1 := cipher.NewCBCEncrypter(block1, make([]byte, 8))
	encrypter1.CryptBlocks(ciphertext, DataBytes)

	if len(Key) != 16 {
		decrypter1 := cipher.NewCBCDecrypter(block2, make([]byte, 8))
		decrypter1.CryptBlocks(ciphertext, ciphertext)

		encrypter2 := cipher.NewCBCEncrypter(block3, make([]byte, 8))
		encrypter2.CryptBlocks(ciphertext, ciphertext)
	}

	*EData = hex.EncodeToString(ciphertext)
	return 1
}

func DESEncryptHex(Key string, Data string, EData *string, ErrDesc *string) int {
	var lEData string
	if len(Data) != 16 && len(Data) != 32 && len(Data) != 48 {
		*ErrDesc = fmt.Sprintf("Data:%s length should be 16/32/48", Data)
		return -1
	}
	if !IsValidHexString(Data) {
		*ErrDesc = fmt.Sprintf("Data:%s should be hexadecimal", Data)
		return -1
	}
	if len(Data) == 16 {
		if desEncryptHexBlock(Key, Data[:16], &lEData, ErrDesc) < 0 {
			return -1
		}
		*EData += lEData
	} else if len(Data) == 32 {
		if desEncryptHexBlock(Key, Data[:16], &lEData, ErrDesc) < 0 {
			return -1
		}
		*EData += lEData
		if desEncryptHexBlock(Key, Data[16:32], &lEData, ErrDesc) < 0 {
			return -1
		}
		*EData += lEData
	} else {
		if desEncryptHexBlock(Key, Data[:16], &lEData, ErrDesc) < 0 {
			return -1
		}
		*EData += lEData
		if desEncryptHexBlock(Key, Data[16:32], &lEData, ErrDesc) < 0 {
			return -1
		}
		*EData += lEData
		if desEncryptHexBlock(Key, Data[32:48], &lEData, ErrDesc) < 0 {
			return -1
		}
		*EData += lEData
	}
	*EData = strings.ToUpper(*EData)
	return 1
}

func desDecryptHexBlock(Key string, EData string, Data *string, ErrDesc *string) int {
	var block1, block2, block3 cipher.Block
	var err error

	if len(Key) != 16 && len(Key) != 32 && len(Key) != 48 {
		*ErrDesc = fmt.Sprintf("Key:%s length should be 16/32/48", Key)
		return -1
	}
	if !IsValidHexString(Key) {
		*ErrDesc = fmt.Sprintf("Key:%s should be hexadecimal", Key)
		return -1
	}
	KeyBytes, _ := hex.DecodeString(Key)
	if len(Key) == 16 {
		block1, err = des.NewCipher(KeyBytes)
		if err != nil {
			*ErrDesc = fmt.Sprintf("%s", err.Error())
			return -1
		}
	} else if len(Key) == 32 {
		block1, err = des.NewCipher(KeyBytes[:8])
		if err != nil {
			*ErrDesc = fmt.Sprintf("%s", err.Error())
			return -1
		}
		block2, err = des.NewCipher(KeyBytes[8:16])
		if err != nil {
			*ErrDesc = fmt.Sprintf("%s", err.Error())
			return -1
		}
		block3, err = des.NewCipher(KeyBytes[:8])
		if err != nil {
			*ErrDesc = fmt.Sprintf("%s", err.Error())
			return -1
		}
	} else {
		block1, err = des.NewCipher(KeyBytes[:8])
		if err != nil {
			*ErrDesc = fmt.Sprintf("%s", err.Error())
			return -1
		}
		block2, err = des.NewCipher(KeyBytes[8:16])
		if err != nil {
			*ErrDesc = fmt.Sprintf("%s", err.Error())
			return -1
		}
		block3, err = des.NewCipher(KeyBytes[16:24])
		if err != nil {
			*ErrDesc = fmt.Sprintf("%s", err.Error())
			return -1
		}
	}
	EDataBytes, _ := hex.DecodeString(EData)
	ciphertext := make([]byte, len(EDataBytes))

	encrypter1 := cipher.NewCBCDecrypter(block1, make([]byte, 8))
	encrypter1.CryptBlocks(ciphertext, EDataBytes)

	if len(Key) != 16 {
		decrypter1 := cipher.NewCBCEncrypter(block2, make([]byte, 8))
		decrypter1.CryptBlocks(ciphertext, ciphertext)

		encrypter2 := cipher.NewCBCDecrypter(block3, make([]byte, 8))
		encrypter2.CryptBlocks(ciphertext, ciphertext)
	}

	*Data = hex.EncodeToString(ciphertext)
	return 1
}

func DESDecryptHex(Key string, EData string, Data *string, ErrDesc *string) int {
	var lData string
	if len(EData) != 16 && len(EData) != 32 && len(EData) != 48 {
		*ErrDesc = fmt.Sprintf("EData:%s length should be 16/32/48", EData)
		return -1
	}
	if !IsValidHexString(EData) {
		*ErrDesc = fmt.Sprintf("EData:%s should be hexadecimal", EData)
		return -1
	}
	if len(EData) == 16 {
		if desDecryptHexBlock(Key, EData[:16], &lData, ErrDesc) < 0 {
			return -1
		}
		*Data += lData
	} else if len(EData) == 32 {
		if desDecryptHexBlock(Key, EData[:16], &lData, ErrDesc) < 0 {
			return -1
		}
		*Data += lData
		if desDecryptHexBlock(Key, EData[16:32], &lData, ErrDesc) < 0 {
			return -1
		}
		*Data += lData
	} else {
		if desDecryptHexBlock(Key, EData[:16], &lData, ErrDesc) < 0 {
			return -1
		}
		*Data += lData
		if desDecryptHexBlock(Key, EData[16:32], &lData, ErrDesc) < 0 {
			return -1
		}
		*Data += lData
		if desDecryptHexBlock(Key, EData[32:48], &lData, ErrDesc) < 0 {
			return -1
		}
		*Data += lData
	}
	*Data = strings.ToUpper(*Data)
	return 1
}
