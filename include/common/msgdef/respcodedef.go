package msgdef

const (
	RespApprovedStr = "Approved"
)

const (
	RCapproved        = "0"
	RCsystemerror     = "96"
	RCunabletoprocess = "5"
)

var CoreDefineRespCodeListMap = map[string]string{
	RCapproved:        "Approved",
	RCunabletoprocess: "Unabletoprocess",
	RCsystemerror:     "Systemerror",
}
