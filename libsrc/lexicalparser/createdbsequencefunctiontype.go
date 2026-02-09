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
	"lmsapieng/libsrc/utils/dbutil"
	"lmsapieng/libsrc/utils/lexicalparserutil"
	"lmsapieng/libsrc/utils/msgutil"
	"lmsapieng/libsrc/utils/schemainfo"
	"lmsapieng/libsrc/utils/sequtil"
	"lmsapieng/libsrc/utils/templateutil"
	"strconv"
	"strings"
)

func CreateDBSequenceFunctionType(execFunction lexicalparserdef.ExecFunctionStruct, reqBrokerDataMap map[string]interface{}) int {
	var rejectDesc string
	var dbErr, dbRejectReason string
	var resultDBContext dbdef.DBContextDef
	resultMap := make(map[string]interface{})
	if len(execFunction.CreateDBSequenceInfo.ObjectName) == 0 {
		rejectDesc = fmt.Sprintf("ObjectName is NULL in CreateDBSequenceInfoFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_CopyObjectFTObjNameNullErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppCorrectiveAction), rejectDesc)
		return -1
	}
	objectMap := make(map[string]interface{})
	if len(execFunction.CreateDBSequenceInfo.TableName) == 0 {
		rejectDesc := fmt.Sprintf("TableName is NULL in CreateDBSequenceInfoFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_CreateDBSequenceFTTableNullErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppCorrectiveAction), rejectDesc)
		return -1
	}
	selectQueryStr := fmt.Sprintf("SELECT MAX(%s) %s FROM %s", strings.ToUpper(execFunction.CreateDBSequenceInfo.SequenceDbData.Key), execFunction.CreateDBSequenceInfo.SequenceDbData.Key, execFunction.CreateDBSequenceInfo.TableName)
	if len(execFunction.CreateDBSequenceInfo.FilterInfo.FilterData) != 0 {
		selectQueryStr += " WHERE "
		for i := 0; i < len(execFunction.CreateDBSequenceInfo.FilterInfo.FilterData); i++ {
			if len(execFunction.CreateDBSequenceInfo.FilterInfo.FilterData[i].FilterName) == 0 {
				rejectDesc = fmt.Sprintf("FilterName is NULL in CreateDBSequenceInfoFunctionType for FunctionName[%s]", execFunction.FunctionName)
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_CreateDBSequenceFTFilterDataErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppCorrectiveAction), rejectDesc)
				return -1
			}
			if len(execFunction.CreateDBSequenceInfo.FilterInfo.FilterData[i].FilterID) == 0 {
				rejectDesc = fmt.Sprintf("FilterID is NULL in DBSingleReadFunctionType for FunctionName[%s]", execFunction.FunctionName)
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_CreateDBSequenceFTFilterDataErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppCorrectiveAction), rejectDesc)
				return -1
			}
			if len(execFunction.CreateDBSequenceInfo.FilterInfo.FilterData[i].FilterDataType) == 0 {
				rejectDesc = fmt.Sprintf("FilterDataType is NULL in DBSingleReadFunctionType for FunctionName[%s]", execFunction.FunctionName)
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_CreateDBSequenceFTFilterDataErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppCorrectiveAction), rejectDesc)
				return -1
			}
			if len(execFunction.CreateDBSequenceInfo.FilterInfo.FilterData[i].FilterCondition) == 0 {
				rejectDesc = fmt.Sprintf("FilterCondition is NULL in DBSingleReadFunctionType for FunctionName[%s]", execFunction.FunctionName)
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_CreateDBSequenceFTFilterDataErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppCorrectiveAction), rejectDesc)
				return -1
			}
			fldValue := ""
			dataValue := lexicalparserutil.ReadValueFrmDataMap(reqBrokerDataMap, execFunction.CreateDBSequenceInfo.FilterInfo.FilterData[i].FilterID, execFunction.CreateDBSequenceInfo.FilterInfo.FilterData[i].FilterDataType, &rejectDesc)
			if dataValue == nil {
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ReadValueErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppCorrectiveAction), rejectDesc)
				return -1
			}
			if strings.EqualFold(execFunction.CreateDBSequenceInfo.FilterInfo.FilterData[i].FilterDataType, datatypedef.DataTypeString) {
				fldValue = dataValue.(string)
			} else if strings.EqualFold(execFunction.CreateDBSequenceInfo.FilterInfo.FilterData[i].FilterDataType, datatypedef.DataTypeInt) {
				fldValue = fmt.Sprintf("%d", dataValue.(int))
			} else {
				rejectDesc = fmt.Sprintf("FilterDataType[%s] not supported for FilterID[%s]", execFunction.CreateDBSequenceInfo.FilterInfo.FilterData[i].FilterDataType, execFunction.CreateDBSequenceInfo.FilterInfo.FilterData[i].FilterID)
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_FldTypeInvalidErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppCorrectiveAction), rejectDesc)
				return -1
			}
			if i == len(execFunction.CreateDBSequenceInfo.FilterInfo.FilterData)-1 {
				selectQueryStr += fmt.Sprintf("%s='%s'", execFunction.CreateDBSequenceInfo.FilterInfo.FilterData[i].FilterName, fldValue)
			} else {
				selectQueryStr += fmt.Sprintf("%s='%s' %s ", execFunction.CreateDBSequenceInfo.FilterInfo.FilterData[i].FilterName, fldValue, execFunction.CreateDBSequenceInfo.FilterInfo.FilterData[i].FilterCondition)
			}
		}
	}
	// trace.Lg("selectQueryStr[%s]", selectQueryStr)
	retval, contextParams := microsv_schema.GetDBContextParams(reqBrokerDataMap, execFunction.CreateDBSequenceInfo.SchemaInfo, &rejectDesc)
	if retval < 0 {
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GetDBContextParamsErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppCorrectiveAction), rejectDesc)
		return -1
	}
	// trace.Lg("contextParams[%s]", contextParams)
	if schemainfo.GetActiveDBContext(reqBrokerDataMap, &resultDBContext, &rejectDesc, contextParams[0], contextParams[1], contextParams[2], contextParams[3], contextParams[4], contextParams[5]) < 0 {
		rejectDesc = fmt.Sprintf("GetActiveDBContext() failed for Module:%s", contextParams)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GetActiveDBContextError, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppCorrectiveAction), rejectDesc)
		return -1
	}
	dbResult := dbtab_tableinfo.ReadFromDBTable(selectQueryStr, resultMap, resultDBContext, &dbErr, &dbRejectReason)
	if dbResult < 0 {
		if !dbutil.IsNoRows(resultDBContext.DBType, dbErr) {
			rejectDesc = fmt.Sprintf("ReadFromDBTable for Table[%s]Query[%s] with Err[%s]", execFunction.CreateDBSequenceInfo.TableName, selectQueryStr, dbErr)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ReadFromDBTableErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppCorrectiveAction), rejectDesc)
			return -1
		}
	}
	_, ok := resultMap[execFunction.CreateDBSequenceInfo.SequenceDbData.Key]
	if !ok {
		rejectDesc = fmt.Sprintf("ResultMap does not have [%s] as Index", execFunction.CreateDBSequenceInfo.SequenceDbData.Key)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_CreateDBSequenceFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppCorrectiveAction), rejectDesc)
		return -1
	}
	fldValue := "2"
	/*
		fldValue := ""
		dataValue := lexicalparserutil.ReadValueFrmDataMap(reqBrokerDataMap, execFunction.CreateDBSequenceInfo.SequenceDbData.SequenceLen, execFunction.CreateDBSequenceInfo.SequenceDbData.DataType, &rejectDesc)
		if dataValue == nil {
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ReadValueErr, rejectDesc)
			return -1
		}
		if strings.EqualFold(execFunction.CreateDBSequenceInfo.SequenceDbData.DataType, datatypedef.DataTypeString) {
			fldValue = dataValue.(string)
		} else if strings.EqualFold(execFunction.CreateDBSequenceInfo.SequenceDbData.DataType, datatypedef.DataTypeInt) {
			fldValue = fmt.Sprintf("%d", dataValue.(int))
		} else {
			rejectDesc = fmt.Sprintf("SequenceDataType[%s] not supported for Key[%s]", execFunction.CreateDBSequenceInfo.SequenceDbData.DataType, execFunction.CreateDBSequenceInfo.SequenceDbData.Key)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_FldTypeInvalidErr, rejectDesc)
			return -1
		}
	*/
	sequenceLen, _ := strconv.Atoi(fldValue)
	currenceSequence := ""
	if resultMap[execFunction.CreateDBSequenceInfo.SequenceDbData.Key] == nil {
		currenceSequence = fmt.Sprintf("%0*d", sequenceLen, 0)
	} else {
		currenceSequence = resultMap[execFunction.CreateDBSequenceInfo.SequenceDbData.Key].(string)
	}
	//trace.Lg("currenceSequence[%s]", currenceSequence)
	retval, nextSequence := sequtil.GetNextSequence(currenceSequence, sequenceLen)
	if retval < 0 {
		rejectDesc = fmt.Sprintf("GetNextSequence() Error for with currenceSequence[%s] sequenceLen[%d]", currenceSequence, sequenceLen)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GenNextSequenceErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppCorrectiveAction), rejectDesc)
		return -1
	}
	resultMap[execFunction.CreateDBSequenceInfo.SequenceDbData.Key] = nextSequence
	_, ok = reqBrokerDataMap[execFunction.CreateDBSequenceInfo.ObjectName]
	if !ok {
		reqBrokerDataMap[execFunction.CreateDBSequenceInfo.ObjectName] = resultMap
	} else {
		subMap, ok := reqBrokerDataMap[execFunction.CreateDBSequenceInfo.ObjectName].(map[string]interface{})
		if !ok {
			rejectDesc := fmt.Sprintf("ObjectName[%s] Type Assertion Failed for FunctionName[%s]", execFunction.CopyObjectInfo.ObjectName, execFunction.FunctionName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_CreateDBSequenceFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppCorrectiveAction), rejectDesc)
			return -1
		}
		for key, value := range objectMap {
			subMap[key] = value
		}
	}
	return 1
}
