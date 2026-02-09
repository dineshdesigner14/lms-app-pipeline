package lexicalparser

import (
	"encoding/json"
	"fmt"
	"lmsapieng/include/common/datatypedef"
	"lmsapieng/include/common/dbdef"
	"lmsapieng/include/common/filtersourcetypedef"
	"lmsapieng/include/common/lexicalparserdef"
	"lmsapieng/include/common/rejectdef"
	"lmsapieng/include/common/reqbrokerdef"
	"lmsapieng/libsrc/dbtab/dbtab_tableinfo"
	"lmsapieng/libsrc/microsv/microsv_schema"
	"lmsapieng/libsrc/utils/exprutil"
	"lmsapieng/libsrc/utils/lexicalparserutil"
	"lmsapieng/libsrc/utils/maputil"
	"lmsapieng/libsrc/utils/msgutil"
	"lmsapieng/libsrc/utils/schemainfo"
	"lmsapieng/libsrc/utils/templateutil"
	"lmsapieng/libsrc/utils/trace"
	"strconv"
	"strings"
)

func removeFilterConditionLastOccurrence(queryStr string) string {
	queryStr = strings.TrimSpace(queryStr)

	lastAndIndex := strings.LastIndex(queryStr, " AND")
	lastOrIndex := strings.LastIndex(queryStr, " OR")

	var lastIndex int
	var lengthToRemove int

	if lastAndIndex > lastOrIndex {
		lastIndex = lastAndIndex
		lengthToRemove = len(" AND")
	} else {
		lastIndex = lastOrIndex
		lengthToRemove = len(" OR")
	}

	// If no AND or OR found, return the original query
	if lastIndex == -1 {
		return queryStr
	}

	// Check if AND/OR is the last condition by verifying the substring following it
	if lastIndex+lengthToRemove == len(queryStr) {
		result := queryStr[:lastIndex]
		return strings.TrimSpace(result)
	}

	return queryStr
}

func DBMultiReadFunctionType(execFunction lexicalparserdef.ExecFunctionStruct, reqBrokerDataMap map[string]interface{}) int {
	var rejectDesc string
	var dbErr, dbRejectReason string
	var resultDBContext dbdef.DBContextDef
	var resultMapArray []map[string]interface{}
	if len(execFunction.DBMultiReadInfo.TableName) == 0 {
		rejectDesc := fmt.Sprintf("TableName is NULL in DBMultiReadFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_DBMultiReadFTTableNullErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBMultiReadInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBMultiReadInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	retval, contextParams := microsv_schema.GetDBContextParams(reqBrokerDataMap, execFunction.DBMultiReadInfo.SchemaInfo, &rejectDesc)
	if retval < 0 {
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GetDBContextParamsErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBMultiReadInfo.AppErrDesc, execFunction.DBMultiReadInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBMultiReadInfo.AppCorrectiveAction, execFunction.DBMultiReadInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	// trace.Lg("contextParams[%s]", contextParams)
	if schemainfo.GetActiveDBContext(reqBrokerDataMap, &resultDBContext, &rejectDesc, contextParams[0], contextParams[1], contextParams[2], contextParams[3], contextParams[4], contextParams[5]) < 0 {
		rejectDesc = fmt.Sprintf("GetActiveDBContext() failed for Module:%s", contextParams)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GetActiveDBContextError, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBMultiReadInfo.AppErrDesc, execFunction.DBMultiReadInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBMultiReadInfo.AppCorrectiveAction, execFunction.DBMultiReadInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	selectQueryStr := ""
	if resultDBContext.DBType == dbdef.DBTypePostgreSQL {
		//trace.Lg("selectQueryStr...1")
		selectQueryStr = fmt.Sprintf(`SELECT * FROM "%s"."%s"`, resultDBContext.SchemaName, execFunction.DBMultiReadInfo.TableName)
	} else {
		//trace.Lg("selectQueryStr...2")
		selectQueryStr = fmt.Sprintf("SELECT * FROM %s", execFunction.DBMultiReadInfo.TableName)
	}
	if len(execFunction.DBMultiReadInfo.FilterInfo.FilterData) != 0 {
		selectQueryStr += " WHERE "
		for i := 0; i < len(execFunction.DBMultiReadInfo.FilterInfo.FilterData); i++ {
			//trace.Lg("Filter Operator[%s] FilterName[%s]", execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterOperator, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName)
			if len(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterApply) != 0 {
				if !exprutil.IsExprTrue(reqBrokerDataMap, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterApply) {
					continue
				}
			}
			if len(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName) == 0 {
				rejectDesc := fmt.Sprintf("FilterName is NULL in DBMultiReadFunctionType for FunctionName[%s]", execFunction.FunctionName)
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_DBMultiReadFTFilterDataErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBMultiReadInfo.AppErrDesc, execFunction.DBMultiReadInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBMultiReadInfo.AppCorrectiveAction, execFunction.DBMultiReadInfo.AppCorrectiveAction), rejectDesc)
				return -1
			}
			if len(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterID) == 0 {
				rejectDesc := fmt.Sprintf("FilterID is NULL in DBMultiReadFunctionType for FunctionName[%s]", execFunction.FunctionName)
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_DBMultiReadFTFilterDataErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBMultiReadInfo.AppErrDesc, execFunction.DBMultiReadInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBMultiReadInfo.AppCorrectiveAction, execFunction.DBMultiReadInfo.AppCorrectiveAction), rejectDesc)
				return -1
			}
			if len(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDataType) == 0 {
				rejectDesc := fmt.Sprintf("FilterDataType is NULL in DBMultiReadFunctionType for FunctionName[%s]", execFunction.FunctionName)
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_DBMultiReadFTFilterDataErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBMultiReadInfo.AppErrDesc, execFunction.DBMultiReadInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBMultiReadInfo.AppCorrectiveAction, execFunction.DBMultiReadInfo.AppCorrectiveAction), rejectDesc)
				return -1
			}
			if len(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterCondition) == 0 {
				rejectDesc := fmt.Sprintf("FilterCondition is NULL in DBMultiReadFunctionType for FunctionName[%s]", execFunction.FunctionName)
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_DBMultiReadFTFilterDataErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBMultiReadInfo.AppErrDesc, execFunction.DBMultiReadInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBMultiReadInfo.AppCorrectiveAction, execFunction.DBMultiReadInfo.AppCorrectiveAction), rejectDesc)
				return -1
			}
			fldValue := ""
			fldValue1 := ""
			if execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterOperator == lexicalparserdef.FilterOperatorBetween {
				filterIDArray := strings.Split(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterID, ",")
				if len(filterIDArray) != 2 {
					rejectDesc := fmt.Sprintf("FilterID[%s] is Invalid for [%s] FilterOperator in DBMultiReadFunctionType for FunctionName[%s]", execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterID, lexicalparserdef.FilterOperatorBetween, execFunction.FunctionName)
					reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_DBMultiReadFTFilterDataErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBMultiReadInfo.AppErrDesc, execFunction.DBMultiReadInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBMultiReadInfo.AppCorrectiveAction, execFunction.DBMultiReadInfo.AppCorrectiveAction), rejectDesc)
					return -1
				}
				for j := 0; j < len(filterIDArray); j++ {
					dataValue := lexicalparserutil.ReadValueFrmDataMap(reqBrokerDataMap, filterIDArray[j], execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDataType, &rejectDesc)
					if dataValue == nil {
						reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ReadValueErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBMultiReadInfo.AppErrDesc, execFunction.DBMultiReadInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBMultiReadInfo.AppCorrectiveAction, execFunction.DBMultiReadInfo.AppCorrectiveAction), rejectDesc)
						return -1
					}
					if j == 0 {
						if strings.EqualFold(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDataType, datatypedef.DataTypeString) {
							fldValue = dataValue.(string)
						} else if strings.EqualFold(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDataType, datatypedef.DataTypeInt) {
							fldValue = fmt.Sprintf("%d", dataValue.(int))
						} else {
							rejectDesc := fmt.Sprintf("FilterDataType[%s] not supported for FilterID[%s]", execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDataType, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterID)
							reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_FldTypeInvalidErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBMultiReadInfo.AppErrDesc, execFunction.DBMultiReadInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBMultiReadInfo.AppCorrectiveAction, execFunction.DBMultiReadInfo.AppCorrectiveAction), rejectDesc)
							return -1
						}
					} else {
						if strings.EqualFold(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDataType, datatypedef.DataTypeString) {
							fldValue1 = dataValue.(string)
						} else if strings.EqualFold(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDataType, datatypedef.DataTypeInt) {
							fldValue1 = fmt.Sprintf("%d", dataValue.(int))
						} else {
							rejectDesc := fmt.Sprintf("FilterDataType[%s] not supported for FilterID[%s]", execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDataType, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterID)
							reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_FldTypeInvalidErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBMultiReadInfo.AppErrDesc, execFunction.DBMultiReadInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBMultiReadInfo.AppCorrectiveAction, execFunction.DBMultiReadInfo.AppCorrectiveAction), rejectDesc)
							return -1
						}
					}
				}
			} else {
				if len(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterSourceType) != 0 {
					if execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterSourceType == filtersourcetypedef.FilterSourceTypeRawValue {
						fldValue = execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterID
					} else {
						dataValue := lexicalparserutil.ReadValueFrmDataMap(reqBrokerDataMap, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterID, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDataType, &rejectDesc)
						if dataValue == nil {
							reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ReadValueErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBMultiReadInfo.AppErrDesc, execFunction.DBMultiReadInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBMultiReadInfo.AppCorrectiveAction, execFunction.DBMultiReadInfo.AppCorrectiveAction), rejectDesc)
							return -1
						}
						if strings.EqualFold(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDataType, datatypedef.DataTypeString) {
							fldValue = dataValue.(string)
						} else if strings.EqualFold(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDataType, datatypedef.DataTypeInt) {
							fldValue = fmt.Sprintf("%d", dataValue.(int))
						} else {
							rejectDesc := fmt.Sprintf("FilterDataType[%s] not supported for FilterID[%s]", execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDataType, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterID)
							reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_FldTypeInvalidErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBMultiReadInfo.AppErrDesc, execFunction.DBMultiReadInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBMultiReadInfo.AppCorrectiveAction, execFunction.DBMultiReadInfo.AppCorrectiveAction), rejectDesc)
							return -1
						}
					}
				} else {
					dataValue := lexicalparserutil.ReadValueFrmDataMap(reqBrokerDataMap, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterID, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDataType, &rejectDesc)
					if dataValue == nil {
						reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ReadValueErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBMultiReadInfo.AppErrDesc, execFunction.DBMultiReadInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBMultiReadInfo.AppCorrectiveAction, execFunction.DBMultiReadInfo.AppCorrectiveAction), rejectDesc)
						return -1
					}
					if strings.EqualFold(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDataType, datatypedef.DataTypeString) {
						fldValue = dataValue.(string)
					} else if strings.EqualFold(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDataType, datatypedef.DataTypeInt) {
						fldValue = fmt.Sprintf("%d", dataValue.(int))
					} else {
						rejectDesc := fmt.Sprintf("FilterDataType[%s] not supported for FilterID[%s]", execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDataType, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterID)
						reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_FldTypeInvalidErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBMultiReadInfo.AppErrDesc, execFunction.DBMultiReadInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBMultiReadInfo.AppCorrectiveAction, execFunction.DBMultiReadInfo.AppCorrectiveAction), rejectDesc)
						return -1
					}
				}
			}
			//trace.Lg(" Changed FilterOperator...1[%s]", execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterOperator)
			if resultDBContext.DBType == dbdef.DBTypePostgreSQL {
				if len(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterOperator) == 0 {
					//trace.Lg("FilterOperator is NULL")
					if i == len(execFunction.DBMultiReadInfo.FilterInfo.FilterData)-1 {
						if len(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType) != 0 && strings.EqualFold(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType, datatypedef.DataTypeDate) {
							selectQueryStr += fmt.Sprintf(`TO_CHAR("%s", 'DDMMYYYY')='%s'`, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue)
						} else {
							selectQueryStr += fmt.Sprintf(`"%s"='%s'`, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue)
						}
					} else {
						if len(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType) != 0 && strings.EqualFold(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType, datatypedef.DataTypeDate) {
							selectQueryStr += fmt.Sprintf(`TO_CHAR("%s", 'DDMMYYYY')='%s' %s `, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterCondition)
						} else {
							selectQueryStr += fmt.Sprintf(`"%s"='%s' %s `, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterCondition)
						}
					}
				} else {
					// trace.Lg(" Changed FilterOperator[%s]", execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterOperator)
					if execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterOperator == lexicalparserdef.FilterOperatorEqual {
						if i == len(execFunction.DBMultiReadInfo.FilterInfo.FilterData)-1 {
							if len(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType) != 0 && strings.EqualFold(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType, datatypedef.DataTypeDate) {
								selectQueryStr += fmt.Sprintf(`TO_CHAR("%s", 'DDMMYYYY')='%s'`, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue)
							} else {
								selectQueryStr += fmt.Sprintf(`"%s"='%s'`, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue)
							}
						} else {
							if len(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType) != 0 && strings.EqualFold(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType, datatypedef.DataTypeDate) {
								selectQueryStr += fmt.Sprintf(`TO_CHAR("%s", 'DDMMYYYY')='%s' %s `, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterCondition)
							} else {
								selectQueryStr += fmt.Sprintf(`"%s"='%s' %s `, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterCondition)
							}
						}
					} else if execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterOperator == lexicalparserdef.FilterOperatorNotEqual {
						if i == len(execFunction.DBMultiReadInfo.FilterInfo.FilterData)-1 {
							if len(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType) != 0 && strings.EqualFold(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType, datatypedef.DataTypeDate) {
								selectQueryStr += fmt.Sprintf(`TO_CHAR("%s", 'DDMMYYYY')!='%s'`, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue)
							} else {
								selectQueryStr += fmt.Sprintf(`"%s"!='%s'`, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue)
							}
						} else {
							if len(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType) != 0 && strings.EqualFold(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType, datatypedef.DataTypeDate) {
								selectQueryStr += fmt.Sprintf(`TO_CHAR("%s", 'DDMMYYYY')!='%s' %s `, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterCondition)
							} else {
								selectQueryStr += fmt.Sprintf(`"%s"!='%s' %s `, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterCondition)
							}
						}
					} else if execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterOperator == lexicalparserdef.FilterOperatorBetween {
						if i == len(execFunction.DBMultiReadInfo.FilterInfo.FilterData)-1 {
							if len(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType) != 0 && strings.EqualFold(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType, datatypedef.DataTypeDate) {
								selectQueryStr += fmt.Sprintf(`"%s" BETWEEN TO_DATE('%s', 'DD-MM-YY') AND TO_DATE('%s', 'DD-MM-YY')`, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue, fldValue1)
							} else {
								selectQueryStr += fmt.Sprintf(`"%s" BETWEEN '%s' AND '%s'`, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue, fldValue1)
							}
						} else {
							if len(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType) != 0 && strings.EqualFold(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType, datatypedef.DataTypeDate) {
								selectQueryStr += fmt.Sprintf(`"%s" BETWEEN TO_DATE('%s', 'DD-MM-YY') AND TO_DATE('%s', 'DD-MM-YY') %s `, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue, fldValue1, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterCondition)
							} else {
								selectQueryStr += fmt.Sprintf(`"%s" BETWEEN '%s' AND '%s' %s `, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue, fldValue1, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterCondition)
							}
						}
					} else if execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterOperator == lexicalparserdef.FilterOperatorLessThan {
						if i == len(execFunction.DBMultiReadInfo.FilterInfo.FilterData)-1 {
							if len(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType) != 0 && strings.EqualFold(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType, datatypedef.DataTypeDate) {
								selectQueryStr += fmt.Sprintf(`TO_CHAR("%s", 'DDMMYYYY')<'%s'`, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue)
							} else {
								selectQueryStr += fmt.Sprintf(`"%s"<'%s'`, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue)
							}
						} else {
							if len(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType) != 0 && strings.EqualFold(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType, datatypedef.DataTypeDate) {
								selectQueryStr += fmt.Sprintf(`TO_CHAR("%s", 'DDMMYYYY')<'%s' %s `, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterCondition)
							} else {
								selectQueryStr += fmt.Sprintf(`"%s"<'%s' %s `, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterCondition)
							}
						}
					} else if execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterOperator == lexicalparserdef.FilterOperatorLessThanEqual {
						if i == len(execFunction.DBMultiReadInfo.FilterInfo.FilterData)-1 {
							if len(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType) != 0 && strings.EqualFold(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType, datatypedef.DataTypeDate) {
								selectQueryStr += fmt.Sprintf(`TO_CHAR("%s", 'DDMMYYYY')<='%s'`, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue)
							} else {
								selectQueryStr += fmt.Sprintf(`"%s"<='%s'`, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue)
							}
						} else {
							if len(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType) != 0 && strings.EqualFold(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType, datatypedef.DataTypeDate) {
								selectQueryStr += fmt.Sprintf(`TO_CHAR("%s", 'DDMMYYYY')<='%s' %s `, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterCondition)
							} else {
								selectQueryStr += fmt.Sprintf(`"%s"<='%s' %s `, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterCondition)
							}
						}
					} else if execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterOperator == lexicalparserdef.FilterOperatorGreaterThan {
						if i == len(execFunction.DBMultiReadInfo.FilterInfo.FilterData)-1 {
							if len(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType) != 0 && strings.EqualFold(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType, datatypedef.DataTypeDate) {
								selectQueryStr += fmt.Sprintf(`TO_CHAR("%s", 'DDMMYYYY')>'%s'`, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue)
							} else {
								selectQueryStr += fmt.Sprintf(`"%s">'%s'`, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue)
							}
						} else {
							if len(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType) != 0 && strings.EqualFold(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType, datatypedef.DataTypeDate) {
								selectQueryStr += fmt.Sprintf(`TO_CHAR("%s", 'DDMMYYYY')>'%s' %s `, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterCondition)
							} else {
								selectQueryStr += fmt.Sprintf(`"%s">'%s' %s `, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterCondition)
							}
						}
					} else if execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterOperator == lexicalparserdef.FilterOperatorGreaterThanEqual {
						if i == len(execFunction.DBMultiReadInfo.FilterInfo.FilterData)-1 {
							if len(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType) != 0 && strings.EqualFold(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType, datatypedef.DataTypeDate) {
								selectQueryStr += fmt.Sprintf(`TO_CHAR("%s", 'DDMMYYYY')>='%s'`, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue)
							} else {
								selectQueryStr += fmt.Sprintf(`"%s">='%s'`, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue)
							}
						} else {
							if len(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType) != 0 && strings.EqualFold(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType, datatypedef.DataTypeDate) {
								selectQueryStr += fmt.Sprintf(`TO_CHAR("%s", 'DDMMYYYY')>='%s' %s `, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterCondition)
							} else {
								selectQueryStr += fmt.Sprintf(`"%s">='%s' %s `, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterCondition)
							}
						}
					} else if execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterOperator == lexicalparserdef.FilterOperatorLike {
						if i == len(execFunction.DBMultiReadInfo.FilterInfo.FilterData)-1 {
							if len(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType) != 0 && strings.EqualFold(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType, datatypedef.DataTypeDate) {
								selectQueryStr += fmt.Sprintf(`TO_CHAR("%s", 'DD-MM-YY') LIKE '%%%s%%'`, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue)
							} else {
								selectQueryStr += fmt.Sprintf(`"%s" LIKE '%%%s%%'`, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue)
							}
						} else {
							if len(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType) != 0 && strings.EqualFold(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType, datatypedef.DataTypeDate) {
								selectQueryStr += fmt.Sprintf(`TO_CHAR("%s", 'DD-MM-YY') LIKE '%%%s%%' %s `, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterCondition)
							} else {
								selectQueryStr += fmt.Sprintf(
									`LOWER(REGEXP_REPLACE("%s", '\s+', ' ', 'g')) LIKE LOWER(REGEXP_REPLACE('%%%s%%', '\s+', ' ', 'g'))`,
									execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName,
									strings.TrimSpace(fldValue),
								)
							}
							selectQueryStr += " AND "
						}
					} else {
						rejectDesc = fmt.Sprintf("Invalid Filter Operator[%s]", execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterOperator)
						reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_DBMultiReadFTFilterDataErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBMultiReadInfo.AppErrDesc, execFunction.DBMultiReadInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBMultiReadInfo.AppCorrectiveAction, execFunction.DBMultiReadInfo.AppCorrectiveAction), rejectDesc)
						return -1
					}
				}
			} else {
				if len(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterOperator) == 0 {
					//trace.Lg("FilterOperator is NULL")
					if i == len(execFunction.DBMultiReadInfo.FilterInfo.FilterData)-1 {
						if len(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType) != 0 && strings.EqualFold(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType, datatypedef.DataTypeDate) {
							selectQueryStr += fmt.Sprintf(`TO_CHAR("%s", 'DDMMYYYY')='%s'`, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue)
						} else {
							selectQueryStr += fmt.Sprintf(`"%s"='%s'`, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue)
						}
					} else {
						if len(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType) != 0 && strings.EqualFold(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType, datatypedef.DataTypeDate) {
							selectQueryStr += fmt.Sprintf(`TO_CHAR("%s", 'DDMMYYYY')='%s' %s `, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterCondition)
						} else {
							selectQueryStr += fmt.Sprintf(`"%s"='%s' %s `, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterCondition)
						}
					}
				} else {
					// trace.Lg(" Changed FilterOperator1[%s]", execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterOperator)
					if execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterOperator == lexicalparserdef.FilterOperatorEqual {
						if i == len(execFunction.DBMultiReadInfo.FilterInfo.FilterData)-1 {
							if len(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType) != 0 && strings.EqualFold(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType, datatypedef.DataTypeDate) {
								selectQueryStr += fmt.Sprintf(`TO_CHAR("%s", 'DDMMYYYY')='%s'`, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue)
							} else {
								selectQueryStr += fmt.Sprintf(`"%s"='%s'`, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue)
							}
						} else {
							if len(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType) != 0 && strings.EqualFold(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType, datatypedef.DataTypeDate) {
								selectQueryStr += fmt.Sprintf(`TO_CHAR("%s", 'DDMMYYYY')='%s' %s `, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterCondition)
							} else {
								selectQueryStr += fmt.Sprintf(`"%s"='%s' %s `, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterCondition)
							}
						}
					} else if execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterOperator == lexicalparserdef.FilterOperatorNotEqual {
						if i == len(execFunction.DBMultiReadInfo.FilterInfo.FilterData)-1 {
							if len(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType) != 0 && strings.EqualFold(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType, datatypedef.DataTypeDate) {
								selectQueryStr += fmt.Sprintf(`TO_CHAR("%s", 'DDMMYYYY')!='%s'`, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue)
							} else {
								selectQueryStr += fmt.Sprintf(`"%s"!='%s'`, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue)
							}
						} else {
							if len(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType) != 0 && strings.EqualFold(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType, datatypedef.DataTypeDate) {
								selectQueryStr += fmt.Sprintf(`TO_CHAR("%s", 'DDMMYYYY')!='%s' %s `, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterCondition)
							} else {
								selectQueryStr += fmt.Sprintf(`"%s"!='%s' %s `, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterCondition)
							}
						}
					} else if execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterOperator == lexicalparserdef.FilterOperatorBetween {
						if i == len(execFunction.DBMultiReadInfo.FilterInfo.FilterData)-1 {
							if len(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType) != 0 && strings.EqualFold(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType, datatypedef.DataTypeDate) {
								selectQueryStr += fmt.Sprintf(`%s BETWEEN TO_DATE('%s', 'DD-MM-YY') AND TO_DATE('%s', 'DD-MM-YY')`, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue, fldValue1)
							} else {
								selectQueryStr += fmt.Sprintf(`%s BETWEEN '%s' AND '%s'`, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue, fldValue1)
							}
						} else {
							if len(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType) != 0 && strings.EqualFold(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType, datatypedef.DataTypeDate) {
								selectQueryStr += fmt.Sprintf(`%s BETWEEN TO_DATE('%s', 'DD-MM-YY') AND TO_DATE('%s', 'DD-MM-YY') %s`, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue, fldValue1, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterCondition)
							} else {
								selectQueryStr += fmt.Sprintf(`%s BETWEEN '%s' AND '%s' %s`, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue, fldValue1, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterCondition)
							}
						}
					} else if execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterOperator == lexicalparserdef.FilterOperatorLessThan {
						if i == len(execFunction.DBMultiReadInfo.FilterInfo.FilterData)-1 {
							if len(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType) != 0 && strings.EqualFold(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType, datatypedef.DataTypeDate) {
								selectQueryStr += fmt.Sprintf(`TO_CHAR("%s", 'DDMMYYYY')<'%s'`, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue)
							} else {
								selectQueryStr += fmt.Sprintf(`"%s"<'%s'`, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue)
							}
						} else {
							if len(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType) != 0 && strings.EqualFold(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType, datatypedef.DataTypeDate) {
								selectQueryStr += fmt.Sprintf(`TO_CHAR("%s", 'DDMMYYYY')<'%s' %s `, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterCondition)
							} else {
								selectQueryStr += fmt.Sprintf(`"%s"<'%s' %s `, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterCondition)
							}
						}
					} else if execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterOperator == lexicalparserdef.FilterOperatorLessThanEqual {
						if i == len(execFunction.DBMultiReadInfo.FilterInfo.FilterData)-1 {
							if len(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType) != 0 && strings.EqualFold(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType, datatypedef.DataTypeDate) {
								selectQueryStr += fmt.Sprintf(`TO_CHAR("%s", 'DDMMYYYY')<='%s'`, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue)
							} else {
								selectQueryStr += fmt.Sprintf(`"%s"<='%s'`, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue)
							}
						} else {
							if len(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType) != 0 && strings.EqualFold(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType, datatypedef.DataTypeDate) {
								selectQueryStr += fmt.Sprintf(`TO_CHAR("%s", 'DDMMYYYY')<='%s' %s `, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterCondition)
							} else {
								selectQueryStr += fmt.Sprintf(`"%s"<='%s' %s `, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterCondition)
							}
						}
					} else if execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterOperator == lexicalparserdef.FilterOperatorGreaterThan {
						if i == len(execFunction.DBMultiReadInfo.FilterInfo.FilterData)-1 {
							if len(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType) != 0 && strings.EqualFold(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType, datatypedef.DataTypeDate) {
								selectQueryStr += fmt.Sprintf(`TO_CHAR("%s", 'DDMMYYYY')>'%s'`, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue)
							} else {
								selectQueryStr += fmt.Sprintf(`"%s">'%s'`, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue)
							}
						} else {
							if len(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType) != 0 && strings.EqualFold(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType, datatypedef.DataTypeDate) {
								selectQueryStr += fmt.Sprintf(`TO_CHAR("%s", 'DDMMYYYY')>'%s' %s `, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterCondition)
							} else {
								selectQueryStr += fmt.Sprintf(`"%s">'%s' %s `, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterCondition)
							}
						}
					} else if execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterOperator == lexicalparserdef.FilterOperatorGreaterThanEqual {
						if i == len(execFunction.DBMultiReadInfo.FilterInfo.FilterData)-1 {
							if len(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType) != 0 && strings.EqualFold(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType, datatypedef.DataTypeDate) {
								selectQueryStr += fmt.Sprintf(`TO_CHAR("%s", 'DDMMYYYY')>='%s'`, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue)
							} else {
								selectQueryStr += fmt.Sprintf(`"%s">='%s'`, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue)
							}
						} else {
							if len(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType) != 0 && strings.EqualFold(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType, datatypedef.DataTypeDate) {
								selectQueryStr += fmt.Sprintf(`TO_CHAR("%s", 'DDMMYYYY')>='%s' %s `, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterCondition)
							} else {
								selectQueryStr += fmt.Sprintf(`"%s">='%s' %s `, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterCondition)
							}
						}
					} else if execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterOperator == lexicalparserdef.FilterOperatorLike {
						if i == len(execFunction.DBMultiReadInfo.FilterInfo.FilterData)-1 {
							if len(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType) != 0 && strings.EqualFold(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType, datatypedef.DataTypeDate) {
								selectQueryStr += fmt.Sprintf(`TO_CHAR("%s", 'DD-MM-YY') LIKE '%%%s%%'`, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue)
							} else {
								selectQueryStr += fmt.Sprintf(`"%s" LIKE '%%%s%%'`, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue)
							}
						} else {
							if len(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType) != 0 && strings.EqualFold(execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterDBDataType, datatypedef.DataTypeDate) {
								selectQueryStr += fmt.Sprintf(`TO_CHAR("%s", 'DD-MM-YY') LIKE '%%%s%%' %s `, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName, fldValue, execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterCondition)
							} else {
								selectQueryStr += fmt.Sprintf(
									`LOWER(REGEXP_REPLACE("%s", '\s+', ' ', 'g')) LIKE LOWER(REGEXP_REPLACE('%%%s%%', '\s+', ' ', 'g'))`,
									execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterName,
									strings.TrimSpace(fldValue),
								)
							}
							selectQueryStr += " AND "
						}
					} else {
						rejectDesc = fmt.Sprintf("Invalid Filter Operator[%s]", execFunction.DBMultiReadInfo.FilterInfo.FilterData[i].FilterOperator)
						reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_DBMultiReadFTFilterDataErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBMultiReadInfo.AppErrDesc, execFunction.DBMultiReadInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBMultiReadInfo.AppCorrectiveAction, execFunction.DBMultiReadInfo.AppCorrectiveAction), rejectDesc)
						return -1
					}
				}
			}
		}
	}
	// trace.Lg("selectQueryStr[%s]", selectQueryStr)
	selectQueryStr = removeFilterConditionLastOccurrence(selectQueryStr)

	// GLOBAL SEARCH FILTER
	if resultDBContext.DBType == dbdef.DBTypePostgreSQL && len(execFunction.DBMultiReadInfo.SearchStr) > 0 {
		searchVal := lexicalparserutil.ReadValueFrmDataMap(reqBrokerDataMap, execFunction.DBMultiReadInfo.SearchStr, datatypedef.DataTypeString, &rejectDesc)
		if searchVal != nil {
			searchStr := strings.TrimSpace(searchVal.(string))
			if searchStr != "" {
				if !strings.Contains(strings.ToUpper(selectQueryStr), " WHERE ") {
					selectQueryStr += " WHERE "
				} else {
					selectQueryStr += " AND "
				}
				selectQueryStr += fmt.Sprintf(`
				EXISTS (
					SELECT 1
					FROM jsonb_each_text(to_jsonb("%s")) kv
					WHERE kv.value ILIKE '%%%s%%'
				)
			`,
					execFunction.DBMultiReadInfo.TableName,
					searchStr,
				)
			}
		}
	}

	limit := -1
	offset := -1

	limitKey := execFunction.DBMultiReadInfo.Limit
	offsetKey := execFunction.DBMultiReadInfo.Offset
	orderBy := execFunction.DBMultiReadInfo.OrderBy
	sortType := execFunction.DBMultiReadInfo.SortType
	if len(limitKey) > 0 {
		val := lexicalparserutil.ReadValueFrmDataMap(reqBrokerDataMap, limitKey, "string", &rejectDesc)
		if val != nil {
			if i, err := strconv.Atoi(val.(string)); err == nil {
				limit = i
			}
		}

	}
	if len(offsetKey) > 0 {
		val := lexicalparserutil.ReadValueFrmDataMap(reqBrokerDataMap, offsetKey, "string", &rejectDesc)

		if val != nil {
			s := strings.TrimSpace(val.(string))
			if f, err := strconv.ParseFloat(s, 64); err == nil {
				offset = int(f)
			}
		}
	}
	// trace.Lg("OffsetValue Raw = %+v", offset)

	hasPagination := limit > 0

	if hasPagination {

		if len(sortType) == 0 {
			sortType = "DESC"
		}
		if len(orderBy) == 0 {
			orderBy = "UPDATED_DATE"
		}
		baseQuery := selectQueryStr

		// split columns and sort types
		orderColumns := strings.Split(orderBy, ",")
		sortTypes := strings.Split(sortType, ",")

		var orderParts []string
		var orderAliases []string

		for i, col := range orderColumns {
			col = strings.TrimSpace(col)
			if col == "" {
				continue
			}

			dir := "DESC"

			if i < len(sortTypes) {
				st := strings.TrimSpace(strings.ToUpper(sortTypes[i]))
				if st == "ASC" || st == "DESC" {
					dir = st
				}
			}

			orderParts = append(orderParts, fmt.Sprintf(`"%s" %s`, col, dir))

			orderAliases = append(orderAliases, fmt.Sprintf(`%s %s`, strings.ToLower(col), dir))
		}

		orderClause := " ORDER BY " + strings.Join(orderParts, ", ")

		innerQuery := fmt.Sprintf(
			`%s%s LIMIT %d OFFSET %d`,
			baseQuery,
			orderClause,
			limit,
			offset,
		)

		selectQueryStr = fmt.Sprintf(`
WITH data AS (%s),
count_data AS (
    SELECT COUNT(*) AS total_items
    FROM (%s) base
)
SELECT
(
    SELECT json_agg(lower_row ORDER BY %s)::text
    FROM (
        SELECT
            d."UPDATED_DATE" AS updated_date,
            d."UPDATED_TIME" AS updated_time,
            jsonb_object_agg(lower(k), v) AS lower_row
        FROM data d,
             jsonb_each(to_jsonb(d)) AS e(k, v)
        WHERE lower(k) <> 'password_hash'
        GROUP BY d."UPDATED_DATE", d."UPDATED_TIME", d
    ) t
) AS items,
(SELECT total_items FROM count_data) AS total_items;
`,
			innerQuery,
			baseQuery,
			strings.Join(orderAliases, ", "),
		)
	}
	trace.Lg("selectQueryStr[%s]", selectQueryStr)

	if dbtab_tableinfo.LoadFromDBTable(selectQueryStr, &resultMapArray, resultDBContext, &dbErr, &dbRejectReason) < 0 {
		rejectDesc = fmt.Sprintf("MultiTableReadError for Table[%s]Query[%s] with Err[%s]", execFunction.DBMultiReadInfo.TableName, selectQueryStr, dbErr)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_LoadFromDBTableErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBMultiReadInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBMultiReadInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	if len(execFunction.DBMultiReadInfo.DBResultInfo.ResultCode) != 0 && execFunction.DBMultiReadInfo.DBResultInfo.ResultCode == dbdef.DBNoRows {
		if len(resultMapArray) == 0 && execFunction.DBMultiReadInfo.DBResultInfo.Result == "failed" {
			rejectDesc = fmt.Sprintf("No Data Found")
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_LoadFromDBTableErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBMultiReadInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBMultiReadInfo.AppCorrectiveAction), rejectDesc)
			return -1
		}
	}

	if hasPagination {
		if len(resultMapArray) > 0 {

			paginatedResult := resultMapArray[0]

			if itemsVal, ok := paginatedResult["items"]; ok {

				switch v := itemsVal.(type) {

				case []byte:
					fmt.Println("items is []byte")
					var decoded []map[string]interface{}
					if err := json.Unmarshal(v, &decoded); err == nil {
						paginatedResult["items"] = decoded
					}

				case string:
					fmt.Println("items is string")
					var decoded []map[string]interface{}
					if err := json.Unmarshal([]byte(v), &decoded); err == nil {
						paginatedResult["items"] = decoded
					}
				}
			}

			resultMapArray[0] = paginatedResult
		}
		if len(execFunction.DBMultiReadInfo.StoreObjName) != 0 {
			if maputil.StoreObject(reqBrokerDataMap, execFunction.DBMultiReadInfo.StoreObjName, resultMapArray[0]) < 0 {
				rejectDesc = fmt.Sprintf("StoreObject for StoreObjName[%s]", execFunction.DBMultiReadInfo.StoreObjName)
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_StoreObjectErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBMultiReadInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBMultiReadInfo.AppCorrectiveAction), rejectDesc)
				return -1
			}
		}

	} else {
		if len(execFunction.DBMultiReadInfo.StoreObjName) != 0 {
			if maputil.StoreObject(reqBrokerDataMap, execFunction.DBMultiReadInfo.StoreObjName, resultMapArray) < 0 {
				rejectDesc = fmt.Sprintf("StoreObject for StoreObjName[%s]", execFunction.DBMultiReadInfo.StoreObjName)
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_StoreObjectErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBMultiReadInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.DBMultiReadInfo.AppCorrectiveAction), rejectDesc)
				return -1
			}
		}
	}
	return 1
}
