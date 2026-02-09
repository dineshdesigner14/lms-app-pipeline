package lexicalparser

import (
	"fmt"
	"lmsapieng/include/common/lexicalparserdef"
	"lmsapieng/include/common/rejectdef"
	"lmsapieng/include/common/reqbrokerdef"
	"lmsapieng/libsrc/utils/msgutil"
	"lmsapieng/libsrc/utils/templateutil"
)

func CreateEmptyListFunctionType(execFunction lexicalparserdef.ExecFunctionStruct, reqBrokerDataMap map[string]interface{}) int {
	var rejectDesc string

	if len(execFunction.CreateEmptyList.ObjectName) == 0 {
		rejectDesc = fmt.Sprintf("ObjectName is NULL in CreateEmptyListFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_DecodeTokenInfoFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DecodeTokenInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DecodeTokenInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	//trace.Lg("execFunction.CreateEmptyList.ObjectName[%s]", execFunction.CreateEmptyList.ObjectName)
	reqBrokerDataMap[execFunction.CreateEmptyList.ObjectName] = []string{}
	return 1
}
