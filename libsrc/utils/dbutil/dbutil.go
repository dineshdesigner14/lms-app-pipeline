package dbutil

import (
	"lmsapieng/include/common/dbdef"
	"lmsapieng/libsrc/utils/cryptoutil"
	"strings"
)

func IsNoRows(DBType string, DBErr string) bool {
	if strings.EqualFold(DBType, dbdef.DBTypeOracle) || strings.EqualFold(DBType, dbdef.DBTypePostgreSQL) {
		if strings.EqualFold(DBErr, "sql: no rows in result set") {
			return true
		}
	}
	return false
}

func IsDuplicateRows(DBType string, DBErr string) bool {
	if strings.EqualFold(DBType, dbdef.DBTypeOracle) || strings.EqualFold(DBType, dbdef.DBTypePostgreSQL) {
		if strings.EqualFold(DBErr, "sql: duplicate rows in result set") {
			return true
		}
	}
	return false
}

func IsDataAlreadyExists(DBType string, DBErr string) bool {
	if strings.EqualFold(DBType, dbdef.DBTypeOracle) || strings.EqualFold(DBType, dbdef.DBTypePostgreSQL) {
		if strings.EqualFold(DBErr, "Data Already Exists") {
			return true
		}
	}
	return false
}

func DecryptDBCredential(inputVal string) (int, string) {
	var ouputVal, errDesc string
	if cryptoutil.AESDecryptText(dbdef.DBKey, inputVal, &ouputVal, &errDesc) < 0 {
		//trace.Lg("AESDecryptText() failed with errDesc:%s", errDesc)
		return -1, ""
	}
	return 1, ouputVal
}

func PrintActiveDBConnections(dbcontextmap map[string]interface{}) {
	_, ok := dbcontextmap[dbdef.DBContextArrayObj]
	if !ok {
		return
	}
	activeContextArray := dbcontextmap[dbdef.DBContextArrayObj].([]dbdef.DBContextDef)
	//trace.Lg("No. of active dbconnections [%d]", len(activeContextArray))
	for i := 0; i < len(activeContextArray); i++ {
		//trace.Lg("ActiveConnection->RecordNum<%s> ConnectionStr<%s>", activeContextArray[i].RecordNum, activeContextArray[i].ConnectionStr)
	}
}
