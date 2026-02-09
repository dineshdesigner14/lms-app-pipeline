package fldvalidateutil

import (
	"fmt"
	"lmsapieng/include/common/lexicalparserdef"
	"net/mail"
	"regexp"
	"strconv"
	"strings"
)

const (
	fldTypeNumber              = "number"
	fldTypeAlphabet            = "alphabet"
	fldTypeAlphaNumeric        = "an"
	fldTypeAlphaNumericSpecial = "ans"
	fldTypeHex                 = "hex"
	fldTypeEmail               = "email"
	fldTypePattern             = "pattern"
)

func isValidNumericString(value string) bool {
	flag := true
	for i := 0; i < len(value); i++ {
		if value[i] < 48 || value[i] > 57 {
			flag = false
			break
		}
	}
	return flag
}

func isValidHexString(value string) bool {
	flag := true
	for i := 0; i < len(value); i++ {
		if value[i] < 48 {
			flag = false
			break
		}
		if value[i] > 57 && value[i] < 65 {
			flag = false
			break
		}
		if value[i] > 70 && value[i] < 97 {
			flag = false
			break
		}
		if value[i] > 101 {
			flag = false
			break
		}
	}
	return flag
}

func isAlphabet(value string) bool {
	var alphanumeric = regexp.MustCompile("^[a-zA-Z]*$")
	return alphanumeric.MatchString(value)
}

func isAlphaNumeric(value string) bool {
	var alphanumeric = regexp.MustCompile("^[a-zA-Z0-9_]*$")
	return alphanumeric.MatchString(value)
}

func isAlphaNumericSpecial(fldValue string, fldAllowedChars string) bool {
	allowedCharList := strings.Split(fldAllowedChars, ",")
	allowedStr := ""
	for i := 0; i < len(allowedCharList); i++ {
		allowedStr += allowedCharList[i]
	}
	patternStr := fmt.Sprintf("^[a-zA-Z0-9%s]+$", allowedStr)
	alphanumericSpecial := regexp.MustCompile(patternStr)
	return alphanumericSpecial.MatchString(fldValue)
}

func isValidEMail(email string) bool {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return false
	}
	return true
}

func ValidateFld(fldValue string, validateInfo lexicalparserdef.LPValidateFldStruct, rejectDesc *string) int {
	if len(validateInfo.FldMinLen) != 0 {
		minLen, err := strconv.Atoi(validateInfo.FldMinLen)
		if err != nil {
			*rejectDesc = fmt.Sprintf("minLen[%s] is not valid for fldValue[%s] in validation file", validateInfo.FldMinLen, fldValue)
			return -1
		}
		if len(fldValue) < minLen {
			*rejectDesc = fmt.Sprintf("fldValue[%s] is lesser than minLen[%d]", fldValue, minLen)
			return -1
		}
	}
	if len(validateInfo.FldMaxLen) != 0 {
		maxLen, err := strconv.Atoi(validateInfo.FldMaxLen)
		if err != nil {
			*rejectDesc = fmt.Sprintf("maxLen[%s] is not valid for fldValue[%s] in validation file", validateInfo.FldMaxLen, fldValue)
			return -1
		}
		if len(fldValue) > maxLen {
			*rejectDesc = fmt.Sprintf("fldValue[%s] is greater than maxLen[%d]", fldValue, maxLen)
			return -1
		}
	}
	if len(validateInfo.FldAllowedLen) != 0 {
		allowedLenList := strings.Split(validateInfo.FldAllowedLen, ",")
		allowedLenFlag := false
		for i := 0; i < len(allowedLenList); i++ {
			fldLen, err := strconv.Atoi(allowedLenList[i])
			if err != nil {
				*rejectDesc = fmt.Sprintf("FldAllowedLen[%s] is not valid for fldValue[%s] in validation file", validateInfo.FldAllowedLen, fldValue)
				return -1
			}
			if len(fldValue) == fldLen {
				allowedLenFlag = true
				break
			}
		}
		if !allowedLenFlag {
			*rejectDesc = fmt.Sprintf("fldValue[%s] is not having allowed length[%s]", fldValue, validateInfo.FldAllowedLen)
			return -1
		}
	}
	if validateInfo.FldType == fldTypeNumber {
		if !isValidNumericString(fldValue) {
			*rejectDesc = fmt.Sprintf("fldValue[%s] should be a number", fldValue)
			return -1
		}
	} else if validateInfo.FldType == fldTypeHex {
		if !isValidHexString(fldValue) {
			*rejectDesc = fmt.Sprintf("fldValue[%s] should be a hex", fldValue)
			return -1
		}
	} else if validateInfo.FldType == fldTypeAlphabet {
		if !isAlphabet(fldValue) {
			*rejectDesc = fmt.Sprintf("fldValue[%s] should be a alphabet", fldValue)
			return -1
		}
	} else if validateInfo.FldType == fldTypeAlphaNumeric {
		if !isAlphaNumeric(fldValue) {
			*rejectDesc = fmt.Sprintf("fldValue[%s] should be a valid alphanumeric", fldValue)
			return -1
		}
	} else if validateInfo.FldType == fldTypeEmail {
		if !isValidEMail(fldValue) {
			*rejectDesc = fmt.Sprintf("fldValue[%s] should be a valid email", fldValue)
			return -1
		}
	} else if validateInfo.FldType == fldTypeAlphaNumericSpecial {
		if len(validateInfo.FldAllowedChars) == 0 {
			if !isAlphaNumeric(fldValue) {
				*rejectDesc = fmt.Sprintf("fldValue[%s] should be a valid alphanumeric", fldValue)
				return -1
			}
		} else {
			if !isAlphaNumericSpecial(fldValue, validateInfo.FldAllowedChars) {
				*rejectDesc = fmt.Sprintf("fldValue[%s] should be a valid alphanumeric with allowed special chars[%s]", fldValue, validateInfo.FldAllowedChars)
				return -1
			}
		}
	} else if validateInfo.FldType == fldTypePattern {
		if len(validateInfo.FldPattern) == 0 {
			*rejectDesc = fmt.Sprintf("FldPattern xml tag is NULL for fldType[%s]", fldTypePattern)
			return -1
		}
		patternStr := regexp.MustCompile(validateInfo.FldPattern)
		if !patternStr.MatchString(fldValue) {
			*rejectDesc = fmt.Sprintf("fldValue[%s] does not match the pattern[%s]", fldValue, validateInfo.FldPattern)
			return -1
		}
	} else {
		*rejectDesc = fmt.Sprintf("FldType[%s] is not valid", validateInfo.FldType)
		return -1
	}
	return 1
}
