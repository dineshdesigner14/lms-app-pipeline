package seqdef

type SeqEngMsgInfo struct {
	Command string
}

const (
	GenSequenceCommand = "GenSequence"
	GenReqInfoCommand  = "GenReqInfo"
)

const (
	SeqEngReqTypeJSONObj = "ReqType"
	SeqEngDataJSONObj    = "SeqData"
	SeqEngSeqNumObj      = "SequenceNum"
)
