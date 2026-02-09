package lexicaldef

import lexicalfile_lms_course_item "lmsapieng/libsrc/lexicalfile/lms_course_item"

func (serviceRef *LexicalDefServiceInfo) GetInsertLMSCourseItemXML() interface{} {
	return lexicalfile_lms_course_item.GetInsertLMSCourseItemServiceXML()
}

func (serviceRef *LexicalDefServiceInfo) GetEditLMSCourseItemXML() interface{} {
	return lexicalfile_lms_course_item.GetEditLMSCourseItemServiceXML()
}

func (serviceRef *LexicalDefServiceInfo) GetDeleteLMSCourseItemXML() interface{} {
	return lexicalfile_lms_course_item.GetDeleteLMSCourseItemServiceXML()
}

func (serviceRef *LexicalDefServiceInfo) GetViewLMSCourseItemXML() interface{} {
	return lexicalfile_lms_course_item.GetViewLMSCourseItemServiceXML()
}

func (serviceRef *LexicalDefServiceInfo) GetViewLMSCourseItemProdXML() interface{} {
	return lexicalfile_lms_course_item.GetViewLMSCourseItemProdServiceXML()
}

func (serviceRef *LexicalDefServiceInfo) GetFetchLMSMyCourseItemXML() interface{} {
	return lexicalfile_lms_course_item.GetFetchLMSMyCourseItemServiceXML()
}
func (serviceRef *LexicalDefServiceInfo) GetEditLMSCourseItemProgressXML() interface{} {
	return lexicalfile_lms_course_item.GetEditLMSCourseItemProgressServiceXML()
}

func (serviceRef *LexicalDefServiceInfo) GetEditLMSCourseItemProgressCompleteXML() interface{} {
	return lexicalfile_lms_course_item.GetEditLMSCourseItemProgressCompleteServiceXML()
}
