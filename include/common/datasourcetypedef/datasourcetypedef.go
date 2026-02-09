package datasourcetypedef

const (
	DataSourceTypeReqbrokerDataMap = "reqbrokermap"
	DataSourceTypeKey              = "key"
	DataSourceTypeRawValue         = "raw_value"
	DataSourceTypeComputeStrExpr   = "compute_str_expr"
	DataSourceTypePasswd           = "password"
	DataSourceTypeGenPasswd        = "gen-password"
)

const (
	DataSourceDefGetDate                 = "get_date"
	DataSourceDefGetTime                 = "get_time"
	DataSourceDefGetRecordNum            = "get_record_num"
	DataSourceDefGenTxnBatchNum          = "gen_txn_batch_num"
	DataSourceDefGenTxnRecordNum         = "gen_txn_record_num"
	DataSourceDefGenRRN                  = "gen_rrn"
	DataSourceDefGenStan                 = "gen_stan"
	DataSourceDefGetDBRecordID           = "get_db_record_id"
	DataSourceDefGetNA                   = "get_na"
	DataSourceDefWaitingForAuth          = "waiting_for_auth"
	DataSourceDefActionInsert            = "action_insert"
	DataSourceDefGetIntegrityRowCheckSum = "get_row_integrity_checksum"
	DataSourceDefGetTimeStamp            = "get_time_stamp"
	DataSourceDefGetSchemaName           = "get_schema_name"
)
