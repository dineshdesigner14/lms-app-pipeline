package lexicalfile_lms_course

func GetViewLMSCourseProdServiceXML() interface{} {
	lexicalParserStr := `
	<sem_lexical_parser microservice="view_lms_course">
	<exec_group group_name="view_lms_course">
	<exec_function function_name="view_lms_course_create_object" function_type="create_object">
			<create_object_info object_name="view_lms_course_prod_temp_obj">
				<object_data key="status" data_source_type="raw_value" data_source="1" data_type="string"></object_data>
			</create_object_info>
		</exec_function>
		<exec_function function_name="view_lms_course_db_multi_read" function_type="db_multi_read">
			<db_multi_read_info>
				<table_name>LMS_COURSE</table_name>
				<filter_info>
					<filter_data filter_name="STATUS" filter_id="view_lms_course_prod_temp_obj.status" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_course_table_list</store_obj_name>
			</db_multi_read_info>
		</exec_function>
		<exec_function function_name="create_to_course_map_array" function_type="create_map_array">
			<map_array_info object_name="lms_course_array" array_size="len(lms_course_table_list)" />
		</exec_function>
		<exec_function function_name="lms_course_loop" function_type="start_loop" index_name="i" start_index="0" end_index="len(lms_course_table_list)" />
		<exec_function function_name="lms_course_copy_object" function_type="copy_object">
			<copy_object_info object_name="lms_course_array[i]">
				<object_data key="id" data_source_type="reqbrokermap" data_source="lms_course_table_list[i].id" data_type="string"></object_data>
				<object_data key="course_name" data_source_type="reqbrokermap" data_source="lms_course_table_list[i].course_name" data_type="string"></object_data>
 			</copy_object_info>
		</exec_function>
				<exec_function function_name="lms_course_loop" function_type="end_loop" index_name="i" />

		<exec_function function_name="view_lms_course_compose_response" function_type="compose_response">
			<response_info>
				<resp_obj obj_name="lms_course_array" obj_type="objectarray" data_source="lms_course_array"></resp_obj>
			</response_info>
		</exec_function>
	</exec_group>
</sem_lexical_parser>
	`
	return lexicalParserStr
}
