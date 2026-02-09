package lexicalparser

import (
	"fmt"
	"lmsapieng/include/common/lexicalparserdef"
	"lmsapieng/include/common/rejectdef"
	"lmsapieng/include/common/reqbrokerdef"
	"lmsapieng/libsrc/utils/amtutil"
	"lmsapieng/libsrc/utils/lexicalparserutil"
	"lmsapieng/libsrc/utils/msgutil"
	"lmsapieng/libsrc/utils/templateutil"
	"lmsapieng/libsrc/utils/trace"
)

func ArithmeticFunctionType(execFunction lexicalparserdef.ExecFunctionStruct, reqBrokerDataMap map[string]interface{}) int {
	var rejectDesc string
	if len(execFunction.ArithFunctionInfo.Operation) == 0 {
		rejectDesc = fmt.Sprintf("Operation is NULL in ArithFunctionInfo for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ArithFunctionFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ArithFunctionInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ArithFunctionInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	if len(execFunction.ArithFunctionInfo.LeftOperand) == 0 {
		rejectDesc = fmt.Sprintf("LeftOperand is NULL in MathFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ArithFunctionFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ArithFunctionInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ArithFunctionInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	dataValue := lexicalparserutil.ReadValueFrmDataMap(reqBrokerDataMap, execFunction.ArithFunctionInfo.LeftOperand, "string", &rejectDesc)
	if dataValue == nil {
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ArithFunctionFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ArithFunctionInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ArithFunctionInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	lLeftOperand := dataValue.(string)
	if len(execFunction.ArithFunctionInfo.RightOperand) == 0 {
		rejectDesc = fmt.Sprintf("RightOperand is NULL in MathFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ArithFunctionFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ArithFunctionInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ArithFunctionInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	dataValue = lexicalparserutil.ReadValueFrmDataMap(reqBrokerDataMap, execFunction.ArithFunctionInfo.RightOperand, "string", &rejectDesc)
	if dataValue == nil {
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ArithFunctionFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ArithFunctionInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ArithFunctionInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	lRightOperand := dataValue.(string)
	if len(execFunction.ArithFunctionInfo.DestObject) == 0 {
		rejectDesc = fmt.Sprintf("DestObject is NULL in MathFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ArithFunctionFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ArithFunctionInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ArithFunctionInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	op := execFunction.ArithFunctionInfo.Operation
	if op == "" {
		rejectDesc = fmt.Sprintf("Operation is NULL in MathFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ArithFunctionFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ArithFunctionInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ArithFunctionInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	var result string

	switch op {
	case "sum":
		result = amtutil.AddAmt(lLeftOperand, lRightOperand, 0)

	case "sub":
		result = amtutil.SubAmt(lLeftOperand, lRightOperand, 0)

	case "multi":
		result = amtutil.MulAmt(lLeftOperand, lRightOperand, 2)

	case "div":
		if lRightOperand == "0" || lRightOperand == "0.00" {
			rejectDesc = fmt.Sprintf("Division by zero in MathFunctionType for FunctionName[%s]", execFunction.FunctionName)
			reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(
				rejectdef.LMS_Admin_Service_Reject_ArithFunctionFTErr,
				templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ArithFunctionInfo.AppErrDesc),
				templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ArithFunctionInfo.AppCorrectiveAction),
				rejectDesc)
			return -1
		}
		result = amtutil.DivAmt(lLeftOperand, lRightOperand, 2)

	default:
		rejectDesc = fmt.Sprintf("Unknown Operation[%s] in MathFunctionType FunctionName[%s]",
			op, execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_ArithFunctionFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ArithFunctionInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ArithFunctionInfo.AppCorrectiveAction), rejectDesc)
		return -1

	}
	trace.Lg("result [%s]", result)
	reqBrokerDataMap[execFunction.ArithFunctionInfo.DestObject] = result
	return 1
}
