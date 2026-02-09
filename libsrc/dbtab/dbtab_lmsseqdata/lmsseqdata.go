package dbtab_seqinfo

import (
	"fmt"
	"lmsapieng/include/common/dbdef"
	"lmsapieng/include/dbtabdef"
	"strings"
)

func ReadLMSSeqDataTable(InstID string, EntityType string, SeqPrefix string, LMSSeqDataPtr *dbtabdef.LMSSeqDataTable, dbContext dbdef.DBContextDef, dberr *string, dbrejectreason *string) int {
	if strings.EqualFold(dbContext.DBType, dbdef.DBTypeOracle) {
		if ReadLMSSeqDataTableOracle(InstID, EntityType, SeqPrefix, LMSSeqDataPtr, dbContext, dberr) < 0 {
			//trace.Lg("ReadLMSSeqDataTableOracle() failed for InstID[%s] EntityType[%s] SeqPrefix[%s] with DbErr[%s]", InstID, EntityType, SeqPrefix, *dberr)
			*dbrejectreason = fmt.Sprintf("ReadLMSSeqDataTableOracle() failed for InstID[%s] EntityType[%s] SeqPrefix[%s] with DbErr[%s]", InstID, EntityType, SeqPrefix, *dberr)
			return -1
		}
	} else if strings.EqualFold(dbContext.DBType, dbdef.DBTypePostgreSQL) {
		if ReadLMSSeqDataTablePostgres(InstID, EntityType, SeqPrefix, LMSSeqDataPtr, dbContext, dberr) < 0 {
			//trace.Lg("ReadLMSSeqDataTablePostgres() failed for InstID[%s] EntityType[%s] SeqPrefix[%s] with DbErr[%s]", InstID, EntityType, SeqPrefix, *dberr)
			*dbrejectreason = fmt.Sprintf("ReadLMSSeqDataTablePostgres() failed for InstID[%s] EntityType[%s] SeqPrefix[%s] with DbErr[%s]", InstID, EntityType, SeqPrefix, *dberr)
			return -1
		}
	} else {
		//trace.Lg("db_type(%s) is Invalid...ReadLMSSeqDataTable() failed", dbContext.DBType)
		*dbrejectreason = fmt.Sprintf("db_type %s is Invalid...ReadLMSSeqDataTable() failed", dbContext.DBType)
		return -1
	}
	//trace.Lg("ReadLMSSeqDataTable() Success")
	return 1
}

func InsertLMSSeqDataTable(LMSSeqDataPtr *dbtabdef.LMSSeqDataTable, dbContext dbdef.DBContextDef, dberr *string, dbrejectreason *string) int {
	if !dbContext.DBTxFlag {
		//trace.Lg("DBTxFlag not set...InsertLMSSeqDataTable() failed")
		*dbrejectreason = fmt.Sprintf("DBTxFlag not set...InsertLMSSeqDataTable() failed")
		return -1
	}
	if strings.EqualFold(dbContext.DBType, dbdef.DBTypeOracle) {
		if InsertLMSSeqDataTableOracle(LMSSeqDataPtr, dbContext, dberr) < 0 {
			//trace.Lg("InsertLMSSeqDataTableOracle() failed with dberr<%s>", *dberr)
			*dbrejectreason = fmt.Sprintf("InsertLMSSeqDataTableOracle failed with DBErr %s", *dberr)
			return -1
		}
	} else if strings.EqualFold(dbContext.DBType, dbdef.DBTypePostgreSQL) {
		if InsertLMSSeqDataTablePostgres(LMSSeqDataPtr, dbContext, dberr) < 0 {
			//trace.Lg("InsertLMSSeqDataTablePostgres() failed with dberr<%s>", *dberr)
			*dbrejectreason = fmt.Sprintf("InsertLMSSeqDataTablePostgres failed with DBErr %s", *dberr)
			return -1
		}
	} else {
		//trace.Lg("db_type(%s) is Invalid...InsertLMSSeqDataTable() failed", dbContext.DBType)
		*dbrejectreason = fmt.Sprintf("db_type %s is Invalid...InsertLMSSeqDataTable() failed", dbContext.DBType)
		return -1
	}
	//trace.Lg("InsertLMSSeqDataTable() Success")
	return 1
}

func UpdateLMSSeqDataTable(LMSSeqDataPtr *dbtabdef.LMSSeqDataTable, dbContext dbdef.DBContextDef, dberr *string, dbrejectreason *string) int {
	if !dbContext.DBTxFlag {
		//trace.Lg("DBTxFlag not set...Update_1_LMSSeqDataTable() failed")
		*dbrejectreason = fmt.Sprintf("DBTxFlag not set...Update_1_LMSSeqDataTable() failed")
		return -1
	}
	if strings.EqualFold(dbContext.DBType, dbdef.DBTypeOracle) {
		if UpdateLMSSeqDataTableOracle(LMSSeqDataPtr, dbContext, dberr) < 0 {
			//trace.Lg("UpdateLMSSeqDataTableOracle() failed for InstID[%s] EntityType[%s] SeqPrefix[%s] SeqLen[%d] OverflowFlag[%s] SeqNum[%s] with DbErr[%s]", LMSSeqDataPtr.InstID, LMSSeqDataPtr.EntityType, LMSSeqDataPtr.SeqPrefix, LMSSeqDataPtr.SeqLen, LMSSeqDataPtr.OverflowFlag, LMSSeqDataPtr.SeqNum, *dberr)
			*dbrejectreason = fmt.Sprintf("UpdateLMSSeqDataTableOracle() failed for InstID[%s] EntityType[%s] SeqPrefix[%s] SeqLen[%d] OverflowFlag[%s] SeqNum[%s] with DbErr[%s]", LMSSeqDataPtr.InstID, LMSSeqDataPtr.EntityType, LMSSeqDataPtr.SeqPrefix, LMSSeqDataPtr.SeqLen, LMSSeqDataPtr.OverflowFlag, LMSSeqDataPtr.SeqNum, *dberr)
			return -1
		}
	} else if strings.EqualFold(dbContext.DBType, dbdef.DBTypePostgreSQL) {
		if UpdateLMSSeqDataTablePostgres(LMSSeqDataPtr, dbContext, dberr) < 0 {
			//trace.Lg("UpdateLMSSeqDataTablePostgres() failed for InstID[%s] EntityType[%s] SeqPrefix[%s] SeqLen[%d] OverflowFlag[%s] SeqNum[%s] with DbErr[%s]", LMSSeqDataPtr.InstID, LMSSeqDataPtr.EntityType, LMSSeqDataPtr.SeqPrefix, LMSSeqDataPtr.SeqLen, LMSSeqDataPtr.OverflowFlag, LMSSeqDataPtr.SeqNum, *dberr)
			*dbrejectreason = fmt.Sprintf("UpdateLMSSeqDataTablePostgres() failed for InstID[%s] EntityType[%s] SeqPrefix[%s] SeqLen[%d] OverflowFlag[%s] SeqNum[%s] with DbErr[%s]", LMSSeqDataPtr.InstID, LMSSeqDataPtr.EntityType, LMSSeqDataPtr.SeqPrefix, LMSSeqDataPtr.SeqLen, LMSSeqDataPtr.OverflowFlag, LMSSeqDataPtr.SeqNum, *dberr)
			return -1
		}
	} else {
		//trace.Lg("db_type(%s) is Invalid...UpdateLMSSeqDataTable() failed", dbContext.DBType)
		*dbrejectreason = fmt.Sprintf("db_type %s is Invalid...UpdateLMSSeqDataTable() failed", dbContext.DBType)
		return -1
	}
	//trace.Lg("UpdateLMSSeqDataTable() Success")
	return 1
}
