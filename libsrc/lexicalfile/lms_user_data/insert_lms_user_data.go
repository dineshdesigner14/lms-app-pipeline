package lexicalfile_lms_user_data

func GetInsertLMSUserDataServiceXML() interface{} {
	lexicalParserStr := `
	<sem_lexical_parser microservice="insert_lms_user_data">
		<exec_group group_name="insert_lms_user_data" group_condition="ReqBrokerReqObj.header_data.channel=='web_admin_portal'" >
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
		<exec_function function_name="lms_user_data_validate_fld" function_type="validate_fld">
			<validate_info app_err_desc="Validation failed for user data" app_corrective_action="Please validate the data">
				<validate_fld fld_name="ReqBrokerReqObj.request_data.full_name" fld_data_type="string" fld_min_len="1" fld_max_len="255" fld_type="pattern" fld_pattern="^[A-Za-z ]+$" app_err_desc="Invalid FULL_NAME" app_corrective_action="Only letters and spaces allowed."></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.password" fld_data_type="string" fld_min_len="8" fld_max_len="255" fld_type="pattern" fld_pattern=".*" app_err_desc="Pattern validation failed for password" app_corrective_action="Please enter a valid password with a minimum of 6 characters, including at least one uppercase letter, one numeric digit, and one special character."></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.email" fld_data_type="string" fld_min_len="5" fld_max_len="255" fld_type="pattern" fld_pattern="^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,3}$" app_err_desc="Pattern validation failed for primary mail address" app_corrective_action="Primary mail address failed pattern validation. Please correct the format"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.mobile_number" fld_data_type="string" fld_min_len="5" fld_max_len="20" fld_type="pattern" fld_pattern="^\+\d{1,3}-\d{10}$" app_err_desc="Invalid MOBILE_NUMBER format" app_corrective_action="Format must be cc followed by mobile num."></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.role" fld_data_type="string" fld_min_len="5" fld_max_len="20" fld_type="pattern" fld_pattern="^(SUPERADMIN|ADMIN|STUDENT|TRAINER|INSTRUCTOR)$" app_err_desc="Invalid ROLE" app_corrective_action="Allowed values are ADMIN, STUDENT, TRAINER, INSTRUCTOR."></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.status" fld_data_type="boolean" fld_min_len="1" fld_max_len="5" fld_type="pattern" fld_pattern="^(true|false)$" app_err_desc="Invalid status" app_corrective_action="Only Boolean value is allowed."></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.profile_picture_url" fld_data_type="string" fld_min_len="0" fld_max_len="500" fld_type="pattern" fld_pattern=".*" app_err_desc="Invalid URL" app_corrective_action="Provide a valid URL."></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.date_of_birth" fld_data_type="string" fld_min_len="1" fld_max_len="20" fld_type="pattern" fld_pattern=".*" app_err_desc="Invalid DATE_OF_BIRTH" app_corrective_action="Only digits and special characters allowed."></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.gender" fld_data_type="string" fld_min_len="1" fld_max_len="10" fld_type="pattern" fld_pattern=".*" app_err_desc="Invalid GENDER" app_corrective_action="Allowed: male, female, others."></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.bio" fld_data_type="string" fld_min_len="0" fld_max_len="500" fld_type="pattern" fld_pattern=".*" app_err_desc="Invalid BIO" app_corrective_action="Only letters, digits, spaces and special characters allowed."></validate_fld>
			</validate_info>
		</exec_function>
			
			<exec_function function_name="insert_lms_user_data_db_single_read" function_type="db_single_read">
				<db_single_read_info app_err_desc="lms user data data table is not found" app_corrective_action="User data cannot be configured as the table does not exist">
					<table_name>LMS_USER_DATA</table_name>
					<filter_info>
						<filter_data filter_name="MOBILE_NUMBER" filter_id="ReqBrokerReqObj.request_data.mobile_number" filter_data_type="string" filter_condition="AND"></filter_data>
					</filter_info>
					<result_info app_err_desc="The mobile_number:{{.ReqBrokerReqObj.request_data.mobile_number}} already exists." app_corrective_action="Please configure new mobile number and proceed">
						<result_success>NoRows</result_success>
					</result_info>
				</db_single_read_info>
				<err_info>
					<err_desc err_code="Duplicate" app_err_desc="Duplicate records found for the mobile_number:{{.ReqBrokerReqObj.request_data.mobile_number}} provided in the input." app_corrective_action="Make sure each Mobile Number shows up only once." />   
				</err_info>
			</exec_function>

			<exec_function function_name="lms_user_data_password_transform" function_type="transform_object">
			<transform_object_info object_name="password_transform_obj">
				<object_data algo="hash" key="hash_password" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.password" data_type="string" />
			</transform_object_info>
		</exec_function>
			<exec_function function_name="insert_lms_user_data_create_object" function_type="copy_object">
				<copy_object_info object_name="insert_lms_user_data_table_obj">
					<object_data key="full_name" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.full_name" data_type="string"></object_data>
					<object_data key="email" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.email" data_type="string"></object_data>
					<object_data key="mobile_number" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.mobile_number" data_type="string"></object_data>
					<object_data key="password_hash" data_source_type="reqbrokermap" data_source="password_transform_obj.hash_password" data_type="string"></object_data>
					<object_data key="role" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.role" data_type="string"></object_data>
					<object_data key="profile_picture_url" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.profile_picture_url" data_type="string"></object_data>
					<object_data key="date_of_birth" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.date_of_birth" data_type="string"></object_data>
					<object_data key="gender" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.gender" data_type="string"></object_data>
					<object_data key="address" data_source_type="reqbrokermap"  data_source="ReqBrokerReqObj.request_data.address" data_type="string"></object_data>
					<object_data key="bio" data_source_type="reqbrokermap"  data_source="ReqBrokerReqObj.request_data.bio" data_type="string"></object_data>
                    <object_data key="created_by" data_source_type="reqbrokermap" data_source="decode_token_obj.email" data_type="string"></object_data>
					<object_data key="created_date" data_source_type="key" data_source="get_date" data_type="string"></object_data>
					<object_data key="created_time" data_source_type="key" data_source="get_time" data_type="string"></object_data>
					<object_data key="updated_by" data_source_type="key" data_source="get_na" data_type="string"></object_data>
					<object_data key="updated_date" data_source_type="key" data_source="get_date" data_type="string"></object_data>
					<object_data key="updated_time" data_source_type="key" data_source="get_time" data_type="string"></object_data>
					<object_data key="deleted_by" data_source_type="key" data_source="get_na" data_type="string"></object_data>
					<object_data key="deleted_date" data_source_type="key" data_source="get_date" data_type="string"></object_data>
					<object_data key="deleted_time" data_source_type="key" data_source="get_time" data_type="string"></object_data>
					<object_data key="passwd_retry_count"  data_source_type="raw_value" data_source="0" data_type="string"></object_data>
					<object_data key="passwd_modified_date" data_source_type="key" data_source="get_date" data_type="string"></object_data>
					<object_data key="passwd_modified_time"  data_source_type="key" data_source="get_time" data_type="string"></object_data>
					<object_data key="last_login_date" data_source_type="key" data_source="get_date" data_type="string"></object_data>
					<object_data key="last_login_time"  data_source_type="key" data_source="get_time" data_type="string"></object_data>
                    <object_data key="passwd_status"  data_source_type="raw_value" data_source="2" data_type="string"></object_data>
				</copy_object_info>
			</exec_function>

			<exec_function function_name="insert_lms_user_data_db_insert" function_type="db_insert">
				<db_insert_info>
					<table_name>LMS_USER_DATA</table_name>
					<insert_data_list>
					    <insert_data insert_data_name="FULL_NAME" insert_db_data_type="string" insert_data_source="insert_lms_user_data_table_obj.full_name" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="EMAIL" insert_db_data_type="string" insert_data_source="insert_lms_user_data_table_obj.email" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="MOBILE_NUMBER" insert_db_data_type="string" insert_data_source="insert_lms_user_data_table_obj.mobile_number" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="PASSWORD_HASH" insert_db_data_type="string" insert_data_source="insert_lms_user_data_table_obj.password_hash" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="ROLE" insert_db_data_type="string" insert_data_source="insert_lms_user_data_table_obj.role" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="STATUS" insert_db_data_type="boolean" insert_data_source="ReqBrokerReqObj.request_data.status" insert_data_type="boolean"></insert_data>
						<insert_data insert_data_name="PROFILE_PICTURE_URL" insert_db_data_type="string" insert_data_source="insert_lms_user_data_table_obj.profile_picture_url" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="DATE_OF_BIRTH" insert_db_data_type="date" insert_data_source="insert_lms_user_data_table_obj.date_of_birth" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="GENDER" insert_db_data_type="string" insert_data_source="insert_lms_user_data_table_obj.gender" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="ADDRESS" insert_db_data_type="string" insert_data_source="insert_lms_user_data_table_obj.address" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="BIO" insert_db_data_type="string" insert_data_source="insert_lms_user_data_table_obj.bio" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="CREATED_BY" insert_db_data_type="string" insert_data_source="insert_lms_user_data_table_obj.created_by" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="CREATED_DATE" insert_db_data_type="date" insert_data_source="insert_lms_user_data_table_obj.created_date" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="CREATED_TIME" insert_db_data_type="string" insert_data_source="insert_lms_user_data_table_obj.created_time" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="UPDATED_BY" insert_db_data_type="string" insert_data_source="insert_lms_user_data_table_obj.updated_by" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="UPDATED_DATE" insert_db_data_type="date" insert_data_source="insert_lms_user_data_table_obj.updated_date" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="UPDATED_TIME" insert_db_data_type="string" insert_data_source="insert_lms_user_data_table_obj.updated_time" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="DELETED_BY" insert_db_data_type="string" insert_data_source="insert_lms_user_data_table_obj.deleted_by" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="DELETED_DATE" insert_db_data_type="date" insert_data_source="insert_lms_user_data_table_obj.deleted_date" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="DELETED_TIME" insert_db_data_type="string" insert_data_source="insert_lms_user_data_table_obj.deleted_time" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="PASSWD_RETRY_COUNT" insert_db_data_type="string" insert_data_source="insert_lms_user_data_table_obj.passwd_retry_count" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="PASSWD_MODIFIED_DATE" insert_db_data_type="date" insert_data_source="insert_lms_user_data_table_obj.passwd_modified_date" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="PASSWD_MODIFIED_TIME" insert_db_data_type="string" insert_data_source="insert_lms_user_data_table_obj.passwd_modified_time" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="LAST_LOGIN_DATE" insert_db_data_type="date" insert_data_source="insert_lms_user_data_table_obj.last_login_date" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="LAST_LOGIN_TIME" insert_db_data_type="string" insert_data_source="insert_lms_user_data_table_obj.last_login_time" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="PASSWD_STATUS" insert_db_data_type="string" insert_data_source="insert_lms_user_data_table_obj.passwd_status" insert_data_type="string"></insert_data>
						</insert_data_list>
				</db_insert_info>
			</exec_function>
            <exec_function function_name="insert_lms_user_token_info_copy_object" function_type="copy_object">
				<copy_object_info object_name="insert_lms_user_token_info_table_obj">
					<object_data key="access_token" data_source_type="raw_value" data_source="NA" data_type="string"></object_data>
					<object_data key="refresh_token" data_source_type="raw_value" data_source="NA" data_type="string"></object_data>
				</copy_object_info>
			</exec_function>
			<exec_function function_name="insert_lms_token_info_db_insert" function_type="db_insert">
				<db_insert_info>
					<table_name>LMS_TOKEN_INFO</table_name>
					<insert_data_list>
					    <insert_data insert_data_name="USER_ID" insert_db_data_type="string" insert_data_source="ReqBrokerReqObj.request_data.email" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="ACCESS_TOKEN" insert_db_data_type="string" insert_data_source="insert_lms_user_token_info_table_obj.access_token" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="REFRESH_TOKEN" insert_db_data_type="string" insert_data_source="insert_lms_user_token_info_table_obj.refresh_token" insert_data_type="string"></insert_data>
						</insert_data_list>
				</db_insert_info>
			</exec_function>
			      
		<exec_function function_name="lms_enquiry_personal_info_auth_db_single_read" function_type="db_single_read">
			    <db_single_read_info app_err_desc="Data Already Exists or No Rows in Result Set" app_corrective_action="check the inputs and provide proper values">
			     	<table_name>LMS_USER_DATA</table_name>
				<filter_info>
					<filter_data filter_name="MOBILE_NUMBER" filter_id="ReqBrokerReqObj.request_data.mobile_number" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_user_data_table_record</store_obj_name>
			</db_single_read_info>
				<err_info>
					<err_desc err_code="NoRows" app_err_desc="The mobile_number:{{.ReqBrokerReqObj.request_data.mobile_number}} not found." app_corrective_action="Please create the user and proceed"/>   
					<err_desc err_code="Duplicate" app_err_desc="The mobile_number:{{.ReqBrokerReqObj.request_data.mobile_number}} already exists." app_corrective_action="Please configure new mobile number and proceed" />   
				</err_info>
			</exec_function>
           <exec_function function_name="edit_lms_enquiry_personal_info_auth_db_multi_read" function_type="db_multi_read">
			<db_multi_read_info app_err_desc="Data Already Exists or No Rows in Result Set" app_corrective_action="check the inputs and provide proper values">
				<table_name>LMS_ENQUIRY</table_name>
				<filter_info>
					<filter_data filter_name="PHONE_NUMBER" filter_id="ReqBrokerReqObj.request_data.mobile_number" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_enquiry_table_list</store_obj_name>
			</db_multi_read_info>
		</exec_function> 
            <exec_function function_name="lms_enquiry_loop" function_type="start_loop" index_name="i" start_index="0" end_index="len(lms_enquiry_table_list)" />
		    <exec_function function_name="insert_lms_user_data_create_object" function_type="copy_object" group_condition="len(lms_enquiry_table_list)!=0">
				<copy_object_info object_name="user_data_response_obj">
				   <object_data key="enquiry_id" data_source_type="reqbrokermap" data_source="lms_enquiry_table_list[i].id" data_type="string"></object_data>
                   <object_data key="id" data_source_type="reqbrokermap" data_source="lms_user_data_table_record.id" data_type="string"></object_data>
				  <object_data key="resp_msg" data_source_type="raw_value" data_source="User created successfully!!!" data_type="string"></object_data>
			</copy_object_info>
			</exec_function>
             		<exec_function function_name="lms_enquiry_loop" function_type="end_loop" index_name="i" />

			<exec_function function_name="insert_lms_user_data_create_object" function_type="copy_object" group_condition="len(lms_enquiry_table_list)==0">
				<copy_object_info object_name="user_data_response_obj">
				 <object_data key="enquiry_id" data_source_type="raw_value" data_source="NA" data_type="string"></object_data>
				<object_data key="id" data_source_type="reqbrokermap" data_source="lms_user_data_table_record.id" data_type="string"></object_data>
				<object_data key="resp_msg" data_source_type="raw_value" data_source="User created successfully!!!" data_type="string"></object_data>
			</copy_object_info>
			</exec_function>
			<exec_function function_name="insert_lms_user_data_compose_response" function_type="compose_response">
				<response_info>
					<resp_obj obj_name="user_data_response_obj" obj_type="object" data_source="user_data_response_obj"></resp_obj>
				</response_info>
			</exec_function>
		</exec_group>
	</sem_lexical_parser>
	`
	return lexicalParserStr
}
