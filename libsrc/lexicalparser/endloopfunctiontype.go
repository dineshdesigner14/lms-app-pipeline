package lexicalparser

import (
	"fmt"
	"lmsapieng/include/common/lexicalparserdef"
	"lmsapieng/include/common/rejectdef"
	"lmsapieng/include/common/reqbrokerdef"
	"lmsapieng/libsrc/utils/lexicalparserutil"
	"lmsapieng/libsrc/utils/msgutil"
)

func EndLoopFunctionType(execFunction lexicalparserdef.ExecFunctionStruct, reqBrokerDataMap map[string]interface{}) int {
	var rejectDesc, correctiveAction string
	if len(execFunction.IndexName) == 0 {
		rejectDesc = fmt.Sprintf("index_name is null for FunctionType[%s]", execFunction.FunctionType)
		correctiveAction = fmt.Sprintf("Check for index_name is null for FunctionType[%s]", execFunction.FunctionType)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_EndLoopFTErr, []byte(rejectDesc), []byte(correctiveAction), rejectDesc)
		return -1
	}
	if lexicalparserutil.StoreIndexValueForEndLoop(reqBrokerDataMap, execFunction.IndexName) < 0 {
		rejectDesc = fmt.Sprintf("StoreIndexValueForEndLoop() failed for IndexName[%s]", execFunction.IndexName)
		correctiveAction = fmt.Sprintf("Check for StoreIndexValueForEndLoop() failed for IndexName[%s]", execFunction.IndexName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_EndLoopFTErr, []byte(rejectDesc), []byte(correctiveAction), rejectDesc)
		return -1
	}
	return 1
}
