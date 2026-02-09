package lexicalparser

import (
	"encoding/json"
	"fmt"
	"lmsapieng/include/common/lexicalparserdef"
	"lmsapieng/include/common/rejectdef"
	"lmsapieng/include/common/reqbrokerdef"
	"lmsapieng/libsrc/utils/msgutil"
	"lmsapieng/libsrc/utils/serviceutil"
	"lmsapieng/libsrc/utils/templateutil"
)

func SendToServiceFunctionType(execFunction lexicalparserdef.ExecFunctionStruct, reqBrokerDataMap map[string]interface{}) int {
	var rejectDesc string
	var reqBuffer, respBuffer []byte
	var err error
	var objectMap map[string]interface{}

	if len(execFunction.SendToService.ServiceModule) == 0 {
		//trace.Lg("ServiceModule is NULL in SendToServiceFunctionType for FunctionName[%s]", execFunction.FunctionName)
		rejectDesc = fmt.Sprintf("ServiceModule is NULL in SendToServiceFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Reject_SendToServiceFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ValidateConditionInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ValidateConditionInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	if len(execFunction.SendToService.ReqObj) == 0 {
		//trace.Lg("ReqObj is NULL in SendToServiceFunctionType for FunctionName[%s]", execFunction.FunctionName)
		rejectDesc = fmt.Sprintf("ReqObj is NULL in SendToServiceFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Reject_SendToServiceFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ValidateConditionInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ValidateConditionInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	if len(execFunction.SendToService.RespObj) == 0 {
		//trace.Lg("RespObj is NULL in SendToServiceFunctionType for FunctionName[%s]", execFunction.FunctionName)
		rejectDesc = fmt.Sprintf("RespObj is NULL in SendToServiceFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Reject_SendToServiceFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ValidateConditionInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ValidateConditionInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	if execFunction.SendToService.TimeOut == 0 {
		//trace.Lg("TimeOut is Zero in SendToServiceFunctionType for FunctionName[%s]", execFunction.FunctionName)
		rejectDesc = fmt.Sprintf("TimeOut is Zero in SendToServiceFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Reject_SendToServiceFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ValidateConditionInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ValidateConditionInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	reqBuffer, err = json.Marshal(reqBrokerDataMap[execFunction.SendToService.ReqObj])
	if err != nil {
		//trace.Lg("json.Marshal() failed for ReqObj[%s] with err[%s] in SendToServiceFunctionType for FunctionName[%s]", execFunction.SendToService.ReqObj, err, execFunction.FunctionName)
		rejectDesc = fmt.Sprintf("json.Marshal() failed for ReqObj[%s] with err[%s] in SendToServiceFunctionType for FunctionName[%s]", execFunction.SendToService.ReqObj, err, execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Reject_SendToServiceFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ValidateConditionInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ValidateConditionInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	//trace.Lg("reqBuffer[%s]", reqBuffer)
	if serviceutil.SendToService(reqBrokerDataMap, execFunction.SendToService.ServiceModule, reqBuffer, &respBuffer, execFunction.SendToService.TimeOut) < 0 {
		rejectDesc = fmt.Sprintf("SendToService() Error in SendToServiceFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Reject_SendToServiceFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ValidateConditionInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ValidateConditionInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	//trace.Lg("Send to server response[%s]", respBuffer)
	err = json.Unmarshal(respBuffer, &objectMap)
	if err != nil {
		//trace.Lg("json.Unmarshal() failed for RespObj[%s] with err[%s] in SendToServiceFunctionType for FunctionName[%s]", execFunction.SendToService.RespObj, err, execFunction.FunctionName)
		rejectDesc = fmt.Sprintf("json.Unmarshal() failed for RespObj[%s] with err[%s] in SendToServiceFunctionType for FunctionName[%s]", execFunction.SendToService.RespObj, err, execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Reject_SendToServiceFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ValidateConditionInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ValidateConditionInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	reqBrokerDataMap[execFunction.SendToService.RespObj] = objectMap
	return 1
}
