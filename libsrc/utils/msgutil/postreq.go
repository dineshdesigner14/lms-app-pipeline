package msgutil

import (
	"fmt"
	"lmsapieng/libsrc/utils/httplib"
	"lmsapieng/libsrc/utils/tlsconfigutil"
)

func PostReq(ModuleName string, IpAddr string, PortNum int, TimeOut int, ReqData []byte, RespData *[]byte, headerList ...string) int {
	var httpStatus int
	var serverCrt, serverName string
	if !tlsconfigutil.IsTLSReq(IpAddr, fmt.Sprintf("%d", PortNum), &serverCrt, &serverName) {
		URL := fmt.Sprintf("http://%s:%d/%s", IpAddr, PortNum, ModuleName)
		rval := httplib.SendPOSTRequest(ReqData, URL, "application/json", TimeOut, RespData, &httpStatus, headerList...)
		if rval < 0 {
			//trace.Lg("SendPOSTRequest() failed for URL(%s)", URL)
		}
		return rval
	}
	URL := fmt.Sprintf("https://%s:%d/%s", IpAddr, PortNum, ModuleName)
	rval := httplib.SendPOSTTLSRequest(serverName, serverCrt, ReqData, URL, "application/json", TimeOut, RespData, &httpStatus, headerList...)
	if rval < 0 {
		//trace.Lg("SendPOSTRequest() failed for URL(%s)", URL)
	}
	return rval
}
