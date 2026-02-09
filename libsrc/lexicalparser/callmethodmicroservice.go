package lexicalparser

import (
	"lmsapieng/libsrc/microsv/microsv_urmg"
	services_option "lmsapieng/libsrc/services/service_option"
	"lmsapieng/libsrc/services/services_calc"
	"lmsapieng/libsrc/services/services_email"
	"lmsapieng/libsrc/services/services_otp"
)

func (serviceClassRef *CallMethodFTServiceInfo) VerifyAdmPortalAccessToken(reqBrokerDataMap map[string]interface{}) int {
	return microsv_urmg.VerifyAdmPortalAccessToken(reqBrokerDataMap)
}
func (serviceClassRef *CallMethodFTServiceInfo) HandleOTPGen(reqBrokerDataMap map[string]interface{}) int {
	return services_otp.HandleOTPGen(reqBrokerDataMap)
}

func (serviceClassRef *CallMethodFTServiceInfo) HandleOTPVerify(reqBrokerDataMap map[string]interface{}) int {
	return services_otp.HandleOTPVerify(reqBrokerDataMap)
}

func (serviceClassRef *CallMethodFTServiceInfo) HandleOTPResend(reqBrokerDataMap map[string]interface{}) int {
	return services_otp.HandleOTPResend(reqBrokerDataMap)
}

func (serviceClassRef *CallMethodFTServiceInfo) VerifyLMSAdminAccessToken(reqBrokerDataMap map[string]interface{}) int {
	return microsv_urmg.VerifyLMSAdminAccessToken(reqBrokerDataMap)
}

func (serviceClassRef *CallMethodFTServiceInfo) HandleSendEmail(reqBrokerDataMap map[string]interface{}) int {
	return services_email.HandleSendEmail(reqBrokerDataMap)
}

func (serviceClassRef *CallMethodFTServiceInfo) HandleQuizScoreCalc(reqBrokerDataMap map[string]interface{}) int {
	return services_calc.HandleQuizScoreCalc(reqBrokerDataMap)
}

func (serviceClassRef *CallMethodFTServiceInfo) HandleQuizAverageCalc(reqBrokerDataMap map[string]interface{}) int {
	return services_calc.HandleQuizAverageCalc(reqBrokerDataMap)
}

func (serviceClassRef *CallMethodFTServiceInfo) BuildHasCorrectOptionFlag(reqBrokerDataMap map[string]interface{}) int {
	return services_option.BuildHasCorrectOptionFlag(reqBrokerDataMap)
}

func (serviceClassRef *CallMethodFTServiceInfo) HandleRandomQuestionShuffle(reqBrokerDataMap map[string]interface{}) int {
	return services_option.HandleRandomQuestionShuffle(reqBrokerDataMap)
}
