package lexicalparser

import (
	"encoding/json"
	"fmt"
	"lmsapieng/include/common/lexicalparserdef"
	"lmsapieng/include/common/rejectdef"
	"lmsapieng/include/common/reqbrokerdef"
	"lmsapieng/libsrc/utils/msgutil"
	"lmsapieng/libsrc/utils/reqbrokerutil"
	"lmsapieng/libsrc/utils/templateutil"
)

func SendToBrokerFunctionType(execFunction lexicalparserdef.ExecFunctionStruct, reqBrokerDataMap map[string]interface{}) int {
	var rejectDesc string
	var reqBuffer, respBuffer []byte
	var err error
	var brokerSelectionArgs string
	var brokerInstance []reqbrokerdef.ReqBrokerInstanceStruct
	var parsedResp map[string]interface{}
	if len(execFunction.SendToBrokerInfo.ReqType) == 0 {
		//trace.Lg("ReqType is NULL in SendToBrokerFunctionType for FunctionName[%s]", execFunction.FunctionName)
		rejectDesc = fmt.Sprintf("ReqType is NULL in SendToBrokerFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_SendToBrokerFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.SendToBrokerInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.SendToBrokerInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	if len(execFunction.SendToBrokerInfo.ReqObj) == 0 {
		//trace.Lg("ReqObj is NULL in SendToBrokerFunctionType for FunctionName[%s]", execFunction.FunctionName)
		rejectDesc = fmt.Sprintf("ReqObj is NULL in SendToBrokerFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_SendToBrokerFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.SendToBrokerInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.SendToBrokerInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	if len(execFunction.SendToBrokerInfo.RespObj) == 0 {
		//trace.Lg("RespObj is NULL in SendToBrokerFunctionType for FunctionName[%s]", execFunction.FunctionName)
		rejectDesc = fmt.Sprintf("RespObj is NULL in SendToBrokerFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_SendToBrokerFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.SendToBrokerInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.SendToBrokerInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	reqBuffer, err = json.Marshal(reqBrokerDataMap[execFunction.SendToBrokerInfo.ReqObj])
	if err != nil {
		//trace.Lg("json.Marshal() failed for ReqObj[%s] with err[%s] in SendToBrokerFunctionType for FunctionName[%s]", execFunction.SendToBrokerInfo.ReqObj, err, execFunction.FunctionName)
		rejectDesc = fmt.Sprintf("json.Marshal() failed for ReqObj[%s] with err[%s] in SendToBrokerFunctionType for FunctionName[%s]", execFunction.SendToBrokerInfo.ReqObj, err, execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_SendToBrokerFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.SendToBrokerInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.SendToBrokerInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	//trace.Lg("reqBuffer [%s]",string(reqBuffer))
	var dataMap map[string]interface{}

	// Unmarshal JSON into the map
	error_resp := json.Unmarshal(reqBuffer, &dataMap)
	if error_resp != nil {
		fmt.Println("Error:", error_resp)
		return -1
	}
	//trace.Lg("DataMap", dataMap)

	brokerReqBuffer := reqbrokerutil.TranslateToBrokerMsg(execFunction.SendToBrokerInfo.ReqType, dataMap)
	//trace.Lg("reqBuffer [%s]", string(brokerReqBuffer))
	if reqbrokerutil.GetBrokerSelectionArgs(brokerReqBuffer, &brokerSelectionArgs, &rejectDesc) < 0 {
		//trace.Lg("GetBrokerSelectionArgs failed with err[%s] in SendToBrokerFunctionType for FunctionName[%s]", rejectDesc, execFunction.FunctionName)
		rejectDesc = fmt.Sprintf("GetBrokerSelectionArgs failed with err[%s] in SendToBrokerFunctionType for FunctionName[%s]", rejectDesc, execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_SendToBrokerFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.SendToBrokerInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.SendToBrokerInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}

	brokerInstance = make([]reqbrokerdef.ReqBrokerInstanceStruct, 0)
	if reqbrokerutil.GetBrokerInstance(brokerSelectionArgs, &brokerInstance) < 0 {
		//trace.Lg("GetBrokerInstance failed in SendToBrokerFunctionType for FunctionName[%s]", execFunction.FunctionName)
		rejectDesc = fmt.Sprintf("GetBrokerInstance failed in SendToBrokerFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_SendToBrokerFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.SendToBrokerInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.SendToBrokerInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	msgSuccess := false
	for i := 0; i < len(brokerInstance); i++ {
		if msgutil.PostReq(brokerInstance[i].BrokerModule, brokerInstance[i].IpAddr, brokerInstance[i].PortNum, brokerInstance[i].TimeOut, brokerReqBuffer, &respBuffer) < 0 {
			continue
		}
		msgSuccess = true
	}
	if !msgSuccess {
		//trace.Lg("Send To Broker Failed failed in SendToBrokerFunctionType for FunctionName[%s]", execFunction.FunctionName)
		rejectDesc = fmt.Sprintf("Send To Broker Failed failed in SendToBrokerFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_SendToBrokerFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.SendToBrokerInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.SendToBrokerInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	//trace.Lg("respBuffer[%s]", string(respBuffer))

	// reqBrokerDataMap[execFunction.SendToBrokerInfo.RespObj] = string(respBuffer)

	// err = json.Unmarshal(respBuffer, reqBrokerDataMap[execFunction.SendToBrokerInfo.RespObj])

	// if err != nil {
	// 	////trace.Lg("json.Unmarshal() failed for RespObj[%s] with err[%s] in SendToBrokerFunctionType for FunctionName[%s]", execFunction.SendToBrokerInfo.RespObj, err, execFunction.FunctionName)
	// 	rejectDesc = fmt.Sprintf("json.Unmarshal() failed for RespObj[%s] with err[%s] in SendToBrokerFunctionType for FunctionName[%s]", execFunction.SendToBrokerInfo.RespObj, err, execFunction.FunctionName)
	// 	reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_SendToBrokerFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.SendToBrokerInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.SendToBrokerInfo.AppCorrectiveAction), rejectDesc)
	// 	return -1
	// }

	error_response := json.Unmarshal(respBuffer, &parsedResp)
	if error_response != nil {
		//trace.Lg("json.Unmarshal() failed: %s", error_response)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_SendToBrokerFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.SendToBrokerInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.SendToBrokerInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}

	reqBrokerDataMap[execFunction.SendToBrokerInfo.RespObj] = parsedResp
	return 1
}
