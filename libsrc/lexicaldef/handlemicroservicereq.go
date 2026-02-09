package lexicaldef

import (
	"encoding/xml"
	"fmt"
	"lmsapieng/include/common/lexicalparserdef"
	"lmsapieng/include/common/rejectdef"
	"lmsapieng/include/common/reqbrokerdef"
	"lmsapieng/libsrc/lexicalparser"
	"lmsapieng/libsrc/utils/msgutil"
	"reflect"
)

type LexicalDefServiceInfo struct{}

func HandleMicroServiceReq(microServiceName string, reqBrokerDataMap map[string]interface{}) int {
	var serviceRef LexicalDefServiceInfo
	var parserInfo lexicalparserdef.LexicalParserStruct
	var err error
	var rejectDesc, correctiveAction string
	functionName := fmt.Sprintf("Get%sXML", microServiceName)
	method := reflect.ValueOf(&serviceRef).MethodByName(functionName)
	if !method.IsValid() {
		//trace.Lg("MicroService[%s] does not exist", functionName)
		rejectDesc = fmt.Sprintf("MicroService[%s] does not exist", functionName)
		correctiveAction = fmt.Sprintf("Please Include the MicroService[%s]", functionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_MicroServiceFunctionNotFoundErr, []byte(rejectDesc), []byte(correctiveAction), rejectDesc)
		return -2
	}
	retval := reflect.ValueOf(&serviceRef).MethodByName(functionName).Call(nil)
	if len(retval) == 0 {
		//trace.Lg("MicroService[%s] does not return anything", functionName)
		rejectDesc = fmt.Sprintf("MicroService[%s] does not return anything", functionName)
		correctiveAction = fmt.Sprintf("Ensure that MicroService[%s] returns proper value", functionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_MicroServiceFunctionErr, []byte(rejectDesc), []byte(correctiveAction), rejectDesc)
		return -1
	}
	resultInterface := retval[0].Interface()
	if resultInterface == nil {
		//trace.Lg("MicroService[%s] returns NULL", functionName)
		rejectDesc = fmt.Sprintf("MicroService[%s] returns NULL", functionName)
		correctiveAction = fmt.Sprintf("MicroService[%s] should not return NULL", functionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_MicroServiceFunctionErr, []byte(rejectDesc), []byte(correctiveAction), rejectDesc)
		return -1
	}
	lexicalParserStr := resultInterface.(string)
	err = xml.Unmarshal([]byte(lexicalParserStr), &parserInfo)
	if err != nil {
		//trace.Lg("xml.Unmarshal() failed with err[%s] for functionName[%s]", err, functionName)
		rejectDesc = fmt.Sprintf("XMLLexicalParseErr for [%s]", functionName)
		correctiveAction = fmt.Sprintf("Check the XML for Sytax Error in MicroService[%s]", functionName)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_XMLLexicalParseErr, []byte(rejectDesc), []byte(correctiveAction), rejectDesc)
		return -1
	}
	return lexicalparser.ExecLexicalParser(parserInfo, reqBrokerDataMap)
}
