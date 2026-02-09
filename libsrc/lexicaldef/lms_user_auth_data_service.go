package lexicaldef

import lexicalfile_lms_user_data "lmsapieng/libsrc/lexicalfile/lms_user_data"

func (serviceRef *LexicalDefServiceInfo) GetInsertLMSUserDataXML() interface{} {
	return lexicalfile_lms_user_data.GetInsertLMSUserDataServiceXML()
}

func (serviceRef *LexicalDefServiceInfo) GetViewLMSUserDataXML() interface{} {
	return lexicalfile_lms_user_data.GetViewLMSUserDataServiceXML()
}

func (serviceRef *LexicalDefServiceInfo) GetEditLMSUserDataXML() interface{} {
	return lexicalfile_lms_user_data.GetEditLMSUserDataServiceXML()
}

func (serviceRef *LexicalDefServiceInfo) GetViewLMSCurrentUserDataXML() interface{} {
	return lexicalfile_lms_user_data.GetViewLMSCurrentUserDataServiceXML()
}
