package dbtab_tableinfo

import (
	"fmt"
	"lmsapieng/include/common/dbdef"
	"strings"
)

func LoadFromDBTable(queryStr string, resultMapArray *[]map[string]interface{}, dbContext dbdef.DBContextDef, dberr *string, dbrejectreason *string) int {
	if strings.EqualFold(dbContext.DBType, dbdef.DBTypeOracle) {
		if LoadFromDBTableOracle(queryStr, resultMapArray, dbContext, dberr) < 0 {
			//trace.Lg("LoadFromDBTableOracle() failed for queryStr[%s] with dberr[%s]", queryStr, *dberr)
			*dbrejectreason = fmt.Sprintf("LoadFromDBTableOracle failed queryStr[%s] with dberr[%s]", queryStr, *dberr)
			return -1
		}
	} else if strings.EqualFold(dbContext.DBType, dbdef.DBTypePostgreSQL) {
		if LoadFromDBTablePostgres(queryStr, resultMapArray, dbContext, dberr) < 0 {
			//trace.Lg("LoadFromDBTablePostgres() failed for queryStr[%s] with dberr[%s]", queryStr, *dberr)
			*dbrejectreason = fmt.Sprintf("LoadFromDBTablePostgres failed queryStr[%s] with dberr[%s]", queryStr, *dberr)
			return -1
		}
	} else {
		//trace.Lg("db_type(%s) is Invalid...LoadFromDBTable() failed", dbContext.DBType)
		*dbrejectreason = fmt.Sprintf("db_type %s is Invalid...LoadFromDBTable() failed", dbContext.DBType)
		return -1
	}
	//trace.Lg("LoadFromDBTable() Success for queryStr[%s]", queryStr)
	return 1
}

func ReadFromDBTable(queryStr string, resultMap map[string]interface{}, dbContext dbdef.DBContextDef, dberr *string, dbrejectreason *string) int {
	if strings.EqualFold(dbContext.DBType, dbdef.DBTypeOracle) {
		if ReadFromDBTableOracle(queryStr, resultMap, dbContext, dberr) < 0 {
			//trace.Lg("ReadFromDBTableOracle() failed for queryStr[%s] with dberr[%s]", queryStr, *dberr)
			*dbrejectreason = fmt.Sprintf("ReadFromDBTableOracle failed queryStr[%s] with dberr[%s]", queryStr, *dberr)
			return -1
		}
	} else if strings.EqualFold(dbContext.DBType, dbdef.DBTypePostgreSQL) {
		if ReadFromDBTablePostgres(queryStr, resultMap, dbContext, dberr) < 0 {
			//trace.Lg("ReadFromDBTablePostgres() failed for queryStr[%s] with dberr[%s]", queryStr, *dberr)
			*dbrejectreason = fmt.Sprintf("ReadFromDBTablePostgres failed queryStr[%s] with dberr[%s]", queryStr, *dberr)
			return -1
		}
	} else {
		//trace.Lg("db_type(%s) is Invalid...ReadFromDBTable() failed", dbContext.DBType)
		*dbrejectreason = fmt.Sprintf("db_type %s is Invalid...ReadFromDBTable() failed", dbContext.DBType)
		return -1
	}
	// trace.Lg("ReadFromDBTable() Success for queryStr[%s]", queryStr)
	return 1
}

func InsertTable(tableName string, queryStr string, dbContext dbdef.DBContextDef, dberr *string, dbrejectreason *string) int {
	if !dbContext.DBTxFlag {
		//trace.Lg("DBTxFlag not set...InsertTable() failed for Table[%s] queryStr[%s]", tableName, queryStr)
		*dbrejectreason = fmt.Sprintf("DBTxFlag not set...InsertTable() failed for Table[%s] queryStr[%s]", tableName, queryStr)
		return -1
	}
	if strings.EqualFold(dbContext.DBType, dbdef.DBTypeOracle) {
		if InsertTableOracle(tableName, queryStr, dbContext, dberr) < 0 {
			//trace.Lg("InsertTableOracle() failed for Table[%s] queryStr[%s] with dberr[%s]", tableName, queryStr, *dberr)
			*dbrejectreason = fmt.Sprintf("InsertTableOracle failed for Table[%s] queryStr[%s] with dberr[%s]", tableName, queryStr, *dberr)
			return -1
		}
	} else if strings.EqualFold(dbContext.DBType, dbdef.DBTypePostgreSQL) {
		if InsertTablePostgres(tableName, queryStr, dbContext, dberr) < 0 {
			//trace.Lg("InsertTablePostgres() failed for Table[%s] queryStr[%s] with dberr[%s]", tableName, queryStr, *dberr)
			*dbrejectreason = fmt.Sprintf("InsertTablePostgres failed for Table[%s] queryStr[%s] with dberr[%s]", tableName, queryStr, *dberr)
			return -1
		}
	} else {
		//trace.Lg("db_type(%s) is Invalid...InsertTable() for Table[%s] queryStr[%s]", dbContext.DBType, tableName, queryStr)
		*dbrejectreason = fmt.Sprintf("db_type %s is Invalid...InsertTable() for Table[%s] queryStr[%s]", dbContext.DBType, tableName, queryStr)
		return -1
	}
	// trace.Lg("InsertTable() Success for Table[%s]", tableName, queryStr)
	return 1
}

func UpdateTable(tableName string, queryStr string, dbContext dbdef.DBContextDef, dberr *string, dbrejectreason *string) int {
	if !dbContext.DBTxFlag {
		//trace.Lg("DBTxFlag not set...UpdateTable() failed for tableName[%s] queryStr[%s]", tableName, queryStr)
		*dbrejectreason = fmt.Sprintf("DBTxFlag not set...UpdateTable() failed for tableName[%s] queryStr[%s]", tableName, queryStr)
		return -1
	}
	if strings.EqualFold(dbContext.DBType, dbdef.DBTypeOracle) {
		if UpdateTableOracle(tableName, queryStr, dbContext, dberr) < 0 {
			//trace.Lg("UpdateTableOracle() failed for tableName[%s] queryStr[%s] with dberr[%s]", tableName, queryStr, *dberr)
			*dbrejectreason = fmt.Sprintf("UpdateTableOracle failed for tableName[%s] queryStr[%s] with dberr[%s]", tableName, queryStr, *dberr)
			return -1
		}
	} else if strings.EqualFold(dbContext.DBType, dbdef.DBTypePostgreSQL) {
		if UpdateTablePostgres(tableName, queryStr, dbContext, dberr) < 0 {
			//trace.Lg("UpdateTablePostgres() failed for tableName[%s] queryStr[%s] with dberr[%s]", tableName, queryStr, *dberr)
			*dbrejectreason = fmt.Sprintf("UpdateTablePostgres failed for tableName[%s] queryStr[%s] with dberr[%s]", tableName, queryStr, *dberr)
			return -1
		}
	} else {
		//trace.Lg("db_type(%s) is Invalid...UpdateTable() failed for tableName[%s] queryStr[%s]", dbContext.DBType, tableName, queryStr)
		*dbrejectreason = fmt.Sprintf("db_type %s is Invalid...UpdateTable() failed for tableName[%s] queryStr[%s]", dbContext.DBType, tableName, queryStr)
		return -1
	}
	// trace.Lg("UpdateTable() Success for tableName[%s] queryStr[%s]", tableName, queryStr)
	return 1
}

func DeleteTable(tableName string, queryStr string, dbContext dbdef.DBContextDef, dberr *string, dbrejectreason *string) int {
	if !dbContext.DBTxFlag {
		//trace.Lg("DBTxFlag not set...DeleteTable() failed for tableName[%s] queryStr[%s]", tableName, queryStr)
		*dbrejectreason = fmt.Sprintf("DBTxFlag not set...DeleteTable() failed for tableName[%s] queryStr[%s]", tableName, queryStr)
		return -1
	}
	if strings.EqualFold(dbContext.DBType, dbdef.DBTypeOracle) {
		if DeleteTableOracle(tableName, queryStr, dbContext, dberr) < 0 {
			//trace.Lg("DeleteTableOracle() failed for tableName[%s] queryStr[%s] with dberr[%s]", tableName, queryStr, *dberr)
			*dbrejectreason = fmt.Sprintf("DeleteTableOracle failed for tableName[%s] queryStr[%s] with dberr[%s]", tableName, queryStr, *dberr)
			return -1
		}
	} else if strings.EqualFold(dbContext.DBType, dbdef.DBTypePostgreSQL) {
		if DeleteTablePostgres(tableName, queryStr, dbContext, dberr) < 0 {
			//trace.Lg("DeleteTablePostgres() failed for tableName[%s] queryStr[%s] with dberr[%s]", tableName, queryStr, *dberr)
			*dbrejectreason = fmt.Sprintf("DeleteTablePostgres failed for tableName[%s] queryStr[%s] with dberr[%s]", tableName, queryStr, *dberr)
			return -1
		}
	} else {
		//trace.Lg("db_type(%s) is Invalid...DeleteTable() failed for tableName[%s] queryStr[%s]", dbContext.DBType, tableName, queryStr)
		*dbrejectreason = fmt.Sprintf("db_type %s is Invalid...DeleteTable() failed for tableName[%s] queryStr[%s]", dbContext.DBType, tableName, queryStr)
		return -1
	}
	//trace.Lg("DeleteTable() Success for tableName[%s] queryStr[%s]", tableName, queryStr)
	return 1
}
