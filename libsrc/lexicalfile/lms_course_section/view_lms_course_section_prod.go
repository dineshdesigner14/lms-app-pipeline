package lexicalfile_lms_course_section

func GetViewLMSCourseSectionProdServiceXML() interface{} {
	lexicalParserStr := `
	<sem_lexical_parser microservice="view_lms_course_section">
	<exec_group group_name="validate_view_lms_course_section"  group_condition="ReqBrokerReqObj.header_data.channel=='web_admin_portal'" >
		<exec_function function_name="read_view_lms_course_section_db_multi_read" function_type="db_multi_read">
			<db_multi_read_info>
				<table_name>LMS_COURSE_SECTION</table_name>
				<store_obj_name>lms_course_section_table_list</store_obj_name>
			</db_multi_read_info>
		</exec_function>
		<exec_function function_name="view_lms_course_section_compose_response" function_type="compose_response">
			<response_info>
				<resp_obj obj_name="lms_course_section_table_list" obj_type="objectarray" data_source="lms_course_section_table_list"></resp_obj>
			</response_info>
		</exec_function>
	</exec_group>
</sem_lexical_parser>
	`
	return lexicalParserStr
}
