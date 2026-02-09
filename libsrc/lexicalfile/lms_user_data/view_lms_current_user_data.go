package lexicalfile_lms_user_data

func GetViewLMSCurrentUserDataServiceXML() interface{} {
	lexicalParserStr := `
	<sem_lexical_parser microservice="view_lms_current_user_data">
    <exec_group group_name="lms_current_user_data_validate_fld">
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
        <exec_function function_name="view_lms_current_user_data_db_single_read" function_type="db_single_read">
				<db_single_read_info app_err_desc="User data not found with UserID:{{.decode_token_obj.email}}" app_corrective_action="Please register and then try to log in.">
					<table_name>LMS_USER_DATA</table_name>
					<filter_info>
							<filter_data filter_name="ID" filter_id="decode_token_obj.user_id" filter_data_type="string" filter_condition="AND"></filter_data>
                            <filter_data filter_name="EMAIL" filter_id="decode_token_obj.email" filter_data_type="string" filter_condition="AND"></filter_data>
					</filter_info>
					<store_obj_name>lms_user_data_record</store_obj_name>
				</db_single_read_info>
				<err_info>
					<err_desc err_code="Duplicate" app_err_desc="Duplicate records found for the UserID:{{.decode_token_obj.email}}." app_corrective_action="Make sure each UserID shows up only once." />   
					<err_desc err_code="NoRows" app_err_desc="User data not found with UserID:{{.decode_token_obj.email}}" app_corrective_action="Please register and then try to log in." />
				</err_info>
			</exec_function>
			<exec_function function_name="view_lms_current_user_data_compose_response" function_type="compose_response">
				<response_info>
					<resp_obj obj_name="lms_user_data_record" obj_type="object" data_source="lms_user_data_record"></resp_obj>
				</response_info>
			</exec_function>
		</exec_group>
	</sem_lexical_parser>
	`
	return lexicalParserStr
}
