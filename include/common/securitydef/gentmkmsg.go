package securitydef

import "encoding/json"

type GenTMKReqMsgStruct struct {
	Component1        string
	Component2        string
	AdditionalReqData json.RawMessage
}

type GenTMKRespMsgStruct struct {
	ETMK               string
	TMKKCV             string
	Component1KCV      string
	Component2KCV      string
	AdditionalRespData json.RawMessage
}
