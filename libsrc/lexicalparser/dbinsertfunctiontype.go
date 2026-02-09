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

func DBInsertFunctionType(execFunction lexicalparserdef.ExecFunctionStruct, reqBrokerDataMap map[string]interface{}) int {
	var rejectDesc string
	var dbErr, dbRejectReason string
	var resultDBContext dbdef.DBContextDef

	retval, contextParams := microsv_schema.GetDBContextParams(reqBrokerDataMap, execFunction.DBInsertInfo.SchemaInfo, &rejectDesc)
	if retval < 0 {
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GetDBContextParamsErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBInsertInfo.AppErrDesc, execFunction.DBInsertInfo.SchemaInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBInsertInfo.AppCorrectiveAction, execFunction.DBInsertInfo.SchemaInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	trace.Lg("contextParams[%s]", contextParams)
	if schemainfo.GetActiveDBContext(reqBrokerDataMap, &resultDBContext, &rejectDesc, contextParams[0], contextParams[1], contextParams[2], contextParams[3], contextParams[4], contextParams[5]) < 0 {
		rejectDesc = fmt.Sprintf("GetActiveDBContext() failed for Module:%s", contextParams)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GetActiveDBContextError, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBInsertInfo.AppErrDesc, execFunction.DBInsertInfo.SchemaInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBInsertInfo.AppCorrectiveAction, execFunction.DBInsertInfo.SchemaInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	//trace.Lg("reqBrokerDataMapBuffer[%s]", reqbrokerutil.GetReqBrokerDataMapBuffer(reqBrokerDataMap))
	if len(execFunction.DBInsertInfo.TableName) == 0 {
		rejectDesc := fmt.Sprintf("TableName is NULL in DBInsertFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_DBInsertFTTableNullErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBInsertInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBInsertInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	fldValue := ""
	insertQueryStr := ""
	insertQueryStr += "("
	for i := 0; i < len(execFunction.DBInsertInfo.InsertDataList.InsertData); i++ {
		if len(execFunction.DBInsertInfo.InsertDataList.InsertData[i].InsertDataName) == 0 {
			rejectDesc := fmt.Sprintf("InsertDataName is NULL in DBInsertFunctionType for FunctionName[%s]", execFunction.FunctionName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_DBInsertFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBInsertInfo.AppErrDesc, execFunction.DBInsertInfo.InsertDataList.InsertData[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBInsertInfo.AppCorrectiveAction, execFunction.DBInsertInfo.InsertDataList.InsertData[i].AppCorrectiveAction), rejectDesc)
			return -1
		}
		if resultDBContext.DBType == dbdef.DBTypePostgreSQL {
			if i == len(execFunction.DBInsertInfo.InsertDataList.InsertData)-1 {
				insertQueryStr += fmt.Sprintf(`"%s"`, execFunction.DBInsertInfo.InsertDataList.InsertData[i].InsertDataName)
			} else {
				insertQueryStr += fmt.Sprintf(`"%s",`, execFunction.DBInsertInfo.InsertDataList.InsertData[i].InsertDataName)
			}
		} else {
			if i == len(execFunction.DBInsertInfo.InsertDataList.InsertData)-1 {
				insertQueryStr += fmt.Sprintf("%s", execFunction.DBInsertInfo.InsertDataList.InsertData[i].InsertDataName)
			} else {
				insertQueryStr += fmt.Sprintf("%s,", execFunction.DBInsertInfo.InsertDataList.InsertData[i].InsertDataName)
			}
		}
	}
	insertQueryStr += ") VALUES ("
	for i := 0; i < len(execFunction.DBInsertInfo.InsertDataList.InsertData); i++ {
		if len(execFunction.DBInsertInfo.InsertDataList.InsertData[i].Condition) != 0 {
			if !exprutil.IsExprTrue(reqBrokerDataMap, execFunction.DBInsertInfo.InsertDataList.InsertData[i].Condition) {
				continue
			}
		}
		if len(execFunction.DBInsertInfo.InsertDataList.InsertData[i].InsertDBDataType) == 0 {
			rejectDesc := fmt.Sprintf("InsertDBDataType is NULL in DBInsertFunctionType for FunctionName[%s]", execFunction.FunctionName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_DBInsertFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBInsertInfo.AppErrDesc, execFunction.DBInsertInfo.InsertDataList.InsertData[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBInsertInfo.AppCorrectiveAction, execFunction.DBInsertInfo.InsertDataList.InsertData[i].AppCorrectiveAction), rejectDesc)
			return -1
		}
		if len(execFunction.DBInsertInfo.InsertDataList.InsertData[i].InsertDataSource) == 0 {
			rejectDesc := fmt.Sprintf("InsertDataSource is NULL in DBInsertFunctionType for FunctionName[%s]", execFunction.FunctionName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_DBInsertFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBInsertInfo.AppErrDesc, execFunction.DBInsertInfo.InsertDataList.InsertData[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBInsertInfo.AppCorrectiveAction, execFunction.DBInsertInfo.InsertDataList.InsertData[i].AppCorrectiveAction), rejectDesc)
			return -1
		}
		if len(execFunction.DBInsertInfo.InsertDataList.InsertData[i].InsertDataType) == 0 {
			rejectDesc := fmt.Sprintf("InsertDataType is NULL in DBInsertFunctionType for FunctionName[%s]", execFunction.FunctionName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_DBInsertFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBInsertInfo.AppErrDesc, execFunction.DBInsertInfo.InsertDataList.InsertData[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBInsertInfo.AppCorrectiveAction, execFunction.DBInsertInfo.InsertDataList.InsertData[i].AppCorrectiveAction), rejectDesc)
			return -1
		}
		dataValue := lexicalparserutil.ReadValueFrmDataMap(reqBrokerDataMap, execFunction.DBInsertInfo.InsertDataList.InsertData[i].InsertDataSource, execFunction.DBInsertInfo.InsertDataList.InsertData[i].InsertDataType, &rejectDesc)
		if dataValue == nil {
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ReadValueErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBInsertInfo.AppErrDesc, execFunction.DBInsertInfo.InsertDataList.InsertData[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBInsertInfo.AppCorrectiveAction, execFunction.DBInsertInfo.InsertDataList.InsertData[i].AppCorrectiveAction), rejectDesc)
			return -1
		}
		//trace.Lg("InsertDataType[%s]InsertDataName[%s] type[%T]", execFunction.DBInsertInfo.InsertDataList.InsertData[i].InsertDataType, execFunction.DBInsertInfo.InsertDataList.InsertData[i].InsertDataName, dataValue)
		if strings.EqualFold(execFunction.DBInsertInfo.InsertDataList.InsertData[i].InsertDataType, datatypedef.DataTypeString) {
			fldValue = dataValue.(string)
		} else if strings.EqualFold(execFunction.DBInsertInfo.InsertDataList.InsertData[i].InsertDataType, datatypedef.DataTypeInt) {
			switch v := dataValue.(type) {
			case int, int64:
				{
					fldValue = fmt.Sprintf("%d", v)
					break
				}
			default:
				rejectDesc = fmt.Sprintf("InsertDataType[%s] not supported for InsertDataName[%s]", execFunction.DBInsertInfo.InsertDataList.InsertData[i].InsertDataType, execFunction.DBInsertInfo.InsertDataList.InsertData[i].InsertDataName)
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_FldTypeInvalidErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBInsertInfo.AppErrDesc, execFunction.DBInsertInfo.InsertDataList.InsertData[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBInsertInfo.AppCorrectiveAction, execFunction.DBInsertInfo.InsertDataList.InsertData[i].AppCorrectiveAction), rejectDesc)
				return -1
			}
		} else if strings.EqualFold(execFunction.DBInsertInfo.InsertDataList.InsertData[i].InsertDataType, datatypedef.DataTypeFloat) {
			switch v := dataValue.(type) {
			case float32, float64:
				{
					fldValue = fmt.Sprintf("%f", v)
					break
				}
			default:
				rejectDesc = fmt.Sprintf("InsertDataType[%s] not supported for InsertDataName[%s]", execFunction.DBInsertInfo.InsertDataList.InsertData[i].InsertDataType, execFunction.DBInsertInfo.InsertDataList.InsertData[i].InsertDataName)
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_FldTypeInvalidErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBInsertInfo.AppErrDesc, execFunction.DBInsertInfo.InsertDataList.InsertData[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBInsertInfo.AppCorrectiveAction, execFunction.DBInsertInfo.InsertDataList.InsertData[i].AppCorrectiveAction), rejectDesc)
				return -1
			}
		} else if strings.EqualFold(execFunction.DBInsertInfo.InsertDataList.InsertData[i].InsertDataType, datatypedef.DataTypeBoolean) {
			var boolVal bool
			switch v := dataValue.(type) {

			case bool:
				boolVal = v

			case string:
				raw := strings.ToLower(strings.TrimSpace(v))
				if raw == "true" || raw == "1" {
					boolVal = true
				} else if raw == "false" || raw == "0" {
					boolVal = false
				} else {
					rejectDesc = fmt.Sprintf(
						"Invalid boolean string '%s' for field %s",
						v,
						execFunction.DBInsertInfo.InsertDataList.InsertData[i].InsertDataName)
					return -1
				}

			default:
				rejectDesc = fmt.Sprintf(
					"Invalid boolean type '%T' for field %s",
					v,
					execFunction.DBInsertInfo.InsertDataList.InsertData[i].InsertDataName)
				return -1
			}
			if boolVal {
				fldValue = "true"
			} else {
				fldValue = "false"
			}
		} else {
			rejectDesc = fmt.Sprintf("InsertDataType[%s] not supported for InsertDataName[%s]", execFunction.DBInsertInfo.InsertDataList.InsertData[i].InsertDataType, execFunction.DBInsertInfo.InsertDataList.InsertData[i].InsertDataName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_FldTypeInvalidErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBInsertInfo.AppErrDesc, execFunction.DBInsertInfo.InsertDataList.InsertData[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBInsertInfo.AppCorrectiveAction, execFunction.DBInsertInfo.InsertDataList.InsertData[i].AppCorrectiveAction), rejectDesc)
			return -1
		}
		if strings.EqualFold(execFunction.DBInsertInfo.InsertDataList.InsertData[i].InsertDBDataType, dbdef.DBDataTypeDate) {
			if resultDBContext.DBType == dbdef.DBTypeOracle {
				if i == len(execFunction.DBInsertInfo.InsertDataList.InsertData)-1 {
					insertQueryStr += fmt.Sprintf("TO_DATE('%s', 'DDMMYYYY')", fldValue)
				} else {
					insertQueryStr += fmt.Sprintf("TO_DATE('%s', 'DDMMYYYY'),", fldValue)
				}
			} else if resultDBContext.DBType == dbdef.DBTypePostgreSQL {
				if i == len(execFunction.DBInsertInfo.InsertDataList.InsertData)-1 {
					insertQueryStr += fmt.Sprintf(`TO_DATE('%s', 'DDMMYYYY')`, fldValue)
				} else {
					insertQueryStr += fmt.Sprintf(`TO_DATE('%s', 'DDMMYYYY'),`, fldValue)
				}
			} else {
				rejectDesc = fmt.Sprintf("Date Format Not Supported for DBType[%s]", resultDBContext.DBType)
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_DBInsertFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBInsertInfo.AppErrDesc, execFunction.DBInsertInfo.InsertDataList.InsertData[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBInsertInfo.AppCorrectiveAction, execFunction.DBInsertInfo.InsertDataList.InsertData[i].AppCorrectiveAction), rejectDesc)
				return -1
			}
		} else if strings.EqualFold(execFunction.DBInsertInfo.InsertDataList.InsertData[i].InsertDBDataType, dbdef.DBDataTypeByteArray) {
			if resultDBContext.DBType == dbdef.DBTypeOracle {
				if i == len(execFunction.DBInsertInfo.InsertDataList.InsertData)-1 {
					insertQueryStr += fmt.Sprintf("decode('%s', 'base64')", fldValue)
				} else {
					insertQueryStr += fmt.Sprintf("decode('%s', 'base64'),", fldValue)
				}
			} else if resultDBContext.DBType == dbdef.DBTypePostgreSQL {
				if i == len(execFunction.DBInsertInfo.InsertDataList.InsertData)-1 {
					insertQueryStr += fmt.Sprintf(`decode('%s', 'base64')`, fldValue)
				} else {
					insertQueryStr += fmt.Sprintf(`decode('%s', 'base64'),`, fldValue)
				}
			} else {
				rejectDesc = fmt.Sprintf("Binary data(bytea) Not Supported for DBType[%s]", resultDBContext.DBType)
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_DBInsertFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBInsertInfo.AppErrDesc, execFunction.DBInsertInfo.InsertDataList.InsertData[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBInsertInfo.AppCorrectiveAction, execFunction.DBInsertInfo.InsertDataList.InsertData[i].AppCorrectiveAction), rejectDesc)
				return -1
			}
		} else if strings.EqualFold(execFunction.DBInsertInfo.InsertDataList.InsertData[i].InsertDBDataType, dbdef.DBDataTypeBoolean) {
			if resultDBContext.DBType == dbdef.DBTypePostgreSQL {
				if i == len(execFunction.DBInsertInfo.InsertDataList.InsertData)-1 {
					insertQueryStr += fmt.Sprintf(`%s`, fldValue)
				} else {
					insertQueryStr += fmt.Sprintf(`%s,`, fldValue)
				}

			} else if resultDBContext.DBType == dbdef.DBTypeOracle {

				// Oracle: store boolean as Y/N
				boolChar := "N"
				if strings.EqualFold(fldValue, "true") {
					boolChar = "Y"
				}

				if i == len(execFunction.DBInsertInfo.InsertDataList.InsertData)-1 {
					insertQueryStr += fmt.Sprintf("'%s'", boolChar)
				} else {
					insertQueryStr += fmt.Sprintf("'%s',", boolChar)
				}

			} else {
				rejectDesc = fmt.Sprintf(
					"Boolean not supported for DB type [%s]",
					resultDBContext.DBType,
				)
				return -1
			}
		} else {
			if i == len(execFunction.DBInsertInfo.InsertDataList.InsertData)-1 {
				insertQueryStr += fmt.Sprintf("'%s'", fldValue)
			} else {
				insertQueryStr += fmt.Sprintf("'%s',", fldValue)
			}
		}
	}
	insertQueryStr += ")"
	trace.Lg("insertQueryStr[%s]", insertQueryStr)
	if schemainfo.GetActiveDBContextWithTxn(reqBrokerDataMap, &resultDBContext, &rejectDesc, contextParams[0], contextParams[1], contextParams[2], contextParams[3], contextParams[4], contextParams[5]) < 0 {
		rejectDesc = fmt.Sprintf("GetActiveDBContextWithTxn() failed for Module:%s", contextParams)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GetActiveDBContextWithTxnError, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBInsertInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBInsertInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	if dbtab_tableinfo.InsertTable(execFunction.DBInsertInfo.TableName, insertQueryStr, resultDBContext, &dbErr, &dbRejectReason) < 0 {
		rejectDesc = fmt.Sprintf("InsertTable for Table[%s] with Err[%s]", execFunction.DBInsertInfo.TableName, dbErr)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_InsertDBTableErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBInsertInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBInsertInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	return 1
}
