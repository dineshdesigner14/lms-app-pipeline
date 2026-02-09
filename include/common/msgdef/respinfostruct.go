package msgdef

import "encoding/json"

type RespInfoStruct struct {
	RespInfo struct {
		RejectCode          string          `json:"reject_code"`
		RejectFunction      string          `json:"reject_function"`
		RejectLongDesc      string          `json:"reject_long_desc"`
		RejectModuleType    string          `json:"reject_module_type"`
		RejectModule        string          `json:"reject_module"`
		RejectShortDesc     string          `json:"reject_short_desc"`
		RejectSrc           string          `json:"reject_src"`
		RespCode            string          `json:"resp_code"`
		RespData            json.RawMessage `json:"resp_data"`
		RespDesc            string          `json:"resp_desc"`
		RespStatus          int             `json:"resp_status"`
		AppErrDesc          string          `json:"app_err_desc"`
		AppCorrectiveAction string          `json:"app_corrective_action"`
	} `json:"resp_info"`
}
