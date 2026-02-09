package lexicaldef

import lexicalfile_lms_question_bank "lmsapieng/libsrc/lexicalfile/lms_question_bank"

func (serviceRef *LexicalDefServiceInfo) GetInsertLMSQuestionBankXML() interface{} {
	return lexicalfile_lms_question_bank.GetInsertLMSQuestionBankServiceXML()
}

func (serviceRef *LexicalDefServiceInfo) GetEditLMSQuestionBankXML() interface{} {
	return lexicalfile_lms_question_bank.GetEditLMSQuestionBankServiceXML()
}

func (serviceRef *LexicalDefServiceInfo) GetViewLMSQuestionBankXML() interface{} {
	return lexicalfile_lms_question_bank.GetViewLMSQuestionBankServiceXML()
}

func (serviceRef *LexicalDefServiceInfo) GetFetchLMSQuestionBankXML() interface{} {
	return lexicalfile_lms_question_bank.GetFetchLMSQuestionBankServiceXML()
}

func (serviceRef *LexicalDefServiceInfo) GetFetchLMSLaunchQuizXML() interface{} {
	return lexicalfile_lms_question_bank.GetFetchLMSLaunchQuizServiceXML()
}

func (serviceRef *LexicalDefServiceInfo) GetViewLMSQuestionBankProdXML() interface{} {
	return lexicalfile_lms_question_bank.GetViewLMSQuestionBankProdServiceXML()
}

func (serviceRef *LexicalDefServiceInfo) GetHandleLMSQuizAttemptXML() interface{} {
	return lexicalfile_lms_question_bank.GetHandleLMSQuizAttemptServiceXML()
}

func (serviceRef *LexicalDefServiceInfo) GetLMSQuestionBankBulkUploadXML() interface{} {
	return lexicalfile_lms_question_bank.GetLMSQuestionBankBulkUploadServiceXML()
}
