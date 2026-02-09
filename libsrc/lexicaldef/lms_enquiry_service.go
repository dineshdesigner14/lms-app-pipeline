package lexicaldef

import lexicalfile_lms_enquiry "lmsapieng/libsrc/lexicalfile/lms_enquiry"

func (serviceRef *LexicalDefServiceInfo) GetInsertLMSEnquiryXML() interface{} {
	return lexicalfile_lms_enquiry.GetInsertLMSEnquiryServiceXML()
}

func (serviceRef *LexicalDefServiceInfo) GetViewLMSEnquiryXML() interface{} {
	return lexicalfile_lms_enquiry.GetViewLMSEnquiryServiceXML()
}

func (serviceRef *LexicalDefServiceInfo) GetEditLMSEnquiryXML() interface{} {
	return lexicalfile_lms_enquiry.GetEditLMSEnquiryServiceXML()
}

func (serviceRef *LexicalDefServiceInfo) GetUpdateLMSEnquiryCommentXML() interface{} {
	return lexicalfile_lms_enquiry.GetUpdateLMSEnquiryCommentServiceXML()
}
func (serviceRef *LexicalDefServiceInfo) GetViewLMSEnquiryCommentXML() interface{} {
	return lexicalfile_lms_enquiry.GetViewLMSEnquiryCommentServiceXML()
}

func (serviceRef *LexicalDefServiceInfo) GetLMSEnquiryEmailVerifyXML() interface{} {
	return lexicalfile_lms_enquiry.GetLMSEnquiryEmailVerifyServiceXML()
}

func (serviceRef *LexicalDefServiceInfo) GetLMSEnquiryResendOTPXML() interface{} {
	return lexicalfile_lms_enquiry.GetLMSEnquiryResendOTPServiceXML()
}

func (serviceRef *LexicalDefServiceInfo) GetHandleLMSEnquiryCourseXML() interface{} {
	return lexicalfile_lms_enquiry.GetHandleLMSEnquiryCourseServiceXML()
}
