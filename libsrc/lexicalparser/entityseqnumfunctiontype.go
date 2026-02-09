package lexicalparser

import (
	"fmt"
	"lmsapieng/include/common/lexicalparserdef"
	"lmsapieng/include/common/rejectdef"
	"lmsapieng/include/common/reqbrokerdef"
	"lmsapieng/libsrc/microsv/microsv_schema"
	"lmsapieng/libsrc/services/services_sequence"
	"lmsapieng/libsrc/utils/msgutil"
	"lmsapieng/libsrc/utils/templateutil"
)

func EntitySeqNumFunctionType(execFunction lexicalparserdef.ExecFunctionStruct, reqBrokerDataMap map[string]interface{}) int {
	var rejectDesc string
	if len(execFunction.EntitySeqNumInfo.ObjectName) == 0 {
		rejectDesc = fmt.Sprintf("ObjectName is NULL in EntitySeqNumFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_EntitySeqNumFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.EntitySeqNumInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.EntitySeqNumInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	retval, contextParams := microsv_schema.GetDBContextParams(reqBrokerDataMap, execFunction.EntitySeqNumInfo.SchemaInfo, &rejectDesc)
	if retval < 0 {
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GetDBContextParamsErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.EntitySeqNumInfo.AppErrDesc, execFunction.EntitySeqNumInfo.SchemaInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.EntitySeqNumInfo.AppCorrectiveAction, execFunction.EntitySeqNumInfo.SchemaInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	if services_sequence.GenEntitySeqNum(reqBrokerDataMap, reqBrokerDataMap[execFunction.EntitySeqNumInfo.ObjectName].(map[string]interface{}), contextParams...) < 0 {
		rejectDesc := fmt.Sprintf("GenEntitySeqNum() Failed")
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_EntitySeqNumFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.EntitySeqNumInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.EntitySeqNumInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	return 1
}
