package rejectdef

import "lmsapieng/include/common/msgdef"

const (
	LMS_Admin_Service_Reject_GetActiveDBContextError           = "LMS_Admin_Service.1"
	LMS_Admin_Service_Reject_DecomposeReqError                 = "LMS_Admin_Service.2"
	LMS_Admin_Service_Reject_InvalidMicroService               = "LMS_Admin_Service.3"
	LMS_Admin_Service_Reject_GenSequenceErr                    = "LMS_Admin_Service.4"
	LMS_Admin_Service_Reject_XMLLexicalParseErr                = "LMS_Admin_Service.5"
	LMS_Admin_Service_Reject_FldTypeInvalidErr                 = "LMS_Admin_Service.6"
	LMS_Admin_Service_Reject_FldValidationErr                  = "LMS_Admin_Service.7"
	LMS_Admin_Service_Reject_ReadValueErr                      = "LMS_Admin_Service.8"
	LMS_Admin_Service_Reject_DBSingleReadFTTableNullErr        = "LMS_Admin_Service.9"
	LMS_Admin_Service_Reject_ReadFromDBTableErr                = "LMS_Admin_Service.10"
	LMS_Admin_Service_Reject_DBSingleReadFTFilterDataErr       = "LMS_Admin_Service.11"
	LMS_Admin_Service_Reject_DBMultiReadFTTableNullErr         = "LMS_Admin_Service.12"
	LMS_Admin_Service_Reject_LoadFromDBTableErr                = "LMS_Admin_Service.13"
	LMS_Admin_Service_Reject_DBMultiReadFTFilterDataErr        = "LMS_Admin_Service.14"
	LMS_Admin_Service_Reject_FunctionTypeInvalidErr            = "LMS_Admin_Service.15"
	LMS_Admin_Service_Reject_JsonMarshalErr                    = "LMS_Admin_Service.16"
	LMS_Admin_Service_Reject_DBInsertFTTableNullErr            = "LMS_Admin_Service.17"
	LMS_Admin_Service_Reject_DBInsertFTErr                     = "LMS_Admin_Service.18"
	LMS_Admin_Service_Reject_GetActiveDBContextWithTxnError    = "LMS_Admin_Service.19"
	LMS_Admin_Service_Reject_InsertDBTableErr                  = "LMS_Admin_Service.20"
	LMS_Admin_Service_Reject_CreateObjectFTObjNameNullErr      = "LMS_Admin_Service.21"
	LMS_Admin_Service_Reject_CreateObjectFTObjKeyNullErr       = "LMS_Admin_Service.22"
	LMS_Admin_Service_Reject_CreateObjectFTObjDataSrcNullErr   = "LMS_Admin_Service.23"
	LMS_Admin_Service_Reject_CreateObjectFTObjDataTypeNullErr  = "LMS_Admin_Service.24"
	LMS_Admin_Service_Reject_DBUpdateFTTableNullErr            = "LMS_Admin_Service.25"
	LMS_Admin_Service_Reject_DBUpdateFTErr                     = "LMS_Admin_Service.26"
	LMS_Admin_Service_Reject_CopyObjectFTObjNameNullErr        = "LMS_Admin_Service.27"
	LMS_Admin_Service_Reject_CopyObjectFTObjKeyNullErr         = "LMS_Admin_Service.28"
	LMS_Admin_Service_Reject_CopyObjectFTObjDataSrcTypeNullErr = "LMS_Admin_Service.29"
	LMS_Admin_Service_Reject_CopyObjectFTObjDataSrcNullErr     = "LMS_Admin_Service.30"
	LMS_Admin_Service_Reject_CopyObjectFTObjDataTypeNullErr    = "LMS_Admin_Service.31"
	LMS_Admin_Service_Reject_UpdateDBTableErr                  = "LMS_Admin_Service.32"
	LMS_Admin_Service_Reject_ValidateConditionFTErr            = "LMS_Admin_Service.33"
	LMS_Admin_Service_Reject_CreateObjectFTErr                 = "LMS_Admin_Service.34"
	LMS_Admin_Service_Reject_CopyObjectFTErr                   = "LMS_Admin_Service.35"
	LMS_Admin_Service_Reject_DBDeleteFTErr                     = "LMS_Admin_Service.36"
	LMS_Admin_Service_Reject_DBDeleteFTTableNullErr            = "LMS_Admin_Service.37"
	LMS_Admin_Service_Reject_DeleteDBTableErr                  = "LMS_Admin_Service.38"
	LMS_Admin_Service_Reject_MicroServiceFunctionNotFoundErr   = "LMS_Admin_Service.39"
	LMS_Admin_Service_Reject_MicroServiceFunctionErr           = "LMS_Admin_Service.40"
	LMS_Admin_Service_Reject_MicroServiceNameFormatErr         = "LMS_Admin_Service.41"
	LMS_Admin_Service_Reject_GetDBContextParamsErr             = "LMS_Admin_Service.42"
	LMS_Admin_Service_Reject_CreateDBSequenceFTTableNullErr    = "LMS_Admin_Service.43"
	LMS_Admin_Service_Reject_CreateDBSequenceFTFilterDataErr   = "LMS_Admin_Service.44"
	LMS_Admin_Service_Reject_CreateDBSequenceFTErr             = "LMS_Admin_Service.45"
	LMS_Admin_Service_Reject_GenNextSequenceErr                = "LMS_Admin_Service.46"
	LMS_Admin_Service_Reject_ExecGroupStartIndexNullErr        = "LMS_Admin_Service.47"
	LMS_Admin_Service_Reject_ExecGroupEndIndexNullErr          = "LMS_Admin_Service.48"
	LMS_Admin_Service_Reject_StartLoopFTErr                    = "LMS_Admin_Service.49"
	LMS_Admin_Service_Reject_EndLoopFTErr                      = "LMS_Admin_Service.50"
	LMS_Admin_Service_Reject_SetExecGroupIndexErr              = "LMS_Admin_Service.51"
	LMS_Admin_Service_Reject_GetExecGroupIndexErr              = "LMS_Admin_Service.52"
	LMS_Admin_Service_Reject_StoreObjectErr                    = "LMS_Admin_Service.53"
	LMS_Admin_Service_Reject_GetDBRecordIDErr                  = "LMS_Admin_Service.54"
	LMS_Admin_Service_Reject_CallMethodFTErr                   = "LMS_Admin_Service.55"
	LMS_Admin_Service_Reject_EntitySeqNumFTErr                 = "LMS_Admin_Service.56"
	LMS_Admin_Service_Reject_TransformObjectFTErr              = "LMS_Admin_Service.57"
	LMS_Admin_Service_Reject_SecurityEngInvalidReq             = "LMS_Admin_Service.58"
	LMS_Admin_Service_Reject_LocateSecurityEngAddrFailed       = "LMS_Admin_Service.59"
	LMS_Admin_Service_Reject_SendToSecurityEngTimedOut         = "LMS_Admin_Service.60"
	LMS_Admin_Service_Reject_SendToSecurityEngFailed           = "LMS_Admin_Service.61"
	LMS_Admin_Service_Reject_SecurityEngRespUnmarshalFailed    = "LMS_Admin_Service.62"
	LMS_Admin_Service_Reject_GenCVVFTErr                       = "LMS_Admin_Service.63"
	LMS_Admin_Service_Reject_GenTrackDataFTErr                 = "LMS_Admin_Service.64"
	LMS_Admin_Service_Reject_CreateMapArrayFTErr               = "LMS_Admin_Service.65"
	LMS_Admin_Service_Reject_GenReqInfoErr                     = "LMS_Admin_Service.66"
	LMS_Admin_Service_Reject_MathFunctionFTErr                 = "LMS_Admin_Service.67"
	LMS_Admin_Service_Reject_GenRandomKeyFTErr                 = "LMS_Admin_Service.68"
	LMS_Admin_Service_Reject_RawQueryFunctionFTErr             = "LMS_Admin_Service.69"
	LMS_Admin_Service_Reject_LocateEmailEngAddrFailed          = "LMS_Admin_Service.70"
	LMS_Admin_Service_Reject_SendToEmailEngTimedOut            = "LMS_Admin_Service.71"
	LMS_Admin_Service_Reject_SendToEmailEngFailed              = "LMS_Admin_Service.72"
	LMS_Admin_Service_Reject_SendEmailFTErr                    = "LMS_Admin_Service.73"
	LMS_Admin_Service_Reject_GenPersoFileFTErr                 = "LMS_Admin_Service.74"
	LMS_Admin_Service_Reject_LocateUrmgApiEngAddr              = "LMS_Admin_Service.75"
	LMS_Admin_Service_Reject_SendToUrmgApiEngTimedOut          = "LMS_Admin_Service.76"
	LMS_Admin_Service_Reject_SendToUrmgApiEngFailed            = "LMS_Admin_Service.77"
	LMS_Admin_Service_Reject_VerifyAdmPortalAccessTokenFailed  = "LMS_Admin_Service.78"
	LMS_Admin_Service_Reject_VerifyTokenInfoFTErr              = "LMS_Admin_Service.79"
	LMS_Admin_Service_Reject_DBErr                             = "LMS_Admin_Service.80"
	LMS_Admin_Service_Reject_GenTokenInfoFTErr                 = "LMS_Admin_Service.81"
	LMS_Reject_SendToServiceFTErr                              = "LMS_Admin_Service.82"
	LMS_Admin_Service_Reject_DecodeTokenInfoFTErr              = "LMS_Admin_Service.83"
	LMS_Admin_Service_Reject_SendToBrokerFTErr                 = "LMS_Admin_Service.84"
	LMS_Admin_Service_Reject_HandleSendOTPErr                  = "LMS_Admin_Service.85"
	LMS_Admin_Service_Reject_ArithFunctionFTErr                = "LMS_Admin_Service.86"
	LMS_Admin_Service_Reject_VerifyLMSAdminAccessTokenFTErr    = "LMS_Admin_Service.87"
	LMS_Reject_HandleSendEmailErr                              = "LMS_Admin_Service.88"
)

var LMS_Admin_Service_RejectTable = []RejectDefInfo{
	{LMS_Admin_Service_Reject_GetActiveDBContextError, msgdef.RCsystemerror, "System error", "GetActiveDBContextErr"},
	{LMS_Admin_Service_Reject_DecomposeReqError, msgdef.RCsystemerror, "System error", "DecomposeReqError"},
	{LMS_Admin_Service_Reject_InvalidMicroService, msgdef.RCsystemerror, "System error", "InvalidMicroService"},
	{LMS_Admin_Service_Reject_GenSequenceErr, msgdef.RCsystemerror, "System error", "GenSequenceErr"},
	{LMS_Admin_Service_Reject_XMLLexicalParseErr, msgdef.RCsystemerror, "System error", "XMLLexicalParseErr"},
	{LMS_Admin_Service_Reject_FldTypeInvalidErr, msgdef.RCsystemerror, "System error", "FldTypeInvalidErr"},
	{LMS_Admin_Service_Reject_FldValidationErr, msgdef.RCsystemerror, "System error", "FldValidationErr"},
	{LMS_Admin_Service_Reject_ReadValueErr, msgdef.RCsystemerror, "System error", "ReadValueErr"},
	{LMS_Admin_Service_Reject_DBSingleReadFTTableNullErr, msgdef.RCsystemerror, "System error", "DBSingleReadFTTableNullErr"},
	{LMS_Admin_Service_Reject_ReadFromDBTableErr, msgdef.RCsystemerror, "System error", "ReadFromDBTableErr"},
	{LMS_Admin_Service_Reject_DBSingleReadFTFilterDataErr, msgdef.RCsystemerror, "System error", "DBSingleReadFTFilterDataErr"},
	{LMS_Admin_Service_Reject_DBMultiReadFTTableNullErr, msgdef.RCsystemerror, "System error", "DBMultiReadFTTableNullErr"},
	{LMS_Admin_Service_Reject_LoadFromDBTableErr, msgdef.RCsystemerror, "System error", "LoadFromDBTableErr"},
	{LMS_Admin_Service_Reject_DBMultiReadFTFilterDataErr, msgdef.RCsystemerror, "System error", "DBMultiReadFTFilterDataErr"},
	{LMS_Admin_Service_Reject_FunctionTypeInvalidErr, msgdef.RCsystemerror, "System error", "FunctionTypeInvalidErr"},
	{LMS_Admin_Service_Reject_JsonMarshalErr, msgdef.RCsystemerror, "System error", "JsonMarshalErr"},
	{LMS_Admin_Service_Reject_DBInsertFTTableNullErr, msgdef.RCsystemerror, "System error", "DBInsertFTTableNullErr"},
	{LMS_Admin_Service_Reject_DBInsertFTErr, msgdef.RCsystemerror, "System error", "DBInsertFTErr"},
	{LMS_Admin_Service_Reject_GetActiveDBContextWithTxnError, msgdef.RCsystemerror, "System error", "GetActiveDBContextWithTxnErr"},
	{LMS_Admin_Service_Reject_InsertDBTableErr, msgdef.RCsystemerror, "System error", "InsertDBTableErr"},
	{LMS_Admin_Service_Reject_CreateObjectFTObjNameNullErr, msgdef.RCsystemerror, "System error", "CreateObjectFTObjNameNullErr"},
	{LMS_Admin_Service_Reject_CreateObjectFTObjKeyNullErr, msgdef.RCsystemerror, "System error", "CreateObjectFTObjKeyNullErr"},
	{LMS_Admin_Service_Reject_CreateObjectFTObjDataSrcNullErr, msgdef.RCsystemerror, "System error", "CreateObjectFTObjDataSrcNullErr"},
	{LMS_Admin_Service_Reject_CreateObjectFTObjDataTypeNullErr, msgdef.RCsystemerror, "System error", "CreateObjectFTObjDataTypeNullErr"},
	{LMS_Admin_Service_Reject_DBUpdateFTTableNullErr, msgdef.RCsystemerror, "System error", "DBUpdateFTTableNullErr"},
	{LMS_Admin_Service_Reject_DBUpdateFTErr, msgdef.RCsystemerror, "System error", "DBUpdateFTErr"},
	{LMS_Admin_Service_Reject_CopyObjectFTObjNameNullErr, msgdef.RCsystemerror, "System error", "CopyObjectFTObjNameNullErr"},
	{LMS_Admin_Service_Reject_CopyObjectFTObjKeyNullErr, msgdef.RCsystemerror, "System error", "CopyObjectFTObjKeyNullErr"},
	{LMS_Admin_Service_Reject_CopyObjectFTObjDataSrcTypeNullErr, msgdef.RCsystemerror, "System error", "CopyObjectFTObjDataSrcTypeNullErr"},
	{LMS_Admin_Service_Reject_CopyObjectFTObjDataSrcNullErr, msgdef.RCsystemerror, "System error", "CopyObjectFTObjDataSrcNullErr"},
	{LMS_Admin_Service_Reject_CopyObjectFTObjDataTypeNullErr, msgdef.RCsystemerror, "System error", "CopyObjectFTObjDataTypeNullErr"},
	{LMS_Admin_Service_Reject_UpdateDBTableErr, msgdef.RCsystemerror, "System error", "UpdateDBTableErr"},
	{LMS_Admin_Service_Reject_ValidateConditionFTErr, msgdef.RCsystemerror, "System error", "ValidateConditionFTErr"},
	{LMS_Admin_Service_Reject_CreateObjectFTErr, msgdef.RCsystemerror, "System error", "CreateObjectFTErr"},
	{LMS_Admin_Service_Reject_CopyObjectFTErr, msgdef.RCsystemerror, "System error", "CopyObjectFTErr"},
	{LMS_Admin_Service_Reject_DBDeleteFTErr, msgdef.RCsystemerror, "System error", "DBDeleteFTErr"},
	{LMS_Admin_Service_Reject_DBDeleteFTTableNullErr, msgdef.RCsystemerror, "System error", "DBDeleteFTTableNullErr"},
	{LMS_Admin_Service_Reject_DeleteDBTableErr, msgdef.RCsystemerror, "System error", "DeleteDBTableErr"},
	{LMS_Admin_Service_Reject_MicroServiceFunctionNotFoundErr, msgdef.RCsystemerror, "System error", "MicroServiceFunctionNotFoundErr"},
	{LMS_Admin_Service_Reject_MicroServiceFunctionErr, msgdef.RCsystemerror, "System error", "MicroServiceFunctionErr"},
	{LMS_Admin_Service_Reject_MicroServiceNameFormatErr, msgdef.RCsystemerror, "System error", "MicroServiceNameFormatErr"},
	{LMS_Admin_Service_Reject_GetDBContextParamsErr, msgdef.RCsystemerror, "System error", "GetDBContextParamsErr"},
	{LMS_Admin_Service_Reject_CreateDBSequenceFTTableNullErr, msgdef.RCsystemerror, "System error", "CreateDBSequenceFTTableNullErr"},
	{LMS_Admin_Service_Reject_CreateDBSequenceFTFilterDataErr, msgdef.RCsystemerror, "System error", "CreateDBSequenceFTFilterDataErr"},
	{LMS_Admin_Service_Reject_CreateDBSequenceFTErr, msgdef.RCsystemerror, "System error", "CreateDBSequenceFTErr"},
	{LMS_Admin_Service_Reject_GenNextSequenceErr, msgdef.RCsystemerror, "System error", "GenNextSequenceErr"},
	{LMS_Admin_Service_Reject_ExecGroupStartIndexNullErr, msgdef.RCsystemerror, "System error", "ExecGroupStartIndexNullErr"},
	{LMS_Admin_Service_Reject_ExecGroupEndIndexNullErr, msgdef.RCsystemerror, "System error", "ExecGroupEndIndexNullErr"},
	{LMS_Admin_Service_Reject_StartLoopFTErr, msgdef.RCsystemerror, "System error", "StartLoopFTErr"},
	{LMS_Admin_Service_Reject_EndLoopFTErr, msgdef.RCsystemerror, "System error", "EndLoopFTErr"},
	{LMS_Admin_Service_Reject_SetExecGroupIndexErr, msgdef.RCsystemerror, "System error", "SetExecGroupIndexErr"},
	{LMS_Admin_Service_Reject_GetExecGroupIndexErr, msgdef.RCsystemerror, "System error", "GetExecGroupIndexErr"},
	{LMS_Admin_Service_Reject_StoreObjectErr, msgdef.RCsystemerror, "System error", "StoreObjectErr"},
	{LMS_Admin_Service_Reject_GetDBRecordIDErr, msgdef.RCsystemerror, "System error", "GetDBRecordIDErr"},
	{LMS_Admin_Service_Reject_CallMethodFTErr, msgdef.RCsystemerror, "System error", "CallMethodFTErr"},
	{LMS_Admin_Service_Reject_EntitySeqNumFTErr, msgdef.RCsystemerror, "System error", "EntitySeqNumFTErr"},
	{LMS_Admin_Service_Reject_TransformObjectFTErr, msgdef.RCsystemerror, "System error", "TransformObjectFTErr"},
	{LMS_Admin_Service_Reject_SecurityEngInvalidReq, msgdef.RCsystemerror, "System error", "SecurityEngInvalidReq"},
	{LMS_Admin_Service_Reject_LocateSecurityEngAddrFailed, msgdef.RCsystemerror, "System error", "LocateSecurityEngAddrFailed"},
	{LMS_Admin_Service_Reject_SendToSecurityEngTimedOut, msgdef.RCsystemerror, "System error", "SendToSecurityEngTimedOut"},
	{LMS_Admin_Service_Reject_SendToSecurityEngFailed, msgdef.RCsystemerror, "System error", "SendToSecurityEngFailed"},
	{LMS_Admin_Service_Reject_SecurityEngRespUnmarshalFailed, msgdef.RCsystemerror, "System error", "SecurityEngRespUnmarshalFailed"},
	{LMS_Admin_Service_Reject_GenCVVFTErr, msgdef.RCsystemerror, "System error", "GenCVVFTErr"},
	{LMS_Admin_Service_Reject_GenTrackDataFTErr, msgdef.RCsystemerror, "System error", "GenTrackDataFTErr"},
	{LMS_Admin_Service_Reject_CreateMapArrayFTErr, msgdef.RCsystemerror, "System error", "CreateMapArrayFTErr"},
	{LMS_Admin_Service_Reject_GenReqInfoErr, msgdef.RCsystemerror, "System error", "GenReqInfoErr"},
	{LMS_Admin_Service_Reject_MathFunctionFTErr, msgdef.RCsystemerror, "System error", "MathFunctionFTErr"},
	{LMS_Admin_Service_Reject_GenRandomKeyFTErr, msgdef.RCsystemerror, "System error", "GenRandomKeyFTErr"},
	{LMS_Admin_Service_Reject_RawQueryFunctionFTErr, msgdef.RCsystemerror, "System error", "RawQueryFunctionFTErr"},
	{LMS_Admin_Service_Reject_LocateEmailEngAddrFailed, msgdef.RCsystemerror, "System error", "LocateEmailEngAddrFailed"},
	{LMS_Admin_Service_Reject_SendToEmailEngTimedOut, msgdef.RCsystemerror, "System error", "SendToEmailEngTimedOut"},
	{LMS_Admin_Service_Reject_SendToEmailEngFailed, msgdef.RCsystemerror, "System error", "SendToEmailEngFailed"},
	{LMS_Admin_Service_Reject_SendEmailFTErr, msgdef.RCsystemerror, "System error", "SendEmailFTErr"},
	{LMS_Admin_Service_Reject_GenPersoFileFTErr, msgdef.RCsystemerror, "System error", "GenPersoFileFTErr"},
	{LMS_Admin_Service_Reject_LocateUrmgApiEngAddr, msgdef.RCsystemerror, "System error", "LocateUrmgApiEngAddr"},
	{LMS_Admin_Service_Reject_SendToUrmgApiEngTimedOut, msgdef.RCsystemerror, "System error", "SendToUrmgApiEngTimedOut"},
	{LMS_Admin_Service_Reject_SendToUrmgApiEngFailed, msgdef.RCsystemerror, "System error", "SendToUrmgApiEngFailed"},
	{LMS_Admin_Service_Reject_VerifyAdmPortalAccessTokenFailed, msgdef.RCsystemerror, "System error", "VerifyAdmPortalAccessTokenFailed"},
	{LMS_Admin_Service_Reject_DecodeTokenInfoFTErr, msgdef.RCsystemerror, "System error", "DecodeTokenFailed"},
	{LMS_Reject_SendToServiceFTErr, msgdef.RCsystemerror, "System error", "SendToServiceFTErr"},
	{LMS_Admin_Service_Reject_SendToBrokerFTErr, msgdef.RCsystemerror, "System error", "SendToBrokerFTErr"},
	{LMS_Admin_Service_Reject_HandleSendOTPErr, msgdef.RCsystemerror, "System error", "LMS_Admin_Service_Reject_HandleSendOTPErr"},
	{LMS_Admin_Service_Reject_ArithFunctionFTErr, msgdef.RCsystemerror, "System error", "LMS_Admin_Service_Reject_ArithFunctionFTErr"},
	{LMS_Admin_Service_Reject_VerifyLMSAdminAccessTokenFTErr, msgdef.RCsystemerror, "System error", "LMS_Admin_Service_Reject_VerifyLMSAdminAccessTokenFTErr"},
	{LMS_Reject_HandleSendEmailErr, msgdef.RCsystemerror, "System error", "LMS_Reject_HandleSendEmailErr"},
}

func GetLMSAdminServiceRejectInfo(RejectCode string, rejectInfo *RejectDefInfo) int {
	rejectOffset := -1
	for i := 0; i < len(LMS_Admin_Service_RejectTable); i++ {
		if RejectCode == LMS_Admin_Service_RejectTable[i].RejectCode {
			rejectOffset = i
			break
		}
	}
	if rejectOffset < 0 {
		return -1
	}
	*rejectInfo = LMS_Admin_Service_RejectTable[rejectOffset]
	return 1
}
