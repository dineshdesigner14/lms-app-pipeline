package maputil

func CopyStringFromMap(dataMap map[string]interface{}, mapIndex string, destValue *string) int {
	var ok bool
	if *destValue, ok = dataMap[mapIndex].(string); ok {
		return 1
	}
	return -1
}

func CopyIntFromMap(dataMap map[string]interface{}, mapIndex string, destValue *int) int {
	var lValue int64
	var ok bool
	if lValue, ok = dataMap[mapIndex].(int64); ok {
		return 1
	}
	*destValue = int(lValue)
	return -1
}
