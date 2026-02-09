package lexicalparserutil

import (
	"fmt"
	"lmsapieng/libsrc/utils/datatypeutil"
)

func GetExecGroupIndex(reqBrokerDataMap map[string]interface{}) int {
	_, ok := reqBrokerDataMap["ExecGroupIndex"]
	if !ok {
		//trace.Lg("IndexKey[%s] search error", "ExecGroupIndex")
		return -1
	}
	return reqBrokerDataMap["ExecGroupIndex"].(int)
}

func SetExecGroupIndex(reqBrokerDataMap map[string]interface{}, indexValue int) int {
	_, ok := reqBrokerDataMap["ExecGroupIndex"]
	if !ok {
		//trace.Lg("IndexKey[%s] search error", "ExecGroupIndex")
		return -1
	}
	// trace.Lg("SetExecGroupIndex() called with CurrentValue[%d] to NewValue[%d]", reqBrokerDataMap["ExecGroupIndex"].(int), indexValue)
	reqBrokerDataMap["ExecGroupIndex"] = indexValue
	return 1
}

func StoreIndexValueForStartLoop(reqBrokerDataMap map[string]interface{}, indexName string, startIndex int, endIndex int) int {
	execGroupIndex := GetExecGroupIndex(reqBrokerDataMap)
	if execGroupIndex < 0 {
		return -1
	}
	//trace.Lg("StoreIndexValueForStartLoop() execGroupIndex[%d] for indexName[%s]", execGroupIndex, indexName)

	indexKey := fmt.Sprintf("%s_%s", "IndexKey", indexName)
	reqBrokerDataMap[indexKey] = startIndex
	startIndexKey := fmt.Sprintf("%s_%s", "StartIndex", indexName)
	reqBrokerDataMap[startIndexKey] = startIndex
	endIndexKey := fmt.Sprintf("%s_%s", "EndIndex", indexName)
	reqBrokerDataMap[endIndexKey] = endIndex
	startLoopIndexKey := fmt.Sprintf("%s_%s", "StartLoopIndexKey", indexName)
	reqBrokerDataMap[startLoopIndexKey] = execGroupIndex

	return 1
}

func StoreIndexValueForEndLoop(reqBrokerDataMap map[string]interface{}, indexName string) int {
	execGroupIndex := GetExecGroupIndex(reqBrokerDataMap)
	if execGroupIndex < 0 {
		return -1
	}
	indexKey := fmt.Sprintf("%s_%s", "IndexKey", indexName)
	_, ok := reqBrokerDataMap[indexKey]
	if !ok {
		//trace.Lg("IndexKey[%s] search error", indexKey)
		return -1
	}
	startIndexKey := fmt.Sprintf("%s_%s", "StartIndex", indexName)
	_, ok = reqBrokerDataMap[startIndexKey]
	if !ok {
		//trace.Lg("IndexKey[%s] search error", startIndexKey)
		return -1
	}
	endIndexKey := fmt.Sprintf("%s_%s", "EndIndex", indexName)
	_, ok = reqBrokerDataMap[endIndexKey]
	if !ok {
		//trace.Lg("IndexKey[%s] search error", endIndexKey)
		return -1
	}
	startLoopIndexKey := fmt.Sprintf("%s_%s", "StartLoopIndexKey", indexName)
	_, ok = reqBrokerDataMap[startLoopIndexKey]
	if !ok {
		//trace.Lg("IndexKey[%s] search error", startLoopIndexKey)
		return -1
	}
	indexKeyValue := reqBrokerDataMap[indexKey].(int)
	endIndexKeyValue := reqBrokerDataMap[endIndexKey].(int)
	startLoopIndexKeyValue := reqBrokerDataMap[startLoopIndexKey].(int)
	//trace.Lg("indexKeyValue[%d] endIndexKeyValue[%d]", indexKeyValue, endIndexKeyValue)
	if indexKeyValue >= endIndexKeyValue-1 {
		/*if SetExecGroupIndex(reqBrokerDataMap, execGroupIndex+1) < 0 {
			return -1
		}*/
		delete(reqBrokerDataMap, indexKey)
		delete(reqBrokerDataMap, startIndexKey)
		delete(reqBrokerDataMap, endIndexKey)
		delete(reqBrokerDataMap, startLoopIndexKey)
	} else {
		indexKeyValue += 1
		reqBrokerDataMap[indexKey] = indexKeyValue
		if SetExecGroupIndex(reqBrokerDataMap, startLoopIndexKeyValue) < 0 {
			return -1
		}
	}
	return 1
}

func GetIndexValue(reqBrokerDataMap map[string]interface{}, indexName string, indexValue *int) int {
	indexKey := fmt.Sprintf("%s_%s", "IndexKey", indexName)
	_, ok := reqBrokerDataMap[indexKey]
	if !ok {
		//trace.Lg("IndexKey[%s] search error", indexKey)
		return -1
	}
	if !datatypeutil.IsInt(reqBrokerDataMap[indexKey]) {
		//trace.Lg("IndexKey[%s] is not a integer", indexKey)
		return -1
	}
	*indexValue = reqBrokerDataMap[indexKey].(int)
	return 1
}
