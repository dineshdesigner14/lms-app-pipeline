package services_option

import (
	"lmsapieng/libsrc/utils/trace"
	"math/rand"
	"time"
)

func HandleRandomQuestionShuffle(reqBrokerDataMap map[string]interface{}) int {
	rawObj, exists := reqBrokerDataMap["questions"]
	if !exists || rawObj == nil {
		trace.Lg("RandomQuestion | questions not found")
		return 1
	}

	// default count = 10
	requiredCount := 10
	if v, ok := reqBrokerDataMap["random_question_count"]; ok {
		switch t := v.(type) {
		case int:
			requiredCount = t
		case int64:
			requiredCount = int(t)
		case float64:
			requiredCount = int(t)
		}
	}

	var questions []map[string]interface{}

	switch arr := rawObj.(type) {

	case []interface{}:
		for _, item := range arr {
			if q, ok := item.(map[string]interface{}); ok {
				questions = append(questions, q)
			}
		}

	case []map[string]interface{}:
		questions = arr

	default:
		trace.Lg("RandomQuestion | unsupported questions type=%T", rawObj)
		return 1
	}

	if len(questions) == 0 {
		trace.Lg("RandomQuestion | empty question list")
		reqBrokerDataMap["question_bank_list"] = []interface{}{}
		return 1
	}

	if requiredCount > len(questions) {
		requiredCount = len(questions)
	}

	// ---- SHUFFLE QUESTIONS ----
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(questions), func(i, j int) {
		questions[i], questions[j] = questions[j], questions[i]
	})

	selected := questions[:requiredCount]

	// ---- SHUFFLE OPTIONS ----
	for _, q := range selected {

		rawOpts, ok := q["options"]
		if !ok || rawOpts == nil {
			continue
		}

		switch opts := rawOpts.(type) {

		case []interface{}:
			if len(opts) > 1 {
				rand.Shuffle(len(opts), func(i, j int) {
					opts[i], opts[j] = opts[j], opts[i]
				})
			}
			q["options"] = opts

		case []map[string]interface{}:
			if len(opts) > 1 {
				rand.Shuffle(len(opts), func(i, j int) {
					opts[i], opts[j] = opts[j], opts[i]
				})
			}
			q["options"] = opts
		}
	}

	reqBrokerDataMap["question_bank_list"] = selected

	trace.Lg(
		"RandomQuestion | total=%d selected=%d",
		len(questions),
		len(selected),
	)

	return 1
}
