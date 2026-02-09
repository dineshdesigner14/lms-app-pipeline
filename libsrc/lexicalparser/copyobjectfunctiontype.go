package lexicalparser

import (
	"fmt"
	"lmsapieng/include/common/authstatusdef"
	"lmsapieng/include/common/datasourcetypedef"
	"lmsapieng/include/common/datatypedef"
	"lmsapieng/include/common/globaldef"
	"lmsapieng/include/common/lexicalparserdef"
	"lmsapieng/include/common/rejectdef"
	"lmsapieng/include/common/reqbrokerdef"
	"lmsapieng/libsrc/microsv/microsv_common"
	"lmsapieng/libsrc/services/services_sequence"
	"lmsapieng/libsrc/utils/dtutil"
	"lmsapieng/libsrc/utils/exprutil"
	"lmsapieng/libsrc/utils/lexicalparserutil"
	"lmsapieng/libsrc/utils/maputil"
	"lmsapieng/libsrc/utils/msgutil"
	"lmsapieng/libsrc/utils/templateutil"
	"strconv"
	"strings"
	"time"
)

func CopyObjectFunctionType(execFunction lexicalparserdef.ExecFunctionStruct, reqBrokerDataMap map[string]interface{}) int {
	var rejectDesc string
	var ok bool

	if len(execFunction.CopyObjectInfo.ObjectName) == 0 {
		rejectDesc = fmt.Sprintf("ObjectName is NULL in CopyObjectFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_CopyObjectFTObjNameNullErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.CopyObjectInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.CopyObjectInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	objectMap := make(map[string]interface{})
	for i := 0; i < len(execFunction.CopyObjectInfo.ObjectData); i++ {
		if len(execFunction.CopyObjectInfo.ObjectData[i].Condition) != 0 {
			if !exprutil.IsExprTrue(reqBrokerDataMap, execFunction.CopyObjectInfo.ObjectData[i].Condition) {
				continue
			}
		}
		if len(execFunction.CopyObjectInfo.ObjectData[i].Key) == 0 {
			rejectDesc = fmt.Sprintf("Key is NULL in CopyObjectFunctionType for FunctionName[%s]", execFunction.FunctionName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_CopyObjectFTObjKeyNullErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.CopyObjectInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.CopyObjectInfo.AppCorrectiveAction), rejectDesc)
			return -1
		}
		if len(execFunction.CopyObjectInfo.ObjectData[i].DataSourceType) == 0 {
			rejectDesc = fmt.Sprintf("DataSourceType is NULL in CopyObjectFunctionType for FunctionName[%s]", execFunction.FunctionName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_CopyObjectFTObjDataSrcTypeNullErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.CopyObjectInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.CopyObjectInfo.AppCorrectiveAction), rejectDesc)
			return -1
		}
		if len(execFunction.CopyObjectInfo.ObjectData[i].DataSource) == 0 {
			rejectDesc = fmt.Sprintf("DataSource is NULL in CopyObjectFunctionType for FunctionName[%s]", execFunction.FunctionName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_CopyObjectFTObjDataSrcNullErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.CopyObjectInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.CopyObjectInfo.AppCorrectiveAction), rejectDesc)
			return -1
		}
		if len(execFunction.CopyObjectInfo.ObjectData[i].DataType) == 0 {
			rejectDesc = fmt.Sprintf("DataType is NULL in CopyObjectFunctionType for FunctionName[%s]", execFunction.FunctionName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_CopyObjectFTObjDataTypeNullErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.CopyObjectInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.CopyObjectInfo.AppCorrectiveAction), rejectDesc)
			return -1
		}
		fldValue := ""
		if strings.EqualFold(execFunction.CopyObjectInfo.ObjectData[i].DataSourceType, datasourcetypedef.DataSourceTypeReqbrokerDataMap) {
			dataValue := lexicalparserutil.ReadValueFrmDataMap(reqBrokerDataMap, execFunction.CopyObjectInfo.ObjectData[i].DataSource, execFunction.CopyObjectInfo.ObjectData[i].DataType, &rejectDesc)
			if dataValue == nil {
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ReadValueErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.CopyObjectInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.CopyObjectInfo.AppCorrectiveAction), rejectDesc)
				return -1
			}
			if strings.EqualFold(execFunction.CopyObjectInfo.ObjectData[i].DataType, datatypedef.DataTypeString) {
				fldValue = dataValue.(string)
			} else if strings.EqualFold(execFunction.CopyObjectInfo.ObjectData[i].DataType, datatypedef.DataTypeInt) {
				fldValue = fmt.Sprintf("%d", dataValue.(int))
			} else if strings.EqualFold(execFunction.CopyObjectInfo.ObjectData[i].DataType, datatypedef.DataTypeBoolean) {
				fldValue = fmt.Sprintf("%v", dataValue.(bool))
			} else {
				rejectDesc := fmt.Sprintf("ObjectDataType[%s] not supported for ObjectData[%s]", execFunction.CopyObjectInfo.ObjectData[i].DataType, execFunction.CopyObjectInfo.ObjectData[i].Key)
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_FldTypeInvalidErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.CopyObjectInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.CopyObjectInfo.AppCorrectiveAction), rejectDesc)
				return -1
			}
		} else if strings.EqualFold(execFunction.CopyObjectInfo.ObjectData[i].DataSourceType, datasourcetypedef.DataSourceTypeComputeStrExpr) {
			if exprutil.EvaulateStrExpr(reqBrokerDataMap, execFunction.CopyObjectInfo.ObjectData[i].DataSource, &fldValue) < 0 {
				rejectDesc := fmt.Sprintf("EvaulateStrExpr() failed for exprValue[%s])", execFunction.CopyObjectInfo.ObjectData[i].DataSource)
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_CopyObjectFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.CopyObjectInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.CopyObjectInfo.AppCorrectiveAction), rejectDesc)
				return -1
			}
			//trace.Lg("EvaulateStrExpr() success for exprValue[%s] with Value[%s]", execFunction.CopyObjectInfo.ObjectData[i].DataSource, fldValue)
		} else if strings.EqualFold(execFunction.CopyObjectInfo.ObjectData[i].DataSourceType, datasourcetypedef.DataSourceTypeKey) {
			if strings.EqualFold(execFunction.CopyObjectInfo.ObjectData[i].DataSource, datasourcetypedef.DataSourceDefGetDate) {
				fldValue = dtutil.GetDate("DDMMYYYY")
			} else if strings.EqualFold(execFunction.CopyObjectInfo.ObjectData[i].DataSource, datasourcetypedef.DataSourceDefGetTime) {
				fldValue = dtutil.GetTime("HHMMSS")
			} else if strings.EqualFold(execFunction.CopyObjectInfo.ObjectData[i].DataSource, datasourcetypedef.DataSourceDefGetTimeStamp) {
				fldValue = time.Now().Format("2006-01-02T15:04:05-07:00")
			} else if strings.EqualFold(execFunction.CopyObjectInfo.ObjectData[i].DataSource, datasourcetypedef.DataSourceDefGetRecordNum) {
				fldValue = dtutil.GetDateTimeVal()
			} else if strings.EqualFold(execFunction.CopyObjectInfo.ObjectData[i].DataSource, datasourcetypedef.DataSourceDefGetNA) {
				fldValue = globaldef.NOT_INITIALIZED
			} else if strings.EqualFold(execFunction.CopyObjectInfo.ObjectData[i].DataSource, datasourcetypedef.DataSourceDefWaitingForAuth) {
				fldValue = authstatusdef.AuthStatusWaitingForAuthorization
			} else if strings.EqualFold(execFunction.CopyObjectInfo.ObjectData[i].DataSource, datasourcetypedef.DataSourceDefActionInsert) {
				fldValue = authstatusdef.AuthActionInsert
			} else if strings.EqualFold(execFunction.CopyObjectInfo.ObjectData[i].DataSource, datasourcetypedef.DataSourceDefGetIntegrityRowCheckSum) {
				fldValue = globaldef.NOT_INITIALIZED
			} else if strings.EqualFold(execFunction.CopyObjectInfo.ObjectData[i].DataSource, datasourcetypedef.DataSourceDefGetDBRecordID) {
				fldValue = services_sequence.GetDBRecordID()
				if len(fldValue) == 0 {
					rejectDesc := "GetDBRecordID() Error"
					reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GetDBRecordIDErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.CopyObjectInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.CopyObjectInfo.AppCorrectiveAction), rejectDesc)
					return -1
				}
			} else if strings.EqualFold(execFunction.CopyObjectInfo.ObjectData[i].DataSource, datasourcetypedef.DataSourceDefGenTxnBatchNum) {
				if microsv_common.GenTxnBatchNum(&fldValue) < 0 {
					rejectDesc := "GenTxnBatchNum() Error"
					reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GetDBRecordIDErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.CopyObjectInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.CopyObjectInfo.AppCorrectiveAction), rejectDesc)
					return -1
				}
			} else if strings.EqualFold(execFunction.CopyObjectInfo.ObjectData[i].DataSource, datasourcetypedef.DataSourceDefGenTxnRecordNum) {
				if microsv_common.GenTxnRecordNum(&fldValue) < 0 {
					rejectDesc := "GenTxnRecordNum() Error"
					reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GetDBRecordIDErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.CopyObjectInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.CopyObjectInfo.AppCorrectiveAction), rejectDesc)
					return -1
				}
			} else if strings.EqualFold(execFunction.CopyObjectInfo.ObjectData[i].DataSource, datasourcetypedef.DataSourceDefGenRRN) {
				if microsv_common.GenRRN(&fldValue) < 0 {
					rejectDesc := "GenRRN() Error"
					reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GetDBRecordIDErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.CopyObjectInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.CopyObjectInfo.AppCorrectiveAction), rejectDesc)
					return -1
				}
			} else if strings.EqualFold(execFunction.CopyObjectInfo.ObjectData[i].DataSource, datasourcetypedef.DataSourceDefGenStan) {
				if microsv_common.GenStan(&fldValue) < 0 {
					rejectDesc := "GenStan() Error"
					reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GetDBRecordIDErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.CopyObjectInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.CopyObjectInfo.AppCorrectiveAction), rejectDesc)
					return -1
				}
			} else if strings.EqualFold(execFunction.CopyObjectInfo.ObjectData[i].DataSource, datasourcetypedef.DataSourceDefGetSchemaName) {
				if microsv_common.GetSchemaName(&fldValue, reqBrokerDataMap) < 0 {
					rejectDesc := "GenStan() Error"
					reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GetDBRecordIDErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.CopyObjectInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.CopyObjectInfo.AppCorrectiveAction), rejectDesc)
					return -1
				}
			} else {
				rejectDesc := fmt.Sprintf("DataSource[%s] not supported for ObjectData[%s]", execFunction.CopyObjectInfo.ObjectData[i].DataSource, execFunction.CopyObjectInfo.ObjectData[i].Key)
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_FldTypeInvalidErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.CopyObjectInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.CopyObjectInfo.AppCorrectiveAction), rejectDesc)
				return -1
			}
		} else if strings.EqualFold(execFunction.CopyObjectInfo.ObjectData[i].DataSourceType, datasourcetypedef.DataSourceTypeRawValue) {
			fldValue = execFunction.CopyObjectInfo.ObjectData[i].DataSource
		} else {
			rejectDesc := fmt.Sprintf("DataSourceType[%s] not supported for ObjectData[%s]", execFunction.CopyObjectInfo.ObjectData[i].DataSourceType, execFunction.CopyObjectInfo.ObjectData[i].Key)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_FldTypeInvalidErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.CopyObjectInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.CopyObjectInfo.AppCorrectiveAction), rejectDesc)
			return -1
		}
		key := execFunction.CopyObjectInfo.ObjectData[i].Key
		if strings.EqualFold(execFunction.CopyObjectInfo.ObjectData[i].DataType, datatypedef.DataTypeBoolean) {

			boolVal, _ := strconv.ParseBool(strings.TrimSpace(fldValue))
			if fldValue == "" {
				rejectDesc := fmt.Sprintf(
					"Empty boolean value for key [%s]",
					key,
				)

				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] =
					msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_FldTypeInvalidErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.CopyObjectInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.CopyObjectInfo.AppCorrectiveAction), rejectDesc)
				return -1
			}

			objectMap[key] = boolVal

		} else {
			objectMap[execFunction.CopyObjectInfo.ObjectData[i].Key] = fldValue
		}
	}
	if lexicalparserutil.IsFldNameArray(execFunction.CopyObjectInfo.ObjectName) {
		objectName := ""
		if lexicalparserutil.ReplaceArrayIndex(reqBrokerDataMap, execFunction.CopyObjectInfo.ObjectName, &objectName) < 0 {
			rejectDesc := fmt.Sprintf("ReplaceArrayIndex() failed for ObjectName[%s]", execFunction.CopyObjectInfo.ObjectName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_CopyObjectFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.CopyObjectInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.CopyObjectInfo.AppCorrectiveAction), rejectDesc)
			return -1
		}
		for key, value := range objectMap {
			lnewObjectName := fmt.Sprintf("%s.%s", objectName, key)
			if maputil.SetValueFromString(reqBrokerDataMap, lnewObjectName, value) < 0 {
				rejectDesc := fmt.Sprintf("SetValueFromString() failed for ObjectName[%s]", execFunction.CopyObjectInfo.ObjectName)
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_CopyObjectFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.CopyObjectInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.CopyObjectInfo.AppCorrectiveAction), rejectDesc)
				return -1
			}
		}
		return 1
	}
	if strings.Contains(execFunction.CopyObjectInfo.ObjectName, ".") {
		for key, value := range objectMap {
			lnewObjectName := fmt.Sprintf("%s.%s", execFunction.CopyObjectInfo.ObjectName, key)
			if maputil.SetValueFromString(reqBrokerDataMap, lnewObjectName, value) < 0 {
				rejectDesc := fmt.Sprintf("SetValueFromString() failed for ObjectName[%s]", execFunction.CopyObjectInfo.ObjectName)
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_CopyObjectFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.CopyObjectInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.CopyObjectInfo.AppCorrectiveAction), rejectDesc)
				return -1
			}
		}
		return 1
	}
	_, ok = reqBrokerDataMap[execFunction.CopyObjectInfo.ObjectName]
	if !ok {
		reqBrokerDataMap[execFunction.CopyObjectInfo.ObjectName] = objectMap
	} else {
		subMap, ok := reqBrokerDataMap[execFunction.CopyObjectInfo.ObjectName].(map[string]interface{})
		if !ok {
			rejectDesc := fmt.Sprintf("ObjectName[%s] Type Assertion Failed for FunctionName[%s]", execFunction.CopyObjectInfo.ObjectName, execFunction.FunctionName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_CopyObjectFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.CopyObjectInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.CopyObjectInfo.AppCorrectiveAction), rejectDesc)
			return -1
		}
		for key, value := range objectMap {
			subMap[key] = value
		}
	}
	// jsonBytes, err := json.MarshalIndent(objectMap, "", "  ")
	// if err != nil {
	// 	fmt.Println("Error marshalling map:", err)
	// 	return -1
	// }
	// trace.Lg("string(jsonBytes):%s", string(jsonBytes))
	return 1

}
