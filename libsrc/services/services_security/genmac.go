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
	"lmsapieng/libsrc/utils/trace"
)

func GenMAC(MACKey string, MACData string, MAC *string, MACKCV *string, secArgs ...string) (int, []byte) {
	var securityEngAddr httpdef.HttpServerAddr
	var respData []byte
	var respInfo msgdef.RespInfoStruct
	var rejectDesc, correctiveAction, securityID string
	var genMACReqMsg securitydef.GenMACReqMsgStruct
	var genMACRespMsg securitydef.GenMACRespMsgStruct

	if len(MACKey) == 0 {
		rejectDesc = "MACKey is NULL"
		return -1, msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_SecurityEngInvalidReq, []byte(rejectDesc), []byte(correctiveAction), rejectDesc)
	}
	if len(MACData) == 0 {
		rejectDesc = "MACData is NULL"
		return -1, msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_SecurityEngInvalidReq, []byte(rejectDesc), []byte(correctiveAction), rejectDesc)
	}
	if len(secArgs) == 0 {
		securityID = globaldef.NOT_INITIALIZED
	} else {
		securityID = secArgs[0]
	}
	genMACReqMsg.MACKey = MACKey
	genMACReqMsg.MACData = MACData
	reqMap := make(map[string]interface{})
	reqMap[securitydef.SecurityEngCommandJSONObj] = securitydef.GenMACCommand
	reqMap[securitydef.SecurityEngIDJSONObj] = securityID
	reqMap[securitydef.SecurityEngDataJSONObj] = genMACReqMsg
	reqData, _ := json.Marshal(&reqMap)

	if locateSecurityEngAddr(&securityEngAddr) < 0 {
		trace.Lg("locateSAFEngAddr() failed")
		return -1, msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_LocateSecurityEngAddrFailed, []byte(rejectDesc), []byte(correctiveAction), rejectDesc)
	}
	rval := msgutil.PostReq(moduledef.SecurityEngModule, securityEngAddr.ServerIpAddr, securityEngAddr.ServerPort, securityEngAddr.ServerTimeout, reqData, &respData)
	if rval < 0 {
		if rval == -2 {
			return -1, msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_SendToSecurityEngTimedOut, []byte(rejectDesc), []byte(correctiveAction), rejectDesc)
		}
		return -1, msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_SendToSecurityEngFailed, []byte(rejectDesc), []byte(correctiveAction), rejectDesc)
	}
	if msgutil.ParseResp(respData, &respInfo) < 0 {
		return -1, respData
	}
	err := json.Unmarshal(respInfo.RespInfo.RespData, &genMACRespMsg)
	if err != nil {
		return -1, msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_SecurityEngRespUnmarshalFailed, []byte(rejectDesc), []byte(correctiveAction), rejectDesc)
	}
	*MAC = genMACRespMsg.MAC
	*MACKCV = genMACRespMsg.MACKCV
	return 1, respData
}
