package lexicalparser

import (
	"fmt"
	"lmsapieng/include/common/debugdef"
	"lmsapieng/include/common/emaildef"
	"lmsapieng/include/common/lexicalparserdef"
	"lmsapieng/include/common/rejectdef"
	"lmsapieng/include/common/reqbrokerdef"
	"lmsapieng/libsrc/services/services_email"
	"lmsapieng/libsrc/utils/msgutil"
	"lmsapieng/libsrc/utils/templateutil"
	"lmsapieng/libsrc/utils/trace"
)

func SendEmailFunctionType(execFunction lexicalparserdef.ExecFunctionStruct, reqBrokerDataMap map[string]interface{}) int {
	var emailMsgInfo emaildef.EMailEngMsgInfo
	var rejectDesc string
	if len(execFunction.SendEmailInfo.GatewayName) == 0 {
		rejectDesc = fmt.Sprintf("GatewayName is NULL in SendEmailFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_SendEmailFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.SendEmailInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.SendEmailInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	if len(execFunction.SendEmailInfo.FromAddress) == 0 {
		rejectDesc = fmt.Sprintf("FromAddress is NULL in SendEmailFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_SendEmailFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.SendEmailInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.SendEmailInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	if len(execFunction.SendEmailInfo.ToAddress) == 0 {
		rejectDesc = fmt.Sprintf("ToAddress is NULL in SendEmailFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_SendEmailFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.SendEmailInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.SendEmailInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	if len(execFunction.SendEmailInfo.Subject) == 0 {
		rejectDesc = fmt.Sprintf("Subject is NULL in SendEmailFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_SendEmailFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.SendEmailInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.SendEmailInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	if len(execFunction.SendEmailInfo.Body) == 0 {
		rejectDesc = fmt.Sprintf("Body is NULL in SendEmailFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_SendEmailFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.SendEmailInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.SendEmailInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	emailMsgInfo.EmailData.GatewayName = templateutil.GetTemplateString(reqBrokerDataMap, execFunction.SendEmailInfo.GatewayName)
	emailMsgInfo.EmailData.FromAddress = templateutil.GetTemplateString(reqBrokerDataMap, execFunction.SendEmailInfo.FromAddress)
	for i := 0; i < len(execFunction.SendEmailInfo.ToAddress); i++ {
		emailMsgInfo.EmailData.ToAddressList = append(emailMsgInfo.EmailData.ToAddressList, templateutil.GetTemplateString(reqBrokerDataMap, execFunction.SendEmailInfo.ToAddress[i]))
	}
	for i := 0; i < len(execFunction.SendEmailInfo.CCAddress); i++ {
		emailMsgInfo.EmailData.CCAddressList = append(emailMsgInfo.EmailData.CCAddressList, templateutil.GetTemplateString(reqBrokerDataMap, execFunction.SendEmailInfo.CCAddress[i]))
	}
	emailMsgInfo.EmailData.Subject = templateutil.GetTemplateString(reqBrokerDataMap, execFunction.SendEmailInfo.Subject)
	emailMsgInfo.EmailData.Body = templateutil.GetTemplateString(reqBrokerDataMap, execFunction.SendEmailInfo.Body)
	emailMsgInfo.EmailData.RequestType = templateutil.GetTemplateString(reqBrokerDataMap, "{{ .ReqBrokerReqObj.request_key.req_type }}")
	emailMsgInfo.EmailData.RequestNum = emailMsgInfo.EmailData.RequestType
	emailMsgInfo.EmailData.MsgID = emailMsgInfo.EmailData.RequestType
	retVal, respData := services_email.SendMail(emailMsgInfo)
	trace.Log(debugdef.DEBUG_LEVEL_TEST, "SendMail() got respData[%s] in SendEmailFunctionType", respData)
	if retVal < 0 {
		rejectDesc = fmt.Sprintf("SendMail() failed for SendEmailFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_SendEmailFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.SendEmailInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.SendEmailInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	return 1
}
