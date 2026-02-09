package dbtabdef

// LMSSeqDataTable defines the structure of LMSSeqDataTable

type LMSSeqDataTable struct {
	InstID       string
	EntityType   string
	SeqPrefix    string
	SeqLen       int
	OverflowFlag string
	SeqNum       string
}
