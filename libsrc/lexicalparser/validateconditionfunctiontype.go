package lexicalparser

import (
	"fmt"
	"lmsapieng/include/common/lexicalparserdef"
	"lmsapieng/include/common/rejectdef"
	"lmsapieng/include/common/reqbrokerdef"
	"lmsapieng/libsrc/utils/exprutil"
	"lmsapieng/libsrc/utils/msgutil"
	"lmsapieng/libsrc/utils/templateutil"
)

func ValidateConditionFunctionType(execFunction lexicalparserdef.ExecFunctionStruct, reqBrokerDataMap map[string]interface{}) int {
	var rejectDesc string
	for i := 0; i < len(execFunction.ValidateConditionInfo.ValidateConditionExpression); i++ {
		if len(execFunction.ValidateConditionInfo.ValidateConditionExpression[i]) == 0 {
			rejectDesc = fmt.Sprintf("ValidateConditionExpression is NULL for FunctionName[%s]", execFunction.FunctionName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ValidateConditionFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ValidateConditionInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ValidateConditionInfo.AppCorrectiveAction), rejectDesc)
			return -1
		}
		if !exprutil.IsExprTrue(reqBrokerDataMap, execFunction.ValidateConditionInfo.ValidateConditionExpression[i]) {
			rejectDesc = fmt.Sprintf("ValidateConditionExpression[%s] is False for FunctionName[%s]", execFunction.ValidateConditionInfo.ValidateConditionExpression[i], execFunction.FunctionName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ValidateConditionFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ValidateConditionInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ValidateConditionInfo.AppCorrectiveAction), rejectDesc)
			return -1
		}
	}
	return 1
}
