package lexicaldef

import lexicalfile_lms_user_auth_data "lmsapieng/libsrc/lexicalfile/lms_user_auth_data"

func (serviceRef *LexicalDefServiceInfo) GetLMSUserForgetPasswdXML() interface{} {
	return lexicalfile_lms_user_auth_data.GetLMSUserForgetPasswdServiceXML()
}

func (serviceRef *LexicalDefServiceInfo) GetLMSUserLoginXML() interface{} {
	return lexicalfile_lms_user_auth_data.GetLMSUserLoginServiceXML()
}

func (serviceRef *LexicalDefServiceInfo) GetLMSUserRenewTokenXML() interface{} {
	return lexicalfile_lms_user_auth_data.GetLMSUserRenewTokenServiceXML()
}

func (serviceRef *LexicalDefServiceInfo) GetLMSUserResetPasswdXML() interface{} {
	return lexicalfile_lms_user_auth_data.GetLMSUserResetPasswdServiceXML()
}

func (serviceRef *LexicalDefServiceInfo) GetLMSUserEmailVerifyXML() interface{} {
	return lexicalfile_lms_user_auth_data.GetLMSUserEmailVerifyServiceXML()
}
