package dbtab_tableinfo

import (
	"context"
	"database/sql"
	"fmt"
	"lmsapieng/include/common/dbdef"
	"strings"
	"time"
)

func LoadFromDBTableOracle(queryStr string, resultMapArray *[]map[string]interface{}, dbContext dbdef.DBContextDef, dberr *string) int {
	var err error
	var rows *sql.Rows
	var cols []string
	timeoutCtx, cancelFunc := context.WithTimeout(dbContext.HTTPReqCtx, time.Second*time.Duration(dbContext.DBTimeOut))
	defer cancelFunc()
	if dbContext.DBTxFlag {
		rows, err = dbContext.DBTx.QueryContext(timeoutCtx, queryStr)
	} else {
		rows, err = dbContext.DBID.QueryContext(timeoutCtx, queryStr)
	}
	if err != nil {
		//trace.Lg("db.Query() failed for queryStmt(%s) with err(%s)", queryStr, err)
		*dberr = err.Error()
		return -1
	}
	cols, err = rows.Columns()
	if err != nil {
		//trace.Lg("rows.Columns() failed for queryStmt(%s) with err(%s)", queryStr, err)
		*dberr = err.Error()
		return -1
	}
	defer rows.Close()
	rowcount := 0
	for rows.Next() {
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}
		if err = rows.Scan(columnPointers...); err != nil {
			//trace.Lg("rows.Scan() failed for queryStmt(%s) with err(%s)", queryStr, err)
			*dberr = err.Error()
			return -1
		}
		resultMap := make(map[string]interface{})
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			resultMapColumnName := strings.ToLower(colName)
			if dateVal, ok := (*val).(time.Time); ok {
				resultMap[resultMapColumnName] = dateVal.Format("02012006")
			} else {
				resultMap[resultMapColumnName] = *val
			}
		}
		*resultMapArray = append(*resultMapArray, resultMap)
		rowcount++
	}
	return rowcount
}

func ReadFromDBTableOracle(queryStr string, resultMap map[string]interface{}, dbContext dbdef.DBContextDef, dberr *string) int {
	var err error
	var rows *sql.Rows
	var cols []string
	timeoutCtx, cancelFunc := context.WithTimeout(dbContext.HTTPReqCtx, time.Second*time.Duration(dbContext.DBTimeOut))
	defer cancelFunc()
	if dbContext.DBTxFlag {
		rows, err = dbContext.DBTx.QueryContext(timeoutCtx, queryStr)
	} else {
		rows, err = dbContext.DBID.QueryContext(timeoutCtx, queryStr)
	}
	if err != nil {
		//trace.Lg("db.Query() failed for queryStmt(%s) with err(%s)", queryStr, err)
		*dberr = err.Error()
		return -1
	}
	cols, err = rows.Columns()
	if err != nil {
		//trace.Lg("rows.Columns() failed for queryStmt(%s) with err(%s)", queryStr, err)
		*dberr = err.Error()
		return -1
	}
	defer rows.Close()
	rowcount := 0
	for rows.Next() {
		if rowcount == 0 {
			columns := make([]interface{}, len(cols))
			columnPointers := make([]interface{}, len(cols))
			for i, _ := range columns {
				columnPointers[i] = &columns[i]
			}
			if err = rows.Scan(columnPointers...); err != nil {
				//trace.Lg("rows.Scan() failed for queryStmt(%s) with err(%s)", queryStr, err)
				*dberr = err.Error()
				return -1
			}
			for i, colName := range cols {
				val := columnPointers[i].(*interface{})
				resultMapColumnName := strings.ToLower(colName)
				if dateVal, ok := (*val).(time.Time); ok {
					resultMap[resultMapColumnName] = dateVal.Format("02012006")
				} else {
					resultMap[resultMapColumnName] = *val
				}
			}
		}
		rowcount++
	}
	if rowcount == 0 {
		*dberr = "sql: no rows in result set"
		return -1
	}
	if rowcount > 1 {
		*dberr = "sql: duplicate rows in result set"
		return -1
	}
	return rowcount
}

func placeholders(n int) string {
	if n == 0 {
		return ""
	}
	namedPlaceholders := make([]string, n)
	for i := 1; i <= n; i++ {
		namedPlaceholders[i-1] = fmt.Sprintf(":%d", i)
	}
	return strings.Join(namedPlaceholders, ",")
}

func InsertTableOracle(tableName string, queryStr string, dbContext dbdef.DBContextDef, dberr *string) int {
	var err error
	timeoutCtx, cancelFunc := context.WithTimeout(dbContext.HTTPReqCtx, time.Second*time.Duration(dbContext.DBTimeOut))
	defer cancelFunc()

	execStmt := fmt.Sprintf("INSERT INTO %s %s", tableName, queryStr)
	_, err = dbContext.DBTx.ExecContext(timeoutCtx, execStmt)
	if err != nil {
		//trace.Lg("Exec() for (%s) failed with err(%s)", execStmt, err)
		*dberr = err.Error()
		return -1
	}
	return 1
}

func UpdateTableOracle(tableName string, queryStr string, dbContext dbdef.DBContextDef, dberr *string) int {
	var err error

	timeoutCtx, cancelFunc := context.WithTimeout(dbContext.HTTPReqCtx, time.Second*time.Duration(dbContext.DBTimeOut))
	defer cancelFunc()

	exectStmt := fmt.Sprintf("UPDATE %s SET %s", tableName, queryStr)
	res, err := dbContext.DBTx.ExecContext(timeoutCtx, exectStmt)
	if err != nil {
		//trace.Lg("Exec() for (%s) failed with err(%s)", exectStmt, err)
		*dberr = err.Error()
		return -1
	}
	count, err := res.RowsAffected()
	if err != nil {
		//trace.Lg("Exec() for (%s) failed with err(%s)", exectStmt, err)
		*dberr = err.Error()
		return -1
	}
	if count == 0 {
		*dberr = "No Rows Updated"
		return -1
	}
	// trace.Lg("Exec() for (%s) Success with count(%d)", exectStmt, count)
	return 1
}

func DeleteTableOracle(tableName string, queryStr string, dbContext dbdef.DBContextDef, dberr *string) int {
	var err error

	timeoutCtx, cancelFunc := context.WithTimeout(dbContext.HTTPReqCtx, time.Second*time.Duration(dbContext.DBTimeOut))
	defer cancelFunc()
	exectStmt := fmt.Sprintf("DELETE FROM %s WHERE %s", tableName, queryStr)
	_, err = dbContext.DBTx.ExecContext(timeoutCtx, exectStmt)
	if err != nil {
		//trace.Lg("Exec() for (%s) failed with err(%s)", queryStr, err)
		*dberr = err.Error()
		return -1
	}
	return 1
}
