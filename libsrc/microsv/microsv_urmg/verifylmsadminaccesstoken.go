package microsv_urmg

import (
	"fmt"
	"lmsapieng/include/common/rejectdef"
	"lmsapieng/include/common/reqbrokerdef"
	"lmsapieng/include/common/urmgdef"
	"lmsapieng/libsrc/utils/msgutil"
	"lmsapieng/libsrc/utils/tokenutil"
	"lmsapieng/libsrc/utils/trace"
)

func VerifyLMSAdminAccessToken(reqBrokerDataMap map[string]interface{}) int {
	var rejectDesc, correctiveAction string
	verifyTokenBuffer := getVerifyAdmPortalAccessTokenBuffer(reqBrokerDataMap, urmgdef.VerifyTokenMicroService)
	trace.Lg("verifyTokenBuffer[%s]", verifyTokenBuffer)
	tokenVal, ok := reqBrokerDataMap["auth_token"].(string)
	if !ok || len(tokenVal) == 0 {
		rejectDesc = "auth_token missing in request"
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] =
			msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_VerifyLMSAdminAccessTokenFTErr, []byte(rejectDesc), []byte(correctiveAction))
		return -1
	}

	var payload map[string]interface{}
	ret := tokenutil.DecodeToken(tokenVal, &payload, &rejectDesc)
	if ret < 0 {
		rejectDesc = "Invalid or expired token"
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] =
			msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_VerifyLMSAdminAccessTokenFTErr, []byte(rejectDesc), []byte(correctiveAction))

		return -1
	}

	userIDVal, ok := payload["UserID"].(float64)
	if !ok {
		rejectDesc = "UserID missing in token claims"
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] =
			msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_VerifyLMSAdminAccessTokenFTErr, []byte(rejectDesc), []byte(correctiveAction))

		return -1
	}
	userID := int(userIDVal)

	email, ok := payload["Email"].(string)
	if !ok {
		rejectDesc = "Email missing in token claims"
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_VerifyLMSAdminAccessTokenFTErr, []byte(rejectDesc), []byte(correctiveAction))

		return -1
	}

	role, ok := payload["Role"].(string)
	if !ok {
		rejectDesc = "Role missing in token claims"
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_VerifyLMSAdminAccessTokenFTErr, []byte(rejectDesc), []byte(correctiveAction))

		return -1
	}

	if role != "ADMIN" {
		rejectDesc = fmt.Sprintf("Access denied. Role[%s] not authorized", role)
		reqBrokerDataMap[reqbrokerdef.ReqBrokerDataMapRespObj] = msgutil.SetResp(rejectdef.LMS_Admin_Service_Reject_VerifyLMSAdminAccessTokenFTErr, []byte(rejectDesc), []byte(correctiveAction))

		return -1
	}

	reqBrokerDataMap["AuthClaims"] = map[string]interface{}{
		"user_id": userID,
		"email":   email,
		"role":    role,
	}

	return 1
}
