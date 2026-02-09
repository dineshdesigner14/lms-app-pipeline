package reqbrokerutil

import (
	"encoding/json"
	"fmt"
	"lmsapieng/include/common/globaldef"
	"lmsapieng/include/common/reqbrokerdef"
)

func GetReqBrokerDataMapBuffer(reqBrokerDataMap map[string]interface{}) string {
	pay_load_buffer, _ := json.MarshalIndent(&reqBrokerDataMap, "", "\t")
	return string(pay_load_buffer)
}

func GetDataMapBuffer(dataMap map[string]interface{}) string {
	pay_load_buffer, _ := json.MarshalIndent(&dataMap, "", "\t")
	return string(pay_load_buffer)
}

func GetReqID(reqBrokerDataMap map[string]interface{}) string {
	_, ok := reqBrokerDataMap[reqbrokerdef.ReqBrokerReqIDJSONObj]
	if !ok {
		return globaldef.NOT_INITIALIZED
	}
	tempType := fmt.Sprintf("%T", reqBrokerDataMap[reqbrokerdef.ReqBrokerReqIDJSONObj])
	if tempType != "string" {
		return globaldef.NOT_INITIALIZED
	}
	return reqBrokerDataMap[reqbrokerdef.ReqBrokerReqIDJSONObj].(string)
}
