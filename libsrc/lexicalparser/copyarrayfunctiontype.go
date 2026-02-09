package lexicalparser

import (
	"lmsapieng/include/common/lexicalparserdef"
	"lmsapieng/libsrc/utils/exprutil"
	"lmsapieng/libsrc/utils/lexicalparserutil"
	"strconv"
	"strings"
)

func CopyArrayFunctionType(
	execFunction lexicalparserdef.ExecFunctionStruct,
	reqBrokerDataMap map[string]interface{},
) int {

	var rejectDesc string

	objectMap := make(map[string]interface{})

	for _, od := range execFunction.CopyObjectInfo.ObjectData {

		if od.Condition != "" &&
			!exprutil.IsExprTrue(reqBrokerDataMap, od.Condition) {
			continue
		}

		val := lexicalparserutil.ReadValueFrmDataMap(
			reqBrokerDataMap,
			od.DataSource,
			od.DataType,
			&rejectDesc,
		)
		if val == nil {
			return -1
		}

		objectMap[od.Key] = val
	}

	fullPath := execFunction.CopyObjectInfo.ObjectName

	if lexicalparserutil.IsFldNameArray(fullPath) {
		if lexicalparserutil.ReplaceArrayIndex(
			reqBrokerDataMap,
			fullPath,
			&fullPath,
		) < 0 {
			return -1
		}
	}

	lastDot := strings.LastIndex(fullPath, ".")
	if lastDot == -1 {
		return -1
	}

	parentIndexPath := fullPath[:lastDot]
	fieldWithIndex := fullPath[lastDot+1:]
	fieldName := fieldWithIndex[:strings.Index(fieldWithIndex, "[")]

	rootArrayName := parentIndexPath[:strings.Index(parentIndexPath, "[")]
	indexStr := parentIndexPath[strings.Index(parentIndexPath, "[")+1 : len(parentIndexPath)-1]

	idx, err := strconv.Atoi(indexStr)
	if err != nil {
		return -1
	}

	rootAny, ok := reqBrokerDataMap[rootArrayName]
	if !ok {
		return -1
	}

	rootArr, ok := rootAny.([]interface{})
	if !ok || idx >= len(rootArr) {
		return -1
	}

	parentMap, ok := rootArr[idx].(map[string]interface{})
	if !ok {
		parentMap = make(map[string]interface{})
		rootArr[idx] = parentMap
	}

	childAny, exists := parentMap[fieldName]
	var childArr []interface{}

	if !exists {
		childArr = []interface{}{}
	} else {
		childArr, ok = childAny.([]interface{})
		if !ok {
			return -1
		}
	}

	childArr = append(childArr, objectMap)
	parentMap[fieldName] = childArr

	reqBrokerDataMap[rootArrayName] = rootArr

	return 1
}
