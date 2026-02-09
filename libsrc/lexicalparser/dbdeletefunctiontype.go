package lexicalparser

import (
	"fmt"
	"lmsapieng/include/common/datatypedef"
	"lmsapieng/include/common/dbdef"
	"lmsapieng/include/common/lexicalparserdef"
	"lmsapieng/include/common/rejectdef"
	"lmsapieng/include/common/reqbrokerdef"
	"lmsapieng/libsrc/dbtab/dbtab_tableinfo"
	"lmsapieng/libsrc/microsv/microsv_schema"
	"lmsapieng/libsrc/utils/exprutil"
	"lmsapieng/libsrc/utils/lexicalparserutil"
	"lmsapieng/libsrc/utils/msgutil"
	"lmsapieng/libsrc/utils/schemainfo"
	"lmsapieng/libsrc/utils/templateutil"
	"lmsapieng/libsrc/utils/trace"
	"strings"
)

func DBDeleteFunctionType(execFunction lexicalparserdef.ExecFunctionStruct, reqBrokerDataMap map[string]interface{}) int {
	var rejectDesc, fldValue string
	var dbErr, dbRejectReason string
	var resultDBContext dbdef.DBContextDef
	//trace.Lg("DBDeleteFunctionType...reqBrokerDataMapBuffer[%s]", reqbrokerutil.GetReqBrokerDataMapBuffer(reqBrokerDataMap))
	retval, contextParams := microsv_schema.GetDBContextParams(reqBrokerDataMap, execFunction.DBDeleteInfo.SchemaInfo, &rejectDesc)
	if retval < 0 {
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GetDBContextParamsErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBDeleteInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBDeleteInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	if schemainfo.GetActiveDBContext(reqBrokerDataMap, &resultDBContext, &rejectDesc, contextParams[0], contextParams[1], contextParams[2], contextParams[3], contextParams[4], contextParams[5]) < 0 {
		rejectDesc = fmt.Sprintf("GetActiveDBContext() failed for Module:%s", contextParams)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GetActiveDBContextError, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBDeleteInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBDeleteInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}

	if len(execFunction.DBDeleteInfo.TableName) == 0 {
		rejectDesc := fmt.Sprintf("TableName is NULL in DBDeleteFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_DBDeleteFTTableNullErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBDeleteInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBDeleteInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	deleteFilterQuery := ""
	for i := 0; i < len(execFunction.DBDeleteInfo.DeleteFilterList.DeleteFilter); i++ {
		if len(execFunction.DBDeleteInfo.DeleteFilterList.DeleteFilter[i].Condition) != 0 {
			if !exprutil.IsExprTrue(reqBrokerDataMap, execFunction.DBDeleteInfo.DeleteFilterList.DeleteFilter[i].Condition) {
				continue
			}
		}
		if len(execFunction.DBDeleteInfo.DeleteFilterList.DeleteFilter[i].DeleteFilterName) == 0 {
			rejectDesc := fmt.Sprintf("DeleteFilterName is NULL in DBDeleteFunctionType for FunctionName[%s]", execFunction.FunctionName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_DBDeleteFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBDeleteInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBDeleteInfo.AppCorrectiveAction), rejectDesc)
			return -1
		}
		if len(execFunction.DBDeleteInfo.DeleteFilterList.DeleteFilter[i].DeleteFilterDBDataType) == 0 {
			rejectDesc := fmt.Sprintf("DeleteFilterDBDataType is NULL in DBDeleteFunctionType for FunctionName[%s]", execFunction.FunctionName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_DBDeleteFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBDeleteInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBDeleteInfo.AppCorrectiveAction), rejectDesc)
			return -1
		}
		if len(execFunction.DBDeleteInfo.DeleteFilterList.DeleteFilter[i].DeleteFilterDataSource) == 0 {
			rejectDesc := fmt.Sprintf("DeleteFilterDataSource is NULL in DBDeleteFunctionType for FunctionName[%s]", execFunction.FunctionName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_DBDeleteFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBDeleteInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBDeleteInfo.AppCorrectiveAction), rejectDesc)
			return -1
		}
		if len(execFunction.DBDeleteInfo.DeleteFilterList.DeleteFilter[i].DeleteFilterDataType) == 0 {
			rejectDesc := fmt.Sprintf("DeleteFilterDataType is NULL in DBDeleteFunctionType for FunctionName[%s]", execFunction.FunctionName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_DBDeleteFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBDeleteInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBDeleteInfo.AppCorrectiveAction), rejectDesc)
			return -1
		}
		dataValue := lexicalparserutil.ReadValueFrmDataMap(reqBrokerDataMap, execFunction.DBDeleteInfo.DeleteFilterList.DeleteFilter[i].DeleteFilterDataSource, execFunction.DBDeleteInfo.DeleteFilterList.DeleteFilter[i].DeleteFilterDataType, &rejectDesc)
		if dataValue == nil {
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ReadValueErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBDeleteInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBDeleteInfo.AppCorrectiveAction), rejectDesc)
			return -1
		}
		if strings.EqualFold(execFunction.DBDeleteInfo.DeleteFilterList.DeleteFilter[i].DeleteFilterDataType, datatypedef.DataTypeString) {
			fldValue = dataValue.(string)
		} else if strings.EqualFold(execFunction.DBDeleteInfo.DeleteFilterList.DeleteFilter[i].DeleteFilterDataType, datatypedef.DataTypeInt) {
			fldValue = fmt.Sprintf("%d", dataValue.(int))
		} else {
			rejectDesc = fmt.Sprintf("DeleteFilterDataType[%s] not supported for DeleteFilterName[%s]", execFunction.DBDeleteInfo.DeleteFilterList.DeleteFilter[i].DeleteFilterDataType, execFunction.DBDeleteInfo.DeleteFilterList.DeleteFilter[i].DeleteFilterName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_FldTypeInvalidErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBDeleteInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBDeleteInfo.AppCorrectiveAction), rejectDesc)
			return -1
		}
		if strings.EqualFold(execFunction.DBDeleteInfo.DeleteFilterList.DeleteFilter[i].DeleteFilterDBDataType, dbdef.DBDataTypeDate) {
			if resultDBContext.DBType == dbdef.DBTypeOracle {
				if i == len(execFunction.DBDeleteInfo.DeleteFilterList.DeleteFilter)-1 {
					deleteFilterQuery += fmt.Sprintf("%s=TO_DATE('%s', 'DDMMYYYY')", execFunction.DBDeleteInfo.DeleteFilterList.DeleteFilter[i].DeleteFilterName, fldValue)
				} else {
					deleteFilterQuery += fmt.Sprintf("%s=TO_DATE('%s', 'DDMMYYYY') %s", execFunction.DBDeleteInfo.DeleteFilterList.DeleteFilter[i].DeleteFilterName, fldValue, execFunction.DBDeleteInfo.DeleteFilterList.DeleteFilter[i].DeleteFilterCondition)
				}
			} else if resultDBContext.DBType == dbdef.DBTypePostgreSQL {
				if i == len(execFunction.DBDeleteInfo.DeleteFilterList.DeleteFilter)-1 {
					deleteFilterQuery += fmt.Sprintf(`"%s"=TO_DATE('%s', 'DDMMYYYY')`, execFunction.DBDeleteInfo.DeleteFilterList.DeleteFilter[i].DeleteFilterName, fldValue)
				} else {
					deleteFilterQuery += fmt.Sprintf(`"%s"=TO_DATE('%s', 'DDMMYYYY') %s`, execFunction.DBDeleteInfo.DeleteFilterList.DeleteFilter[i].DeleteFilterName, fldValue, execFunction.DBDeleteInfo.DeleteFilterList.DeleteFilter[i].DeleteFilterCondition)
				}
			} else {
				rejectDesc = fmt.Sprintf("Date Format Not Supported for DBType[%s]", resultDBContext.DBType)
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_DBDeleteFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBDeleteInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBDeleteInfo.AppCorrectiveAction), rejectDesc)
				return -1
			}
		} else {
			if resultDBContext.DBType == dbdef.DBTypePostgreSQL {
				if i == len(execFunction.DBDeleteInfo.DeleteFilterList.DeleteFilter)-1 {
					deleteFilterQuery += fmt.Sprintf(`"%s"='%s'`, execFunction.DBDeleteInfo.DeleteFilterList.DeleteFilter[i].DeleteFilterName, fldValue)
				} else {
					deleteFilterQuery += fmt.Sprintf(`"%s"='%s' %s `, execFunction.DBDeleteInfo.DeleteFilterList.DeleteFilter[i].DeleteFilterName, fldValue, execFunction.DBDeleteInfo.DeleteFilterList.DeleteFilter[i].DeleteFilterCondition)
				}
			} else {
				if i == len(execFunction.DBDeleteInfo.DeleteFilterList.DeleteFilter)-1 {
					deleteFilterQuery += fmt.Sprintf("%s='%s'", execFunction.DBDeleteInfo.DeleteFilterList.DeleteFilter[i].DeleteFilterName, fldValue)
				} else {
					deleteFilterQuery += fmt.Sprintf("%s='%s' %s ", execFunction.DBDeleteInfo.DeleteFilterList.DeleteFilter[i].DeleteFilterName, fldValue, execFunction.DBDeleteInfo.DeleteFilterList.DeleteFilter[i].DeleteFilterCondition)
				}
			}
		}
	}
	if schemainfo.GetActiveDBContextWithTxn(reqBrokerDataMap, &resultDBContext, &rejectDesc, contextParams[0], contextParams[1], contextParams[2], contextParams[3], contextParams[4], contextParams[5]) < 0 {
		rejectDesc = fmt.Sprintf("GetActiveDBContextWithTxn() failed for Module:%s", contextParams)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GetActiveDBContextWithTxnError, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBDeleteInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBDeleteInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	trace.Lg("deleteFilterQuery[%s]", deleteFilterQuery)
	if dbtab_tableinfo.DeleteTable(execFunction.DBDeleteInfo.TableName, deleteFilterQuery, resultDBContext, &dbErr, &dbRejectReason) < 0 {
		rejectDesc = fmt.Sprintf("DeleteTable for Table[%s] with Err[%s]", execFunction.DBDeleteInfo.TableName, dbErr)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_DeleteDBTableErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBDeleteInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBDeleteInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	return 1
}
