package jsonutil

import (
	"encoding/json"
	"fmt"
)

func GetValueFromJSONObj(JsonObj json.RawMessage, JsonKey string, JsonValue *string) int {
	JSONObjMap := make(map[string]interface{})
	err := json.Unmarshal(JsonObj, &JSONObjMap)
	if err != nil {
//trace.Lg("json.Unmarshal() failed for JsonObj(%s)", JsonObj)
		return -1
	}
	_, ok := JSONObjMap[JsonKey]
	if !ok {
//trace.Lg("No JsonKey %s", JsonKey)
		return -1
	}
	tempType := fmt.Sprintf("%T", JSONObjMap[JsonKey])
	if tempType != "string" {
//trace.Lg("%s should be a string", JsonKey)
		return -1
	}
	*JsonValue = JSONObjMap[JsonKey].(string)
	return 1
}

func GetIntValueFromJSONObj(JsonObj json.RawMessage, JsonKey string, JsonValue *float64) int {
	JSONObjMap := make(map[string]interface{})
	err := json.Unmarshal(JsonObj, &JSONObjMap)
	if err != nil {
//trace.Lg("json.Unmarshal() failed for JsonObj(%s)", JsonObj)
		return -1
	}
	_, ok := JSONObjMap[JsonKey]
	if !ok {
//trace.Lg("No JsonKey %s", JsonKey)
		return -1
	}
	tempType := fmt.Sprintf("%T", JSONObjMap[JsonKey])
//trace.Lg("tempType<%s>", tempType)
	if tempType != "float64" {
//trace.Lg("%s should be a float64", JsonKey)
		return -1
	}
	*JsonValue = JSONObjMap[JsonKey].(float64)
	return 1
}
