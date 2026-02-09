package lexicaldef

import (
	lexicalfile_lms_course_progress "lmsapieng/libsrc/lexicalfile/lms_course_progress"
)

func (serviceRef *LexicalDefServiceInfo) GetInsertLMSCourseProgressXML() interface{} {
	return lexicalfile_lms_course_progress.GetInsertLMSCourseProgressServiceXML()
}

func (serviceRef *LexicalDefServiceInfo) GetEditLMSCourseProgressXML() interface{} {
	return lexicalfile_lms_course_progress.GetEditLMSCourseProgressServiceXML()
}

func (serviceRef *LexicalDefServiceInfo) GetViewLMSCourseProgressXML() interface{} {
	return lexicalfile_lms_course_progress.GetViewLMSCourseProgressServiceXML()
}

func (serviceRef *LexicalDefServiceInfo) GetViewLMSCourseProgressProdXML() interface{} {
	return lexicalfile_lms_course_progress.GetViewLMSCourseProgressProdServiceXML()
}
