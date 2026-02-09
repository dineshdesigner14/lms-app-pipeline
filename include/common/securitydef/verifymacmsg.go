package securitydef

import "encoding/json"

type VerifyMACReqMsgStruct struct {
	MACKey            string
	MACData           string
	MAC               string
	AdditionalReqData json.RawMessage
}

type VerifyMACRespMsgStruct struct {
	AdditionalRespData json.RawMessage
}
