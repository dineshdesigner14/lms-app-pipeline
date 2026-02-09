package cardutil

import (
	"crypto/sha512"
	"fmt"
	"lmsapieng/libsrc/utils/dtutil"
	"strconv"
)

func sumStr(InputStr string) int {
	var lSum int
	lSum = 0
	for i := 0; i < len(InputStr); i++ {
		lSum += int(InputStr[i] - 48)
	}
	return lSum
}

func GetExpiryDate(expiryPeriod int) string {
	tempYear, _ := strconv.Atoi(dtutil.GetDate("DDMMYYYY")[4:])
	return dtutil.GetDate("DDMMYYYY")[:4] + fmt.Sprintf("%04d", expiryPeriod+tempYear)
}

func GenLuhnCheckDigit(CardNum string) string {
	lBytes := []byte(CardNum)
	for i := len(CardNum) - 1; i >= 0; i -= 2 {
		if ((CardNum[i] - 48) * 2) > 9 {
			lBytes[i] = (((CardNum[i] - 48) * 2) - 9) + 48
		} else {
			lBytes[i] = ((CardNum[i] - 48) * 2) + 48
		}
	}
	lcheckDigit := (10 - (sumStr(string(lBytes)) % 10))
	if lcheckDigit == 10 {
		lcheckDigit = 0
	}
	return fmt.Sprintf("%d", lcheckDigit)
}

func VerifyLuhnCheckDigit(CardNum string) bool {
	if GenLuhnCheckDigit(CardNum[:len(CardNum)-1]) == fmt.Sprintf("%c", CardNum[len(CardNum)-1]) {
		return true
	}
	return false
}

func GetHashCardNum(CardNum string) string {
	lCardNum := CardNum
	SalStr := ""
	HashStr := fmt.Sprintf("%s%s", lCardNum, SalStr)
	sha512 := sha512.New()
	sha512.Write([]byte(HashStr))
	shavalue := fmt.Sprintf("%x", sha512.Sum(nil))
	return shavalue
}
