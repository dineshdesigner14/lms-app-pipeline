package lexicalparser

import (
	"fmt"
	"lmsapieng/include/common/lexicalparserdef"
	"lmsapieng/include/common/rejectdef"
	"lmsapieng/include/common/reqbrokerdef"
	"lmsapieng/libsrc/utils/exprutil"
	"lmsapieng/libsrc/utils/lexicalparserutil"
	"lmsapieng/libsrc/utils/maputil"
	"lmsapieng/libsrc/utils/msgutil"
	"lmsapieng/libsrc/utils/templateutil"
	"lmsapieng/libsrc/utils/trace"
	"strings"
)

func CreateMapArrayFunctionType(
	execFunction lexicalparserdef.ExecFunctionStruct,
	reqBrokerDataMap map[string]interface{},
) int {

	var rejectDesc string

	trace.Lg("TRACE: Enter CreateMapArrayFunctionType")
	trace.Lg("TRACE: FunctionName = %s", execFunction.FunctionName)
	trace.Lg("TRACE: ObjectName = %s", execFunction.MapArrayInfo.ObjectName)
	trace.Lg("TRACE: ArraySizeExpr = %s", execFunction.MapArrayInfo.ArraySize)

	// ================= VALIDATION =================
	if len(execFunction.MapArrayInfo.ObjectName) == 0 {
		rejectDesc = fmt.Sprintf(
			"ObjectName is NULL in CreateMapArrayFunctionType for FunctionName[%s]",
			execFunction.FunctionName,
		)
		trace.Lg("ERROR: %s", rejectDesc)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] =
			msgutil.SetResp(
				rejectdef.LMS_Admin_Service_Reject_CreateMapArrayFTErr,
				templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppErrDesc),
				templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppCorrectiveAction),
				rejectDesc,
			)
		return -1
	}

	if len(execFunction.MapArrayInfo.ArraySize) == 0 {
		rejectDesc = fmt.Sprintf(
			"ArraySize is NULL in CreateMapArrayFunctionType for FunctionName[%s]",
			execFunction.FunctionName,
		)
		trace.Lg("ERROR: %s", rejectDesc)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] =
			msgutil.SetResp(
				rejectdef.LMS_Admin_Service_Reject_CreateMapArrayFTErr,
				templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppErrDesc),
				templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppCorrectiveAction),
				rejectDesc,
			)
		return -1
	}

	arraySize := 0
	if exprutil.EvaulateIntExpr(
		reqBrokerDataMap,
		execFunction.MapArrayInfo.ArraySize,
		&arraySize,
	) < 0 {
		rejectDesc = fmt.Sprintf(
			"Expr[%s] failed in CreateMapArrayFunctionType for FunctionName[%s]",
			execFunction.MapArrayInfo.ArraySize,
			execFunction.FunctionName,
		)
		trace.Lg("ERROR: %s", rejectDesc)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] =
			msgutil.SetResp(
				rejectdef.LMS_Admin_Service_Reject_CreateMapArrayFTErr,
				templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppErrDesc),
				templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ResponseArray.AppCorrectiveAction),
				rejectDesc,
			)
		return -1
	}

	trace.Lg("TRACE: Evaluated arraySize = %d", arraySize)

	objectName := execFunction.MapArrayInfo.ObjectName

	// ======================================================
	// CASE 1: SIMPLE ARRAY → questions
	// ======================================================
	if !strings.Contains(objectName, ".") && !strings.Contains(objectName, "[") {

		trace.Lg("TRACE: Detected SIMPLE array creation for [%s]", objectName)

		arr := make([]interface{}, arraySize)
		for i := 0; i < arraySize; i++ {
			arr[i] = make(map[string]interface{})
		}

		trace.Lg("TRACE: Creating root array [%s] with %d slots", objectName, arraySize)

		if maputil.SetValueFromString(reqBrokerDataMap, objectName, arr) < 0 {
			trace.Lg("ERROR: SetValueFromString failed for [%s]", objectName)
			return -1
		}

		trace.Lg("TRACE: Successfully created root array [%s]", objectName)
		trace.Lg("TRACE: Exit CreateMapArrayFunctionType")
		return 1
	}

	// ======================================================
	// CASE 2: NESTED ARRAY → questions[i].options
	// ======================================================
	trace.Lg("TRACE: Detected NESTED array creation")

	parts := strings.Split(objectName, ".")
	currentPath := ""

	for _, part := range parts {

		if currentPath == "" {
			currentPath = part
		} else {
			currentPath = currentPath + "." + part
		}

		trace.Lg("TRACE: Processing path segment [%s]", currentPath)

		resolvedPath := currentPath

		// Resolve index (i, j, k...)
		if lexicalparserutil.IsFldNameArray(resolvedPath) {
			trace.Lg("TRACE: Resolving array index in [%s]", resolvedPath)

			if lexicalparserutil.ReplaceArrayIndex(
				reqBrokerDataMap,
				resolvedPath,
				&resolvedPath,
			) < 0 {
				trace.Lg("ERROR: ReplaceArrayIndex failed for [%s]", resolvedPath)
				return -1
			}

			trace.Lg("TRACE: Resolved path = [%s]", resolvedPath)
		}

		// Create array slots if indexed
		if strings.Contains(part, "[") {

			trace.Lg(
				"TRACE: Creating array slots at [%s] with size [%d]",
				resolvedPath,
				arraySize,
			)

			arr := make([]interface{}, arraySize)
			for i := 0; i < arraySize; i++ {
				arr[i] = make(map[string]interface{})
			}

			if maputil.SetValueFromString(
				reqBrokerDataMap,
				resolvedPath,
				arr,
			) < 0 {
				trace.Lg("ERROR: SetValueFromString failed for [%s]", resolvedPath)
				return -1
			}

			trace.Lg("TRACE: Successfully created array at [%s]", resolvedPath)
		}
	}

	trace.Lg("TRACE: Exit CreateMapArrayFunctionType")
	return 1
}
