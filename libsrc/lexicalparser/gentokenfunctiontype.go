package lexicalparser

import (
	"fmt"
	"lmsapieng/include/common/lexicalparserdef"
	"lmsapieng/include/common/rejectdef"
	"lmsapieng/include/common/reqbrokerdef"
	"lmsapieng/libsrc/utils/lexicalparserutil"
	"lmsapieng/libsrc/utils/mathutil"
	"lmsapieng/libsrc/utils/msgutil"
	"lmsapieng/libsrc/utils/templateutil"
	"lmsapieng/libsrc/utils/tokenutil"
	"strconv"
)

func GenTokenFunctionType(execFunction lexicalparserdef.ExecFunctionStruct, reqBrokerDataMap map[string]interface{}) int {
	var rejectDesc string

	if len(execFunction.GenTokenInfo.UserID) == 0 {
		rejectDesc = fmt.Sprintf("UserID is NULL in GenTokenInfoFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GenTokenInfoFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenTokenInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenTokenInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	dataValue := lexicalparserutil.ReadValueFrmDataMap(reqBrokerDataMap, execFunction.GenTokenInfo.UserID, "string", &rejectDesc)
	if dataValue == nil {
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ReadValueErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenTokenInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenTokenInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	lUserID, _ := strconv.Atoi(dataValue.(string))

	if len(execFunction.GenTokenInfo.Email) == 0 {
		rejectDesc = fmt.Sprintf("Email is NULL in GenTokenInfoFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GenTokenInfoFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenTokenInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenTokenInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	dataValue = lexicalparserutil.ReadValueFrmDataMap(reqBrokerDataMap, execFunction.GenTokenInfo.Email, "string", &rejectDesc)
	if dataValue == nil {
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ReadValueErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenTokenInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenTokenInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	lEmail := dataValue.(string)

	if len(execFunction.GenTokenInfo.Role) == 0 {
		rejectDesc = fmt.Sprintf("Email is NULL in GenTokenInfoFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GenTokenInfoFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenTokenInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenTokenInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	dataValue = lexicalparserutil.ReadValueFrmDataMap(reqBrokerDataMap, execFunction.GenTokenInfo.Role, "string", &rejectDesc)
	if dataValue == nil {
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ReadValueErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenTokenInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenTokenInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	lRole := dataValue.(string)

	if len(execFunction.GenTokenInfo.TokenIssuer) == 0 {
		rejectDesc = fmt.Sprintf("TokenIssuer is NULL in GenTokenInfoFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GenTokenInfoFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenTokenInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenTokenInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	dataValue = lexicalparserutil.ReadValueFrmDataMap(reqBrokerDataMap, execFunction.GenTokenInfo.TokenIssuer, "string", &rejectDesc)
	if dataValue == nil {
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ReadValueErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenTokenInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenTokenInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	lTokenIssuer := dataValue.(string)

	if len(execFunction.GenTokenInfo.TokenSecret) == 0 {
		rejectDesc = fmt.Sprintf("TokenSecret is NULL in GenTokenInfoFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GenTokenInfoFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenTokenInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenTokenInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	dataValue = lexicalparserutil.ReadValueFrmDataMap(reqBrokerDataMap, execFunction.GenTokenInfo.TokenSecret, "string", &rejectDesc)
	if dataValue == nil {
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ReadValueErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenTokenInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenTokenInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	lTokenSecret := dataValue.(string)

	if len(execFunction.GenTokenInfo.TokenExpUnit) == 0 {
		rejectDesc = fmt.Sprintf("TokenExpUnit is NULL in GenTokenInfoFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GenTokenInfoFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenTokenInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenTokenInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	dataValue = lexicalparserutil.ReadValueFrmDataMap(reqBrokerDataMap, execFunction.GenTokenInfo.TokenExpUnit, "string", &rejectDesc)
	if dataValue == nil {
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ReadValueErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenTokenInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenTokenInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	lTokenExpUnit := dataValue.(string)

	if len(execFunction.GenTokenInfo.TokenExpiry) == 0 {
		rejectDesc = fmt.Sprintf("TokenExpiry is NULL in GenTokenInfoFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GenTokenInfoFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenTokenInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenTokenInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	dataValue = lexicalparserutil.ReadValueFrmDataMap(reqBrokerDataMap, execFunction.GenTokenInfo.TokenExpiry, "string", &rejectDesc)
	if dataValue == nil {
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ReadValueErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenTokenInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenTokenInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	if !mathutil.IsValidNumericString(dataValue.(string)) {
		rejectDesc = fmt.Sprintf("TokenExpiry[%s] is Invalid in GenTokenInfoFunctionType for FunctionName[%s]", dataValue.(string), execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GenTokenInfoFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenTokenInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenTokenInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	lTokenExpiry, _ := strconv.Atoi(dataValue.(string))

	if len(execFunction.GenTokenInfo.TokenObject) == 0 {
		rejectDesc = fmt.Sprintf("TokenObject is NULL in GenTokenInfoFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GenTokenInfoFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenTokenInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenTokenInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	retVal, tokenValue := tokenutil.GenToken(lUserID, lEmail, lRole, lTokenIssuer, lTokenSecret, lTokenExpUnit, lTokenExpiry)
	if retVal < 0 {
		rejectDesc = fmt.Sprintf("GenToken() failed in GenTokenInfoFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GenTokenInfoFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenTokenInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenTokenInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	objectMap := make(map[string]interface{})
	objectMap["token_value"] = tokenValue
	reqBrokerDataMap[execFunction.GenTokenInfo.TokenObject] = objectMap
	return 1
}
