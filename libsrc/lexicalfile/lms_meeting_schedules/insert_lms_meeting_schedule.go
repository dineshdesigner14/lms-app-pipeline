package lexicalfile_lms_meeting_schedules

func GetInsertLMSMeetingScheduleServiceXML() interface{} {
	lexicalParserStr := `
	<sem_lexical_parser microservice="insert_lms_meeting_schedule">
		<exec_group group_name="insert_lms_meeting_schedule" group_condition="ReqBrokerReqObj.header_data.channel=='web_client_portal'" >
		<exec_function function_name="decode_token_function_type" function_type="decode_token">
			<decode_token_info token_value="ReqBrokerReqObj.header_data.auth_token" token_object="decode_token_obj" app_err_desc="Decode token failed!. Invalid access token." app_corrective_action="Please provide valid access token."/>
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
		<exec_function function_name="lms_meeting_schedule_validate_fld" function_type="validate_fld">
			<validate_info app_err_desc="Validation failed for meeting_schedule" app_corrective_action="Please validate the data">
				<validate_fld fld_name="decode_token_obj.user_id" fld_data_type="string" fld_min_len="1" fld_max_len="255" fld_type="pattern" fld_pattern="^[0-9]+$" app_err_desc="Invalid student id" app_corrective_action="Only numbers are allowed."></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.course_id" fld_data_type="string" fld_min_len="1" fld_max_len="255" fld_type="pattern" fld_pattern="^[0-9]+$" app_err_desc="Pattern validation failed for course id" app_corrective_action="Only numbers are allowed."></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.meeting_date" fld_data_type="string" fld_min_len="5" fld_max_len="30" fld_type="pattern" fld_pattern="^(0[1-9]|[12][0-9]|3[01])-(0[1-9]|1[0-2])-\d{4}$" app_err_desc="Invalid meeting date format. Please enter the date in DD-MM-YYYY format." app_corrective_action="Enter a valid date in the format DD-MM-YYYY (for example, 06-01-2026)."></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.meeting_time" fld_data_type="string" fld_min_len="5" fld_max_len="30" fld_type="pattern" fld_pattern="^([01][0-9]|2[0-3]):[0-5][0-9]$" app_err_desc="Invalid meeting time format. Please enter the time in HH:MM (24-hour) format." app_corrective_action="Enter a valid time between 00:00 and 23:59."></validate_fld>
			</validate_info>
		</exec_function>
			
			<exec_function function_name="insert_lms_meeting_schedule_db_single_read" function_type="db_single_read">
				<db_single_read_info app_err_desc="lms meeting schedule table is not found" app_corrective_action="meeting schedule data cannot be configured as the table does not exist">
					<table_name>LMS_MEETING_SCHEDULE</table_name>
					<filter_info>
						<filter_data filter_name="user_id" filter_id="decode_token_obj.user_id" filter_data_type="string" filter_condition="AND"></filter_data>
						<filter_data filter_name="meeting_date" filter_id="ReqBrokerReqObj.request_data.meeting_date" filter_data_type="string" filter_condition="AND"></filter_data>
                        <filter_data filter_name="meeting_time" filter_id="ReqBrokerReqObj.request_data.meeting_time" filter_data_type="string" filter_condition="AND"></filter_data>
					</filter_info>
					<result_info app_err_desc="The user_id:{{.decode_token_obj.user_id}} already scheduled meeting with date {{.ReqBrokerReqObj.request_data.meeting_date}} and time {{.ReqBrokerReqObj.request_data.meeting_time}}." app_corrective_action="Please choose a different meeting date or time and try again.">
						<result_success>NoRows</result_success>
					</result_info>
				</db_single_read_info>
				<err_info>
					<err_desc err_code="Duplicate" app_err_desc="Duplicate records found for the user_id:{{.decode_token_obj.user_id}} already scheduled meeting with date {{.ReqBrokerReqObj.request_data.meeting_date}} and time {{.ReqBrokerReqObj.request_data.meeting_time}}." app_corrective_action="Please verify the details and avoid submitting duplicate entries." />   
				</err_info>
			</exec_function>

			<exec_function function_name="insert_lms_meeting_schedule_create_object" function_type="copy_object">
				<copy_object_info object_name="insert_lms_meeting_schedule_table_obj">
					<object_data key="user_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.user_id" data_type="string"></object_data>
					<object_data key="course_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.course_id" data_type="string"></object_data>
					<object_data key="meeting_date" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.meeting_date" data_type="string"></object_data>
					<object_data key="meeting_time" data_source_type="reqbrokermap" data_source="password_transform_obj.meeting_time" data_type="string"></object_data>
					<object_data key="meeting_link" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.meeting_link" data_type="string"></object_data>
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

			<exec_function function_name="insert_lms_meeting_schedule_db_insert" function_type="db_insert">
				<db_insert_info>
					<table_name>LMS_MEETING_SCHEDULE</table_name>
					<insert_data_list>
					    <insert_data insert_data_name="USER_ID" insert_db_data_type="string" insert_data_source="insert_lms_meeting_schedule_table_obj.user_id" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="COURSE_ID" insert_db_data_type="string" insert_data_source="insert_lms_meeting_schedule_table_obj.course_id" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="MEETING_DATE" insert_db_data_type="string" insert_data_source="insert_lms_meeting_schedule_table_obj.meeting_date" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="MEETING_TIME" insert_db_data_type="string" insert_data_source="insert_lms_meeting_schedule_table_obj.meeting_time" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="MEETING_LINK" insert_db_data_type="string" insert_data_source="insert_lms_meeting_schedule_table_obj.meeting_link" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="CREATED_BY" insert_db_data_type="string" insert_data_source="insert_lms_meeting_schedule_table_obj.created_by" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="CREATED_DATE" insert_db_data_type="date" insert_data_source="insert_lms_meeting_schedule_table_obj.created_date" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="CREATED_TIME" insert_db_data_type="string" insert_data_source="insert_lms_meeting_schedule_table_obj.created_time" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="UPDATED_BY" insert_db_data_type="string" insert_data_source="insert_lms_meeting_schedule_table_obj.updated_by" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="UPDATED_DATE" insert_db_data_type="date" insert_data_source="insert_lms_meeting_schedule_table_obj.updated_date" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="UPDATED_TIME" insert_db_data_type="string" insert_data_source="insert_lms_meeting_schedule_table_obj.updated_time" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="DELETED_BY" insert_db_data_type="string" insert_data_source="insert_lms_meeting_schedule_table_obj.deleted_by" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="DELETED_DATE" insert_db_data_type="date" insert_data_source="insert_lms_meeting_schedule_table_obj.deleted_date" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="DELETED_TIME" insert_db_data_type="string" insert_data_source="insert_lms_meeting_schedule_table_obj.deleted_time" insert_data_type="string"></insert_data>
						</insert_data_list>
				</db_insert_info>
			</exec_function>
        
			<exec_function function_name="insert_lms_meeting_schedule_create_object" function_type="copy_object" group_condition="len(lms_enquiry_table_list)==0">
				<copy_object_info object_name="meeting_response_obj">
				<object_data key="resp_msg" data_source_type="raw_value" data_source="Meeting scheduled successfully!!" data_type="string"></object_data>
			</copy_object_info>
			</exec_function>
			<exec_function function_name="insert_lms_meeting_schedule_compose_response" function_type="compose_response">
				<response_info>
					<resp_obj obj_name="meeting_response_obj" obj_type="object" data_source="meeting_response_obj"></resp_obj>
				</response_info>
			</exec_function>
		</exec_group>
	</sem_lexical_parser>
	`
	return lexicalParserStr
}
