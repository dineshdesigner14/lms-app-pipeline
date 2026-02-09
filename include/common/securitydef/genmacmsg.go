package securitydef

import "encoding/json"

type GenMACReqMsgStruct struct {
	MACKey            string
	MACData           string
	AdditionalReqData json.RawMessage
}

type GenMACRespMsgStruct struct {
	MAC                string
	MACKCV             string
	AdditionalRespData json.RawMessage
}
