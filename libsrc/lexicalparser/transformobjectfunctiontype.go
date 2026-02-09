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
	"lmsapieng/include/common/transformobjdef"
	"lmsapieng/libsrc/services/services_security"
	"lmsapieng/libsrc/services/services_sequence"
	"lmsapieng/libsrc/utils/cardutil"
	"lmsapieng/libsrc/utils/dtutil"
	"lmsapieng/libsrc/utils/exprutil"
	"lmsapieng/libsrc/utils/lexicalparserutil"
	"lmsapieng/libsrc/utils/maskutil"
	"lmsapieng/libsrc/utils/msgutil"
	"lmsapieng/libsrc/utils/templateutil"
	"strconv"
	"strings"
)

func TransformObjectFunctionType(execFunction lexicalparserdef.ExecFunctionStruct, reqBrokerDataMap map[string]interface{}) int {
	var rejectDesc string
	var ok bool

	if len(execFunction.TransformObjectInfo.ObjectName) == 0 {
		rejectDesc = fmt.Sprintf("ObjectName is NULL in TransformObjectFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_TransformObjectFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.TransformObjectInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.TransformObjectInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	objectMap := make(map[string]interface{})
	for i := 0; i < len(execFunction.TransformObjectInfo.ObjectData); i++ {
		if len(execFunction.TransformObjectInfo.ObjectData[i].Condition) != 0 {
			if !exprutil.IsExprTrue(reqBrokerDataMap, execFunction.TransformObjectInfo.ObjectData[i].Condition) {
				continue
			}
		}
		if len(execFunction.TransformObjectInfo.ObjectData[i].Algo) == 0 {
			rejectDesc = fmt.Sprintf("Algo is NULL in TransformObjectFunctionType for FunctionName[%s]", execFunction.FunctionName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_TransformObjectFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.TransformObjectInfo.AppErrDesc, execFunction.TransformObjectInfo.ObjectData[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.TransformObjectInfo.AppCorrectiveAction, execFunction.TransformObjectInfo.ObjectData[i].AppCorrectiveAction), rejectDesc)
			return -1
		}
		if len(execFunction.TransformObjectInfo.ObjectData[i].Key) == 0 {
			rejectDesc = fmt.Sprintf("Key is NULL in TransformObjectFunctionType for FunctionName[%s]", execFunction.FunctionName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_TransformObjectFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.TransformObjectInfo.AppErrDesc, execFunction.TransformObjectInfo.ObjectData[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.TransformObjectInfo.AppCorrectiveAction, execFunction.TransformObjectInfo.ObjectData[i].AppCorrectiveAction), rejectDesc)
			return -1
		}
		if len(execFunction.TransformObjectInfo.ObjectData[i].DataSourceType) == 0 {
			rejectDesc = fmt.Sprintf("DataSourceType is NULL in TransformObjectFunctionType for FunctionName[%s]", execFunction.FunctionName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_TransformObjectFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.TransformObjectInfo.AppErrDesc, execFunction.TransformObjectInfo.ObjectData[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.TransformObjectInfo.AppCorrectiveAction, execFunction.TransformObjectInfo.ObjectData[i].AppCorrectiveAction), rejectDesc)
			return -1
		}
		if len(execFunction.TransformObjectInfo.ObjectData[i].DataSource) == 0 {
			rejectDesc = fmt.Sprintf("DataSource is NULL in TransformObjectFunctionType for FunctionName[%s]", execFunction.FunctionName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_TransformObjectFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.TransformObjectInfo.AppErrDesc, execFunction.TransformObjectInfo.ObjectData[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.TransformObjectInfo.AppCorrectiveAction, execFunction.TransformObjectInfo.ObjectData[i].AppCorrectiveAction), rejectDesc)
			return -1
		}
		if len(execFunction.TransformObjectInfo.ObjectData[i].DataType) == 0 {
			rejectDesc = fmt.Sprintf("DataType is NULL in TransformObjectFunctionType for FunctionName[%s]", execFunction.FunctionName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_TransformObjectFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.TransformObjectInfo.AppErrDesc, execFunction.TransformObjectInfo.ObjectData[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.TransformObjectInfo.AppCorrectiveAction, execFunction.TransformObjectInfo.ObjectData[i].AppCorrectiveAction), rejectDesc)
			return -1
		}
		fldValue := ""
		if strings.EqualFold(execFunction.TransformObjectInfo.ObjectData[i].DataSourceType, datasourcetypedef.DataSourceTypeReqbrokerDataMap) {
			dataValue := lexicalparserutil.ReadValueFrmDataMap(reqBrokerDataMap, execFunction.TransformObjectInfo.ObjectData[i].DataSource, execFunction.TransformObjectInfo.ObjectData[i].DataType, &rejectDesc)
			if dataValue == nil {
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ReadValueErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.TransformObjectInfo.AppErrDesc, execFunction.TransformObjectInfo.ObjectData[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.TransformObjectInfo.AppCorrectiveAction, execFunction.TransformObjectInfo.ObjectData[i].AppCorrectiveAction), rejectDesc)
				return -1
			}
			if strings.EqualFold(execFunction.TransformObjectInfo.ObjectData[i].DataType, datatypedef.DataTypeString) {
				fldValue = dataValue.(string)
			} else if strings.EqualFold(execFunction.TransformObjectInfo.ObjectData[i].DataType, datatypedef.DataTypeInt) {
				fldValue = fmt.Sprintf("%d", dataValue.(int))
			} else {
				rejectDesc := fmt.Sprintf("ObjectDataType[%s] not supported for ObjectData[%s]", execFunction.TransformObjectInfo.ObjectData[i].DataType, execFunction.TransformObjectInfo.ObjectData[i].Key)
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_FldTypeInvalidErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.TransformObjectInfo.AppErrDesc, execFunction.TransformObjectInfo.ObjectData[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.TransformObjectInfo.AppCorrectiveAction, execFunction.TransformObjectInfo.ObjectData[i].AppCorrectiveAction), rejectDesc)
				return -1
			}
		} else if strings.EqualFold(execFunction.TransformObjectInfo.ObjectData[i].DataSourceType, datasourcetypedef.DataSourceTypeComputeStrExpr) {
			if exprutil.EvaulateStrExpr(reqBrokerDataMap, execFunction.TransformObjectInfo.ObjectData[i].DataSource, &fldValue) < 0 {
				rejectDesc := fmt.Sprintf("EvaulateStrExpr() failed for exprValue[%s])", execFunction.TransformObjectInfo.ObjectData[i].DataSource)
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_TransformObjectFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.TransformObjectInfo.AppErrDesc, execFunction.TransformObjectInfo.ObjectData[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.TransformObjectInfo.AppCorrectiveAction, execFunction.TransformObjectInfo.ObjectData[i].AppCorrectiveAction), rejectDesc)
				return -1
			}
			//trace.Lg("EvaulateStrExpr() success for exprValue[%s] with Value[%s]", execFunction.TransformObjectInfo.ObjectData[i].DataSource, fldValue)
		} else if strings.EqualFold(execFunction.TransformObjectInfo.ObjectData[i].DataSourceType, datasourcetypedef.DataSourceTypeKey) {
			if strings.EqualFold(execFunction.TransformObjectInfo.ObjectData[i].DataSource, datasourcetypedef.DataSourceDefGetDate) {
				fldValue = dtutil.GetDate("DDMMYYYY")
			} else if strings.EqualFold(execFunction.TransformObjectInfo.ObjectData[i].DataSource, datasourcetypedef.DataSourceDefGetTime) {
				fldValue = dtutil.GetTime("HHMMSS")
			} else if strings.EqualFold(execFunction.TransformObjectInfo.ObjectData[i].DataSource, datasourcetypedef.DataSourceDefGetRecordNum) {
				fldValue = dtutil.GetDateTimeVal()
			} else if strings.EqualFold(execFunction.TransformObjectInfo.ObjectData[i].DataSource, datasourcetypedef.DataSourceDefGetNA) {
				fldValue = globaldef.NOT_INITIALIZED
			} else if strings.EqualFold(execFunction.TransformObjectInfo.ObjectData[i].DataSource, datasourcetypedef.DataSourceDefWaitingForAuth) {
				fldValue = authstatusdef.AuthStatusWaitingForAuthorization
			} else if strings.EqualFold(execFunction.TransformObjectInfo.ObjectData[i].DataSource, datasourcetypedef.DataSourceDefActionInsert) {
				fldValue = authstatusdef.AuthActionInsert
			} else if strings.EqualFold(execFunction.TransformObjectInfo.ObjectData[i].DataSource, datasourcetypedef.DataSourceDefGetIntegrityRowCheckSum) {
				fldValue = globaldef.NOT_INITIALIZED
			} else if strings.EqualFold(execFunction.TransformObjectInfo.ObjectData[i].DataSource, datasourcetypedef.DataSourceDefGetDBRecordID) {
				fldValue = services_sequence.GetDBRecordID()
				if len(fldValue) == 0 {
					rejectDesc := "GetDBRecordID() Error"
					reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GetDBRecordIDErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.TransformObjectInfo.AppErrDesc, execFunction.TransformObjectInfo.ObjectData[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.TransformObjectInfo.AppCorrectiveAction, execFunction.TransformObjectInfo.ObjectData[i].AppCorrectiveAction), rejectDesc)
					return -1
				}
			} else {
				rejectDesc := fmt.Sprintf("DataSource[%s] not supported for ObjectData[%s]", execFunction.TransformObjectInfo.ObjectData[i].DataSource, execFunction.TransformObjectInfo.ObjectData[i].Key)
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_FldTypeInvalidErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.TransformObjectInfo.AppErrDesc, execFunction.TransformObjectInfo.ObjectData[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.TransformObjectInfo.AppCorrectiveAction, execFunction.TransformObjectInfo.ObjectData[i].AppCorrectiveAction), rejectDesc)
				return -1
			}
		} else if strings.EqualFold(execFunction.TransformObjectInfo.ObjectData[i].DataSourceType, datasourcetypedef.DataSourceTypeRawValue) {
			fldValue = execFunction.TransformObjectInfo.ObjectData[i].DataSource
		} else {
			rejectDesc := fmt.Sprintf("DataSourceType[%s] not supported for ObjectData[%s]", execFunction.TransformObjectInfo.ObjectData[i].DataSourceType, execFunction.TransformObjectInfo.ObjectData[i].Key)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_FldTypeInvalidErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.TransformObjectInfo.AppErrDesc, execFunction.TransformObjectInfo.ObjectData[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.TransformObjectInfo.AppCorrectiveAction, execFunction.TransformObjectInfo.ObjectData[i].AppCorrectiveAction), rejectDesc)
			return -1
		}

		if execFunction.TransformObjectInfo.ObjectData[i].Algo == transformobjdef.TransformObjectAlgoMask {
			if len(execFunction.TransformObjectInfo.ObjectData[i].StartLen) == 0 {
				rejectDesc = fmt.Sprintf("StartLen is NULL in TransformObjectFunctionType for FunctionName[%s]", execFunction.FunctionName)
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_TransformObjectFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.TransformObjectInfo.AppErrDesc, execFunction.TransformObjectInfo.ObjectData[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.TransformObjectInfo.AppCorrectiveAction, execFunction.TransformObjectInfo.ObjectData[i].AppCorrectiveAction), rejectDesc)
				return -1
			}
			if len(execFunction.TransformObjectInfo.ObjectData[i].EndLen) == 0 {
				rejectDesc = fmt.Sprintf("EndLen is NULL in TransformObjectFunctionType for FunctionName[%s]", execFunction.FunctionName)
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_TransformObjectFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.TransformObjectInfo.AppErrDesc, execFunction.TransformObjectInfo.ObjectData[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.TransformObjectInfo.AppCorrectiveAction, execFunction.TransformObjectInfo.ObjectData[i].AppCorrectiveAction), rejectDesc)
				return -1
			}
			startlen := 0
			if exprutil.EvaulateIntExpr(reqBrokerDataMap, execFunction.TransformObjectInfo.ObjectData[i].StartLen, &startlen) < 0 {
				rejectDesc = fmt.Sprintf("EvaulateIntExpr() failed for Expr[%s] in TransformObjectFunctionType for FunctionName[%s]", execFunction.TransformObjectInfo.ObjectData[i].DataSource, execFunction.FunctionName)
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_TransformObjectFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.TransformObjectInfo.AppErrDesc, execFunction.TransformObjectInfo.ObjectData[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.TransformObjectInfo.AppCorrectiveAction, execFunction.TransformObjectInfo.ObjectData[i].AppCorrectiveAction), rejectDesc)
				return -1
			}
			endlen := 0
			if exprutil.EvaulateIntExpr(reqBrokerDataMap, execFunction.TransformObjectInfo.ObjectData[i].EndLen, &endlen) < 0 {
				rejectDesc = fmt.Sprintf("EvaulateIntExpr() failed for Expr[%s] in TransformObjectFunctionType for FunctionName[%s]", execFunction.TransformObjectInfo.ObjectData[i].DataSource, execFunction.FunctionName)
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_TransformObjectFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.TransformObjectInfo.AppErrDesc, execFunction.TransformObjectInfo.ObjectData[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.TransformObjectInfo.AppCorrectiveAction, execFunction.TransformObjectInfo.ObjectData[i].AppCorrectiveAction), rejectDesc)
				return -1
			}
			fldValue = maskutil.GetMaskStr(fldValue, startlen, endlen)
		} else if execFunction.TransformObjectInfo.ObjectData[i].Algo == transformobjdef.TransformObjectAlgExpiryDate {
			expiryPeriod, err := strconv.Atoi(fldValue)
			if err != nil {
				rejectDesc = fmt.Sprintf("ExpiryPeriod[%s] should be numeric and should be in years in TransformObjectFunctionType for FunctionName[%s]", execFunction.TransformObjectInfo.ObjectData[i].DataSource, execFunction.FunctionName)
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_TransformObjectFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.TransformObjectInfo.AppErrDesc, execFunction.TransformObjectInfo.ObjectData[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.TransformObjectInfo.AppCorrectiveAction, execFunction.TransformObjectInfo.ObjectData[i].AppCorrectiveAction), rejectDesc)
				return -1
			}
			fldValue = cardutil.GetExpiryDate(expiryPeriod)
			if len(fldValue) == 0 {
				rejectDesc = fmt.Sprintf("GetExpiryDate() failed in TransformObjectFunctionType for FunctionName[%s]", execFunction.FunctionName)
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_TransformObjectFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.TransformObjectInfo.AppErrDesc, execFunction.TransformObjectInfo.ObjectData[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.TransformObjectInfo.AppCorrectiveAction, execFunction.TransformObjectInfo.ObjectData[i].AppCorrectiveAction), rejectDesc)
				return -1
			}
		} else if execFunction.TransformObjectInfo.ObjectData[i].Algo == transformobjdef.TransformObjectAlgoLuhns {
			fldValue += cardutil.GenLuhnCheckDigit(fldValue)
		} else if execFunction.TransformObjectInfo.ObjectData[i].Algo == transformobjdef.TransformObjectAlgoHash {
			fldValue = cardutil.GetHashCardNum(fldValue)
			if len(fldValue) == 0 {
				rejectDesc = fmt.Sprintf("GetHashCardNum() failed in TransformObjectFunctionType for FunctionName[%s]", execFunction.FunctionName)
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_TransformObjectFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.TransformObjectInfo.AppErrDesc, execFunction.TransformObjectInfo.ObjectData[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.TransformObjectInfo.AppCorrectiveAction, execFunction.TransformObjectInfo.ObjectData[i].AppCorrectiveAction), rejectDesc)
				return -1
			}
		} else if execFunction.TransformObjectInfo.ObjectData[i].Algo == transformobjdef.TransformObjectAlgoEncrypt {
			eData := ""
			retval, _ := services_security.EncryptData(fldValue, &eData)
			if retval < 0 {
				rejectDesc = fmt.Sprintf("EncryptData() failed in TransformObjectFunctionType for FunctionName[%s]", execFunction.FunctionName)
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_TransformObjectFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.TransformObjectInfo.AppErrDesc, execFunction.TransformObjectInfo.ObjectData[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.TransformObjectInfo.AppCorrectiveAction, execFunction.TransformObjectInfo.ObjectData[i].AppCorrectiveAction), rejectDesc)
				return -1
			}
			fldValue = eData
		} else if execFunction.TransformObjectInfo.ObjectData[i].Algo == transformobjdef.TransformObjectAlgoDecrypt {
			clearData := ""
			retval, _ := services_security.DecryptData(fldValue, &clearData)
			if retval < 0 {
				rejectDesc = fmt.Sprintf("DecryptData() failed in TransformObjectFunctionType for FunctionName[%s]", execFunction.FunctionName)
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_TransformObjectFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.TransformObjectInfo.AppErrDesc, execFunction.TransformObjectInfo.ObjectData[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.TransformObjectInfo.AppCorrectiveAction, execFunction.TransformObjectInfo.ObjectData[i].AppCorrectiveAction), rejectDesc)
				return -1
			}
			fldValue = clearData
		} else {
			rejectDesc := fmt.Sprintf("Algo[%s] not supported for ObjectData[%s]", execFunction.TransformObjectInfo.ObjectData[i].Algo, execFunction.TransformObjectInfo.ObjectData[i].Key)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_TransformObjectFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.TransformObjectInfo.AppErrDesc, execFunction.TransformObjectInfo.ObjectData[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.TransformObjectInfo.AppCorrectiveAction, execFunction.TransformObjectInfo.ObjectData[i].AppCorrectiveAction), rejectDesc)
			return -1
		}
		objectMap[execFunction.TransformObjectInfo.ObjectData[i].Key] = fldValue
	}
	_, ok = reqBrokerDataMap[execFunction.TransformObjectInfo.ObjectName]
	if !ok {
		reqBrokerDataMap[execFunction.TransformObjectInfo.ObjectName] = objectMap
	} else {
		subMap, ok := reqBrokerDataMap[execFunction.TransformObjectInfo.ObjectName].(map[string]interface{})
		if !ok {
			rejectDesc := fmt.Sprintf("ObjectName[%s] Type Assertion Failed for FunctionName[%s]", execFunction.TransformObjectInfo.ObjectName, execFunction.FunctionName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_TransformObjectFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.TransformObjectInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.TransformObjectInfo.AppCorrectiveAction), rejectDesc)
			return -1
		}
		for key, value := range objectMap {
			subMap[key] = value
		}
	}
	return 1
}
