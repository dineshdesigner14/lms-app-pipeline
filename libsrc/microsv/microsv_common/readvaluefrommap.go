package microsv_common

import (
	"fmt"
	"lmsapieng/include/common/datatypedef"
	"lmsapieng/libsrc/utils/datatypeutil"
	"strings"
)

func ReadValueFromDataMap(reqBrokerDataMap map[string]interface{}, fldName string, fldType string, errDesc *string) interface{} {

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
					if datatypeutil.IsInt(val) {
						val = fmt.Sprintf("%d", val)
					} else {
						*errDesc = fmt.Sprintf("ReadValueFrmDataMapErr Fld[%s] Not a Valid DataType[%s]", fldName, fldType)
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
