package lexicaldef

import lexicalfile_lms_course "lmsapieng/libsrc/lexicalfile/lms_course"

func (serviceRef *LexicalDefServiceInfo) GetViewLMSCourseXML() interface{} {
	return lexicalfile_lms_course.GetViewLMSCourseServiceXML()
}

func (serviceRef *LexicalDefServiceInfo) GetViewLMSCourseProdXML() interface{} {
	return lexicalfile_lms_course.GetViewLMSCourseProdServiceXML()
}

func (serviceRef *LexicalDefServiceInfo) GetInsertLMSCourseXML() interface{} {
	return lexicalfile_lms_course.GetInsertLMSCourseServiceXML()
}

func (serviceRef *LexicalDefServiceInfo) GetEditLMSCourseXML() interface{} {
	return lexicalfile_lms_course.GetEditLMSCourseServiceXML()
}

func (serviceRef *LexicalDefServiceInfo) GetDeleteLMSCourseXML() interface{} {
	return lexicalfile_lms_course.GetDeleteLMSCourseServiceXML()
}

func (serviceRef *LexicalDefServiceInfo) GetFetchLMSMyCourseSummaryXML() interface{} {
	return lexicalfile_lms_course.GetFetchLMSMyCourseSummaryServiceXML()
}

func (serviceRef *LexicalDefServiceInfo) GetFetchLMSMyCourseXML() interface{} {
	return lexicalfile_lms_course.GetFetchLMSMyCourseServiceXML()
}

func (serviceRef *LexicalDefServiceInfo) GetFetchLMSLaunchCourseXML() interface{} {
	return lexicalfile_lms_course.GetFetchLMSLaunchCourseServiceXML()
}

func (serviceRef *LexicalDefServiceInfo) GetViewLMSCourseStudentXML() interface{} {
	return lexicalfile_lms_course.GetViewLMSCourseStudentServiceXML()
}

func (serviceRef *LexicalDefServiceInfo) GetViewLMSCourseCurrentUserXML() interface{} {
	return lexicalfile_lms_course.GetViewLMSCourseCurrentUserServiceXML()
}

func (serviceRef *LexicalDefServiceInfo) GetLMSAssignCourseToStudentXML() interface{} {
	return lexicalfile_lms_course.GetLMSAssignCourseToStudentServiceXML()
}
