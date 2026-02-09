package services_sequence

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"lmsapieng/include/common/globaldef"
	"lmsapieng/include/common/httpdef"
	"lmsapieng/include/common/moduledef"
	"lmsapieng/include/common/msgdef"
	"lmsapieng/include/common/rejectdef"
	"lmsapieng/include/common/seqdef"
	"lmsapieng/libsrc/utils/msgutil"
	"os"
)

var locateSeqEngFlag = false
var addrInfo seqdef.SeqEngAddrInfo

func locateSeqEngAddr(SeqEngAddr *httpdef.HttpServerAddr) int {
	if !locateSeqEngFlag {
		configFile := fmt.Sprintf("%s/config/seqeng/seqeng_addr.xml", globaldef.GetAppBaseDir())
		xmlFile, err := os.Open(configFile)
		if err != nil {
			//trace.Lg("os.Open() failed for configFile(%s)...locateSeqEngAddr() Failed", configFile)
			return -1
		}
		defer xmlFile.Close()
		byteValue, _ := ioutil.ReadAll(xmlFile)
		err = xml.Unmarshal(byteValue, &addrInfo)
		if err != nil {
			//trace.Lg("xml.Unmarshal() failed for configFile(%s)...locateSeqEngAddr() Failed", configFile)
			return -1
		}
		locateSeqEngFlag = true
	}
	SeqEngAddr.ServerIpAddr = addrInfo.IpAddr
	SeqEngAddr.ServerPort = addrInfo.Port
	SeqEngAddr.ServerTimeout = addrInfo.Timeout
	return 1
}

func GenNextSequence(RelType string, SeqParams ...string) (int, []byte, string) {
	var EntityType, EntityID, rejectDesc, correctiveAction string
	var seqendmsg seqdef.SeqEngMsgStruct
	var SeqEngAddr httpdef.HttpServerAddr
	var respData []byte
	var respInfo msgdef.RespInfoStruct

	if len(RelType) == 0 {
		rejectDesc = "RelType is Null"
		correctiveAction = "Check For RelType is Null"
		//trace.Lg("RelType is Null")
		return -1, msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GenSequenceErr, []byte(rejectDesc), []byte(correctiveAction), rejectDesc), ""
	}
	if len(SeqParams) != 0 && len(SeqParams) != 2 {
		rejectDesc = "GenNextSequence() should have 2 or 4 arguments"
		correctiveAction = "Check For GenNextSequence() should have 2 or 4 arguments"
		//trace.Lg("GenNextSequence() should have 2 or 4 arguments")
		return -1, msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GenSequenceErr, []byte(rejectDesc), []byte(correctiveAction), rejectDesc), ""
	}
	if len(SeqParams) == 0 {
		EntityType = RelType
		EntityID = RelType
	} else {
		EntityType = SeqParams[0]
		EntityID = SeqParams[1]
	}
	seqendmsg.EntityType = EntityType
	seqendmsg.EntityID = EntityID
	seqendmsg.RelType = RelType
	ReqInfoMap := make(map[string]interface{})
	ReqInfoMap[seqdef.SeqEngReqTypeJSONObj] = seqdef.GenSequenceCommand
	ReqInfoMap[seqdef.SeqEngDataJSONObj] = seqendmsg
	reqinfobuffer, _ := json.MarshalIndent(&ReqInfoMap, "", "\t")

	if locateSeqEngAddr(&SeqEngAddr) < 0 {
		rejectDesc = "locateSeqEngAddr() failed"
		correctiveAction = "Check For locateSeqEngAddr() failed"
		//trace.Lg("locateSeqEngAddr() failed")
		return -1, msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GenSequenceErr, []byte(rejectDesc), []byte(correctiveAction), rejectDesc), ""
	}
	rval := msgutil.PostReq(moduledef.SeqEngModule, SeqEngAddr.ServerIpAddr, SeqEngAddr.ServerPort, SeqEngAddr.ServerTimeout, reqinfobuffer, &respData)
	if rval < 0 {
		if rval == -2 {
			rejectDesc = "SendToSeqEngTimedOut"
			correctiveAction = "Check For SendToSeqEngTimedOut"
			return -1, msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GenSequenceErr, []byte(rejectDesc), []byte(correctiveAction), rejectDesc), ""
		}
		rejectDesc = "SendToSeqEngFailed"
		correctiveAction = "Check For SendToSeqEngFailed"
		return -1, msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GenSequenceErr, []byte(rejectDesc), []byte(correctiveAction), rejectDesc), ""
	}
	if msgutil.ParseResp(respData, &respInfo) < 0 {
		return -1, respData, ""
	}
	respDataMap := make(map[string]interface{})
	json.Unmarshal(respInfo.RespInfo.RespData, &respDataMap)
	_, ok := respDataMap[seqdef.SeqEngSeqNumObj]
	if !ok {
		rejectDesc = "NoSeqNumObjInResp"
		correctiveAction = "Check For NoSeqNumObjInResp"
		return -1, msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GenSequenceErr, []byte(rejectDesc), []byte(correctiveAction), rejectDesc), ""
	}
	tempType := fmt.Sprintf("%T", respDataMap[seqdef.SeqEngSeqNumObj])
	if tempType != "string" {
		rejectDesc = "InValidSeqNumTypeInResp"
		correctiveAction = "Check For InValidSeqNumTypeInResp"
		return -1, msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GenSequenceErr, []byte(rejectDesc), []byte(correctiveAction), rejectDesc), ""
	}
	return 1, respData, respDataMap[seqdef.SeqEngSeqNumObj].(string)
}
