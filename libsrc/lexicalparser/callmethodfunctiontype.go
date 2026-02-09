package lexicalparser

import (
	"fmt"
	"lmsapieng/include/common/lexicalparserdef"
	"lmsapieng/include/common/rejectdef"
	"lmsapieng/include/common/reqbrokerdef"
	"lmsapieng/libsrc/utils/msgutil"
	"reflect"
)

type CallMethodFTServiceInfo struct{}

func CallMethodFunctionType(execFunction lexicalparserdef.ExecFunctionStruct, reqBrokerDataMap map[string]interface{}) int {
	var rejectDesc, correctiveAction string
	var serviceRef CallMethodFTServiceInfo
	functionName := execFunction.FunctionName
	method := reflect.ValueOf(&serviceRef).MethodByName(functionName)
	if !method.IsValid() {
		//trace.Lg("MethodName[%s] does not exist", functionName)
		rejectDesc = fmt.Sprintf("MethodName[%s] does not exist", functionName)
		correctiveAction = fmt.Sprintf("Implement the MethodName[%s] Properly", functionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_CallMethodFTErr, []byte(rejectDesc), []byte(correctiveAction), rejectDesc)
		return -2
	}
	retval := reflect.ValueOf(&serviceRef).MethodByName(functionName).Call([]reflect.Value{reflect.ValueOf(reqBrokerDataMap)})
	if len(retval) == 0 {
		//trace.Lg("MethodName[%s] does not return anything", functionName)
		rejectDesc = fmt.Sprintf("MethodName[%s] does not return anything", functionName)
		correctiveAction = fmt.Sprintf("Ensure that the MethodName[%s] Returns Proper Value", functionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_CallMethodFTErr, []byte(rejectDesc), []byte(correctiveAction), rejectDesc)
		return -1
	}
	return int(retval[0].Int())
}
