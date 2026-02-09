package lexicalfile_lms_user_auth_data

func GetLMSUserLoginServiceXML() interface{} {
	lexicalParserStr := `
	<sem_lexical_parser microservice="lms_user_login">
		<exec_group group_name="lms_admin_user_login" group_condition="ReqBrokerReqObj.header_data.channel=='web_admin_portal'" >
		<exec_function function_name="lms_user_data_validate_fld" function_type="validate_fld">
			<validate_info app_err_desc="Validation failed for user data" app_corrective_action="Please validate the data">
				<validate_fld fld_name="ReqBrokerReqObj.request_data.user_id" fld_data_type="string" fld_min_len="1" fld_max_len="255" fld_type="pattern" fld_pattern="^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,3}$" app_err_desc="Invalid UserID" app_corrective_action="Only send valid email format."></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.password" fld_data_type="string" fld_min_len="5" fld_max_len="255" fld_type="pattern" fld_pattern=".*" app_err_desc="Password is not secure" app_corrective_action="Password must have atleast one captial case, small case, number and special character"></validate_fld>
			</validate_info>
		</exec_function>
		<exec_function function_name="login_lms_user_data_db_single_read" function_type="db_single_read">
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
		 <exec_function function_name="role_validate_condition" function_type="validate_condition">
			<validate_condition_info app_err_desc="Invalid User Role" app_corrective_action="Only ADMIN or SUPERADMIN allowed to login.">
				<validate_condition_expression>lms_user_data_record.role=='ADMIN'||lms_user_data_record.role=='SUPERADMIN'</validate_condition_expression>
			</validate_condition_info>
		</exec_function>
		<exec_function function_name="lms_user_data_password_transform" function_type="transform_object">
			<transform_object_info object_name="password_transform_obj">
				<object_data algo="hash" key="hash_password" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.password" data_type="string" />
			</transform_object_info>
		</exec_function>
		<exec_function function_name="lms_user_data_password_validate_condition" function_type="validate_condition">
			<validate_condition_info app_err_desc="Password retry exceeded" app_corrective_action="User account is blocked.">
				<validate_condition_expression>lms_user_data_record.passwd_status!='0'</validate_condition_expression>
			</validate_condition_info>
		</exec_function>
		<exec_function function_name="lms_user_data_password_validate_condition" function_type="validate_condition">
			<validate_condition_info app_err_desc="Invalid Password" app_corrective_action="Kindly enter a valid password.">
				<validate_condition_expression>password_transform_obj.hash_password==lms_user_data_record.password_hash</validate_condition_expression>
			</validate_condition_info>
		</exec_function>
		<exec_function function_name="login_lms_token_guard_db_single_read" function_type="db_single_read">
			<db_single_read_info app_err_desc="Token guard data not found with Role:{{.lms_user_data_record.role}}" app_corrective_action="Please configure token guard for Role:{{.lms_user_data_record.role}}">
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
		<exec_function function_name="mapp_token_info_copy_object" function_type="copy_object">
			<copy_object_info object_name="token_obj">
				<object_data key="token_issuer" data_source_type="raw_value" data_source="MindFlix" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
		<exec_function function_name="gen_access_token_function_type" function_type="gen_token" >
			<gen_token_info user_id="lms_user_data_record.id" email="ReqBrokerReqObj.request_data.user_id" role="lms_user_data_record.role" token_issuer="token_obj.token_issuer"  token_secret="lms_token_guard_record.access_token_secret"  token_exp_unit="lms_token_guard_record.access_token_exp_unit"  token_expiry="lms_token_guard_record.access_token_exp" token_object="gen_access_token_obj" app_err_desc="Token Generation Failed" app_corrective_action="Please config token parameters properly"/>
		</exec_function>
		<exec_function function_name="gen_refresh_token_function_type" function_type="gen_token" >
			<gen_token_info user_id="lms_user_data_record.id" email="ReqBrokerReqObj.request_data.user_id" role="lms_user_data_record.role" token_issuer="token_obj.token_issuer"  token_secret="lms_token_guard_record.refresh_token_secret"  token_exp_unit="lms_token_guard_record.refresh_token_exp_unit"  token_expiry="lms_token_guard_record.refresh_token_exp" token_object="gen_refresh_token_obj" app_err_desc="Token Generation Failed" app_corrective_action="Please config token parameters properly"/>
		</exec_function>
		<exec_function function_name="mapp_token_db_update" function_type="db_update">
			<db_update_info>
				<table_name>LMS_TOKEN_INFO</table_name>
				<update_data_list>
					<update_data update_data_name="ACCESS_TOKEN" update_db_data_type="string" update_data_source="gen_access_token_obj.token_value" update_data_type="string"></update_data>
					<update_data update_data_name="REFRESH_TOKEN" update_db_data_type="string" update_data_source="gen_refresh_token_obj.token_value" update_data_type="string"></update_data>					
				</update_data_list>
				<update_filter_list>
					<update_filter update_filter_name="USER_ID" update_filter_db_data_type="string" update_filter_data_source="ReqBrokerReqObj.request_data.user_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
				</update_filter_list>
			</db_update_info>
		</exec_function>
		<exec_function function_name="login_response_copy_object" function_type="copy_object">
			<copy_object_info object_name="login_resp_obj">
				<object_data key="full_name" data_source_type="reqbrokermap" data_source="lms_user_data_record.full_name" data_type="string"></object_data>
				<object_data key="email" data_source_type="reqbrokermap" data_source="lms_user_data_record.email" data_type="string"></object_data>
				<object_data key="role" data_source_type="reqbrokermap" data_source="lms_user_data_record.role" data_type="string"></object_data>
				<object_data key="passwd_status" data_source_type="reqbrokermap" data_source="lms_user_data_record.passwd_status" data_type="string"></object_data>
				<object_data key="last_login_date" data_source_type="reqbrokermap" data_source="lms_user_data_record.last_login_date" data_type="string"></object_data>
				<object_data key="last_login_time" data_source_type="reqbrokermap" data_source="lms_user_data_record.last_login_time" data_type="string"></object_data>
				<object_data key="profile_picture_url" data_source_type="reqbrokermap" data_source="lms_user_data_record.profile_picture_url" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
		<exec_function function_name="login_token_response_copy_object" function_type="copy_object">
			<copy_object_info object_name="login_resp_obj.token_info">
				<object_data key="access_token" data_source_type="reqbrokermap" data_source="gen_access_token_obj.token_value" data_type="string"></object_data>
				<object_data key="refresh_token" data_source_type="reqbrokermap" data_source="gen_refresh_token_obj.token_value" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
		<exec_function function_name="login_lms_user_data_compose_response" function_type="compose_response">
			<response_info>
				<resp_obj obj_name="login_resp_obj" obj_type="object" data_source="login_resp_obj"></resp_obj>
			</response_info>
		</exec_function>
		</exec_group>
		<exec_group group_name="lms_student_user_login" group_condition="ReqBrokerReqObj.header_data.channel=='web_client_portal'" >
		<exec_function function_name="lms_user_data_validate_fld" function_type="validate_fld">
			<validate_info app_err_desc="Validation failed for user data" app_corrective_action="Please validate the data">
				<validate_fld fld_name="ReqBrokerReqObj.request_data.user_id" fld_data_type="string" fld_min_len="1" fld_max_len="255" fld_type="pattern" fld_pattern="^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,3}$" app_err_desc="Invalid UserID" app_corrective_action="Only send valid email format."></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.password" fld_data_type="string" fld_min_len="5" fld_max_len="255" fld_type="pattern" fld_pattern=".*" app_err_desc="Password is not secure" app_corrective_action="Password must have atleast one captial case, small case, number and special character"></validate_fld>
			</validate_info>
		</exec_function>
		<exec_function function_name="login_lms_user_data_db_single_read" function_type="db_single_read">
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
		 <exec_function function_name="role_validate_condition" function_type="validate_condition">
			<validate_condition_info app_err_desc="Invalid User Role" app_corrective_action="Only STUDENT allowed to login.">
				<validate_condition_expression>lms_user_data_record.role=='STUDENT'</validate_condition_expression>
			</validate_condition_info>
		</exec_function>
		<exec_function function_name="lms_user_data_password_transform" function_type="transform_object">
			<transform_object_info object_name="password_transform_obj">
				<object_data algo="hash" key="hash_password" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.password" data_type="string" />
			</transform_object_info>
		</exec_function>
		<exec_function function_name="lms_user_data_password_validate_condition" function_type="validate_condition">
			<validate_condition_info app_err_desc="Password retry exceeded" app_corrective_action="User account is blocked.">
				<validate_condition_expression>lms_user_data_record.passwd_status!='0'</validate_condition_expression>
			</validate_condition_info>
		</exec_function>
		<exec_function function_name="lms_user_data_password_validate_condition" function_type="validate_condition">
			<validate_condition_info app_err_desc="Invalid Password" app_corrective_action="Kindly enter a valid password.">
				<validate_condition_expression>password_transform_obj.hash_password==lms_user_data_record.password_hash</validate_condition_expression>
			</validate_condition_info>
		</exec_function>
		<exec_function function_name="login_lms_token_guard_db_single_read" function_type="db_single_read">
			<db_single_read_info app_err_desc="Token guard data not found with Role:{{.lms_user_data_record.role}}" app_corrective_action="Please configure token guard for Role:{{.lms_user_data_record.role}}">
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
		<exec_function function_name="mapp_token_info_copy_object" function_type="copy_object">
			<copy_object_info object_name="token_obj">
				<object_data key="token_issuer" data_source_type="raw_value" data_source="MindFlix" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
		<exec_function function_name="gen_access_token_function_type" function_type="gen_token" >
			<gen_token_info user_id="lms_user_data_record.id" email="ReqBrokerReqObj.request_data.user_id" role="lms_user_data_record.role" token_issuer="token_obj.token_issuer"  token_secret="lms_token_guard_record.access_token_secret"  token_exp_unit="lms_token_guard_record.access_token_exp_unit"  token_expiry="lms_token_guard_record.access_token_exp" token_object="gen_access_token_obj" app_err_desc="Token Generation Failed" app_corrective_action="Please config token parameters properly"/>
		</exec_function>
		<exec_function function_name="gen_refresh_token_function_type" function_type="gen_token" >
			<gen_token_info user_id="lms_user_data_record.id" email="ReqBrokerReqObj.request_data.user_id" role="lms_user_data_record.role" token_issuer="token_obj.token_issuer"  token_secret="lms_token_guard_record.refresh_token_secret"  token_exp_unit="lms_token_guard_record.refresh_token_exp_unit"  token_expiry="lms_token_guard_record.refresh_token_exp" token_object="gen_refresh_token_obj" app_err_desc="Token Generation Failed" app_corrective_action="Please config token parameters properly"/>
		</exec_function>
		<exec_function function_name="mapp_token_db_update" function_type="db_update">
			<db_update_info>
				<table_name>LMS_TOKEN_INFO</table_name>
				<update_data_list>
					<update_data update_data_name="ACCESS_TOKEN" update_db_data_type="string" update_data_source="gen_access_token_obj.token_value" update_data_type="string"></update_data>
					<update_data update_data_name="REFRESH_TOKEN" update_db_data_type="string" update_data_source="gen_refresh_token_obj.token_value" update_data_type="string"></update_data>					
				</update_data_list>
				<update_filter_list>
					<update_filter update_filter_name="USER_ID" update_filter_db_data_type="string" update_filter_data_source="ReqBrokerReqObj.request_data.user_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
				</update_filter_list>
			</db_update_info>
		</exec_function>
		<exec_function function_name="login_response_copy_object" function_type="copy_object">
			<copy_object_info object_name="login_resp_obj">
				<object_data key="full_name" data_source_type="reqbrokermap" data_source="lms_user_data_record.full_name" data_type="string"></object_data>
				<object_data key="email" data_source_type="reqbrokermap" data_source="lms_user_data_record.email" data_type="string"></object_data>
				<object_data key="role" data_source_type="reqbrokermap" data_source="lms_user_data_record.role" data_type="string"></object_data>
				<object_data key="passwd_status" data_source_type="reqbrokermap" data_source="lms_user_data_record.passwd_status" data_type="string"></object_data>
				<object_data key="last_login_date" data_source_type="reqbrokermap" data_source="lms_user_data_record.last_login_date" data_type="string"></object_data>
				<object_data key="last_login_time" data_source_type="reqbrokermap" data_source="lms_user_data_record.last_login_time" data_type="string"></object_data>
				<object_data key="profile_picture_url" data_source_type="reqbrokermap" data_source="lms_user_data_record.profile_picture_url" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
		<exec_function function_name="login_token_response_copy_object" function_type="copy_object">
			<copy_object_info object_name="login_resp_obj.token_info">
				<object_data key="access_token" data_source_type="reqbrokermap" data_source="gen_access_token_obj.token_value" data_type="string"></object_data>
				<object_data key="refresh_token" data_source_type="reqbrokermap" data_source="gen_refresh_token_obj.token_value" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
		<exec_function function_name="login_lms_user_data_compose_response" function_type="compose_response">
			<response_info>
				<resp_obj obj_name="login_resp_obj" obj_type="object" data_source="login_resp_obj"></resp_obj>
			</response_info>
		</exec_function>
		</exec_group>
	</sem_lexical_parser>
	`
	return lexicalParserStr
}
