package cryptoutil

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func IsValidHexString(value string) bool {
	_, err := hex.DecodeString(value)
	return err == nil
}

func atohex(val uint8) uint8 {
	if (val >= 48) && (val <= 57) {
		return (val - 48)
	} else if (val >= 65) && (val <= 70) {
		return (val - 55)
	} else if (val >= 97) && (val <= 102) {
		return (val - 87)
	} else {
		return val
	}
}

func AsciiToBcd(ASCIIStr string, StrSize int) []byte {
	var lnibble, rnibble uint8
	var i, j, lbyteLen int
	var lASCIIStr string
	if StrSize%2 == 0 {
		lASCIIStr = ASCIIStr
		lbyteLen = StrSize / 2
	} else {
		lASCIIStr = "0" + ASCIIStr
		lbyteLen = StrSize/2 + 1
	}
	ByteVal := make([]byte, lbyteLen)
	for i, j, lnibble, rnibble = 0, 0, 0, 0; i < lbyteLen; i++ {
		lnibble = lASCIIStr[j]
		lnibble = atohex(lnibble)
		rnibble = lASCIIStr[j+1]
		rnibble = atohex(rnibble)
		ByteVal[i] = (lnibble << 4) | (rnibble & 0xf)
		j += 2
	}
	return ByteVal
}

func BcdToAscii(byteStr []byte, StrSize int) string {
	i := 0
	j := 0
	ASCIIStr := make([]byte, StrSize)
	for {
		if j >= StrSize {
			break
		}
		ASCIIStr[j] = ((byteStr[i] & 0xf0) >> 4)
		if ASCIIStr[j] >= 0 && ASCIIStr[j] <= 9 {
			ASCIIStr[j] += 48
		} else {
			ASCIIStr[j] += 55
		}
		ASCIIStr[j+1] = byteStr[i] & 0x0f
		if ASCIIStr[j+1] >= 0 && ASCIIStr[j+1] <= 9 {
			ASCIIStr[j+1] += 48
		} else {
			ASCIIStr[j+1] += 55
		}
		j += 2
		i++
	}
	return string(ASCIIStr)
}

func IsValidKey(Key string) bool {
	// if len(Key) != 16 && len(Key) != 32 && len(Key) != 48 {
	// 	return false
	// }
	if !IsValidHexString(Key) {
		return false
	}
	return true
}

func XorStrings(s1, s2 string) (int, string) {
	b1, _ := hex.DecodeString(s1)
	b2, _ := hex.DecodeString(s2)
	if len(b1) != len(b2) {
		return -1, ""
	}
	result := make([]byte, len(b1))
	for i := 0; i < len(b1); i++ {
		result[i] = b1[i] ^ b2[i]
	}
	// //trace.LogHex(debugdef.DEBUG_LEVEL_TEST, result)
	return 1, strings.ToUpper(hex.EncodeToString(result))
}

func IsValidAnsiPinBlock(PinBlock string, PINLen int) bool {
	tempStr := PinBlock[:2]
	tLen, err := strconv.Atoi(tempStr)
	if err != nil {
		return false
	}
	if tLen != PINLen {
		return false
	}
	return true
}

func GenAnsiPinBlock(CardNum string, PIN string, PinBlock *string) {
	ClearBlock := fmt.Sprintf("%02d%s", len(PIN), PIN)
	ClearBlock += strings.Repeat("F", 16-len(ClearBlock))
	AcctBlock := strings.Repeat("0", 4)
	AcctBlock += CardNum[len(CardNum)-13 : len(CardNum)-13+12]
	_, *PinBlock = XorStrings(ClearBlock, AcctBlock)
}

func DecimilizeData(Data string) string {
	decData := make([]byte, 0)
	for i := 0; i < len(Data); i++ {
		if Data[i] == 'A' || Data[i] == 'a' {
			decData = append(decData, '0')
		} else if Data[i] == 'B' || Data[i] == 'b' {
			decData = append(decData, '1')
		} else if Data[i] == 'C' || Data[i] == 'c' {
			decData = append(decData, '2')
		} else if Data[i] == 'D' || Data[i] == 'd' {
			decData = append(decData, '3')
		} else if Data[i] == 'E' || Data[i] == 'e' {
			decData = append(decData, '4')
		} else if Data[i] == 'F' || Data[i] == 'f' {
			decData = append(decData, '5')
		} else {
			decData = append(decData, Data[i])
		}
	}
	return string(decData)
}

func GetClearPIN(CardNum string, ClearPinBlock string, PINLen int) string {
	AnsiPanBlock := strings.Repeat("0", 4)
	AnsiPanBlock += CardNum[len(CardNum)-13 : len(CardNum)-13+12]
	_, tempStr := XorStrings(ClearPinBlock, AnsiPanBlock)
	return tempStr[2 : 2+PINLen]
}

func SubtractWithoutBorrow(Data1 string, Data2 string) string {
	resultData := make([]byte, len(Data1))
	for i := 0; i < len(Data1); i++ {
		if Data1[i] < Data2[i] {
			resultData[i] = Data1[i] - Data2[i] + 58
		} else {
			resultData[i] = Data1[i] - Data2[i] + 48
		}
	}
	return string(resultData)
}
