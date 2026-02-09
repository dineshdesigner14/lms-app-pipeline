package seqdef

type SeqEngMsgStruct struct {
	EntityType string `json:"entity_type,omitempty"`
	EntityID   string `json:"entity_id,omitempty"`
	RelType    string `json:"rel_type,omitempty"`
	RelSeq     string `json:"rel_seq,omitempty"`
	SeqLen     int    `json:"seq_len,omitempty"`
}

type SeqEngReqMsgStruct struct {
	RequestNumFlag string `json:"request_num_flag,omitempty"`
	RecordNumFlag  string `json:"record_num_flag,omitempty"`
	RRNFlag        string `json:"rrn_flag,omitempty"`
	StanFlag       string `json:"stan_flag,omitempty"`
}

type SeqReqInfoRespStruct struct {
	RequestNum string
	RecordNum  string
	RRN        string
	Stan       string
}
