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

func ValidateConditionAndUpdateOnErrorFunctionType(execFunction lexicalparserdef.ExecFunctionStruct, reqBrokerDataMap map[string]interface{}) int {
	var rejectDesc string
	for i := 0; i < len(execFunction.ValidateConditionInfo.ValidateConditionExpression); i++ {
		if len(execFunction.ValidateConditionInfo.ValidateConditionExpression[i]) == 0 {
			rejectDesc = fmt.Sprintf("ValidateConditionExpression is NULL for FunctionName[%s]", execFunction.FunctionName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ValidateConditionFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ValidateConditionInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ValidateConditionInfo.AppCorrectiveAction), rejectDesc)
			return -1
		}
		if !exprutil.IsExprTrue(reqBrokerDataMap, execFunction.ValidateConditionInfo.ValidateConditionExpression[i]) {
			DBUpdate(execFunction, reqBrokerDataMap)
			rejectDesc = fmt.Sprintf("ValidateConditionExpression[%s] is False for FunctionName[%s]", execFunction.ValidateConditionInfo.ValidateConditionExpression[i], execFunction.FunctionName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ValidateConditionFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ValidateConditionInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ValidateConditionInfo.AppCorrectiveAction), rejectDesc)
			return -1
		}
	}
	return 1
}

func DBUpdate(execFunction lexicalparserdef.ExecFunctionStruct, reqBrokerDataMap map[string]interface{}) int {

	var rejectDesc, fldValue string
	var dbErr, dbRejectReason string
	var resultDBContext dbdef.DBContextDef
	//trace.Lg("DBUpdateFunctionType...reqBrokerDataMapBuffer[%s]", reqbrokerutil.GetReqBrokerDataMapBuffer(reqBrokerDataMap))

	retval, contextParams := microsv_schema.GetDBContextParams(reqBrokerDataMap, execFunction.DBUpdateInfo.SchemaInfo, &rejectDesc)
	if retval < 0 {
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GetDBContextParamsErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBUpdateInfo.AppErrDesc, execFunction.DBUpdateInfo.SchemaInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBUpdateInfo.AppCorrectiveAction, execFunction.DBUpdateInfo.SchemaInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	//trace.Lg("contextParams[%s]", contextParams)

	if schemainfo.GetActiveDBContext(reqBrokerDataMap, &resultDBContext, &rejectDesc, contextParams[0], contextParams[1], contextParams[2], contextParams[3], contextParams[4], contextParams[5]) < 0 {
		rejectDesc = fmt.Sprintf("GetActiveDBContext() failed for Module:%s", contextParams)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GetActiveDBContextError, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBUpdateInfo.AppErrDesc, execFunction.DBUpdateInfo.SchemaInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBUpdateInfo.AppCorrectiveAction, execFunction.DBUpdateInfo.SchemaInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}

	if len(execFunction.DBUpdateInfo.TableName) == 0 {
		rejectDesc := fmt.Sprintf("TableName is NULL in DBUpdateFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_DBUpdateFTTableNullErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBUpdateInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBUpdateInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	updateSetQuery := ""
	for i := 0; i < len(execFunction.DBUpdateInfo.UpdateDataList.UpdateData); i++ {
		if len(execFunction.DBUpdateInfo.UpdateDataList.UpdateData[i].Condition) != 0 {
			if !exprutil.IsExprTrue(reqBrokerDataMap, execFunction.DBUpdateInfo.UpdateDataList.UpdateData[i].Condition) {
				continue
			}
		}
		if len(execFunction.DBUpdateInfo.UpdateDataList.UpdateData[i].UpdateDataName) == 0 {
			rejectDesc := fmt.Sprintf("UpdateDataName is NULL in DBUpdateFunctionType for FunctionName[%s]", execFunction.FunctionName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_DBUpdateFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBUpdateInfo.AppErrDesc, execFunction.DBUpdateInfo.UpdateDataList.UpdateData[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBUpdateInfo.AppCorrectiveAction, execFunction.DBUpdateInfo.UpdateDataList.UpdateData[i].AppCorrectiveAction), rejectDesc)
			return -1
		}
		if len(execFunction.DBUpdateInfo.UpdateDataList.UpdateData[i].UpdateDBDataType) == 0 {
			rejectDesc := fmt.Sprintf("UpdateDBDataType is NULL in DBUpdateFunctionType for FunctionName[%s]", execFunction.FunctionName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_DBUpdateFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBUpdateInfo.AppErrDesc, execFunction.DBUpdateInfo.UpdateDataList.UpdateData[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBUpdateInfo.AppCorrectiveAction, execFunction.DBUpdateInfo.UpdateDataList.UpdateData[i].AppCorrectiveAction), rejectDesc)
			return -1
		}
		if len(execFunction.DBUpdateInfo.UpdateDataList.UpdateData[i].UpdateDataSource) == 0 {
			rejectDesc := fmt.Sprintf("UpdateDataSource is NULL in DBUpdateFunctionType for FunctionName[%s]", execFunction.FunctionName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_DBUpdateFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBUpdateInfo.AppErrDesc, execFunction.DBUpdateInfo.UpdateDataList.UpdateData[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBUpdateInfo.AppCorrectiveAction, execFunction.DBUpdateInfo.UpdateDataList.UpdateData[i].AppCorrectiveAction), rejectDesc)
			return -1
		}
		if len(execFunction.DBUpdateInfo.UpdateDataList.UpdateData[i].UpdateDataType) == 0 {
			rejectDesc := fmt.Sprintf("UpdateDataType is NULL in DBUpdateFunctionType for FunctionName[%s]", execFunction.FunctionName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_DBUpdateFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBUpdateInfo.AppErrDesc, execFunction.DBUpdateInfo.UpdateDataList.UpdateData[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBUpdateInfo.AppCorrectiveAction, execFunction.DBUpdateInfo.UpdateDataList.UpdateData[i].AppCorrectiveAction), rejectDesc)
			return -1
		}
		dataValue := lexicalparserutil.ReadValueFrmDataMap(reqBrokerDataMap, execFunction.DBUpdateInfo.UpdateDataList.UpdateData[i].UpdateDataSource, execFunction.DBUpdateInfo.UpdateDataList.UpdateData[i].UpdateDataType, &rejectDesc)
		if dataValue == nil {
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ReadValueErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBUpdateInfo.AppErrDesc, execFunction.DBUpdateInfo.UpdateDataList.UpdateData[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBUpdateInfo.AppCorrectiveAction, execFunction.DBUpdateInfo.UpdateDataList.UpdateData[i].AppCorrectiveAction), rejectDesc)
			return -1
		}
		if strings.EqualFold(execFunction.DBUpdateInfo.UpdateDataList.UpdateData[i].UpdateDataType, datatypedef.DataTypeString) {
			fldValue = dataValue.(string)
		} else if strings.EqualFold(execFunction.DBUpdateInfo.UpdateDataList.UpdateData[i].UpdateDataType, datatypedef.DataTypeInt) {
			switch v := dataValue.(type) {
			case int, int64:
				{
					fldValue = fmt.Sprintf("%d", v)
					break
				}
			default:
				rejectDesc = fmt.Sprintf("UpdateDataType[%s] not supported for UpdateDataName[%s]", execFunction.DBUpdateInfo.UpdateDataList.UpdateData[i].UpdateDataType, execFunction.DBUpdateInfo.UpdateDataList.UpdateData[i].UpdateDataName)
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_FldTypeInvalidErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBUpdateInfo.AppErrDesc, execFunction.DBUpdateInfo.UpdateDataList.UpdateData[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBUpdateInfo.AppCorrectiveAction, execFunction.DBUpdateInfo.UpdateDataList.UpdateData[i].AppCorrectiveAction), rejectDesc)
				return -1
			}
		} else if strings.EqualFold(execFunction.DBUpdateInfo.UpdateDataList.UpdateData[i].UpdateDataType, datatypedef.DataTypeBoolean) {

			// NEW BOOLEAN HANDLER
			switch v := dataValue.(type) {

			case bool:
				if v {
					fldValue = "true"
				} else {
					fldValue = "false"
				}

			case string:
				l := strings.ToLower(v)
				if l == "1" || l == "true" || l == "y" || l == "yes" {
					fldValue = "true"
				} else {
					fldValue = "false"
				}

			default:
				rejectDesc = fmt.Sprintf("UpdateDataType[bool] not supported for UpdateDataName[%s]", execFunction.DBUpdateInfo.UpdateDataList.UpdateData[i].UpdateDataName)
				// reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_FldTypeInvalidErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBUpdateInfo.AppErrDesc, execFunction.DBUpdateDataList.UpdateData[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBUpdateInfo.AppCorrectiveAction, execFunction.DBUpdateDataList.UpdateData[i].AppCorrectiveAction), rejectDesc)
				return -1
			}

		} else {
			rejectDesc = fmt.Sprintf("UpdateDataType[%s] not supported for UpdateDataName[%s]", execFunction.DBUpdateInfo.UpdateDataList.UpdateData[i].UpdateDataType, execFunction.DBUpdateInfo.UpdateDataList.UpdateData[i].UpdateDataName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_FldTypeInvalidErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBUpdateInfo.AppErrDesc, execFunction.DBUpdateInfo.UpdateDataList.UpdateData[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBUpdateInfo.AppCorrectiveAction, execFunction.DBUpdateInfo.UpdateDataList.UpdateData[i].AppCorrectiveAction), rejectDesc)
			return -1
		}
		if strings.EqualFold(execFunction.DBUpdateInfo.UpdateDataList.UpdateData[i].UpdateDBDataType, dbdef.DBDataTypeDate) {
			if resultDBContext.DBType == dbdef.DBTypeOracle {
				if i == len(execFunction.DBUpdateInfo.UpdateDataList.UpdateData)-1 {
					updateSetQuery += fmt.Sprintf("%s=TO_DATE('%s', 'DDMMYYYY')", execFunction.DBUpdateInfo.UpdateDataList.UpdateData[i].UpdateDataName, fldValue)
				} else {
					updateSetQuery += fmt.Sprintf("%s=TO_DATE('%s', 'DDMMYYYY'),", execFunction.DBUpdateInfo.UpdateDataList.UpdateData[i].UpdateDataName, fldValue)
				}
			} else if resultDBContext.DBType == dbdef.DBTypePostgreSQL {
				if i == len(execFunction.DBUpdateInfo.UpdateDataList.UpdateData)-1 {
					updateSetQuery += fmt.Sprintf(`"%s"=TO_DATE('%s', 'DDMMYYYY')`, execFunction.DBUpdateInfo.UpdateDataList.UpdateData[i].UpdateDataName, fldValue)
				} else {
					updateSetQuery += fmt.Sprintf(`"%s"=TO_DATE('%s', 'DDMMYYYY'),`, execFunction.DBUpdateInfo.UpdateDataList.UpdateData[i].UpdateDataName, fldValue)
				}
			} else {
				rejectDesc = fmt.Sprintf("Date Format Not Supported for DBType[%s]", resultDBContext.DBType)
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_DBUpdateFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBUpdateInfo.AppErrDesc, execFunction.DBUpdateInfo.UpdateDataList.UpdateData[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBUpdateInfo.AppCorrectiveAction, execFunction.DBUpdateInfo.UpdateDataList.UpdateData[i].AppCorrectiveAction), rejectDesc)
				return -1
			}
		} else if strings.EqualFold(execFunction.DBUpdateInfo.UpdateDataList.UpdateData[i].UpdateDBDataType, dbdef.DBDataTypeBoolean) {

			colName := execFunction.DBUpdateInfo.UpdateDataList.UpdateData[i].UpdateDataName

			if resultDBContext.DBType == dbdef.DBTypePostgreSQL {

				// booleans must be unquoted: true / false
				if i == len(execFunction.DBUpdateInfo.UpdateDataList.UpdateData)-1 {
					updateSetQuery += fmt.Sprintf(`"%s"=%s`, colName, fldValue)
				} else {
					updateSetQuery += fmt.Sprintf(`"%s"=%s,`, colName, fldValue)
				}

			} else if resultDBContext.DBType == dbdef.DBTypeOracle {

				boolChar := "N"
				if strings.EqualFold(fldValue, "true") {
					boolChar = "Y"
				}

				if i == len(execFunction.DBUpdateInfo.UpdateDataList.UpdateData)-1 {
					updateSetQuery += fmt.Sprintf(`"%s"='%s'`, colName, boolChar)
				} else {
					updateSetQuery += fmt.Sprintf(`"%s"='%s',`, colName, boolChar)
				}

			} else {
				rejectDesc = fmt.Sprintf("Boolean Not Supported for DBType[%s]", resultDBContext.DBType)
				// return reject
				return -1
			}
		} else {
			if resultDBContext.DBType == dbdef.DBTypePostgreSQL {
				if i == len(execFunction.DBUpdateInfo.UpdateDataList.UpdateData)-1 {
					updateSetQuery += fmt.Sprintf(`"%s"='%s'`, execFunction.DBUpdateInfo.UpdateDataList.UpdateData[i].UpdateDataName, fldValue)
				} else {
					updateSetQuery += fmt.Sprintf(`"%s"='%s',`, execFunction.DBUpdateInfo.UpdateDataList.UpdateData[i].UpdateDataName, fldValue)
				}
			} else {
				if i == len(execFunction.DBUpdateInfo.UpdateDataList.UpdateData)-1 {
					updateSetQuery += fmt.Sprintf("%s='%s'", execFunction.DBUpdateInfo.UpdateDataList.UpdateData[i].UpdateDataName, fldValue)
				} else {
					updateSetQuery += fmt.Sprintf("%s='%s',", execFunction.DBUpdateInfo.UpdateDataList.UpdateData[i].UpdateDataName, fldValue)
				}
			}
		}
	}
	trace.Lg("updateSetQuery[%s]", updateSetQuery)

	updateFilterQuery := ""
	for i := 0; i < len(execFunction.DBUpdateInfo.UpdateFilterList.UpdateFilter); i++ {
		if len(execFunction.DBUpdateInfo.UpdateFilterList.UpdateFilter[i].Condition) != 0 {
			if !exprutil.IsExprTrue(reqBrokerDataMap, execFunction.DBUpdateInfo.UpdateFilterList.UpdateFilter[i].Condition) {
				continue
			}
		}
		if len(execFunction.DBUpdateInfo.UpdateFilterList.UpdateFilter[i].UpdateFilterName) == 0 {
			rejectDesc := fmt.Sprintf("UpdateFilterName is NULL in DBUpdateFunctionType for FunctionName[%s]", execFunction.FunctionName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_DBUpdateFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBUpdateInfo.AppErrDesc, execFunction.DBUpdateInfo.UpdateFilterList.UpdateFilter[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBUpdateInfo.AppCorrectiveAction, execFunction.DBUpdateInfo.UpdateFilterList.UpdateFilter[i].AppCorrectiveAction), rejectDesc)
			return -1
		}
		if len(execFunction.DBUpdateInfo.UpdateFilterList.UpdateFilter[i].UpdateFilterDBDataType) == 0 {
			rejectDesc := fmt.Sprintf("UpdateFilterDBDataType is NULL in DBUpdateFunctionType for FunctionName[%s]", execFunction.FunctionName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_DBUpdateFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBUpdateInfo.AppErrDesc, execFunction.DBUpdateInfo.UpdateFilterList.UpdateFilter[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBUpdateInfo.AppCorrectiveAction, execFunction.DBUpdateInfo.UpdateFilterList.UpdateFilter[i].AppCorrectiveAction), rejectDesc)
			return -1
		}
		if len(execFunction.DBUpdateInfo.UpdateFilterList.UpdateFilter[i].UpdateFilterDataSource) == 0 {
			rejectDesc := fmt.Sprintf("UpdateFilterDataSource is NULL in DBUpdateFunctionType for FunctionName[%s]", execFunction.FunctionName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_DBUpdateFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBUpdateInfo.AppErrDesc, execFunction.DBUpdateInfo.UpdateFilterList.UpdateFilter[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBUpdateInfo.AppCorrectiveAction, execFunction.DBUpdateInfo.UpdateFilterList.UpdateFilter[i].AppCorrectiveAction), rejectDesc)
			return -1
		}
		if len(execFunction.DBUpdateInfo.UpdateFilterList.UpdateFilter[i].UpdateFilterDataType) == 0 {
			rejectDesc := fmt.Sprintf("UpdateFilterDataType is NULL in DBUpdateFunctionType for FunctionName[%s]", execFunction.FunctionName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_DBUpdateFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBUpdateInfo.AppErrDesc, execFunction.DBUpdateInfo.UpdateFilterList.UpdateFilter[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBUpdateInfo.AppCorrectiveAction, execFunction.DBUpdateInfo.UpdateFilterList.UpdateFilter[i].AppCorrectiveAction), rejectDesc)
			return -1
		}
		dataValue := lexicalparserutil.ReadValueFrmDataMap(reqBrokerDataMap, execFunction.DBUpdateInfo.UpdateFilterList.UpdateFilter[i].UpdateFilterDataSource, execFunction.DBUpdateInfo.UpdateFilterList.UpdateFilter[i].UpdateFilterDataType, &rejectDesc)
		if dataValue == nil {
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ReadValueErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBUpdateInfo.AppErrDesc, execFunction.DBUpdateInfo.UpdateFilterList.UpdateFilter[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBUpdateInfo.AppCorrectiveAction, execFunction.DBUpdateInfo.UpdateFilterList.UpdateFilter[i].AppCorrectiveAction), rejectDesc)
			return -1
		}
		if strings.EqualFold(execFunction.DBUpdateInfo.UpdateFilterList.UpdateFilter[i].UpdateFilterDataType, datatypedef.DataTypeString) {
			fldValue = dataValue.(string)
		} else if strings.EqualFold(execFunction.DBUpdateInfo.UpdateFilterList.UpdateFilter[i].UpdateFilterDataType, datatypedef.DataTypeInt) {
			fldValue = fmt.Sprintf("%d", dataValue.(int))
		} else {
			rejectDesc = fmt.Sprintf("UpdateFilterDataType[%s] not supported for UpdateFilterName[%s]", execFunction.DBUpdateInfo.UpdateFilterList.UpdateFilter[i].UpdateFilterDataType, execFunction.DBUpdateInfo.UpdateFilterList.UpdateFilter[i].UpdateFilterName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_FldTypeInvalidErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBUpdateInfo.AppErrDesc, execFunction.DBUpdateInfo.UpdateFilterList.UpdateFilter[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBUpdateInfo.AppCorrectiveAction, execFunction.DBUpdateInfo.UpdateFilterList.UpdateFilter[i].AppCorrectiveAction), rejectDesc)
			return -1
		}
		if strings.EqualFold(execFunction.DBUpdateInfo.UpdateFilterList.UpdateFilter[i].UpdateFilterDBDataType, dbdef.DBDataTypeDate) {
			if resultDBContext.DBType == dbdef.DBTypeOracle {
				if i == len(execFunction.DBUpdateInfo.UpdateFilterList.UpdateFilter)-1 {
					updateFilterQuery += fmt.Sprintf("%s=TO_DATE('%s', 'DDMMYYYY')", execFunction.DBUpdateInfo.UpdateFilterList.UpdateFilter[i].UpdateFilterName, fldValue)
				} else {
					updateFilterQuery += fmt.Sprintf("%s=TO_DATE('%s', 'DDMMYYYY') %s", execFunction.DBUpdateInfo.UpdateFilterList.UpdateFilter[i].UpdateFilterName, fldValue, execFunction.DBUpdateInfo.UpdateFilterList.UpdateFilter[i].UpdateFilterCondition)
				}
			} else if resultDBContext.DBType == dbdef.DBTypePostgreSQL {
				if i == len(execFunction.DBUpdateInfo.UpdateFilterList.UpdateFilter)-1 {
					updateFilterQuery += fmt.Sprintf(`"%s"=TO_DATE('%s', 'DDMMYYYY')`, execFunction.DBUpdateInfo.UpdateFilterList.UpdateFilter[i].UpdateFilterName, fldValue)
				} else {
					updateFilterQuery += fmt.Sprintf(`"%s"=TO_DATE('%s', 'DDMMYYYY') %s`, execFunction.DBUpdateInfo.UpdateFilterList.UpdateFilter[i].UpdateFilterName, fldValue, execFunction.DBUpdateInfo.UpdateFilterList.UpdateFilter[i].UpdateFilterCondition)
				}
			} else {
				rejectDesc = fmt.Sprintf("Date Format Not Supported for DBType[%s]", resultDBContext.DBType)
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_DBUpdateFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBUpdateInfo.AppErrDesc, execFunction.DBUpdateInfo.UpdateFilterList.UpdateFilter[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBUpdateInfo.AppCorrectiveAction, execFunction.DBUpdateInfo.UpdateFilterList.UpdateFilter[i].AppCorrectiveAction), rejectDesc)
				return -1
			}
		} else if strings.EqualFold(execFunction.DBUpdateInfo.UpdateFilterList.UpdateFilter[i].UpdateFilterDBDataType, dbdef.DBDataTypeByteArray) {
			if resultDBContext.DBType == dbdef.DBTypeOracle {
				if i == len(execFunction.DBUpdateInfo.UpdateDataList.UpdateData)-1 {
					updateFilterQuery += fmt.Sprintf("decode('%s', 'base64')", fldValue)
				} else {
					updateFilterQuery += fmt.Sprintf("decode('%s', 'base64'),", fldValue)
				}
			} else if resultDBContext.DBType == dbdef.DBTypePostgreSQL {
				if i == len(execFunction.DBUpdateInfo.UpdateDataList.UpdateData)-1 {
					updateFilterQuery += fmt.Sprintf(`decode('%s', 'base64')`, fldValue)
				} else {
					updateFilterQuery += fmt.Sprintf(`decode('%s', 'base64'),`, fldValue)
				}
			} else {
				rejectDesc = fmt.Sprintf("Binary data(bytea) Not Supported for DBType[%s]", resultDBContext.DBType)
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_DBUpdateFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBUpdateInfo.AppErrDesc, execFunction.DBUpdateInfo.UpdateDataList.UpdateData[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBUpdateInfo.AppCorrectiveAction, execFunction.DBUpdateInfo.UpdateDataList.UpdateData[i].AppCorrectiveAction), rejectDesc)
				return -1
			}
		} else {
			if resultDBContext.DBType == dbdef.DBTypePostgreSQL {
				if i == len(execFunction.DBUpdateInfo.UpdateFilterList.UpdateFilter)-1 {
					updateFilterQuery += fmt.Sprintf(`"%s"='%s'`, execFunction.DBUpdateInfo.UpdateFilterList.UpdateFilter[i].UpdateFilterName, fldValue)
				} else {
					updateFilterQuery += fmt.Sprintf(`"%s"='%s' %s `, execFunction.DBUpdateInfo.UpdateFilterList.UpdateFilter[i].UpdateFilterName, fldValue, execFunction.DBUpdateInfo.UpdateFilterList.UpdateFilter[i].UpdateFilterCondition)
				}
			} else {
				if i == len(execFunction.DBUpdateInfo.UpdateFilterList.UpdateFilter)-1 {
					updateFilterQuery += fmt.Sprintf("%s='%s'", execFunction.DBUpdateInfo.UpdateFilterList.UpdateFilter[i].UpdateFilterName, fldValue)
				} else {
					updateFilterQuery += fmt.Sprintf("%s='%s' %s ", execFunction.DBUpdateInfo.UpdateFilterList.UpdateFilter[i].UpdateFilterName, fldValue, execFunction.DBUpdateInfo.UpdateFilterList.UpdateFilter[i].UpdateFilterCondition)
				}
			}
		}
	}
	// trace.Lg("updateFilterQuery[%s]", updateFilterQuery)
	updateQuery := fmt.Sprintf("%s WHERE %s", updateSetQuery, updateFilterQuery)

	if schemainfo.GetActiveDBContextWithTxn(reqBrokerDataMap, &resultDBContext, &rejectDesc, contextParams[0], contextParams[1], contextParams[2], contextParams[3], contextParams[4], contextParams[5]) < 0 {
		rejectDesc = fmt.Sprintf("GetActiveDBContextWithTxn() failed for Module:%s", contextParams)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GetActiveDBContextWithTxnError, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBUpdateInfo.AppErrDesc, execFunction.DBUpdateInfo.SchemaInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBUpdateInfo.AppCorrectiveAction, execFunction.DBUpdateInfo.SchemaInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	if dbtab_tableinfo.UpdateTable(execFunction.DBUpdateInfo.TableName, updateQuery, resultDBContext, &dbErr, &dbRejectReason) < 0 {
		rejectDesc = fmt.Sprintf("UpdateTable for Table[%s] with Err[%s]", execFunction.DBUpdateInfo.TableName, dbErr)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_UpdateDBTableErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBUpdateInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBUpdateInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	if resultDBContext.DBTxFlag && resultDBContext.DBTx != nil {
		if err := resultDBContext.DBTx.Commit(); err != nil {
			rejectDesc = fmt.Sprintf("DB Commit failed: %v", err)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_DBUpdateFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, "", ""), templateutil.GetTemplateStr(reqBrokerDataMap, "", ""), rejectDesc)
			return -1
		}
	}
	return 1
}
