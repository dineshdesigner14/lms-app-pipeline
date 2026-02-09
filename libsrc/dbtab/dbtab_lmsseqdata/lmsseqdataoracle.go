package dbtab_seqinfo

import (
	"context"
	"database/sql"
	"fmt"
	"lmsapieng/include/common/dbdef"
	"lmsapieng/include/dbtabdef"
	"time"
)

func ReadLMSSeqDataTableOracle(InstID string, EntityType string, SeqPrefix string, LMSSeqDataPtr *dbtabdef.LMSSeqDataTable, dbContext dbdef.DBContextDef, dberr *string) int {
	var err error
	var lSeqLen sql.NullInt64
	var lOverflowFlag sql.NullString
	var lSeqNum sql.NullString

	timeoutCtx, cancelFunc := context.WithTimeout(dbContext.HTTPReqCtx, time.Second*time.Duration(dbContext.DBTimeOut))
	defer cancelFunc()

	queryStmt := fmt.Sprintf(
		`SELECT 
			SEQ_LEN,
			OVERFLOW_FLAG,
			SEQ_NUM
		FROM LMS_SEQ_DATA WHERE INST_ID='%s' AND ENTITY_TYPE='%s' AND SEQ_PREFIX='%s'`, InstID, EntityType, SeqPrefix)

	if dbContext.DBTxFlag {
		err = dbContext.DBTx.QueryRowContext(timeoutCtx, queryStmt).Scan(
			&lSeqLen,
			&lOverflowFlag,
			&lSeqNum,
		)
	} else {
		err = dbContext.DBID.QueryRowContext(timeoutCtx, queryStmt).Scan(
			&lSeqLen,
			&lOverflowFlag,
			&lSeqNum,
		)
	}
	if err != nil {
		//trace.Lg("QueryRow() for (%s) failed with err(%s)", queryStmt, err)
		*dberr = err.Error()
		return -1
	}
	LMSSeqDataPtr.InstID = InstID
	LMSSeqDataPtr.EntityType = EntityType
	LMSSeqDataPtr.SeqPrefix = SeqPrefix
	LMSSeqDataPtr.SeqLen = int(lSeqLen.Int64)
	LMSSeqDataPtr.OverflowFlag = lOverflowFlag.String
	LMSSeqDataPtr.SeqNum = lSeqNum.String
	return 1
}

func InsertLMSSeqDataTableOracle(LMSSeqDataPtr *dbtabdef.LMSSeqDataTable, dbContext dbdef.DBContextDef, dberr *string) int {
	var err error

	timeoutCtx, cancelFunc := context.WithTimeout(dbContext.HTTPReqCtx, time.Second*time.Duration(dbContext.DBTimeOut))
	defer cancelFunc()

	execStmt := fmt.Sprintf(`
		INSERT INTO LMS_SEQ_DATA
		(
			INST_ID,
			ENTITY_TYPE, 
			SEQ_PREFIX,
			SEQ_LEN,
			OVERFLOW_FLAG,
			SEQ_NUM
		)
		VALUES
	(       '%s',
			'%s',
			'%s',
			'%d',
			'%s',
			'%s'
		)
		`,
		LMSSeqDataPtr.InstID,
		LMSSeqDataPtr.EntityType,
		LMSSeqDataPtr.SeqPrefix,
		LMSSeqDataPtr.SeqLen,
		LMSSeqDataPtr.OverflowFlag,
		LMSSeqDataPtr.SeqNum,
	)
	_, err = dbContext.DBTx.ExecContext(timeoutCtx, execStmt)
	if err != nil {
		//trace.Lg("Exec() for (%s) failed with err(%s)", execStmt, err)
		*dberr = err.Error()
		return -1
	}
	return 1
}

func UpdateLMSSeqDataTableOracle(LMSSeqDataPtr *dbtabdef.LMSSeqDataTable, dbContext dbdef.DBContextDef, dberr *string) int {
	var err error

	timeoutCtx, cancelFunc := context.WithTimeout(dbContext.HTTPReqCtx, time.Second*time.Duration(dbContext.DBTimeOut))
	defer cancelFunc()

	execStmt := fmt.Sprintf(`
	UPDATE LMS_SEQ_DATA SET
		SEQ_NUM = '%s'
	WHERE INST_ID='%s' AND ENTITY_TYPE='%s' AND SEQ_PREFIX='%s'`,
		LMSSeqDataPtr.SeqNum,
		LMSSeqDataPtr.InstID,
		LMSSeqDataPtr.EntityType,
		LMSSeqDataPtr.SeqPrefix,
	)
	res, err := dbContext.DBTx.ExecContext(timeoutCtx, execStmt)
	if err != nil {
		//trace.Lg("Exec() for (%s) failed with err(%s)", execStmt, err)
		*dberr = err.Error()
		return -1
	}
	count, err := res.RowsAffected()
	if err != nil {
		//trace.Lg("Exec() for (%s) failed with err(%s)", execStmt, err)
		*dberr = err.Error()
		return -1
	}
	if count == 0 {
		*dberr = "No Rows Updated"
		return -1
	}
	// trace.Lg("Exec() for (%s) Success with count(%d)", execStmt, count)
	return 1
}
