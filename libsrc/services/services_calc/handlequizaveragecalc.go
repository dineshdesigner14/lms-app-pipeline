package services_calc

import (
	"fmt"
	"lmsapieng/libsrc/utils/trace"
	"reflect"
	"strconv"
)

func HandleQuizAverageCalc(reqBrokerDataMap map[string]interface{}) int {

	trace.Lg("HandleQuizAverageCalc() called")

	quizObj, ok := reqBrokerDataMap["quiz_obj"].(map[string]interface{})
	if !ok {
		trace.Lg("ERROR: quiz_obj not found in reqBrokerDataMap")
		trace.Lg("reqBrokerDataMap keys: %+v", reflect.ValueOf(reqBrokerDataMap).MapKeys())
		return -1
	}

	trace.Lg("quiz_obj BEFORE calculation: %+v", quizObj)

	totalStr, ok1 := quizObj["total_points"].(string)
	obtainedStr, ok2 := quizObj["obtained_points"].(string)

	if !ok1 || !ok2 {
		trace.Lg(
			"ERROR: Invalid types | total_points=%T obtained_points=%T",
			quizObj["total_points"],
			quizObj["obtained_points"],
		)
		return -1
	}

	total, err1 := strconv.Atoi(totalStr)
	obtained, err2 := strconv.Atoi(obtainedStr)

	if err1 != nil || err2 != nil {
		trace.Lg(
			"ERROR: Atoi failed | total='%s' obtained='%s' err1=%v err2=%v",
			totalStr, obtainedStr, err1, err2,
		)
		return -1
	}

	trace.Lg("Parsed values | total=%d obtained=%d", total, obtained)

	if total == 0 {
		trace.Lg("Total points is 0 → setting average to 0")
		quizObj["average"] = "0"
		trace.Lg("quiz_obj AFTER calculation: %+v", quizObj)
		return 1
	}

	avg := (float64(obtained) / float64(total)) * 100
	avgStr := fmt.Sprintf("%.2f", avg)

	quizObj["average"] = avgStr

	trace.Lg("Average calculated successfully: %s%%", avgStr)
	trace.Lg("quiz_obj AFTER calculation: %+v", quizObj)

	return 1
}
