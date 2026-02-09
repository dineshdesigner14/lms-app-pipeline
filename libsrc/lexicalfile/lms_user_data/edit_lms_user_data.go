package lexicalfile_lms_user_data

func GetEditLMSUserDataServiceXML() interface{} {
	lexicalParserStr := `
	<sem_lexical_parser microservice="edit_lms_user_data" group_condition="ReqBrokerReqObj.header_data.channel=='web_admin_portal'">
		<exec_group group_name="edit_lms_user_data">
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
			<exec_function function_name="edit_lms_user_data_validate_fld" function_type="validate_fld">
			 <validate_info app_err_desc="Validation failed for user data" app_corrective_action="Please validate the data">
				<validate_fld fld_name="ReqBrokerReqObj.request_data.id" fld_data_type="string" fld_min_len="1" fld_max_len="255" fld_type="pattern" fld_pattern="^[A-Za-z0-9 ]+$" app_err_desc="Invalid ID" app_corrective_action="Only Numbers Allowed"></validate_fld>
                <validate_fld fld_name="ReqBrokerReqObj.request_data.full_name" fld_data_type="string" fld_min_len="1" fld_max_len="255" fld_type="pattern" fld_pattern="^[A-Za-z ]+$" app_err_desc="Invalid FULL_NAME" app_corrective_action="Only letters and spaces allowed."></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.email" fld_data_type="string" fld_min_len="5" fld_max_len="255" fld_type="pattern" fld_pattern="^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,3}$" app_err_desc="Pattern validation failed for primary mail address" app_corrective_action="Primary mail address failed pattern validation. Please correct the format"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.mobile_number" fld_data_type="string" fld_min_len="5" fld_max_len="20" fld_type="pattern" fld_pattern="^\+\d{1,3}-\d{10}$" app_err_desc="Invalid MOBILE_NUMBER format" app_corrective_action="Format must be cc followed by mobile num."></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.role" fld_data_type="string" fld_min_len="5" fld_max_len="20" fld_type="pattern" fld_pattern="(SUPER ADMIN|ADMIN|STUDENT|TRAINER|INSTRUCTOR)$" app_err_desc="Invalid ROLE" app_corrective_action="Allowed values are ADMIN, STUDENT, TRAINER, INSTRUCTOR."></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.status" fld_data_type="boolean" fld_min_len="1" fld_max_len="5" fld_type="pattern" fld_pattern="^(true|false|0|1)$" app_err_desc="Invalid status" app_corrective_action="Only Boolean value is allowed."></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.profile_picture_url" fld_data_type="string" fld_min_len="0" fld_max_len="500" fld_type="pattern" fld_pattern=".*" app_err_desc="Invalid URL" app_corrective_action="Provide a valid URL."></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.date_of_birth" fld_data_type="string" fld_min_len="1" fld_max_len="20" fld_type="pattern" fld_pattern=".*" app_err_desc="Invalid DATE_OF_BIRTH" app_corrective_action="Only digits and special characters allowed."></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.gender" fld_data_type="string" fld_min_len="4" fld_max_len="10" fld_type="pattern" fld_pattern=".*" app_err_desc="Invalid GENDER" app_corrective_action="Allowed: male, female, others."></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.bio" fld_data_type="string" fld_min_len="0" fld_max_len="500" fld_type="pattern" fld_pattern=".*" app_err_desc="Invalid BIO" app_corrective_action="Only letters, digits, spaces and special characters allowed."></validate_fld>
			</validate_info>
			</exec_function>
			<exec_function function_name="edit_lms_enquiry_personal_info_auth_db_single_read" function_type="db_single_read">
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
		
			<exec_function function_name="edit_lms_user_data_copy_object" function_type="copy_object">
				<copy_object_info object_name="edit_lms_user_data_additional_obj">
				<object_data key="full_name" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.full_name" data_type="string"></object_data>
					<object_data key="email" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.email" data_type="string"></object_data>
					<object_data key="mobile_number" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.mobile_number" data_type="string"></object_data>
					<object_data key="role" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.role" data_type="string"></object_data>
					<object_data key="status" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.status" data_type="Boolean"></object_data>
					<object_data key="profile_picture_url" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.profile_picture_url" data_type="string"></object_data>
					<object_data key="date_of_birth" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.date_of_birth" data_type="string"></object_data>
					<object_data key="gender" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.gender" data_type="string"></object_data>
					<object_data key="address" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.address" data_type="string"></object_data>
					<object_data key="bio" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.bio" data_type="string"></object_data>
					<object_data key="updated_by" data_source_type="reqbrokermap" data_source="decode_token_obj.email" data_type="string"></object_data>
					<object_data key="updated_date" data_source_type="key" data_source="get_date" data_type="string"></object_data>
					<object_data key="updated_time" data_source_type="key" data_source="get_time" data_type="string"></object_data>
					<object_data key="passwd_retry_count"  data_source_type="reqbrokermap"  data_source="lms_user_data_table_record.passwd_retry_count" data_type="string"></object_data>
					<object_data key="passwd_modified_date" data_source_type="reqbrokermap" data_source="lms_user_data_table_record.passwd_modified_date" data_type="string"></object_data>
					<object_data key="passwd_modified_time" data_source_type="reqbrokermap" data_source="lms_user_data_table_record.passwd_modified_time" data_type="string"></object_data>
					<object_data key="last_login_date" data_source_type="reqbrokermap" data_source="lms_user_data_table_record.last_login_date" data_type="string"></object_data>
					<object_data key="last_login_time" data_source_type="reqbrokermap" data_source="lms_user_data_table_record.last_login_time" data_type="string"></object_data>
                    <object_data key="passwd_status" data_source_type="reqbrokermap" data_source="lms_user_data_table_record.passwd_status" data_type="string"></object_data>
				</copy_object_info>
			</exec_function>
			<exec_function function_name="edit_lms_user_data_db_update" function_type="db_update">
				<db_update_info>
					<table_name>LMS_USER_DATA</table_name>
					<update_data_list>
						<update_data update_data_name="FULL_NAME" update_db_data_type="string" update_data_source="edit_lms_user_data_additional_obj.full_name" update_data_type="string"></update_data>
                       <update_data update_data_name="EMAIL" update_db_data_type="string" update_data_source="edit_lms_user_data_additional_obj.email" update_data_type="string"></update_data>
                       <update_data update_data_name="ROLE" update_db_data_type="string" update_data_source="edit_lms_user_data_additional_obj.role" update_data_type="string"></update_data>
                       <update_data update_data_name="STATUS" update_db_data_type="boolean" update_data_source="ReqBrokerReqObj.request_data.status" update_data_type="boolean"></update_data>
                       <update_data update_data_name="PROFILE_PICTURE_URL" update_db_data_type="string" update_data_source="edit_lms_user_data_additional_obj.profile_picture_url" update_data_type="string"></update_data>
                       <update_data update_data_name="DATE_OF_BIRTH" update_db_data_type="date" update_data_source="edit_lms_user_data_additional_obj.date_of_birth" update_data_type="string"></update_data>
                       <update_data update_data_name="GENDER" update_db_data_type="string" update_data_source="edit_lms_user_data_additional_obj.gender" update_data_type="string"></update_data>
                       <update_data update_data_name="ADDRESS" update_db_data_type="string" update_data_source="edit_lms_user_data_additional_obj.address" update_data_type="string"></update_data>
                       <update_data update_data_name="BIO" update_db_data_type="string" update_data_source="edit_lms_user_data_additional_obj.bio" update_data_type="string"></update_data>
						<update_data update_data_name="UPDATED_BY" update_db_data_type="string" update_data_source="edit_lms_user_data_additional_obj.updated_by" update_data_type="string"></update_data>
						<update_data update_data_name="UPDATED_DATE" update_db_data_type="date" update_data_source="edit_lms_user_data_additional_obj.updated_date" update_data_type="string"></update_data>
						<update_data update_data_name="UPDATED_TIME" update_db_data_type="string" update_data_source="edit_lms_user_data_additional_obj.updated_time" update_data_type="string"></update_data>
					</update_data_list>
					<update_filter_list>
					    <update_filter update_filter_name="ID" update_filter_db_data_type="string" update_filter_data_source="ReqBrokerReqObj.request_data.id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
						<update_filter update_filter_name="MOBILE_NUMBER" update_filter_db_data_type="string" update_filter_data_source="ReqBrokerReqObj.request_data.mobile_number" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
					</update_filter_list>
				</db_update_info>
			</exec_function>
		</exec_group>
		<exec_group group_name="edit_lms_user_course_info" group_condition="ReqBrokerReqObj.header_data.channel=='web_admin_portal'">
		<exec_function function_name="lms_enquiry_course_loop" function_type="start_loop" index_name="i" start_index="0" end_index="len(ReqBrokerReqObj.request_data.courses)" />
		<exec_function function_name="lms_enquiry_course_validate_fld" function_type="validate_fld">
			<validate_info app_err_desc="pattern validation failed for input" app_corrective_action="check the inputs and provide proper values">
				<validate_fld fld_name="ReqBrokerReqObj.request_data.courses[i].course_id" fld_data_type="string" fld_min_len="1" fld_max_len="128" fld_type="pattern" fld_pattern="^[a-zA-Z0-9.,|_* ]+$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.courses[i].db_operation" fld_data_type="string" fld_min_len="1" fld_max_len="128" fld_type="pattern" fld_pattern="^[INSERT|EDIT|DELETE]+$"></validate_fld>
			</validate_info>
		</exec_function>
		<exec_function function_name="insert_lms_enquiry_course_db_single_read" function_type="db_single_read" function_condition="ReqBrokerReqObj.request_data.courses[i].db_operation=='INSERT'">
			<db_single_read_info app_err_desc="Data Already Exists or No Rows in Result Set" app_corrective_action="check the inputs and provide proper values">
				<table_name>LMS_COURSE_PROGRESS</table_name>
				<filter_info>
					<filter_data filter_name="USER_ID" filter_id="ReqBrokerReqObj.request_data.id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="COURSE_ID" filter_id="ReqBrokerReqObj.request_data.courses[i].course_id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<result_info>
					<result_success>NoRows</result_success>
				</result_info>
			</db_single_read_info>
		</exec_function>
		<exec_function function_name="insert_lms_course_progress_copy_object" function_type="copy_object" function_condition="ReqBrokerReqObj.request_data.courses[i].db_operation=='INSERT'">
			<copy_object_info object_name="insert_lms_course_progress_additional_obj">
				<object_data key="user_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.id" data_type="string"></object_data>
				<object_data key="course_id"  data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.courses[i].course_id" data_type="string"></object_data>
				<object_data key="percentage" data_source_type="raw_value" data_source="0" data_type="string"></object_data>
				<object_data key="created_by" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.email" data_type="string"></object_data>
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
		<exec_function function_name="insert_lms_course_progress_db_insert" function_type="db_insert" function_condition="ReqBrokerReqObj.request_data.courses[i].db_operation=='INSERT'">
			<db_insert_info>
				<table_name>LMS_COURSE_PROGRESS</table_name>
				<insert_data_list>
					<insert_data insert_data_name="USER_ID" insert_db_data_type="string" insert_data_source="insert_lms_course_progress_additional_obj.user_id" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="COURSE_ID" insert_db_data_type="string" insert_data_source="insert_lms_course_progress_additional_obj.course_id" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="CREATED_BY" insert_db_data_type="string" insert_data_source="insert_lms_course_progress_additional_obj.created_by" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="CREATED_DATE" insert_db_data_type="date" insert_data_source="insert_lms_course_progress_additional_obj.created_date" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="CREATED_TIME" insert_db_data_type="string" insert_data_source="insert_lms_course_progress_additional_obj.created_time" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="UPDATED_BY" insert_db_data_type="string" insert_data_source="insert_lms_course_progress_additional_obj.updated_by" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="UPDATED_DATE" insert_db_data_type="date" insert_data_source="insert_lms_course_progress_additional_obj.updated_date" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="UPDATED_TIME" insert_db_data_type="string" insert_data_source="insert_lms_course_progress_additional_obj.updated_time" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="DELETED_BY" insert_db_data_type="string" insert_data_source="insert_lms_course_progress_additional_obj.deleted_by" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="DELETED_DATE" insert_db_data_type="date" insert_data_source="insert_lms_course_progress_additional_obj.deleted_date" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="DELETED_TIME" insert_db_data_type="string" insert_data_source="insert_lms_course_progress_additional_obj.deleted_time" insert_data_type="string"></insert_data>
				</insert_data_list>
			</db_insert_info>
		</exec_function>
		<exec_function function_name="insert_lms_course_section_db_single_read" function_type="db_multi_read" function_condition="ReqBrokerReqObj.request_data.courses[i].db_operation=='INSERT'">
			<db_multi_read_info app_err_desc="lms course  section data table is not found" app_corrective_action="course section data cannot be configured as the table does not exist">
				<table_name>LMS_COURSE_SECTION</table_name>
				<filter_info>
					<filter_data filter_name="COURSE_ID" filter_id="ReqBrokerReqObj.request_data.courses[i].course_id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_course_section_table_list</store_obj_name>
			</db_multi_read_info>
			<err_info>
				<err_desc err_code="Duplicate" app_err_desc="Duplicate records found for course_id provided in the input." app_corrective_action="Make sure each course should be unique." />
			</err_info>
		</exec_function>
		<exec_function function_name="lms_enquiry_course_section_loop" function_type="start_loop" index_name="j" start_index="0" end_index="len(lms_course_section_table_list)" />
		<exec_function function_name="insert_lms_course_progress_copy_object" function_type="copy_object" function_condition="lms_course_section_table_list[j].order!=1">
			<copy_object_info object_name="insert_lms_course_progress_additional_obj">
				<object_data key="user_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.id" data_type="string"></object_data>
				<object_data key="course_id"  data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.courses[i].course_id" data_type="string"></object_data>
				<object_data key="section_id"  data_source_type="reqbrokermap" data_source="lms_course_section_table_list[j].id" data_type="string"></object_data>
				<object_data key="is_enabled"  data_source_type="raw_value" data_source="false" data_type="Boolean"></object_data>
				<object_data key="percentage" data_source_type="raw_value" data_source="0" data_type="string"></object_data>
				<object_data key="created_by" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.email" data_type="string"></object_data>
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
		<exec_function function_name="insert_lms_course_progress_copy_object" function_type="copy_object" function_condition="lms_course_section_table_list[j].order==1">
			<copy_object_info object_name="insert_lms_course_progress_additional_obj">
				<object_data key="user_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.id" data_type="string"></object_data>
				<object_data key="course_id"  data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.courses[i].course_id" data_type="string"></object_data>
				<object_data key="section_id"  data_source_type="reqbrokermap" data_source="lms_course_section_table_list[j].id" data_type="string"></object_data>
				<object_data key="is_enabled"  data_source_type="raw_value" data_source="true" data_type="Boolean"></object_data>
				<object_data key="percentage" data_source_type="raw_value" data_source="0" data_type="string"></object_data>
				<object_data key="created_by" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.email" data_type="string"></object_data>
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
		<exec_function function_name="insert_lms_course_section_progress_db_insert" function_type="db_insert">
			<db_insert_info>
				<table_name>LMS_COURSE_SECTION_PROGRESS</table_name>
				<insert_data_list>
					<insert_data insert_data_name="USER_ID" insert_db_data_type="string" insert_data_source="insert_lms_course_progress_additional_obj.user_id" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="COURSE_ID" insert_db_data_type="string" insert_data_source="insert_lms_course_progress_additional_obj.course_id" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="SECTION_ID" insert_db_data_type="string" insert_data_source="insert_lms_course_progress_additional_obj.section_id" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="IS_ENABLED" insert_db_data_type="boolean" insert_data_source="insert_lms_course_progress_additional_obj.is_enabled" insert_data_type="boolean"></insert_data>
					<insert_data insert_data_name="CREATED_BY" insert_db_data_type="string" insert_data_source="insert_lms_course_progress_additional_obj.created_by" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="CREATED_DATE" insert_db_data_type="date" insert_data_source="insert_lms_course_progress_additional_obj.created_date" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="CREATED_TIME" insert_db_data_type="string" insert_data_source="insert_lms_course_progress_additional_obj.created_time" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="UPDATED_BY" insert_db_data_type="string" insert_data_source="insert_lms_course_progress_additional_obj.updated_by" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="UPDATED_DATE" insert_db_data_type="date" insert_data_source="insert_lms_course_progress_additional_obj.updated_date" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="UPDATED_TIME" insert_db_data_type="string" insert_data_source="insert_lms_course_progress_additional_obj.updated_time" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="DELETED_BY" insert_db_data_type="string" insert_data_source="insert_lms_course_progress_additional_obj.deleted_by" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="DELETED_DATE" insert_db_data_type="date" insert_data_source="insert_lms_course_progress_additional_obj.deleted_date" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="DELETED_TIME" insert_db_data_type="string" insert_data_source="insert_lms_course_progress_additional_obj.deleted_time" insert_data_type="string"></insert_data>
				</insert_data_list>
			</db_insert_info>
		</exec_function>
		<exec_function function_name="lms_course_item_db_single_read" function_type="db_multi_read">
			<db_multi_read_info app_err_desc="lms course item data table is not found" app_corrective_action="course item data cannot be configured as the table does not exist">
				<table_name>LMS_COURSE_ITEM</table_name>
				<filter_info>
					<filter_data filter_name="COURSE_ID" filter_id="ReqBrokerReqObj.request_data.courses[i].course_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="SECTION_ID" filter_id="lms_course_section_table_list[j].id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_course_item_table_list</store_obj_name>
			</db_multi_read_info>
			<err_info>
				<err_desc err_code="Duplicate" app_err_desc="Duplicate records found for course and section provided in the input." app_corrective_action="Make sure each course item should be linked with one section" />
			</err_info>
		</exec_function>
		<exec_function function_name="lms_enquiry_course_item_loop" function_type="start_loop" index_name="k" start_index="0" end_index="len(lms_course_item_table_list)" />
		<exec_function function_name="insert_lms_course_progress_copy_object" function_type="copy_object" function_condition="lms_course_item_table_list[k].order!=1||lms_course_section_table_list[j].order!=1">
			<copy_object_info object_name="insert_lms_course_item_progress_additional_obj">
				<object_data key="user_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.id" data_type="string"></object_data>
				<object_data key="course_id"  data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.courses[i].course_id" data_type="string"></object_data>
				<object_data key="section_id"  data_source_type="reqbrokermap" data_source="lms_course_section_table_list[j].id" data_type="string"></object_data>
				<object_data key="item_id"  data_source_type="reqbrokermap" data_source="lms_course_item_table_list[k].id" data_type="string"></object_data>
				<object_data key="is_enabled"  data_source_type="raw_value" data_source="false" data_type="Boolean"></object_data>
				<object_data key="percentage" data_source_type="raw_value" data_source="0" data_type="string"></object_data>
				<object_data key="created_by" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.email" data_type="string"></object_data>
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
		<exec_function function_name="insert_lms_course_progress_copy_object" function_type="copy_object" function_condition="lms_course_section_table_list[j].order==1&amp;&amp;lms_course_item_table_list[k].order==1">
			<copy_object_info object_name="insert_lms_course_item_progress_additional_obj">
				<object_data key="user_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.id" data_type="string"></object_data>
				<object_data key="course_id"  data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.courses[i].course_id" data_type="string"></object_data>
				<object_data key="section_id"  data_source_type="reqbrokermap" data_source="lms_course_section_table_list[j].id" data_type="string"></object_data>
				<object_data key="item_id"  data_source_type="reqbrokermap" data_source="lms_course_item_table_list[k].id" data_type="string"></object_data>
				<object_data key="is_enabled"  data_source_type="raw_value" data_source="true" data_type="Boolean"></object_data>
				<object_data key="percentage" data_source_type="raw_value" data_source="0" data_type="string"></object_data>
				<object_data key="created_by" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.email" data_type="string"></object_data>
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
		<exec_function function_name="insert_lms_course_item_progress_db_insert" function_type="db_insert" function_condition="ReqBrokerReqObj.request_data.courses[i].db_operation=='INSERT'">
			<db_insert_info>
				<table_name>LMS_COURSE_ITEM_PROGRESS</table_name>
				<insert_data_list>
					<insert_data insert_data_name="USER_ID" insert_db_data_type="string" insert_data_source="insert_lms_course_item_progress_additional_obj.user_id" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="COURSE_ID" insert_db_data_type="string" insert_data_source="insert_lms_course_item_progress_additional_obj.course_id" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="SECTION_ID" insert_db_data_type="string" insert_data_source="insert_lms_course_item_progress_additional_obj.section_id" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="ITEM_ID" insert_db_data_type="string" insert_data_source="insert_lms_course_item_progress_additional_obj.item_id" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="IS_ENABLED" insert_db_data_type="boolean" insert_data_source="insert_lms_course_item_progress_additional_obj.is_enabled" insert_data_type="boolean"></insert_data>
					<insert_data insert_data_name="CREATED_BY" insert_db_data_type="string" insert_data_source="insert_lms_course_item_progress_additional_obj.created_by" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="CREATED_DATE" insert_db_data_type="date" insert_data_source="insert_lms_course_item_progress_additional_obj.created_date" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="CREATED_TIME" insert_db_data_type="string" insert_data_source="insert_lms_course_item_progress_additional_obj.created_time" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="UPDATED_BY" insert_db_data_type="string" insert_data_source="insert_lms_course_item_progress_additional_obj.updated_by" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="UPDATED_DATE" insert_db_data_type="date" insert_data_source="insert_lms_course_item_progress_additional_obj.updated_date" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="UPDATED_TIME" insert_db_data_type="string" insert_data_source="insert_lms_course_item_progress_additional_obj.updated_time" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="DELETED_BY" insert_db_data_type="string" insert_data_source="insert_lms_course_item_progress_additional_obj.deleted_by" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="DELETED_DATE" insert_db_data_type="date" insert_data_source="insert_lms_course_item_progress_additional_obj.deleted_date" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="DELETED_TIME" insert_db_data_type="string" insert_data_source="insert_lms_course_item_progress_additional_obj.deleted_time" insert_data_type="string"></insert_data>
				</insert_data_list>
			</db_insert_info>
		</exec_function>
		<exec_function function_name="lms_enquiry_course_item_loop" function_type="end_loop" index_name="k" />
		<exec_function function_name="lms_enquiry_course_section_loop" function_type="end_loop" index_name="j" />
		<exec_function function_name="delete_lms_course_progress_table_delete" function_type="db_delete" function_condition="ReqBrokerReqObj.request_data.courses[i].db_operation=='DELETE'">
			<db_delete_info>
				<table_name>LMS_COURSE_PROGRESS</table_name>
				<delete_filter_list>
					<delete_filter delete_filter_name="USER_ID" delete_filter_db_data_type="string" delete_filter_data_source="ReqBrokerReqObj.request_data.id" delete_filter_data_type="string" delete_filter_condition="AND"></delete_filter>
					<delete_filter delete_filter_name="COURSE_ID" delete_filter_db_data_type="string" delete_filter_data_source="ReqBrokerReqObj.request_data.courses[i].course_id" delete_filter_data_type="string" delete_filter_condition="AND"></delete_filter>
				</delete_filter_list>
			</db_delete_info>
		</exec_function>
		<exec_function function_name="delete_lms_course_progress_table_delete" function_type="db_delete" function_condition="ReqBrokerReqObj.request_data.courses[i].db_operation=='DELETE'">
			<db_delete_info>
				<table_name>LMS_COURSE_SECTION_PROGRESS</table_name>
				<delete_filter_list>
					<delete_filter delete_filter_name="USER_ID" delete_filter_db_data_type="string" delete_filter_data_source="ReqBrokerReqObj.request_data.id" delete_filter_data_type="string" delete_filter_condition="AND"></delete_filter>
					<delete_filter delete_filter_name="COURSE_ID" delete_filter_db_data_type="string" delete_filter_data_source="ReqBrokerReqObj.request_data.courses[i].course_id" delete_filter_data_type="string" delete_filter_condition="AND"></delete_filter>
				</delete_filter_list>
			</db_delete_info>
		</exec_function>
		<exec_function function_name="delete_lms_course_item_progress_table_delete" function_type="db_delete" function_condition="ReqBrokerReqObj.request_data.courses[i].db_operation=='DELETE'">
			<db_delete_info>
				<table_name>LMS_COURSE_ITEM_PROGRESS</table_name>
				<delete_filter_list>
					<delete_filter delete_filter_name="USER_ID" delete_filter_db_data_type="string" delete_filter_data_source="ReqBrokerReqObj.request_data.id" delete_filter_data_type="string" delete_filter_condition="AND"></delete_filter>
					<delete_filter delete_filter_name="COURSE_ID" delete_filter_db_data_type="string" delete_filter_data_source="ReqBrokerReqObj.request_data.courses[i].course_id" delete_filter_data_type="string" delete_filter_condition="AND"></delete_filter>
				</delete_filter_list>
			</db_delete_info>
		</exec_function>
		<exec_function function_name="lms_enquiry_course_loop" function_type="end_loop" index_name="i" />
	</exec_group>
	<exec_group group_name="edit_lms_user_data_compose_response">
	<exec_function function_name="lms_user_data_copy_object" function_type="copy_object">
		<copy_object_info object_name="user_data_response_obj">
				<object_data key="resp_msg" data_source_type="raw_value" data_source="User Data Updated successfully!!!" data_type="string"></object_data>
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
