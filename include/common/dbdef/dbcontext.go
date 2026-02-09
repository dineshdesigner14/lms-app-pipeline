package dbdef

import (
	"context"
	"database/sql"
)

const (
	DBContextArrayObj = "dbcontextarray"
)

type DBContextDef struct {
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
	DBID             *sql.DB
	DBTx             *sql.Tx
	DBTxFlag         bool
	HTTPReqCtx       context.Context
	DBTimeOut        int
	ConnectionStr    string
	DBDriverName     string
	SchemaName       string
}
