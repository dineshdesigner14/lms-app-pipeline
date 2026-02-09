package diagutil

import (
	"encoding/json"
	"fmt"
	"lmsapieng/include/common/diagdef"
	"lmsapieng/include/common/headerdef"
	"lmsapieng/include/common/moduledef"
	"lmsapieng/include/common/msgdef"
	"lmsapieng/include/common/nodeconfig"
	"lmsapieng/libsrc/utils/dtutil"
	"lmsapieng/libsrc/utils/msgutil"
	"lmsapieng/libsrc/utils/nodeutil"
	"lmsapieng/libsrc/utils/rtutil"
	"lmsapieng/libsrc/utils/serviceutil"
	"os"
)

func SendDiagMsg(msgType string) int {
	var lmsapiengreq diagdef.DiagEngReqMsgStruct
	var diagInfo diagdef.DiagInfoStruct
	var nodeInfo nodeconfig.NodeInfo
	var respData []byte
	var respInfo msgdef.RespInfoStruct
	currentNode := rtutil.GetCurrentNodeName()
	diagInfo.ServiceName = currentNode + "." + serviceutil.GetServiceName()
	diagInfo.ServiceID = fmt.Sprintf("%d", os.Getpid())
	diagInfo.StartDate = serviceutil.GetServiceStartDate()
	diagInfo.StartTime = serviceutil.GetServiceStartTime()
	diagInfo.MsgType = msgType
	diagInfo.MsgSize = fmt.Sprintf("%d", serviceutil.GetServiceQueue())
	if msgType == headerdef.App_Header_Value_HeartBeat {
		diagInfo.LastKeepaliveDate = dtutil.GetDate("DDMMYYYY")
		diagInfo.LastKeepaliveTime = dtutil.GetTime("HHMMSS")
	} else {
		diagInfo.LastRequestDate = dtutil.GetDate("DDMMYYYY")
		diagInfo.LastRequestTime = dtutil.GetTime("HHMMSS")
	}
	lmsapiengreq.Command = diagdef.DIAG_ENG_COMMAND_SERVICE_DIAG
	lmsapiengreq.DiagReqData, _ = json.Marshal(&diagInfo)
	reqData, _ := json.Marshal(&lmsapiengreq)
	if nodeutil.GetNodeInfo(currentNode, &nodeInfo) < 0 {
		//trace.Lg("GetNodeInfo() failed for Node:%s", currentNode)
		return -1
	}
	if msgutil.PostReq(moduledef.DiagEngModule, nodeInfo.DiagEngIpaddr, nodeInfo.DiagEngPort, nodeInfo.ServiceEngTimeout, reqData, &respData) < 0 {
		//trace.Lg("PostReq() failed for %s with Ipaddr:%s Port:%d", moduledef.DiagEngModule, nodeInfo.DiagEngIpaddr, nodeInfo.DiagEngPort)
		return -1
	}
	if msgutil.ParseResp(respData, &respInfo) < 0 {
		//trace.Lg("ParseResp() failed for for %s with Ipaddr:%s Port:%d", moduledef.DiagEngModule, nodeInfo.DiagEngIpaddr, nodeInfo.DiagEngPort)
		return -1
	}
	return 1
}
