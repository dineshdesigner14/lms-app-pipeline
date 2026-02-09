package lexicalparser

import (
	"fmt"
	"lmsapieng/include/common/lexicalparserdef"
	"lmsapieng/include/common/rejectdef"
	"lmsapieng/include/common/reqbrokerdef"
	"lmsapieng/libsrc/services/services_security"
	"lmsapieng/libsrc/utils/dtutil"
	"lmsapieng/libsrc/utils/lexicalparserutil"
	"lmsapieng/libsrc/utils/mathutil"
	"lmsapieng/libsrc/utils/msgutil"
	"lmsapieng/libsrc/utils/templateutil"
)

func GenCVVFunctionType(execFunction lexicalparserdef.ExecFunctionStruct, reqBrokerDataMap map[string]interface{}) int {
	var rejectDesc string

	if len(execFunction.GenCVVInfo.ObjectName) == 0 {
		rejectDesc = fmt.Sprintf("ObjectName is NULL in GenCVVInfoFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GenCVVFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenCVVInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenCVVInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	if len(execFunction.GenCVVInfo.CardNum) == 0 {
		rejectDesc = fmt.Sprintf("CardNum is NULL in GenCVVInfoFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GenCVVFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenCVVInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenCVVInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	dataValue := lexicalparserutil.ReadValueFrmDataMap(reqBrokerDataMap, execFunction.GenCVVInfo.CardNum, "string", &rejectDesc)
	if dataValue == nil {
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ReadValueErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenCVVInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenCVVInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	lCardNum := dataValue.(string)
	if !mathutil.IsValidNumericStr(lCardNum, 16, 19) {
		rejectDesc = fmt.Sprintf("CardNum[%s] is Invalid in GenCVVInfoFunctionType for FunctionName[%s]", lCardNum, execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GenCVVFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenCVVInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenCVVInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	if len(execFunction.GenCVVInfo.ExpDate) == 0 {
		rejectDesc = fmt.Sprintf("ExpDate is NULL in GenCVVInfoFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GenCVVFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenCVVInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenCVVInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	dataValue = lexicalparserutil.ReadValueFrmDataMap(reqBrokerDataMap, execFunction.GenCVVInfo.ExpDate, "string", &rejectDesc)
	if dataValue == nil {
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ReadValueErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenCVVInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenCVVInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	lExpDate := dataValue.(string)
	if len(lExpDate) != 8 {
		rejectDesc = fmt.Sprintf("ExpDate[%s] should be 8 digit in GenCVVInfoFunctionType for FunctionName[%s]", lExpDate, execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GenCVVFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenCVVInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenCVVInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	if !dtutil.IsValidDateFormat(lExpDate, "DDMMYYYY") {
		rejectDesc = fmt.Sprintf("ExpDate[%s] is invalid in GenCVVInfoFunctionType for FunctionName[%s]", lExpDate, execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GenCVVFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenCVVInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenCVVInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	if len(execFunction.GenCVVInfo.ServiceCode) == 0 {
		rejectDesc = fmt.Sprintf("ServiceCode is NULL in GenCVVInfoFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GenCVVFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenCVVInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenCVVInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	dataValue = lexicalparserutil.ReadValueFrmDataMap(reqBrokerDataMap, execFunction.GenCVVInfo.ServiceCode, "string", &rejectDesc)
	if dataValue == nil {
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ReadValueErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenCVVInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenCVVInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	lServiceCode := dataValue.(string)
	if !mathutil.IsValidNumericStr(lServiceCode, 3, 3) {
		rejectDesc = fmt.Sprintf("ServiceCode[%s] is Invalid in GenCVVInfoFunctionType for FunctionName[%s]", lServiceCode, execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GenCVVFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenCVVInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenCVVInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	if len(execFunction.GenCVVInfo.CVK1) == 0 {
		rejectDesc = fmt.Sprintf("CVK1 is NULL in GenCVVInfoFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GenCVVFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenCVVInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenCVVInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	dataValue = lexicalparserutil.ReadValueFrmDataMap(reqBrokerDataMap, execFunction.GenCVVInfo.CVK1, "string", &rejectDesc)
	if dataValue == nil {
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ReadValueErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenCVVInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenCVVInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	lCVK1 := dataValue.(string)
	if len(execFunction.GenCVVInfo.CVK2) == 0 {
		rejectDesc = fmt.Sprintf("CVK2 is NULL in GenCVVInfoFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GenCVVFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenCVVInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenCVVInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	dataValue = lexicalparserutil.ReadValueFrmDataMap(reqBrokerDataMap, execFunction.GenCVVInfo.CVK2, "string", &rejectDesc)
	if dataValue == nil {
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ReadValueErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenCVVInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenCVVInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	lCVK2 := dataValue.(string)

	objectMap := make(map[string]interface{})

	lexpDate := lExpDate[6:] + lExpDate[2:4]
	cvv := ""
	retval, _ := services_security.GenCVV(lCardNum, lexpDate, lServiceCode, lCVK1, lCVK2, &cvv)
	if retval < 0 {
		rejectDesc = fmt.Sprintf("GenCVV1() failed in GenCVVInfoFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GenCVVFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenCVVInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenCVVInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	objectMap["cvv1"] = cvv

	lexpDate = lExpDate[2:4] + lExpDate[6:]
	cvv = ""
	retval, _ = services_security.GenCVV(lCardNum, lexpDate, "000", lCVK1, lCVK2, &cvv)
	if retval < 0 {
		rejectDesc = fmt.Sprintf("GenCVV2() failed in GenCVVInfoFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GenCVVFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenCVVInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenCVVInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	objectMap["cvv2"] = cvv

	lexpDate = lExpDate[6:] + lExpDate[2:4]
	cvv = ""
	retval, _ = services_security.GenCVV(lCardNum, lexpDate, "999", lCVK1, lCVK2, &cvv)
	if retval < 0 {
		rejectDesc = fmt.Sprintf("GenICVV() failed in GenCVVInfoFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GenCVVFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenCVVInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenCVVInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	objectMap["icvv"] = cvv
	reqBrokerDataMap[execFunction.GenCVVInfo.ObjectName] = objectMap
	return 1
}
