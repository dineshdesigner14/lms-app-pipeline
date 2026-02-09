package lexicalfile_lms_user_auth_data

func GetLMSUserResetPasswdServiceXML() interface{} {
	lexicalParserStr := `
	<sem_lexical_parser microservice="lms_user_reset_passwd">
		<exec_group group_name="lms_user_reset_passwd" group_condition="ReqBrokerReqObj.request_data.old_password!=nil &amp;&amp;ReqBrokerReqObj.request_data.verify_code==nil" >
		<exec_function function_name="lms_user_data_validate_fld" function_type="validate_fld">
			<validate_info app_err_desc="Validation failed for user data" app_corrective_action="Please validate the data">
				<validate_fld fld_name="ReqBrokerReqObj.request_data.user_id" fld_data_type="string" fld_min_len="1" fld_max_len="255" fld_type="pattern" fld_pattern="^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,3}$" app_err_desc="Invalid UserID" app_corrective_action="Only send valid email format."></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.old_password" fld_data_type="string" fld_min_len="5" fld_max_len="255" fld_type="pattern" fld_pattern=".*" app_corrective_action="Password will have atleast one captial cases, small cases, number and special character"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.new_password" fld_data_type="string" fld_min_len="5" fld_max_len="255" fld_type="pattern" fld_pattern=".*" app_corrective_action="Password must have atleast one captial cases, small cases, number and special character"></validate_fld>
			</validate_info>
		</exec_function>
		<exec_function function_name="reset_passwd_lms_user_data_db_single_read" function_type="db_single_read">
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
		<exec_function function_name="lms_user_data_password_transform" function_type="transform_object">
			<transform_object_info object_name="password_transform_obj">
				<object_data algo="hash" key="hash_password" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.old_password" data_type="string" />
				<object_data algo="hash" key="hash_new_password" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.new_password" data_type="string" />
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
		<exec_function function_name="mapp_token_info_copy_object" function_type="copy_object">
			<copy_object_info object_name="passwd_obj">
				<object_data key="passwd_modified_date" data_source_type="key" data_source="get_date" data_type="string"></object_data>
				<object_data key="passwd_modified_time" data_source_type="key" data_source="get_time" data_type="string"></object_data>
				<object_data key="passwd_status" data_source_type="raw_value" data_source="1" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
		<exec_function function_name="edit_lms_user_data_db_update" function_type="db_update">
			<db_update_info>
				<table_name>LMS_USER_DATA</table_name>
				<update_data_list>
					<update_data update_data_name="PASSWORD_HASH" update_db_data_type="string" update_data_source="password_transform_obj.hash_new_password" update_data_type="string"></update_data>
					<update_data update_data_name="PASSWD_MODIFIED_DATE" update_db_data_type="date" update_data_source="passwd_obj.passwd_modified_date" update_data_type="string"></update_data>
					<update_data update_data_name="PASSWD_MODIFIED_TIME" update_db_data_type="string" update_data_source="passwd_obj.passwd_modified_time" update_data_type="string"></update_data>
					<update_data update_data_name="PASSWD_STATUS" update_db_data_type="string" update_data_source="passwd_obj.passwd_status" update_data_type="string"></update_data>
				</update_data_list>
				<update_filter_list>
					<update_filter update_filter_name="EMAIL" update_filter_db_data_type="string" update_filter_data_source="ReqBrokerReqObj.request_data.user_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
				</update_filter_list>
			</db_update_info>
		</exec_function>
		<exec_function function_name="reset_password_response_copy_object" function_type="copy_object">
			<copy_object_info object_name="reset_password_obj">
				<object_data key="status" data_source_type="raw_value" data_source="Password updated successfully!" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
		<exec_function function_name="reset_passwd_lms_user_data_compose_response" function_type="compose_response">
			<response_info>
				<resp_obj obj_name="reset_password_obj" obj_type="object" data_source="reset_password_obj"></resp_obj>
			</response_info>
		</exec_function>
		</exec_group>
		<exec_group group_name="lms_user_reset_passwd" group_condition="ReqBrokerReqObj.request_data.old_password==nil &amp;&amp; ReqBrokerReqObj.request_data.verify_code!=nil" >
			<exec_function function_name="lms_user_data_validate_fld" function_type="validate_fld">
				<validate_info app_err_desc="Validation failed for user data" app_corrective_action="Please validate the data">
					<validate_fld fld_name="ReqBrokerReqObj.request_data.user_id" fld_data_type="string" fld_min_len="1" fld_max_len="255" fld_type="pattern" fld_pattern="^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,3}$" app_err_desc="Invalid UserID pattern" app_corrective_action="Only send valid email format."></validate_fld>
					<validate_fld fld_name="ReqBrokerReqObj.request_data.verify_code" fld_data_type="string" fld_min_len="5" fld_max_len="255" fld_type="pattern" fld_pattern="^[a-zA-Z0-9.,|_* -]+$" app_err_desc="Verification code pattern error" app_corrective_action="Please send verification code in valid pattern"></validate_fld>
					<validate_fld fld_name="ReqBrokerReqObj.request_data.new_password" fld_data_type="string" fld_min_len="5" fld_max_len="255" fld_type="pattern" fld_pattern=".*" app_corrective_action="Password must have atleast one captial cases, small cases, number and special character"></validate_fld>
				</validate_info>
			</exec_function>
			<exec_function function_name="reset_passwd_lms_user_data_db_single_read" function_type="db_single_read">
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
			<exec_function function_name="lms_user_data_verification_code_validate_condition" function_type="validate_condition">
				<validate_condition_info app_err_desc="Invalid verification code" app_corrective_action="Kindly enter a valid verification code.">
					<validate_condition_expression>ReqBrokerReqObj.request_data.verify_code=='123456'</validate_condition_expression>
				</validate_condition_info>
			</exec_function>
			<exec_function function_name="mapp_token_info_copy_object" function_type="copy_object">
				<copy_object_info object_name="passwd_obj">
					<object_data key="passwd_modified_date" data_source_type="key" data_source="get_date" data_type="string"></object_data>
					<object_data key="passwd_modified_time" data_source_type="key" data_source="get_time" data_type="string"></object_data>
				</copy_object_info>
			</exec_function>
			<exec_function function_name="lms_user_data_password_transform" function_type="transform_object">
				<transform_object_info object_name="password_transform_obj">
					<object_data algo="hash" key="hash_new_password" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.new_password" data_type="string" />
				</transform_object_info>
			</exec_function>
			<exec_function function_name="edit_lms_user_data_db_update" function_type="db_update">
				<db_update_info>
					<table_name>LMS_USER_DATA</table_name>
					<update_data_list>
						<update_data update_data_name="PASSWORD_HASH" update_db_data_type="string" update_data_source="password_transform_obj.hash_new_password" update_data_type="string"></update_data>
						<update_data update_data_name="PASSWD_MODIFIED_DATE" update_db_data_type="date" update_data_source="passwd_obj.passwd_modified_date" update_data_type="string"></update_data>
						<update_data update_data_name="PASSWD_MODIFIED_TIME" update_db_data_type="string" update_data_source="passwd_obj.passwd_modified_time" update_data_type="string"></update_data>
					</update_data_list>
					<update_filter_list>
						<update_filter update_filter_name="EMAIL" update_filter_db_data_type="string" update_filter_data_source="ReqBrokerReqObj.request_data.user_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
					</update_filter_list>
				</db_update_info>
			</exec_function>
			<exec_function function_name="reset_password_response_copy_object" function_type="copy_object">
				<copy_object_info object_name="reset_password_obj">
					<object_data key="status" data_source_type="raw_value" data_source="Password updated successfully!" data_type="string"></object_data>
				</copy_object_info>
			</exec_function>
			<exec_function function_name="reset_passwd_lms_user_data_compose_response" function_type="compose_response">
				<response_info>
					<resp_obj obj_name="reset_password_obj" obj_type="object" data_source="reset_password_obj"></resp_obj>
				</response_info>
			</exec_function>
		</exec_group>
		<exec_group group_name="lms_user_reset_passwd_admin" group_condition="ReqBrokerReqObj.header_data.channel=='web_admin_portal'">
			<exec_function function_name="lms_user_data_validate_fld" function_type="validate_fld">
				<validate_info app_err_desc="Validation failed for user data" app_corrective_action="Please validate the data">
				   <validate_fld fld_name="ReqBrokerReqObj.request_data.user_id" fld_data_type="string" fld_min_len="6" fld_max_len="255" fld_type="pattern" fld_pattern="^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,3}$" app_err_desc="Pattern validation failed for user id" app_corrective_action="User ID failed pattern validation. Please correct the format"></validate_fld>
                   <validate_fld fld_name="ReqBrokerReqObj.request_data.new_password" fld_data_type="string" fld_min_len="8" fld_max_len="255" fld_type="pattern" fld_pattern=".*" app_corrective_action="Password must have atleast one captial cases, small cases, number and special character"></validate_fld>
				</validate_info>
			</exec_function>
			<exec_function function_name="decode_token_function_type" function_type="decode_token">
			<decode_token_info token_value="ReqBrokerReqObj.header_data.auth_token" token_object="decode_token_obj" app_err_desc="Decode token failed!. Invalid access token." app_corrective_action="Please provide valid access token."/>
		</exec_function>
			<exec_function function_name="reset_passwd_lms_user_data_db_single_read" function_type="db_single_read">
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
			<exec_function function_name="lms_user_data_role_validate_condition" function_type="validate_condition">
				<validate_condition_info app_err_desc="Invalid Role" app_corrective_action="User role Should be ADMIN.">
					<validate_condition_expression>lms_user_data_record.role=='ADMIN'|| lms_user_data_record.role=='SUPERADMIN'</validate_condition_expression>
				</validate_condition_info>
			</exec_function>
			<exec_function function_name="lms_user_data_db_single_read" function_type="db_single_read">
			    <db_single_read_info app_err_desc="Data Already Exists or No Rows in Result Set" app_corrective_action="check the inputs and provide proper values">
			     	<table_name>LMS_USER_DATA</table_name>
				<filter_info>
					<filter_data filter_name="EMAIL" filter_id="ReqBrokerReqObj.request_data.user_id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_user_data_table_record</store_obj_name>
			</db_single_read_info>
				<err_info>
					<err_desc err_code="NoRows" app_err_desc="The user_id:{{.ReqBrokerReqObj.request_data.user_id}} not found." app_corrective_action="Please create the user and proceed"/>   
					<err_desc err_code="Duplicate" app_err_desc="The user_id:{{.ReqBrokerReqObj.request_data.user_id}} already exists." app_corrective_action="Please configure new user number and proceed" />   
				</err_info>
			</exec_function>
		
			<exec_function function_name="mapp_token_info_copy_object" function_type="copy_object">
				<copy_object_info object_name="passwd_obj">
					<object_data key="passwd_modified_date" data_source_type="key" data_source="get_date" data_type="string"></object_data>
					<object_data key="passwd_modified_time" data_source_type="key" data_source="get_time" data_type="string"></object_data>
				</copy_object_info>
			</exec_function>
			<exec_function function_name="lms_user_data_password_transform" function_type="transform_object">
				<transform_object_info object_name="password_transform_obj">
					<object_data algo="hash" key="hash_new_password" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.new_password" data_type="string" />
				</transform_object_info>
			</exec_function>
			<exec_function function_name="edit_lms_user_data_db_update" function_type="db_update">
				<db_update_info>
					<table_name>LMS_USER_DATA</table_name>
					<update_data_list>
						<update_data update_data_name="PASSWORD_HASH" update_db_data_type="string" update_data_source="password_transform_obj.hash_new_password" update_data_type="string"></update_data>
						<update_data update_data_name="PASSWD_MODIFIED_DATE" update_db_data_type="date" update_data_source="passwd_obj.passwd_modified_date" update_data_type="string"></update_data>
						<update_data update_data_name="PASSWD_MODIFIED_TIME" update_db_data_type="string" update_data_source="passwd_obj.passwd_modified_time" update_data_type="string"></update_data>
					</update_data_list>
					<update_filter_list>
							<update_filter update_filter_name="ID" update_filter_db_data_type="string" update_filter_data_source="lms_user_data_table_record.id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
                            <update_filter update_filter_name="EMAIL" update_filter_db_data_type="string" update_filter_data_source="lms_user_data_table_record.email" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
					</update_filter_list>
				</db_update_info>
			</exec_function>
			<exec_function function_name="reset_password_response_copy_object" function_type="copy_object">
				<copy_object_info object_name="reset_password_obj">
					<object_data key="status" data_source_type="raw_value" data_source="Password updated successfully!" data_type="string"></object_data>
				</copy_object_info>
			</exec_function>
			<exec_function function_name="reset_passwd_lms_user_data_compose_response" function_type="compose_response">
				<response_info>
					<resp_obj obj_name="reset_password_obj" obj_type="object" data_source="reset_password_obj"></resp_obj>
				</response_info>
			</exec_function>
		</exec_group>
	</sem_lexical_parser>
	`
	return lexicalParserStr
}
