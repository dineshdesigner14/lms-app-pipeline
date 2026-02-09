package emaildef

type EMailDataInfo struct {
	GatewayName   string   `json:"gateway_name,omitempty"`
	RequestType   string   `json:"request_type,omitempty"`
	RequestNum    string   `json:"request_num,omitempty"`
	MsgID         string   `json:"msg_id,omitempty"`
	FromAddress   string   `json:"from_address,omitempty"`
	ToAddressList []string `json:"to_address_list,omitempty"`
	CCAddressList []string `json:"cc_address_list,omitempty"`
	Subject       string   `json:"subject,omitempty"`
	Body          string   `json:"body,omitempty"`
}

type EMailEngMsgInfo struct {
	EmailData EMailDataInfo `json:"mail_data,omitempty"`
}
