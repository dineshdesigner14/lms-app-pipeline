package lexicalparserutil

import (
	"encoding/json"
	"fmt"
	"lmsapieng/include/common/datatypedef"
	"lmsapieng/libsrc/utils/datatypeutil"
	"strconv"
	"strings"
)

func ReadValueFromDataMapArray(reqBrokerDataMap map[string]interface{}, fldName string, fldType string, errDesc *string) interface{} {
	//trace.Lg("ReadValueFromDataMapArray() called for fldName[%s]", fldName)

	current := reqBrokerDataMap
	steps := strings.Split(fldName, ".")
	for i, step := range steps {
		// Check if step contains array index
		if strings.Contains(step, "[") && strings.Contains(step, "]") {
			arraySteps := strings.Split(step, "[")
			field := arraySteps[0]
			indexStr := strings.TrimRight(arraySteps[1], "]")
			index, err := strconv.Atoi(indexStr)
			if err != nil {
				*errDesc = fmt.Sprintf("ReadValueFrmDataMapErr Fld[%s] Invalid array index", fldName)
				return nil
			}
			val, ok := current[field]
			if !ok {
				*errDesc = fmt.Sprintf("ReadValueFrmDataMapErr Fld[%s] Not Found", fldName)
				return nil
			}
			if marr, ok := val.([]map[string]interface{}); ok {
				if index < 0 || index >= len(marr) {
					*errDesc = fmt.Sprintf("ReadValueFrmDataMapErr Fld[%s] Array index out of range", fldName)
					return nil
				}
				val = marr[index]
			}
			if arr, ok := val.([]interface{}); ok {
				if index < 0 || index >= len(arr) {
					*errDesc = fmt.Sprintf("ReadValueFrmDataMapErr Fld[%s] Array index out of range", fldName)
					return nil
				}
				val = arr[index]
			}
			if i == len(steps)-1 {
				if strings.EqualFold(fldType, datatypedef.DataTypeString) {
					if !datatypeutil.IsString(val) {
						if datatypeutil.IsInt(val) {
							val = fmt.Sprintf("%d", val)
						} else if datatypeutil.IsFloat(val) {
							val = fmt.Sprintf("%f", val)
						} else {
							*errDesc = fmt.Sprintf("ReadValueFrmDataMapErr Fld[%s] Not a Valid DataType[%s]", fldName, fldType)
							return nil
						}
					}
				}
				return val
			}
			if m, ok := val.(map[string]interface{}); ok {
				current = m
			} else {
				*errDesc = fmt.Sprintf("ReadValueFrmDataMapErr Fld[%s] Not a map", fldName)
				return nil
			}
		} else {
			val, ok := current[step]
			if !ok {
				*errDesc = fmt.Sprintf("ReadValueFrmDataMapErr Fld[%s] Not Found", fldName)
				return nil
			}

			if i == len(steps)-1 {
				if strings.EqualFold(fldType, datatypedef.DataTypeString) {
					if !datatypeutil.IsString(val) {
						if datatypeutil.IsInt(val) {
							val = fmt.Sprintf("%d", val)
						} else if datatypeutil.IsFloat(val) {
							val = fmt.Sprintf("%f", val)
						} else {
							*errDesc = fmt.Sprintf("ReadValueFrmDataMapErr Fld[%s] Not a Valid DataType[%s]", fldName, fldType)
							return nil
						}
					}
				}
				return val
			}

			if m, ok := val.(map[string]interface{}); ok {
				current = m
			} else {
				*errDesc = fmt.Sprintf("ReadValueFrmDataMapErr Fld[%s] Not a map", fldName)
				return nil
			}
		}
	}
	*errDesc = fmt.Sprintf("ReadValueFrmDataMapErr Fld[%s] Not Found", fldName)
	return nil
}

func IsFldNameArray(fldName string) bool {
	if strings.Contains(fldName, "[") && strings.Contains(fldName, "]") {
		return true
	}
	return false
}

func replaceArrayIndex(reqBrokerDataMap map[string]interface{}, fldName string, indexKey string, newFldName *string) int {
	//trace.Lg("replaceArrayIndex() called for flName[%s]", fldName)
	startIndexFound := false
	indexName := ""
	indexValue := 0
	startIndex := 0
	endIndex := 0
	replacedString := ""
	for i, fldNameChar := range fldName {
		if !startIndexFound && fldName[i] == '[' {
			//trace.Lg("loop...1")
			startIndexFound = true
			startIndex = i
			continue
		}
		if startIndexFound && fldName[i] != '[' && fldName[i] != ']' && fldName[i] != ' ' && fldName[i] != '\t' && fldName[i] != '\n' {
			//trace.Lg("loop...2")
			indexName += fmt.Sprintf("%c", fldNameChar)
			continue
		}
		if startIndexFound && fldName[i] == ']' {
			endIndex = i
			//trace.Lg("fldName[%s]indexName[%s]", fldName, indexName)
			if indexKey == indexName {
				if GetIndexValue(reqBrokerDataMap, string(indexName), &indexValue) < 0 {
					return -1
				}
				replacedString = fldName[:startIndex+1] + fmt.Sprintf("%d", indexValue) + fldName[endIndex:]
				//trace.Lg("replacedString[%s]", replacedString)
				*newFldName = replacedString
			}
			indexName = ""
			startIndexFound = false
			if indexKey == indexName {
				break
			}
		}
	}
	return 1
}

func ReplaceArrayIndex(reqBrokerDataMap map[string]interface{}, fldName string, newFldName *string) int {
	//trace.Lg("ReplaceArrayIndex() called for fldName[%s]")
	indexArray := make([]string, 0)
	startIndexFound := false
	indexName := ""
	for i, fldNameChar := range fldName {
		if !startIndexFound && fldName[i] == '[' {
			startIndexFound = true
			continue
		}
		if startIndexFound && fldName[i] != '[' && fldName[i] != ']' && fldName[i] != ' ' && fldName[i] != '\t' && fldName[i] != '\n' {
			indexName += fmt.Sprintf("%c", fldNameChar)
			continue
		}
		if startIndexFound && fldName[i] == ']' {
			indexArray = append(indexArray, indexName)
			indexName = ""
			startIndexFound = false
		}
	}
	tempFldName := fldName
	for i := 0; i < len(indexArray); i++ {
		if replaceArrayIndex(reqBrokerDataMap, tempFldName, indexArray[i], newFldName) < 0 {
			return -1
		}
		//trace.Lg("tempFldName[%s] newFldName[%s]", tempFldName, *newFldName)
		tempFldName = *newFldName
	}
	//trace.Lg("ReplaceArrayIndex() success for fldName[%s] newFldName[%s]", fldName, *newFldName)
	return 1
}

func ReadValueFrmDataMap(reqBrokerDataMap map[string]interface{}, fldName string, fldType string, errDesc *string) interface{} {
	var replacedString string
	if IsFldNameArray(fldName) {
		startIndex := strings.Index(fldName, "[")
		endIndex := strings.Index(fldName, "]")
		if startIndex == -1 || endIndex == -1 || endIndex <= startIndex+1 {
			*errDesc = fmt.Sprintf("ReadValueFrmDataMapErr Fld[%s] Not is not a valid Array", fldName)
			return nil
		}
		if ReplaceArrayIndex(reqBrokerDataMap, fldName, &replacedString) < 0 {
			*errDesc = fmt.Sprintf("replaceArrayIndex() failed for fldName[%s]", fldName)
			return nil
		}
		//trace.Lg("replacedString[%s]", replacedString)
		return ReadValueFromDataMapArray(reqBrokerDataMap, replacedString, fldType, errDesc)
	}
	current := reqBrokerDataMap
	steps := strings.Split(fldName, ".")
	for i, step := range steps {
		val, ok := current[step]
		if !ok {
			*errDesc = fmt.Sprintf("ReadValueFrmDataMapErr Fld[%s] Not Found", fldName)
			return nil
		}
		if i == len(steps)-1 {
			if strings.EqualFold(fldType, datatypedef.DataTypeString) {
				if !datatypeutil.IsString(val) {
					if i, ok := val.(int); ok {
						val = fmt.Sprintf("%d", i)
					} else if i64, ok := val.(int64); ok {
						val = fmt.Sprintf("%d", i64)
					} else if f64, ok := val.(float64); ok {
						if f64 == float64(int64(f64)) {
							val = fmt.Sprintf("%d", int64(f64))
						} else {
							val = fmt.Sprintf("%f", f64)
						}
					} else if num, ok := val.(json.Number); ok {
						val = num.String()
					} else {
						*errDesc = fmt.Sprintf("ReadValueFrmDataMapErr Fld[%s] Not a Valid DataType[%s], got [%T]", fldName, fldType, val)
						return nil
					}
				}
			} else if strings.EqualFold(fldType, datatypedef.DataTypeInt) {
				if !datatypeutil.IsInt(val) {
					*errDesc = fmt.Sprintf("ReadValueFrmDataMapErr Fld[%s] Not a Valid DataType[%s]", fldName, fldType)
					return nil
				}
			} else if strings.EqualFold(fldType, datatypedef.DataTypeObject) {
				if !datatypeutil.IsObject(val) {
					*errDesc = fmt.Sprintf("ReadValueFrmDataMapErr Fld[%s] Not a Valid DataType[%s]", fldName, fldType)
					return nil
				}
			} else if strings.EqualFold(fldType, datatypedef.DataTypeObjectArray) {
				if !datatypeutil.IsObjectArray(val) {
					*errDesc = fmt.Sprintf("ReadValueFrmDataMapErr Fld[%s] Not a Valid DataType[%s]", fldName, fldType)
					return nil
				}
			} else if strings.EqualFold(fldType, datatypedef.DataTypeBoolean) {
				if !datatypeutil.IsBool(val) {
					//trace.Lg("Bool[%s]", fldType)
					*errDesc = fmt.Sprintf("ReadValueFrmDataMapErr Fld[%s] Not a Valid DataType[%s]", fldName, fldType)
					return nil
				}
			} else {
				*errDesc = fmt.Sprintf("ReadValueFrmDataMapErr DataType[%s] Not Supported for Fld[%s]", fldType, fldName)
				return nil
			}
			return val
		}
		if m, ok := val.(map[string]interface{}); ok {
			current = m
		}
	}
	*errDesc = fmt.Sprintf("ReadValueFrmDataMapErr Fld[%s] Not Found", fldName)
	return nil
}
