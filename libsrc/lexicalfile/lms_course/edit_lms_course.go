package lexicalfile_lms_course

func GetEditLMSCourseServiceXML() interface{} {
	lexicalParserStr := `
	<sem_lexical_parser microservice="edit_lms_course">
	<exec_group group_name="edit_lms_course" group_condition="ReqBrokerReqObj.header_data.channel=='web_admin_portal'">
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
		<exec_function function_name="edit_lms_course_validate_fld" function_type="validate_fld">
			<validate_info app_err_desc="Validation failed for course data" app_corrective_action="Please validate the data">
			    <validate_fld fld_name="ReqBrokerReqObj.request_data.id" fld_data_type="string" fld_min_len="1" fld_max_len="255" fld_type="pattern" fld_pattern="^[0-9.]+$" app_err_desc="Invalid course id" app_corrective_action="Only letters are allowed"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.course_name" fld_data_type="string" fld_min_len="1" fld_max_len="255" fld_type="pattern" fld_pattern=".*" app_err_desc="Invalid course name" app_corrective_action="Only letters, numbers and spaces allowed."></validate_fld>
				 <validate_fld fld_name="ReqBrokerReqObj.request_data.level" fld_data_type="string" fld_min_len="1" fld_max_len="155" fld_type="pattern" fld_pattern="^(BEGINNER|INTERMEDIATE|ADVANCED)$" app_err_desc="Invalid course level" app_corrective_action="Only BEGINNER,INTERMEDIATE,ADVANCED allowed"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.course_price" fld_data_type="string" fld_min_len="1" fld_max_len="10" fld_type="pattern" fld_pattern="^[0-9.]+$" app_err_desc="Invalid course price" app_corrective_action="Only numbers will be allowed"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.duration_min" fld_data_type="string" fld_min_len="1" fld_max_len="10" fld_type="pattern" fld_pattern="^[0-9]+$" app_err_desc="Invalid duration" app_corrective_action="Only numbers will be allowed"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.status" fld_data_type="string" fld_min_len="1" fld_max_len="5" fld_type="pattern" fld_pattern="^(true|false|0|1)$" app_err_desc="Invalid Status" app_corrective_action="Only true,false are allowed"></validate_fld>'
			</validate_info>
		</exec_function>
		<exec_function function_name="edit_lms_coursedb_single_read" function_type="db_single_read">
			<db_single_read_info>
				<table_name>LMS_COURSE</table_name>
				<filter_info>
					<filter_data filter_name="ID" filter_id="ReqBrokerReqObj.request_data.id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_course_auth_table_record</store_obj_name>
			</db_single_read_info>
			<err_info>
					<err_desc err_code="NoRows" app_err_desc="The course_id:{{.ReqBrokerReqObj.request_data.id}} not found." app_corrective_action="Please create the course and proceed"/>   
					<err_desc err_code="Duplicate" app_err_desc="The course_id:{{.ReqBrokerReqObj.request_data.id}} already exists." app_corrective_action="Please configure new course and proceed" />   
				</err_info>
		</exec_function>
		<exec_function function_name="edit_lms_course_copy_object" function_type="copy_object">
			<copy_object_info object_name="edit_lms_course_additional_obj">
				<object_data key="updated_by" data_source_type="reqbrokermap" data_source="decode_token_obj.email" data_type="string"></object_data>
				<object_data key="updated_date" data_source_type="key" data_source="get_date" data_type="string"></object_data>
				<object_data key="updated_time" data_source_type="key" data_source="get_time" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
		<exec_function function_name="edit_lms_course_create_object" function_type="create_object">
			<create_object_info object_name="edit_lms_course_table_obj">
				<object_data key="id" data_source="ReqBrokerReqObj.request_data.id" data_type="string"></object_data>
				<object_data key="course_name" data_source="ReqBrokerReqObj.request_data.course_name" data_type="string"></object_data>
				<object_data key="level" data_source="ReqBrokerReqObj.request_data.level" data_type="string"></object_data>
				<object_data key="course_price" data_source="ReqBrokerReqObj.request_data.course_price" data_type="string"></object_data>
				<object_data key="duration_min" data_source="ReqBrokerReqObj.request_data.duration_min" data_type="string"></object_data>
				<object_data key="status" data_source="ReqBrokerReqObj.request_data.status" data_type="string"></object_data>
				<object_data key="course_img" data_source="ReqBrokerReqObj.request_data.course_img" data_type="string"></object_data>
			</create_object_info>
		</exec_function>
		<exec_function function_name="edit_lms_course_db_update" function_type="db_update">
			<db_update_info>
				<table_name>LMS_COURSE</table_name>
				<update_data_list>
					<update_data update_data_name="COURSE_NAME" update_db_data_type="string" update_data_source="edit_lms_course_table_obj.course_name" update_data_type="string"></update_data>
					<update_data update_data_name="COURSE_PRICE" update_db_data_type="string" update_data_source="edit_lms_course_table_obj.course_price" update_data_type="string"></update_data>
                    <update_data update_data_name="LEVEL" update_db_data_type="string" update_data_source="edit_lms_course_table_obj.level" update_data_type="string"></update_data>
					<update_data update_data_name="DURATION_MIN" update_db_data_type="string" update_data_source="edit_lms_course_table_obj.duration_min" update_data_type="string"></update_data>
					<update_data update_data_name="STATUS" update_db_data_type="string" update_data_source="edit_lms_course_table_obj.status" update_data_type="string"></update_data>
					<update_data update_data_name="COURSE_IMG" update_db_data_type="string" update_data_source="edit_lms_course_table_obj.course_img" update_data_type="string"></update_data>
					<update_data update_data_name="UPDATED_BY" update_db_data_type="string" update_data_source="edit_lms_course_additional_obj.updated_by" update_data_type="string"></update_data>
					<update_data update_data_name="UPDATED_DATE" update_db_data_type="date" update_data_source="edit_lms_course_additional_obj.updated_date" update_data_type="string"></update_data>
					<update_data update_data_name="UPDATED_TIME" update_db_data_type="string" update_data_source="edit_lms_course_additional_obj.updated_time" update_data_type="string"></update_data>
				</update_data_list>
				<update_filter_list>
					<update_filter update_filter_name="ID" update_filter_db_data_type="string" update_filter_data_source="ReqBrokerReqObj.request_data.id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
				</update_filter_list>
			</db_update_info>
		</exec_function>
		<exec_function function_name="insert_lms_course_db_single_read" function_type="db_single_read">
			<db_single_read_info app_err_desc="lms course data table is not found" app_corrective_action="lms course data cannot be configured as the table does not exist">
				<table_name>LMS_COURSE</table_name>
				<filter_info>
					<filter_data filter_name="COURSE_NAME" filter_id="ReqBrokerReqObj.request_data.course_name" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_course_table_record</store_obj_name>
			</db_single_read_info>
		</exec_function>
		<exec_function function_name="insert_lms_course_compose_response" function_type="compose_response">
			<response_info>
				<resp_obj obj_name="lms_course_table_record" obj_type="object" data_source="lms_course_table_record"></resp_obj>
			</response_info>
		</exec_function>
	</exec_group>
</sem_lexical_parser>
	`
	return lexicalParserStr
}
