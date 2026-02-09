package lexicalparser

import (
	"fmt"
	"lmsapieng/include/common/datatypedef"
	"lmsapieng/include/common/dbdef"
	"lmsapieng/include/common/filtersourcetypedef"
	"lmsapieng/include/common/lexicalparserdef"
	"lmsapieng/include/common/rejectdef"
	"lmsapieng/include/common/reqbrokerdef"
	"lmsapieng/libsrc/dbtab/dbtab_tableinfo"
	"lmsapieng/libsrc/microsv/microsv_schema"
	"lmsapieng/libsrc/utils/dbutil"
	"lmsapieng/libsrc/utils/exprutil"
	"lmsapieng/libsrc/utils/lexicalparserutil"
	"lmsapieng/libsrc/utils/maputil"
	"lmsapieng/libsrc/utils/msgutil"
	"lmsapieng/libsrc/utils/schemainfo"
	"lmsapieng/libsrc/utils/templateutil"
	"strings"
)

func DBSingleReadFunctionType(execFunction lexicalparserdef.ExecFunctionStruct, reqBrokerDataMap map[string]interface{}) int {
	var rejectDesc string
	var dbErr, dbRejectReason string
	var resultDBContext dbdef.DBContextDef
	resultMap := make(map[string]interface{})
	if len(execFunction.DBSingleReadInfo.TableName) == 0 {
		rejectDesc := fmt.Sprintf("TableName is NULL in DBSingleReadFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_DBSingleReadFTTableNullErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBSingleReadInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBSingleReadInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	retval, contextParams := microsv_schema.GetDBContextParams(reqBrokerDataMap, execFunction.DBSingleReadInfo.SchemaInfo, &rejectDesc)
	if retval < 0 {
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GetDBContextParamsErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBSingleReadInfo.AppErrDesc, execFunction.DBSingleReadInfo.SchemaInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBSingleReadInfo.AppCorrectiveAction, execFunction.DBSingleReadInfo.SchemaInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	// trace.Lg("contextParams[%s]", contextParams)
	if schemainfo.GetActiveDBContext(reqBrokerDataMap, &resultDBContext, &rejectDesc, contextParams[0], contextParams[1], contextParams[2], contextParams[3], contextParams[4], contextParams[5]) < 0 {
		rejectDesc = fmt.Sprintf("GetActiveDBContext() failed for Module:%s", contextParams)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GetActiveDBContextError, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBSingleReadInfo.AppErrDesc, execFunction.DBSingleReadInfo.SchemaInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBSingleReadInfo.AppCorrectiveAction, execFunction.DBSingleReadInfo.SchemaInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	selectQueryStr := ""
	if resultDBContext.DBType == dbdef.DBTypePostgreSQL {
		selectQueryStr = fmt.Sprintf(`SELECT * FROM "%s"."%s"`, resultDBContext.SchemaName, execFunction.DBSingleReadInfo.TableName)
	} else {
		selectQueryStr = fmt.Sprintf("SELECT * FROM %s", execFunction.DBSingleReadInfo.TableName)
	}
	if len(execFunction.DBSingleReadInfo.FilterInfo.FilterData) != 0 {
		selectQueryStr += " WHERE "
		for i := 0; i < len(execFunction.DBSingleReadInfo.FilterInfo.FilterData); i++ {
			if len(execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].FilterApply) != 0 {
				if !exprutil.IsExprTrue(reqBrokerDataMap, execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].FilterApply) {
					continue
				}
			}
			if len(execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].FilterName) == 0 {
				rejectDesc = fmt.Sprintf("FilterName is NULL in DBSingleReadFunctionType for FunctionName[%s]", execFunction.FunctionName)
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_DBSingleReadFTFilterDataErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBSingleReadInfo.AppErrDesc, execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBSingleReadInfo.AppCorrectiveAction, execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].AppCorrectiveAction), rejectDesc)
				return -1
			}
			if len(execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].FilterID) == 0 {
				rejectDesc = fmt.Sprintf("FilterID is NULL in DBSingleReadFunctionType for FunctionName[%s]", execFunction.FunctionName)
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_DBSingleReadFTFilterDataErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBSingleReadInfo.AppErrDesc, execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBSingleReadInfo.AppCorrectiveAction, execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].AppCorrectiveAction), rejectDesc)
				return -1
			}
			if len(execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].FilterDataType) == 0 {
				rejectDesc = fmt.Sprintf("FilterDataType is NULL in DBSingleReadFunctionType for FunctionName[%s]", execFunction.FunctionName)
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_DBSingleReadFTFilterDataErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBSingleReadInfo.AppErrDesc, execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBSingleReadInfo.AppCorrectiveAction, execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].AppCorrectiveAction), rejectDesc)
				return -1
			}
			if len(execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].FilterCondition) == 0 {
				rejectDesc = fmt.Sprintf("FilterCondition is NULL in DBSingleReadFunctionType for FunctionName[%s]", execFunction.FunctionName)
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_DBSingleReadFTFilterDataErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBSingleReadInfo.AppErrDesc, execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBSingleReadInfo.AppCorrectiveAction, execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].AppCorrectiveAction), rejectDesc)
				return -1
			}
			fldValue := ""
			if len(execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].FilterSourceType) != 0 {
				if execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].FilterSourceType == filtersourcetypedef.FilterSourceTypeRawValue {
					fldValue = execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].FilterID
				} else {
					dataValue := lexicalparserutil.ReadValueFrmDataMap(reqBrokerDataMap, execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].FilterID, execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].FilterDataType, &rejectDesc)
					if dataValue == nil {
						reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ReadValueErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBSingleReadInfo.AppErrDesc, execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBSingleReadInfo.AppCorrectiveAction, execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].AppCorrectiveAction), rejectDesc)
						return -1
					}
					if strings.EqualFold(execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].FilterDataType, datatypedef.DataTypeString) {
						fldValue = dataValue.(string)
					} else if strings.EqualFold(execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].FilterDataType, datatypedef.DataTypeDate) {
						fldValue = dataValue.(string)
					} else if strings.EqualFold(execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].FilterDataType, datatypedef.DataTypeInt) {
						fldValue = fmt.Sprintf("%d", dataValue.(int))
					} else if strings.EqualFold(execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].FilterDataType, datatypedef.DataTypeBoolean) {
						fldValue = fmt.Sprintf("%t", dataValue.(bool))
					} else {
						rejectDesc = fmt.Sprintf("FilterDataType[%s] not supported for FilterID[%s]", execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].FilterDataType, execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].FilterID)
						reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_FldTypeInvalidErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBSingleReadInfo.AppErrDesc, execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBSingleReadInfo.AppCorrectiveAction, execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].AppCorrectiveAction), rejectDesc)
						return -1
					}
				}
			} else {
				dataValue := lexicalparserutil.ReadValueFrmDataMap(reqBrokerDataMap, execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].FilterID, execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].FilterDataType, &rejectDesc)
				if dataValue == nil {
					reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ReadValueErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBSingleReadInfo.AppErrDesc, execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBSingleReadInfo.AppCorrectiveAction, execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].AppCorrectiveAction), rejectDesc)
					return -1
				}
				if strings.EqualFold(execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].FilterDataType, datatypedef.DataTypeString) {
					fldValue = dataValue.(string)
				} else if strings.EqualFold(execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].FilterDataType, datatypedef.DataTypeDate) {
					fldValue = dataValue.(string)
				} else if strings.EqualFold(execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].FilterDataType, datatypedef.DataTypeInt) {
					fldValue = fmt.Sprintf("%d", dataValue.(int))
				} else if strings.EqualFold(execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].FilterDataType, datatypedef.DataTypeBoolean) {
					fldValue = fmt.Sprintf("%t", dataValue.(bool))
				} else {
					rejectDesc = fmt.Sprintf("FilterDataType[%s] not supported for FilterID[%s]", execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].FilterDataType, execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].FilterID)
					reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_FldTypeInvalidErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBSingleReadInfo.AppErrDesc, execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBSingleReadInfo.AppCorrectiveAction, execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].AppCorrectiveAction), rejectDesc)
					return -1
				}
			}
			if resultDBContext.DBType == dbdef.DBTypePostgreSQL {
				if i == len(execFunction.DBSingleReadInfo.FilterInfo.FilterData)-1 {
					if len(execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].FilterDBDataType) != 0 && strings.EqualFold(execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].FilterDBDataType, datatypedef.DataTypeDate) {
						selectQueryStr += fmt.Sprintf(`TO_CHAR("%s", 'DDMMYYYY')='%s'`, execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].FilterName, fldValue)
					} else {
						selectQueryStr += fmt.Sprintf(`"%s"='%s'`, execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].FilterName, fldValue)
					}
				} else {
					if len(execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].FilterDBDataType) != 0 && strings.EqualFold(execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].FilterDBDataType, datatypedef.DataTypeDate) {
						selectQueryStr += fmt.Sprintf(`TO_CHAR("%s", 'DDMMYYYY')='%s' %s `, execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].FilterName, fldValue, execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].FilterCondition)
					} else {
						selectQueryStr += fmt.Sprintf(`"%s"='%s' %s `, execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].FilterName, fldValue, execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].FilterCondition)
					}
				}
			} else {
				if i == len(execFunction.DBSingleReadInfo.FilterInfo.FilterData)-1 {
					if len(execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].FilterDBDataType) != 0 && strings.EqualFold(execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].FilterDBDataType, datatypedef.DataTypeDate) {
						selectQueryStr += fmt.Sprintf(`TO_CHAR("%s", 'DDMMYYYY')='%s'`, execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].FilterName, fldValue)
					} else {
						selectQueryStr += fmt.Sprintf("%s='%s'", execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].FilterName, fldValue)
					}
				} else {
					if len(execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].FilterDBDataType) != 0 && strings.EqualFold(execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].FilterDBDataType, datatypedef.DataTypeDate) {
						selectQueryStr += fmt.Sprintf(`TO_CHAR("%s", 'DDMMYYYY')='%s' %s `, execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].FilterName, fldValue, execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].FilterCondition)
					} else {
						selectQueryStr += fmt.Sprintf("%s='%s' %s ", execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].FilterName, fldValue, execFunction.DBSingleReadInfo.FilterInfo.FilterData[i].FilterCondition)
					}
				}
			}
		}
	}
	// trace.Lg("selectQueryStr[%s]", selectQueryStr)
	defer handleDBErr(execFunction, reqBrokerDataMap, &dbErr, &resultDBContext)
	dbResult := dbtab_tableinfo.ReadFromDBTable(selectQueryStr, resultMap, resultDBContext, &dbErr, &dbRejectReason)
	if len(execFunction.DBSingleReadInfo.ResultInfo.ResultSuccess) != 0 {
		if strings.EqualFold(execFunction.DBSingleReadInfo.ResultInfo.ResultSuccess, dbdef.DBNoRows) {
			if dbutil.IsNoRows(resultDBContext.DBType, dbErr) {
				return 1
			} else {
				if dbResult < 0 {
					rejectDesc = fmt.Sprintf("SingleTableReadError for Table[%s]Query[%s] with Err[%s]", execFunction.DBSingleReadInfo.TableName, selectQueryStr, dbErr)
					reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ReadFromDBTableErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBSingleReadInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBSingleReadInfo.AppCorrectiveAction), rejectDesc)
					return -1
				} else {
					rejectDesc = fmt.Sprintf("SingleTableReadError for Table[%s]Query[%s] with Err[Data Already Exists]", execFunction.DBSingleReadInfo.TableName, selectQueryStr)
					reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ReadFromDBTableErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBSingleReadInfo.ResultInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBSingleReadInfo.ResultInfo.AppCorrectiveAction), rejectDesc)
					return -1
				}
			}
		}
	}
	if dbResult < 0 {
		rejectDesc = fmt.Sprintf("SingleTableReadError for Table[%s]Query[%s] with Err[%s]", execFunction.DBSingleReadInfo.TableName, selectQueryStr, dbErr)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ReadFromDBTableErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBSingleReadInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBSingleReadInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	if len(execFunction.DBSingleReadInfo.StoreObjName) != 0 {
		if maputil.StoreObject(reqBrokerDataMap, execFunction.DBSingleReadInfo.StoreObjName, resultMap) < 0 {
			rejectDesc = fmt.Sprintf("StoreObject() failed for StoreObjName[%s]", execFunction.DBSingleReadInfo.StoreObjName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_StoreObjectErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBSingleReadInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBSingleReadInfo.AppCorrectiveAction), rejectDesc)
			return -1
		}
	}
	return 1
}
