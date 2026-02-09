package lexicalparser

import (
	"fmt"
	"lmsapieng/include/common/datasourcetypedef"
	"lmsapieng/include/common/datatypedef"
	"lmsapieng/include/common/lexicalparserdef"
	"lmsapieng/include/common/rejectdef"
	"lmsapieng/include/common/reqbrokerdef"
	"lmsapieng/libsrc/utils/exprutil"
	"lmsapieng/libsrc/utils/lexicalparserutil"
	"lmsapieng/libsrc/utils/msgutil"
	"lmsapieng/libsrc/utils/templateutil"
	"strings"
)

func CreateObjectFunctionType(execFunction lexicalparserdef.ExecFunctionStruct, reqBrokerDataMap map[string]interface{}) int {
	var rejectDesc string
	var ok bool
	if len(execFunction.CreateObjectInfo.ObjectName) == 0 {
		rejectDesc = fmt.Sprintf("ObjectName is NULL in CreateObjectFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_CreateObjectFTObjNameNullErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.CreateObjectInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.CreateObjectInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	objectMap := make(map[string]interface{})
	fldValue := ""
	for i := 0; i < len(execFunction.CreateObjectInfo.ObjectData); i++ {
		if len(execFunction.CreateObjectInfo.ObjectData[i].Condition) != 0 {
			if !exprutil.IsExprTrue(reqBrokerDataMap, execFunction.CreateObjectInfo.ObjectData[i].Condition) {
				continue
			}
		}
		if len(execFunction.CreateObjectInfo.ObjectData[i].Key) == 0 {
			rejectDesc = fmt.Sprintf("Key is NULL in CreateObjectFunctionType for FunctionName[%s]", execFunction.FunctionName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_CreateObjectFTObjKeyNullErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.CreateObjectInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.CreateObjectInfo.AppCorrectiveAction), rejectDesc)
			return -1
		}
		if len(execFunction.CreateObjectInfo.ObjectData[i].DataSource) == 0 {
			rejectDesc = fmt.Sprintf("DataSource is NULL in CreateObjectFunctionType for FunctionName[%s]", execFunction.FunctionName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_CreateObjectFTObjDataSrcNullErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.CreateObjectInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.CreateObjectInfo.AppCorrectiveAction), rejectDesc)
			return -1
		}
		if len(execFunction.CreateObjectInfo.ObjectData[i].DataType) == 0 {
			rejectDesc = fmt.Sprintf("DataType is NULL in CreateObjectFunctionType for FunctionName[%s]", execFunction.FunctionName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_CreateObjectFTObjDataTypeNullErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.CreateObjectInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.CreateObjectInfo.AppCorrectiveAction), rejectDesc)
			return -1
		}
		if strings.EqualFold(execFunction.CreateObjectInfo.ObjectData[i].DataSourceType, datasourcetypedef.DataSourceTypeRawValue) {
			fldValue = execFunction.CreateObjectInfo.ObjectData[i].DataSource
		} else {
			dataValue := lexicalparserutil.ReadValueFrmDataMap(reqBrokerDataMap, execFunction.CreateObjectInfo.ObjectData[i].DataSource, execFunction.CreateObjectInfo.ObjectData[i].DataType, &rejectDesc)
			if dataValue == nil {
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ReadValueErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.CreateObjectInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.CreateObjectInfo.AppCorrectiveAction), rejectDesc)
				return -1
			}
			if strings.EqualFold(execFunction.CreateObjectInfo.ObjectData[i].DataType, datatypedef.DataTypeString) {
				fldValue = dataValue.(string)
			} else if strings.EqualFold(execFunction.CreateObjectInfo.ObjectData[i].DataType, datatypedef.DataTypeInt) {
				fldValue = fmt.Sprintf("%d", dataValue.(int))
			} else {
				rejectDesc := fmt.Sprintf("ObjectDataType[%s] not supported for ObjectData[%s]", execFunction.CreateObjectInfo.ObjectData[i].DataType, execFunction.CreateObjectInfo.ObjectData[i].Key)
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_FldTypeInvalidErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.CreateObjectInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.CreateObjectInfo.AppCorrectiveAction), rejectDesc)
				return -1
			}
		}
		objectMap[execFunction.CreateObjectInfo.ObjectData[i].Key] = fldValue
	}
	_, ok = reqBrokerDataMap[execFunction.CreateObjectInfo.ObjectName]
	if !ok {
		reqBrokerDataMap[execFunction.CreateObjectInfo.ObjectName] = objectMap
	} else {
		subMap, ok := reqBrokerDataMap[execFunction.CreateObjectInfo.ObjectName].(map[string]interface{})
		if !ok {
			rejectDesc := fmt.Sprintf("ObjectName[%s] Type Assertion Failed for FunctionName[%s]", execFunction.CreateObjectInfo.ObjectName, execFunction.FunctionName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_CreateObjectFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.CreateObjectInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.CreateObjectInfo.AppCorrectiveAction), rejectDesc)
			return -1
		}
		for key, value := range objectMap {
			subMap[key] = value
		}
	}
	return 1
}
