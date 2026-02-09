package diagdef

import "encoding/json"

type DiagEngReqMsgStruct struct {
	Command     string          `json:"command,omitempty"`
	DiagReqData json.RawMessage `json:"diag_req_data,omitempty"`
}
