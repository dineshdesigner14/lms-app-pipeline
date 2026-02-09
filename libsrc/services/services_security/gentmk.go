package services_security

import (
	"encoding/json"
	"lmsapieng/include/common/globaldef"
	"lmsapieng/include/common/httpdef"
	"lmsapieng/include/common/moduledef"
	"lmsapieng/include/common/msgdef"
	"lmsapieng/include/common/rejectdef"
	"lmsapieng/include/common/securitydef"
	"lmsapieng/libsrc/utils/cryptoutil"
	"lmsapieng/libsrc/utils/msgutil"
)

func GenTMK(Component1 string, Component2 string, eTMK *string, TMKKCV *string, Component1KCV *string, Component2KCV *string, secArgs ...string) (int, []byte) {
	var securityEngAddr httpdef.HttpServerAddr
	var respData []byte
	var respInfo msgdef.RespInfoStruct
	var rejectDesc, correctiveAction, securityID string
	var genTMKReqMsg securitydef.GenTMKReqMsgStruct
	var genTMKRespMsg securitydef.GenTMKRespMsgStruct

	if len(Component1) == 0 {
		rejectDesc = "Component1 is NULL"
		return -1, msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_SecurityEngInvalidReq, []byte(rejectDesc), []byte(correctiveAction), rejectDesc)
	}
	if !cryptoutil.IsValidKey(Component1) {
		rejectDesc = "Component1 is Invalid"
		return -1, msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_SecurityEngInvalidReq, []byte(rejectDesc), []byte(correctiveAction), rejectDesc)
	}
	if len(Component2) == 0 {
		rejectDesc = "Component2 is NULL"
		return -1, msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_SecurityEngInvalidReq, []byte(rejectDesc), []byte(correctiveAction), rejectDesc)
	}
	if !cryptoutil.IsValidKey(Component2) {
		rejectDesc = "Component2 is Invalid"
		return -1, msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_SecurityEngInvalidReq, []byte(rejectDesc), []byte(correctiveAction), rejectDesc)
	}
	if len(secArgs) == 0 {
		securityID = globaldef.NOT_INITIALIZED
	} else {
		securityID = secArgs[0]
	}
	genTMKReqMsg.Component1 = Component1
	genTMKReqMsg.Component2 = Component2
	reqMap := make(map[string]interface{})
	reqMap[securitydef.SecurityEngCommandJSONObj] = securitydef.GenTMKCommand
	reqMap[securitydef.SecurityEngIDJSONObj] = securityID
	reqMap[securitydef.SecurityEngDataJSONObj] = genTMKReqMsg
	reqData, _ := json.Marshal(&reqMap)

	if locateSecurityEngAddr(&securityEngAddr) < 0 {
		//trace.Lg("locateSecurityEngAddr() failed")
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
	err := json.Unmarshal(respInfo.RespInfo.RespData, &genTMKRespMsg)
	if err != nil {
		return -1, msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_SecurityEngRespUnmarshalFailed, []byte(rejectDesc), []byte(correctiveAction), rejectDesc)
	}
	*eTMK = genTMKRespMsg.ETMK
	*TMKKCV = genTMKRespMsg.TMKKCV
	*Component1KCV = genTMKRespMsg.Component1KCV
	*Component2KCV = genTMKRespMsg.Component2KCV
	return 1, respData
}
