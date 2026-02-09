package lexicalparser

import (
	"lmsapieng/include/common/dbdef"
	"lmsapieng/include/common/lexicalparserdef"
	"lmsapieng/include/common/rejectdef"
	"lmsapieng/include/common/reqbrokerdef"
	"lmsapieng/libsrc/utils/dbutil"
	"lmsapieng/libsrc/utils/msgutil"
	"lmsapieng/libsrc/utils/templateutil"
	"strings"
)

func handleDBErr(execFunction lexicalparserdef.ExecFunctionStruct, reqBrokerDataMap map[string]interface{}, dbErr *string, resultDBContext *dbdef.DBContextDef) {
	for i := 0; i < len(execFunction.ErrInfo.ErrDesc); i++ {
		if strings.EqualFold(execFunction.ErrInfo.ErrDesc[i].ErrCode, dbdef.DBNoRows) {
			if dbutil.IsNoRows(resultDBContext.DBType, *dbErr) {
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_DBErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ErrInfo.ErrDesc[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ErrInfo.ErrDesc[i].AppCorrectiveAction))
				break
			}
		} else if strings.EqualFold(execFunction.ErrInfo.ErrDesc[i].ErrCode, dbdef.DBDuplicateRows) {
			if dbutil.IsDuplicateRows(resultDBContext.DBType, *dbErr) {
				reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_DBErr, templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ErrInfo.ErrDesc[i].AppErrDesc), templateutil.GetTemplateStr(reqBrokerDataMap, execFunction.ErrInfo.ErrDesc[i].AppCorrectiveAction))
				break
			}
		}
	}
}
