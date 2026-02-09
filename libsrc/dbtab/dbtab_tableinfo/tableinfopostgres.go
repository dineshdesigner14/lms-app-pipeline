package dbtab_tableinfo

import (
	"context"
	"database/sql"
	"encoding/base64"
	"fmt"
	"lmsapieng/include/common/dbdef"
	"strconv"
	"strings"
	"time"
)

func LoadFromDBTablePostgres(queryStr string, resultMapArray *[]map[string]interface{}, dbContext dbdef.DBContextDef, dberr *string) int {
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
	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		//trace.Lg("rows.ColumnTypes() failed with err(%s)", err)
		*dberr = err.Error()
		return -1
	}
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
			columnType := columnTypes[i].DatabaseTypeName()
			//trace.Lg("columnName[%s] columnType[%s]", resultMapColumnName, columnType)
			switch columnType {
			case "VARCHAR", "TEXT":
				switch v := (*val).(type) {
				case []byte:
					resultMap[resultMapColumnName] = string(v)
				case string:
					resultMap[resultMapColumnName] = v
				default:
					//trace.Lg("Unexpected type for VARCHAR/TEXT value: %T", v)
					resultMap[resultMapColumnName] = fmt.Sprintf("%v", v)
				}
			case "NUMERIC":
				// Check for precision and scale
				_, scale, _ := columnTypes[i].DecimalSize()
				if scale == 0 {
					// Integer value
					switch v := (*val).(type) {
					case int64:
						resultMap[resultMapColumnName] = v
					case []byte:
						if numVal, err := strconv.ParseInt(string(v), 10, 64); err == nil {
							resultMap[resultMapColumnName] = numVal
						} else {
							//trace.Lg("Failed to parse NUMERIC value as int64: %s", err)
							resultMap[resultMapColumnName] = string(v)
						}
					default:
						//trace.Lg("Unexpected type for NUMERIC value: %T", v)
						resultMap[resultMapColumnName] = fmt.Sprintf("%v", v)
					}
				} else {
					// Handling for numeric with scale (float)
					switch v := (*val).(type) {
					case float64:
						resultMap[resultMapColumnName] = v
					case []byte:
						if numVal, err := strconv.ParseFloat(string(v), 64); err == nil {
							resultMap[resultMapColumnName] = numVal
						} else {
							//trace.Lg("Failed to parse NUMERIC value as float64: %s", err)
							resultMap[resultMapColumnName] = string(v)
						}
					default:
						//trace.Lg("Unexpected type for NUMERIC value: %T", v)
						resultMap[resultMapColumnName] = fmt.Sprintf("%v", v)
					}
				}
			case "INT4", "INT8":
				switch v := (*val).(type) {
				case int64:
					resultMap[resultMapColumnName] = v
				case []byte:
					if numVal, err := strconv.ParseInt(string(v), 10, 64); err == nil {
						resultMap[resultMapColumnName] = numVal
					} else {
						//trace.Lg("Failed to parse INT value: %s", err)
						resultMap[resultMapColumnName] = string(v)
					}
				default:
					//trace.Lg("Unexpected type for INT value: %T", v)
					resultMap[resultMapColumnName] = fmt.Sprintf("%v", v)
				}
			case "BOOL":
				if boolVal, ok := (*val).(bool); ok {
					resultMap[resultMapColumnName] = boolVal
				} else {
					//trace.Lg("Unexpected type for BOOL value: %T", *val)
					resultMap[resultMapColumnName] = fmt.Sprintf("%v", *val)
				}
			case "TIMESTAMP", "TIMESTAMPTZ", "DATE":
				if dateVal, ok := (*val).(time.Time); ok {
					resultMap[resultMapColumnName] = dateVal.Format("02012006")
				} else {
					//trace.Lg("Unexpected type for DATE value: %T", *val)
					resultMap[resultMapColumnName] = fmt.Sprintf("%v", *val)
				}
			case "BYTEA":
				switch v := (*val).(type) {
				case []byte:
					resultMap[resultMapColumnName] = base64.StdEncoding.EncodeToString(v)
				default:
					//trace.Lg("Unexpected type for BYTEA value: %T", v)
					resultMap[resultMapColumnName] = fmt.Sprintf("%v", v)
				}
			default:
				resultMap[resultMapColumnName] = fmt.Sprintf("%v", *val)
			}
		}
		*resultMapArray = append(*resultMapArray, resultMap)
		rowcount++
	}
	return rowcount
}

func ReadFromDBTablePostgres(queryStr string, resultMap map[string]interface{}, dbContext dbdef.DBContextDef, dberr *string) int {
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

	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		//trace.Lg("rows.ColumnTypes() failed with err(%s)", err)
		*dberr = err.Error()
		return -1
	}

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
				columnType := columnTypes[i].DatabaseTypeName()
				switch columnType {
				case "VARCHAR", "TEXT":
					switch v := (*val).(type) {
					case []byte:
						resultMap[resultMapColumnName] = string(v)
					case string:
						resultMap[resultMapColumnName] = v
					default:
						//trace.Lg("Unexpected type for VARCHAR/TEXT value: %T", v)
						resultMap[resultMapColumnName] = fmt.Sprintf("%v", v)
					}
				case "NUMERIC":
					// Check for precision and scale
					_, scale, _ := columnTypes[i].DecimalSize()
					if scale == 0 {
						// Integer value
						switch v := (*val).(type) {
						case int64:
							resultMap[resultMapColumnName] = v
						case []byte:
							if numVal, err := strconv.ParseInt(string(v), 10, 64); err == nil {
								resultMap[resultMapColumnName] = numVal
							} else {
								//trace.Lg("Failed to parse NUMERIC value as int64: %s", err)
								resultMap[resultMapColumnName] = string(v)
							}
						default:
							//trace.Lg("Unexpected type for NUMERIC value: %T", v)
							resultMap[resultMapColumnName] = fmt.Sprintf("%v", v)
						}
					} else {
						// Handling for numeric with scale (float)
						switch v := (*val).(type) {
						case float64:
							resultMap[resultMapColumnName] = v
						case []byte:
							if numVal, err := strconv.ParseFloat(string(v), 64); err == nil {
								resultMap[resultMapColumnName] = numVal
							} else {
								//trace.Lg("Failed to parse NUMERIC value as float64: %s", err)
								resultMap[resultMapColumnName] = string(v)
							}
						default:
							//trace.Lg("Unexpected type for NUMERIC value: %T", v)
							resultMap[resultMapColumnName] = fmt.Sprintf("%v", v)
						}
					}
				case "INT4", "INT8":
					switch v := (*val).(type) {
					case int64:
						resultMap[resultMapColumnName] = v
					case []byte:
						if numVal, err := strconv.ParseInt(string(v), 10, 64); err == nil {
							resultMap[resultMapColumnName] = numVal
						} else {
							//trace.Lg("Failed to parse INT value: %s", err)
							resultMap[resultMapColumnName] = string(v)
						}
					default:
						//trace.Lg("Unexpected type for INT value: %T", v)
						resultMap[resultMapColumnName] = fmt.Sprintf("%v", v)
					}
				case "BOOL":
					if boolVal, ok := (*val).(bool); ok {
						resultMap[resultMapColumnName] = boolVal
					} else {
						//trace.Lg("Unexpected type for BOOL value: %T", *val)
						resultMap[resultMapColumnName] = fmt.Sprintf("%v", *val)
					}
				case "TIMESTAMP", "TIMESTAMPTZ", "DATE":
					if dateVal, ok := (*val).(time.Time); ok {
						resultMap[resultMapColumnName] = dateVal.Format("02012006")
					} else {
						//trace.Lg("Unexpected type for DATE value: %T", *val)
						resultMap[resultMapColumnName] = fmt.Sprintf("%v", *val)
					}
				case "BYTEA":
					switch v := (*val).(type) {
					case []byte:
						resultMap[resultMapColumnName] = base64.StdEncoding.EncodeToString(v)
					default:
						//trace.Lg("Unexpected type for BYTEA value: %T", v)
						resultMap[resultMapColumnName] = fmt.Sprintf("%v", v)
					}
				default:
					resultMap[resultMapColumnName] = fmt.Sprintf("%v", *val)
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

func placeholderspostgres(n int) string {
	if n == 0 {
		return ""
	}
	namedPlaceholders := make([]string, n)
	for i := 1; i <= n; i++ {
		namedPlaceholders[i-1] = fmt.Sprintf(":%d", i)
	}
	return strings.Join(namedPlaceholders, ",")
}

func InsertTablePostgres(tableName string, queryStr string, dbContext dbdef.DBContextDef, dberr *string) int {
	var err error
	timeoutCtx, cancelFunc := context.WithTimeout(dbContext.HTTPReqCtx, time.Second*time.Duration(dbContext.DBTimeOut))
	defer cancelFunc()

	execStmt := fmt.Sprintf(`INSERT INTO "%s"."%s" %s`, dbContext.SchemaName, tableName, queryStr)
	_, err = dbContext.DBTx.ExecContext(timeoutCtx, execStmt)
	if err != nil {
		//trace.Lg("Exec() for (%s) failed with err(%s)", execStmt, err)
		*dberr = err.Error()
		return -1
	}
	return 1
}

func UpdateTablePostgres(tableName string, queryStr string, dbContext dbdef.DBContextDef, dberr *string) int {
	var err error

	timeoutCtx, cancelFunc := context.WithTimeout(dbContext.HTTPReqCtx, time.Second*time.Duration(dbContext.DBTimeOut))
	defer cancelFunc()

	exectStmt := fmt.Sprintf(`UPDATE "%s"."%s" SET %s`, dbContext.SchemaName, tableName, queryStr)
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

func DeleteTablePostgres(tableName string, queryStr string, dbContext dbdef.DBContextDef, dberr *string) int {
	var err error

	timeoutCtx, cancelFunc := context.WithTimeout(dbContext.HTTPReqCtx, time.Second*time.Duration(dbContext.DBTimeOut))
	defer cancelFunc()
	exectStmt := fmt.Sprintf(`DELETE FROM "%s"."%s" WHERE %s`, dbContext.SchemaName, tableName, queryStr)
	_, err = dbContext.DBTx.ExecContext(timeoutCtx, exectStmt)
	if err != nil {
		//trace.Lg("Exec() for (%s) failed with err(%s)", queryStr, err)
		*dberr = err.Error()
		return -1
	}
	return 1
}
