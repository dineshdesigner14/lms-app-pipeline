package microsv_sequence

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

func GenReqInfo(seqRespInfo *seqdef.SeqReqInfoRespStruct, SeqParams ...string) (int, []byte) {
	var rejectDesc, correctiveAction string
	var seqengreqmsg seqdef.SeqEngReqMsgStruct
	var SeqEngAddr httpdef.HttpServerAddr
	var respData []byte
	var respInfo msgdef.RespInfoStruct

	seqengreqmsg.RequestNumFlag = "N"
	seqengreqmsg.RecordNumFlag = "N"
	seqengreqmsg.RRNFlag = "N"
	seqengreqmsg.StanFlag = "N"

	if len(SeqParams) >= 1 {
		seqengreqmsg.RequestNumFlag = SeqParams[0]
	}
	if len(SeqParams) >= 2 {
		seqengreqmsg.RecordNumFlag = SeqParams[1]
	}
	if len(SeqParams) >= 3 {
		seqengreqmsg.RRNFlag = SeqParams[2]
	}
	if len(SeqParams) >= 4 {
		seqengreqmsg.StanFlag = SeqParams[3]
	}

	ReqInfoMap := make(map[string]interface{})
	ReqInfoMap[seqdef.SeqEngReqTypeJSONObj] = seqdef.GenReqInfoCommand
	ReqInfoMap[seqdef.SeqEngDataJSONObj] = seqengreqmsg
	reqinfobuffer, _ := json.MarshalIndent(&ReqInfoMap, "", "\t")

	if locateSeqEngAddr(&SeqEngAddr) < 0 {
		rejectDesc = "locateSeqEngAddr() failed"
		correctiveAction = "Check For locateSeqEngAddr() failed"
		//trace.Lg("locateSeqEngAddr() failed")
		return -1, msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GenReqInfoErr, []byte(rejectDesc), []byte(correctiveAction), rejectDesc)
	}
	rval := msgutil.PostReq(moduledef.SeqEngModule, SeqEngAddr.ServerIpAddr, SeqEngAddr.ServerPort, SeqEngAddr.ServerTimeout, reqinfobuffer, &respData)
	if rval < 0 {
		if rval == -2 {
			rejectDesc = "SendToSeqEngTimedOut"
			correctiveAction = "Check For SendToSeqEngTimedOut"
			return -1, msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GenReqInfoErr, []byte(rejectDesc), []byte(correctiveAction), rejectDesc)
		}
		rejectDesc = "SendToSeqEngFailed"
		correctiveAction = "Check For SendToSeqEngFailed"
		return -1, msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GenReqInfoErr, []byte(rejectDesc), []byte(correctiveAction), rejectDesc)
	}
	if msgutil.ParseResp(respData, &respInfo) < 0 {
		rejectDesc = fmt.Sprintf("msgutil.ParseResp() failed for respData[%s]", respData)
		correctiveAction = fmt.Sprintf("Check For msgutil.ParseResp() failed for respData[%s]", respData)
		//trace.Lg("msgutil.ParseResp() failed for respData[%s]", respData)
		return -1, msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GenReqInfoErr, []byte(rejectDesc), []byte(correctiveAction), rejectDesc)
	}
	err := json.Unmarshal(respInfo.RespInfo.RespData, seqRespInfo)
	if err != nil {
		rejectDesc = fmt.Sprintf("json.Unmarshal failed for respData[%s] with err[%s]", respData, err)
		correctiveAction = fmt.Sprintf("Check For json.Unmarshal failed for respData[%s] with err[%s]", respData, err)
		//trace.Lg("json.Unmarshal failed for respData[%s] with err[%s]", respData, err)
		return -1, msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GenReqInfoErr, []byte(rejectDesc), []byte(correctiveAction), rejectDesc)
	}
	return 1, respData
}
