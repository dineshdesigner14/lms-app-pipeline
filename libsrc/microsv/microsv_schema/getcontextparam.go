package microsv_schema

import (
	"lmsapieng/include/common/datatypedef"
	"lmsapieng/include/common/dbdef"
	"lmsapieng/include/common/lexicalparserdef"
	"lmsapieng/libsrc/microsv/microsv_common"
	"strings"
)

func GetDBContextParams(reqBrokerDataMap map[string]interface{}, schemaInfo lexicalparserdef.LPSchemaInfoStruct, rejectDesc *string) (int, []string) {
	contextParams := ""
	if len(schemaInfo.ModuleName) != 0 {
		contextParams += schemaInfo.ModuleName
	} else {
		contextParams += dbdef.DBModuleLMS
	}
	contextParams += ","
	if len(schemaInfo.SubModuleName) != 0 {
		contextParams += schemaInfo.SubModuleName
	} else {
		contextParams += "*"
	}
	contextParams += ","
	if len(schemaInfo.InstID) != 0 {
		dataValue := microsv_common.ReadValueFromDataMap(reqBrokerDataMap, schemaInfo.InstID, datatypedef.DataTypeString, rejectDesc)
		if dataValue == nil {
			//trace.Lg("ReadValueFromDataMap() failed for schemaInfo.InstID[%s]", schemaInfo.InstID)
			return -1, nil
		}
		contextParams += dataValue.(string)
	} else {
		contextParams += "*"
	}
	contextParams += ","
	contextParams += "*"
	contextParams += ","
	if len(schemaInfo.BinID) != 0 {
		dataValue := microsv_common.ReadValueFromDataMap(reqBrokerDataMap, schemaInfo.BinID, datatypedef.DataTypeString, rejectDesc)
		if dataValue == nil {
			//trace.Lg("ReadValueFromDataMap() failed for schemaInfo.BinID[%s]", schemaInfo.BinID)
			return -1, nil
		}
		contextParams += dataValue.(string)
	} else {
		contextParams += "*"
	}
	contextParams += ","
	contextParams += "*"
	return 1, strings.Split(contextParams, ",")
}
