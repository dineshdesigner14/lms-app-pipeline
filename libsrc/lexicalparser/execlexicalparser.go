package lexicalparser

import (
	"fmt"
	"lmsapieng/include/common/functiondef"
	"lmsapieng/include/common/lexicalparserdef"
	"lmsapieng/include/common/rejectdef"
	"lmsapieng/include/common/reqbrokerdef"
	"lmsapieng/libsrc/utils/exprutil"
	"lmsapieng/libsrc/utils/lexicalparserutil"
	"lmsapieng/libsrc/utils/msgutil"
	"lmsapieng/libsrc/utils/trace"
)

func ExecLexicalParser(parserInfo lexicalparserdef.LexicalParserStruct, reqBrokerDataMap map[string]interface{}) int {
	// trace.Lg("ExecLexicalParser() running Microservice[%s]", parserInfo.Microservice)
	reqBrokerDataMap["ExecGroupIndex"] = 0
	for i := 0; i < len(parserInfo.ExecGroup); i++ {
		if len(parserInfo.ExecGroup[i].GroupCondition) != 0 {
			if !exprutil.IsExprTrue(reqBrokerDataMap, parserInfo.ExecGroup[i].GroupCondition) {
				//trace.Lg("isExecGroupConditionSatisfied() failed ExecGroupName[%s] for GroupCondition[%s]", parserInfo.ExecGroup[i].GroupName, parserInfo.ExecGroup[i].GroupCondition)
				continue
			}
		}
		if executeGroup(parserInfo.ExecGroup[i], reqBrokerDataMap) < 0 {
			return -1
		}
	}
	return 1
}

func executeGroup(execGroup lexicalparserdef.ExecGroupStruct, reqBrokerDataMap map[string]interface{}) int {
	var rejectDesc, correctiveAction string
	if len(execGroup.ArrayIndex) == 0 {
		for i := 0; i < len(execGroup.ExecFunction); i++ {
			// trace.Lg("ExecFunction->FunctionName[%s] FunctionType[%s]", execGroup.ExecFunction[i].FunctionName, execGroup.ExecFunction[i].FunctionType)
			if lexicalparserutil.SetExecGroupIndex(reqBrokerDataMap, i) < 0 {
				rejectDesc = fmt.Sprintf("SetExecGroupIndex Failed for execGroup[%s]", execGroup.GroupName)
				correctiveAction = fmt.Sprintf("Correct the SetExecGroupIndex for execGroup[%s]", execGroup.GroupName)
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_SetExecGroupIndexErr, []byte(rejectDesc), []byte(correctiveAction), rejectDesc)
				return -1
			}
			if len(execGroup.ExecFunction[i].FunctionCondition) != 0 {
				if !exprutil.IsExprTrue(reqBrokerDataMap, execGroup.ExecFunction[i].FunctionCondition) {
					continue
				}
				//trace.Lg("FunctionCondition[%s] is true for functionName[%s]", execGroup.ExecFunction[i].FunctionCondition, execGroup.ExecFunction[i].FunctionName)
			}
			if execGroup.ExecFunction[i].FunctionType == functiondef.FunctionTypeStartLoop {
				endIndex := exprutil.GetNumericValueFromExpr(reqBrokerDataMap, execGroup.ExecFunction[i].EndIndex)
				if endIndex <= 0 {
					endIndexFound := false
					newIndex := 0
					for j := i + 1; j < len(execGroup.ExecFunction); j++ {
						if execGroup.ExecFunction[j].FunctionType == functiondef.FunctionTypeEndLoop && execGroup.ExecFunction[j].IndexName == execGroup.ExecFunction[i].IndexName {
							newIndex = j
							endIndexFound = true
							break
						}
					}
					if !endIndexFound {
						rejectDesc = fmt.Sprintf("EndLoop Not Found for StartLoop With FunctionName[%s]", execGroup.ExecFunction[i].FunctionName)
						correctiveAction = fmt.Sprintf("Include EndLoop for StartLoop With FunctionName[%s]", execGroup.ExecFunction[i].FunctionName)
						reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GetExecGroupIndexErr, []byte(rejectDesc), []byte(correctiveAction), rejectDesc)
						return -1
					}
					i = newIndex
					continue
				}
			}
			if executeFunction(execGroup.ExecFunction[i], reqBrokerDataMap) < 0 {
				return -1
			}
			execGroupIndex := lexicalparserutil.GetExecGroupIndex(reqBrokerDataMap)
			if execGroupIndex < 0 {
				rejectDesc = fmt.Sprintf("GetExecGroupIndex() failed for execGroup[%s]", execGroup.GroupName)
				correctiveAction = fmt.Sprintf("Check Properly the GetExecGroupIndex() for execGroup[%s]", execGroup.GroupName)
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GetExecGroupIndexErr, []byte(rejectDesc), []byte(correctiveAction), rejectDesc)
				return -1
			}
			i = execGroupIndex
		}
		return 1
	}
	if len(execGroup.StartIndex) == 0 {
		rejectDesc = fmt.Sprintf("start_index is null for execGroup[%s]", execGroup.GroupName)
		correctiveAction = fmt.Sprintf("check why start_index is null for execGroup[%s]", execGroup.GroupName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ExecGroupStartIndexNullErr, []byte(rejectDesc), []byte(correctiveAction), rejectDesc)
		return -1
	}
	if len(execGroup.EndIndex) == 0 {
		rejectDesc = fmt.Sprintf("end_index is null for execGroup[%s]", execGroup.GroupName)
		correctiveAction = fmt.Sprintf("check why end_index is null for execGroup[%s]", execGroup.GroupName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ExecGroupEndIndexNullErr, []byte(rejectDesc), []byte(correctiveAction), rejectDesc)
		return -1
	}
	startIndex := exprutil.GetNumericValueFromExpr(reqBrokerDataMap, execGroup.StartIndex)
	if startIndex < 0 {
		rejectDesc = fmt.Sprintf("start_index is invalid for execGroup[%s]", execGroup.GroupName)
		correctiveAction = fmt.Sprintf("check why start_index is invalid for execGroup[%s]", execGroup.GroupName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ExecGroupStartIndexNullErr, []byte(rejectDesc), []byte(correctiveAction), rejectDesc)
		return -1
	}
	endIndex := exprutil.GetNumericValueFromExpr(reqBrokerDataMap, execGroup.EndIndex)
	if endIndex < 0 {
		rejectDesc = fmt.Sprintf("end_index is invalid for execGroup[%s]", execGroup.GroupName)
		correctiveAction = fmt.Sprintf("check why end_index is invalid for execGroup[%s]", execGroup.GroupName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ExecGroupEndIndexNullErr, []byte(rejectDesc), []byte(correctiveAction), rejectDesc)
		return -1
	}
	//trace.Lg("startIndex[%d] endIndex[%d] for execGroup[%s]", startIndex, endIndex, execGroup.GroupName)
	for i := startIndex; i < endIndex; i++ {
		reqBrokerDataMap[lexicalparserdef.LexicalParserCurrentIndex] = fmt.Sprintf("%d", i)
		for j := 0; j < len(execGroup.ExecFunction); j++ {
			//trace.Lg("ExecFunction->FunctionName[%s] FunctionType[%s] for Index[%d]", execGroup.ExecFunction[j].FunctionName, execGroup.ExecFunction[j].FunctionType, i)
			if len(execGroup.ExecFunction[i].FunctionCondition) != 0 {
				if !exprutil.IsExprTrue(reqBrokerDataMap, execGroup.ExecFunction[i].FunctionCondition) {
					continue
				}
			}
			if executeFunction(execGroup.ExecFunction[j], reqBrokerDataMap) < 0 {
				return -1
			}
		}
	}
	return 1
}

func executeFunction(execFunction lexicalparserdef.ExecFunctionStruct, reqBrokerDataMap map[string]interface{}) int {
	trace.Lg("executeFunction() called for FunctionName[%s] FunctionType[%s]", execFunction.FunctionName, execFunction.FunctionType)
	//trace.Lg("reqBrokerDataBuffer[%s]", reqbrokerutil.GetReqBrokerDataMapBuffer(reqBrokerDataMap))
	if execFunction.FunctionType == functiondef.FunctionTypeCallMethod {
		return CallMethodFunctionType(execFunction, reqBrokerDataMap)
	} else if execFunction.FunctionType == functiondef.FunctionTypeValidateField {
		return ValidateFunctionType(execFunction, reqBrokerDataMap)
	} else if execFunction.FunctionType == functiondef.FunctionTypeValidateCondition {
		return ValidateConditionFunctionType(execFunction, reqBrokerDataMap)
	} else if execFunction.FunctionType == functiondef.FunctionTypeDBSingleRead {
		return DBSingleReadFunctionType(execFunction, reqBrokerDataMap)
	} else if execFunction.FunctionType == functiondef.FunctionTypeDBMultiRead {
		return DBMultiReadFunctionType(execFunction, reqBrokerDataMap)
	} else if execFunction.FunctionType == functiondef.FunctionTypeDBInsert {
		return DBInsertFunctionType(execFunction, reqBrokerDataMap)
	} else if execFunction.FunctionType == functiondef.FunctionTypeDBUpdate {
		return DBUpdateFunctionType(execFunction, reqBrokerDataMap)
	} else if execFunction.FunctionType == functiondef.FunctionTypeDBDelete {
		return DBDeleteFunctionType(execFunction, reqBrokerDataMap)
	} else if execFunction.FunctionType == functiondef.FunctionTypeCreateObject {
		return CreateObjectFunctionType(execFunction, reqBrokerDataMap)
	} else if execFunction.FunctionType == functiondef.FunctionTypeCopyObject {
		return CopyObjectFunctionType(execFunction, reqBrokerDataMap)
	} else if execFunction.FunctionType == functiondef.FunctionTypeComposeResponse {
		return ComposeRespFunctionType(execFunction, reqBrokerDataMap)
	} else if execFunction.FunctionType == functiondef.FunctionTypeCreateDBSequence {
		return CreateDBSequenceFunctionType(execFunction, reqBrokerDataMap)
	} else if execFunction.FunctionType == functiondef.FunctionTypeStartLoop {
		return StartLoopFunctionType(execFunction, reqBrokerDataMap)
	} else if execFunction.FunctionType == functiondef.FunctionTypeEndLoop {
		return EndLoopFunctionType(execFunction, reqBrokerDataMap)
	} else if execFunction.FunctionType == functiondef.FunctionTypeEntitySeqNum {
		return EntitySeqNumFunctionType(execFunction, reqBrokerDataMap)
	} else if execFunction.FunctionType == functiondef.FunctionTypeTransformObject {
		return TransformObjectFunctionType(execFunction, reqBrokerDataMap)
	} else if execFunction.FunctionType == functiondef.FunctionTypeGenCVV {
		return GenCVVFunctionType(execFunction, reqBrokerDataMap)
	} else if execFunction.FunctionType == functiondef.FunctionTypeGenTrackData {
		return GenTrackDataFunctionType(execFunction, reqBrokerDataMap)
	} else if execFunction.FunctionType == functiondef.FunctionTypeCreateMapArray {
		return CreateMapArrayFunctionType(execFunction, reqBrokerDataMap)
	} else if execFunction.FunctionType == functiondef.FunctionTypeGenPersoFile {
		return GenPersoFileFunctionType(execFunction, reqBrokerDataMap)
	} else if execFunction.FunctionType == functiondef.FunctionTypeMathFunction {
		return MathFunctionType(execFunction, reqBrokerDataMap)
	} else if execFunction.FunctionType == functiondef.FunctionTypeGenRandomKey {
		return GenRandomKeyFunctionType(execFunction, reqBrokerDataMap)
	} else if execFunction.FunctionType == functiondef.FunctionTypeRawQuery {
		return RawQueryFunctionType(execFunction, reqBrokerDataMap)
	} else if execFunction.FunctionType == functiondef.FunctionTypeSendEmail {
		return SendEmailFunctionType(execFunction, reqBrokerDataMap)
	} else if execFunction.FunctionType == functiondef.FunctionTypeSendSMS {
		return SendSMSFunctionType(execFunction, reqBrokerDataMap)
	} else if execFunction.FunctionType == functiondef.FunctionTypeGenToken {
		return GenTokenFunctionType(execFunction, reqBrokerDataMap)
	} else if execFunction.FunctionType == functiondef.FunctionTypeVerifyToken {
		return VerifyTokenFunctionType(execFunction, reqBrokerDataMap)
	} else if execFunction.FunctionType == functiondef.FunctionTypeSendToService {
		return SendToServiceFunctionType(execFunction, reqBrokerDataMap)
	} else if execFunction.FunctionType == functiondef.FunctionTypeDecodeToken {
		return DecodeTokenFunctionType(execFunction, reqBrokerDataMap)
	} else if execFunction.FunctionType == functiondef.FunctionTypeCreateEmptyList {
		return CreateEmptyListFunctionType(execFunction, reqBrokerDataMap)
	} else if execFunction.FunctionType == functiondef.FunctionTypeSendToBroker {
		return SendToBrokerFunctionType(execFunction, reqBrokerDataMap)
	} else if execFunction.FunctionType == functiondef.FunctionTypeArithmetic {
		return ArithmeticFunctionType(execFunction, reqBrokerDataMap)
	} else if execFunction.FunctionType == functiondef.FunctionTypeValidateConditionAndUpdateOnError {
		return ValidateConditionAndUpdateOnErrorFunctionType(execFunction, reqBrokerDataMap)
	} else if execFunction.FunctionType == functiondef.FunctionTypeCopyArray {
		return CopyArrayFunctionType(execFunction, reqBrokerDataMap)
	} else {
		rejectDesc := fmt.Sprintf("FunctionType[%s] not valid", execFunction.FunctionType)
		correctiveAction := fmt.Sprintf("Include Valid FunctionType and Not Invalid FunctionType[%s]", execFunction.FunctionType)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_FunctionTypeInvalidErr, []byte(rejectDesc), []byte(correctiveAction), rejectDesc)
		return -1
	}
}
