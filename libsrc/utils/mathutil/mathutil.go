package mathutil

func IsValidNumericString(value string) bool {
	flag := true
	for i := 0; i < len(value); i++ {
		if value[i] < 48 || value[i] > 57 {
			flag = false
			break
		}
	}
	return flag
}

func IsValidNumericStr(value string, MinLen int, MaxLen int) bool {
	if len(value) < MinLen || len(value) > MaxLen {
		return false
	}
	return IsValidNumericString(value)
}

func IsValidHexString(value string) bool {
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

func IsValidHexStr(value string, MinLen int, MaxLen int) bool {
	if len(value) < MinLen || len(value) > MaxLen {
		return false
	}
	return IsValidHexString(value)
}
