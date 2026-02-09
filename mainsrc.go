package main

import (
	"lmsapieng/include/common/debugdef"
	"lmsapieng/include/common/globaldef"
	"lmsapieng/include/common/headerdef"
	"lmsapieng/include/common/moduledef"
	"lmsapieng/libsrc/utils/diagutil"
	"lmsapieng/libsrc/utils/memusageutil"
	"lmsapieng/libsrc/utils/serviceutil"
	"lmsapieng/libsrc/utils/trace"

	"io/ioutil"
	"net/http"
	"os"
)

func init() {
	if serviceutil.InitService(moduledef.LMSApiEngModule, mainSrcVersion) < 0 {
		os.Exit(globaldef.EXIT_INIT_FAILED)
	}
	serviceutil.SendInitMsg(globaldef.STATUS_OK)
}

func handleServiceNotAvailable(HTTPResponseWriter http.ResponseWriter, HTTPRequest *http.Request) {
}

func handleAPIResponse(HTTPResponseWriter http.ResponseWriter, lrespbuffer *[]byte) {
	trace.Log(debugdef.DEBUG_LEVEL_TEST, "Sending Response <%s>", *lrespbuffer)
	HTTPResponseWriter.Header().Set("Content-Type", "application/json")
	HTTPResponseWriter.WriteHeader(http.StatusOK)
	HTTPResponseWriter.Write(*lrespbuffer)
}

func handleAPIReq(HTTPResponseWriter http.ResponseWriter, HTTPRequest *http.Request) {
	var lrespbuffer []byte
	var err error
	memusageutil.LogMemUsage()
	//trace.Lg("handleAPIReq() called..........")
	// httprequest, _ := httputil.DumpRequest(HTTPRequest, true)
	//trace.LgHex(httprequest)
	// //trace.LogHex(debugdef.DEBUG_LEVEL_TEST, httprequest)
	diagutil.SendDiagMsg(HTTPRequest.Header.Get(headerdef.App_Header_Type_MsgType))
	if serviceutil.IsServiceIntimationMsg(HTTPResponseWriter, HTTPRequest) {
		return
	}
	if serviceutil.IsServiceStopFlagTrue() {
		handleServiceNotAvailable(HTTPResponseWriter, HTTPRequest)
		return
	}
	requestBody, err := ioutil.ReadAll(HTTPRequest.Body)
	if err != nil {
		//trace.Lg("ioutil.ReadAll() failed for request")
		return
	}
	trace.Log(debugdef.DEBUG_LEVEL_TEST, "Got Request(%s)", requestBody)
	defer handleAPIResponse(HTTPResponseWriter, &lrespbuffer)
	serviceutil.IncrementServiceQueue()
	defer serviceutil.DecrementServiceQueue()
	lrespbuffer = execMicroServices(requestBody, HTTPRequest.Context())
}

func main() {
	http.HandleFunc("/", handleAPIReq)
	serviceutil.StartHttpServer()
}
