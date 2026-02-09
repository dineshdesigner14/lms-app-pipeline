package schemainfo

import (
	"database/sql"
	"fmt"
)

func loadGlobalSchemaInfoMSSql(globalDBConfig SchemaDBConfig, globalSchemaInfo *[]SchemaInfoTable, globalSchemaInfoDBID *sql.DB) int {
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

	var lglobalSchemaInfo SchemaInfoTable

	queryStmt := fmt.Sprintf(
		`SELECT 
			RECORD_NUM,
			MODULE_NAME, 
			SUB_MODULE_NAME,
			INST_ID, 
			INST_SUBID, 
			BIN_ID,
			BIN_SUBID,
			DB_TYPE, 
			DB_USER, 
			DB_PWD, 
			DB_CONNSTR, 
			DB_TIMEOUT
		FROM SCHEMA_INFO`)

	rows, err := globalSchemaInfoDBID.Query(queryStmt)

	if err != nil {
		//trace.Log(debugdef.DEBUG_LEVEL_TEST, "db.Query() failed for queryStmt(%s) with err(%s)", queryStmt, err)
		return -1
	}
	defer rows.Close()
	//trace.Log(debugdef.DEBUG_LEVEL_TEST, "db.Query() success for queryStmt(%s)", queryStmt)
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
		)
		if err != nil {
			//trace.Log(debugdef.DEBUG_LEVEL_TEST, "rows.Scan() failed for queryStmt(%s) with err(%s)", queryStmt, err)
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
		*globalSchemaInfo = append(*globalSchemaInfo, lglobalSchemaInfo)
		rowcount++
	}
	err = rows.Err()
	if err != nil {
		//trace.Log(debugdef.DEBUG_LEVEL_TEST, "rows.Scan() failed for queryStmt(%s) with err(%s)", queryStmt, err)
		return -1
	}
	if rowcount == 0 {
		//trace.Log(debugdef.DEBUG_LEVEL_TEST, "db.Query() success for queryStr(%s) returned zero rows", queryStmt)
		return -1
	}
	return rowcount
}
