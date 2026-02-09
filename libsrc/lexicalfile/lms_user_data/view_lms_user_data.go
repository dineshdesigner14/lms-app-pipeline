package lexicalfile_lms_user_data

func GetViewLMSUserDataServiceXML() interface{} {
	lexicalParserStr := `
	<sem_lexical_parser microservice="view_lms_user_data">
	<exec_group group_name="user_token_validation" group_condition="ReqBrokerReqObj.header_data.channel=='web_admin_portal'">
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
	</exec_group>
	<exec_group group_name="view_lms_user_data" group_condition="ReqBrokerReqObj.request_data.page!=nil &amp;&amp; ReqBrokerReqObj.request_data.per_page!=nil &amp;&amp;ReqBrokerReqObj.request_data.search==nil">
		<exec_function function_name="lms_user_data_validate_fld" function_type="validate_fld">
			<validate_info app_err_desc="Validation failed for user data" app_corrective_action="Please validate the data">
				<validate_fld fld_name="ReqBrokerReqObj.request_data.page" fld_data_type="string" fld_min_len="1" fld_max_len="20" fld_type="pattern" fld_pattern="^[0-9]+$" app_err_desc="Invalid Page" app_corrective_action="Only letters are allowed."></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.per_page" fld_data_type="string" fld_min_len="1" fld_max_len="20" fld_type="pattern" fld_pattern="^[0-9]+$" app_err_desc="Invalid Limit" app_corrective_action="Only letters are allowed."></validate_fld>
			</validate_info>
		</exec_function>
		<exec_function function_name="login_response_copy_object" function_type="copy_object">
			<copy_object_info object_name="login_resp_obj">
				<object_data key="current_page" data_source_type="raw_value" data_source="1" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
		<exec_function function_name="calculate_math" function_type="arith_operation">
			<arith_operation_info operation="sub" left_operand="ReqBrokerReqObj.request_data.page" right_operand="login_resp_obj.current_page" dest_object="page" />
		</exec_function>
		<exec_function function_name="calculate_math" function_type="arith_operation">
			<arith_operation_info operation="multi" left_operand="page" right_operand="ReqBrokerReqObj.request_data.per_page" dest_object="offset" />
		</exec_function>
		<exec_function function_name="offset_copy_object" function_type="copy_object">
			<copy_object_info object_name="offset_resp_obj">
				<object_data key="offset" data_source_type="reqbrokermap" data_source="offset" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
		<exec_function function_name="lms_user_info_auth_db_multi_read" function_type="db_multi_read">
			<db_multi_read_info  order_by="UPDATED_DATE,UPDATED_TIME" limit="ReqBrokerReqObj.request_data.per_page" sort_type="desc" offset="offset_resp_obj.offset" app_err_desc="Data Already Exists or No Rows in Result Set" app_corrective_action="check the inputs and provide proper values">
				<table_name>LMS_USER_DATA</table_name>
				<store_obj_name>lms_user_table_record</store_obj_name>
			</db_multi_read_info>
		</exec_function>
		<exec_function function_name="login_response_copy_object" function_type="copy_object">
			<copy_object_info object_name="lms_user_table_record">
				<object_data key="page" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.page" data_type="string"></object_data>
				<object_data key="page_size" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.per_page" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
		<exec_function function_name="view_lms_user_data_compose_response" function_type="compose_response">
			<response_info>
				<resp_obj obj_name="lms_user_table_record" obj_type="object" data_source="lms_user_table_record"></resp_obj>
			</response_info>
		</exec_function>
	</exec_group>
	<exec_group group_name="view_lms_user_data_student" group_condition="ReqBrokerReqObj.request_data.page==nil &amp;&amp; ReqBrokerReqObj.request_data.per_page==nil &amp;&amp; ReqBrokerReqObj.request_data.id!=nil">
		<exec_function function_name="insert_lms_enquiry_course_db_single_read" function_type="db_single_read">
			<db_single_read_info app_err_desc="lms course section data table is not found" app_corrective_action="lms course section data cannot be configured as the table does not exist">
				<table_name>LMS_USER_DATA</table_name>
				<filter_info>
					<filter_data filter_name="ID" filter_id="ReqBrokerReqObj.request_data.id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_user_data_record</store_obj_name>
			</db_single_read_info>
		</exec_function>
		<exec_function function_name="lms_user_data_copy_object" function_type="copy_object">
				<copy_object_info object_name="lms_user_data_table_obj">
					<object_data key="id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.id" data_type="string"></object_data>
                    <object_data key="full_name" data_source_type="reqbrokermap" data_source="lms_user_data_record.full_name" data_type="string"></object_data>
					<object_data key="email" data_source_type="reqbrokermap" data_source="lms_user_data_record.email" data_type="string"></object_data>
					<object_data key="mobile_number" data_source_type="reqbrokermap" data_source="lms_user_data_record.mobile_number" data_type="string"></object_data>
					<object_data key="role" data_source_type="reqbrokermap" data_source="lms_user_data_record.role" data_type="string"></object_data>
					<object_data key="profile_picture_url" data_source_type="reqbrokermap" data_source="lms_user_data_record.profile_picture_url" data_type="string"></object_data>
					<object_data key="date_of_birth" data_source_type="reqbrokermap" data_source="lms_user_data_record.date_of_birth" data_type="string"></object_data>
					<object_data key="status" data_source_type="reqbrokermap" data_source="lms_user_data_record.status" data_type="Boolean"></object_data>
                	<object_data key="gender" data_source_type="reqbrokermap" data_source="lms_user_data_record.gender" data_type="string"></object_data>
					<object_data key="address" data_source_type="reqbrokermap"  data_source="lms_user_data_record.address" data_type="string"></object_data>
					<object_data key="bio" data_source_type="reqbrokermap"  data_source="lms_user_data_record.bio" data_type="string"></object_data>
                    <object_data key="created_by" data_source_type="reqbrokermap" data_source="lms_user_data_record.created_by" data_type="string"></object_data>
					<object_data key="created_date" data_source_type="reqbrokermap" data_source="lms_user_data_record.created_date" data_type="string"></object_data>
					<object_data key="created_time" data_source_type="reqbrokermap" data_source="lms_user_data_record.created_time" data_type="string"></object_data>
					<object_data key="updated_by" data_source_type="reqbrokermap" data_source="lms_user_data_record.updated_by" data_type="string"></object_data>
					<object_data key="updated_date" data_source_type="reqbrokermap" data_source="lms_user_data_record.updated_date" data_type="string"></object_data>
					<object_data key="updated_time" data_source_type="reqbrokermap" data_source="lms_user_data_record.updated_time" data_type="string"></object_data>
					<object_data key="deleted_by" data_source_type="reqbrokermap" data_source="lms_user_data_record.deleted_by" data_type="string"></object_data>
					<object_data key="deleted_date" data_source_type="reqbrokermap" data_source="lms_user_data_record.deleted_date" data_type="string"></object_data>
					<object_data key="deleted_time" data_source_type="reqbrokermap" data_source="lms_user_data_record.deleted_time" data_type="string"></object_data>
					<object_data key="last_login_date" data_source_type="reqbrokermap" data_source="lms_user_data_record.last_login_date" data_type="string"></object_data>
					<object_data key="last_login_time"  data_source_type="reqbrokermap" data_source="lms_user_data_record.last_login_time" data_type="string"></object_data>
				</copy_object_info>
			</exec_function>
		<exec_function function_name="lms_course_item_count" function_type="raw_query" function_condition="lms_user_data_record.role=='STUDENT'">
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
				<store_obj_name>lms_user_data_table_obj.lms_course_item_progress_list</store_obj_name>
			</raw_query_info>
		</exec_function>
       

		<exec_function function_name="lms_course_item_loop" function_type="start_loop" index_name="i" start_index="0" end_index="len(lms_course_item_progress_list)" />

		<exec_function function_name="edit_lms_course_copy_object" function_type="copy_object">
			<copy_object_info object_name="edit_lms_course_additional_obj">
				<object_data key="user_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.id" data_type="string"></object_data>
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
		
		<exec_function function_name="view_lms_user_data_compose_response" function_type="compose_response" function_condition="lms_user_data_record.role=='STUDENT'">
			<response_info>
				<resp_obj obj_name="lms_user_data_table_obj" obj_type="object" data_source="lms_user_data_table_obj"></resp_obj>
			</response_info>
		</exec_function>
		<exec_function function_name="view_lms_user_data_compose_response" function_type="compose_response" function_condition="lms_user_data_record.role!='STUDENT'">
			<response_info>
				<resp_obj obj_name="lms_user_data_table_obj" obj_type="object" data_source="lms_user_data_table_obj"></resp_obj>
			</response_info>
		</exec_function>
	</exec_group>
	<exec_group group_name="view_lms_user_data" group_condition="ReqBrokerReqObj.request_data.page==nil &amp;&amp; ReqBrokerReqObj.request_data.per_page==nil &amp;&amp; ReqBrokerReqObj.request_data.id==nil">
		<exec_function function_name="user_raw_query_info" function_type="raw_query" >
			<raw_query_info>
				<postgres_query_str>SELECT
	COALESCE(e."ID"::text, 'NA') AS enquiry_id,
    u."ID" AS user_id,
    u."FULL_NAME",
    u."EMAIL",
    u."MOBILE_NUMBER",
    u."ROLE",
    u."STATUS",
    u."PROFILE_PICTURE_URL",
    u."DATE_OF_BIRTH",
    u."GENDER",
    u."ADDRESS",
    u."BIO",
    u."CREATED_BY",
    u."CREATED_DATE",
    u."CREATED_TIME",
    u."UPDATED_BY",
    u."UPDATED_DATE",
    u."UPDATED_TIME",
    u."DELETED_BY",
    u."DELETED_DATE",
    u."DELETED_TIME",
    u."PASSWD_RETRY_COUNT",
    u."PASSWD_MODIFIED_DATE",
    u."PASSWD_MODIFIED_TIME",
    u."LAST_LOGIN_DATE",
    u."LAST_LOGIN_TIME",
    u."PASSWD_STATUS"
FROM "LMS"."LMS_USER_DATA" u
LEFT JOIN "LMS"."LMS_ENQUIRY" e
    ON u."MOBILE_NUMBER" = e."PHONE_NUMBER";
</postgres_query_str>
				<store_obj_name>lms_user_data_table_list</store_obj_name>
			</raw_query_info>
		</exec_function>
		<exec_function function_name="view_lms_user_data_compose_response" function_type="compose_response">
			<response_info>
				<resp_obj obj_name="lms_user_data_table_list" obj_type="objectarray" data_source="lms_user_data_table_list"></resp_obj>
			</response_info>
		</exec_function>
	</exec_group>
	<exec_group group_name="view_lms_user_search_data" group_condition="ReqBrokerReqObj.request_data.page!=nil &amp;&amp; ReqBrokerReqObj.request_data.per_page!=nil &amp;&amp; ReqBrokerReqObj.request_data.search!=nil">
		<exec_function function_name="lms_user_data_search_validate_fld" function_type="validate_fld">
			<validate_info app_err_desc="Validation failed for user data" app_corrective_action="Please validate the data">
				<validate_fld fld_name="ReqBrokerReqObj.request_data.page" fld_data_type="string" fld_min_len="1" fld_max_len="20" fld_type="pattern" fld_pattern="^[0-9]+$" app_err_desc="Invalid Page" app_corrective_action="Only letters are allowed."></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.per_page" fld_data_type="string" fld_min_len="1" fld_max_len="20" fld_type="pattern" fld_pattern="^[0-9]+$" app_err_desc="Invalid Limit" app_corrective_action="Only letters are allowed."></validate_fld>
                <validate_fld fld_name="ReqBrokerReqObj.request_data.search" fld_data_type="string" fld_min_len="3" fld_max_len="128" fld_type="pattern" fld_pattern=".*" app_err_desc="Invalid search key" app_corrective_action="Only letters,alphabets are allowed and minimum length of search key is 3."></validate_fld>
			</validate_info>
		</exec_function>
		<exec_function function_name="user_pagination_search_copy_object" function_type="copy_object">
			<copy_object_info object_name="login_resp_obj">
				<object_data key="current_page" data_source_type="raw_value" data_source="1" data_type="string"></object_data>
                <object_data key="search_key" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.search" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
		<exec_function function_name="calculate_math_search" function_type="arith_operation">
			<arith_operation_info operation="sub" left_operand="ReqBrokerReqObj.request_data.page" right_operand="login_resp_obj.current_page" dest_object="page" />
		</exec_function>
		<exec_function function_name="calculate_math_search" function_type="arith_operation">
			<arith_operation_info operation="multi" left_operand="page" right_operand="ReqBrokerReqObj.request_data.per_page" dest_object="offset" />
		</exec_function>
		<exec_function function_name="offset_search_copy_object" function_type="copy_object">
			<copy_object_info object_name="offset_resp_obj">
				<object_data key="offset" data_source_type="reqbrokermap" data_source="offset" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>

		<exec_function function_name="lms_user_info_search_auth_db_multi_read" function_type="db_multi_read">
			<db_multi_read_info  order_by="UPDATED_DATE" limit="ReqBrokerReqObj.request_data.per_page" sort_type="desc" offset="offset_resp_obj.offset" search_str="login_resp_obj.search_key" app_err_desc="Data Already Exists or No Rows in Result Set" app_corrective_action="check the inputs and provide proper values">
				<table_name>LMS_USER_DATA</table_name>
				<store_obj_name>lms_user_data_table_list</store_obj_name>
			</db_multi_read_info>
		</exec_function>
		<exec_function function_name="view_lms_user_data_compose_response" function_type="compose_response">
			<response_info>
				<resp_obj obj_name="lms_user_data_table_list" obj_type="object" data_source="lms_user_data_table_list"></resp_obj>
			</response_info>
		</exec_function>
	</exec_group>
</sem_lexical_parser>
	`
	return lexicalParserStr
}
