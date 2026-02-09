package lexicalfile_lms_course_section

func GetEditLMSCourseSectionServiceXML() interface{} {
	lexicalParserStr := `
	<sem_lexical_parser microservice="edit_lms_course_section">
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
	<exec_group group_name="edit_lms_course_section" group_condition="ReqBrokerReqObj.header_data.channel=='web_admin_portal'">
		<exec_function function_name="edit_lms_course_section_validate_fld" function_type="validate_fld">
			<validate_info app_err_desc="Validation failed for course section data" app_corrective_action="Please validate the data">
				<validate_fld fld_name="ReqBrokerReqObj.request_data.id" fld_data_type="string" fld_min_len="1" fld_max_len="255" fld_type="pattern" fld_pattern="^[A-Za-z0-9 ]+$" app_err_desc="Invalid course section id" app_corrective_action="Only Numbers Allowed"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.course_id" fld_data_type="string" fld_min_len="1" fld_max_len="255" fld_type="pattern" fld_pattern="^[0-9.]+$" app_err_desc="Invalid course id" app_corrective_action="Only numbers are allowed."></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.title" fld_data_type="string" fld_min_len="1" fld_max_len="255" fld_type="pattern" fld_pattern=".*" app_err_desc="Invalid title" app_corrective_action="Only letters, numbers and spaces allowed."></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.description" fld_data_type="string" fld_min_len="1" fld_max_len="255" fld_type="pattern" fld_pattern=".*" app_err_desc="Invalid description" app_corrective_action="Only letters, numbers and spaces allowed."></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.order" fld_data_type="string" fld_min_len="1" fld_max_len="155" fld_type="pattern" fld_pattern="^[0-9.]+$" app_err_desc="Invalid order" app_corrective_action="Only numbers are allowed."></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.is_locked" fld_data_type="boolean" fld_min_len="1" fld_max_len="10" fld_type="pattern" fld_pattern="^(true|false|0|1)$" app_err_desc="Invalid is locked" app_corrective_action="Only Boolean value is allowed."></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.unlock_condition" fld_data_type="string" fld_min_len="1" fld_max_len="155" fld_type="pattern" fld_pattern="^(PREVIOUS_SECTION_COMPLETE|MANUAL)$" app_err_desc="Invalid unlock condition" app_corrective_action="Only PREVIOUS_SECTION_COMPLETE,MANUAL allowed"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.req_complete_percentage" fld_data_type="string" fld_min_len="1" fld_max_len="10" fld_type="pattern" fld_pattern="^[0-9.]+$" app_err_desc="Invalid required completion percentage" app_corrective_action="Only numbers will be allowed"></validate_fld>
			</validate_info>
		</exec_function>
		<exec_function function_name="edit_lms_course_section_db_single_read" function_type="db_single_read">
			<db_single_read_info app_err_desc="Course section order already exists" app_corrective_action="Please use a different section order and proceed.">
				<table_name>LMS_COURSE_SECTION</table_name>
				<filter_info>
					<filter_data filter_name="COURSE_ID" filter_id="ReqBrokerReqObj.request_data.course_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="ID" filter_id="ReqBrokerReqObj.request_data.id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="ORDER" filter_id="ReqBrokerReqObj.request_data.order" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_course_section_table_record</store_obj_name>
			</db_single_read_info>
			<err_info>
				<err_desc err_code="Duplicate" app_err_desc="Duplicate records found for the course_id:{{.ReqBrokerReqObj.request_data.course}} provided in the input." app_corrective_action="Make sure each course id shows up only once." />
			</err_info>
		</exec_function>
		<exec_function function_name="edit_lms_course_section_table_copy_object" function_type="copy_object">
			<copy_object_info object_name="lms_course_section_table_obj">
				<object_data key="id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.id" data_type="string"></object_data>
				<object_data key="course" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.course_id" data_type="string"></object_data>
				<object_data key="title" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.title" data_type="string"></object_data>
				<object_data key="description" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.description" data_type="string"></object_data>
				<object_data key="order" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.order" data_type="string" condition="lms_course_section_table_record.order!=ReqBrokerReqObj.request_data.order"></object_data>
				<object_data key="order" data_source_type="reqbrokermap" data_source="lms_course_section_table_record.order" data_type="string" condition="lms_course_section_table_record.order==ReqBrokerReqObj.request_data.order"></object_data>
				<object_data key="is_locked" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.is_locked" data_type="Boolean"></object_data>
				<object_data key="unlock_condition" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.unlock_condition" data_type="string"></object_data>
				<object_data key="req_complete_percent" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.req_complete_percentage" data_type="string"></object_data>
				<object_data key="updated_by" data_source_type="reqbrokermap" data_source="decode_token_obj.email" data_type="string"></object_data>
				<object_data key="updated_date" data_source_type="key" data_source="get_date" data_type="string"></object_data>
				<object_data key="updated_time" data_source_type="key" data_source="get_time" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
		<exec_function function_name="lms_course_section_info_db_update" function_type="db_update">
			<db_update_info>
				<table_name>LMS_COURSE_SECTION</table_name>
				<update_data_list>
					<update_data update_data_name="TITLE" update_db_data_type="string" update_data_source="lms_course_section_table_obj.title" update_data_type="string"></update_data>
					<update_data update_data_name="DESCRIPTION" update_db_data_type="string" update_data_source="lms_course_section_table_obj.description" update_data_type="string"></update_data>
					<update_data update_data_name="IS_LOCKED" update_db_data_type="boolean" update_data_source="lms_course_section_table_obj.is_locked" update_data_type="boolean"></update_data>
					<update_data update_data_name="UNLOCK_CONDITION" update_db_data_type="string" update_data_source="lms_course_section_table_obj.unlock_condition" update_data_type="string"></update_data>
					<update_data update_data_name="REQ_COMPLETE_PERCENTAGE" update_db_data_type="string" update_data_source="lms_course_section_table_obj.req_complete_percent" update_data_type="string"></update_data>
					<update_data update_data_name="UPDATED_BY" update_db_data_type="string" update_data_source="lms_course_section_table_obj.updated_by" update_data_type="string"></update_data>
					<update_data update_data_name="UPDATED_DATE" update_db_data_type="date" update_data_source="lms_course_section_table_obj.updated_date" update_data_type="string"></update_data>
					<update_data update_data_name="UPDATED_TIME" update_db_data_type="string" update_data_source="lms_course_section_table_obj.updated_time" update_data_type="string"></update_data>
				</update_data_list>
				<update_filter_list>
					<update_filter update_filter_name="COURSE_ID" update_filter_db_data_type="string" update_filter_data_source="ReqBrokerReqObj.request_data.course_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
					<update_filter update_filter_name="ID" update_filter_db_data_type="string" update_filter_data_source="ReqBrokerReqObj.request_data.id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
					<update_filter update_filter_name="ORDER" update_filter_db_data_type="string" update_filter_data_source="ReqBrokerReqObj.request_data.order" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
				</update_filter_list>
			</db_update_info>
		</exec_function>
		<exec_function function_name="insert_lms_course_section_db_single_read" function_type="db_single_read">
			<db_single_read_info app_err_desc="lms course data table is not found" app_corrective_action="lms course data cannot be configured as the table does not exist">
				<table_name>LMS_COURSE_SECTION</table_name>
				<filter_info>
					<filter_data filter_name="COURSE_ID" filter_id="ReqBrokerReqObj.request_data.course_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="ID" filter_id="ReqBrokerReqObj.request_data.id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_course_section_table_record</store_obj_name>
			</db_single_read_info>
		</exec_function>
		<exec_function function_name="course_section_response_copy_object" function_type="copy_object">
			<copy_object_info object_name="course_section_obj">
				<object_data key="status" data_source_type="raw_value" data_source="Course Section Edited Successfully!" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
		<exec_function function_name="insert_lms_course_section_compose_response" function_type="compose_response">
			<response_info>
				<resp_obj obj_name="course_section_obj" obj_type="object" data_source="course_section_obj"></resp_obj>
			</response_info>
		</exec_function>
	</exec_group>
</sem_lexical_parser>
	`
	return lexicalParserStr
}
