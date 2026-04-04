package serviceutil

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lmsapieng/include/common/debugdef"
	"lmsapieng/include/common/globaldef"
	"lmsapieng/include/common/headerdef"
	"lmsapieng/include/common/moduledef"
	"lmsapieng/include/common/msgdef"
	"lmsapieng/include/common/nodeconfig"
	"lmsapieng/include/common/servicedef"
	"lmsapieng/libsrc/utils/dtutil"
	"lmsapieng/libsrc/utils/genutil"
	"lmsapieng/libsrc/utils/jsonutil"
	"lmsapieng/libsrc/utils/msgutil"
	"lmsapieng/libsrc/utils/nodeutil"
	"lmsapieng/libsrc/utils/osutil"
	"lmsapieng/libsrc/utils/rtutil"
	"lmsapieng/libsrc/utils/trace"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var listenPort string
var serviceName string
var serviceStartDate string
var serviceStartTime string
var serviceQueueNum int
var serviceQueueMutex = &sync.Mutex{}
var serviceStopFlag bool

func InitService(ModuleName string, ModuleVersion string, ServiceArgs ...string) int {
	if !globaldef.IsAppBaseDirExists() {
		fmt.Printf("\nIsAppBaseDirExists() failed.failed....\n")
		return -1
	}
	serviceName = osutil.GetServiceName()
	if trace.OpenTrace(serviceName, serviceName) < 0 {
		fmt.Printf("\nOpenTrace() failed for processName(%s)\n", serviceName)
		return -1
	}
	genutil.SetModule(serviceName, ModuleName, ModuleVersion)
	var rval int
	rval, listenPort = genutil.GetListeningPort()
	if rval < 0 {
		return -1
	}
	if rtutil.LoadAppConfig() < 0 {
		return -1
	}
	trace.SetTraceLevel(rtutil.GetServiceDebugLevel(serviceName))
	serviceQueueNum = 0
	serviceStopFlag = false
	serviceStartDate = dtutil.GetDate("DDMMYYYY")
	serviceStartTime = dtutil.GetTime("HHMMSS")
	return 1
}

func GetServiceStartDate() string {
	return serviceStartDate
}

func GetServiceStartTime() string {
	return serviceStartTime
}

func GetServiceName() string {
	return serviceName
}

func GetListenPort() string {
	return listenPort
}

func SendInitMsg(Status string) int {
	var serviceEngReq servicedef.ServiceEngReqMsg
	var respData []byte
	var nodeInfo nodeconfig.NodeInfo
	var respInfo msgdef.RespInfoStruct

	serviceEngReq.Command = servicedef.ServiceEngCommand_StartServiceAck
	serviceEngReq.ServiceName = GetServiceName()
	serviceEngReq.ServiceID = fmt.Sprintf("%d", os.Getpid())
	serviceEngReq.StartDate = dtutil.GetDate("DDMMYYYY")
	serviceEngReq.StartTime = dtutil.GetTime("HHMMSS")
	serviceEngReq.Status = Status
	tLevel := trace.GetTraceLevel()
	switch tLevel {
	case debugdef.DEBUG_LEVEL_NORMAL:
		serviceEngReq.DebugLevel = debugdef.DEBUG_LEVEL_NORMAL_STR
	case debugdef.DEBUG_LEVEL_SECURED:
		serviceEngReq.DebugLevel = debugdef.DEBUG_LEVEL_SECURED_STR
	case debugdef.DEBUG_LEVEL_TEST:
		serviceEngReq.DebugLevel = debugdef.DEBUG_LEVEL_TEST_STR
	case debugdef.DEBUG_LEVEL_ERROR:
		serviceEngReq.DebugLevel = debugdef.DEBUG_LEVEL_ERROR_STR
	default:
		serviceEngReq.DebugLevel = globaldef.NOT_INITIALIZED
	}

	reqData, _ := json.Marshal(&serviceEngReq)
	if nodeutil.GetNodeInfo(rtutil.GetCurrentNodeName(), &nodeInfo) < 0 {
		return -1
	}
	if msgutil.PostReq(moduledef.ServiceEngModule, nodeInfo.Ipaddr, nodeInfo.ServiceEngPort, nodeInfo.ServiceEngTimeout, reqData, &respData) < 0 {
		return -1
	}
	if msgutil.ParseResp(respData, &respInfo) < 0 {
		return -1
	}
	return 1
}

func IncrementServiceQueue() {
	serviceQueueMutex.Lock()
	defer serviceQueueMutex.Unlock()
	serviceQueueNum++
}

func DecrementServiceQueue() {
	serviceQueueMutex.Lock()
	defer serviceQueueMutex.Unlock()
	serviceQueueNum--
}

func GetServiceQueue() int {
	serviceQueueMutex.Lock()
	defer serviceQueueMutex.Unlock()
	return serviceQueueNum
}

func SetServiceStopFlag() {
	serviceQueueMutex.Lock()
	defer serviceQueueMutex.Unlock()
	serviceStopFlag = true
}

func IsServiceStopFlagTrue() bool {
	serviceQueueMutex.Lock()
	defer serviceQueueMutex.Unlock()
	return serviceStopFlag
}

func CanServiceBeStopped() bool {
	serviceQueueMutex.Lock()
	defer serviceQueueMutex.Unlock()
	if serviceStopFlag && serviceQueueNum == 0 {
		return true
	}
	return false
}

func SendServiceAlive(serviceName string) int {
	var respData []byte
	var respInfo msgdef.RespInfoStruct
	reqData := []byte(`{"msg":"keepalive"}`)
	lServicePort, _ := strconv.Atoi(genutil.GetListeningPortFromServiceName(serviceName))
	if msgutil.PostReq(serviceName, "localhost", lServicePort, rtutil.GetServiceAliveTimeOut(serviceName), reqData, &respData, fmt.Sprintf("%s:%s", headerdef.App_Header_Type_MsgType, headerdef.App_Header_Value_HeartBeat)) < 0 {
		return -1
	}
	if msgutil.ParseResp(respData, &respInfo) < 0 {
		return -1
	}
	return 1
}

func SendServiceSetDebugLevel(serviceName string, debugLevel string) int {
	var respData []byte
	var respInfo msgdef.RespInfoStruct
	reqData := []byte(fmt.Sprintf(`{"debug_level":"%s"}`, debugLevel))
	lServicePort, _ := strconv.Atoi(genutil.GetListeningPortFromServiceName(serviceName))
	if msgutil.PostReq(serviceName, "localhost", lServicePort, rtutil.GetServiceAliveTimeOut(serviceName), reqData, &respData, fmt.Sprintf("%s:%s", headerdef.App_Header_Type_MsgType, headerdef.App_Header_Value_SetDebugLevel)) < 0 {
		return -1
	}
	if msgutil.ParseResp(respData, &respInfo) < 0 {
		return -1
	}
	return 1
}

func SendServiceStop(serviceName string) int {
	var respData []byte
	var respInfo msgdef.RespInfoStruct
	reqData := []byte(`{"msg":"stop-service"}`)
	lServicePort, _ := strconv.Atoi(genutil.GetListeningPortFromServiceName(serviceName))
	if msgutil.PostReq(serviceName, "localhost", lServicePort, rtutil.GetServiceAliveTimeOut(serviceName), reqData, &respData, fmt.Sprintf("%s:%s", headerdef.App_Header_Type_MsgType, headerdef.App_Header_Value_StopService)) < 0 {
		return -1
	}
	if msgutil.ParseResp(respData, &respInfo) < 0 {
		return -1
	}
	return 1
}

func sendServiceIntimationMsg(HTTPResponseWriter http.ResponseWriter, HTTPRequest *http.Request) {
	HTTPResponseWriter.WriteHeader(http.StatusOK)
	HTTPResponseWriter.Write(msgutil.SetResp(msgdef.RCapproved, []byte(msgdef.RespApprovedStr), []byte(msgdef.RespApprovedStr)))
}

func IsServiceIntimationMsg(HTTPResponseWriter http.ResponseWriter, HTTPRequest *http.Request) bool {
	serviceIntimationMsgFlag := false
	if HTTPRequest.Header.Get(headerdef.App_Header_Type_MsgType) == headerdef.App_Header_Value_HeartBeat ||
		HTTPRequest.Header.Get(headerdef.App_Header_Type_MsgType) == headerdef.App_Header_Value_SetDebugLevel ||
		HTTPRequest.Header.Get(headerdef.App_Header_Type_MsgType) == headerdef.App_Header_Value_StopService {
		serviceIntimationMsgFlag = true
	}
	if serviceIntimationMsgFlag {
		defer sendServiceIntimationMsg(HTTPResponseWriter, HTTPRequest)
		if HTTPRequest.Header.Get(headerdef.App_Header_Type_MsgType) == headerdef.App_Header_Value_SetDebugLevel {
			var debugLevel string
			var traceLevel int
			RequestBody, err := ioutil.ReadAll(HTTPRequest.Body)
			if err != nil {
				return serviceIntimationMsgFlag
			}
			if jsonutil.GetValueFromJSONObj(RequestBody, "debug_level", &debugLevel) < 0 {
				return serviceIntimationMsgFlag
			}
			if strings.EqualFold(debugLevel, debugdef.DEBUG_LEVEL_NORMAL_STR) {
				traceLevel = debugdef.DEBUG_LEVEL_NORMAL
			} else if strings.EqualFold(debugLevel, debugdef.DEBUG_LEVEL_SECURED_STR) {
				traceLevel = debugdef.DEBUG_LEVEL_SECURED
			} else if strings.EqualFold(debugLevel, debugdef.DEBUG_LEVEL_TEST_STR) {
				traceLevel = debugdef.DEBUG_LEVEL_TEST
			} else if strings.EqualFold(debugLevel, debugdef.DEBUG_LEVEL_ERROR_STR) {
				traceLevel = debugdef.DEBUG_LEVEL_ERROR
			} else {
				return serviceIntimationMsgFlag
			}
			trace.SetTraceLevel(traceLevel)
		} else if HTTPRequest.Header.Get(headerdef.App_Header_Type_MsgType) == headerdef.App_Header_Value_StopService {
			SetServiceStopFlag()
			go cleanupRoutine()
		}
	}
	return serviceIntimationMsgFlag
}

func cleanupRoutine() {
	for {
		time.Sleep(time.Second * time.Duration(5))
		if CanServiceBeStopped() {
			os.Exit(globaldef.EXIT_NORMAL)
		}
	}
}

func StartHttpServer() {
	if rtutil.IsServiceTLSFlagSet(serviceName) {
		serverCrt := fmt.Sprintf("%s/%s/%s.%s.crt", globaldef.GetAppBaseDir(), "config/tls/certificates/server", rtutil.GetCurrentNodeName(), serviceName)
		serverKey := fmt.Sprintf("%s/%s/%s.%s.key", globaldef.GetAppBaseDir(), "config/tls/certificates/server", rtutil.GetCurrentNodeName(), serviceName)
		cert, err := tls.LoadX509KeyPair(serverCrt, serverKey)
		if err != nil {
			os.Exit(0)
		}
		config := &tls.Config{
			Certificates:             []tls.Certificate{cert},
			MinVersion:               tls.VersionTLS13,
			CipherSuites:             []uint16{tls.TLS_AES_128_GCM_SHA256, tls.TLS_AES_256_GCM_SHA384},
			PreferServerCipherSuites: true,
		}
		server := &http.Server{
			Addr:      fmt.Sprintf(":%s", GetListenPort()),
			TLSConfig: config,
		}
		err = server.ListenAndServeTLS("", "")
		if err != nil {
			os.Exit(0)
		}
	} else {
		err := http.ListenAndServe(fmt.Sprintf(":%s", GetListenPort()), nil)
		if err != nil {
			os.Exit(0)
		}
	}
}
