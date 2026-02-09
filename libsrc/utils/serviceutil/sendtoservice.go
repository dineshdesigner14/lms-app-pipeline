package serviceutil

import (
	"lmsapieng/include/common/dbdef"
	"lmsapieng/include/dbtabdef"
	"lmsapieng/libsrc/dbtab/dbtab_serviceaddrinfo"
	"lmsapieng/libsrc/utils/msgutil"
	"lmsapieng/libsrc/utils/schemainfo"
	"lmsapieng/libsrc/utils/trace"
)

func SendToService(reqBrokerDataMap map[string]interface{}, serviceModule string, reqBuffer []byte, respBuffer *[]byte, serviceTimeout int) int {
	var resultDBContext dbdef.DBContextDef
	var rejectDesc string
	var dbErr, dbRejectReason string
	if schemainfo.GetActiveDBContext(reqBrokerDataMap, &resultDBContext, &rejectDesc, dbdef.DBModuleRT) < 0 {
		trace.Lg("GetActiveDBContext() failed for Module[%s]", dbdef.DBModuleRT)
		trace.Lg("SendToService() failed for serviceModule[%s]", serviceModule)
		return -1
	}
	serviceAddrInfoPtr := make([]dbtabdef.ServiceAddrInfoTable, 0)
	if dbtab_serviceaddrinfo.Load_1_ServiceAddrInfoTable(serviceModule, &serviceAddrInfoPtr, resultDBContext, &dbErr, &dbRejectReason) < 0 {
		trace.Lg("Load_1_ServiceAddrInfoTable() failed for serviceModule[%s] with dberr[%s]", serviceModule, dbErr)
		trace.Lg("SendToService() failed for serviceModule[%s]", serviceModule)
		return -1
	}
	msgSuccess := false
	for i := 0; i < len(serviceAddrInfoPtr); i++ {
		if msgutil.PostReq(serviceModule, serviceAddrInfoPtr[i].ServiceIpAddr, serviceAddrInfoPtr[i].ServicePort, serviceAddrInfoPtr[i].ServiceTimeout, reqBuffer, respBuffer) < 0 {
			continue
		}
		msgSuccess = true
	}
	if !msgSuccess {
		trace.Lg("PostReq() failed for serviceModule[%s]", serviceModule)
		trace.Lg("SendToService() failed for serviceModule[%s]", serviceModule)
		return -1
	}
	return 1
}
