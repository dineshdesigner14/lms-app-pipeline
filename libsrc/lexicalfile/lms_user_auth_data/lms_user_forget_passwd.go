package lexicalfile_lms_user_auth_data

func GetLMSUserForgetPasswdServiceXML() interface{} {
	lexicalParserStr := `
	<sem_lexical_parser microservice="lms_user_forget_passwd">
		<exec_group group_name="lms_user_forget_passwd" group_condition="ReqBrokerReqObj.header_data.channel=='web_admin_portal'" >
		<exec_function function_name="lms_user_data_validate_fld" function_type="validate_fld">
			<validate_info app_err_desc="Validation failed for user data" app_corrective_action="Please validate the data">
				<validate_fld fld_name="ReqBrokerReqObj.request_data.user_id" fld_data_type="string" fld_min_len="1" fld_max_len="255" fld_type="pattern" fld_pattern="^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,3}$" app_err_desc="Invalid UserID" app_corrective_action="Only send valid email format."></validate_fld>
			</validate_info>
		</exec_function>
		<exec_function function_name="forget_passwd_lms_user_data_db_single_read" function_type="db_single_read">
			<db_single_read_info app_err_desc="User data not found with UserID:{{.ReqBrokerReqObj.request_data.user_id}}" app_corrective_action="Please register and then try to log in.">
				<table_name>LMS_USER_DATA</table_name>
				<filter_info>
					<filter_data filter_name="EMAIL" filter_id="ReqBrokerReqObj.request_data.user_id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_user_data_record</store_obj_name>
			</db_single_read_info>
			<err_info>
				<err_desc err_code="Duplicate" app_err_desc="Duplicate records found for the UserID:{{.ReqBrokerReqObj.request_data.user_id}}." app_corrective_action="Make sure each UserID shows up only once." />   
				<err_desc err_code="NoRows" app_err_desc="User data not found with UserID:{{.ReqBrokerReqObj.request_data.user_id}}" app_corrective_action="Please register and then try to log in." />
			</err_info>
		</exec_function>
		<exec_function function_name="lms_user_data_password_validate_condition" function_type="validate_condition">
			<validate_condition_info app_err_desc="Password retry exceeded" app_corrective_action="User account is blocked.">
				<validate_condition_expression>lms_user_data_record.passwd_status!='0'</validate_condition_expression>
			</validate_condition_info>
		</exec_function>
		<exec_function function_name="lms_user_data_password_validate_condition" function_type="validate_condition">
			<validate_condition_info app_err_desc="This user password is still system generated one!" app_corrective_action="Please reset the password using the one sent to your email.">
				<validate_condition_expression>lms_user_data_record.passwd_status!='2'</validate_condition_expression>
			</validate_condition_info>
		</exec_function>
		<exec_function function_name="reset_password_response_copy_object" function_type="copy_object">
			<copy_object_info object_name="reset_password_obj">
				<object_data key="status" data_source_type="raw_value" data_source="Verification code sent to emailID successfully!" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
		<exec_function function_name="forget_passwd_lms_user_data_compose_response" function_type="compose_response">
			<response_info>
				<resp_obj obj_name="reset_password_obj" obj_type="object" data_source="reset_password_obj"></resp_obj>
			</response_info>
		</exec_function>
		</exec_group>
	</sem_lexical_parser>
	`
	return lexicalParserStr
}
