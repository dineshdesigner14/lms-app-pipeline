package lexicalfile_lms_course_item

func GetEditLMSCourseItemProgressServiceXML() interface{} {
	lexicalParserStr := `
	<sem_lexical_parser microservice="edit_lms_course_item_progress">
	<exec_group group_name="user_token_validation">
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
        <exec_function function_name="lms_course_item_progress_validate_fld" function_type="validate_fld">
			<validate_info app_err_desc="Validation failed for course item data" app_corrective_action="Please validate the data">
				<validate_fld fld_name="ReqBrokerReqObj.request_data.course_id" fld_data_type="string" fld_min_len="1" fld_max_len="255" fld_type="pattern" fld_pattern="^[0-9.]+$" app_err_desc="Invalid course id" app_corrective_action="Only numbers are allowed."></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.section_id" fld_data_type="string" fld_min_len="1" fld_max_len="255" fld_type="pattern" fld_pattern="^[0-9.]+$" app_err_desc="Invalid course section id" app_corrective_action="Only numbers are allowed."></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.item_id" fld_data_type="string" fld_min_len="1" fld_max_len="255" fld_type="pattern" fld_pattern="^[0-9.]+$" app_err_desc="Invalid item id" app_corrective_action="Only numbers are allowed."></validate_fld>
			    <validate_fld fld_name="ReqBrokerReqObj.request_data.is_course_content_completed" fld_data_type="boolean" fld_min_len="1" fld_max_len="10" fld_type="pattern" fld_pattern="^(true|false|0|1)$" app_err_desc="Invalid is required" app_corrective_action="Only Boolean value is allowed."></validate_fld>
		   </validate_info>
		</exec_function>
         <exec_function function_name="lms_course_item_progress_db_single_read" function_type="db_single_read">
			<db_single_read_info app_err_desc="lms course item progress data table is not found" app_corrective_action="course item progress data cannot be configured as the table does not exist">
				<table_name>LMS_COURSE_ITEM_PROGRESS</table_name>
				<filter_info>
					<filter_data filter_name="USER_ID" filter_id="decode_token_obj.user_id" filter_data_type="string" filter_condition="AND"></filter_data>
                    <filter_data filter_name="COURSE_ID" filter_id="ReqBrokerReqObj.request_data.course_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="SECTION_ID" filter_id="ReqBrokerReqObj.request_data.section_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="ITEM_ID" filter_id="ReqBrokerReqObj.request_data.item_id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_course_item_progress_table_record</store_obj_name>
			</db_single_read_info>
			<err_info>
				<err_desc err_code="NoRows" app_err_desc="No progress records found for the given course, section, and item for the specified user.." app_corrective_action="Verify that the user is enrolled in the course and that progress tracking has been initialized for the selected course item."/>
				<err_desc err_code="Duplicate" app_err_desc="Multiple progress records found for the given course item for the specified user.." app_corrective_action="Ensure that each course item has only one unique progress record." />
			</err_info>
		</exec_function>  
		<exec_function function_name="lms_course_item_copy_object" function_type="copy_object">
			<copy_object_info object_name="lms_course_item_table_obj">
				<object_data key="user_id" data_source_type="reqbrokermap" data_source="decode_token_obj.user_id" data_type="string"></object_data>
				<object_data key="course_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.course_id" data_type="string"></object_data>
				<object_data key="section_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.section_id" data_type="string"></object_data>
				<object_data key="item_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.item_id" data_type="string"></object_data>
				<object_data key="is_course_content_completed" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.is_course_content_completed" data_type="Boolean"></object_data>
				<object_data key="updated_by" data_source_type="reqbrokermap" data_source="decode_token_obj.email" data_type="string"></object_data>
				<object_data key="updated_date" data_source_type="key" data_source="get_date" data_type="string"></object_data>
				<object_data key="updated_time" data_source_type="key" data_source="get_time" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
		<exec_function function_name="edit_lms_course_item_progress_db_update" function_type="db_update">
			<db_update_info>
				<table_name>LMS_COURSE_ITEM_PROGRESS</table_name>
				<update_data_list>
					<update_data update_data_name="IS_COURSE_CONTENT_COMPLETED" update_db_data_type="boolean" update_data_source="lms_course_item_table_obj.is_course_content_completed" update_data_type="Boolean"></update_data>
					<update_data update_data_name="UPDATED_BY" update_db_data_type="string" update_data_source="lms_course_item_table_obj.updated_by" update_data_type="string"></update_data>
					<update_data update_data_name="UPDATED_DATE" update_db_data_type="date" update_data_source="lms_course_item_table_obj.updated_date" update_data_type="string"></update_data>
					<update_data update_data_name="UPDATED_TIME" update_db_data_type="string" update_data_source="lms_course_item_table_obj.updated_time" update_data_type="string"></update_data>
				</update_data_list>
				<update_filter_list>
					<update_filter update_filter_name="USER_ID" update_filter_db_data_type="string" update_filter_data_source="lms_course_item_table_obj.user_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
					<update_filter update_filter_name="COURSE_ID" update_filter_db_data_type="string" update_filter_data_source="lms_course_item_table_obj.course_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
					<update_filter update_filter_name="SECTION_ID" update_filter_db_data_type="string" update_filter_data_source="lms_course_item_table_obj.section_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
					<update_filter update_filter_name="ITEM_ID" update_filter_db_data_type="string" update_filter_data_source="lms_course_item_table_obj.item_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
				</update_filter_list>
			</db_update_info>
		</exec_function>
          <exec_function function_name="lms_course_item_progress_db_single_read_compose_response" function_type="db_single_read">
			<db_single_read_info app_err_desc="lms course item progress data table is not found" app_corrective_action="course item progress data cannot be configured as the table does not exist">
				<table_name>LMS_COURSE_ITEM_PROGRESS</table_name>
				<filter_info>
					<filter_data filter_name="USER_ID" filter_id="decode_token_obj.user_id" filter_data_type="string" filter_condition="AND"></filter_data>
                    <filter_data filter_name="COURSE_ID" filter_id="ReqBrokerReqObj.request_data.course_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="SECTION_ID" filter_id="ReqBrokerReqObj.request_data.section_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="ITEM_ID" filter_id="ReqBrokerReqObj.request_data.item_id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_course_item_progress_table_record</store_obj_name>
			</db_single_read_info>
			<err_info>
				<err_desc err_code="NoRows" app_err_desc="No progress records found for the given course, section, and item for the specified user.." app_corrective_action="Verify that the user is enrolled in the course and that progress tracking has been initialized for the selected course item."/>
				<err_desc err_code="Duplicate" app_err_desc="Multiple progress records found for the given course item for the specified user." app_corrective_action="Ensure that each course item has only one unique progress record." />
			</err_info>
		</exec_function>  
		<exec_function function_name="lms_course_item_copy_object_response" function_type="copy_object">
			<copy_object_info object_name="lms_course_item_progress_table_obj">
				<object_data key="user_id" data_source_type="reqbrokermap" data_source="lms_course_item_progress_table_record.user_id" data_type="string"></object_data>
				<object_data key="course_id" data_source_type="reqbrokermap" data_source="lms_course_item_progress_table_record.course_id" data_type="string"></object_data>
				<object_data key="section_id" data_source_type="reqbrokermap" data_source="lms_course_item_progress_table_record.section_id" data_type="string"></object_data>
				<object_data key="item_id" data_source_type="reqbrokermap" data_source="lms_course_item_progress_table_record.item_id" data_type="string"></object_data>
				<object_data key="is_course_content_completed" data_source_type="reqbrokermap" data_source="lms_course_item_progress_table_record.is_course_content_completed" data_type="Boolean"></object_data>
			</copy_object_info>
		</exec_function>
		<exec_function function_name="insert_lms_course_item_compose_response" function_type="compose_response">
			<response_info>
				<resp_obj obj_name="lms_course_item_progress_table_obj" obj_type="object" data_source="lms_course_item_progress_table_obj"></resp_obj>
			</response_info>
		</exec_function>
	</exec_group>
</sem_lexical_parser>
	`
	return lexicalParserStr
}
