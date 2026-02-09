package services_otp

import (
	"encoding/json"
	"fmt"
	"lmsapieng/include/common/lmsapiengdef"
	"lmsapieng/include/common/msgdef"
	"lmsapieng/include/common/rejectdef"
	"lmsapieng/include/common/reqbrokerdef"
	"lmsapieng/libsrc/utils/datatypeutil"
	"lmsapieng/libsrc/utils/msgutil"
	"lmsapieng/libsrc/utils/reqbrokerutil"
	"lmsapieng/libsrc/utils/serviceutil"
	"lmsapieng/libsrc/utils/trace"
)

func HandleOTPGen(reqBrokerDataMap map[string]interface{}) int {
	var respBuffer []byte
	var respInfo msgdef.RespInfoStruct
	_, ok := reqBrokerDataMap["otp_obj"]
	if !ok {
		appErrDesc := fmt.Sprintf("ObjectName[%s] does not exist in reqBrokerDataMap", "otp_obj")
		correctiveAction := fmt.Sprintf("Ensure that ObjectName[%s] exist in reqBrokerDataMap", "otp_obj")
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_HandleSendOTPErr, []byte(appErrDesc), []byte(correctiveAction))
		return -1
	}
	otpObjMap := reqBrokerDataMap["otp_obj"].(map[string]interface{})
	_, ok = otpObjMap["otp_ref_data"]
	if !ok {
		appErrDesc := fmt.Sprintf("otp_ref_data does not exist in otp_obj", "otp_obj")
		correctiveAction := fmt.Sprintf("Ensure that otp_ref_data does exist in otp_obj")
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_HandleSendOTPErr, []byte(appErrDesc), []byte(correctiveAction))
		return -1
	}
	if !datatypeutil.IsString(otpObjMap["otp_ref_data"]) {
		appErrDesc := fmt.Sprintf("otp_ref_data[%T] is not string in reqBrokerDataMap", otpObjMap["otp_ref_data"])
		correctiveAction := fmt.Sprintf("Ensure that otp_ref_data[%T] is a string in reqBrokerDataMap", otpObjMap["otp_ref_data"])
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_HandleSendOTPErr, []byte(appErrDesc), []byte(correctiveAction))
		return -1
	}
	otp_ref_data := otpObjMap["otp_ref_data"].(string)
	if len(otp_ref_data) == 0 {
		appErrDesc := fmt.Sprintf("otp_ref_data should not be null in reqBrokerDataMap")
		correctiveAction := fmt.Sprintf("Ensure that otp_ref_data is not a null in reqBrokerDataMap")
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_HandleSendOTPErr, []byte(appErrDesc), []byte(correctiveAction))
		return -1
	}
	otpObjMapBuffer, _ := json.Marshal(otpObjMap)
	trace.Lg("otpObjMapBuffer[%s]", otpObjMapBuffer)
	reqBuffer := reqbrokerutil.TranslateInternalMsg("GenOTP", serviceutil.GetServiceName(), lmsapiengdef.LMSApiEngVersion, json.RawMessage(otpObjMapBuffer))
	trace.Lg("reqBuffer[%s]", reqBuffer)
	if serviceutil.SendToService(reqBrokerDataMap, "OTPEng", reqBuffer, &respBuffer, 30) < 0 {
		appErrDesc := fmt.Sprintf("SendToService failed for OTPEng")
		correctiveAction := fmt.Sprintf("Ensure that SendToService does not failed for OTPEng")
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_HandleSendOTPErr, []byte(appErrDesc), []byte(correctiveAction))
		return -1
	}
	trace.Lg("sendToServie completed")
	if msgutil.ParseResp(respBuffer, &respInfo) < 0 {
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_HandleSendOTPErr, []byte(respInfo.RespInfo.AppErrDesc), []byte(respInfo.RespInfo.AppCorrectiveAction))
		return -1
	}
	respDataMap := make(map[string]interface{})
	json.Unmarshal(respInfo.RespInfo.RespData, &respDataMap)
	trace.Lg("respDataMap[%s]", respDataMap)
	reqBrokerDataMap["otp_resp_obj"] = respDataMap
	return 1
}
