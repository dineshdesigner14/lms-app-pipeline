package securitydef

import "encoding/json"

type GenCVVReqMsgStruct struct {
	CardNum           string
	ExpDate           string
	ServiceCode       string
	CVK1              string
	CVK2              string
	AdditionalReqData json.RawMessage
}

type GenCVVRespMsgStruct struct {
	CVV                string
	AdditionalRespData json.RawMessage
}
