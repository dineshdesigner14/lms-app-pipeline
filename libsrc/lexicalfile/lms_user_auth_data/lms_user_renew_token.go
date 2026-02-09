package lexicalfile_lms_user_auth_data

func GetLMSUserRenewTokenServiceXML() interface{} {
	lexicalParserStr := `
	<sem_lexical_parser microservice="lms_user_renew_token">
		<exec_group group_name="lms_user_renew_token" group_condition="ReqBrokerReqObj.header_data.channel=='web_admin_portal'" >
		<exec_function function_name="decode_token_function_type" function_type="decode_token">
			<decode_token_info token_value="ReqBrokerReqObj.request_data.auth_token" token_object="decode_token_obj" app_err_desc="Decode token failed!. Invalid access token." app_corrective_action="Please provide valid access token."/>
		</exec_function>
		<exec_function function_name="forget_passwd_lms_user_data_db_single_read" function_type="db_single_read">
			<db_single_read_info app_err_desc="User data not found with UserID:{{.decode_token_obj.email}}" app_corrective_action="Please register and then try to log in.">
				<table_name>LMS_USER_DATA</table_name>
				<filter_info>
					<filter_data filter_name="ID" filter_id="decode_token_obj.user_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="EMAIL" filter_id="decode_token_obj.email" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="ROLE" filter_id="decode_token_obj.role" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_user_data_record</store_obj_name>
			</db_single_read_info>
			<err_info>
				<err_desc err_code="Duplicate" app_err_desc="Duplicate records found for the UserID:{{.decode_token_obj.email}}." app_corrective_action="Make sure each UserID shows up only once." />   
				<err_desc err_code="NoRows" app_err_desc="User data not found with UserID:{{.decode_token_obj.email}}" app_corrective_action="Please register and then try to log in." />
			</err_info>
		</exec_function>
		<exec_function function_name="login_lms_token_guard_db_single_read" function_type="db_single_read">
			<db_single_read_info app_err_desc="Token guard data not found with Role:{{.decode_token_obj.email}}" app_corrective_action="Please configure token guard for Role:{{.lms_user_data_record.role}}">
				<table_name>LMS_TOKEN_GUARD</table_name>
				<filter_info>
					<filter_data filter_name="ROLE" filter_id="lms_user_data_record.role" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_token_guard_record</store_obj_name>
			</db_single_read_info>
			<err_info>
				<err_desc err_code="Duplicate" app_err_desc="Duplicate token guard data found for the Role:{{.lms_user_data_record.role}}." app_corrective_action="Make sure each Role shows up only once in table." />   
				<err_desc err_code="NoRows" app_err_desc="Token guard data not found with Role:{{.lms_user_data_record.role}}" app_corrective_action="Please configure token guard for Role:{{.lms_user_data_record.role}}" />
			</err_info>
		</exec_function>
		<exec_function function_name="verify_token_function_type" function_type="verify_token">
			<verify_token_info token_value="ReqBrokerReqObj.request_data.auth_token" token_secret="lms_token_guard_record.access_token_secret" token_object="verify_token_obj" app_err_desc="Token Validation failed!. Invalid Refresh token." app_corrective_action="Please provide valid refresh token."/>
		</exec_function>
		<exec_function function_name="mapp_token_info_copy_object" function_type="copy_object">
			<copy_object_info object_name="token_obj">
				<object_data key="token_issuer" data_source_type="raw_value" data_source="MindFlix" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
		<exec_function function_name="gen_access_token_function_type" function_type="gen_token" >
			<gen_token_info user_id="lms_user_data_record.id" email="lms_user_data_record.email" role="lms_user_data_record.role" token_issuer="token_obj.token_issuer"  token_secret="lms_token_guard_record.access_token_secret"  token_exp_unit="lms_token_guard_record.access_token_exp_unit"  token_expiry="lms_token_guard_record.access_token_exp" token_object="gen_access_token_obj" app_err_desc="Token Generation Failed" app_corrective_action="Please config token parameters properly"/>
		</exec_function>
		<exec_function function_name="gen_refresh_token_function_type" function_type="gen_token" >
			<gen_token_info user_id="lms_user_data_record.id" email="lms_user_data_record.email" role="lms_user_data_record.role" token_issuer="token_obj.token_issuer"  token_secret="lms_token_guard_record.refresh_token_secret"  token_exp_unit="lms_token_guard_record.refresh_token_exp_unit"  token_expiry="lms_token_guard_record.refresh_token_exp" token_object="gen_refresh_token_obj" app_err_desc="Token Generation Failed" app_corrective_action="Please config token parameters properly"/>
		</exec_function>
		<exec_function function_name="mapp_token_db_update" function_type="db_update">
			<db_update_info>
				<table_name>LMS_TOKEN_INFO</table_name>
				<update_data_list>
					<update_data update_data_name="ACCESS_TOKEN" update_db_data_type="string" update_data_source="gen_access_token_obj.token_value" update_data_type="string"></update_data>
					<update_data update_data_name="REFRESH_TOKEN" update_db_data_type="string" update_data_source="gen_refresh_token_obj.token_value" update_data_type="string"></update_data>					
				</update_data_list>
				<update_filter_list>
					<update_filter update_filter_name="USER_ID" update_filter_db_data_type="string" update_filter_data_source="decode_token_obj.email" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
				</update_filter_list>
			</db_update_info>
		</exec_function>
		<exec_function function_name="renew_token_response_copy_object" function_type="copy_object">
			<copy_object_info object_name="token_info">
				<object_data key="access_token" data_source_type="reqbrokermap" data_source="gen_access_token_obj.token_value" data_type="string"></object_data>
				<object_data key="refresh_token" data_source_type="reqbrokermap" data_source="gen_refresh_token_obj.token_value" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
		<exec_function function_name="renew_lms_user_data_compose_response" function_type="compose_response">
			<response_info>
				<resp_obj obj_name="token_info" obj_type="object" data_source="token_info"></resp_obj>
			</response_info>
		</exec_function>
		</exec_group>
	</sem_lexical_parser>
	`
	return lexicalParserStr
}
