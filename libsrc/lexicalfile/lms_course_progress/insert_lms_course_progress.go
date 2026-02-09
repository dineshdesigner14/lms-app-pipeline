package lexicalfile_lms_course_progress

func GetInsertLMSCourseProgressServiceXML() interface{} {
	lexicalParserStr := `
	<sem_lexical_parser microservice="insert_lms_course_progress">
	<exec_group group_name="insert_lms_course_progress" group_condition="ReqBrokerReqObj.header_data.channel=='web_admin_portal'">
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
		<exec_function function_name="insert_lms_course_progress_validate_fld" function_type="validate_fld">
			<validate_info app_err_desc="Validation failed for course data" app_corrective_action="Please validate the data">
				<validate_fld fld_name="ReqBrokerReqObj.request_data.user_id" fld_data_type="string" fld_min_len="1" fld_max_len="255" fld_type="pattern" fld_pattern="^[0-9.]+$" app_err_desc="Invalid user id" app_corrective_action="Only letters, numbers and spaces allowed."></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.course_id" fld_data_type="string" fld_min_len="1" fld_max_len="255" fld_type="pattern" fld_pattern="^[0-9.]+$" app_err_desc="Invalid course id" app_corrective_action="Only letters, numbers and spaces allowed."></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.percentage" fld_data_type="string" fld_min_len="1" fld_max_len="255" fld_type="pattern" fld_pattern="^[0-9.]+$" app_err_desc="Invalid percentage" app_corrective_action="Only letters, numbers and spaces allowed."></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.is_completed" fld_data_type="boolean" fld_min_len="1" fld_max_len="10" fld_type="pattern" fld_pattern="^(true|false|0|1)$" app_err_desc="Invalid Is_Completed" app_corrective_action="Only numbers will be allowed"></validate_fld>
                <validate_fld fld_name="ReqBrokerReqObj.request_data.time_spend_minutes" fld_data_type="string" fld_min_len="1" fld_max_len="155" fld_type="pattern" fld_pattern="^[0-9.]+$" app_err_desc="Invalid Time Spend In Minutes" app_corrective_action="Only numbers will be allowed"></validate_fld>
			</validate_info>
		</exec_function>
		<exec_function function_name="insert_lms_course_progress_single_read" function_type="db_single_read">
			<db_single_read_info app_err_desc="lms course progress table is not found" app_corrective_action="course data cannot be configured as the table does not exist">
				<table_name>LMS_COURSE_PROGRESS</table_name>
				<filter_info>
					<filter_data filter_name="USER_ID" filter_id="ReqBrokerReqObj.request_data.user_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="COURSE_ID" filter_id="ReqBrokerReqObj.request_data.course_id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<result_info app_err_desc="The course_id:{{.ReqBrokerReqObj.request_data.course_id}} already exists." app_corrective_action="Please configure new course id and proceed">
						<result_success>NoRows</result_success>
					</result_info>
				</db_single_read_info>
				<err_info>
					<err_desc err_code="Duplicate" app_err_desc="Duplicate records found for the course_id:{{.ReqBrokerReqObj.request_data.course_id}} provided in the input." app_corrective_action="Make sure each course id shows up only once." />   
				</err_info>
		</exec_function>
		
		<exec_function function_name="insert_lms_course_progress_copy_object" function_type="copy_object">
			<copy_object_info object_name="insert_lms_course_progress_additional_obj">
				<object_data key="user_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.user_id" data_type="string"></object_data>
				<object_data key="course_id"  data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.course_id" data_type="string"></object_data>
				<object_data key="percentage" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.percentage" data_type="string"></object_data>
				<object_data key="is_completed" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.is_completed" data_type="Boolean"></object_data>
				<object_data key="completed_at" data_source_type="key" data_source="get_time_stamp" data_type="string"></object_data>
				<object_data key="time_spend_minutes" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.time_spend_minutes" data_type="string"></object_data>
				<object_data key="last_accessed_at" data_source_type="reqbrokermap" data_source_type="key" data_source="get_time_stamp" data_type="string"></object_data>
				<object_data key="created_by" data_source_type="reqbrokermap" data_source="decode_token_obj.email" data_type="string"></object_data>
				<object_data key="created_date" data_source_type="key" data_source="get_date" data_type="string"></object_data>
				<object_data key="created_time" data_source_type="key" data_source="get_time" data_type="string"></object_data>
				<object_data key="updated_by" data_source_type="key" data_source="get_na" data_type="string"></object_data>
				<object_data key="updated_date" data_source_type="key" data_source="get_date" data_type="string"></object_data>
				<object_data key="updated_time" data_source_type="key" data_source="get_time" data_type="string"></object_data>
				<object_data key="deleted_by" data_source_type="key" data_source="get_na" data_type="string"></object_data>
				<object_data key="deleted_date" data_source_type="key" data_source="get_date" data_type="string"></object_data>
               <object_data key="deleted_time" data_source_type="key" data_source="get_time" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
		<exec_function function_name="insert_lms_course_progress_db_insert" function_type="db_insert">
			<db_insert_info>
				<table_name>LMS_COURSE_PROGRESS</table_name>
				<insert_data_list>
					<insert_data insert_data_name="USER_ID" insert_db_data_type="string" insert_data_source="insert_lms_course_progress_additional_obj.user_id" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="COURSE_ID" insert_db_data_type="string" insert_data_source="insert_lms_course_progress_additional_obj.course_id" insert_data_type="string"></insert_data>
                    <insert_data insert_data_name="PERCENTAGE" insert_db_data_type="string" insert_data_source="insert_lms_course_progress_additional_obj.percentage" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="IS_COMPLETED" insert_db_data_type="boolean" insert_data_source="ReqBrokerReqObj.request_data.is_completed" insert_data_type="Boolean"></insert_data>
					<insert_data insert_data_name="COMPLETED_AT" insert_db_data_type="string" insert_data_source="insert_lms_course_progress_additional_obj.completed_at" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="TIME_SPENT_MINUTES" insert_db_data_type="string" insert_data_source="insert_lms_course_progress_additional_obj.time_spend_minutes" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="LAST_ACCESSED_AT" insert_db_data_type="string" insert_data_source="insert_lms_course_progress_additional_obj.last_accessed_at" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="CREATED_BY" insert_db_data_type="string" insert_data_source="insert_lms_course_progress_additional_obj.created_by" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="CREATED_DATE" insert_db_data_type="date" insert_data_source="insert_lms_course_progress_additional_obj.created_date" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="CREATED_TIME" insert_db_data_type="string" insert_data_source="insert_lms_course_progress_additional_obj.created_time" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="UPDATED_BY" insert_db_data_type="string" insert_data_source="insert_lms_course_progress_additional_obj.updated_by" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="UPDATED_DATE" insert_db_data_type="date" insert_data_source="insert_lms_course_progress_additional_obj.updated_date" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="UPDATED_TIME" insert_db_data_type="string" insert_data_source="insert_lms_course_progress_additional_obj.updated_time" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="DELETED_BY" insert_db_data_type="string" insert_data_source="insert_lms_course_progress_additional_obj.deleted_by" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="DELETED_DATE" insert_db_data_type="date" insert_data_source="insert_lms_course_progress_additional_obj.deleted_date" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="DELETED_TIME" insert_db_data_type="string" insert_data_source="insert_lms_course_progress_additional_obj.deleted_time" insert_data_type="string"></insert_data>
				</insert_data_list>
			</db_insert_info>
		</exec_function>
		<exec_function function_name="insert_lms_course_progress_db_single_read" function_type="db_single_read">
			<db_single_read_info app_err_desc="lms course progress table is not found" app_corrective_action="lms course progress cannot be configured as the table does not exist">
				<table_name>LMS_COURSE_PROGRESS</table_name>
				<filter_info>
					<filter_data filter_name="USER_ID" filter_id="ReqBrokerReqObj.request_data.user_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="COURSE_ID" filter_id="ReqBrokerReqObj.request_data.course_id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_course_progress_table_record</store_obj_name>
			</db_single_read_info>
		</exec_function>
		<exec_function function_name="insert_lms_course_progress_compose_response" function_type="compose_response">
			<response_info>
				<resp_obj obj_name="lms_course_progress_table_record" obj_type="object" data_source="lms_course_progress_table_record"></resp_obj>
			</response_info>
		</exec_function>
	</exec_group>
</sem_lexical_parser>
	`
	return lexicalParserStr
}
