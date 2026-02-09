package reqbrokerdef

import "encoding/json"

type MsgBrokerReqInfo struct {
	MicroserviceName  string          `json:"microservice_name"`
	BrokerRequestData json.RawMessage `json:"broker_request_data"`
}
