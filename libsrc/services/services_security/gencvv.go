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

func GenCVV(CardNum string, ExpDate string, ServiceCode string, CVK1 string, CVK2 string, CVV *string, secArgs ...string) (int, []byte) {
	var securityEngAddr httpdef.HttpServerAddr
	var respData []byte
	var respInfo msgdef.RespInfoStruct
	var rejectDesc, correctiveAction, securityID string
	var genCVVReqMsg securitydef.GenCVVReqMsgStruct
	var genCVVRespMsg securitydef.GenCVVRespMsgStruct

	//trace.Lg("GenCVV() called for CardNum[%s] ExpDate[%s] ServiceCode[%s] CVK1[%s] CVK2[%s]", CardNum, ExpDate, ServiceCode, CVK1, CVK2)
	if len(CardNum) == 0 {
		rejectDesc = "CardNum is NULL"
		//trace.Lg("CardNum is NULL")
		return -1, msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_SecurityEngInvalidReq, []byte(rejectDesc), []byte(correctiveAction), rejectDesc)
	}
	if len(ExpDate) == 0 {
		rejectDesc = "ExpDate is NULL"
		//trace.Lg("ExpDate is NULL")
		return -1, msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_SecurityEngInvalidReq, []byte(rejectDesc), []byte(correctiveAction), rejectDesc)
	}
	if len(ServiceCode) == 0 {
		rejectDesc = "ServiceCode is NULL"
		//trace.Lg("ServiceCode is NULL")
		return -1, msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_SecurityEngInvalidReq, []byte(rejectDesc), []byte(correctiveAction), rejectDesc)
	}
	if !cryptoutil.IsValidKey(CVK1) {
		//trace.Lg("CVK1 is Invalid")
		rejectDesc = "CVK1 is Invalid"
		return -1, msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_SecurityEngInvalidReq, []byte(rejectDesc), []byte(correctiveAction), rejectDesc)
	}
	if !cryptoutil.IsValidKey(CVK2) {
		//trace.Lg("CVK2 is Invalid")
		rejectDesc = "CVK2 is Invalid"
		return -1, msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_SecurityEngInvalidReq, []byte(rejectDesc), []byte(correctiveAction), rejectDesc)
	}
	if len(secArgs) == 0 {
		securityID = globaldef.NOT_INITIALIZED
	} else {
		securityID = secArgs[0]
	}
	genCVVReqMsg.CardNum = CardNum
	genCVVReqMsg.ExpDate = ExpDate
	genCVVReqMsg.ServiceCode = ServiceCode
	genCVVReqMsg.CVK1 = CVK1
	genCVVReqMsg.CVK2 = CVK2

	reqMap := make(map[string]interface{})
	reqMap[securitydef.SecurityEngCommandJSONObj] = securitydef.GenCVVCommand
	reqMap[securitydef.SecurityEngIDJSONObj] = securityID
	reqMap[securitydef.SecurityEngDataJSONObj] = genCVVReqMsg
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
	err := json.Unmarshal(respInfo.RespInfo.RespData, &genCVVRespMsg)
	if err != nil {
		return -1, msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_SecurityEngRespUnmarshalFailed, []byte(rejectDesc), []byte(correctiveAction), rejectDesc)
	}
	*CVV = genCVVRespMsg.CVV
	return 1, respData
}
