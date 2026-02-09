package lexicalparser

import (
	"fmt"
	"lmsapieng/include/common/datatypedef"
	"lmsapieng/include/common/lexicalparserdef"
	"lmsapieng/include/common/rejectdef"
	"lmsapieng/include/common/reqbrokerdef"
	"lmsapieng/libsrc/utils/amtutil"
	"lmsapieng/libsrc/utils/datatypeutil"
	"lmsapieng/libsrc/utils/lexicalparserutil"
	"lmsapieng/libsrc/utils/msgutil"
	"lmsapieng/libsrc/utils/templateutil"
)

func MathFunctionType(execFunction lexicalparserdef.ExecFunctionStruct, reqBrokerDataMap map[string]interface{}) int {
	var rejectDesc string
	if len(execFunction.MathFunctionInfo.Algo) == 0 {
		rejectDesc = fmt.Sprintf("Algo is NULL in MathFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_MathFunctionFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.MathFunctionInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.MathFunctionInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	if execFunction.MathFunctionInfo.Algo == "calculate_sum" {
		if len(execFunction.MathFunctionInfo.SrcObject) == 0 {
			rejectDesc = fmt.Sprintf("SrcObject is NULL in MathFunctionType for FunctionName[%s]", execFunction.FunctionName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_MathFunctionFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.MathFunctionInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.MathFunctionInfo.AppCorrectiveAction), rejectDesc)
			return -1
		}
		dataValue := lexicalparserutil.ReadValueFrmDataMap(reqBrokerDataMap, execFunction.MathFunctionInfo.SrcObject, datatypedef.DataTypeObjectArray, &rejectDesc)
		if dataValue == nil {
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ReadValueErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.MathFunctionInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.MathFunctionInfo.AppCorrectiveAction), rejectDesc)
			return -1
		}
		if !datatypeutil.IsObjectArray(dataValue) {
			rejectDesc = fmt.Sprintf("SrcObject[%s] is Not ObjectArray in MathFunctionType for FunctionName[%s]", execFunction.MathFunctionInfo.SrcObject, execFunction.FunctionName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_MathFunctionFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.MathFunctionInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.MathFunctionInfo.AppCorrectiveAction), rejectDesc)
			return -1
		}
		if len(execFunction.MathFunctionInfo.SrcKey) == 0 {
			rejectDesc = fmt.Sprintf("SrcKey is NULL in MathFunctionType for FunctionName[%s]", execFunction.FunctionName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_MathFunctionFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.MathFunctionInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.MathFunctionInfo.AppCorrectiveAction), rejectDesc)
			return -1
		}
		if len(execFunction.MathFunctionInfo.DestObject) == 0 {
			rejectDesc = fmt.Sprintf("DestObject is NULL in MathFunctionType for FunctionName[%s]", execFunction.FunctionName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_MathFunctionFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.MathFunctionInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.MathFunctionInfo.AppCorrectiveAction), rejectDesc)
			return -1
		}
		objectArray := dataValue.([]interface{})
		NetAmount := "0.00"
		amountStr := "0.00"
		for i := 0; i < len(objectArray); i++ {
			found := false
			for key, value := range objectArray[i].(map[string]interface{}) {
				if key == execFunction.MathFunctionInfo.SrcKey {
					amountStr = value.(string)
					if !amtutil.IsValidAmount(amountStr, 1, 10, 2, 2) {
						rejectDesc = fmt.Sprintf("Amount[%s] is Not Valid in MathFunctionType for FunctionName[%s]", amountStr, execFunction.FunctionName)
						reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_MathFunctionFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.MathFunctionInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.MathFunctionInfo.AppCorrectiveAction), rejectDesc)
						return -1
					}
					found = true
					break
				}
			}
			if !found {
				rejectDesc = fmt.Sprintf("SrcKey[%s] Not Present in MathFunctionType for FunctionName[%s]", execFunction.MathFunctionInfo.SrcKey, execFunction.FunctionName)
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_MathFunctionFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.MathFunctionInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.MathFunctionInfo.AppCorrectiveAction), rejectDesc)
				return -1
			}
			NetAmount = amtutil.AddAmt(NetAmount, amountStr, 2)
		}
		//trace.Lg("NetAmount[%s]", NetAmount)
		reqBrokerDataMap[execFunction.MathFunctionInfo.DestObject] = NetAmount
	} else {
		rejectDesc = fmt.Sprintf("Algo[%s] is Invalid in MathFunctionType for FunctionName[%s]", execFunction.MathFunctionInfo.Algo, execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_MathFunctionFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.MathFunctionInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.MathFunctionInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	return 1
}
