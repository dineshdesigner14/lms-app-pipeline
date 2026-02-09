package services_security

import (
	"encoding/json"
	"lmsapieng/include/common/globaldef"
	"lmsapieng/include/common/httpdef"
	"lmsapieng/include/common/moduledef"
	"lmsapieng/include/common/msgdef"
	"lmsapieng/include/common/rejectdef"
	"lmsapieng/include/common/securitydef"
	"lmsapieng/libsrc/utils/msgutil"
)

func DecryptData(EData string, Data *string, secArgs ...string) (int, []byte) {
	var securityEngAddr httpdef.HttpServerAddr
	var respData []byte
	var respInfo msgdef.RespInfoStruct
	var rejectDesc, correctiveAction, securityID string
	var decryptDataReqMsg securitydef.DecryptDataReqMsgStruct
	var decryptDataRespMsg securitydef.DecryptDataRespMsgStruct

	if len(EData) == 0 {
		rejectDesc = "DecryptData is NULL"
		correctiveAction = "Check the DecryptData Used for Decryption for NULL"
		return -1, msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_SecurityEngInvalidReq, []byte(rejectDesc), []byte(correctiveAction), rejectDesc)
	}
	if len(secArgs) == 0 {
		securityID = globaldef.NOT_INITIALIZED
	} else {
		securityID = secArgs[0]
	}
	decryptDataReqMsg.EData = EData
	reqMap := make(map[string]interface{})
	reqMap[securitydef.SecurityEngCommandJSONObj] = securitydef.DecryptDataCommand
	reqMap[securitydef.SecurityEngIDJSONObj] = securityID
	reqMap[securitydef.SecurityEngDataJSONObj] = decryptDataReqMsg
	reqData, _ := json.Marshal(&reqMap)

	if locateSecurityEngAddr(&securityEngAddr) < 0 {
		//trace.Lg("locateSecurityEngAddr() failed")
		rejectDesc = "locateSecurityEngAddr Error"
		correctiveAction = "Check the SecurityEngAddr"
		return -1, msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_LocateSecurityEngAddrFailed, []byte(rejectDesc), []byte(correctiveAction))
	}
	rval := msgutil.PostReq(moduledef.SecurityEngModule, securityEngAddr.ServerIpAddr, securityEngAddr.ServerPort, securityEngAddr.ServerTimeout, reqData, &respData)
	if rval < 0 {
		if rval == -2 {
			rejectDesc = "SecurityEng TimedOut Error"
			correctiveAction = "Check Why SecurityEng TimedOut Error Occured"
			return -1, msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_SendToSecurityEngTimedOut, []byte(rejectDesc), []byte(correctiveAction))
		}
		rejectDesc = "SecurityEng Send Error"
		correctiveAction = "Check Why SecurityEng Send Error Occured"
		return -1, msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_SendToSecurityEngFailed, []byte(rejectDesc), []byte(correctiveAction))
	}
	if msgutil.ParseResp(respData, &respInfo) < 0 {
		return -1, respData
	}
	err := json.Unmarshal(respInfo.RespInfo.RespData, &decryptDataRespMsg)
	if err != nil {
		rejectDesc = "SecurityEng Response Unmarshal Error"
		correctiveAction = "Check SecurityEng Response Unmarshal Error Occured"
		return -1, msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_SecurityEngRespUnmarshalFailed, []byte(rejectDesc), []byte(correctiveAction))
	}
	*Data = decryptDataRespMsg.Data
	return 1, respData
}
