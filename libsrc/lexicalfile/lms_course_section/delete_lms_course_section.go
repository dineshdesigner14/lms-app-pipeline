package lexicalfile_lms_course_section

func GetDeleteLMSCourseSectionServiceXML() interface{} {
	lexicalParserStr := `
	<sem_lexical_parser microservice="delete_lms_course_section">
	<exec_group group_name="user_token_validation" group_condition="ReqBrokerReqObj.header_data.channel=='web_admin_portal'">
		<exec_function function_name="decode_token_function_type" function_type="decode_token">
			<decode_token_info token_value="ReqBrokerReqObj.header_data.auth_token" token_object="decode_token_obj" app_err_desc="Decode token failed!. Invalid access token." app_corrective_action="Please provide valid access token."/>
		</exec_function>
		<exec_function function_name="role_validate_condition" function_type="validate_condition">
			<validate_condition_info app_err_desc="Invalid Role" app_corrective_action="Only ADMIN allowed to create User.">
				<validate_condition_expression>decode_token_obj.role=='ADMIN'||decode_token_obj.role=='SUPERADMIN'</validate_condition_expression>
			</validate_condition_info>
		</exec_function>
		<exec_function function_name="login_lms_token_guard_db_single_read" function_type="db_single_read">
			<db_single_read_info app_err_desc="Token guard data not found with Role:{{decode_token_obj.email}}" app_corrective_action="Please configure token guard for Role:{{decode_token_obj.role}}">
				<table_name>LMS_TOKEN_GUARD</table_name>
				<filter_info>
					<filter_data filter_name="ROLE" filter_id="decode_token_obj.role" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_token_guard_record</store_obj_name>
			</db_single_read_info>
			<err_info>
				<err_desc err_code="Duplicate" app_err_desc="Duplicate token guard data found for the Role:{{decode_token_obj.role}}." app_corrective_action="Make sure each Role shows up only once in table." />
				<err_desc err_code="NoRows" app_err_desc="Token guard data not found with Role:{{decode_token_obj.role}}" app_corrective_action="Please configure token guard for Role:{{decode_token_obj.role}}" />
			</err_info>
		</exec_function>
		<exec_function function_name="verify_token_function_type" function_type="verify_token">
			<verify_token_info token_value="ReqBrokerReqObj.header_data.auth_token" token_secret="lms_token_guard_record.access_token_secret" token_object="verify_token_obj" app_err_desc="Token Validation failed!. Invalid Access token." app_corrective_action="Please provide valid Access token."/>
		</exec_function>
	</exec_group>
	<exec_group group_name="delete_lms_course_section">
		<exec_function function_name="delete_lms_course_section_validate_fld" function_type="validate_fld">
			<validate_info>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.id" fld_data_type="string" fld_min_len="1" fld_max_len="10" fld_type="pattern" fld_pattern="^[0-9]+$"></validate_fld>
			</validate_info>
		</exec_function>
		<exec_function function_name="delete_lms_course_section_auth_db_single_read" function_type="db_single_read">
			<db_single_read_info>
				<table_name>LMS_COURSE_SECTION</table_name>
				<filter_info>
					<filter_data filter_name="ID" filter_id="ReqBrokerReqObj.request_data.id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_course_section_table_record</store_obj_name>
			</db_single_read_info>
		</exec_function>
		<exec_function function_name="delete_lms_course_section_copy_object" function_type="copy_object">
			<copy_object_info object_name="delete_lms_course_section_additional_obj">
				<object_data key="deleted_by" data_source_type="reqbrokermap" data_source="lms_course_section_table_record.created_by" data_type="string"></object_data>
				<object_data key="deleted_date" data_source_type="key" data_source="get_date" data_type="string"></object_data>
				<object_data key="deleted_time" data_source_type="key" data_source="get_time" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
		<exec_function function_name="delete_lms_course_section_db_update" function_type="db_update">
			<db_update_info>
				<table_name>LMS_COURSE_SECTION</table_name>
				<update_data_list>
					<update_data update_data_name="DELETED_BY" update_db_data_type="string" update_data_source="delete_lms_course_section_additional_obj.deleted_by" update_data_type="string"></update_data>
					<update_data update_data_name="DELETED_DATE" update_db_data_type="date" update_data_source="delete_lms_course_section_additional_obj.deleted_date" update_data_type="string"></update_data>
					<update_data update_data_name="DELETED_TIME" update_db_data_type="string" update_data_source="delete_lms_course_section_additional_obj.deleted_time" update_data_type="string"></update_data>
				</update_data_list>
				<update_filter_list>
					<update_filter update_filter_name="ID" update_filter_db_data_type="string" update_filter_data_source="ReqBrokerReqObj.request_data.id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
				</update_filter_list>
			</db_update_info>
		</exec_function>
		<exec_function function_name="insert_lms_course_section_db_single_read" function_type="db_single_read">
			<db_single_read_info app_err_desc="lms course section data table is not found" app_corrective_action="lms course section data cannot be configured as the table does not exist">
				<table_name>LMS_COURSE_SECTION</table_name>
				<filter_info>
					<filter_data filter_name="ID" filter_id="ReqBrokerReqObj.request_data.id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_course_section_table_record</store_obj_name>
			</db_single_read_info>
		</exec_function>
		<exec_function function_name="insert_lms_course_section_compose_response" function_type="compose_response">
			<response_info>
				<resp_obj obj_name="lms_course_section_table_record" obj_type="object" data_source="lms_course_section_table_record"></resp_obj>
			</response_info>
		</exec_function>
	</exec_group>
</sem_lexical_parser>
	`
	return lexicalParserStr
}
