package maputil

import (
	"strconv"
	"strings"
)

func CreateMapArray(reqBrokerDataMap map[string]interface{}, inputStr string, arraySize int) {
	keys := strings.Split(inputStr, ".")
	currentMap := reqBrokerDataMap

	for i, key := range keys {
		if i == len(keys)-1 {
			// Last key, create an array of specified size
			arr := make([]map[string]interface{}, arraySize)
			for j := range arr {
				arr[j] = make(map[string]interface{})
			}
			currentMap[key] = arr
		} else {
			// Intermediate keys, create nested maps if they don't exist
			if _, ok := currentMap[key]; !ok {
				currentMap[key] = make(map[string]interface{})
			}
			currentMap = currentMap[key].(map[string]interface{})
		}
	}
}

func SetValueFromString(reqBrokerDataMap map[string]interface{}, inputStr string, value interface{}) int {
	inputSlice := strings.Split(inputStr, ".")
	currentMap := reqBrokerDataMap
	for i := 0; i < len(inputSlice)-1; i++ {
		key := inputSlice[i]
		if strings.Contains(key, "[") {
			objectName := strings.Split(key, "[")[0]
			indexStr := strings.Split(strings.Split(key, "[")[1], "]")[0]
			index, err := strconv.Atoi(indexStr)
			if err != nil {
				return -1
			}
			if currentMap[objectName] == nil {
				currentMap[objectName] = make([]interface{}, index+1)
			}
			if arr, ok := currentMap[objectName].([]interface{}); ok {
				if len(arr) <= index {
					arr = append(arr, make(map[string]interface{}))
				}
				currentMap[objectName] = arr
				currentMap = arr[index].(map[string]interface{})
			}
			if marr, ok := currentMap[objectName].([]map[string]interface{}); ok {
				if len(marr) <= index {
					marr = append(marr, make(map[string]interface{}))
				}
				currentMap[objectName] = marr
				currentMap = marr[index]
			}
		} else {
			if currentMap[key] == nil {
				currentMap[key] = make(map[string]interface{})
			}
			currentMap = currentMap[key].(map[string]interface{})
		}
	}
	lastKey := inputSlice[len(inputSlice)-1]
	currentMap[lastKey] = value
	return 1
}

func GetInt(data map[string]interface{}, key string, defaultVal int) int {

	// Key does not exist
	v, ok := data[key]
	if !ok || v == nil {
		return defaultVal
	}

	switch val := v.(type) {

	case int:
		return val

	case int8:
		return int(val)

	case int16:
		return int(val)

	case int32:
		return int(val)

	case int64:
		return int(val)

	case float32:
		return int(val)

	case float64:
		return int(val)

	case string:
		parsed, err := strconv.Atoi(val)
		if err == nil {
			return parsed
		}
		return defaultVal

	default:
		return defaultVal
	}
}
