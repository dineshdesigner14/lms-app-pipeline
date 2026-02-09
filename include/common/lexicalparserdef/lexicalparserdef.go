package lexicalparserdef

import "encoding/xml"

type LPDBResultInfoStruct struct {
	Text       string `xml:",chardata"`
	ResultCode string `xml:"result_code,attr"`
	Result     string `xml:"result,attr"`
}

type LPSchemaInfoStruct struct {
	Text                string `xml:",chardata"`
	ModuleName          string `xml:"module_name,attr"`
	SubModuleName       string `xml:"sub_module_name,attr"`
	InstID              string `xml:"inst_id,attr"`
	InstSubID           string `xml:"inst_sub_id,attr"`
	BinID               string `xml:"bin_id,attr"`
	BinSubID            string `xml:"bin_sub_id,attr"`
	AppErrDesc          string `xml:"app_err_desc,attr"`
	AppCorrectiveAction string `xml:"app_corrective_action,attr"`
}

type LPValidateFldStruct struct {
	Text                string `xml:",chardata"`
	FldPresent          string `xml:"fld_present,attr"`
	FldDefaultSource    string `xml:"fld_default_source,attr"`
	FldDefaultValue     string `xml:"fld_default_value,attr"`
	FldName             string `xml:"fld_name,attr"`
	FldDataType         string `xml:"fld_data_type,attr"`
	FldMinLen           string `xml:"fld_min_len,attr"`
	FldMaxLen           string `xml:"fld_max_len,attr"`
	FldType             string `xml:"fld_type,attr"`
	FldAllowedChars     string `xml:"fld_allowed_chars,attr"`
	FldAllowedLen       string `xml:"fld_allowed_len,attr"`
	FldPattern          string `xml:"fld_pattern,attr"`
	AppErrDesc          string `xml:"app_err_desc,attr"`
	AppCorrectiveAction string `xml:"app_corrective_action,attr"`
}

type LPValidateInfoStruct struct {
	Text                string                `xml:",chardata"`
	ValidateFld         []LPValidateFldStruct `xml:"validate_fld"`
	AppErrDesc          string                `xml:"app_err_desc,attr"`
	AppCorrectiveAction string                `xml:"app_corrective_action,attr"`
}

type LPFilterDataInfoStruct struct {
	Text                string `xml:",chardata"`
	FilterName          string `xml:"filter_name,attr"`
	FilterSourceType    string `xml:"filter_source_type,attr"`
	FilterID            string `xml:"filter_id,attr"`
	FilterOperator      string `xml:"filter_operator,attr"`
	FilterDataType      string `xml:"filter_data_type,attr"`
	FilterDBDataType    string `xml:"filter_db_data_type,attr"`
	FilterCondition     string `xml:"filter_condition,attr"`
	FilterApply         string `xml:"apply_filter,attr"`
	AppErrDesc          string `xml:"app_err_desc,attr"`
	AppCorrectiveAction string `xml:"app_corrective_action,attr"`
}

type LPCreateDBSequenceInfoStruct struct {
	Text                string             `xml:",chardata"`
	ObjectName          string             `xml:"object_name,attr"`
	TableName           string             `xml:"table_name"`
	AppErrDesc          string             `xml:"app_err_desc,attr"`
	AppCorrectiveAction string             `xml:"app_corrective_action,attr"`
	SchemaInfo          LPSchemaInfoStruct `xml:"schema_info"`
	SequenceDbData      struct {
		Text                string `xml:",chardata"`
		Key                 string `xml:"key,attr"`
		SequenceLen         string `xml:"sequence_len,attr"`
		DataType            string `xml:"data_type,attr"`
		AppErrDesc          string `xml:"app_err_desc,attr"`
		AppCorrectiveAction string `xml:"app_corrective_action,attr"`
	} `xml:"sequence_db_data"`
	FilterInfo struct {
		Text       string                   `xml:",chardata"`
		FilterData []LPFilterDataInfoStruct `xml:"filter_data"`
	} `xml:"filter_info"`
}

type LPDBSingleReadInfoStruct struct {
	Text                string             `xml:",chardata"`
	TableName           string             `xml:"table_name"`
	AppErrDesc          string             `xml:"app_err_desc,attr"`
	AppCorrectiveAction string             `xml:"app_corrective_action,attr"`
	SchemaInfo          LPSchemaInfoStruct `xml:"schema_info"`
	FilterInfo          struct {
		Text       string                   `xml:",chardata"`
		FilterData []LPFilterDataInfoStruct `xml:"filter_data"`
	} `xml:"filter_info"`
	ResultInfo struct {
		Text                string `xml:",chardata"`
		ResultSuccess       string `xml:"result_success"`
		AppErrDesc          string `xml:"app_err_desc,attr"`
		AppCorrectiveAction string `xml:"app_corrective_action,attr"`
	} `xml:"result_info"`
	StoreObjName string `xml:"store_obj_name"`
}

type LPDBMultiReadInfoStruct struct {
	Text                string             `xml:",chardata"`
	TableName           string             `xml:"table_name"`
	AppErrDesc          string             `xml:"app_err_desc,attr"`
	AppCorrectiveAction string             `xml:"app_corrective_action,attr"`
	SchemaInfo          LPSchemaInfoStruct `xml:"schema_info"`
	FilterInfo          struct {
		Text       string                   `xml:",chardata"`
		FilterData []LPFilterDataInfoStruct `xml:"filter_data"`
	} `xml:"filter_info"`
	StoreObjName string `xml:"store_obj_name"`

	DBResultInfo LPDBResultInfoStruct `xml:"dbresult_info"`
	OrderBy      string               `xml:"order_by,attr"`
	Limit        string               `xml:"limit,attr"`
	SortType     string               `xml:"sort_type,attr"`
	Offset       string               `xml:"offset,attr"`
	SearchStr    string               `xml:"search_str,attr"`
}

type LPRespInfoStruct struct {
	Text                string `xml:",chardata"`
	AppErrDesc          string `xml:"app_err_desc,attr"`
	AppCorrectiveAction string `xml:"app_corrective_action,attr"`
	RespObj             []struct {
		Text       string `xml:",chardata"`
		ObjName    string `xml:"obj_name,attr"`
		ObjType    string `xml:"obj_type,attr"`
		ObjReq     string `xml:"obj_req,attr"`
		DataSource string `xml:"data_source,attr"`
	} `xml:"resp_obj"`
}

type LPCreateObjectInfoStruct struct {
	Text                string `xml:",chardata"`
	ObjectName          string `xml:"object_name,attr"`
	AppErrDesc          string `xml:"app_err_desc,attr"`
	AppCorrectiveAction string `xml:"app_corrective_action,attr"`
	ObjectData          []struct {
		Text                string `xml:",chardata"`
		Key                 string `xml:"key,attr"`
		DataSourceType      string `xml:"data_source_type,attr"`
		DataSource          string `xml:"data_source,attr"`
		DataType            string `xml:"data_type,attr"`
		Condition           string `xml:"condition,attr"`
		AppErrDesc          string `xml:"app_err_desc,attr"`
		AppCorrectiveAction string `xml:"app_corrective_action,attr"`
	} `xml:"object_data"`
}

type LPCopyObjectInfoStruct struct {
	Text                string `xml:",chardata"`
	ObjectName          string `xml:"object_name,attr"`
	AppErrDesc          string `xml:"app_err_desc,attr"`
	AppCorrectiveAction string `xml:"app_corrective_action,attr"`
	ObjectData          []struct {
		Text                string `xml:",chardata"`
		Key                 string `xml:"key,attr"`
		DataSourceType      string `xml:"data_source_type,attr"`
		DataSource          string `xml:"data_source,attr"`
		DataType            string `xml:"data_type,attr"`
		Condition           string `xml:"condition,attr"`
		AppErrDesc          string `xml:"app_err_desc,attr"`
		AppCorrectiveAction string `xml:"app_corrective_action,attr"`
	} `xml:"object_data"`
}

type LPTransformObjectInfoStruct struct {
	Text                string `xml:",chardata"`
	ObjectName          string `xml:"object_name,attr"`
	AppErrDesc          string `xml:"app_err_desc,attr"`
	AppCorrectiveAction string `xml:"app_corrective_action,attr"`
	ObjectData          []struct {
		Text                string `xml:",chardata"`
		Algo                string `xml:"algo,attr"`
		StartLen            string `xml:"start_len,attr"`
		EndLen              string `xml:"end_len,attr"`
		Key                 string `xml:"key,attr"`
		DataSourceType      string `xml:"data_source_type,attr"`
		DataSource          string `xml:"data_source,attr"`
		DataType            string `xml:"data_type,attr"`
		Condition           string `xml:"condition,attr"`
		AppErrDesc          string `xml:"app_err_desc,attr"`
		AppCorrectiveAction string `xml:"app_corrective_action,attr"`
	} `xml:"object_data"`
}

type LPDBInsertInfoStruct struct {
	Text                string             `xml:",chardata"`
	TableName           string             `xml:"table_name"`
	SchemaInfo          LPSchemaInfoStruct `xml:"schema_info"`
	AppErrDesc          string             `xml:"app_err_desc,attr"`
	AppCorrectiveAction string             `xml:"app_corrective_action,attr"`
	InsertDataList      struct {
		Text       string `xml:",chardata"`
		InsertData []struct {
			Text                string `xml:",chardata"`
			InsertDataName      string `xml:"insert_data_name,attr"`
			InsertDBDataType    string `xml:"insert_db_data_type,attr"`
			InsertDataSource    string `xml:"insert_data_source,attr"`
			InsertDataType      string `xml:"insert_data_type,attr"`
			Condition           string `xml:"condition,attr"`
			AppErrDesc          string `xml:"app_err_desc,attr"`
			AppCorrectiveAction string `xml:"app_corrective_action,attr"`
		} `xml:"insert_data"`
	} `xml:"insert_data_list"`
}

type LPDBUpdateInfoStruct struct {
	Text                string             `xml:",chardata"`
	TableName           string             `xml:"table_name"`
	SchemaInfo          LPSchemaInfoStruct `xml:"schema_info"`
	AppErrDesc          string             `xml:"app_err_desc,attr"`
	AppCorrectiveAction string             `xml:"app_corrective_action,attr"`
	UpdateDataList      struct {
		Text       string `xml:",chardata"`
		UpdateData []struct {
			Text                string `xml:",chardata"`
			UpdateDataName      string `xml:"update_data_name,attr"`
			UpdateDBDataType    string `xml:"update_db_data_type,attr"`
			UpdateDataSource    string `xml:"update_data_source,attr"`
			UpdateDataType      string `xml:"update_data_type,attr"`
			Condition           string `xml:"condition,attr"`
			AppErrDesc          string `xml:"app_err_desc,attr"`
			AppCorrectiveAction string `xml:"app_corrective_action,attr"`
		} `xml:"update_data"`
	} `xml:"update_data_list"`
	UpdateFilterList struct {
		Text         string `xml:",chardata"`
		UpdateFilter []struct {
			Text                   string `xml:",chardata"`
			UpdateFilterName       string `xml:"update_filter_name,attr"`
			UpdateFilterDBDataType string `xml:"update_filter_db_data_type,attr"`
			UpdateFilterDataSource string `xml:"update_filter_data_source,attr"`
			UpdateFilterDataType   string `xml:"update_filter_data_type,attr"`
			UpdateFilterCondition  string `xml:"update_filter_condition,attr"`
			Condition              string `xml:"condition,attr"`
			AppErrDesc             string `xml:"app_err_desc,attr"`
			AppCorrectiveAction    string `xml:"app_corrective_action,attr"`
		} `xml:"update_filter"`
	} `xml:"update_filter_list"`
}

type LPDBDeleteInfoStruct struct {
	Text                string             `xml:",chardata"`
	TableName           string             `xml:"table_name"`
	SchemaInfo          LPSchemaInfoStruct `xml:"schema_info"`
	AppErrDesc          string             `xml:"app_err_desc,attr"`
	AppCorrectiveAction string             `xml:"app_corrective_action,attr"`
	DeleteFilterList    struct {
		Text         string `xml:",chardata"`
		DeleteFilter []struct {
			Text                   string `xml:",chardata"`
			DeleteFilterName       string `xml:"delete_filter_name,attr"`
			DeleteFilterDBDataType string `xml:"delete_filter_db_data_type,attr"`
			DeleteFilterDataSource string `xml:"delete_filter_data_source,attr"`
			DeleteFilterDataType   string `xml:"delete_filter_data_type,attr"`
			DeleteFilterCondition  string `xml:"delete_filter_condition,attr"`
			Condition              string `xml:"condition,attr"`
			AppErrDesc             string `xml:"app_err_desc,attr"`
			AppCorrectiveAction    string `xml:"app_corrective_action,attr"`
		} `xml:"delete_filter"`
	} `xml:"delete_filter_list"`
}

type LPValidateConditionInfoStruct struct {
	Text                        string   `xml:",chardata"`
	ValidateConditionExpression []string `xml:"validate_condition_expression"`
	AppErrDesc                  string   `xml:"app_err_desc,attr"`
	AppCorrectiveAction         string   `xml:"app_corrective_action,attr"`
}

type LPGenPersoFileInfoStruct struct {
	Text                string `xml:",chardata"`
	PersoObject         string `xml:"perso_object,attr"`
	PersoObjectArray    string `xml:"perso_object_array,attr"`
	AppErrDesc          string `xml:"app_err_desc,attr"`
	AppCorrectiveAction string `xml:"app_corrective_action,attr"`
}

type LPEntitySeqNumInfoStruct struct {
	Text                string             `xml:",chardata"`
	ObjectName          string             `xml:"object_name,attr"`
	SchemaInfo          LPSchemaInfoStruct `xml:"schema_info"`
	AppErrDesc          string             `xml:"app_err_desc,attr"`
	AppCorrectiveAction string             `xml:"app_corrective_action,attr"`
}

type LPGenTrackDataInfoStruct struct {
	Text                string `xml:",chardata"`
	ObjectName          string `xml:"object_name,attr"`
	AppErrDesc          string `xml:"app_err_desc,attr"`
	AppCorrectiveAction string `xml:"app_corrective_action,attr"`
}

type LPMapArrayInfoStruct struct {
	Text                string `xml:",chardata"`
	ObjectName          string `xml:"object_name,attr"`
	ArraySize           string `xml:"array_size,attr"`
	AppErrDesc          string `xml:"app_err_desc,attr"`
	AppCorrectiveAction string `xml:"app_corrective_action,attr"`
}

type LPGenCVVInfoStruct struct {
	Text                string `xml:",chardata"`
	ObjectName          string `xml:"object_name,attr"`
	CardNum             string `xml:"card_num,attr"`
	ExpDate             string `xml:"exp_date,attr"`
	ServiceCode         string `xml:"service_code,attr"`
	CVK1                string `xml:"cvk1,attr"`
	CVK2                string `xml:"cvk2,attr"`
	AppErrDesc          string `xml:"app_err_desc,attr"`
	AppCorrectiveAction string `xml:"app_corrective_action,attr"`
}

type LPGenRandomKeyInfoStruct struct {
	Text                string `xml:",chardata"`
	ObjectName          string `xml:"object_name,attr"`
	KeyType             string `xml:"key_type,attr"`
	Key1                string `xml:"key_1,attr"`
	Comp1               string `xml:"comp_1,attr"`
	Comp2               string `xml:"comp_2,attr"`
	AppErrDesc          string `xml:"app_err_desc,attr"`
	AppCorrectiveAction string `xml:"app_corrective_action,attr"`
}

type LPMathFunctionInfoStruct struct {
	Text                string `xml:",chardata"`
	Algo                string `xml:"algo,attr"`
	SrcObject           string `xml:"src_object,attr"`
	SrcKey              string `xml:"src_key,attr"`
	DestObject          string `xml:"dest_object,attr"`
	AppErrDesc          string `xml:"app_err_desc,attr"`
	AppCorrectiveAction string `xml:"app_corrective_action,attr"`
}

type LPRawQueryInfoStruct struct {
	Text                   string               `xml:",chardata"`
	OracleQueryStr         string               `xml:"oracle_query_str"`
	PostgresOracleQueryStr string               `xml:"postgres_query_str"`
	SchemaInfo             LPSchemaInfoStruct   `xml:"schema_info"`
	StoreObjName           string               `xml:"store_obj_name"`
	DBResultInfo           LPDBResultInfoStruct `xml:"dbresult_info"`
	AppErrDesc             string               `xml:"app_err_desc,attr"`
	AppCorrectiveAction    string               `xml:"app_corrective_action,attr"`
}

type LPSendEmailInfoStruct struct {
	Text                string   `xml:",chardata"`
	AppErrDesc          string   `xml:"app_err_desc,attr"`
	AppCorrectiveAction string   `xml:"app_corrective_action,attr"`
	GatewayName         string   `xml:"gateway_name"`
	FromAddress         string   `xml:"from_address"`
	ToAddress           []string `xml:"to_address"`
	CCAddress           []string `xml:"cc_address"`
	Subject             string   `xml:"subject"`
	Body                string   `xml:"body"`
}

type LPSendSMSInfoStruct struct {
	Text                string `xml:",chardata"`
	AppErrDesc          string `xml:"app_err_desc,attr"`
	AppCorrectiveAction string `xml:"app_corrective_action,attr"`
}
type LPGenTokenInfoStruct struct {
	Text                string `xml:",chardata"`
	AppErrDesc          string `xml:"app_err_desc,attr"`
	AppCorrectiveAction string `xml:"app_corrective_action,attr"`
	UserID              string `xml:"user_id,attr"`
	Email               string `xml:"email,attr"`
	Role                string `xml:"role,attr"`
	TokenIssuer         string `xml:"token_issuer,attr"`
	TokenSecret         string `xml:"token_secret,attr"`
	TokenExpUnit        string `xml:"token_exp_unit,attr"`
	TokenExpiry         string `xml:"token_expiry,attr"`
	TokenObject         string `xml:"token_object,attr"`
}

type LPVerifyTokenInfoStruct struct {
	Text                string `xml:",chardata"`
	AppErrDesc          string `xml:"app_err_desc,attr"`
	AppCorrectiveAction string `xml:"app_corrective_action,attr"`
	TokenValue          string `xml:"token_value,attr"`
	UserID              string `xml:"user_id,attr"`
	Email               string `xml:"email,attr"`
	Role                string `xml:"role,attr"`
	TokenSecret         string `xml:"token_secret,attr"`
	TokenObject         string `xml:"token_object,attr"`
}

type LPDecodeTokenInfoStruct struct {
	Text                string `xml:",chardata"`
	AppErrDesc          string `xml:"app_err_desc,attr"`
	AppCorrectiveAction string `xml:"app_corrective_action,attr"`
	TokenValue          string `xml:"token_value,attr"`
	TokenObject         string `xml:"token_object,attr"`
}

type LPErrDescStruct struct {
	Text                string `xml:",chardata"`
	ErrCode             string `xml:"err_code,attr"`
	AppErrDesc          string `xml:"app_err_desc,attr"`
	AppCorrectiveAction string `xml:"app_corrective_action,attr"`
}

type LPErrInfoStruct struct {
	Text    string            `xml:",chardata"`
	ErrDesc []LPErrDescStruct `xml:"err_desc"`
}

type LPSendToServiceInfoStruct struct {
	Text          string `xml:",chardata"`
	ServiceModule string `xml:"service_module,attr"`
	ReqObj        string `xml:"req_obj,attr"`
	RespObj       string `xml:"resp_obj,attr"`
	TimeOut       int    `xml:"time_out,attr"`
}

type LPCreateEmptyListStruct struct {
	Text                string `xml:",chardata"`
	ObjectName          string `xml:"object_name,attr"`
	AppErrDesc          string `xml:"app_err_desc,attr"`
	AppCorrectiveAction string `xml:"app_corrective_action,attr"`
}

type LPSendToBrokerInfoStruct struct {
	Text                string `xml:",chardata"`
	AppErrDesc          string `xml:"app_err_desc,attr"`
	AppCorrectiveAction string `xml:"app_corrective_action,attr"`
	ReqType             string `xml:"req_type,attr"`
	ReqObj              string `xml:"req_obj,attr"`
	RespObj             string `xml:"resp_obj,attr"`
}

type LPArithOperationInfo struct {
	Text                string `xml:",chardata"`
	Operation           string `xml:"operation,attr"`
	LeftOperand         string `xml:"left_operand,attr"`
	RightOperand        string `xml:"right_operand,attr"`
	DestObject          string `xml:"dest_object,attr"`
	AppErrDesc          string `xml:"app_err_desc,attr"`
	AppCorrectiveAction string `xml:"app_corrective_action,attr"`
}

type ExecFunctionStruct struct {
	Text                  string                        `xml:",chardata"`
	FunctionName          string                        `xml:"function_name,attr"`
	FunctionType          string                        `xml:"function_type,attr"`
	FunctionCondition     string                        `xml:"function_condition,attr"`
	IndexName             string                        `xml:"index_name,attr"`
	StartIndex            string                        `xml:"start_index,attr"`
	EndIndex              string                        `xml:"end_index,attr"`
	ErrInfo               LPErrInfoStruct               `xml:"err_info"`
	ValidateInfo          LPValidateInfoStruct          `xml:"validate_info"`
	DBSingleReadInfo      LPDBSingleReadInfoStruct      `xml:"db_single_read_info"`
	DBMultiReadInfo       LPDBMultiReadInfoStruct       `xml:"db_multi_read_info"`
	DBInsertInfo          LPDBInsertInfoStruct          `xml:"db_insert_info"`
	DBUpdateInfo          LPDBUpdateInfoStruct          `xml:"db_update_info"`
	DBDeleteInfo          LPDBDeleteInfoStruct          `xml:"db_delete_info"`
	ResponseArray         LPRespInfoStruct              `xml:"response_info"`
	CreateObjectInfo      LPCreateObjectInfoStruct      `xml:"create_object_info"`
	CopyObjectInfo        LPCopyObjectInfoStruct        `xml:"copy_object_info"`
	TransformObjectInfo   LPTransformObjectInfoStruct   `xml:"transform_object_info"`
	ValidateConditionInfo LPValidateConditionInfoStruct `xml:"validate_condition_info"`
	CreateDBSequenceInfo  LPCreateDBSequenceInfoStruct  `xml:"create_db_sequence_info"`
	EntitySeqNumInfo      LPEntitySeqNumInfoStruct      `xml:"entity_seq_num_info"`
	GenCVVInfo            LPGenCVVInfoStruct            `xml:"gen_cvv_info"`
	GenTrackDataInfo      LPGenTrackDataInfoStruct      `xml:"gen_track_data_info"`
	MapArrayInfo          LPMapArrayInfoStruct          `xml:"map_array_info"`
	GenPersoFileInfo      LPGenPersoFileInfoStruct      `xml:"gen_perso_file_info"`
	MathFunctionInfo      LPMathFunctionInfoStruct      `xml:"math_function_info"`
	GenRandomKeyInfo      LPGenRandomKeyInfoStruct      `xml:"gen_random_key_info"`
	RawQueryInfo          LPRawQueryInfoStruct          `xml:"raw_query_info"`
	SendEmailInfo         LPSendEmailInfoStruct         `xml:"send_email_info"`
	SendSMSInfo           LPSendSMSInfoStruct           `xml:"send_sms_info"`
	VerifyTokenInfo       LPVerifyTokenInfoStruct       `xml:"verify_token_info"`
	DecodeTokenInfo       LPDecodeTokenInfoStruct       `xml:"decode_token_info"`
	GenTokenInfo          LPGenTokenInfoStruct          `xml:"gen_token_info"`
	SendToService         LPSendToServiceInfoStruct     `xml:"send_service_info"`
	CreateEmptyList       LPCreateEmptyListStruct       `xml:"create_empty_list_info"`
	SendToBrokerInfo      LPSendToBrokerInfoStruct      `xml:"send_broker_info"`
	ArithFunctionInfo     LPArithOperationInfo          `xml:"arith_operation_info"`
}

type ExecGroupStruct struct {
	Text           string               `xml:",chardata"`
	GroupName      string               `xml:"group_name,attr"`
	GroupCondition string               `xml:"group_condition,attr"`
	ArrayIndex     string               `xml:"array_index,attr"`
	StartIndex     string               `xml:"start_index,attr"`
	EndIndex       string               `xml:"end_index,attr"`
	ExecFunction   []ExecFunctionStruct `xml:"exec_function"`
}

type LexicalParserStruct struct {
	XMLName      xml.Name          `xml:"sem_lexical_parser"`
	Text         string            `xml:",chardata"`
	Microservice string            `xml:"microservice,attr"`
	ExecGroup    []ExecGroupStruct `xml:"exec_group"`
}
