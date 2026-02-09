package microsv_common

import (
	"fmt"
	"lmsapieng/include/common/datatypedef"
	"lmsapieng/include/common/dbdef"
	"lmsapieng/include/common/lexicalparserdef"
	"lmsapieng/include/common/rejectdef"
	"lmsapieng/include/common/reqbrokerdef"
	"lmsapieng/include/common/seqdef"
	"lmsapieng/libsrc/microsv/microsv_sequence"
	"lmsapieng/libsrc/utils/dtutil"
	"lmsapieng/libsrc/utils/msgutil"
	"lmsapieng/libsrc/utils/templateutil"
	"strconv"
	"strings"
)

func getJulDay(dateString string) int {

	monthlist := [13]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	ddnum, _ := strconv.Atoi(string(dateString[0:2]))
	mmnum, _ := strconv.Atoi(string(dateString[2:4]))
	yynum, _ := strconv.Atoi(string(dateString[2:4]))
	jday := 0

	for i := 0; i < mmnum-1; i++ {
		jday += monthlist[i]
	}
	if yynum%4 == 0 && mmnum >= 3 {
		jday += ddnum
	}
	return jday
}

func genRRN(stan string) string {
	dateString := dtutil.GetDate("DDMMYY")
	timeString := dtutil.GetDate("HHMMSS")
	return fmt.Sprintf("%c%03d%s%s", dateString[5], getJulDay(dateString), string(timeString[0:2]), stan)
}

func GenTxnBatchNum(txnBatchNum *string) int {
	var seqRespInfo seqdef.SeqReqInfoRespStruct
	RequestNumFlag := "Y"
	RecordNumFlag := "N"
	RRNFlag := "N"
	StanFlag := "N"
	retval, _ := microsv_sequence.GenReqInfo(&seqRespInfo, RequestNumFlag, RecordNumFlag, RRNFlag, StanFlag)
	if retval < 0 {
		return -1
	}
	*txnBatchNum = seqRespInfo.RequestNum
	return 1
}

func GenTxnRecordNum(txnRecordNum *string) int {
	var seqRespInfo seqdef.SeqReqInfoRespStruct
	RequestNumFlag := "N"
	RecordNumFlag := "Y"
	RRNFlag := "N"
	StanFlag := "N"
	retval, _ := microsv_sequence.GenReqInfo(&seqRespInfo, RequestNumFlag, RecordNumFlag, RRNFlag, StanFlag)
	if retval < 0 {
		return -1
	}
	*txnRecordNum = dtutil.GetDateTimeVal() + seqRespInfo.RecordNum
	return 1
}

func GenRRN(RRN *string) int {
	var seqRespInfo seqdef.SeqReqInfoRespStruct
	RequestNumFlag := "N"
	RecordNumFlag := "N"
	RRNFlag := "Y"
	StanFlag := "N"
	retval, _ := microsv_sequence.GenReqInfo(&seqRespInfo, RequestNumFlag, RecordNumFlag, RRNFlag, StanFlag)
	if retval < 0 {
		return -1
	}
	*RRN = genRRN(seqRespInfo.RRN)
	return 1
}

func GenStan(Stan *string) int {
	var seqRespInfo seqdef.SeqReqInfoRespStruct
	RequestNumFlag := "N"
	RecordNumFlag := "N"
	RRNFlag := "N"
	StanFlag := "Y"
	retval, _ := microsv_sequence.GenReqInfo(&seqRespInfo, RequestNumFlag, RecordNumFlag, RRNFlag, StanFlag)
	if retval < 0 {
		return -1
	}
	*Stan = seqRespInfo.Stan
	return 1
}

func GetSchemaName(ModuleName *string, reqBrokerDataMap map[string]interface{}) int {
	var SchemaInfo lexicalparserdef.LPSchemaInfoStruct
	var rejectDesc string
	retval, contextParams := GetDBContextParamsByModuleName(reqBrokerDataMap, SchemaInfo, &rejectDesc)
	if retval < 0 {
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GetDBContextParamsErr, templateutil.GetTemplateStr(reqBrokerDataMap, ""), templateutil.GetTemplateStr(reqBrokerDataMap, ""), rejectDesc)
		return -1
	}
	*ModuleName = contextParams[0]
	return 1

}

func GetDBContextParamsByModuleName(reqBrokerDataMap map[string]interface{}, schemaInfo lexicalparserdef.LPSchemaInfoStruct, rejectDesc *string) (int, []string) {
	contextParams := ""
	if len(schemaInfo.ModuleName) != 0 {
		contextParams += schemaInfo.ModuleName
	} else {
		contextParams += dbdef.DBModuleLMS
	}
	contextParams += ","
	if len(schemaInfo.SubModuleName) != 0 {
		contextParams += schemaInfo.SubModuleName
	} else {
		contextParams += "*"
	}
	contextParams += ","
	if len(schemaInfo.InstID) != 0 {
		dataValue := ReadValueFromDataMap(reqBrokerDataMap, schemaInfo.InstID, datatypedef.DataTypeString, rejectDesc)
		if dataValue == nil {
			//trace.Lg("ReadValueFromDataMap() failed for schemaInfo.InstID[%s]", schemaInfo.InstID)
			return -1, nil
		}
		contextParams += dataValue.(string)
	} else {
		contextParams += "*"
	}
	contextParams += ","
	contextParams += "*"
	contextParams += ","
	if len(schemaInfo.BinID) != 0 {
		dataValue := ReadValueFromDataMap(reqBrokerDataMap, schemaInfo.BinID, datatypedef.DataTypeString, rejectDesc)
		if dataValue == nil {
			//trace.Lg("ReadValueFromDataMap() failed for schemaInfo.BinID[%s]", schemaInfo.BinID)
			return -1, nil
		}
		contextParams += dataValue.(string)
	} else {
		contextParams += "*"
	}
	contextParams += ","
	contextParams += "*"
	return 1, strings.Split(contextParams, ",")
}
