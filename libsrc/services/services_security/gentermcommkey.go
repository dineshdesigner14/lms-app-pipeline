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

func GenTermCommKey(TMK string, TPKUTMK *string, TPKULMK *string, TPKKCV *string, secArgs ...string) (int, []byte) {
	var securityEngAddr httpdef.HttpServerAddr
	var respData []byte
	var respInfo msgdef.RespInfoStruct
	var rejectDesc, correctiveAction, securityID string
	var genTermCommKeyReqMsg securitydef.GenTermCommKeyReqMsgStruct
	var genTermCommKeyRespMsg securitydef.GenTermCommKeyRespMsgStruct

	if len(TMK) == 0 {
		rejectDesc = "TMK is NULL"
		return -1, msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_SecurityEngInvalidReq, []byte(rejectDesc), []byte(correctiveAction), rejectDesc)
	}
	if !cryptoutil.IsValidKey(TMK) {
		rejectDesc = "TMK is Invalid"
		return -1, msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_SecurityEngInvalidReq, []byte(rejectDesc), []byte(correctiveAction), rejectDesc)
	}
	if len(secArgs) == 0 {
		securityID = globaldef.NOT_INITIALIZED
	} else {
		securityID = secArgs[0]
	}
	genTermCommKeyReqMsg.TMK = TMK
	reqMap := make(map[string]interface{})
	reqMap[securitydef.SecurityEngCommandJSONObj] = securitydef.GenTermCommKeyCommand
	reqMap[securitydef.SecurityEngIDJSONObj] = securityID
	reqMap[securitydef.SecurityEngDataJSONObj] = genTermCommKeyReqMsg
	reqData, _ := json.Marshal(&reqMap)

	if locateSecurityEngAddr(&securityEngAddr) < 0 {
		//trace.Lg("locateSAFEngAddr() failed")
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
	err := json.Unmarshal(respInfo.RespInfo.RespData, &genTermCommKeyRespMsg)
	if err != nil {
		return -1, msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_SecurityEngRespUnmarshalFailed, []byte(rejectDesc), []byte(correctiveAction), rejectDesc)
	}
	*TPKUTMK = genTermCommKeyRespMsg.TPKUTMK
	*TPKULMK = genTermCommKeyRespMsg.TPKULMK
	*TPKKCV = genTermCommKeyRespMsg.TPKKCV
	return 1, respData
}
