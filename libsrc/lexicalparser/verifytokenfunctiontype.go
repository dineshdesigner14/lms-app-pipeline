package lexicalparser

import (
	"fmt"
	"lmsapieng/include/common/lexicalparserdef"
	"lmsapieng/include/common/rejectdef"
	"lmsapieng/include/common/reqbrokerdef"
	"lmsapieng/libsrc/utils/lexicalparserutil"
	"lmsapieng/libsrc/utils/msgutil"
	"lmsapieng/libsrc/utils/templateutil"
	"lmsapieng/libsrc/utils/tokenutil"
)

func VerifyTokenFunctionType(execFunction lexicalparserdef.ExecFunctionStruct, reqBrokerDataMap map[string]interface{}) int {
	var rejectDesc string

	if len(execFunction.VerifyTokenInfo.TokenValue) == 0 {
		rejectDesc = fmt.Sprintf("TokenValue is NULL in VerifyTokenInfoFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_VerifyTokenInfoFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.VerifyTokenInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.VerifyTokenInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	dataValue := lexicalparserutil.ReadValueFrmDataMap(reqBrokerDataMap, execFunction.VerifyTokenInfo.TokenValue, "string", &rejectDesc)
	if dataValue == nil {
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ReadValueErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.VerifyTokenInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.VerifyTokenInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	lTokenValue := dataValue.(string)

	if len(execFunction.VerifyTokenInfo.TokenSecret) == 0 {
		rejectDesc = fmt.Sprintf("TokenSecret is NULL in VerifyTokenInfoFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_VerifyTokenInfoFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.VerifyTokenInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.VerifyTokenInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	dataValue = lexicalparserutil.ReadValueFrmDataMap(reqBrokerDataMap, execFunction.VerifyTokenInfo.TokenSecret, "string", &rejectDesc)
	if dataValue == nil {
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ReadValueErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.VerifyTokenInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.VerifyTokenInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	lTokenSecret := dataValue.(string)

	if len(execFunction.VerifyTokenInfo.TokenObject) == 0 {
		rejectDesc = fmt.Sprintf("TokenObject is NULL in VerifyTokenInfoFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_VerifyTokenInfoFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.VerifyTokenInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.VerifyTokenInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	lUserID := 0
	lEmail := ""
	lRole := ""
	lTokenIssuer := ""
	retVal := tokenutil.VerifyToken(lTokenValue, lTokenSecret, &lUserID, &lEmail, &lRole, &lTokenIssuer, &rejectDesc)
	if retVal < 0 {
		rejectDesc = fmt.Sprintf("VerifyToken() failed with rejectDesc[%s] in VerifyTokenInfoFunctionType for FunctionName[%s]", rejectDesc, execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_VerifyTokenInfoFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.VerifyTokenInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.VerifyTokenInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	objectMap := make(map[string]interface{})
	objectMap["user_id"] = lUserID
	objectMap["email"] = lEmail
	objectMap["role"] = lRole
	objectMap["token_issuer"] = lTokenIssuer
	reqBrokerDataMap[execFunction.VerifyTokenInfo.TokenObject] = objectMap
	return 1
}
