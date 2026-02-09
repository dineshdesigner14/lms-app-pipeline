package dtutil

import (
	"fmt"
	"lmsapieng/libsrc/utils/mathutil"
	"strconv"
	"strings"
	"time"
)

const (
	TimeUnitHrs       = "Hrs"
	TimeUnitMins      = "Mins"
	TimeUnitSecs      = "Secs"
	TimeUnitMilliSecs = "MilliSecs"
	TimeUnitMicroSecs = "MicroSecs"
	TimeUnitNanoSecs  = "NanoSecs"
)

func DiffTimeStr(CurrTimeStr string, OldTimeStr string, TimeUnit string) float64 {
	time1, _ := time.Parse("02012006150405", CurrTimeStr)
	time2, _ := time.Parse("02012006150405", OldTimeStr)
	return DiffTime(time1, time2, TimeUnit)
}

func DiffTime(CurrTime time.Time, OldTime time.Time, TimeUnit string) float64 {
	diff := CurrTime.Sub(OldTime)
	if strings.EqualFold(TimeUnit, TimeUnitSecs) {
		return float64(diff.Seconds())
	} else if strings.EqualFold(TimeUnit, TimeUnitNanoSecs) {
		return float64(diff.Nanoseconds())
	} else if strings.EqualFold(TimeUnit, TimeUnitMilliSecs) {
		return float64(diff.Nanoseconds() / 1000000)
	} else if strings.EqualFold(TimeUnit, TimeUnitMicroSecs) {
		return float64(diff.Nanoseconds() / 1000)
	} else if strings.EqualFold(TimeUnit, TimeUnitMins) {
		return float64(diff.Minutes())
	} else if strings.EqualFold(TimeUnit, TimeUnitHrs) {
		return float64(diff.Hours())
	} else {
		return float64(diff.Seconds())
	}
}

func GetDate(DateFormat string) string {
	var dateval string
	currentDate := time.Now()
	lDateFormat := strings.ToUpper(DateFormat)
	switch lDateFormat {
	case "DDMMYY":
		{
			dateval = fmt.Sprintf("%s", currentDate.Format("020106"))
		}
		break
	case "DDMMYYYY":
		{
			dateval = fmt.Sprintf("%s", currentDate.Format("02012006"))
		}
		break
	case "DD/MM/YY":
		{
			dateval = fmt.Sprintf("%s", currentDate.Format("02/01/06"))
		}
		break
	case "DD/MM/YYYY":
		{
			dateval = fmt.Sprintf("%s", currentDate.Format("02/01/2006"))
		}
		break
	case "DD-MM-YY":
		{
			dateval = fmt.Sprintf("%s", currentDate.Format("02-01-06"))
		}
		break
	case "DD-MM-YYYY":
		{
			dateval = fmt.Sprintf("%s", currentDate.Format("02-01-2006"))
		}
		break
	case "YYYYMMDD":
		{
			dateval = fmt.Sprintf("%s", currentDate.Format("20060102"))
		}
		break
	default:
		dateval = fmt.Sprintf("%s", currentDate.Format("02012006"))
	}
	return dateval
}

func GetTime(TimeFormat string) string {
	var timeval string
	lTimeFormat := strings.ToUpper(TimeFormat)
	currentTime := time.Now()
	nanos := currentTime.UnixNano()
	millis := nanos / 1000000
	switch lTimeFormat {
	case "HHMMSS":
		{
			timeval = fmt.Sprintf("%s", currentTime.Format("150405"))
		}
		break
	case "HH:MM:SS":
		{
			timeval = fmt.Sprintf("%s", currentTime.Format("15:04:05"))
		}
		break
	case "HHMMSSMS":
		{
			timeval = fmt.Sprintf("%s%d", currentTime.Format("150405"), millis)
		}
		break
	case "HH:MM:SS:MS":
		{
			timeval = fmt.Sprintf("%s:%d", currentTime.Format("15:04:05"), millis)
		}
		break
	default:
		timeval = fmt.Sprintf("%s", currentTime.Format("150405"))
	}
	return timeval
}

func IsValidYear(InputDate string, DateFormat string) bool {
	lDateFormat := strings.ToUpper(DateFormat)
	switch lDateFormat {
	case "YYMM":
		{
			yy := InputDate[:2]
			lYY, _ := strconv.Atoi(yy)
			if lYY < 0 || lYY > 99 {
				return false
			}
		}
		break
	}
	return true
}

func IsValidMonth(InputDate string, DateFormat string) bool {
	lDateFormat := strings.ToUpper(DateFormat)
	switch lDateFormat {
	case "MMDD":
		{
			mm := InputDate[:2]
			lMM, _ := strconv.Atoi(mm)
			if lMM < 1 || lMM > 12 {
				return false
			}
		}
		break
	case "YYMM":
		{
			mm := InputDate[2:]
			lMM, _ := strconv.Atoi(mm)
			if lMM < 1 || lMM > 12 {
				return false
			}
		}
		break
	}
	return true
}

func IsValidDay(InputDate string, DateFormat string) bool {
	lDateFormat := strings.ToUpper(DateFormat)
	switch lDateFormat {
	case "MMDD":
		{
			mm := InputDate[:2]
			dd := InputDate[2:]
			lMM, _ := strconv.Atoi(mm)
			if lMM < 1 || lMM > 12 {
				return false
			}
			lDD, _ := strconv.Atoi(dd)
			if lMM == 1 || lMM == 3 || lMM == 5 || lMM == 7 || lMM == 8 || lMM == 10 || lMM == 12 {
				if lDD < 0 || lDD > 31 {
					return false
				}
			} else if lMM == 4 || lMM == 6 || lMM == 9 || lMM == 11 {
				if lDD < 0 || lDD > 30 {
					return false
				}
			} else {
				if lDD < 0 || lDD > 29 {
					return false
				}
			}
		}
		break
	}
	return true
}

func IsValidHour(InputHr string) bool {
	lHour, _ := strconv.Atoi(InputHr)
	if lHour < 0 || lHour > 23 {
		return false
	}
	return true
}

func IsValidMin(InputMin string) bool {
	lMin, _ := strconv.Atoi(InputMin)
	if lMin < 0 || lMin > 59 {
		return false
	}
	return true
}

func IsValidSec(InputSec string) bool {
	lSec, _ := strconv.Atoi(InputSec)
	if lSec < 0 || lSec > 59 {
		return false
	}
	return true
}

func TransformDateInternal(InputDate string, DateFormat string) (int, string) {
	var OutputDate string
	lDateFormat := strings.ToUpper(DateFormat)
	switch lDateFormat {
	case "MMDD":
		{
			mm := InputDate[:2]
			dd := InputDate[2:]
			if !IsValidMonth(InputDate, lDateFormat) {
				return -1, ""
			}
			if !IsValidDay(InputDate, lDateFormat) {
				return -2, ""
			}
			TempStr := GetDate("DDMMYYYY")
			OutputDate = dd + mm + TempStr[4:8]
		}
		break
	}
	return 1, OutputDate
}

func TransformTimeInternal(InputTime string, TimeFormat string) (int, string) {
	var OutputTime string
	lTimeFormat := strings.ToUpper(TimeFormat)
	switch lTimeFormat {
	case "HHMMSS":
		{
			hh := InputTime[:2]
			mm := InputTime[2:4]
			ss := InputTime[4:]
			if !IsValidHour(hh) {
				return -1, ""
			}
			if !IsValidMin(mm) {
				return -2, ""
			}
			if !IsValidSec(ss) {
				return -3, ""
			}
			OutputTime = InputTime
		}
		break
	}
	return 1, OutputTime
}

func TransformInternalExpiryDate(InputExpiryYear string, ExpYearFormat string) (int, string) {
	var OutputExpiryYear string
	lExpYearFormat := strings.ToUpper(ExpYearFormat)
	switch lExpYearFormat {
	case "YYMM":
		{
			yy := InputExpiryYear[:2]
			mm := InputExpiryYear[2:]
			if !IsValidYear(InputExpiryYear, lExpYearFormat) {
				return -1, ""
			}
			if !IsValidMonth(InputExpiryYear, lExpYearFormat) {
				return -2, ""
			}
			OutputExpiryYear = mm + yy
		}
		break
	}
	return 1, OutputExpiryYear
}

func IsValidTime(value string) bool {
	var tval int
	if len(value) != 6 {
		return false
	}
	if !mathutil.IsValidNumericString(value) {
		return false
	}
	timeSlice := []rune(value)
	hour := string(timeSlice[0:2])
	min := string(timeSlice[2:4])
	sec := string(timeSlice[4:6])
	tval, _ = strconv.Atoi(hour)
	if tval < 0 || tval > 23 {
		return false
	}
	tval, _ = strconv.Atoi(min)
	if tval < 0 || tval > 59 {
		return false
	}
	tval, _ = strconv.Atoi(sec)
	if tval < 0 || tval > 59 {
		return false
	}
	return true
}

func IsValidDateFormat(dateValue string, dateFormat string) bool {
	if strings.EqualFold(dateFormat, "YYMM") {
		if len(dateValue) != 4 {
			return false
		}
		if !mathutil.IsValidNumericString(dateValue) {
			return false
		}
		timeSlice := []rune(dateValue)
		year := string(timeSlice[:2])
		mon := string(timeSlice[2:])
		tval, _ := strconv.Atoi(mon)
		if tval < 1 || tval > 12 {
			return false
		}
		tval, _ = strconv.Atoi(year)
		if tval < 0 || tval > 99 {
			return false
		}
	} else if strings.EqualFold(dateFormat, "DDMMYYYY") {
		var tval int
		if len(dateValue) != 8 {
			return false
		}
		if !mathutil.IsValidNumericString(dateValue) {
			return false
		}
		timeSlice := []rune(dateValue)
		day := string(timeSlice[0:2])
		mon := string(timeSlice[2:4])
		year := string(timeSlice[6:8])

		if mon == "01" || mon == "03" || mon == "05" || mon == "07" || mon == "08" || mon == "10" || mon == "12" {
			tval, _ = strconv.Atoi(day)
			if tval < 1 || tval > 31 {
				return false
			}
		} else if mon == "04" || mon == "06" || mon == "09" || mon == "11" {
			tval, _ = strconv.Atoi(day)
			if tval < 1 || tval > 30 {
				return false
			}
		} else {
			tval, _ = strconv.Atoi(year)
			if tval%4 == 0 {
				tval, _ = strconv.Atoi(day)
				if tval < 1 || tval > 29 {
					return false
				}
			} else {
				tval, _ = strconv.Atoi(day)
				if tval < 1 || tval > 28 {
					return false
				}
			}
		}
		tval, _ = strconv.Atoi(mon)
		if tval < 1 || tval > 12 {
			return false
		}
		tval, _ = strconv.Atoi(year)
		if tval < 0 || tval > 99 {
			return false
		}
	} else if strings.EqualFold(dateFormat, "DDMMYY") {
		var tval int
		if len(dateValue) != 6 {
			return false
		}
		if !mathutil.IsValidNumericString(dateValue) {
			return false
		}
		timeSlice := []rune(dateValue)
		day := string(timeSlice[0:2])
		mon := string(timeSlice[2:4])
		year := string(timeSlice[4:6])

		if mon == "01" || mon == "03" || mon == "05" || mon == "07" || mon == "08" || mon == "10" || mon == "12" {
			tval, _ = strconv.Atoi(day)
			if tval < 1 || tval > 31 {
				return false
			}
		} else if mon == "04" || mon == "06" || mon == "09" || mon == "11" {
			tval, _ = strconv.Atoi(day)
			if tval < 1 || tval > 30 {
				return false
			}
		} else {
			tval, _ = strconv.Atoi(year)
			if tval%4 == 0 {
				tval, _ = strconv.Atoi(day)
				if tval < 1 || tval > 29 {
					return false
				}
			} else {
				tval, _ = strconv.Atoi(day)
				if tval < 1 || tval > 28 {
					return false
				}
			}
		}
		tval, _ = strconv.Atoi(mon)
		if tval < 1 || tval > 12 {
			return false
		}
		tval, _ = strconv.Atoi(year)
		if tval < 0 || tval > 99 {
			return false
		}
	}
	return true
}

func IsValidDate(value string) bool {
	var tval int
	if len(value) != 6 {
		return false
	}
	if !mathutil.IsValidNumericString(value) {
		return false
	}
	timeSlice := []rune(value)
	day := string(timeSlice[0:2])
	mon := string(timeSlice[2:4])
	year := string(timeSlice[4:6])

	if mon == "01" || mon == "03" || mon == "05" || mon == "07" || mon == "08" || mon == "10" || mon == "12" {
		tval, _ = strconv.Atoi(day)
		if tval < 1 || tval > 31 {
			return false
		}
	} else if mon == "04" || mon == "06" || mon == "09" || mon == "11" {
		tval, _ = strconv.Atoi(day)
		if tval < 1 || tval > 30 {
			return false
		}
	} else {
		tval, _ = strconv.Atoi(year)
		if tval%4 == 0 {
			tval, _ = strconv.Atoi(day)
			if tval < 1 || tval > 29 {
				return false
			}
		} else {
			tval, _ = strconv.Atoi(day)
			if tval < 1 || tval > 28 {
				return false
			}
		}
	}
	tval, _ = strconv.Atoi(mon)
	if tval < 1 || tval > 12 {
		return false
	}
	tval, _ = strconv.Atoi(year)
	if tval < 0 || tval > 99 {
		return false
	}
	return true
}

func TransformDateExternal(InputDate string, DateFormat string) (int, string) {
	var OutputDate string
	dd := InputDate[:2]
	mm := InputDate[2:4]

	lDateFormat := strings.ToUpper(DateFormat)
	switch lDateFormat {
	case "MMDD":
		OutputDate = mm + dd
		break
	}
	return 1, OutputDate
}

func TransformTimeExternal(InputTime string, TimeFormat string) (int, string) {
	var OutputTime string
	hour := InputTime[:2]
	min := InputTime[2:4]
	sec := InputTime[4:6]
	lTimeFormat := strings.ToUpper(TimeFormat)
	switch lTimeFormat {
	case "HHMMSS":
		OutputTime = hour + min + sec
		break
	}
	return 1, OutputTime
}

func GetDateTimeVal() string {
	currentDate := time.Now()
	return fmt.Sprintf("%s%s", currentDate.Format("20060102"), currentDate.Format("150405"))
}
