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

func DecodeTokenFunctionType(execFunction lexicalparserdef.ExecFunctionStruct, reqBrokerDataMap map[string]interface{}) int {
	var rejectDesc string

	if len(execFunction.DecodeTokenInfo.TokenValue) == 0 {
		rejectDesc = fmt.Sprintf("TokenValue is NULL in DecodeTokenFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_DecodeTokenInfoFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DecodeTokenInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DecodeTokenInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	dataValue := lexicalparserutil.ReadValueFrmDataMap(reqBrokerDataMap, execFunction.DecodeTokenInfo.TokenValue, "string", &rejectDesc)
	if dataValue == nil {
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ReadValueErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DecodeTokenInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DecodeTokenInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	lTokenValue := dataValue.(string)

	if len(execFunction.DecodeTokenInfo.TokenObject) == 0 {
		rejectDesc = fmt.Sprintf("TokenObject is NULL in DecodeTokenFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_DecodeTokenInfoFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DecodeTokenInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DecodeTokenInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}

	var payload map[string]interface{}
	retVal := tokenutil.DecodeToken(lTokenValue, &payload, &rejectDesc)
	if retVal < 0 {
		rejectDesc = fmt.Sprintf("VerifyToken() failed with rejectDesc[%s] in DecodeTokenFunctionType for FunctionName[%s]", rejectDesc, execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_DecodeTokenInfoFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DecodeTokenInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DecodeTokenInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	objectMap := make(map[string]interface{})
	userIDVal, ok := payload["UserID"].(float64)
	if !ok {
		rejectDesc = "UserID not found or invalid in token"
		return -1
	}
	lUserID := int(userIDVal)

	email, ok := payload["Email"].(string)
	if !ok {
		rejectDesc = "Email not found or invalid in token"
		return -1
	}

	role, ok := payload["Role"].(string)
	if !ok {
		rejectDesc = "Role not found or invalid in token"
		return -1
	}

	objectMap["user_id"] = lUserID
	objectMap["email"] = email
	objectMap["role"] = role

	reqBrokerDataMap[execFunction.DecodeTokenInfo.TokenObject] = objectMap
	return 1
}
