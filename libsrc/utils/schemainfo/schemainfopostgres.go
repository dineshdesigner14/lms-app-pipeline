package schemainfo

import (
	"database/sql"
	"fmt"
	"lmsapieng/include/common/moduledef"
)

func readModuleFromSchemaPostgres(globalSchemaInfoDBID *sql.DB, recordNum string, schemaName string) int {
	var err error
	var lRecordNum sql.NullString
	var lModuleName sql.NullString
	var lSchemaStatus sql.NullString

	tableName := fmt.Sprintf(`"%s"."%s"`, schemaName, "MODULE_SCHEMA_MAP")

	queryStmt := fmt.Sprintf(
		`SELECT 
			"SCHEMA_RECORD_NUM",
			"MODULE_NAME",
			"SCHEMA_STATUS"
		FROM %s WHERE "SCHEMA_RECORD_NUM"='%s' AND "MODULE_NAME"='%s'`, tableName, recordNum, moduledef.LMSApiEngModule)

	err = globalSchemaInfoDBID.QueryRow(queryStmt).Scan(
		&lRecordNum,
		&lModuleName,
		&lSchemaStatus,
	)
	if err != nil {
		//trace.Lg("QueryRow() for (%s) failed with err(%s)", queryStmt, err)
		return -1
	}
	return 1
}

func loadGlobalSchemaInfoPostgres(globalDBConfig SchemaDBConfig, globalSchemaInfo *[]SchemaInfoTable, globalSchemaInfoDBID *sql.DB) int {
	var lRecordNum sql.NullString
	var lModuleName sql.NullString
	var lSubModuleName sql.NullString
	var lInstID sql.NullString
	var lInstSubID sql.NullString
	var lBinID sql.NullString
	var lBinSubID sql.NullString
	var lDBType sql.NullString
	var lDBUser sql.NullString
	var lDBPwd sql.NullString
	var lDBConnectionName sql.NullString
	var lDBTimeOut sql.NullInt64

	var lglobalSchemaInfo SchemaInfoTable

	tableName := fmt.Sprintf(`"%s"."%s"`, globalDBConfig.SchemaName, "SCHEMA_INFO")

	queryStmt := fmt.Sprintf(
		`SELECT 
			"RECORD_NUM",
			"MODULE_NAME", 
			"SUB_MODULE_NAME",
			"INST_ID", 
			"INST_SUBID", 
			"BIN_ID",
			"BIN_SUBID",
			"DB_TYPE", 
			"DB_USER", 
			"DB_PWD", 
			"DB_CONNSTR",
			"DB_TIMEOUT"
		FROM %s`, tableName)

	rows, err := globalSchemaInfoDBID.Query(queryStmt)

	if err != nil {
		// trace.Lg("db.Query() failed for queryStmt(%s) with err(%s)", queryStmt, err)
		return -1
	}
	defer rows.Close()
	////trace.Log(debugdef.DEBUG_LEVEL_TEST, "db.Query() success for queryStmt(%s)", queryStmt)
	rowcount := 0
	for rows.Next() {
		err := rows.Scan(
			&lRecordNum,
			&lModuleName,
			&lSubModuleName,
			&lInstID,
			&lInstSubID,
			&lBinID,
			&lBinSubID,
			&lDBType,
			&lDBUser,
			&lDBPwd,
			&lDBConnectionName,
			&lDBTimeOut,
		)
		if err != nil {
			// trace.Lg("rows.Scan() failed for queryStmt(%s) with err(%s)", queryStmt, err)
			return -1
		}
		lglobalSchemaInfo.RecordNum = lRecordNum.String
		lglobalSchemaInfo.ModuleName = lModuleName.String
		lglobalSchemaInfo.SubModuleName = lSubModuleName.String
		lglobalSchemaInfo.InstID = lInstID.String
		lglobalSchemaInfo.InstSubID = lInstSubID.String
		lglobalSchemaInfo.BinID = lBinID.String
		lglobalSchemaInfo.BinSubID = lBinSubID.String
		lglobalSchemaInfo.DBType = lDBType.String
		lglobalSchemaInfo.DBUser = lDBUser.String
		lglobalSchemaInfo.DBPwd = lDBPwd.String
		lglobalSchemaInfo.DBConnectionName = lDBConnectionName.String
		lglobalSchemaInfo.DBTimeOut = int(lDBTimeOut.Int64)
		lglobalSchemaInfo.SchemaName = globalDBConfig.SchemaName
		if readModuleFromSchemaPostgres(globalSchemaInfoDBID, lglobalSchemaInfo.RecordNum, lglobalSchemaInfo.SchemaName) >= 0 {
			*globalSchemaInfo = append(*globalSchemaInfo, lglobalSchemaInfo)
		}
		rowcount++
	}
	err = rows.Err()
	if err != nil {
		// trace.Lg("rows.Scan() failed for queryStmt(%s) with err(%s)", queryStmt, err)
		return -1
	}
	if rowcount == 0 {
		// trace.Lg("db.Query() success for queryStr(%s) returned zero rows", queryStmt)
		return -1
	}
	return rowcount
}
