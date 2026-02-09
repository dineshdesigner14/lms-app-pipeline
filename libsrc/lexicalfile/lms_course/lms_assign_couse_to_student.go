package lexicalfile_lms_course

func GetLMSAssignCourseToStudentServiceXML() interface{} {
	lexicalParserStr := `
	<sem_lexical_parser microservice="lms_assign_course_to_student">
	<exec_group group_name="lms_assign_course_to_student" group_condition="ReqBrokerReqObj.header_data.channel=='web_admin_portal' &amp;&amp;(ReqBrokerReqObj.request_data.page==nil &amp;&amp; ReqBrokerReqObj.request_data.per_page==nil)">
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
		<exec_function function_name="view_lms_course_db_multi_read" function_type="db_single_read">
			<db_single_read_info>
				<table_name>LMS_USER_DATA</table_name>
				<filter_info>
					<filter_data filter_name="ID" filter_id="ReqBrokerReqObj.request_data.user_id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_user_data_table_record</store_obj_name>
			</db_single_read_info>
		</exec_function>
		<exec_function function_name="lms_assign_course_to_user_copy_object" function_type="copy_object">
			<copy_object_info object_name="lms_assign_course_to_user">
				<object_data key="user_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.user_id" data_type="string"></object_data>
				<object_data key="course_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.course_id" data_type="string"></object_data>
				<object_data key="updated_by" data_source_type="reqbrokermap" data_source="decode_token_obj.email" data_type="string"></object_data>
				<object_data key="updated_date" data_source_type="key" data_source="get_date" data_type="string"></object_data>
				<object_data key="updated_time" data_source_type="key" data_source="get_time" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
		<exec_function function_name="edit_lms_course_db_update" function_type="db_update">
			<db_update_info>
				<table_name>LMS_USER_DATA</table_name>
				<update_data_list>
					<update_data update_data_name="ENROLLMENT_ID" update_db_data_type="string" update_data_source="lms_assign_course_to_user.course_id" update_data_type="string"></update_data>
					<update_data update_data_name="UPDATED_BY" update_db_data_type="string" update_data_source="lms_assign_course_to_user.updated_by" update_data_type="string"></update_data>
					<update_data update_data_name="UPDATED_DATE" update_db_data_type="date" update_data_source="lms_assign_course_to_user.updated_date" update_data_type="string"></update_data>
					<update_data update_data_name="UPDATED_TIME" update_db_data_type="string" update_data_source="lms_assign_course_to_user.updated_time" update_data_type="string"></update_data>
				</update_data_list>
				<update_filter_list>
					<update_filter update_filter_name="ID" update_filter_db_data_type="string" update_filter_data_source="ReqBrokerReqObj.request_data.user_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
				</update_filter_list>
			</db_update_info>
		</exec_function>
		<exec_function function_name="view_lms_course_db_multi_read" function_type="db_single_read">
			<db_single_read_info>
				<table_name>LMS_USER_DATA</table_name>
				<filter_info>
					<filter_data filter_name="ID" filter_id="ReqBrokerReqObj.request_data.user_id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_user_data_table_record</store_obj_name>
			</db_single_read_info>
		</exec_function>
		<exec_function function_name="lms_assign_course_to_user_copy_object" function_type="copy_object">
			<copy_object_info object_name="lms_assign_course_copy_object">
				<object_data key="user_id" data_source_type="reqbrokermap" data_source="lms_user_data_table_record.id" data_type="string"></object_data>
				<object_data key="enrollment_id" data_source_type="reqbrokermap" data_source="lms_user_data_table_record.enrollment_id" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
		<exec_function function_name="view_lms_course_compose_response" function_type="compose_response">
			<response_info>
				<resp_obj obj_name="lms_assign_course_copy_object" obj_type="object" data_source="lms_assign_course_copy_object"></resp_obj>
			</response_info>
		</exec_function>
	</exec_group>
</sem_lexical_parser>
	`
	return lexicalParserStr
}
