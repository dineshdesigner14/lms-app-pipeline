package exprutil

import (
	"fmt"
	"lmsapieng/libsrc/utils/lexicalparserutil"

	"github.com/expr-lang/expr"
	"github.com/expr-lang/expr/vm"
)

func IsExprTrue(reqBrokerDataMap map[string]interface{}, expStr string) bool {
	newExpStr := expStr
	if lexicalparserutil.IsFldNameArray(expStr) {
		if lexicalparserutil.ReplaceArrayIndex(reqBrokerDataMap, expStr, &newExpStr) < 0 {
			return false
		}
	}
	program, err := expr.Compile(newExpStr)
	if err != nil {
		//trace.Lg("expr.Compile() failed for expStr[%s] with Err[%s]", newExpStr, err)
		return false
	}
	result, err := vm.Run(program, reqBrokerDataMap)
	if err != nil {
		//trace.Lg("vm.Run() failed for expStr[%s] with Err[%s]", newExpStr, err)
		return false
	}
	boolResult, ok := result.(bool)
	if !ok {
		//trace.Lg("Expression result not a boolean for expStr[%s]", newExpStr)
		return false
	}
	return boolResult
}

func GetNumericValueFromExpr(reqBrokerDataMap map[string]interface{}, expStr string) int {
	newExpStr := expStr
	if lexicalparserutil.IsFldNameArray(expStr) {
		if lexicalparserutil.ReplaceArrayIndex(reqBrokerDataMap, expStr, &newExpStr) < 0 {
			return -1
		}
	}
	program, err := expr.Compile(newExpStr)
	if err != nil {
		//trace.Lg("expr.Compile() failed for expStr[%s] with Err[%s]", newExpStr, err)
		return -1
	}
	result, err := vm.Run(program, reqBrokerDataMap)
	if err != nil {
		//trace.Lg("vm.Run() failed for expStr[%s] with Err[%s]", newExpStr, err)
		return -1
	}
	intResult, ok := result.(int)
	if !ok {
		//trace.Lg("Expression result not a integer for expStr[%s]", newExpStr)
		return -1
	}
	return intResult
}

func EvaulateStrExpr(reqBrokerDataMap map[string]interface{}, expStr string, exprValue *string) int {
	var ok bool
	program, err := expr.Compile(expStr)
	if err != nil {
		//trace.Lg("expr.Compile() failed for expStr[%s] with Err[%s]", expStr, err)
		return -1
	}
	result, err := vm.Run(program, reqBrokerDataMap)
	if err != nil {
		//trace.Lg("vm.Run() failed for expStr[%s] with Err[%s]", expStr, err)
		return -1
	}
	*exprValue, ok = result.(string)
	if !ok {
		intValue, ok := result.(int)
		if ok {
			*exprValue = fmt.Sprintf("%d", intValue)
			return 1
		}
		//trace.Lg("Expression result not a string for expStr[%s]", expStr)
		return -1
	}
	return 1
}

func EvaulateIntExpr(reqBrokerDataMap map[string]interface{}, expStr string, exprValue *int) int {
	var ok bool
	program, err := expr.Compile(expStr)
	if err != nil {
		//trace.Lg("expr.Compile() failed for expStr[%s] with Err[%s]", expStr, err)
		return -1
	}
	result, err := vm.Run(program, reqBrokerDataMap)
	if err != nil {
		//trace.Lg("vm.Run() failed for expStr[%s] with Err[%s]", expStr, err)
		return -1
	}
	*exprValue, ok = result.(int)
	if !ok {
		//trace.Lg("Expression result not a int for expStr[%s]", expStr)
		return -1
	}
	return 1
}

func AssignExpr(reqBrokerDataMap map[string]interface{}, expStr string) int {
	program, err := expr.Compile(expStr)
	if err != nil {
		//trace.Lg("expr.Compile() failed for expStr[%s] with Err[%s]", expStr, err)
		return -1
	}
	_, err = vm.Run(program, reqBrokerDataMap)
	if err != nil {
		//trace.Lg("vm.Run() failed for expStr[%s] with Err[%s]", expStr, err)
		return -1
	}
	return 1
}
