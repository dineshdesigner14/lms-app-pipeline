package headerdef

const (
	App_Header_Type_MsgType    = "msgtype"
	App_Header_Type_ReqType    = "reqtype"
	App_Header_Type_ReqSubType = "reqsubtype"
)

const (
	App_Header_Value_HeartBeat     = "heartbeat"
	App_Header_Value_SetDebugLevel = "setdebuglevel"
	App_Header_Value_StopService   = "stopservice"
	App_Header_Value_Request       = "request"
)

const (
	App_Header_XAPI_Lang         = "x-api-lang"
	App_Header_XAPI_Channel      = "x-api-channel"
	App_Header_XAPI_Version      = "x-api-version"
	App_Header_XAPI_AuthToken    = "Authorization"
	App_Header_XAPI_RefreshToken = "x-refresh-token"
)

const (
	App_Header_Broker_MsgSrcType   = "msgsrctype"
	App_Header_Broker_MsgSrc       = "msgsrc"
	App_Header_Broker_Channel      = "channel"
	App_Header_Broker_Lang         = "lang"
	App_Header_Broker_version      = "version"
	App_Header_Broker_AuthToken    = "auth_token"
	App_Header_Broker_RefreshToken = "refresh_token"
	App_Header_Broker_MsgDigest    = "msg_digest"
)
