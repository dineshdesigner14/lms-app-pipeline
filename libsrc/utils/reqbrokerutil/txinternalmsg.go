package reqbrokerutil

import (
	"encoding/json"
)

type InternalMsgStruct struct {
	MicroserviceName  string `json:"microservice_name"`
	BrokerRequestData struct {
		ReqBrokerReqObj struct {
			HeaderData struct {
				Msgsrc     string `json:"msgsrc"`
				MsgsrcType string `json:"msgsrctype"`
				Version    string `json:"version"`
			} `json:"header_data"`
			RequestKey struct {
				ReqType string `json:"req_type"`
			} `json:"request_key"`
			RequestData json.RawMessage `json:"request_data"`
		} `json:"ReqBrokerReqObj"`
	} `json:"broker_request_data"`
}

func TranslateInternalMsg(reqType string, moduleName string, version string, reqBuffer json.RawMessage) []byte {
	var imsg InternalMsgStruct
	imsg.MicroserviceName = reqType
	imsg.BrokerRequestData.ReqBrokerReqObj.HeaderData.Msgsrc = moduleName
	imsg.BrokerRequestData.ReqBrokerReqObj.HeaderData.MsgsrcType = reqType
	imsg.BrokerRequestData.ReqBrokerReqObj.HeaderData.Version = version
	imsg.BrokerRequestData.ReqBrokerReqObj.RequestKey.ReqType = reqType
	imsg.BrokerRequestData.ReqBrokerReqObj.RequestData = reqBuffer
	brokerReqBuffer, _ := json.Marshal(&imsg)
	return brokerReqBuffer
}
