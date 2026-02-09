package lexicalfile_lms_course

func GetFetchLMSLaunchCourseServiceXML() interface{} {
	lexicalParserStr := `
	<sem_lexical_parser microservice="fetch_lms_launch_course">
	<exec_group group_name="view_fetch_lms_my_course">
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
		<exec_function function_name="insert_lms_course_validate_fld" function_type="validate_fld">
			<validate_info app_err_desc="Validation failed for course data" app_corrective_action="Please validate the data">
				<validate_fld fld_name="ReqBrokerReqObj.request_data.course_id" fld_data_type="string" fld_min_len="1" fld_max_len="255" fld_type="pattern" fld_pattern="^[a-zA-Z0-9.,|_* ]+$" ></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.section_id" fld_data_type="string" fld_min_len="1" fld_max_len="10" fld_type="pattern" fld_pattern="^[0-9.]+$"></validate_fld>
                <validate_fld fld_name="ReqBrokerReqObj.request_data.item_id" fld_data_type="string" fld_min_len="1" fld_max_len="155" fld_type="pattern" fld_pattern="^[a-zA-Z0-9.,|_* ]+$"></validate_fld>
			</validate_info>
		</exec_function>
			<exec_function function_name="lms_course_item_db_single_read" function_type="db_single_read">
			<db_single_read_info app_err_desc="lms course item table is not found" app_corrective_action="course item cannot be configured as the table does not exist">
				<table_name>LMS_COURSE</table_name>
				<filter_info>
					<filter_data filter_name="ID" filter_id="ReqBrokerReqObj.request_data.course_id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_course_table_record</store_obj_name>
			</db_single_read_info>
			<err_info>
				<err_desc err_code="Duplicate" app_err_desc="Duplicate records found for the title:{{.ReqBrokerReqObj.request_data.title}} provided in the input." app_corrective_action="Make sure each title shows up only once." />
			</err_info>
		</exec_function>
		<exec_function function_name="insert_lms_course_item_db_single_read" function_type="db_single_read">
			<db_single_read_info app_err_desc="lms course item table is not found" app_corrective_action="course item cannot be configured as the table does not exist">
				<table_name>LMS_COURSE_SECTION</table_name>
				<filter_info>
					<filter_data filter_name="COURSE_ID" filter_id="ReqBrokerReqObj.request_data.course_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="ID" filter_id="ReqBrokerReqObj.request_data.section_id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_course_section_table_record</store_obj_name>
			</db_single_read_info>
			<err_info>
				<err_desc err_code="Duplicate" app_err_desc="Duplicate records found for the title:{{.ReqBrokerReqObj.request_data.title}} provided in the input." app_corrective_action="Make sure each title shows up only once." />
			</err_info>
		</exec_function>
			<exec_function function_name="lms_course_item_db_single_read" function_type="db_single_read">
			<db_single_read_info app_err_desc="lms course item table is not found" app_corrective_action="course item cannot be configured as the table does not exist">
				<table_name>LMS_COURSE_ITEM</table_name>
				<filter_info>
					<filter_data filter_name="COURSE_ID" filter_id="ReqBrokerReqObj.request_data.course_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="SECTION_ID" filter_id="ReqBrokerReqObj.request_data.section_id" filter_data_type="string" filter_condition="AND"></filter_data>
                    <filter_data filter_name="ID" filter_id="ReqBrokerReqObj.request_data.item_id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_course_item_table_record</store_obj_name>
			</db_single_read_info>
			<err_info>
				<err_desc err_code="Duplicate" app_err_desc="Duplicate records found for the title:{{.ReqBrokerReqObj.request_data.title}} provided in the input." app_corrective_action="Make sure each title shows up only once." />
			</err_info>
		</exec_function>
		<exec_function function_name="course_item_response_copy_object" function_type="copy_object">
			<copy_object_info object_name="course_item_obj">
				<object_data key="course_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.course_id" data_type="string"></object_data>
				<object_data key="section_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.section_id" data_type="string"></object_data>
                <object_data key="item_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.item_id" data_type="string"></object_data>
                <object_data key="video_url" data_source_type="reqbrokermap" data_source="lms_course_item_table_record.video_url" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
		<exec_function function_name="view_lms_course_item_compose_response" function_type="compose_response">
			<response_info>
				<resp_obj obj_name="course_item_obj" obj_type="object" data_source="course_item_obj"></resp_obj>
			</response_info>
		</exec_function>
	</exec_group>
</sem_lexical_parser>
	`
	return lexicalParserStr
}
