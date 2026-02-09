package securitydef

import "encoding/json"

type RcvMACKeyReqMsgStruct struct {
	ZMK               string
	MACKeyUZMK        string
	AdditionalReqData json.RawMessage
}

type RcvMACKeyRespMsgStruct struct {
	MACKeyULMK         string
	MACKeyKCV          string
	AdditionalRespData json.RawMessage
}
