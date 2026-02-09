package lexicalparser

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"lmsapieng/include/common/datatypedef"
	"lmsapieng/include/common/globaldef"
	"lmsapieng/include/common/keytypedef"
	"lmsapieng/include/common/lexicalparserdef"
	"lmsapieng/include/common/rejectdef"
	"lmsapieng/include/common/reqbrokerdef"
	"lmsapieng/libsrc/services/services_security"
	"lmsapieng/libsrc/utils/cryptoutil"
	"lmsapieng/libsrc/utils/lexicalparserutil"
	"lmsapieng/libsrc/utils/msgutil"
	"lmsapieng/libsrc/utils/templateutil"
	"strings"
)

func isValidKeyType(KeyType string) bool {
	if KeyType != keytypedef.KeyTypeDMK && KeyType != keytypedef.KeyTypeDEK {
		return false
	}
	return true
}

func generateRandomHexKey() (string, error) {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return strings.ToUpper(hex.EncodeToString(bytes)), nil
}

func GenRandomKeyFunctionType(execFunction lexicalparserdef.ExecFunctionStruct, reqBrokerDataMap map[string]interface{}) int {
	var rejectDesc string
	var err error
	if len(execFunction.GenRandomKeyInfo.ObjectName) == 0 {
		rejectDesc = fmt.Sprintf("ObjectName is NULL in GenRandomKeyFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GenRandomKeyFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenRandomKeyInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenRandomKeyInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	if len(execFunction.GenRandomKeyInfo.KeyType) == 0 {
		rejectDesc = fmt.Sprintf("KeyType is NULL in GenRandomKeyFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GenRandomKeyFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenRandomKeyInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenRandomKeyInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	if !isValidKeyType(execFunction.GenRandomKeyInfo.KeyType) {
		rejectDesc = fmt.Sprintf("KeyType[%s] is Invalid in GenRandomKeyFunctionType for FunctionName[%s]", execFunction.GenRandomKeyInfo.KeyType, execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GenRandomKeyFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenRandomKeyInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenRandomKeyInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	lKey1 := globaldef.NOT_INITIALIZED
	objectMap := make(map[string]interface{})
	objectMap["key_type"] = execFunction.GenRandomKeyInfo.KeyType
	if execFunction.GenRandomKeyInfo.KeyType == keytypedef.KeyTypeDEK {
		if len(execFunction.GenRandomKeyInfo.Key1) == 0 {
			rejectDesc = fmt.Sprintf("Key1 is NULL in GenRandomKeyFunctionType for FunctionName[%s]", execFunction.FunctionName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GenRandomKeyFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenRandomKeyInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenRandomKeyInfo.AppCorrectiveAction), rejectDesc)
			return -1
		}
		dataValue := lexicalparserutil.ReadValueFrmDataMap(reqBrokerDataMap, execFunction.GenRandomKeyInfo.Key1, "string", &rejectDesc)
		if dataValue == nil {
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ReadValueErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenRandomKeyInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenRandomKeyInfo.AppCorrectiveAction), rejectDesc)
			return -1
		}
		lKey1 = dataValue.(string)
	}
	objectMap["key_1"] = lKey1
	if execFunction.GenRandomKeyInfo.KeyType == keytypedef.KeyTypeDMK {
		comp1 := ""
		comp2 := ""
		if execFunction.GenRandomKeyInfo.Comp1 == globaldef.NOT_INITIALIZED {
			comp1, err = generateRandomHexKey()
			if err != nil {
				rejectDesc = fmt.Sprintf("generateRandomHexKey() failed in GenRandomKeyFunctionType for FunctionName[%s]", execFunction.FunctionName)
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GenRandomKeyFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenRandomKeyInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenRandomKeyInfo.AppCorrectiveAction), rejectDesc)
				return -1
			}
			comp2 = "00000000000000000000000000000000"
		} else {
			dataValue := lexicalparserutil.ReadValueFrmDataMap(reqBrokerDataMap, execFunction.GenRandomKeyInfo.Comp1, datatypedef.DataTypeString, &rejectDesc)
			if dataValue == nil {
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ReadValueErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenRandomKeyInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenRandomKeyInfo.AppCorrectiveAction), rejectDesc)
				return -1
			}
			comp1 = dataValue.(string)
			clearComp := ""
			retval, respData := services_security.DecryptData(comp1, &clearComp)
			if retval < 0 {
				rejectDesc = fmt.Sprintf("DecryptData() with respData[%s] failed in GenRandomKeyFunctionType for FunctionName[%s]", respData, execFunction.FunctionName)
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GenRandomKeyFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenRandomKeyInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenRandomKeyInfo.AppCorrectiveAction), rejectDesc)
				return -1
			}
			comp1 = clearComp
			dataValue = lexicalparserutil.ReadValueFrmDataMap(reqBrokerDataMap, execFunction.GenRandomKeyInfo.Comp2, datatypedef.DataTypeString, &rejectDesc)
			if dataValue == nil {
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ReadValueErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenRandomKeyInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenRandomKeyInfo.AppCorrectiveAction), rejectDesc)
				return -1
			}
			comp2 = dataValue.(string)
			clearComp = ""
			retval, respData = services_security.DecryptData(comp2, &clearComp)
			if retval < 0 {
				rejectDesc = fmt.Sprintf("DecryptData() with respData[%s] failed in GenRandomKeyFunctionType for FunctionName[%s]", respData, execFunction.FunctionName)
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GenRandomKeyFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenRandomKeyInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenRandomKeyInfo.AppCorrectiveAction), rejectDesc)
				return -1
			}
			comp2 = clearComp
		}
		eTMK := ""
		TMKKCV := ""
		Component1KCV := ""
		Component2KCV := ""
		//trace.Lg("comp1[%s]", comp1)
		//trace.Lg("comp2[%s]", comp2)
		retval, respData := services_security.GenTMK(comp1, comp2, &eTMK, &TMKKCV, &Component1KCV, &Component2KCV)
		if retval < 0 {
			rejectDesc = fmt.Sprintf("GenTMK() with respData[%s] failed in GenRandomKeyFunctionType for FunctionName[%s]", respData, execFunction.FunctionName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GenRandomKeyFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenRandomKeyInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenRandomKeyInfo.AppCorrectiveAction), rejectDesc)
			return -1
		}
		retVal, finalComp := cryptoutil.XorStrings(comp1, comp2)
		if retVal < 0 {
			rejectDesc = fmt.Sprintf("XorStrings failed in GenRandomKeyFunctionType for FunctionName[%s]", execFunction.FunctionName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GenRandomKeyFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenRandomKeyInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenRandomKeyInfo.AppCorrectiveAction), rejectDesc)
			return -1
		}
		objectMap["e_random_key"] = eTMK
		objectMap["random_key"] = finalComp
		objectMap["random_key_kcv"] = TMKKCV
		reqBrokerDataMap[execFunction.GenRandomKeyInfo.ObjectName] = objectMap
	} else {
		TPKUTMK := ""
		TPKULMK := ""
		TPKKCV := ""
		//trace.Lg("lKey1[%s]", lKey1)
		retval, respData := services_security.GenTermCommKey(lKey1, &TPKUTMK, &TPKULMK, &TPKKCV)
		if retval < 0 {
			rejectDesc = fmt.Sprintf("GenTermCommKey() with respData[%s] failed in GenRandomKeyFunctionType for FunctionName[%s]", respData, execFunction.FunctionName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GenRandomKeyFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenRandomKeyInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenRandomKeyInfo.AppCorrectiveAction), rejectDesc)
			return -1
		}
		objectMap["e_random_key"] = TPKULMK
		objectMap["random_key"] = TPKUTMK
		objectMap["random_key_kcv"] = TPKKCV
		reqBrokerDataMap[execFunction.GenRandomKeyInfo.ObjectName] = objectMap
	}
	return 1
}
