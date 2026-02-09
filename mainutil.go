package main

import (
	"context"
	"encoding/json"
	"fmt"
	"lmsapieng/include/common/dbdef"
	"lmsapieng/include/common/globaldef"
	"lmsapieng/include/common/msgdef"
	"lmsapieng/include/common/rejectdef"
	"lmsapieng/include/common/reqbrokerdef"
	"lmsapieng/libsrc/lexicaldef"
	"lmsapieng/libsrc/utils/msgutil"
	"lmsapieng/libsrc/utils/reqbrokerutil"
	"lmsapieng/libsrc/utils/schemainfo"
	"lmsapieng/libsrc/utils/trace"
)

func decomposeAPIReq(requestBody []byte, brokerReqInfo *reqbrokerdef.MsgBrokerReqInfo) int {
	var err error
	err = json.Unmarshal(requestBody, brokerReqInfo)
	if err != nil {
		//trace.Lg("decomposeAPIReq() failed with err<%s> for requestBody<%s>", err, requestBody)
		return -1
	}
	//trace.Lg("decomposeAPIReq() success with MicroserviceName<%s>", brokerReqInfo.MicroserviceName)
	return 1
}

func completeTxn(reqbrokerdatamap map[string]interface{}, respBuffer *[]byte) {
	var respInfo msgdef.RespInfoStruct
	//trace.Lg("completeTxn() called")
	if msgutil.ParseResp(*respBuffer, &respInfo) < 0 {
		schemainfo.RollbackActiveDBContext(reqbrokerdatamap)
	} else {
		schemainfo.CommitActiveDBContext(reqbrokerdatamap)
	}
}

func execMicroServices(requestBody []byte, httpReqCtx context.Context) []byte {
	var brokerReqInfo reqbrokerdef.MsgBrokerReqInfo
	var reqBrokerDataMap map[string]interface{}
	var activeDBContextArray []dbdef.DBContextDef
	var respBuffer []byte

	if decomposeAPIReq(requestBody, &brokerReqInfo) < 0 {
		return nil
	}
	reqBrokerDataMap = make(map[string]interface{})
	json.Unmarshal(brokerReqInfo.BrokerRequestData, &reqBrokerDataMap)
	activeDBContextArray = make([]dbdef.DBContextDef, 0)
	reqBrokerDataMap[dbdef.DBContextArrayObj] = activeDBContextArray
	reqBrokerDataMap[globaldef.ReqCtxObj] = httpReqCtx
	reqBrokerStorageMap := make(map[string]interface{})
	reqBrokerDataMap[reqbrokerdef.ReqBrokerStorageObj] = reqBrokerStorageMap
	defer completeTxn(reqBrokerDataMap, &respBuffer)
	respBuffer = execService(brokerReqInfo.MicroserviceName, reqBrokerDataMap)
	return respBuffer
}

func execService(microServiceName string, reqBrokerDataMap map[string]interface{}) []byte {
	var rejectDesc, correctiveAction string
	retval := lexicaldef.HandleMicroServiceReq(microServiceName, reqBrokerDataMap)
	trace.Lg("Final reqBrokerDataBuffer[%s]", reqbrokerutil.GetReqBrokerDataMapBuffer(reqBrokerDataMap))
	if retval != -2 {
		return reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj].([]byte)
	}
	rejectDesc = fmt.Sprintf("execService() failed with invalid microServiceName[%s]", microServiceName)
	correctiveAction = fmt.Sprintf("Include the microServiceName[%s]", microServiceName)
	reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_InvalidMicroService, []byte(rejectDesc), []byte(correctiveAction), rejectDesc)
	return reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj].([]byte)
}
