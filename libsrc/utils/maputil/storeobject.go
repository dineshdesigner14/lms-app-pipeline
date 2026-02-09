package maputil

import (
	"lmsapieng/libsrc/utils/lexicalparserutil"
	"strings"
)

func isObjectArray(objectName string) bool {
	if strings.Contains(objectName, "[") && strings.Contains(objectName, "]") {
		return true
	}
	return false
}

func StoreObject(reqBrokerDataMap map[string]interface{}, objectName string, storeData interface{}) int {
	var newobjectName string
	//trace.Lg("StoreObject() called for objectName[%s]", objectName)
	objectSlice := strings.Split(objectName, ".")
	if len(objectSlice) == 0 {
		reqBrokerDataMap[objectName] = storeData
		return 1
	}
	if !isObjectArray(objectName) {
		//trace.Lg("objectName[%s] not an array..hence directly calling SetValueFromString", objectName)
		return SetValueFromString(reqBrokerDataMap, objectName, storeData)
	}
	//trace.Lg("objectName[%s] is an array..hence call ReplaceArrayIndex followed by SetValueFromString", objectName)
	if lexicalparserutil.ReplaceArrayIndex(reqBrokerDataMap, objectName, &newobjectName) < 0 {
		return -1
	}
	return SetValueFromString(reqBrokerDataMap, newobjectName, storeData)
}
