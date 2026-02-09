package diagdef

type DiagInfoStruct struct {
	ServiceName       string `json:"service_name,omitempty"`
	ServiceID         string `json:"service_id,omitempty"`
	StartDate         string `json:"start_date,omitempty"`
	StartTime         string `json:"start_time,omitempty"`
	MsgType           string `json:"msg_type,omitempty"`
	LastKeepaliveDate string `json:"last_keepalive_date,omitempty"`
	LastKeepaliveTime string `json:"last_keepalive_time,omitempty"`
	LastRequestDate   string `json:"last_request_date,omitempty"`
	LastRequestTime   string `json:"last_request_time,omitempty"`
	MsgSize           string `json:"msg_size,omitempty"`
}
