package lexicalfile_lms_course

func GetViewLMSCourseCurrentUserServiceXML() interface{} {
	lexicalParserStr := `
	<sem_lexical_parser microservice="view_lms_course_current_user">
	<exec_group group_name="view_lms_course_student" group_condition="ReqBrokerReqObj.header_data.channel=='web_client_portal'">
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

		<exec_function function_name="user_data_db_single_read" function_type="db_single_read">
			<db_single_read_info app_err_desc="User Data not found" app_corrective_action="Please configure the user data">
				<table_name>LMS_USER_DATA</table_name>
				<filter_info>
					<filter_data filter_name="EMAIL" filter_id="decode_token_obj.email" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_user_data_record</store_obj_name>
			</db_single_read_info>
			<err_info>
				<err_desc err_code="Duplicate" app_err_desc="Duplicate Record for the User Data" app_corrective_action="Make sure each Role shows up only once in table." />   
				<err_desc err_code="NoRows" app_err_desc="No Data Found in the User Data Table" app_corrective_action="Please configure the User Data" />
			</err_info>
		</exec_function>



		<exec_function function_name="lms_course_item_count" function_type="raw_query" >
			<raw_query_info>
				<postgres_query_str>SELECT
    c."ID" AS course_id,
    c."COURSE_NAME",
    ROUND(
        (
            COUNT(*) FILTER (
                WHERE cip."IS_COMPLETED" = true
            )::DECIMAL
            / NULLIF(COUNT(*), 0)
        ) * 100
    ) AS course_percentage
FROM "LMS_COURSE_ITEM_PROGRESS" cip
JOIN "LMS_COURSE" c
    ON c."ID" = cip."COURSE_ID"
WHERE cip."USER_ID" = '{{.lms_user_data_record.id}}'
GROUP BY c."ID", c."COURSE_NAME";
</postgres_query_str>
				<store_obj_name>lms_course_item_progress_list</store_obj_name>
			</raw_query_info>
		</exec_function>


		<exec_function function_name="lms_course_item_loop" function_type="start_loop" index_name="i" start_index="0" end_index="len(lms_course_item_progress_list)" />

		<exec_function function_name="edit_lms_course_copy_object" function_type="copy_object">
			<copy_object_info object_name="edit_lms_course_additional_obj">
				<object_data key="user_id" data_source_type="reqbrokermap" data_source="lms_user_data_record.id" data_type="string"></object_data>
				<object_data key="course_id" data_source_type="reqbrokermap" data_source="lms_course_item_progress_list[i].course_id" data_type="string"></object_data>
				<object_data key="percentage" data_source_type="reqbrokermap" data_source="lms_course_item_progress_list[i].course_percentage" data_type="string"></object_data>
				<object_data key="updated_by" data_source_type="reqbrokermap" data_source="decode_token_obj.email" data_type="string"></object_data>
				<object_data key="updated_date" data_source_type="key" data_source="get_date" data_type="string"></object_data>
				<object_data key="updated_time" data_source_type="key" data_source="get_time" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>


		<exec_function function_name="edit_lms_course_db_update" function_type="db_update">
			<db_update_info>
				<table_name>LMS_COURSE_PROGRESS</table_name>
				<update_data_list>
					<update_data update_data_name="PERCENTAGE" update_db_data_type="string" update_data_source="edit_lms_course_additional_obj.percentage" update_data_type="string"></update_data>
					<update_data update_data_name="UPDATED_DATE" update_db_data_type="date" update_data_source="edit_lms_course_additional_obj.updated_date" update_data_type="string"></update_data>
					<update_data update_data_name="UPDATED_TIME" update_db_data_type="string" update_data_source="edit_lms_course_additional_obj.updated_time" update_data_type="string"></update_data>
				</update_data_list>
				<update_filter_list>
					<update_filter update_filter_name="USER_ID" update_filter_db_data_type="string" update_filter_data_source="edit_lms_course_additional_obj.user_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
					<update_filter update_filter_name="COURSE_ID" update_filter_db_data_type="string" update_filter_data_source="edit_lms_course_additional_obj.course_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
				</update_filter_list>
			</db_update_info>
		</exec_function>


		<exec_function function_name="lms_course_item_loop" function_type="end_loop" index_name="i" />



		<exec_function function_name="insert_lms_course_compose_response" function_type="compose_response">
			<response_info>
				<resp_obj obj_name="lms_course_item_progress_list" obj_type="objectarray" data_source="lms_course_item_progress_list"></resp_obj>
			</response_info>
		</exec_function>
	</exec_group>
</sem_lexical_parser>
	`
	return lexicalParserStr
}
