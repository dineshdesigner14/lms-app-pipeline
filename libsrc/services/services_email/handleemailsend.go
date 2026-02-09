package services_email

import (
	"encoding/json"
	"fmt"
	"lmsapieng/include/common/msgdef"
	"lmsapieng/include/common/rejectdef"
	"lmsapieng/include/common/reqbrokerdef"
	"lmsapieng/libsrc/utils/datatypeutil"
	"lmsapieng/libsrc/utils/maputil"
	"lmsapieng/libsrc/utils/msgutil"
	"lmsapieng/libsrc/utils/serviceutil"
	"lmsapieng/libsrc/utils/templateutil"
	"lmsapieng/libsrc/utils/trace"
)

func HandleSendEmail(reqBrokerDataMap map[string]interface{}) int {
	var respInfo msgdef.RespInfoStruct
	var respBuffer []byte
	_, ok := reqBrokerDataMap["email_obj"]
	if !ok {
		appErrDesc := fmt.Sprintf("ObjectName[%s] does not exist in reqBrokerDataMap", "email_obj")
		correctiveAction := fmt.Sprintf("Ensure that ObjectName[%s] exist in reqBrokerDataMap", "email_obj")
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Reject_HandleSendEmailErr, []byte(appErrDesc), []byte(correctiveAction))
		return -1
	}
	emailObjMap := reqBrokerDataMap["email_obj"].(map[string]interface{})
	if !datatypeutil.IsString(emailObjMap["to_address_list"]) {
		appErrDesc := fmt.Sprintf("to_address_list[%T] is not object array in reqBrokerDataMap", emailObjMap["to_address_list"])
		correctiveAction := fmt.Sprintf("Ensure that to_address_list[%T] is not object array in reqBrokerDataMap", emailObjMap["to_address_list"])
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Reject_HandleSendEmailErr, []byte(appErrDesc), []byte(correctiveAction))
		return -1
	}
	_, ok = reqBrokerDataMap["cc_address_list"]
	if ok {
		if !datatypeutil.IsObjectArray(emailObjMap["cc_address_list"]) {
			appErrDesc := fmt.Sprintf("cc_address_list[%T] is not object array in reqBrokerDataMap", emailObjMap["cc_address_list"])
			correctiveAction := fmt.Sprintf("Ensure that cc_address_list[%T] is not object array in reqBrokerDataMap", emailObjMap["cc_address_list"])
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Reject_HandleSendEmailErr, []byte(appErrDesc), []byte(correctiveAction))
			return -1
		}
	}
	_, ok = reqBrokerDataMap["to_address_list"]
	if ok {
		if !datatypeutil.IsObjectArray(emailObjMap["to_address_list"]) {
			appErrDesc := fmt.Sprintf("to_address_list[%T] is not object array in reqBrokerDataMap", emailObjMap["to_address_list"])
			correctiveAction := fmt.Sprintf("Ensure that to_address_list[%T] is not object array in reqBrokerDataMap", emailObjMap["to_address_list"])
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Reject_HandleSendEmailErr, []byte(appErrDesc), []byte(correctiveAction))
			return -1
		}
	}
	if !datatypeutil.IsString(emailObjMap["subject"]) {
		appErrDesc := fmt.Sprintf("subject[%T] is not string in reqBrokerDataMap", emailObjMap["subject"])
		correctiveAction := fmt.Sprintf("Ensure that subject[%T] is not string in reqBrokerDataMap", emailObjMap["subject"])
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Reject_HandleSendEmailErr, []byte(appErrDesc), []byte(correctiveAction))
		return -1
	}
	if !datatypeutil.IsString(emailObjMap["body"]) {
		appErrDesc := fmt.Sprintf("body[%T] is not string in reqBrokerDataMap", emailObjMap["body"])
		correctiveAction := fmt.Sprintf("Ensure that body[%T] is not string in reqBrokerDataMap", emailObjMap["body"])
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Reject_HandleSendEmailErr, []byte(appErrDesc), []byte(correctiveAction))
		return -1
	}
	if !datatypeutil.IsString(emailObjMap["msg_id"]) {
		appErrDesc := fmt.Sprintf("msg_id[%T] is not string in reqBrokerDataMap", emailObjMap["msg_id"])
		correctiveAction := fmt.Sprintf("Ensure that msg_id[%T] is not string in reqBrokerDataMap", emailObjMap["msg_id"])
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Reject_HandleSendEmailErr, []byte(appErrDesc), []byte(correctiveAction))
		return -1
	}
	if !datatypeutil.IsString(emailObjMap["request_num"]) {
		appErrDesc := fmt.Sprintf("request_num[%T] is not string in reqBrokerDataMap", emailObjMap["request_num"])
		correctiveAction := fmt.Sprintf("Ensure that request_num[%T] is not string in reqBrokerDataMap", emailObjMap["request_num"])
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Reject_HandleSendEmailErr, []byte(appErrDesc), []byte(correctiveAction))
		return -1
	}
	if !datatypeutil.IsString(emailObjMap["request_type"]) {
		appErrDesc := fmt.Sprintf("request_type[%T] is not string in reqBrokerDataMap", emailObjMap["request_type"])
		correctiveAction := fmt.Sprintf("Ensure that request_type[%T] is not string in reqBrokerDataMap", emailObjMap["request_type"])
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Reject_HandleSendEmailErr, []byte(appErrDesc), []byte(correctiveAction))
		return -1
	}
	titleTemplate := ""
	if maputil.CopyStringFromMap(emailObjMap, "body", &titleTemplate) < 0 {
		return 1
	}
	emailObjMap["body"] = templateutil.GetTemplateString(reqBrokerDataMap, titleTemplate)
	emailMap := make(map[string]interface{})
	emailMap["mail_data"] = emailObjMap
	emailMapBuffer, _ := json.Marshal(emailMap)
	if serviceutil.SendToService(reqBrokerDataMap, "EmailEng", emailMapBuffer, &respBuffer, 30) < 0 {
		appErrDesc := fmt.Sprintf("SendToService failed for EmailEng")
		correctiveAction := fmt.Sprintf("Ensure that SendToService does not failed for EmailEng")
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Reject_HandleSendEmailErr, []byte(appErrDesc), []byte(correctiveAction))
		return -1
	}
	if msgutil.ParseResp(respBuffer, &respInfo) < 0 {
		appErrDesc := fmt.Sprintf("SendToService failed for EmailEng with Err[%s]", respBuffer)
		correctiveAction := "Ensure that SendToService does not failed for EmailEng"
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Reject_HandleSendEmailErr, []byte(appErrDesc), []byte(correctiveAction))
		return -1
	}
	trace.Lg("respInfo = %+v", respInfo)

	return 1
}
