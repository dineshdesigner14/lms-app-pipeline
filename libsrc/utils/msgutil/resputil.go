package msgutil

import (
	"encoding/json"
	"fmt"
	"lmsapieng/include/common/commondef"
	"lmsapieng/include/common/globaldef"
	"lmsapieng/include/common/moduledef"
	"lmsapieng/include/common/msgdef"
	"lmsapieng/include/common/rejectdef"
	"net/http"
	"strconv"
)

func SetRespWithData(RejectCode string, AppErrDesc []byte, AppCorrectiveAction []byte, RespData string, RejectArgs ...string) []byte {
	var rejectInfo rejectdef.RejectDefInfo
	RespInfoMap := make(map[string]interface{})
	if RejectCode == msgdef.RCapproved {
		RespInfoMap[msgdef.RespInfoStatusJSONObj] = http.StatusOK
		RespInfoMap[msgdef.RejectCodeJSONObj] = RejectCode
		RespInfoMap[msgdef.RespCodeJSONObj] = RejectCode
		RespInfoMap[msgdef.RespDescJSONObj] = msgdef.CoreDefineRespCodeListMap[RejectCode]
		RespInfoMap[msgdef.RejectShortDescJSONObj] = RespInfoMap[msgdef.RespDescJSONObj]
		RespInfoMap[msgdef.RejectLongDescJSONObj] = RespInfoMap[msgdef.RespDescJSONObj]
	} else {
		RespInfoMap[msgdef.RespInfoStatusJSONObj] = http.StatusInternalServerError
		RespInfoMap[msgdef.RejectCodeJSONObj] = RejectCode
		if rejectdef.GetLMSAdminServiceRejectInfo(RejectCode, &rejectInfo) < 0 {
			RespInfoMap[msgdef.RespCodeJSONObj] = msgdef.RCsystemerror
			RespInfoMap[msgdef.RespDescJSONObj] = msgdef.CoreDefineRespCodeListMap[msgdef.RCsystemerror]
			RespInfoMap[msgdef.RejectShortDescJSONObj] = globaldef.NOT_INITIALIZED
		} else {
			RespInfoMap[msgdef.RespCodeJSONObj] = rejectInfo.RespCode
			RespInfoMap[msgdef.RespDescJSONObj] = msgdef.CoreDefineRespCodeListMap[rejectInfo.RespCode]
			RespInfoMap[msgdef.RejectShortDescJSONObj] = rejectInfo.RejectDesc
		}
	}
	if len(RespData) != 0 {
		RespInfoMap[msgdef.RespDataJSONObj] = json.RawMessage(RespData)
	}
	if len(RejectArgs) > 0 {
		if len(RejectArgs[0]) != 0 {
			RespInfoMap[msgdef.RejectLongDescJSONObj] = RejectArgs[0]
		}
	}
	if len(RejectArgs) > 1 {
		if len(RejectArgs[1]) != 0 {
			RespInfoMap[msgdef.RejectSrcJSONObj] = RejectArgs[1]
		}
	}
	if len(RejectArgs) > 2 {
		if len(RejectArgs[2]) != 0 {
			RespInfoMap[msgdef.RejectModuleJSONObj] = RejectArgs[2]
		}
	}
	if len(RejectArgs) > 3 {
		if len(RejectArgs[3]) != 0 {
			RespInfoMap[msgdef.RejectFuncJSONObj] = RejectArgs[3]
		}
	}
	RespInfoMap[msgdef.RejectModuleTypeJSONObj] = commondef.ModuleTypeInternal
	RespInfoMap[msgdef.RejectModuleJSONObj] = moduledef.LMSApiEngModule
	RespInfoMap[msgdef.AppErrDescJSONObj] = string(AppErrDesc)
	RespInfoMap[msgdef.AppCorrectiveActionJSONObj] = string(AppCorrectiveAction)
	RespMap := make(map[string]interface{})
	RespMap[msgdef.RespInfoJSONObj] = RespInfoMap
	RespBuffer, _ := json.MarshalIndent(&RespMap, "", "\t")
	return RespBuffer
}

func SetResp(RejectCode string, AppErrDesc []byte, AppCorrectiveAction []byte, RejectArgs ...string) []byte {
	var rejectInfo rejectdef.RejectDefInfo
	RespInfoMap := make(map[string]interface{})
	if RejectCode == msgdef.RCapproved {
		RespInfoMap[msgdef.RespInfoStatusJSONObj] = http.StatusOK
		RespInfoMap[msgdef.RejectCodeJSONObj] = RejectCode
		RespInfoMap[msgdef.RespCodeJSONObj] = RejectCode
		RespInfoMap[msgdef.RespDescJSONObj] = msgdef.CoreDefineRespCodeListMap[RejectCode]
		RespInfoMap[msgdef.RejectShortDescJSONObj] = RespInfoMap[msgdef.RespDescJSONObj]
		RespInfoMap[msgdef.RejectLongDescJSONObj] = RespInfoMap[msgdef.RespDescJSONObj]
	} else {
		RespInfoMap[msgdef.RespInfoStatusJSONObj] = http.StatusInternalServerError
		RespInfoMap[msgdef.RejectCodeJSONObj] = RejectCode
		if rejectdef.GetLMSAdminServiceRejectInfo(RejectCode, &rejectInfo) < 0 {
			RespInfoMap[msgdef.RespCodeJSONObj] = msgdef.RCsystemerror
			RespInfoMap[msgdef.RespDescJSONObj] = msgdef.CoreDefineRespCodeListMap[msgdef.RCsystemerror]
			RespInfoMap[msgdef.RejectShortDescJSONObj] = globaldef.NOT_INITIALIZED
		} else {
			RespInfoMap[msgdef.RespCodeJSONObj] = rejectInfo.RespCode
			RespInfoMap[msgdef.RespDescJSONObj] = msgdef.CoreDefineRespCodeListMap[rejectInfo.RespCode]
			RespInfoMap[msgdef.RejectShortDescJSONObj] = rejectInfo.RejectDesc
		}
	}
	if len(RejectArgs) > 0 {
		if len(RejectArgs[0]) != 0 {
			RespInfoMap[msgdef.RejectLongDescJSONObj] = RejectArgs[0]
		}
	}
	if len(RejectArgs) > 1 {
		if len(RejectArgs[1]) != 0 {
			RespInfoMap[msgdef.RejectSrcJSONObj] = RejectArgs[1]
		}
	}
	if len(RejectArgs) > 2 {
		if len(RejectArgs[2]) != 0 {
			RespInfoMap[msgdef.RejectModuleJSONObj] = RejectArgs[2]
		}
	}
	if len(RejectArgs) > 3 {
		if len(RejectArgs[3]) != 0 {
			RespInfoMap[msgdef.RejectFuncJSONObj] = RejectArgs[3]
		}
	}
	RespInfoMap[msgdef.RejectModuleTypeJSONObj] = commondef.ModuleTypeInternal
	RespInfoMap[msgdef.RejectModuleJSONObj] = moduledef.LMSApiEngModule
	RespInfoMap[msgdef.AppErrDescJSONObj] = string(AppErrDesc)
	RespInfoMap[msgdef.AppCorrectiveActionJSONObj] = string(AppCorrectiveAction)
	RespMap := make(map[string]interface{})
	RespMap[msgdef.RespInfoJSONObj] = RespInfoMap
	RespBuffer, _ := json.MarshalIndent(&RespMap, "", "\t")
	return RespBuffer
}

func ParseResp(RespData []byte, respInfo *msgdef.RespInfoStruct) int {
	err := json.Unmarshal(RespData, respInfo)
	//trace.Lg("RespCode:%s", respInfo.RespInfo.RespCode)
	if err != nil {
		return -1
	}
	respcode, err := strconv.Atoi(respInfo.RespInfo.RespCode)
	if err != nil {
		return -1
	}
	if respcode != 0 {
		//fmt.Printf("\nGot Bad Response:%s", respInfo.RespInfo.RejectShortDesc)
		return -1
	}
	return 1
}

func ParseRespStatus(RespData []byte, respInfo *msgdef.RespInfoStruct) int {
	err := json.Unmarshal(RespData, respInfo)
	if err != nil {
		return -1
	}
	return respInfo.RespInfo.RespStatus
}

func GetNormalizeRespCode(RespCode string, RespCodeSize int) string {
	var lRespCode int
	if len(RespCode) > RespCodeSize {
		//trace.Lg("GetNormalizeRespCode() changing response since RespCodeSize > RespCodeSize")
		if RespCodeSize == 1 {
			lRespCode, _ = strconv.Atoi(msgdef.RCunabletoprocess)
		} else {
			lRespCode, _ = strconv.Atoi(msgdef.RCsystemerror)
		}
	} else {
		lRespCode, _ = strconv.Atoi(RespCode)
	}
	return fmt.Sprintf("%0*d", RespCodeSize, lRespCode)
}
