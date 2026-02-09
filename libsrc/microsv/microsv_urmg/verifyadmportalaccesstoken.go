package microsv_urmg

import (
	"encoding/json"
	"fmt"
	"lmsapieng/include/common/dbdef"
	"lmsapieng/include/common/globaldef"
	"lmsapieng/include/common/httpdef"
	"lmsapieng/include/common/moduledef"
	"lmsapieng/include/common/msgdef"
	"lmsapieng/include/common/rejectdef"
	"lmsapieng/include/common/reqbrokerdef"
	"lmsapieng/include/common/urmgdef"
	"lmsapieng/libsrc/utils/msgutil"
)

func getVerifyAdmPortalAccessTokenBuffer(reqBrokerDataMap map[string]interface{}, MicroServiceName string) []byte {
	var brokerReqInfo reqbrokerdef.MsgBrokerReqInfo
	//trace.Lg("VerifyAdmPortalAccessToken() called")
	microServiceBufferMap := make(map[string]interface{})
	for reqBrokerKey, reqBrokerValue := range reqBrokerDataMap {
		if reqBrokerKey != dbdef.DBContextArrayObj &&
			reqBrokerKey != reqbrokerdef.ReqBrokerServiceConfigObj &&
			reqBrokerKey != globaldef.ReqCtxObj &&
			reqBrokerKey != reqbrokerdef.ReqBrokerDataMapRespObj {
			microServiceBufferMap[reqBrokerKey] = reqBrokerValue
		}
	}
	microServiceBuffer, _ := json.Marshal(&microServiceBufferMap)
	brokerReqInfo.MicroserviceName = MicroServiceName
	brokerReqInfo.BrokerRequestData = microServiceBuffer
	brokerReqInfoBuffer, _ := json.Marshal(&brokerReqInfo)
	return brokerReqInfoBuffer
}

func writeRespData(respInfo msgdef.RespInfoStruct, reqBrokerDataMap map[string]interface{}) {
	respDataMap := make(map[string]interface{})
	//trace.Lg("writeRespData() called for respData[%s]", respInfo.RespInfo.RespData)
	err := json.Unmarshal(respInfo.RespInfo.RespData, &respDataMap)
	if err != nil {
		//trace.Lg("json.Unmarshal() failed for RespData[%s]", respInfo.RespInfo.RespData)
		return
	}
	_, ok := respDataMap[reqbrokerdef.ReqBrokerAccumDataJSONObj]
	if !ok {
		//trace.Lg("no json obj %s in the respData..hence nothing to save", reqbrokerdef.ReqBrokerAccumDataJSONObj)
		return
	}
	tempType := fmt.Sprintf("%T", respDataMap[reqbrokerdef.ReqBrokerAccumDataJSONObj])
	if tempType != "map[string]interface {}" {
		//trace.Lg("%s object has invalid type..hence nothing to save", reqbrokerdef.ReqBrokerAccumDataJSONObj)
		return
	}
	accumDataMap := respDataMap[reqbrokerdef.ReqBrokerAccumDataJSONObj].(map[string]interface{})
	for reqBrokerKey, reqBrokerValue := range accumDataMap {
		if reqBrokerKey == urmgdef.UserDataTableKey {
			reqBrokerDataMap[reqBrokerKey] = reqBrokerValue
		}
	}
}

func VerifyAdmPortalAccessToken(reqBrokerDataMap map[string]interface{}) int {
	var urmgApiEngAddr httpdef.HttpServerAddr
	var respData []byte
	var respInfo msgdef.RespInfoStruct
	var rejectDesc, correctiveAction string
	verifyTokenBuffer := getVerifyAdmPortalAccessTokenBuffer(reqBrokerDataMap, urmgdef.VerifyTokenMicroService)
	if locateUrmgApiEngAddr(&urmgApiEngAddr) < 0 {
		//trace.Lg("locateUrmgApiEngAddr() failed")
		rejectDesc = "locateUrmgApiEngAddr() failed"
		correctiveAction = "Check the locateUrmgApiEngAddr() failed"
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_LocateUrmgApiEngAddr, []byte(rejectDesc), []byte(correctiveAction))
		return -1
	}
	rval := msgutil.PostReq(moduledef.URMGApiEngModule, urmgApiEngAddr.ServerIpAddr, urmgApiEngAddr.ServerPort, urmgApiEngAddr.ServerTimeout, verifyTokenBuffer, &respData)
	if rval < 0 {
		if rval == -2 {
			rejectDesc = "SendToUrmgApiEngTimedOut"
			correctiveAction = "Check the SendToUrmgApiEngTimedOut"
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_SendToUrmgApiEngTimedOut, []byte(rejectDesc), []byte(correctiveAction))
			return -1
		}
		rejectDesc = "SendToUrmgApiEngFailed"
		correctiveAction = "Check the SendToUrmgApiEngFailed"
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_SendToUrmgApiEngFailed, []byte(rejectDesc), []byte(correctiveAction))
		return -1
	}
	if msgutil.ParseResp(respData, &respInfo) < 0 {
		rejectDesc = "Session expired or accessed from another device"
		correctiveAction = "Please log in again to proceed"
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_VerifyAdmPortalAccessTokenFailed, []byte(rejectDesc), []byte(correctiveAction))
		return -1
	}
	writeRespData(respInfo, reqBrokerDataMap)
	return 1
}
