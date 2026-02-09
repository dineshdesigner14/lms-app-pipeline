package lexicalfile_lms_course_item

func GetFetchLMSMyCourseItemServiceXML() interface{} {
	lexicalParserStr := `
	<sem_lexical_parser microservice="view_fetch_lms_my_course_item">
	<exec_group group_name="view_fetch_lms_my_course_item">
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
        <exec_function function_name="lms_course_item_validate_fld" function_type="validate_fld">
			<validate_info app_err_desc="Validation failed for course section data" app_corrective_action="Please validate the data">
				<validate_fld fld_name="ReqBrokerReqObj.request_data.course_id" fld_data_type="string" fld_min_len="1" fld_max_len="255" fld_type="pattern" fld_pattern="^[0-9.]+$" app_err_desc="Invalid course id" app_corrective_action="Only numbers are allowed."></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.section_id" fld_data_type="string" fld_min_len="1" fld_max_len="255" fld_type="pattern" fld_pattern="^[0-9.]+$" app_err_desc="Invalid course id" app_corrective_action="Only numbers are allowed."></validate_fld>
			</validate_info>
		</exec_function>	   
          <exec_function function_name="lms_enquiry_db_single_read" function_type="db_single_read">
			<db_single_read_info app_err_desc="lms enquiry table is not found" app_corrective_action="lms enquiry table cannot be configured as the table does not exist">
				<table_name>LMS_USER_DATA</table_name>
				<filter_info>
					<filter_data filter_name="EMAIL" filter_id="decode_token_obj.email" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_user_record</store_obj_name>
			</db_single_read_info>
			<err_info>
				<err_desc err_code="Duplicate" app_err_desc="Duplicate record found for email:{{decode_token_obj.email}}." app_corrective_action="Make sure email should be unique." />   
				<err_desc err_code="NoRows" app_err_desc="{{decode_token_obj.email}} not found" app_corrective_action="Please do enquiry for email:{{decode_token_obj.email}}" />
			</err_info>
		</exec_function>
		
	   <exec_function function_name="lms_course_db_single_read" function_type="db_single_read">
			<db_single_read_info app_err_desc="Data Already Exists or No Rows in Result Set" app_corrective_action="check the inputs and provide proper values">
				<table_name>LMS_COURSE</table_name>
				<filter_info>
					<filter_data filter_name="ID" filter_id="ReqBrokerReqObj.request_data.course_id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_course_table_record</store_obj_name>
			</db_single_read_info>
		</exec_function>
		  <exec_function function_name="read_lms_course_db_multi_read" function_type="db_multi_read">
			<db_multi_read_info app_err_desc="Data Already Exists or No Rows in Result Set" app_corrective_action="check the inputs and provide proper values">
				<table_name>LMS_COURSE_ITEM_PROGRESS</table_name>
				<filter_info>
					<filter_data filter_name="USER_ID" filter_id="decode_token_obj.user_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="COURSE_ID" filter_id="lms_course_table_record.id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="SECTION_ID" filter_id="ReqBrokerReqObj.request_data.section_id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_course_item_progress_list</store_obj_name>
			</db_multi_read_info>
		</exec_function>
		<exec_function function_name="create_to_cust_enq_map_array" function_type="create_map_array">
			<map_array_info object_name="lms_course_item_progress" array_size="len(lms_course_item_progress_list)" />
		</exec_function>
		<exec_function function_name="lms_item_progress_loop" function_type="start_loop" index_name="i" start_index="0" end_index="len(lms_course_item_progress_list)" />

		  <exec_function function_name="read_lms_course_db_multi_read" function_type="db_single_read">
			<db_single_read_info app_err_desc="Data Already Exists or No Rows in Result Set" app_corrective_action="check the inputs and provide proper values">
				<table_name>LMS_COURSE_ITEM</table_name>
				<filter_info>
					<filter_data filter_name="COURSE_ID" filter_id="lms_course_item_progress_list[i].course_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="SECTION_ID" filter_id="lms_course_item_progress_list[i].section_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="ID" filter_id="lms_course_item_progress_list[i].item_id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_course_item_record</store_obj_name>
			</db_single_read_info>
		</exec_function>
		<exec_function function_name="lms_course_item_copy_object" function_type="copy_object">
			<copy_object_info object_name="lms_course_item_progress[i]">
				<object_data key="course_id" data_source_type="reqbrokermap" data_source="lms_course_item_record.course_id" data_type="string"></object_data>
				<object_data key="section_id" data_source_type="reqbrokermap" data_source="lms_course_item_record.section_id" data_type="string"></object_data>
				<object_data key="item_id" data_source_type="reqbrokermap" data_source="lms_course_item_record.id" data_type="string"></object_data>
				<object_data key="title" data_source_type="reqbrokermap" data_source="lms_course_item_record.title" data_type="string"></object_data>
				<object_data key="description" data_source_type="reqbrokermap" data_source="lms_course_item_record.description" data_type="string"></object_data>
				<object_data key="order" data_source_type="reqbrokermap" data_source="lms_course_item_record.order" data_type="string"></object_data>
				<object_data key="is_completed" data_source_type="reqbrokermap" data_source="lms_course_item_progress_list[i].is_completed" data_type="Boolean"></object_data>
				<object_data key="is_enabled" data_source_type="reqbrokermap" data_source="lms_course_item_progress_list[i].is_enabled" data_type="Boolean"></object_data>
				<object_data key="is_course_content_completed" data_source_type="reqbrokermap" data_source="lms_course_item_progress_list[i].is_course_content_completed" data_type="Boolean"></object_data>
 			</copy_object_info>
		</exec_function>

		<exec_function function_name="lms_item_progress_loop" function_type="end_loop" index_name="i" />

		<exec_function function_name="view_lms_course_compose_response" function_type="compose_response">
			<response_info>
				<resp_obj obj_name="lms_course_item_progress" obj_type="objectarray" data_source="lms_course_item_progress"></resp_obj>
			</response_info>
		</exec_function>
	</exec_group>
</sem_lexical_parser>
	`
	return lexicalParserStr
}
