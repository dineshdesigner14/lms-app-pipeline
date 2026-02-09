package services_email

import (
	"encoding/json"
	"lmsapieng/include/common/emaildef"
	"lmsapieng/include/common/httpdef"
	"lmsapieng/include/common/moduledef"
	"lmsapieng/include/common/msgdef"
	"lmsapieng/include/common/rejectdef"
	"lmsapieng/libsrc/utils/msgutil"
)

func SendMail(emailMsgInfo emaildef.EMailEngMsgInfo) (int, []byte) {
	var emailEngAddr httpdef.HttpServerAddr
	var respData []byte
	var respInfo msgdef.RespInfoStruct
	var rejectDesc, correctiveAction string

	reqData, _ := json.Marshal(&emailMsgInfo)

	if locateEmailEngAddr(&emailEngAddr) < 0 {
		//trace.Lg("locateEmailEngAddr() failed")
		rejectDesc = "emailEngAddr() Error"
		correctiveAction = "Check the emailEngAddr"
		return -1, msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_LocateEmailEngAddrFailed, []byte(rejectDesc), []byte(correctiveAction), rejectDesc)
	}
	rval := msgutil.PostReq(moduledef.EmailEngModule, emailEngAddr.ServerIpAddr, emailEngAddr.ServerPort, emailEngAddr.ServerTimeout, reqData, &respData)
	if rval < 0 {
		if rval == -2 {
			rejectDesc = "EmailEng TimedOut Error"
			correctiveAction = "Check Why EmailEng TimedOut Error Occured"
			return -1, msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_SendToEmailEngTimedOut, []byte(rejectDesc), []byte(correctiveAction), rejectDesc)
		}
		rejectDesc = "EmailEng Send Error"
		correctiveAction = "Check Why EmailEng Send Error Occured"
		return -1, msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_SendToEmailEngFailed, []byte(rejectDesc), []byte(correctiveAction), rejectDesc)
	}
	if msgutil.ParseResp(respData, &respInfo) < 0 {
		return -1, respData
	}
	return 1, respData
}
