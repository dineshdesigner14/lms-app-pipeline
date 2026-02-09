package lexicalparser

import (
	"encoding/json"
	"fmt"
	"lmsapieng/include/common/lexicalparserdef"
	"lmsapieng/include/common/msgdef"
	"lmsapieng/include/common/rejectdef"
	"lmsapieng/include/common/reqbrokerdef"
	"lmsapieng/libsrc/utils/lexicalparserutil"
	"lmsapieng/libsrc/utils/msgutil"
	"lmsapieng/libsrc/utils/templateutil"
)

func ComposeRespFunctionType(execFunction lexicalparserdef.ExecFunctionStruct, reqBrokerDataMap map[string]interface{}) int {
	var rejectDesc string

	// trace.Lg("reqBrokerDataMapBuffer[%s]", reqbrokerutil.GetReqBrokerDataMapBuffer(reqBrokerDataMap))

	if len(execFunction.ResponseArray.RespObj) == 1 && len(execFunction.ResponseArray.RespObj[0].ObjReq) != 0 && execFunction.ResponseArray.RespObj[0].ObjReq == "1" {
		payloadMap := make(map[string]interface{})
		if len(execFunction.ResponseArray.RespObj[0].DataSource) == 0 {
			rejectDesc = fmt.Sprintf("DataSource is NULL in ComposeRespFunctionType for FunctionName[%s]", execFunction.FunctionName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_DBSingleReadFTFilterDataErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppCorrectiveAction), rejectDesc)
			return -1
		}
		if len(execFunction.ResponseArray.RespObj[0].ObjName) == 0 {
			rejectDesc = fmt.Sprintf("ObjName is NULL in ComposeRespFunctionType for FunctionName[%s]", execFunction.FunctionName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_DBSingleReadFTFilterDataErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppCorrectiveAction), rejectDesc)
			return -1
		}
		if len(execFunction.ResponseArray.RespObj[0].ObjType) == 0 {
			rejectDesc = fmt.Sprintf("ObjType is NULL in ComposeRespFunctionType for FunctionName[%s]", execFunction.FunctionName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_DBSingleReadFTFilterDataErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppCorrectiveAction), rejectDesc)
			return -1
		}
		dataValue := lexicalparserutil.ReadValueFrmDataMap(reqBrokerDataMap, execFunction.ResponseArray.RespObj[0].DataSource, execFunction.ResponseArray.RespObj[0].ObjType, &rejectDesc)
		if dataValue == nil {
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ReadValueErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppCorrectiveAction), rejectDesc)
			return -1
		}
		payloadMap[execFunction.ResponseArray.RespObj[0].ObjName] = dataValue
		pay_load_buffer, err := json.MarshalIndent(&payloadMap, "", "\t")
		if err != nil {
			rejectDesc = fmt.Sprintf("JsonMarshal Err in ComposeRespFunctionType for FunctionName[%s]", execFunction.FunctionName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ReadValueErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppCorrectiveAction), rejectDesc)
			return -1
		}
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetRespWithData(msgdef.RCapproved, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppCorrectiveAction), string(pay_load_buffer))
		return 1
	}
	if len(execFunction.ResponseArray.RespObj) == 1 {
		if len(execFunction.ResponseArray.RespObj[0].DataSource) == 0 {
			rejectDesc = fmt.Sprintf("DataSource is NULL in ComposeRespFunctionType for FunctionName[%s]", execFunction.FunctionName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_DBSingleReadFTFilterDataErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppCorrectiveAction), rejectDesc)
			return -1
		}
		if len(execFunction.ResponseArray.RespObj[0].ObjName) == 0 {
			rejectDesc = fmt.Sprintf("ObjName is NULL in ComposeRespFunctionType for FunctionName[%s]", execFunction.FunctionName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_DBSingleReadFTFilterDataErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppCorrectiveAction), rejectDesc)
			return -1
		}
		if len(execFunction.ResponseArray.RespObj[0].ObjType) == 0 {
			rejectDesc = fmt.Sprintf("ObjType is NULL in ComposeRespFunctionType for FunctionName[%s]", execFunction.FunctionName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_DBSingleReadFTFilterDataErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppCorrectiveAction), rejectDesc)
			return -1
		}
		dataValue := lexicalparserutil.ReadValueFrmDataMap(reqBrokerDataMap, execFunction.ResponseArray.RespObj[0].DataSource, execFunction.ResponseArray.RespObj[0].ObjType, &rejectDesc)
		if dataValue == nil {
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ReadValueErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppCorrectiveAction), rejectDesc)
			return -1
		}
		pay_load_buffer, err := json.MarshalIndent(&dataValue, "", "\t")
		if err != nil {
			rejectDesc = fmt.Sprintf("JsonMarshal Err in ComposeRespFunctionType for FunctionName[%s]", execFunction.FunctionName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ReadValueErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppCorrectiveAction), rejectDesc)
			return -1
		}
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetRespWithData(msgdef.RCapproved, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppCorrectiveAction), string(pay_load_buffer))
	} else {
		payloadMap := make(map[string]interface{})
		for i := 0; i < len(execFunction.ResponseArray.RespObj); i++ {
			if len(execFunction.ResponseArray.RespObj[i].DataSource) == 0 {
				rejectDesc = fmt.Sprintf("DataSource is NULL in ComposeRespFunctionType for FunctionName[%s]", execFunction.FunctionName)
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_DBSingleReadFTFilterDataErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppCorrectiveAction), rejectDesc)
				return -1
			}
			if len(execFunction.ResponseArray.RespObj[i].ObjName) == 0 {
				rejectDesc = fmt.Sprintf("ObjName is NULL in ComposeRespFunctionType for FunctionName[%s]", execFunction.FunctionName)
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_DBSingleReadFTFilterDataErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppCorrectiveAction), rejectDesc)
				return -1
			}
			if len(execFunction.ResponseArray.RespObj[i].ObjType) == 0 {
				rejectDesc = fmt.Sprintf("ObjType is NULL in ComposeRespFunctionType for FunctionName[%s]", execFunction.FunctionName)
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_DBSingleReadFTFilterDataErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppCorrectiveAction), rejectDesc)
				return -1
			}
			dataValue := lexicalparserutil.ReadValueFrmDataMap(reqBrokerDataMap, execFunction.ResponseArray.RespObj[i].DataSource, execFunction.ResponseArray.RespObj[i].ObjType, &rejectDesc)
			if dataValue == nil {
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ReadValueErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppCorrectiveAction), rejectDesc)
				return -1
			}
			payloadMap[execFunction.ResponseArray.RespObj[i].ObjName] = dataValue
		}
		pay_load_buffer, err := json.MarshalIndent(&payloadMap, "", "\t")
		if err != nil {
			rejectDesc = fmt.Sprintf("JsonMarshal Err in ComposeRespFunctionType for FunctionName[%s]", execFunction.FunctionName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ReadValueErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppCorrectiveAction), rejectDesc)
			return -1
		}
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetRespWithData(msgdef.RCapproved, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppCorrectiveAction), string(pay_load_buffer))
	}
	return 1
}
