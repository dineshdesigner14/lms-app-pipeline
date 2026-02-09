package reqbrokerutil

import (
	"encoding/json"
	"lmsapieng/include/common/headerdef"
	cbsapiengdef "lmsapieng/include/common/lmsapiengdef"
	"lmsapieng/include/common/moduledef"
	"lmsapieng/include/common/reqbrokerdef"
	"lmsapieng/include/common/reqkeydef"
	"lmsapieng/libsrc/utils/serviceutil"
)

func TranslateToBrokerMsg(reqType string, reqBuffer map[string]interface{}) []byte {
	broker_req_map := make(map[string]interface{})
	header_data_map := make(map[string]interface{})
	header_data_map[headerdef.App_Header_Broker_MsgSrcType] = moduledef.LMSApiEngModule
	header_data_map[headerdef.App_Header_Broker_MsgSrc] = serviceutil.GetServiceName()
	header_data_map[headerdef.App_Header_Broker_version] = cbsapiengdef.LMSApiEngVersion
	broker_req_map[reqbrokerdef.ReqBrokerReqHeaderJSONObj] = header_data_map
	request_key_map := make(map[string]interface{})
	request_key_map[reqkeydef.App_ReqKey_Broker_ReqType] = reqType
	broker_req_map[reqbrokerdef.ReqBrokerReqKeyJSONObj] = request_key_map
	broker_req_map[reqbrokerdef.ReqBrokerReqDataJSONObj] = reqBuffer

	brokerReqBuffer, _ := json.Marshal(&broker_req_map)
	return brokerReqBuffer
}
