package services_option

func BuildHasCorrectOptionFlag(reqBrokerDataMap map[string]interface{}) int {

	// initialize once
	if _, ok := reqBrokerDataMap["has_correct_option"]; !ok {
		reqBrokerDataMap["has_correct_option"] = false
	}

	optRaw, ok := reqBrokerDataMap["lms_question_option_table_obj"]
	if !ok {
		return 1
	}

	opt, ok := optRaw.(map[string]interface{})
	if !ok {
		return 1
	}

	// flip flag if any option is correct
	if val, exists := opt["is_correct"]; exists {
		if v, ok := val.(bool); ok && v {
			reqBrokerDataMap["has_correct_option"] = true
		}
	}

	return 1
}
