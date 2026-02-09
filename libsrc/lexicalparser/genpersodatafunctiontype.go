package lexicalparser

import (
	"fmt"
	"lmsapieng/include/common/carddef"
	"lmsapieng/include/common/debugdef"
	"lmsapieng/include/common/globaldef"
	"lmsapieng/include/common/lexicalparserdef"
	"lmsapieng/include/common/rejectdef"
	"lmsapieng/include/common/reqbrokerdef"
	"lmsapieng/libsrc/utils/cardutil"
	"lmsapieng/libsrc/utils/datatypeutil"
	"lmsapieng/libsrc/utils/dtutil"
	"lmsapieng/libsrc/utils/mathutil"
	"lmsapieng/libsrc/utils/msgutil"
	"lmsapieng/libsrc/utils/templateutil"
	"lmsapieng/libsrc/utils/trace"
	"os"
)

func writeCPFFile(embRecordList []string, BIN string, CardSubProduct string, lBatchID string) int {
	fileName := fmt.Sprintf("%s.%s.%s.%s.%s.CPF", BIN, CardSubProduct, lBatchID, dtutil.GetDate("DDMMYYYY"), dtutil.GetTime("HHMMSS"))
	filePath := fmt.Sprintf("%s/%s/%s/%s", globaldef.GetAppBaseDir(), "data", "CPF", fileName)

	file, err := os.Create(filePath)
	if err != nil {
		//trace.Lg("os.Create() failed with err[%s]", err)
		return -1
	}
	defer file.Close()

	for _, record := range embRecordList {
		_, err := file.WriteString(record + "\n")
		if err != nil {
			//trace.Lg("WriteString for record[%s] failed with err[%s]", record, err)
			return -1
		}
	}
	return 1
}

func genPersoFile(reqBrokerDataMap map[string]interface{}, persoObject map[string]interface{}, persoFileObject []map[string]interface{}) int {
	_, ok := persoObject[carddef.CardPersoRecordKeyBIN]
	if !ok {
		//trace.Lg("[%s] not present in persoObject", carddef.CardPersoRecordKeyBIN)
		return -1
	}
	if !datatypeutil.IsString(persoObject[carddef.CardPersoRecordKeyBIN]) {
		//trace.Lg("[%s] not a string in persoObject", carddef.CardPersoRecordKeyBIN)
		return -1
	}
	lBIN := persoObject[carddef.CardPersoRecordKeyBIN].(string)
	if !mathutil.IsValidNumericStr(lBIN, 6, 12) {
		//trace.Lg("BIN[%s] is invalid in persoObject", lBIN)
		return -1
	}

	_, ok = persoObject[carddef.CardPersoRecordKeyCardSubProduct]
	if !ok {
		//trace.Lg("[%s] not present in persoObject", carddef.CardPersoRecordKeyCardSubProduct)
		return -1
	}
	if !datatypeutil.IsString(persoObject[carddef.CardPersoRecordKeyCardSubProduct]) {
		//trace.Lg("[%s] not a string in persoObject", carddef.CardPersoRecordKeyCardSubProduct)
		return -1
	}
	lCardSubProduct := persoObject[carddef.CardPersoRecordKeyCardSubProduct].(string)

	if !mathutil.IsValidNumericStr(lCardSubProduct, 3, 4) {
		//trace.Lg("CardSubProduct[%s] is invalid in persoObject", lCardSubProduct)
		return -1
	}

	_, ok = persoObject[carddef.CardPersoRecordKeyBatchID]
	if !ok {
		//trace.Lg("[%s] not present in persoObject", carddef.CardPersoRecordKeyBatchID)
		return -1
	}
	if !datatypeutil.IsString(persoObject[carddef.CardPersoRecordKeyBatchID]) {
		//trace.Lg("[%s] not a string in persoObject", carddef.CardPersoRecordKeyBatchID)
		return -1
	}
	lBatchID := persoObject[carddef.CardPersoRecordKeyBatchID].(string)

	_, ok = persoObject[carddef.CardPersoRecordKeyServerFlag]
	if !ok {
		//trace.Lg("[%s] not present in persoObject", carddef.CardPersoRecordKeyServerFlag)
		return -1
	}
	if !datatypeutil.IsString(persoObject[carddef.CardPersoRecordKeyServerFlag]) {
		//trace.Lg("[%s] not a string in persoObject", carddef.CardPersoRecordKeyServerFlag)
		return -1
	}
	lServerFlag := persoObject[carddef.CardPersoRecordKeyServerFlag].(string)

	if lServerFlag != "0" && lServerFlag != "1" {
		//trace.Lg("lServerFlag[%s] is invalid in persoObject", lServerFlag)
		return -1
	}

	embRecordList := make([]string, 0)
	for i, persoRecordObject := range persoFileObject {
		var embRecordInfo carddef.CardEmbossingFileStruct
		_, ok := persoRecordObject[carddef.CardPersoRecordKeyCardTypeID]
		if !ok {
			trace.Log(debugdef.DEBUG_LEVEL_TEST, "[%s] not present in persoRecordObject[%d]", carddef.CardPersoRecordKeyCardTypeID, i)
			return -1
		}
		if !datatypeutil.IsString(persoRecordObject[carddef.CardPersoRecordKeyCardTypeID]) {
			//trace.Lg("[%s] not a string in persoRecordObject[%d]", carddef.CardPersoRecordKeyCardTypeID, i)
			return -1
		}
		embRecordInfo.CardTypeID = persoRecordObject[carddef.CardPersoRecordKeyCardNum].(string)

		_, ok = persoRecordObject[carddef.CardPersoRecordKeyCardNum]
		if !ok {
			//trace.Lg("[%s] not present in persoRecordObject[%d]", carddef.CardPersoRecordKeyCardNum, i)
			return -1
		}
		if !datatypeutil.IsString(persoRecordObject[carddef.CardPersoRecordKeyCardNum]) {
			//trace.Lg("[%s] not a string in persoRecordObject[%d]", carddef.CardPersoRecordKeyCardNum, i)
			return -1
		}
		embRecordInfo.CardNum = persoRecordObject[carddef.CardPersoRecordKeyCardNum].(string)

		_, ok = persoRecordObject[carddef.CardPersoRecordKeyCardSerialNum]
		if !ok {
			//trace.Lg("[%s] not present in persoRecordObject[%d]", carddef.CardPersoRecordKeyCardSerialNum, i)
			return -1
		}
		if !datatypeutil.IsString(persoRecordObject[carddef.CardPersoRecordKeyCardSerialNum]) {
			//trace.Lg("[%s] not a string in persoRecordObject[%d]", carddef.CardPersoRecordKeyCardSerialNum, i)
			return -1
		}
		embRecordInfo.CardSerialNum = persoRecordObject[carddef.CardPersoRecordKeyCardSerialNum].(string)

		_, ok = persoRecordObject[carddef.CardPersoRecordKeyCardSeqNum]
		if !ok {
			//trace.Lg("[%s] not present in persoRecordObject[%d]", carddef.CardPersoRecordKeyCardSeqNum, i)
			return -1
		}
		if !datatypeutil.IsString(persoRecordObject[carddef.CardPersoRecordKeyCardSeqNum]) {
			//trace.Lg("[%s] not a string in persoRecordObject[%d]", carddef.CardPersoRecordKeyCardSeqNum, i)
			return -1
		}
		embRecordInfo.CardSeqNum = persoRecordObject[carddef.CardPersoRecordKeyCardSeqNum].(string)

		_, ok = persoRecordObject[carddef.CardPersoRecordKeyExpiryDate]
		if !ok {
			//trace.Lg("[%s] not present in persoRecordObject[%d]", carddef.CardPersoRecordKeyExpiryDate, i)
			return -1
		}
		if !datatypeutil.IsString(persoRecordObject[carddef.CardPersoRecordKeyExpiryDate]) {
			//trace.Lg("[%s] not a string in persoRecordObject[%d]", carddef.CardPersoRecordKeyExpiryDate, i)
			return -1
		}
		embRecordInfo.ExpiryDate = persoRecordObject[carddef.CardPersoRecordKeyExpiryDate].(string)

		embRecordInfo.ExpiryDate = embRecordInfo.ExpiryDate[6:] + embRecordInfo.ExpiryDate[2:4]

		_, ok = persoRecordObject[carddef.CardPersoRecordKeyCardName]
		if !ok {
			//trace.Lg("[%s] not present in persoRecordObject[%d]", carddef.CardPersoRecordKeyCardName, i)
			return -1
		}
		if !datatypeutil.IsString(persoRecordObject[carddef.CardPersoRecordKeyCardName]) {
			//trace.Lg("[%s] not a string in persoRecordObject[%d]", carddef.CardPersoRecordKeyCardName, i)
			return -1
		}
		embRecordInfo.CardName = persoRecordObject[carddef.CardPersoRecordKeyCardName].(string)

		_, ok = persoRecordObject[carddef.CardPersoRecordKeyCoBrandedName]
		if !ok {
			//trace.Lg("[%s] not present in persoRecordObject[%d]", carddef.CardPersoRecordKeyCoBrandedName, i)
			return -1
		}
		if !datatypeutil.IsString(persoRecordObject[carddef.CardPersoRecordKeyCoBrandedName]) {
			//trace.Lg("[%s] not a string in persoRecordObject[%d]", carddef.CardPersoRecordKeyCoBrandedName, i)
			return -1
		}
		embRecordInfo.CoBrandedName = persoRecordObject[carddef.CardPersoRecordKeyCoBrandedName].(string)

		_, ok = persoRecordObject[carddef.CardPersoRecordKeyTrack1]
		if !ok {
			//trace.Lg("[%s] not present in persoRecordObject[%d]", carddef.CardPersoRecordKeyTrack1, i)
			return -1
		}
		if !datatypeutil.IsString(persoRecordObject[carddef.CardPersoRecordKeyTrack1]) {
			//trace.Lg("[%s] not a string in persoRecordObject[%d]", carddef.CardPersoRecordKeyTrack1, i)
			return -1
		}
		embRecordInfo.Track1 = persoRecordObject[carddef.CardPersoRecordKeyTrack1].(string)

		_, ok = persoRecordObject[carddef.CardPersoRecordKeyTrack2]
		if !ok {
			//trace.Lg("[%s] not present in persoRecordObject[%d]", carddef.CardPersoRecordKeyTrack2, i)
			return -1
		}
		if !datatypeutil.IsString(persoRecordObject[carddef.CardPersoRecordKeyTrack2]) {
			//trace.Lg("[%s] not a string in persoRecordObject[%d]", carddef.CardPersoRecordKeyTrack2, i)
			return -1
		}
		embRecordInfo.Track2 = persoRecordObject[carddef.CardPersoRecordKeyTrack2].(string)

		_, ok = persoRecordObject[carddef.CardPersoRecordKeyCVV1]
		if !ok {
			//trace.Lg("[%s] not present in persoRecordObject[%d]", carddef.CardPersoRecordKeyCVV1, i)
			return -1
		}
		if !datatypeutil.IsString(persoRecordObject[carddef.CardPersoRecordKeyCVV1]) {
			//trace.Lg("[%s] not a string in persoRecordObject[%d]", carddef.CardPersoRecordKeyCVV1, i)
			return -1
		}
		embRecordInfo.CVV1 = persoRecordObject[carddef.CardPersoRecordKeyCVV1].(string)

		_, ok = persoRecordObject[carddef.CardPersoRecordKeyCVV2]
		if !ok {
			//trace.Lg("[%s] not present in persoRecordObject[%d]", carddef.CardPersoRecordKeyCVV2, i)
			return -1
		}
		if !datatypeutil.IsString(persoRecordObject[carddef.CardPersoRecordKeyCVV2]) {
			//trace.Lg("[%s] not a string in persoRecordObject[%d]", carddef.CardPersoRecordKeyCVV2, i)
			return -1
		}
		embRecordInfo.CVV2 = persoRecordObject[carddef.CardPersoRecordKeyCVV2].(string)

		_, ok = persoRecordObject[carddef.CardPersoRecordKeyICVV]
		if !ok {
			//trace.Lg("[%s] not present in persoRecordObject[%d]", carddef.CardPersoRecordKeyICVV, i)
			return -1
		}
		if !datatypeutil.IsString(persoRecordObject[carddef.CardPersoRecordKeyICVV]) {
			//trace.Lg("[%s] not a string in persoRecordObject[%d]", carddef.CardPersoRecordKeyICVV, i)
			return -1
		}
		embRecordInfo.ICVV = persoRecordObject[carddef.CardPersoRecordKeyICVV].(string)

		_, ok = persoRecordObject[carddef.CardPersoRecordKeyServiceCode]
		if !ok {
			//trace.Lg("[%s] not present in persoRecordObject[%d]", carddef.CardPersoRecordKeyServiceCode, i)
			return -1
		}
		if !datatypeutil.IsString(persoRecordObject[carddef.CardPersoRecordKeyServiceCode]) {
			//trace.Lg("[%s] not a string in persoRecordObject[%d]", carddef.CardPersoRecordKeyServiceCode, i)
			return -1
		}
		embRecordInfo.ServiceCode = persoRecordObject[carddef.CardPersoRecordKeyServiceCode].(string)

		_, ok = persoRecordObject[carddef.CardPersoRecordKeyCustomerName]
		if !ok {
			//trace.Lg("[%s] not present in persoRecordObject[%d]", carddef.CardPersoRecordKeyCustomerName, i)
			return -1
		}
		if !datatypeutil.IsString(persoRecordObject[carddef.CardPersoRecordKeyCustomerName]) {
			//trace.Lg("[%s] not a string in persoRecordObject[%d]", carddef.CardPersoRecordKeyCustomerName, i)
			return -1
		}
		embRecordInfo.CustomerName = persoRecordObject[carddef.CardPersoRecordKeyCustomerName].(string)

		_, ok = persoRecordObject[carddef.CardPersoRecordKeyCardIssDate]
		if !ok {
			//trace.Lg("[%s] not present in persoRecordObject[%d]", carddef.CardPersoRecordKeyCardIssDate, i)
			return -1
		}
		if !datatypeutil.IsString(persoRecordObject[carddef.CardPersoRecordKeyCardIssDate]) {
			//trace.Lg("[%s] not a string in persoRecordObject[%d]", carddef.CardPersoRecordKeyCardIssDate, i)
			return -1
		}
		embRecordInfo.CardIssDate = persoRecordObject[carddef.CardPersoRecordKeyCardIssDate].(string)
		embRecordInfo.ServerFlag = lServerFlag

		embRecordStr := ""
		rejectDesc := ""
		if cardutil.GetCardEmbRecord(embRecordInfo, &embRecordStr, &rejectDesc) < 0 {
			//trace.Lg("GetCardEmbRecord() failed for persoRecordObject[%d] with rejectDesc[%s]", i, rejectDesc)
			return -1
		}
		embRecordList = append(embRecordList, embRecordStr)

	}
	if writeCPFFile(embRecordList, lBIN, lCardSubProduct, lBatchID) < 0 {
		//trace.Lg("writeCPFFile() failed for BIN[%s] CardSubProduct[%s] BatchID[%s]", lBIN, lCardSubProduct, lBatchID)
		return -1
	}
	return 1
}

func GenPersoFileFunctionType(execFunction lexicalparserdef.ExecFunctionStruct, reqBrokerDataMap map[string]interface{}) int {
	var rejectDesc string
	//trace.Lg("GenPersoFileFunctionType() called")

	//trace.Lg("GetReqBrokerDataMapBuffer[%s]", reqbrokerutil.GetReqBrokerDataMapBuffer(reqBrokerDataMap))

	if len(execFunction.GenPersoFileInfo.PersoObject) == 0 {
		rejectDesc = fmt.Sprintf("PersoObject is NULL in GenPersoFileFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GenPersoFileFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenPersoFileInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenPersoFileInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	_, ok := reqBrokerDataMap[execFunction.GenPersoFileInfo.PersoObject]
	if !ok {
		rejectDesc = fmt.Sprintf("ObjectName[%s] does not exist in GenPersoFileFunctionType for FunctionName[%s]", execFunction.GenPersoFileInfo.PersoObject, execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GenPersoFileFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenPersoFileInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenPersoFileInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	tempType := fmt.Sprintf("%T", reqBrokerDataMap[execFunction.GenPersoFileInfo.PersoObject])
	if tempType != "map[string]interface {}" {
		rejectDesc = fmt.Sprintf("ObjectName[%s] is invalid in GenPersoFileFunctionType for FunctionName[%s]", execFunction.GenPersoFileInfo.PersoObject, execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GenPersoFileFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenPersoFileInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenPersoFileInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	if len(execFunction.GenPersoFileInfo.PersoObjectArray) == 0 {
		rejectDesc = fmt.Sprintf("PersoObjectArray is NULL in GenPersoFileFunctionType for FunctionName[%s]", execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GenPersoFileFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenPersoFileInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenPersoFileInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	_, ok = reqBrokerDataMap[execFunction.GenPersoFileInfo.PersoObjectArray]
	if !ok {
		rejectDesc = fmt.Sprintf("ObjectName[%s] does not exist in GenPersoFileFunctionType for FunctionName[%s]", execFunction.GenPersoFileInfo.PersoObjectArray, execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GenPersoFileFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenPersoFileInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenPersoFileInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	tempType = fmt.Sprintf("%T", reqBrokerDataMap[execFunction.GenPersoFileInfo.PersoObjectArray])
	//trace.Lg("tempType[%s]", tempType)
	if tempType != "[]interface {}" && tempType != "[]map[string]interface {}" {
		rejectDesc = fmt.Sprintf("ObjectName[%s] is invalid in GenPersoFileFunctionType for FunctionName[%s]", execFunction.GenPersoFileInfo.PersoObjectArray, execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GenPersoFileFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenPersoFileInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenPersoFileInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	if genPersoFile(reqBrokerDataMap, reqBrokerDataMap[execFunction.GenPersoFileInfo.PersoObject].(map[string]interface{}), reqBrokerDataMap[execFunction.GenPersoFileInfo.PersoObjectArray].([]map[string]interface{})) < 0 {
		rejectDesc = fmt.Sprintf("GenPersoFile() failed for PersoObject[%s] PersoObjectArray[%s] in GenPersoFileFunctionType for FunctionName[%s]", execFunction.GenPersoFileInfo.PersoObject, execFunction.GenPersoFileInfo.PersoObjectArray, execFunction.FunctionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_GenPersoFileFTErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenPersoFileInfo.AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.GenPersoFileInfo.AppCorrectiveAction), rejectDesc)
		return -1
	}
	return 1
}
