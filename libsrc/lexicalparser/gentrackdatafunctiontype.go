package lexicalparser

import (
	"fmt"
	"lmsapieng/include/common/carddef"
	"lmsapieng/include/common/lexicalparserdef"
	"lmsapieng/include/common/rejectdef"
	"lmsapieng/include/common/reqbrokerdef"
	"lmsapieng/include/common/trackobjdef"
	"lmsapieng/libsrc/utils/cardutil"
	"lmsapieng/libsrc/utils/datatypeutil"
	"lmsapieng/libsrc/utils/dtutil"
	"lmsapieng/libsrc/utils/mathutil"
	"lmsapieng/libsrc/utils/msgutil"
	"lmsapieng/libsrc/utils/templateutil"
)

func genTrackData(reqBrokerDataMap map[string]interface{}, trackObject map[string]interface{}) int {
	_, ok := trackObject[trackobjdef.CardNumKey]
	if !ok {
		//trace.Lg(reqbrokerutil.GetReqID(reqBrokerDataMap), "[%s] not present in trackObject", trackobjdef.CardNumKey)
		return -1
	}
	if !datatypeutil.IsString(trackObject[trackobjdef.CardNumKey]) {
		//trace.Lg(reqbrokerutil.GetReqID(reqBrokerDataMap), "%s trackObject has invalid type[%T]", trackobjdef.CardNumKey, trackObject[trackobjdef.CardNumKey])
		return -1
	}
	CardNum := trackObject[trackobjdef.CardNumKey].(string)

	if len(CardNum) != 16 && len(CardNum) != 19 {
		//trace.LgReq(reqbrokerutil.GetReqID(reqBrokerDataMap), "CardNum Length should 16 or 19")
		return -1
	}
	if !mathutil.IsValidNumericString(CardNum) {
		//trace.LgReq(reqbrokerutil.GetReqID(reqBrokerDataMap), "CardNum[%s] Should be Numeric", CardNum)
		return -1
	}

	_, ok = trackObject[trackobjdef.EncodingNameKey]
	if !ok {
		//trace.Lg(reqbrokerutil.GetReqID(reqBrokerDataMap), "[%s] not present in trackObject", trackobjdef.EncodingNameKey)
		return -1
	}
	if !datatypeutil.IsString(trackObject[trackobjdef.EncodingNameKey]) {
		//trace.Lg(reqbrokerutil.GetReqID(reqBrokerDataMap), "%s trackObject has invalid type[%T]", trackobjdef.EncodingNameKey, trackObject[trackobjdef.EncodingNameKey])
		return -1
	}
	EncodingName := trackObject[trackobjdef.EncodingNameKey].(string)

	_, ok = trackObject[trackobjdef.ExpDateKey]
	if !ok {
		//trace.Lg(reqbrokerutil.GetReqID(reqBrokerDataMap), "[%s] not present in trackObject", trackobjdef.ExpDateKey)
		return -1
	}
	if !datatypeutil.IsString(trackObject[trackobjdef.ExpDateKey]) {
		//trace.Lg(reqbrokerutil.GetReqID(reqBrokerDataMap), "%s trackObject has invalid type[%T]", trackobjdef.ExpDateKey, trackObject[trackobjdef.ExpDateKey])
		return -1
	}
	ExpDate := trackObject[trackobjdef.ExpDateKey].(string)

	_, ok = trackObject[trackobjdef.ServiceCodeKey]
	if !ok {
		//trace.Lg(reqbrokerutil.GetReqID(reqBrokerDataMap), "[%s] not present in trackObject", trackobjdef.ServiceCodeKey)
		return -1
	}
	if !datatypeutil.IsString(trackObject[trackobjdef.ServiceCodeKey]) {
		//trace.Lg(reqbrokerutil.GetReqID(reqBrokerDataMap), "%s trackObject has invalid type[%T]", trackobjdef.ServiceCodeKey, trackObject[trackobjdef.ServiceCodeKey])
		return -1
	}
	ServiceCode := trackObject[trackobjdef.ServiceCodeKey].(string)

	if len(ServiceCode) != 3 {
		//trace.LgReq(reqbrokerutil.GetReqID(reqBrokerDataMap), "CardNum Length should 3")
		return -1
	}
	if !mathutil.IsValidNumericString(CardNum) {
		//trace.LgReq(reqbrokerutil.GetReqID(reqBrokerDataMap), "ServiceCode[%s] Should be Numeric", ServiceCode)
		return -1
	}

	_, ok = trackObject[trackobjdef.CVV1Key]
	if !ok {
		//trace.Lg(reqbrokerutil.GetReqID(reqBrokerDataMap), "[%s] not present in trackObject", trackobjdef.CVV1Key)
		return -1
	}
	if !datatypeutil.IsString(trackObject[trackobjdef.CVV1Key]) {
		//trace.Lg(reqbrokerutil.GetReqID(reqBrokerDataMap), "%s trackObject has invalid type[%T]", trackobjdef.CVV1Key, trackObject[trackobjdef.CVV1Key])
		return -1
	}
	CVV1 := trackObject[trackobjdef.CVV1Key].(string)

	if len(CVV1) != 3 {
		//trace.LgReq(reqbrokerutil.GetReqID(reqBrokerDataMap), "CVV1 Length should 3")
		return -1
	}
	if !mathutil.IsValidNumericString(CVV1) {
		//trace.LgReq(reqbrokerutil.GetReqID(reqBrokerDataMap), "CVV1[%s] Should be Numeric", CVV1)
		return -1
	}

	_, ok = trackObject[trackobjdef.CVV2Key]
	if !ok {
		//trace.Lg(reqbrokerutil.GetReqID(reqBrokerDataMap), "[%s] not present in trackObject", trackobjdef.CVV2Key)
		return -1
	}
	if !datatypeutil.IsString(trackObject[trackobjdef.CVV2Key]) {
		//trace.Lg(reqbrokerutil.GetReqID(reqBrokerDataMap), "%s trackObject has invalid type[%T]", trackobjdef.CVV2Key, trackObject[trackobjdef.CVV2Key])
		return -1
	}

	if len(trackObject[trackobjdef.CVV2Key].(string)) != 3 {
		//trace.LgReq(reqbrokerutil.GetReqID(reqBrokerDataMap), "CVV2 Length should 3")
		return -1
	}
	if !mathutil.IsValidNumericString(trackObject[trackobjdef.CVV2Key].(string)) {
		//trace.LgReq(reqbrokerutil.GetReqID(reqBrokerDataMap), "CVV2[%s] Should be Numeric", trackObject[trackobjdef.CVV2Key].(string))
		return -1
	}

	_, ok = trackObject[trackobjdef.ICVVKey]
	if !ok {
		//trace.Lg(reqbrokerutil.GetReqID(reqBrokerDataMap), "[%s] not present in trackObject", trackobjdef.ICVVKey)
		return -1
	}
	if !datatypeutil.IsString(trackObject[trackobjdef.ICVVKey]) {
		//trace.Lg(reqbrokerutil.GetReqID(reqBrokerDataMap), "%s trackObject has invalid type[%T]", trackobjdef.ICVVKey, trackObject[trackobjdef.ICVVKey])
		return -1
	}

	if len(trackObject[trackobjdef.ICVVKey].(string)) != 3 {
		//trace.LgReq(reqbrokerutil.GetReqID(reqBrokerDataMap), "ICVV Length should 3")
		return -1
	}
	if !mathutil.IsValidNumericString(trackObject[trackobjdef.ICVVKey].(string)) {
		//trace.LgReq(reqbrokerutil.GetReqID(reqBrokerDataMap), "ICVV[%s] Should be Numeric", trackObject[trackobjdef.ICVVKey].(string))
		return -1
	}

	_, ok = trackObject[trackobjdef.ServerFlagKey]
	if !ok {
		//trace.LgReq(reqbrokerutil.GetReqID(reqBrokerDataMap), "[%s] not present in trackObject", trackobjdef.ServerFlagKey)
		return -1
	}
	if !datatypeutil.IsString(trackObject[trackobjdef.ServerFlagKey]) {
		//trace.LgReq(reqbrokerutil.GetReqID(reqBrokerDataMap), "%s trackObject has invalid type[%T]", trackobjdef.ServerFlagKey, trackObject[trackobjdef.ServerFlagKey])
		return -1
	}
	ServerFlag := trackObject[trackobjdef.ServerFlagKey].(string)
	if ServerFlag != "0" && ServerFlag != "1" {
		//trace.LgReq(reqbrokerutil.GetReqID(reqBrokerDataMap), "%s trackObject should be 0 or 1", trackObject[trackobjdef.ServerFlagKey])
		return -1
	}

	ExpDate = ExpDate[6:] + ExpDate[2:4]
	Track1 := ""
	if cardutil.ComposeTrack1Data(CardNum, EncodingName, ExpDate, ServiceCode, CVV1, carddef.Track1DataSize, &Track1) < 0 {
		//trace.Lg("composeTrack1Data() failed")
		return -1
	}
	if len(ExpDate) != 4 {
		//trace.LgReq(reqbrokerutil.GetReqID(reqBrokerDataMap), "ExpDate length Should be 4")
		return -1
	}
	if !dtutil.IsValidDateFormat(ExpDate, "YYMM") {
		//trace.LgReq(reqbrokerutil.GetReqID(reqBrokerDataMap), "ExpiryDate[%s] Should be YYMM", ExpDate)
		return -1
	}

	Track2 := ""
	if cardutil.ComposeTrack2Data(CardNum, ExpDate, ServiceCode, CVV1, ServerFlag, carddef.Track2DataSize, &Track2) < 0 {
		//trace.Lg("composeTrack2Data() failed")
		return -1
	}
	trackObject[trackobjdef.Track1Key] = Track1
	trackObject[trackobjdef.Track2Key] = Track2
	return 1
}

func GenTrackDataFunctionType(execFunction lexicalparserdef.ExecFunctionStruct, reqBrokerDataMap map[string]interface{}) int {
	var rejectDesc string
	if len(execFunction.GenTrackDataInfo.ObjectName) == 0 {
		rejectDesc = fmt.Sprintf("ObjectName is NULL in GenTrackDataFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GenTrackDataFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenTrackDataInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenTrackDataInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	_, ok := reqBrokerDataMap[execFunction.GenTrackDataInfo.ObjectName]
	if !ok {
		rejectDesc = fmt.Sprintf("ObjectName[%s] does not exist in GenTrackDataFunctionType for FunctionName[%s]", execFunction.GenTrackDataInfo.ObjectName, execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GenTrackDataFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenTrackDataInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenTrackDataInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	tempType := fmt.Sprintf("%T", reqBrokerDataMap[execFunction.GenTrackDataInfo.ObjectName])
	if tempType != "map[string]interface {}" {
		rejectDesc = fmt.Sprintf("ObjectName[%s] is invalid in GenTrackDataFunctionType for FunctionName[%s]", execFunction.GenTrackDataInfo.ObjectName, execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GenTrackDataFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenTrackDataInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenTrackDataInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	if genTrackData(reqBrokerDataMap, reqBrokerDataMap[execFunction.GenTrackDataInfo.ObjectName].(map[string]interface{})) < 0 {
		rejectDesc = fmt.Sprintf("genTrackData() failed for ObjectName[%s] in GenTrackDataFunctionType for FunctionName[%s]", execFunction.GenTrackDataInfo.ObjectName, execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GenTrackDataFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenTrackDataInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenTrackDataInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	return 1
}
