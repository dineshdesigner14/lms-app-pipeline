package lexicalparser

import (
	"fmt"
	"lmsapieng/include/common/dbdef"
	"lmsapieng/include/common/lexicalparserdef"
	"lmsapieng/include/common/rejectdef"
	"lmsapieng/include/common/reqbrokerdef"
	"lmsapieng/libsrc/dbtab/dbtab_tableinfo"
	"lmsapieng/libsrc/microsv/microsv_schema"
	"lmsapieng/libsrc/utils/maputil"
	"lmsapieng/libsrc/utils/msgutil"
	"lmsapieng/libsrc/utils/schemainfo"
	"lmsapieng/libsrc/utils/templateutil"
	"lmsapieng/libsrc/utils/trace"
)

func RawQueryFunctionType(execFunction lexicalparserdef.ExecFunctionStruct, reqBrokerDataMap map[string]interface{}) int {
	var rejectDesc string
	var dbErr, dbRejectReason string
	var resultDBContext dbdef.DBContextDef
	var resultMapArray []map[string]interface{}

	retval, contextParams := microsv_schema.GetDBContextParams(reqBrokerDataMap, execFunction.RawQueryInfo.SchemaInfo, &rejectDesc)
	if retval < 0 {
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GetDBContextParamsErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.RawQueryInfo.AppErrDesc, execFunction.RawQueryInfo.SchemaInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.RawQueryInfo.AppCorrectiveAction, execFunction.RawQueryInfo.SchemaInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	//trace.Lg("contextParams[%s]", contextParams)
	if schemainfo.GetActiveDBContext(reqBrokerDataMap, &resultDBContext, &rejectDesc, contextParams[0], contextParams[1], contextParams[2], contextParams[3], contextParams[4], contextParams[5]) < 0 {
		rejectDesc = fmt.Sprintf("GetActiveDBContext() failed for Module:%s", contextParams)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GetActiveDBContextError, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.RawQueryInfo.AppErrDesc, execFunction.RawQueryInfo.SchemaInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.RawQueryInfo.AppCorrectiveAction, execFunction.RawQueryInfo.SchemaInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	queryStr := ""
	if resultDBContext.DBType == dbdef.DBTypeOracle {
		if len(execFunction.RawQueryInfo.OracleQueryStr) == 0 {
			rejectDesc = fmt.Sprintf("OracleQueryStr is NULL in RawQueryFunctionType for FunctionName[%s]", execFunction.FunctionName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_RawQueryFunctionFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.RawQueryInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.RawQueryInfo.AppCorrectiveAction), rejectDesc)
			return -1
		}
		//trace.Lg("Passing OracleQueryStr[%s] to Template...", execFunction.RawQueryInfo.OracleQueryStr)
		queryStr = templateutil.GetTemplateString(reqBrokerDataMap, execFunction.RawQueryInfo.OracleQueryStr)
		//trace.Lg("Converted queryStr[%s]...", queryStr)
	} else if resultDBContext.DBType == dbdef.DBTypePostgreSQL {
		if len(execFunction.RawQueryInfo.PostgresOracleQueryStr) == 0 {
			rejectDesc = fmt.Sprintf("PostgresOracleQueryStr is NULL in RawQueryFunctionType for FunctionName[%s]", execFunction.FunctionName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_RawQueryFunctionFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.RawQueryInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.RawQueryInfo.AppCorrectiveAction), rejectDesc)
			return -1
		}
		//trace.Lg("Passing PostgresOracleQueryStr[%s] to Template...", execFunction.RawQueryInfo.PostgresOracleQueryStr)
		queryStr = templateutil.GetTemplateString(reqBrokerDataMap, execFunction.RawQueryInfo.PostgresOracleQueryStr)
		trace.Lg("Converted queryStr[%s]...", queryStr)
	} else {
		if len(execFunction.RawQueryInfo.OracleQueryStr) == 0 {
			rejectDesc = fmt.Sprintf("OracleQueryStr is NULL in RawQueryFunctionType for FunctionName[%s]", execFunction.FunctionName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_RawQueryFunctionFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.RawQueryInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.RawQueryInfo.AppCorrectiveAction), rejectDesc)
			return -1
		}
		//trace.Lg("Passing OracleQueryStr[%s] to Template...", execFunction.RawQueryInfo.OracleQueryStr)
		queryStr = templateutil.GetTemplateString(reqBrokerDataMap, execFunction.RawQueryInfo.OracleQueryStr)
		trace.Lg("Converted queryStr[%s]...", queryStr)
	}
	if len(queryStr) == 0 {
		rejectDesc = fmt.Sprintf("GetTemplateString Returns NULL for queryStr[%s] in RawQueryFunctionType for FunctionName[%s]", queryStr, execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_RawQueryFunctionFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.RawQueryInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.RawQueryInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	if dbtab_tableinfo.LoadFromDBTable(queryStr, &resultMapArray, resultDBContext, &dbErr, &dbRejectReason) < 0 {
		rejectDesc = fmt.Sprintf("MultiTableReadError for Query[%s] with Err[%s]", queryStr, dbErr)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_LoadFromDBTableErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.RawQueryInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.RawQueryInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	//trace.Lg("RawQuery...StoreObjName[%s]", execFunction.RawQueryInfo.StoreObjName)
	if len(execFunction.RawQueryInfo.DBResultInfo.ResultCode) != 0 && execFunction.RawQueryInfo.DBResultInfo.ResultCode == dbdef.DBNoRows {
		if len(resultMapArray) == 0 && execFunction.RawQueryInfo.DBResultInfo.Result == "failed" {
			rejectDesc = "No Data Found"
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_LoadFromDBTableErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.RawQueryInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.RawQueryInfo.AppCorrectiveAction), rejectDesc)
			return -1
		}
	}
	if len(execFunction.RawQueryInfo.StoreObjName) != 0 {
		if maputil.StoreObject(reqBrokerDataMap, execFunction.RawQueryInfo.StoreObjName, resultMapArray) < 0 {
			rejectDesc = fmt.Sprintf("StoreObject for StoreObjName[%s]", execFunction.RawQueryInfo.StoreObjName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_StoreObjectErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.RawQueryInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.RawQueryInfo.AppCorrectiveAction), rejectDesc)
			return -1
		}
		//trace.Lg("Stored Rawquery object successfully in StoreObj[%s]", execFunction.RawQueryInfo.StoreObjName)
	}
	return 1
}
