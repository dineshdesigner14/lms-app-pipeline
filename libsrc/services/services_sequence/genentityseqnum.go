package services_sequence

import (
	"lmsapieng/include/common/dbdef"
	"lmsapieng/include/common/entitytypedef"
	"lmsapieng/include/common/globaldef"
	"lmsapieng/include/common/seqobjdef"
	"lmsapieng/include/dbtabdef"
	dbtab_seqinfo "lmsapieng/libsrc/dbtab/dbtab_lmsseqdata"
	"lmsapieng/libsrc/utils/datatypeutil"
	"lmsapieng/libsrc/utils/dbutil"
	"lmsapieng/libsrc/utils/schemainfo"
	"lmsapieng/libsrc/utils/sequtil"
	"strconv"
)

func genEntitySeqNum(reqBrokerDataMap map[string]interface{}, InstID string, EntityType string, PrefixStr string, SeqLen int, OverflowFlag string, SeqNum *string, contextParams ...string) int {
	var rejectDesc string
	var dbErr, dbRejectReason string
	var resultDBContext dbdef.DBContextDef
	var cmsSeqDataRec dbtabdef.LMSSeqDataTable
	var retVal int
	if schemainfo.GetActiveDBContext(reqBrokerDataMap, &resultDBContext, &rejectDesc, contextParams[0], contextParams[1], contextParams[2], contextParams[3], contextParams[4], contextParams[5]) < 0 {
		//trace.Lg("GetActiveDBContext() failed for Module:%s", contextParams)
		return -1
	}
	if dbtab_seqinfo.ReadLMSSeqDataTable(InstID, EntityType, PrefixStr, &cmsSeqDataRec, resultDBContext, &dbErr, &dbRejectReason) < 0 {
		if !dbutil.IsNoRows(resultDBContext.DBType, dbErr) {
			//trace.Lg("ReadLMSSeqDataTable() failed for InstID[%s] EntityType[%s] PrefixStr[%s] with DBErr[%s]", InstID, EntityType, PrefixStr, dbErr)
			return -1
		}
	}
	if dbutil.IsNoRows(resultDBContext.DBType, dbErr) {
		if schemainfo.GetActiveDBContextWithTxn(reqBrokerDataMap, &resultDBContext, &rejectDesc, contextParams[0], contextParams[1], contextParams[2], contextParams[3], contextParams[4], contextParams[5]) < 0 {
			//trace.Lg("GetActiveDBContextWithTxn() failed for Module:%s", contextParams)
			return -1
		}
		cmsSeqDataRec.InstID = InstID
		cmsSeqDataRec.EntityType = EntityType
		cmsSeqDataRec.SeqPrefix = PrefixStr
		cmsSeqDataRec.SeqLen = SeqLen
		cmsSeqDataRec.OverflowFlag = OverflowFlag
		retVal, cmsSeqDataRec.SeqNum = sequtil.GetNextSequence("0", SeqLen)
		if retVal < 0 {
			//trace.Lg("sequtil.GetNextSequence() failed for SeqNum[%s]SeqLen[%d]", "0", SeqLen)
			return -1
		}
		if dbtab_seqinfo.InsertLMSSeqDataTable(&cmsSeqDataRec, resultDBContext, &dbErr, &dbRejectReason) < 0 {
			//trace.Lg("InsertLMSSeqDataTable() failed with DBErr[%s]", dbErr)
			return -1
		}
	} else {
		if cmsSeqDataRec.SeqLen != SeqLen {
			//trace.Lg("DBSeqLen[%d] does not match InputSeqLen[%d]", cmsSeqDataRec.SeqLen, SeqLen)
			return -1
		}
		if cmsSeqDataRec.OverflowFlag != OverflowFlag {
			//trace.Lg("DBOverflowFlag[%s] does not match InputOverflowFlag[%s]", cmsSeqDataRec.OverflowFlag, OverflowFlag)
			return -1
		}
		retVal, cmsSeqDataRec.SeqNum = sequtil.GetNextSequence(cmsSeqDataRec.SeqNum, SeqLen)
		if retVal < 0 {
			//trace.Lg("sequtil.GetNextSequence() failed for SeqNum[%s]SeqLen[%d]", cmsSeqDataRec.SeqNum, SeqLen)
			return -1
		}
		if schemainfo.GetActiveDBContextWithTxn(reqBrokerDataMap, &resultDBContext, &rejectDesc, contextParams[0], contextParams[1], contextParams[2], contextParams[3], contextParams[4], contextParams[5]) < 0 {
			//trace.Lg("GetActiveDBContextWithTxn() failed for Module:%s", contextParams)
			return -1
		}
		if dbtab_seqinfo.UpdateLMSSeqDataTable(&cmsSeqDataRec, resultDBContext, &dbErr, &dbRejectReason) < 0 {
			//trace.Lg("UpdateLMSSeqDataTable() failed with DBErr[%s]", dbErr)
			return -1
		}
	}
	*SeqNum = cmsSeqDataRec.SeqNum
	return 1
}

func GenEntitySeqNum(reqBrokerDataMap map[string]interface{}, seqObj map[string]interface{}, contextParams ...string) int {
	var SeqNum string
	if !datatypeutil.IsObject(seqObj) {
		//trace.Lg("seqObj is not a Map")
		return -1
	}
	_, ok := seqObj[seqobjdef.InstIDKey]
	if !ok {
		//trace.Lg("[%s] not present in seqObj", seqobjdef.InstIDKey)
		return -1
	}
	if !datatypeutil.IsString(seqObj[seqobjdef.InstIDKey]) {
		//trace.Lg("%s seqObj has invalid type[%T]", seqobjdef.InstIDKey, seqObj[seqobjdef.InstIDKey])
		return -1
	}
	InstID := seqObj[seqobjdef.InstIDKey].(string)

	_, ok = seqObj[seqobjdef.EntityTypeKey]
	if !ok {
		//trace.Lg("[%s] not present in seqObj", seqobjdef.EntityTypeKey)
		return -1
	}
	if !datatypeutil.IsString(seqObj[seqobjdef.EntityTypeKey]) {
		//trace.Lg("%s seqObj has invalid type[%T]", seqobjdef.EntityTypeKey, seqObj[seqobjdef.EntityTypeKey])
		return -1
	}
	EntityType := seqObj[seqobjdef.EntityTypeKey].(string)

	_, ok = seqObj[seqobjdef.SeqPrefixKey]
	if !ok {
		//trace.Lg("[%s] not present in seqObj", seqobjdef.SeqPrefixKey)
		return -1
	}
	if !datatypeutil.IsString(seqObj[seqobjdef.SeqPrefixKey]) {
		//trace.Lg("%s seqObj has invalid type[%T]", seqobjdef.SeqPrefixKey, seqObj[seqobjdef.SeqPrefixKey])
		return -1
	}
	PrefixStr := seqObj[seqobjdef.SeqPrefixKey].(string)

	_, ok = seqObj[seqobjdef.SeqLenKey]
	if !ok {
		//trace.Lg("[%s] not present in seqObj", seqobjdef.SeqLenKey)
		return -1
	}
	if !datatypeutil.IsString(seqObj[seqobjdef.SeqLenKey]) {
		//trace.Lg("%s seqObj has invalid type[%T]", seqobjdef.SeqLenKey, seqObj[seqobjdef.SeqLenKey])
		return -1
	}
	SeqLen, err := strconv.Atoi(seqObj[seqobjdef.SeqLenKey].(string))
	if err != nil {
		//trace.Lg("%s seqObj value[%s] should be a number", seqobjdef.SeqLenKey, seqObj[seqobjdef.SeqLenKey])
		return -1
	}
	OverflowFlag := "N"
	_, ok = seqObj[seqobjdef.OverflowFlagKey]
	if !ok {
		OverflowFlag = "N"
	} else {
		if !datatypeutil.IsString(seqObj[seqobjdef.OverflowFlagKey]) {
			//trace.Lg("%s seqObj has invalid type[%T]", seqobjdef.OverflowFlagKey, seqObj[seqobjdef.OverflowFlagKey])
			return -1
		}
		OverflowFlag = seqObj[seqobjdef.OverflowFlagKey].(string)
	}
	if OverflowFlag != "Y" && OverflowFlag != "N" {
		//trace.Lg("%s seqObj has invalid value[%s]", seqobjdef.OverflowFlagKey, OverflowFlag)
		return -1
	}
	// trace.Lg("GenEntitySeqNum() called with InstID[%s] EntityType[%s] PrefixStr[%s] SeqLen[%d] OverflowFlag[%s]", InstID, EntityType, PrefixStr, SeqLen, OverflowFlag)

	if EntityType == entitytypedef.EntityTypeDefCardNum {
		if SeqLen != 16 && SeqLen != 19 {
			//trace.Lg("invalid sequence length[%d] for entitytype[%s]..should be 16 or 19", EntityType)
			return -1
		}
	}
	if EntityType == entitytypedef.EntityTypeDefCardRefNum || EntityType == entitytypedef.EntityTypeDefCardSerialNum {
		if SeqLen != 24 {
			//trace.Lg("invalid sequence length[%d] for entitytype[%s]..should be 24", EntityType)
			return -1
		}
	}
	if PrefixStr != globaldef.NOT_INITIALIZED {
		if EntityType == entitytypedef.EntityTypeDefCardNum {
			SeqLen = SeqLen - len(PrefixStr) - 1
		} else {
			SeqLen = SeqLen - len(PrefixStr)
		}
	}
	if genEntitySeqNum(reqBrokerDataMap, InstID, EntityType, PrefixStr, SeqLen, OverflowFlag, &SeqNum, contextParams...) < 0 {
		return -1
	}
	seqObj[seqobjdef.SeqNumKey] = SeqNum
	if PrefixStr != globaldef.NOT_INITIALIZED {
		seqObj[seqobjdef.SeqValueKey] = PrefixStr + SeqNum
	} else {
		seqObj[seqobjdef.SeqValueKey] = SeqNum
	}
	return 1
}
