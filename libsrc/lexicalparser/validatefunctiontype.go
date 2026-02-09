package lexicalparser

import (
	"fmt"
	"lmsapieng/include/common/datasourcetypedef"
	"lmsapieng/include/common/datatypedef"
	"lmsapieng/include/common/fielddef"
	"lmsapieng/include/common/lexicalparserdef"
	"lmsapieng/include/common/rejectdef"
	"lmsapieng/include/common/reqbrokerdef"
	"lmsapieng/libsrc/utils/fldvalidateutil"
	"lmsapieng/libsrc/utils/lexicalparserutil"
	"lmsapieng/libsrc/utils/maputil"
	"lmsapieng/libsrc/utils/msgutil"
	"lmsapieng/libsrc/utils/templateutil"
	"strings"
)

func ValidateFunctionType(execFunction lexicalparserdef.ExecFunctionStruct, reqBrokerDataMap map[string]interface{}) int {
	var fldValue, rejectDesc string
	for i := 0; i < len(execFunction.ValidateInfo.ValidateFld); i++ {
		if len(execFunction.ValidateInfo.ValidateFld[i].FldPresent) != 0 {
			if execFunction.ValidateInfo.ValidateFld[i].FldPresent == fielddef.FldPresentOptional {
				var dataValue interface{}
				dataValue = lexicalparserutil.ReadValueFrmDataMap(reqBrokerDataMap, execFunction.ValidateInfo.ValidateFld[i].FldName, execFunction.ValidateInfo.ValidateFld[i].FldDataType, &rejectDesc)
				if dataValue == nil {
					if execFunction.ValidateInfo.ValidateFld[i].FldDefaultSource == datasourcetypedef.DataSourceTypeReqbrokerDataMap {
						dataValue = lexicalparserutil.ReadValueFrmDataMap(reqBrokerDataMap, execFunction.ValidateInfo.ValidateFld[i].FldDefaultValue, execFunction.ValidateInfo.ValidateFld[i].FldDataType, &rejectDesc)
						if dataValue == nil {
							reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ReadValueErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ValidateInfo.AppErrDesc, execFunction.ValidateInfo.ValidateFld[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ValidateInfo.AppCorrectiveAction, execFunction.ValidateInfo.ValidateFld[i].AppCorrectiveAction), rejectDesc)
							return -1
						}
						//trace.Lg("dataValue[%s] type is [%T]", execFunction.ValidateInfo.ValidateFld[i].FldDefaultValue, dataValue)
						if strings.EqualFold(execFunction.ValidateInfo.ValidateFld[i].FldDataType, datatypedef.DataTypeString) {
							fldValue = dataValue.(string)
						} else if strings.EqualFold(execFunction.ValidateInfo.ValidateFld[i].FldDataType, datatypedef.DataTypeInt) {
							fldValue = fmt.Sprintf("%d", dataValue.(int))
						} else if strings.EqualFold(execFunction.ValidateInfo.ValidateFld[i].FldDataType, datatypedef.DataTypeBoolean) {
							fldValue = fmt.Sprintf("%t", dataValue.(bool))
						} else {
							rejectDesc = fmt.Sprintf("FldType[%s] not supported for FldName[%s]", execFunction.ValidateInfo.ValidateFld[i].FldDataType, execFunction.ValidateInfo.ValidateFld[i].FldName)
							reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_FldTypeInvalidErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ValidateInfo.AppErrDesc, execFunction.ValidateInfo.ValidateFld[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ValidateInfo.AppCorrectiveAction, execFunction.ValidateInfo.ValidateFld[i].AppCorrectiveAction), rejectDesc)
							return -1
						}
					} else {
						fldValue = execFunction.ValidateInfo.ValidateFld[i].FldDefaultValue
					}
					maputil.SetValueFromString(reqBrokerDataMap, execFunction.ValidateInfo.ValidateFld[i].FldName, fldValue)
					//ReqBrokerReqObj.request_data.secondary_mobile_num = fldValue
				} else {
					if strings.EqualFold(execFunction.ValidateInfo.ValidateFld[i].FldDataType, datatypedef.DataTypeString) {
						fldValue = dataValue.(string)
					} else if strings.EqualFold(execFunction.ValidateInfo.ValidateFld[i].FldDataType, datatypedef.DataTypeInt) {
						fldValue = fmt.Sprintf("%d", dataValue.(int))
					} else if strings.EqualFold(execFunction.ValidateInfo.ValidateFld[i].FldDataType, datatypedef.DataTypeBoolean) {
						fldValue = fmt.Sprintf("%t", dataValue.(bool))
					} else {
						rejectDesc = fmt.Sprintf("FldType[%s] not supported for FldName[%s]", execFunction.ValidateInfo.ValidateFld[i].FldDataType, execFunction.ValidateInfo.ValidateFld[i].FldName)
						reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_FldTypeInvalidErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ValidateInfo.AppErrDesc, execFunction.ValidateInfo.ValidateFld[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ValidateInfo.AppCorrectiveAction, execFunction.ValidateInfo.ValidateFld[i].AppCorrectiveAction), rejectDesc)
						return -1
					}
					if len(fldValue) == 0 {
						if execFunction.ValidateInfo.ValidateFld[i].FldDefaultSource == datasourcetypedef.DataSourceTypeReqbrokerDataMap {
							dataValue = lexicalparserutil.ReadValueFrmDataMap(reqBrokerDataMap, execFunction.ValidateInfo.ValidateFld[i].FldDefaultValue, execFunction.ValidateInfo.ValidateFld[i].FldDataType, &rejectDesc)
							if dataValue == nil {
								reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ReadValueErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ValidateInfo.AppErrDesc, execFunction.ValidateInfo.ValidateFld[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ValidateInfo.AppCorrectiveAction, execFunction.ValidateInfo.ValidateFld[i].AppCorrectiveAction), rejectDesc)
								return -1
							}
							//trace.Lg("dataValue[%s] type is [%T]", execFunction.ValidateInfo.ValidateFld[i].FldDefaultValue, dataValue)
							if strings.EqualFold(execFunction.ValidateInfo.ValidateFld[i].FldDataType, datatypedef.DataTypeString) {
								fldValue = dataValue.(string)
							} else if strings.EqualFold(execFunction.ValidateInfo.ValidateFld[i].FldDataType, datatypedef.DataTypeInt) {
								fldValue = fmt.Sprintf("%d", dataValue.(int))
							} else if strings.EqualFold(execFunction.ValidateInfo.ValidateFld[i].FldDataType, datatypedef.DataTypeBoolean) {
								fldValue = fmt.Sprintf("%t", dataValue.(bool))
							} else {
								rejectDesc = fmt.Sprintf("FldType[%s] not supported for FldName[%s]", execFunction.ValidateInfo.ValidateFld[i].FldDataType, execFunction.ValidateInfo.ValidateFld[i].FldName)
								reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_FldTypeInvalidErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ValidateInfo.AppErrDesc, execFunction.ValidateInfo.ValidateFld[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ValidateInfo.AppCorrectiveAction, execFunction.ValidateInfo.ValidateFld[i].AppCorrectiveAction), rejectDesc)
								return -1
							}
						} else {
							fldValue = execFunction.ValidateInfo.ValidateFld[i].FldDefaultValue
						}
						maputil.SetValueFromString(reqBrokerDataMap, execFunction.ValidateInfo.ValidateFld[i].FldName, fldValue)
					}
				}
			}
		}
		dataValue := lexicalparserutil.ReadValueFrmDataMap(reqBrokerDataMap, execFunction.ValidateInfo.ValidateFld[i].FldName, execFunction.ValidateInfo.ValidateFld[i].FldDataType, &rejectDesc)
		if dataValue == nil {
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ReadValueErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ValidateInfo.AppErrDesc, execFunction.ValidateInfo.ValidateFld[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ValidateInfo.AppCorrectiveAction, execFunction.ValidateInfo.ValidateFld[i].AppCorrectiveAction), rejectDesc)
			return -1
		}
		if strings.EqualFold(execFunction.ValidateInfo.ValidateFld[i].FldDataType, datatypedef.DataTypeString) {
			fldValue = dataValue.(string)
		} else if strings.EqualFold(execFunction.ValidateInfo.ValidateFld[i].FldDataType, datatypedef.DataTypeInt) {
			fldValue = fmt.Sprintf("%d", dataValue.(int))
		} else if strings.EqualFold(execFunction.ValidateInfo.ValidateFld[i].FldDataType, datatypedef.DataTypeBoolean) {
			fldValue = fmt.Sprintf("%t", dataValue.(bool))
		} else {
			rejectDesc = fmt.Sprintf("FldType[%s] not supported for FldName[%s]", execFunction.ValidateInfo.ValidateFld[i].FldDataType, execFunction.ValidateInfo.ValidateFld[i].FldName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_FldTypeInvalidErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ValidateInfo.AppErrDesc, execFunction.ValidateInfo.ValidateFld[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ValidateInfo.AppCorrectiveAction, execFunction.ValidateInfo.ValidateFld[i].AppCorrectiveAction), rejectDesc)
			return -1
		}
		if fldvalidateutil.ValidateFld(fldValue, execFunction.ValidateInfo.ValidateFld[i], &rejectDesc) < 0 {
			//trace.Lg("ValidateFld() failed with rejectDesc[%s]", rejectDesc)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_FldValidationErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ValidateInfo.AppErrDesc, execFunction.ValidateInfo.ValidateFld[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ValidateInfo.AppCorrectiveAction, execFunction.ValidateInfo.ValidateFld[i].AppCorrectiveAction), rejectDesc)
			return -1
		}
	}
	return 1
}
