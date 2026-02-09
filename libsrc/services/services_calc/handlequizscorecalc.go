package services_calc

import (
	"lmsapieng/libsrc/utils/trace"
	"math"
	"strconv"
	"strings"
)

func HandleQuizScoreCalc(reqBrokerDataMap map[string]interface{}) int {

	rawQuizObj, exists := reqBrokerDataMap["quiz_obj"]
	if !exists || rawQuizObj == nil {
		trace.Lg("QuizScoreCalc | quiz_obj not found")
		return 1
	}

	// helper: convert interface{} → int
	toInt := func(v interface{}) int {
		switch t := v.(type) {
		case int:
			return t
		case int64:
			return int(t)
		case float64:
			return int(t)
		case string:
			i, _ := strconv.Atoi(strings.TrimSpace(t))
			return i
		default:
			return 0
		}
	}

	obtainedPoints := 0
	maxTotalPoints := 0
	questionCount := 0
	passingScore := 0

	switch arr := rawQuizObj.(type) {

	case []interface{}:
		for i, item := range arr {
			m, ok := item.(map[string]interface{})
			if !ok {
				continue
			}

			points := toInt(m["points"])
			total := toInt(m["total_points"])
			ps := toInt(m["passing_score"])

			trace.Lg(
				"QuizScoreCalc | index=%d | points=%d | total_points=%d",
				i, points, total,
			)

			obtainedPoints += points
			maxTotalPoints += total
			questionCount++

			if passingScore == 0 {
				passingScore = ps // same for all questions
			}
		}

	case []map[string]interface{}:
		for i, m := range arr {
			points := toInt(m["points"])
			total := toInt(m["total_points"])
			ps := toInt(m["passing_score"])

			trace.Lg(
				"QuizScoreCalc | index=%d | points=%d | total_points=%d",
				i, points, total,
			)

			obtainedPoints += points
			maxTotalPoints += total
			questionCount++

			if passingScore == 0 {
				passingScore = ps
			}
		}

	default:
		trace.Lg("QuizScoreCalc | unsupported quiz_obj type=%T", rawQuizObj)
		return 1
	}

	// ---- CALCULATE AVERAGE ----

	average := 0.0
	if maxTotalPoints > 0 {
		rawAvg := (float64(obtainedPoints) / float64(maxTotalPoints)) * 100
		average = math.Round(rawAvg)
	}

	// ---- PASS / FAIL ----
	isCompleted := average >= float64(passingScore)

	// ---- STORE SUMMARY ----
	reqBrokerDataMap["quiz_obj_summary"] = map[string]interface{}{
		"obtained_points": obtainedPoints,
		"total_points":    maxTotalPoints,
		"question_count":  questionCount,
		"average":         average,      // float64
		"passing_score":   passingScore, // int
		"is_completed":    isCompleted,
	}

	trace.Lg(
		"QuizScoreCalc | FINAL | obtained=%d total=%d questions=%d average=%.2f passing=%d completed=%v",
		obtainedPoints,
		maxTotalPoints,
		questionCount,
		average,
		passingScore,
		isCompleted,
	)

	return 1
}
