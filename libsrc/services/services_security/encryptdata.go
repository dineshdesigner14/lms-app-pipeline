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

func EncryptData(Data string, EData *string, secArgs ...string) (int, []byte) {
	var securityEngAddr httpdef.HttpServerAddr
	var respData []byte
	var respInfo msgdef.RespInfoStruct
	var rejectDesc, correctiveAction, securityID string
	var encryptDataReqMsg securitydef.EncryptDataReqMsgStruct
	var encryptDataRespMsg securitydef.EncryptDataRespMsgStruct

	if len(Data) == 0 {
		rejectDesc = "EncryptData is NULL"
		correctiveAction = "Check the EncryptData Used for Encryption for NULL"
		return -1, msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_SecurityEngInvalidReq, []byte(rejectDesc), []byte(correctiveAction), rejectDesc)
	}
	if len(secArgs) == 0 {
		securityID = globaldef.NOT_INITIALIZED
	} else {
		securityID = secArgs[0]
	}
	encryptDataReqMsg.Data = Data
	reqMap := make(map[string]interface{})
	reqMap[securitydef.SecurityEngCommandJSONObj] = securitydef.EncryptDataCommand
	reqMap[securitydef.SecurityEngIDJSONObj] = securityID
	reqMap[securitydef.SecurityEngDataJSONObj] = encryptDataReqMsg
	reqData, _ := json.Marshal(&reqMap)

	if locateSecurityEngAddr(&securityEngAddr) < 0 {
		//trace.Lg("locateSecurityEngAddr() failed")
		rejectDesc = "locateSecurityEngAddr Error"
		correctiveAction = "Check the SecurityEngAddr"
		return -1, msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_LocateSecurityEngAddrFailed, []byte(rejectDesc), []byte(correctiveAction), rejectDesc)
	}
	rval := msgutil.PostReq(moduledef.SecurityEngModule, securityEngAddr.ServerIpAddr, securityEngAddr.ServerPort, securityEngAddr.ServerTimeout, reqData, &respData)
	if rval < 0 {
		if rval == -2 {
			rejectDesc = "SecurityEng TimedOut Error"
			correctiveAction = "Check Why SecurityEng TimedOut Error Occured"
			return -1, msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_SendToSecurityEngTimedOut, []byte(rejectDesc), []byte(correctiveAction), rejectDesc)
		}
		rejectDesc = "SecurityEng Send Error"
		correctiveAction = "Check Why SecurityEng Send Error Occured"
		return -1, msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_SendToSecurityEngFailed, []byte(rejectDesc), []byte(correctiveAction), rejectDesc)
	}
	if msgutil.ParseResp(respData, &respInfo) < 0 {
		return -1, respData
	}
	err := json.Unmarshal(respInfo.RespInfo.RespData, &encryptDataRespMsg)
	if err != nil {
		rejectDesc = "SecurityEng Response Unmarshal Error"
		correctiveAction = "Check SecurityEng Response Unmarshal Error Occured"
		return -1, msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_SecurityEngRespUnmarshalFailed, []byte(rejectDesc), []byte(correctiveAction), rejectDesc)
	}
	*EData = encryptDataRespMsg.EData
	return 1, respData
}
