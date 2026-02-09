package lexicalparser

import (
	"fmt"
	"lmsapieng/include/common/lexicalparserdef"
	"lmsapieng/include/common/rejectdef"
	"lmsapieng/include/common/reqbrokerdef"
	"lmsapieng/libsrc/utils/exprutil"
	"lmsapieng/libsrc/utils/lexicalparserutil"
	"lmsapieng/libsrc/utils/msgutil"
)

func StartLoopFunctionType(execFunction lexicalparserdef.ExecFunctionStruct, reqBrokerDataMap map[string]interface{}) int {
	var rejectDesc, correctiveAction string
	if len(execFunction.IndexName) == 0 {
		rejectDesc = fmt.Sprintf("index_name is null for FunctionType[%s]", execFunction.FunctionType)
		correctiveAction = fmt.Sprintf("Correct the index_name is null for FunctionType[%s]", execFunction.FunctionType)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_StartLoopFTErr, []byte(rejectDesc), []byte(correctiveAction), rejectDesc)
		return -1
	}
	if len(execFunction.StartIndex) == 0 {
		rejectDesc = fmt.Sprintf("start_index is null for FunctionType[%s]", execFunction.FunctionType)
		correctiveAction = fmt.Sprintf("Correct the start_index is null for FunctionType[%s]", execFunction.FunctionType)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_StartLoopFTErr, []byte(rejectDesc), []byte(correctiveAction), rejectDesc)
		return -1
	}
	if len(execFunction.EndIndex) == 0 {
		rejectDesc = fmt.Sprintf("end_index is null for FunctionType[%s]", execFunction.FunctionType)
		correctiveAction = fmt.Sprintf("Correct the end_index is null for FunctionType[%s]", execFunction.FunctionType)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_StartLoopFTErr, []byte(rejectDesc), []byte(correctiveAction), rejectDesc)
		return -1
	}
	startIndex := exprutil.GetNumericValueFromExpr(reqBrokerDataMap, execFunction.StartIndex)
	if startIndex < 0 {
		rejectDesc = fmt.Sprintf("start_index is should be numeric for FunctionType[%s]", execFunction.FunctionType)
		correctiveAction = fmt.Sprintf("Correct the start_index is should be numeric for FunctionType[%s]", execFunction.FunctionType)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_StartLoopFTErr, []byte(rejectDesc), []byte(correctiveAction), rejectDesc)
		return -1
	}
	endIndex := exprutil.GetNumericValueFromExpr(reqBrokerDataMap, execFunction.EndIndex)
	if endIndex < 0 {
		rejectDesc = fmt.Sprintf("end_index is should be numeric for FunctionType[%s]", execFunction.FunctionType)
		correctiveAction = fmt.Sprintf("Correct the end_index is should be numeric for FunctionType[%s]", execFunction.FunctionType)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_StartLoopFTErr, []byte(rejectDesc), []byte(correctiveAction), rejectDesc)
		return -1
	}
	if lexicalparserutil.StoreIndexValueForStartLoop(reqBrokerDataMap, execFunction.IndexName, startIndex, endIndex) < 0 {
		rejectDesc = fmt.Sprintf("StoreIndexValueForStartLoop() failed for IndexName[%s] startIndex[%d] endIndex[%d]", execFunction.IndexName, startIndex, endIndex)
		correctiveAction = fmt.Sprintf("Correct the StoreIndexValueForStartLoop() failed for IndexName[%s] startIndex[%d] endIndex[%d]", execFunction.IndexName, startIndex, endIndex)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_StartLoopFTErr, []byte(rejectDesc), []byte(correctiveAction), rejectDesc)
		return -1
	}
	return 1
}
