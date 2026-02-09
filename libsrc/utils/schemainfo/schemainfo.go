package schemainfo

import (
	"context"
	"database/sql"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"lmsapieng/include/common/dbdef"
	"lmsapieng/include/common/globaldef"
	"lmsapieng/libsrc/utils/dbutil"
	"lmsapieng/libsrc/utils/rtutil"

	"os"
	"strings"

	_ "github.com/lib/pq"
)

type SchemaDBConfig struct {
	XMLName          xml.Name `xml:"global_dbconfig"`
	Text             string   `xml:",chardata"`
	DBType           string   `xml:"db_type"`
	DBUser           string   `xml:"db_user"`
	DBPasswd         string   `xml:"db_passwd"`
	DBConnectionName string   `xml:"db_conn_name"`
	SchemaName       string   `xml:"schema_name"`
}

type SchemaInfoTable struct {
	RecordNum        string
	ModuleName       string
	SubModuleName    string
	InstID           string
	InstSubID        string
	BinID            string
	BinSubID         string
	DBType           string
	DBUser           string
	DBPwd            string
	DBConnectionName string
	DBTimeOut        int
	DBID             *sql.DB
	LoadStatus       bool
	ConnectionStr    string
	DBDriverName     string
	SchemaName       string
}

var schemaDBConfig SchemaDBConfig
var loadSchemaDBConfigFlag = false
var globalSchemaInfo []SchemaInfoTable
var loadGlobalSchemaInfoFlag = false

func loadSchemaInfoConfig() int {
	if !globaldef.IsAppBaseDirExists() {
		//trace.Lg("IsAppBaseDirExists() failed")
		return -1
	}
	configFile := fmt.Sprintf("%s/config/schemainfo/schemainfo_config.xml", globaldef.GetAppBaseDir())
	xmlFile, err := os.Open(configFile)
	if err != nil {
		//trace.Lg("os.Open() failed for configFile(%s)...loadGlobalDBConfig() Failed", configFile)
		return -1
	}
	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)
	err = xml.Unmarshal(byteValue, &schemaDBConfig)
	if err != nil {
		//trace.Lg("xml.Unmarshal() failed for configFile(%s)...loadGlobalDBConfig() Failed", configFile)
		return -1
	}
	loadSchemaDBConfigFlag = true
	return 1
}

func getSchemaDBConfig(sDBConfig *SchemaDBConfig) int {
	if !loadSchemaDBConfigFlag {
		if loadSchemaInfoConfig() < 0 {
			return -1
		}
	}
	if len(schemaDBConfig.DBType) == 0 || len(schemaDBConfig.DBUser) == 0 || len(schemaDBConfig.DBPasswd) == 0 || len(schemaDBConfig.DBConnectionName) == 0 {
		return -1
	}
	*sDBConfig = schemaDBConfig
	return 1
}

func loadGlobalSchemaInfo() int {

	var globalDBConfig SchemaDBConfig
	var globalSchemaInfoTableCount, rVal int
	var err error
	var connStr string
	var clearDBUser, clearDBPasswd, clearDBConnectionName string
	var globalSchemaInfoDBID *sql.DB

	if getSchemaDBConfig(&globalDBConfig) < 0 {
		//trace.Lg("globalconfig() failed")
		return -1
	}
	// trace.Lg("GetGlobalDBConfig() success with dbType(%s) dbUser(%s) dbPasswd(%s) dbConnectionName(%s)", globalDBConfig.DBType, globalDBConfig.DBUser, globalDBConfig.DBPasswd, globalDBConfig.DBConnectionName)
	globalSchemaInfo = make([]SchemaInfoTable, 0)
	if rtutil.IsDBEncryptReq() {
		if len(globalDBConfig.DBUser) != 0 {
			rVal, clearDBUser = dbutil.DecryptDBCredential(globalDBConfig.DBUser)
			if rVal < 0 {
				//trace.Lg("DecryptDBCredential() failed for DBUser:%s", globalDBConfig.DBUser)
				return -1
			}
		}
		if len(globalDBConfig.DBPasswd) != 0 {
			rVal, clearDBPasswd = dbutil.DecryptDBCredential(globalDBConfig.DBPasswd)
			if rVal < 0 {
				//trace.Lg("DecryptDBCredential() failed for DBPasswd:%s", globalDBConfig.DBPasswd)
				return -1
			}
		}
		if len(globalDBConfig.DBConnectionName) != 0 && !strings.EqualFold(globalDBConfig.DBConnectionName, "Nil") {
			rVal, clearDBConnectionName = dbutil.DecryptDBCredential(globalDBConfig.DBConnectionName)
			if rVal < 0 {
				//trace.Lg("DecryptDBCredential() failed for DBConnectionName:%s", globalDBConfig.DBConnectionName)
				return -1
			}
		}
	}
	if strings.EqualFold(globalDBConfig.DBType, dbdef.DBTypeOracle) {
		if rtutil.IsDBEncryptReq() {
			if strings.EqualFold(globalDBConfig.DBConnectionName, "Nil") {
				connStr = fmt.Sprintf("%s/%s", clearDBUser, clearDBPasswd)
			} else {
				connStr = fmt.Sprintf("%s/%s@%s", clearDBUser, clearDBPasswd, clearDBConnectionName)
			}
		} else {
			if strings.EqualFold(globalDBConfig.DBConnectionName, "Nil") {
				connStr = fmt.Sprintf("%s/%s", globalDBConfig.DBUser, globalDBConfig.DBPasswd)
			} else {
				connStr = fmt.Sprintf("%s/%s@%s", globalDBConfig.DBUser, globalDBConfig.DBPasswd, globalDBConfig.DBConnectionName)
			}
		}
		globalSchemaInfoDBID, err = sql.Open("oci8", connStr)
		if err != nil {
			////trace.Log(debugdef.DEBUG_LEVEL_TEST, "sql.Open() failed connStr(%s) with err(%s) loadGlobalSchemaInfo() failed", connStr, err)
			return -1
		}
		globalSchemaInfoTableCount = loadGlobalSchemaInfoOracle(globalDBConfig, &globalSchemaInfo, globalSchemaInfoDBID)
	} else if strings.EqualFold(globalDBConfig.DBType, dbdef.DBTypePostgreSQL) {
		if rtutil.IsDBEncryptReq() {
			dbcredentials := strings.Split(clearDBConnectionName, "|")
			if len(dbcredentials) != 4 {
				//trace.Lg("globalDBConfig.DBConnectionName for PostgreSQL should be configured as (host|port|dbname|schema)")
				return -1
			}
			connStr = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
				dbcredentials[0], dbcredentials[1], clearDBUser, clearDBPasswd, dbcredentials[2])
			globalDBConfig.SchemaName = dbcredentials[3]
		} else {
			dbcredentials := strings.Split(globalDBConfig.DBConnectionName, "|")
			if len(dbcredentials) != 4 {
				//trace.Lg("globalDBConfig.DBConnectionName for PostgreSQL should be configured as (host|port|dbname|schema)")
				return -1
			}
			connStr = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
				dbcredentials[0], dbcredentials[1], globalDBConfig.DBUser, globalDBConfig.DBPasswd, dbcredentials[2])
			globalDBConfig.SchemaName = dbcredentials[3]
		}
		globalSchemaInfoDBID, err = sql.Open("postgres", connStr)
		if err != nil {
			////trace.Log(debugdef.DEBUG_LEVEL_TEST, "sql.Open() failed connStr(%s) with err(%s) loadGlobalSchemaInfo() failed", connStr, err)
			return -1
		}
		globalSchemaInfoTableCount = loadGlobalSchemaInfoPostgres(globalDBConfig, &globalSchemaInfo, globalSchemaInfoDBID)
	} else {
		//trace.Lg("invalid DBType(%s) loadGlobalSchemaInfo() failed", globalDBConfig.DBType)
		return -1
	}

	// trace.Lg("Closing GlobalSchemaDBID....")
	globalSchemaInfoDBID.Close()

	if globalSchemaInfoTableCount <= 0 {
		//trace.Lg("loadGlobalSchemaInfo() failed")
		return -1
	}

	// trace.Lg("loadGlobalSchemaInfo() success")

	for i := 0; i < len(globalSchemaInfo); i++ {
		//trace.Lg("loaded record_num(%s)->(%s/%s/%s/%s/%s/%s/%s/%s/%s/%s)",
		// 	globalSchemaInfo[i].RecordNum,
		// 	globalSchemaInfo[i].ModuleName,
		// 	globalSchemaInfo[i].SubModuleName,
		// 	globalSchemaInfo[i].InstID,
		// 	globalSchemaInfo[i].InstSubID,
		// 	globalSchemaInfo[i].BinID,
		// 	globalSchemaInfo[i].BinSubID,
		// 	globalSchemaInfo[i].DBType,
		// 	globalSchemaInfo[i].DBUser,
		// 	globalSchemaInfo[i].DBPwd,
		// 	globalSchemaInfo[i].DBConnectionName)
		globalSchemaInfo[i].LoadStatus = false

		if len(globalSchemaInfo[i].RecordNum) == 0 {
			//trace.Lg("RecordNum is NULL...")
			continue
		}
		if len(globalSchemaInfo[i].ModuleName) == 0 {
			//trace.Lg("ModuleName is NULL... for RecordNum(%s)", globalSchemaInfo[i].RecordNum)
			continue
		}
		if len(globalSchemaInfo[i].SubModuleName) == 0 {
			//trace.Lg("SubModuleName is NULL... for RecordNum(%s)", globalSchemaInfo[i].RecordNum)
			continue
		}
		if len(globalSchemaInfo[i].InstID) == 0 {
			//trace.Lg("InstID is NULL... for RecordNum(%s)", globalSchemaInfo[i].RecordNum)
			continue
		}
		if len(globalSchemaInfo[i].InstSubID) == 0 {
			//trace.Lg("InstSubID is NULL... for RecordNum(%s)", globalSchemaInfo[i].RecordNum)
			continue
		}
		if len(globalSchemaInfo[i].BinID) == 0 {
			//trace.Lg("BinID is NULL... for RecordNum(%s)", globalSchemaInfo[i].RecordNum)
			continue
		}
		if len(globalSchemaInfo[i].BinSubID) == 0 {
			//trace.Lg("BinSubID is NULL... for RecordNum(%s)", globalSchemaInfo[i].RecordNum)
			continue
		}
		if len(globalSchemaInfo[i].DBType) == 0 {
			//trace.Lg("DBType is NULL... for RecordNum(%s)", globalSchemaInfo[i].RecordNum)
			continue
		}
		if len(globalSchemaInfo[i].DBUser) == 0 {
			//trace.Lg("DBUser is NULL... for RecordNum(%s)", globalSchemaInfo[i].RecordNum)
			continue
		}
		if len(globalSchemaInfo[i].DBPwd) == 0 {
			//trace.Lg("DBPwd is NULL... for RecordNum(%s)", globalSchemaInfo[i].RecordNum)
			continue
		}
		if len(globalSchemaInfo[i].DBConnectionName) == 0 {
			//trace.Lg("DBConnectionName is NULL... for RecordNum(%s)", globalSchemaInfo[i].RecordNum)
			continue
		}

		if rtutil.IsDBEncryptReq() {
			if len(globalSchemaInfo[i].DBUser) != 0 && globalSchemaInfo[i].DBUser != "*" {
				rVal, clearDBUser = dbutil.DecryptDBCredential(globalSchemaInfo[i].DBUser)
				if rVal < 0 {
					//trace.Lg("DecryptDBCredential() failed for DBUser:%s", globalSchemaInfo[i].DBUser)
					return -1
				}
			}
			if len(globalSchemaInfo[i].DBPwd) != 0 && globalSchemaInfo[i].DBPwd != "*" {
				rVal, clearDBPasswd = dbutil.DecryptDBCredential(globalSchemaInfo[i].DBPwd)
				if rVal < 0 {
					////trace.Log(debugdef.DEBUG_LEVEL_TEST, "DecryptDBCredential() failed for DBPasswd:%s", globalSchemaInfo[i].DBPwd)
					return -1
				}
			}
			if len(globalSchemaInfo[i].DBConnectionName) != 0 && !strings.EqualFold(globalSchemaInfo[i].DBConnectionName, "Nil") {
				rVal, clearDBConnectionName = dbutil.DecryptDBCredential(globalSchemaInfo[i].DBConnectionName)
				if rVal < 0 {
					//trace.Lg("DecryptDBCredential() failed for DBConnectionName:%s", globalSchemaInfo[i].DBConnectionName)
					return -1
				}
			}
		}
		if strings.EqualFold(globalSchemaInfo[i].DBType, dbdef.DBTypeOracle) {
			if rtutil.IsDBEncryptReq() {
				if strings.EqualFold(clearDBConnectionName, "Nil") {
					connStr = fmt.Sprintf("%s/%s", clearDBUser, clearDBPasswd)
				} else {
					connStr = fmt.Sprintf("%s/%s@%s", clearDBUser, clearDBPasswd, clearDBConnectionName)
				}
			} else {
				if strings.EqualFold(globalSchemaInfo[i].DBConnectionName, "Nil") {
					connStr = fmt.Sprintf("%s/%s", globalSchemaInfo[i].DBUser, globalSchemaInfo[i].DBPwd)
				} else {
					connStr = fmt.Sprintf("%s/%s@%s", globalSchemaInfo[i].DBUser, globalSchemaInfo[i].DBPwd, globalSchemaInfo[i].DBConnectionName)
				}
			}
			globalSchemaInfo[i].DBID, err = sql.Open("oci8", connStr)
			if err != nil {
				////trace.Log(debugdef.DEBUG_LEVEL_TEST, "sql.Open() failed for connStr(%s) with err(%s) for RecordNum(%s)", connStr, err, globalSchemaInfo[i].RecordNum)
				continue
			}
			globalSchemaInfo[i].ConnectionStr = connStr
			globalSchemaInfo[i].DBDriverName = "oci8"
			// //trace.Log(debugdef.DEBUG_LEVEL_TEST, "sql.Open() success for connStr(%s) for RecordNum(%s)", connStr, globalSchemaInfo[i].RecordNum)
		} else if strings.EqualFold(globalSchemaInfo[i].DBType, dbdef.DBTypeMSSQL) {
			if rtutil.IsDBEncryptReq() {
				dbcredentials := strings.Split(clearDBConnectionName, "|")
				if len(dbcredentials) != 3 {
					//trace.Lg("globalDBConfig.DBConnectionName for MSSql should be configured as (server|port|database)")
					continue
				}
				connStr = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;",
					dbcredentials[0], clearDBUser, clearDBPasswd, dbcredentials[1], dbcredentials[2])
				////trace.Log(debugdef.DEBUG_LEVEL_TEST, "connStr(%s) for MSSql", connStr)
			} else {
				dbcredentials := strings.Split(globalSchemaInfo[i].DBConnectionName, "|")
				if len(dbcredentials) != 3 {
					//trace.Lg("globalDBConfig.DBConnectionName for MSSql should be configured as (server|port|database)")
					continue
				}
				connStr = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;",
					dbcredentials[0], globalSchemaInfo[i].DBUser, globalSchemaInfo[i].DBPwd, dbcredentials[1], dbcredentials[2])
				// trace.Lg( "connStr(%s) for MSSql", connStr)
			}
			globalSchemaInfo[i].DBID, err = sql.Open("mssql", connStr)
			if err != nil {
				// trace.Lg("sql.Open() failed for connStr(%s) with err(%s) for RecordNum(%s)", connStr, err, globalSchemaInfo[i].RecordNum)
				continue
			}
			globalSchemaInfo[i].ConnectionStr = connStr
			globalSchemaInfo[i].DBDriverName = "mssql"
			// //trace.Log(debugdef.DEBUG_LEVEL_TEST, "sql.Open() success for connStr(%s) for RecordNum(%s)", connStr, globalSchemaInfo[i].RecordNum)
		} else if strings.EqualFold(globalDBConfig.DBType, dbdef.DBTypePostgreSQL) {
			if rtutil.IsDBEncryptReq() {
				dbcredentials := strings.Split(clearDBConnectionName, "|")
				if len(dbcredentials) != 4 {
					//trace.Lg("globalDBConfig.DBConnectionName for PostgreSQL should be configured as (host|port|dbname|schema)")
					return -1
				}
				connStr = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
					dbcredentials[0], dbcredentials[1], clearDBUser, clearDBPasswd, dbcredentials[2])
				globalSchemaInfo[i].SchemaName = dbcredentials[3]
			} else {
				dbcredentials := strings.Split(globalSchemaInfo[i].DBConnectionName, "|")
				if len(dbcredentials) != 4 {
					//trace.Lg("globalDBConfig.DBConnectionName for PostgreSQL should be configured as (host|port|dbname|schema)")
					return -1
				}
				connStr = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
					dbcredentials[0], dbcredentials[1], globalSchemaInfo[i].DBUser, globalSchemaInfo[i].DBPwd, dbcredentials[2])
				globalSchemaInfo[i].SchemaName = dbcredentials[3]
			}
			globalSchemaInfo[i].DBID, err = sql.Open("postgres", connStr)
			if err != nil {
				// trace.Lg("sql.Open() failed connStr(%s) with err(%s) loadGlobalSchemaInfo() failed", connStr, err)
				return -1
			}
			globalSchemaInfo[i].ConnectionStr = connStr
			globalSchemaInfo[i].DBDriverName = "postgres"
			// //trace.Log(debugdef.DEBUG_LEVEL_TEST, "sql.Open() success for connStr(%s) for RecordNum(%s)", connStr, globalSchemaInfo[i].RecordNum)
		} else if strings.EqualFold(globalSchemaInfo[i].DBType, dbdef.DBTypeSQLite) {
			if rtutil.IsDBEncryptReq() {
				connStr = fmt.Sprintf("%s/db/%s", globaldef.GetAppBaseDir(), globalSchemaInfo[i].DBConnectionName)
			} else {
				connStr = fmt.Sprintf("%s/db/%s", globaldef.GetAppBaseDir(), clearDBConnectionName)
			}
			globalSchemaInfo[i].DBID, err = sql.Open("sqlite3", connStr)
			if err != nil {
				//	//trace.Log(debugdef.DEBUG_LEVEL_TEST, "sql.Open() failed for connStr(%s) with err(%s) for RecordNum(%s)", connStr, err, globalSchemaInfo[i].RecordNum)
				continue
			}
			globalSchemaInfo[i].ConnectionStr = connStr
			globalSchemaInfo[i].DBDriverName = "sqlite3"
			////trace.Log(debugdef.DEBUG_LEVEL_TEST, "sql.Open() success for connStr(%s) for RecordNum(%s)", connStr, globalSchemaInfo[i].RecordNum)
		} else {
			//trace.Lg("DbType(%s) is invalid for for RecordNum(%s)", globalSchemaInfo[i].DBType, globalSchemaInfo[i].RecordNum)
			continue
		}
		globalSchemaInfo[i].LoadStatus = true
	}
	return 1
}

func getDBContext(contextParams ...string) (int, *sql.DB, string, int, string, string, string, string) {

	if !loadGlobalSchemaInfoFlag {
		if loadGlobalSchemaInfo() < 0 {
			return -1, nil, "", 0, "", "", "", ""
		}
		loadGlobalSchemaInfoFlag = true
	}

	var lglobalSchemaInfo SchemaInfoTable
	var found bool
	lglobalSchemaInfo.ModuleName = "*"
	lglobalSchemaInfo.SubModuleName = "*"
	lglobalSchemaInfo.InstID = "*"
	lglobalSchemaInfo.BinID = "*"
	lglobalSchemaInfo.InstSubID = "*"
	lglobalSchemaInfo.BinSubID = "*"
	if len(contextParams) >= 1 {
		lglobalSchemaInfo.ModuleName = contextParams[0]
	}
	if len(contextParams) >= 2 {
		lglobalSchemaInfo.SubModuleName = contextParams[1]
	}
	if len(contextParams) >= 3 {
		lglobalSchemaInfo.InstID = contextParams[2]
	}
	if len(contextParams) >= 4 {
		lglobalSchemaInfo.BinID = contextParams[3]
	}
	if len(contextParams) >= 5 {
		lglobalSchemaInfo.InstSubID = contextParams[4]
	}
	if len(contextParams) >= 6 {
		lglobalSchemaInfo.BinSubID = contextParams[5]
	}
	found = false
	for i := 0; i < len(globalSchemaInfo); i++ {
		if globalSchemaInfo[i].LoadStatus &&
			globalSchemaInfo[i].ModuleName == lglobalSchemaInfo.ModuleName &&
			globalSchemaInfo[i].SubModuleName == lglobalSchemaInfo.SubModuleName &&
			globalSchemaInfo[i].InstID == lglobalSchemaInfo.InstID &&
			globalSchemaInfo[i].BinID == lglobalSchemaInfo.BinID &&
			globalSchemaInfo[i].InstSubID == lglobalSchemaInfo.InstSubID &&
			globalSchemaInfo[i].BinSubID == lglobalSchemaInfo.BinSubID {
			lglobalSchemaInfo.DBID = globalSchemaInfo[i].DBID
			lglobalSchemaInfo.DBType = globalSchemaInfo[i].DBType
			lglobalSchemaInfo.DBTimeOut = globalSchemaInfo[i].DBTimeOut
			lglobalSchemaInfo.ConnectionStr = globalSchemaInfo[i].ConnectionStr
			lglobalSchemaInfo.DBDriverName = globalSchemaInfo[i].DBDriverName
			if len(globalSchemaInfo[i].SchemaName) == 0 {
				globalSchemaInfo[i].SchemaName = globaldef.NOT_INITIALIZED
			}
			lglobalSchemaInfo.SchemaName = globalSchemaInfo[i].SchemaName
			lglobalSchemaInfo.RecordNum = globalSchemaInfo[i].RecordNum
			found = true
			break
		}
	}
	if found {
		return 1, lglobalSchemaInfo.DBID, lglobalSchemaInfo.DBType, lglobalSchemaInfo.DBTimeOut, lglobalSchemaInfo.ConnectionStr, lglobalSchemaInfo.DBDriverName, lglobalSchemaInfo.SchemaName, lglobalSchemaInfo.RecordNum
	}
	if loadGlobalSchemaInfo() < 0 {
		return -1, nil, "", 0, "", "", "", ""
	}
	found = false
	for i := 0; i < len(globalSchemaInfo); i++ {
		if globalSchemaInfo[i].LoadStatus &&
			globalSchemaInfo[i].ModuleName == lglobalSchemaInfo.ModuleName &&
			globalSchemaInfo[i].SubModuleName == lglobalSchemaInfo.SubModuleName &&
			globalSchemaInfo[i].InstID == lglobalSchemaInfo.InstID &&
			globalSchemaInfo[i].BinID == lglobalSchemaInfo.BinID &&
			globalSchemaInfo[i].InstSubID == lglobalSchemaInfo.InstSubID &&
			globalSchemaInfo[i].BinSubID == lglobalSchemaInfo.BinSubID {
			lglobalSchemaInfo.DBID = globalSchemaInfo[i].DBID
			lglobalSchemaInfo.DBType = globalSchemaInfo[i].DBType
			lglobalSchemaInfo.DBTimeOut = globalSchemaInfo[i].DBTimeOut
			lglobalSchemaInfo.ConnectionStr = globalSchemaInfo[i].ConnectionStr
			lglobalSchemaInfo.DBDriverName = globalSchemaInfo[i].DBDriverName
			if len(globalSchemaInfo[i].SchemaName) == 0 {
				globalSchemaInfo[i].SchemaName = globaldef.NOT_INITIALIZED
			}
			lglobalSchemaInfo.SchemaName = globalSchemaInfo[i].SchemaName
			lglobalSchemaInfo.RecordNum = globalSchemaInfo[i].RecordNum
			found = true
			break
		}
	}
	if found {
		return 1, lglobalSchemaInfo.DBID, lglobalSchemaInfo.DBType, lglobalSchemaInfo.DBTimeOut, lglobalSchemaInfo.ConnectionStr, lglobalSchemaInfo.DBDriverName, lglobalSchemaInfo.SchemaName, lglobalSchemaInfo.RecordNum
	}
	return -1, nil, "", 0, "", "", "", ""
}

func SelectDBContext(resultDBContext *dbdef.DBContextDef, rejectReason *string, contextParams ...string) int {
	retval, DBID, DBType, DBTimeOut, ConnectionStr, DBDriverName, SchemaName, RecordNum := getDBContext(contextParams...)
	if retval < 0 {
		//trace.Lg("getDBContext() Failed for %s/%s/%s/%s/%s/%s", resultDBContext.ModuleName, resultDBContext.SubModuleName, resultDBContext.InstID, resultDBContext.BinID, resultDBContext.InstSubID, resultDBContext.BinSubID)
		*rejectReason = fmt.Sprintf("getDBContext Failed for %s/%s/%s/%s/%s/%s", resultDBContext.ModuleName, resultDBContext.SubModuleName, resultDBContext.InstID, resultDBContext.BinID, resultDBContext.InstSubID, resultDBContext.BinSubID)
		return -1
	}
	//trace.Lg("getDBContext() success with DBType<%s> DBTimeOut<%d>", DBType, DBTimeOut)
	resultDBContext.DBType = DBType
	resultDBContext.DBID = DBID
	resultDBContext.DBTimeOut = DBTimeOut
	resultDBContext.ConnectionStr = ConnectionStr
	resultDBContext.DBDriverName = DBDriverName
	resultDBContext.SchemaName = SchemaName
	resultDBContext.RecordNum = RecordNum
	resultDBContext.DBTxFlag = false
	return 1
}

func SetTxDBContext(dbContext *dbdef.DBContextDef, rejectReason *string) int {
	var err error
	dbContext.DBTx, err = dbContext.DBID.Begin()
	if err != nil {
		//trace.Lg("DBID.Begin() Failed with err(%s)...SetTxDBContext() Failed", err)
		*rejectReason = fmt.Sprintf("DBID.Begin() Failed with err(%s)...SetTxDBContext() Failed", err)
		return -1
	}
	dbContext.DBTxFlag = true
	return 1
}

func CommitDBContext(dbContext *dbdef.DBContextDef) {
	if dbContext.DBTxFlag {
		dbContext.DBTx.Commit()
		dbContext.DBTxFlag = false
	}
}

func RollbackDBContext(dbContext *dbdef.DBContextDef) {
	if dbContext.DBTxFlag {
		dbContext.DBTx.Rollback()
		dbContext.DBTxFlag = false
	}
}

func GetActiveDBContext(dbcontextmap map[string]interface{}, resultDBContext *dbdef.DBContextDef, rejectReason *string, contextParams ...string) int {

	var ldbContextDef dbdef.DBContextDef
	var found bool
	activeContextArray := dbcontextmap[dbdef.DBContextArrayObj].([]dbdef.DBContextDef)
	ldbContextDef.ModuleName = "*"
	ldbContextDef.SubModuleName = "*"
	ldbContextDef.InstID = "*"
	ldbContextDef.BinID = "*"
	ldbContextDef.InstSubID = "*"
	ldbContextDef.BinSubID = "*"
	if len(contextParams) >= 1 {
		ldbContextDef.ModuleName = contextParams[0]
	}
	if len(contextParams) >= 2 {
		ldbContextDef.SubModuleName = contextParams[1]
	}
	if len(contextParams) >= 3 {
		ldbContextDef.InstID = contextParams[2]
	}
	if len(contextParams) >= 4 {
		ldbContextDef.BinID = contextParams[3]
	}
	if len(contextParams) >= 5 {
		ldbContextDef.InstSubID = contextParams[4]
	}
	if len(contextParams) >= 6 {
		ldbContextDef.BinSubID = contextParams[5]
	}
	found = false
	for i := 0; i < len(activeContextArray); i++ {
		if ldbContextDef.ModuleName == activeContextArray[i].ModuleName &&
			ldbContextDef.SubModuleName == activeContextArray[i].SubModuleName &&
			ldbContextDef.InstID == activeContextArray[i].InstID &&
			ldbContextDef.BinID == activeContextArray[i].BinID &&
			ldbContextDef.InstSubID == activeContextArray[i].InstSubID &&
			ldbContextDef.BinSubID == activeContextArray[i].BinSubID {
			*resultDBContext = activeContextArray[i]
			found = true
			break
		}
	}
	if found {
		return 1
	}
	retval, DBID, DBType, DBTimeOut, ConnectionStr, DBDriverName, SchemaName, RecordNum := getDBContext(contextParams...)
	if retval < 0 {
		//trace.Lg("GetDBContext() Failed for %s/%s/%s/%s/%s/%s", ldbContextDef.ModuleName, ldbContextDef.SubModuleName, ldbContextDef.InstID, ldbContextDef.BinID, ldbContextDef.InstSubID, ldbContextDef.BinSubID)
		*rejectReason = fmt.Sprintf("GetDBContext Failed for %s/%s/%s/%s/%s/%s", ldbContextDef.ModuleName, ldbContextDef.SubModuleName, ldbContextDef.InstID, ldbContextDef.BinID, ldbContextDef.InstSubID, ldbContextDef.BinSubID)
		return -1
	}
	ldbContextDef.DBType = DBType
	ldbContextDef.DBID = DBID
	ldbContextDef.DBTimeOut = DBTimeOut
	ldbContextDef.ConnectionStr = ConnectionStr
	ldbContextDef.DBDriverName = DBDriverName
	ldbContextDef.SchemaName = SchemaName
	ldbContextDef.RecordNum = RecordNum
	ldbContextDef.DBTxFlag = false
	ldbContextDef.HTTPReqCtx = dbcontextmap[globaldef.ReqCtxObj].(context.Context)
	*resultDBContext = ldbContextDef
	activeContextArray = append(activeContextArray, ldbContextDef)
	dbcontextmap[dbdef.DBContextArrayObj] = activeContextArray
	return 1
}

func startActiveDBContext(dbcontextmap map[string]interface{}, dbContext *dbdef.DBContextDef, rejectReason *string) int {
	var err error
	activeContextArray := dbcontextmap[dbdef.DBContextArrayObj].([]dbdef.DBContextDef)
	dbContext.DBTx, err = dbContext.DBID.Begin()
	if err != nil {
		//trace.Lg("DBID.Begin() Failed with err(%s)...StartActiveDBContext() Failed", err)
		*rejectReason = fmt.Sprintf("DBID.Begin() Failed with err(%s)...StartActiveDBContext() Failed", err)
		return -1
	}
	dbContext.DBTxFlag = true
	found := false
	for i := 0; i < len(activeContextArray); i++ {
		if activeContextArray[i].InstID == dbContext.InstID &&
			activeContextArray[i].ModuleName == dbContext.ModuleName &&
			activeContextArray[i].BinID == dbContext.BinID {
			activeContextArray[i].DBTx = dbContext.DBTx
			activeContextArray[i].DBTxFlag = dbContext.DBTxFlag
			found = true
			break
		}
	}
	if !found {
		dbContext.DBTx.Commit()
		//trace.Lg("Cannot Search dbContext in activeContextArray...StartActiveDBContext() failed")
		*rejectReason = "Cannot Search dbContext in activeContextArray...StartActiveDBContext() failed"
		return -1
	}
	dbcontextmap[dbdef.DBContextArrayObj] = activeContextArray
	return 1
}

func GetActiveDBContextWithTxn(dbcontextmap map[string]interface{}, resultDBContext *dbdef.DBContextDef, rejectReason *string, contextParams ...string) int {
	var err error
	var ldbContextDef dbdef.DBContextDef
	var found bool
	activeContextArray := dbcontextmap[dbdef.DBContextArrayObj].([]dbdef.DBContextDef)
	ldbContextDef.ModuleName = "*"
	ldbContextDef.SubModuleName = "*"
	ldbContextDef.InstID = "*"
	ldbContextDef.BinID = "*"
	ldbContextDef.InstSubID = "*"
	ldbContextDef.BinSubID = "*"
	if len(contextParams) >= 1 {
		ldbContextDef.ModuleName = contextParams[0]
	}
	if len(contextParams) >= 2 {
		ldbContextDef.SubModuleName = contextParams[1]
	}
	if len(contextParams) >= 3 {
		ldbContextDef.InstID = contextParams[2]
	}
	if len(contextParams) >= 4 {
		ldbContextDef.BinID = contextParams[3]
	}
	if len(contextParams) >= 5 {
		ldbContextDef.InstSubID = contextParams[4]
	}
	if len(contextParams) >= 6 {
		ldbContextDef.BinSubID = contextParams[5]
	}
	found = false
	for i := 0; i < len(activeContextArray); i++ {
		if ldbContextDef.ModuleName == activeContextArray[i].ModuleName &&
			ldbContextDef.SubModuleName == activeContextArray[i].SubModuleName &&
			ldbContextDef.InstID == activeContextArray[i].InstID &&
			ldbContextDef.BinID == activeContextArray[i].BinID &&
			ldbContextDef.InstSubID == activeContextArray[i].InstSubID &&
			ldbContextDef.BinSubID == activeContextArray[i].BinSubID {
			if !activeContextArray[i].DBTxFlag {
				activeContextArray[i].DBTx, err = activeContextArray[i].DBID.Begin()
				if err != nil {
					//trace.Lg("DBID.Begin() Failed with err(%s)...GetActiveDBContextWithTxn() Failed", err)
					*rejectReason = fmt.Sprintf("DBID.Begin() Failed with err(%s)...GetActiveDBContextWithTxn() Failed", err)
					return -1
				}
				activeContextArray[i].DBTxFlag = true
			}
			*resultDBContext = activeContextArray[i]
			found = true
			break
		}
	}
	if found {
		return 1
	}
	retval, DBID, DBType, DBTimeOut, ConnectionStr, DBDriverName, SchemaName, RecordNum := getDBContext(contextParams...)
	if retval < 0 {
		//trace.Lg("GetDBContext() Failed for %s/%s/%s/%s/%s/%s", ldbContextDef.ModuleName, ldbContextDef.SubModuleName, ldbContextDef.InstID, ldbContextDef.BinID, ldbContextDef.InstSubID, ldbContextDef.BinSubID)
		*rejectReason = fmt.Sprintf("GetDBContext Failed for %s/%s/%s/%s/%s/%s", ldbContextDef.ModuleName, ldbContextDef.SubModuleName, ldbContextDef.InstID, ldbContextDef.BinID, ldbContextDef.InstSubID, ldbContextDef.BinSubID)
		return -1
	}
	ldbContextDef.DBType = DBType
	ldbContextDef.DBID = DBID
	ldbContextDef.DBTimeOut = DBTimeOut
	ldbContextDef.ConnectionStr = ConnectionStr
	ldbContextDef.DBDriverName = DBDriverName
	ldbContextDef.SchemaName = SchemaName
	ldbContextDef.RecordNum = RecordNum
	ldbContextDef.DBTxFlag = false
	ldbContextDef.DBTx, err = ldbContextDef.DBID.Begin()
	if err != nil {
		//trace.Lg("DBID.Begin() Failed with err(%s)...GetActiveDBContextWithTxn() Failed", err)
		*rejectReason = fmt.Sprintf("DBID.Begin() Failed with err(%s)...GetActiveDBContextWithTxn() Failed", err)
		return -1
	}
	ldbContextDef.DBTxFlag = true
	ldbContextDef.HTTPReqCtx = dbcontextmap[globaldef.ReqCtxObj].(context.Context)
	*resultDBContext = ldbContextDef
	activeContextArray = append(activeContextArray, ldbContextDef)
	dbcontextmap[dbdef.DBContextArrayObj] = activeContextArray
	return 1
}

func CommitSpecificActiveDBContext(dbcontext dbdef.DBContextDef, dbcontextmap map[string]interface{}) {
	activeContextArray := dbcontextmap[dbdef.DBContextArrayObj].([]dbdef.DBContextDef)
	for i := 0; i < len(activeContextArray); i++ {
		if dbcontext.ModuleName == activeContextArray[i].ModuleName &&
			dbcontext.SubModuleName == activeContextArray[i].SubModuleName &&
			dbcontext.InstID == activeContextArray[i].InstID &&
			dbcontext.BinID == activeContextArray[i].BinID &&
			dbcontext.InstSubID == activeContextArray[i].InstSubID &&
			dbcontext.BinSubID == activeContextArray[i].BinSubID {
			if activeContextArray[i].DBTxFlag {
				//trace.Lg("CommitSpecificActiveDBContext() called for Module<%s>", activeContextArray[i].ModuleName)
				activeContextArray[i].DBTx.Commit()
				activeContextArray[i].DBTxFlag = false
				break
			}
		}
	}
	dbcontextmap[dbdef.DBContextArrayObj] = activeContextArray
}

func CommitActiveDBContext(dbcontextmap map[string]interface{}) {
	activeContextArray := dbcontextmap[dbdef.DBContextArrayObj].([]dbdef.DBContextDef)
	for i := 0; i < len(activeContextArray); i++ {
		if activeContextArray[i].DBTxFlag {
			// trace.Lg("CommitActiveDBContext() called for Module<%s>", activeContextArray[i].ModuleName)
			activeContextArray[i].DBTx.Commit()
			activeContextArray[i].DBTxFlag = false
		}
	}
	dbcontextmap[dbdef.DBContextArrayObj] = activeContextArray
}

func RollbackSpecificActiveDBContext(dbcontext dbdef.DBContextDef, dbcontextmap map[string]interface{}) {
	activeContextArray := dbcontextmap[dbdef.DBContextArrayObj].([]dbdef.DBContextDef)
	for i := 0; i < len(activeContextArray); i++ {
		if dbcontext.ModuleName == activeContextArray[i].ModuleName &&
			dbcontext.SubModuleName == activeContextArray[i].SubModuleName &&
			dbcontext.InstID == activeContextArray[i].InstID &&
			dbcontext.BinID == activeContextArray[i].BinID &&
			dbcontext.InstSubID == activeContextArray[i].InstSubID &&
			dbcontext.BinSubID == activeContextArray[i].BinSubID {
			if activeContextArray[i].DBTxFlag {
				//trace.Lg("CommitSpecificActiveDBContext() called for Module<%s>", activeContextArray[i].ModuleName)
				activeContextArray[i].DBTx.Rollback()
				activeContextArray[i].DBTxFlag = false
				break
			}
		}
	}
	dbcontextmap[dbdef.DBContextArrayObj] = activeContextArray
}

func RollbackActiveDBContext(dbcontextmap map[string]interface{}) {
	activeContextArray := dbcontextmap[dbdef.DBContextArrayObj].([]dbdef.DBContextDef)
	for i := 0; i < len(activeContextArray); i++ {
		if activeContextArray[i].DBTxFlag {
			//trace.Lg("RollbackActiveDBContext() called for Module<%s>", activeContextArray[i].ModuleName)
			activeContextArray[i].DBTx.Rollback()
			activeContextArray[i].DBTxFlag = false
		}
	}
	dbcontextmap[dbdef.DBContextArrayObj] = activeContextArray
}

func PrintActiveDBContext(dbcontextmap map[string]interface{}) {
	activeContextArray := dbcontextmap[dbdef.DBContextArrayObj].([]dbdef.DBContextDef)
	//trace.Lg("PrintActiveDBContext() called with <%d> active dbcontexts", len(activeContextArray))
	for i := 0; i < len(activeContextArray); i++ {
		//trace.Lg("%s#%s#%s#%s#%s#%s->txflag<%t>", activeContextArray[i].ModuleName, activeContextArray[i].SubModuleName, activeContextArray[i].InstID, activeContextArray[i].BinID, activeContextArray[i].InstSubID, activeContextArray[i].BinSubID, activeContextArray[i].DBTxFlag)
	}
}

func IsActiveDBContextOpened(dbcontextmap map[string]interface{}) bool {
	found := false
	activeContextArray := dbcontextmap[dbdef.DBContextArrayObj].([]dbdef.DBContextDef)
	for i := 0; i < len(activeContextArray); i++ {
		if activeContextArray[i].DBTxFlag {
			found = true
			break
		}
	}
	return found
}
