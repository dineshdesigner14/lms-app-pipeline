package lexicaldef

import (
	lexicalfile_lms_course_section "lmsapieng/libsrc/lexicalfile/lms_course_section"
)

func (serviceRef *LexicalDefServiceInfo) GetInsertLMSCourseSectionXML() interface{} {
	return lexicalfile_lms_course_section.GetInsertLMSCourseSectionServiceXML()
}

func (serviceRef *LexicalDefServiceInfo) GetEditLMSCourseSectionXML() interface{} {
	return lexicalfile_lms_course_section.GetEditLMSCourseSectionServiceXML()
}

func (serviceRef *LexicalDefServiceInfo) GetDeleteLMSCourseSectionXML() interface{} {
	return lexicalfile_lms_course_section.GetDeleteLMSCourseSectionServiceXML()
}

func (serviceRef *LexicalDefServiceInfo) GetViewLMSCourseSectionXML() interface{} {
	return lexicalfile_lms_course_section.GetViewLMSCourseSectionServiceXML()
}

func (serviceRef *LexicalDefServiceInfo) GetViewLMSCourseSectionProdXML() interface{} {
	return lexicalfile_lms_course_section.GetViewLMSCourseSectionProdServiceXML()
}
func (serviceRef *LexicalDefServiceInfo) GetFetchLMSMyCourseSectionXML() interface{} {
	return lexicalfile_lms_course_section.GetFetchLMSMyCourseSectionServiceXML()
}
