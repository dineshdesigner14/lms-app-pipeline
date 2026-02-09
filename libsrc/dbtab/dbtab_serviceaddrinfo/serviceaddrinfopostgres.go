package dbtab_serviceaddrinfo

import (
	"context"
	"database/sql"
	"fmt"
	"lmsapieng/include/common/dbdef"
	"lmsapieng/include/dbtabdef"
	"time"
)

func Load_1_ServiceAddrInfoTablePostgres(serviceModule string, ServiceAddrInfoPtr *[]dbtabdef.ServiceAddrInfoTable, dbContext dbdef.DBContextDef, dberr *string) int {
	var err error
	var rows *sql.Rows
	var lServiceID sql.NullString
	var lServiceModule sql.NullString
	var lServiceStatus sql.NullString
	var lServiceIpAddr sql.NullString
	var lServicePort sql.NullInt64
	var lServiceTimeout sql.NullInt64

	var lServiceAddrInfo dbtabdef.ServiceAddrInfoTable

	timeoutCtx, cancelFunc := context.WithTimeout(dbContext.HTTPReqCtx, time.Second*time.Duration(dbContext.DBTimeOut))
	defer cancelFunc()

	tableName := fmt.Sprintf(`"%s"."%s"`, dbContext.SchemaName, "SERVICE_ADDR_INFO")

	queryStmt :=
		fmt.Sprintf(`SELECT 
			"SERVICE_ID",
			"SERVICE_MODULE",
			"SERVICE_STATUS",
			"SERVICE_IP",
			"SERVICE_PORT",
			"SERVICE_TIMEOUT"
		FROM %s WHERE "SERVICE_MODULE"='%s'`, tableName, serviceModule)

	if dbContext.DBTxFlag {
		rows, err = dbContext.DBTx.QueryContext(timeoutCtx, queryStmt)
	} else {
		rows, err = dbContext.DBID.QueryContext(timeoutCtx, queryStmt)
	}
	if err != nil {
		//trace.Lg("db.Query() failed for queryStmt(%s) with err(%s)", queryStmt, err)
		*dberr = err.Error()
		return -1
	}
	defer rows.Close()
	rowcount := 0
	for rows.Next() {
		err := rows.Scan(
			&lServiceID,
			&lServiceModule,
			&lServiceStatus,
			&lServiceIpAddr,
			&lServicePort,
			&lServiceTimeout,
		)
		if err != nil {
			//trace.Lg("rows.Scan() failed for queryStmt(%s) with err(%s)", queryStmt, err)
			*dberr = err.Error()
			return -1
		}
		lServiceAddrInfo.ServiceID = lServiceID.String
		lServiceAddrInfo.ServiceModule = lServiceModule.String
		lServiceAddrInfo.ServiceStatus = lServiceStatus.String
		lServiceAddrInfo.ServiceIpAddr = lServiceIpAddr.String
		lServiceAddrInfo.ServicePort = int(lServicePort.Int64)
		lServiceAddrInfo.ServiceTimeout = int(lServiceTimeout.Int64)
		*ServiceAddrInfoPtr = append(*ServiceAddrInfoPtr, lServiceAddrInfo)
		rowcount++
	}
	if rowcount == 0 {
		*dberr = "sql: no rows in result set"
		return -1
	}
	return rowcount
}
