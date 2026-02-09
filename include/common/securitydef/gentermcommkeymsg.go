package securitydef

import "encoding/json"

type GenTermCommKeyReqMsgStruct struct {
	TMK               string
	AdditionalReqData json.RawMessage
}

type GenTermCommKeyRespMsgStruct struct {
	TPKUTMK            string
	TPKULMK            string
	TPKKCV             string
	AdditionalRespData json.RawMessage
}
