package carddef

const (
	Track1DataSize = 79
	Track2DataSize = 40
	Track3DataSize = 107
)

const (
	CardPersoRecordKeyBIN            = "bin"
	CardPersoRecordKeyCardSubProduct = "card_subproduct"
	CardPersoRecordKeyBatchID        = "batch_id"
	CardPersoRecordKeyCardTypeID     = "card_type_id"
	CardPersoRecordKeyCardNum        = "card_num"
	CardPersoRecordKeyCardSerialNum  = "card_serial_num"
	CardPersoRecordKeyCardSeqNum     = "card_seq_num"
	CardPersoRecordKeyExpiryDate     = "expiry_date"
	CardPersoRecordKeyCardName       = "card_name"
	CardPersoRecordKeyCoBrandedName  = "co_branded_name"
	CardPersoRecordKeyTrack1         = "track1"
	CardPersoRecordKeyTrack2         = "track2"
	CardPersoRecordKeyCVV1           = "cvv1"
	CardPersoRecordKeyCVV2           = "cvv2"
	CardPersoRecordKeyICVV           = "icvv"
	CardPersoRecordKeyServiceCode    = "service_code"
	CardPersoRecordKeyCustomerName   = "customer_name"
	CardPersoRecordKeyCardIssDate    = "card_iss_date"
	CardPersoRecordKeyServerFlag     = "server_flag"
)

type CardEmbossingFileStruct struct {
	CardTypeID    string
	CardNum       string
	CardSerialNum string
	CardSeqNum    string
	ExpiryDate    string
	CardName      string
	CoBrandedName string
	Track1        string
	Track2        string
	CVV1          string
	CVV2          string
	ICVV          string
	ServiceCode   string
	CustomerName  string
	CardIssDate   string
	ServerFlag    string
}
