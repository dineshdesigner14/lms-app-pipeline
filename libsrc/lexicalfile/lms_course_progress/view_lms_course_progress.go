package lexicalfile_lms_course_progress

func GetViewLMSCourseProgressServiceXML() interface{} {
	lexicalParserStr := `
	<sem_lexical_parser microservice="view_lms_course">
	<exec_group group_name="view_lms_course" group_condition="ReqBrokerReqObj.header_data.channel=='web_admin_portal' &amp;&amp;(ReqBrokerReqObj.request_data.page==nil &amp;&amp; ReqBrokerReqObj.request_data.per_page==nil)">
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
		<exec_function function_name="view_lms_course_db_multi_read" function_type="db_multi_read">
			<db_multi_read_info>
				<table_name>LMS_COURSE_PROGRESS</table_name>
				<store_obj_name>lms_course_progress_table_list</store_obj_name>
			</db_multi_read_info>
		</exec_function>
		<exec_function function_name="lms_course_progress_loop" function_type="start_loop" index_name="i" start_index="0" end_index="len(lms_course_progress_table_list)" />
		<exec_function function_name="view_lms_course_db_multi_read" function_type="db_multi_read">
			<db_multi_read_info>
				<table_name>LMS_COURSE_SECTION</table_name>
				<filter_info>
					<filter_data filter_name="COURSE_ID" filter_id="lms_course_progress_table_list[i].course_id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_course_progress_table_list[i].lms_course_section_table_list</store_obj_name>
			</db_multi_read_info>
		</exec_function>
		<exec_function function_name="lms_course_item_loop" function_type="start_loop" index_name="j" start_index="0" end_index="len(lms_course_progress_table_list[i].lms_course_section_table_list)" />
		<exec_function function_name="view_lms_course_db_multi_read" function_type="db_multi_read">
			<db_multi_read_info>
				<table_name>LMS_COURSE_ITEM</table_name>
				<filter_info>
					<filter_data filter_name="COURSE_ID" filter_id="lms_course_progress_table_list[i].lms_course_section_table_list[j].course_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="SECTION_ID" filter_id="lms_course_progress_table_list[i].lms_course_section_table_list[j].id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_course_progress_table_list[i].lms_course_section_table_list[j].lms_course_item_table_list</store_obj_name>
			</db_multi_read_info>
		</exec_function>
		<exec_function function_name="lms_course_item_loop" function_type="end_loop" index_name="j" />
		<exec_function function_name="lms_course_progress_loop" function_type="end_loop" index_name="i" />

		<exec_function function_name="view_lms_course_compose_response" function_type="compose_response">
			<response_info>
				<resp_obj obj_name="lms_course_progress_table_list" obj_type="objectarray" data_source="lms_course_progress_table_list"></resp_obj>
			</response_info>
		</exec_function>
	</exec_group>
	<exec_group group_name="view_lms_course" group_condition="ReqBrokerReqObj.header_data.channel=='web_client_portal' &amp;&amp;(ReqBrokerReqObj.request_data.page==nil &amp;&amp; ReqBrokerReqObj.request_data.per_page==nil)">
	<exec_function function_name="decode_token_function_type" function_type="decode_token">
			<decode_token_info token_value="ReqBrokerReqObj.header_data.auth_token" token_object="decode_token_obj" app_err_desc="Decode token failed!. Invalid access token." app_corrective_action="Please provide valid access token."/>
		</exec_function>
		<exec_function function_name="role_validate_condition" function_type="validate_condition">
			<validate_condition_info app_err_desc="Invalid Role" app_corrective_action="Only ADMIN allowed to create User.">
				<validate_condition_expression>decode_token_obj.role=='STUDENT'</validate_condition_expression>
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
		<exec_function function_name="login_lms_token_guard_db_single_read" function_type="db_single_read">
			<db_single_read_info app_err_desc="Token guard data not found with Role:{{decode_token_obj.email}}" app_corrective_action="Please configure token guard for Role:{{decode_token_obj.role}}">
				<table_name>LMS_USER_DATA</table_name>
				<filter_info>
					<filter_data filter_name="EMAIL" filter_id="decode_token_obj.email" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_user_data_record</store_obj_name>
			</db_single_read_info>
			<err_info>
				<err_desc err_code="Duplicate" app_err_desc="Duplicate token guard data found for the Role:{{decode_token_obj.role}}." app_corrective_action="Make sure each Role shows up only once in table." />
				<err_desc err_code="NoRows" app_err_desc="Token guard data not found with Role:{{decode_token_obj.role}}" app_corrective_action="Please configure token guard for Role:{{decode_token_obj.role}}" />
			</err_info>
		</exec_function>
		<exec_function function_name="verify_token_function_type" function_type="verify_token">
			<verify_token_info token_value="ReqBrokerReqObj.header_data.auth_token" token_secret="lms_token_guard_record.access_token_secret" token_object="verify_token_obj" app_err_desc="Token Validation failed!. Invalid Access token." app_corrective_action="Please provide valid Access token."/>
		</exec_function>
		<exec_function function_name="view_lms_course_db_multi_read" function_type="db_multi_read">
			<db_multi_read_info>
				<table_name>LMS_COURSE_PROGRESS</table_name>
				<filter_info>
					<filter_data filter_name="USER_ID" filter_id="lms_user_data_record.id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_course_progress_table_list</store_obj_name>
			</db_multi_read_info>
		</exec_function>
		<exec_function function_name="insert_lms_course_progress_copy_object" function_type="copy_object">
			<copy_object_info object_name="view_lms_course_progress">
			</copy_object_info>
		</exec_function>
		<exec_function function_name="course_progress_map_array" function_type="create_map_array">
			<map_array_info object_name="view_lms_course_progress.lms_course_progress" array_size="len(lms_course_progress_table_list)" />
		</exec_function>
		<exec_function function_name="lms_course_progress_loop" function_type="start_loop" index_name="i" start_index="0" end_index="len(lms_course_progress_table_list)" />
			<exec_function function_name="lms_course_db_multi_read" function_type="db_single_read">
			<db_single_read_info app_err_desc="Data Already Exists or No Rows in Result Set" app_corrective_action="check the inputs and provide proper values">
				<table_name>LMS_COURSE</table_name>
				<filter_info>
					<filter_data filter_name="ID" filter_id="lms_course_progress_table_list[i].course_id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_course_table_record</store_obj_name>
			</db_single_read_info>
		</exec_function>
		<exec_function function_name="edit_lms_course_copy_object" function_type="copy_object">
			<copy_object_info object_name="view_lms_course_progress.lms_course_progress[i]">
				<object_data key="user_id" data_source_type="reqbrokermap" data_source="lms_course_progress_table_list[i].user_id" data_type="string"></object_data>
				<object_data key="name" data_source_type="reqbrokermap" data_source="lms_user_data_record.full_name" data_type="string"></object_data>
				<object_data key="course_id" data_source_type="reqbrokermap" data_source="lms_course_progress_table_list[i].course_id" data_type="string"></object_data>
				<object_data key="course_name" data_source_type="reqbrokermap" data_source="lms_course_table_record.course_name" data_type="string"></object_data>
				<object_data key="course_duration" data_source_type="reqbrokermap" data_source="lms_course_table_record.duration_min" data_type="string"></object_data>
				<object_data key="course_image" data_source_type="reqbrokermap" data_source="lms_course_table_record.course_img" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>

		<exec_function function_name="view_lms_course_db_multi_read" function_type="db_multi_read">
			<db_multi_read_info>
				<table_name>LMS_COURSE_SECTION_PROGRESS</table_name>
				<filter_info>
					<filter_data filter_name="USER_ID" filter_id="view_lms_course_progress.lms_course_progress[i].user_id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_course_section_progress_list</store_obj_name>
			</db_multi_read_info>
		</exec_function>

        <exec_function function_name="course_progress_map_array" function_type="create_map_array">
			<map_array_info object_name="view_lms_course_progress.lms_section_progress" array_size="len(lms_course_section_progress_list)" />
		</exec_function>

		<exec_function function_name="lms_course_section_progress_loop" function_type="start_loop" index_name="j" start_index="0" end_index="len(lms_course_section_progress_list)" />
         <exec_function function_name="view_lms_course_db_multi_read" function_type="db_single_read">
			<db_single_read_info>
				<table_name>LMS_COURSE_SECTION</table_name>
				<filter_info>
					<filter_data filter_name="COURSE_ID" filter_id="lms_course_section_progress_list[j].course_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="ID" filter_id="lms_course_section_progress_list[j].section_id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_course_section_table_record</store_obj_name>
			</db_single_read_info>
		</exec_function>

		<exec_function function_name="lms_course_section_copy_object" function_type="copy_object">
			<copy_object_info object_name="view_lms_course_progress.lms_section_progress[j]">
				<object_data key="course_id" data_source_type="reqbrokermap" data_source="lms_course_section_table_record.course_id" data_type="string"></object_data>
				<object_data key="section_id" data_source_type="reqbrokermap" data_source="lms_course_section_table_record.id" data_type="string"></object_data>
				<object_data key="title" data_source_type="reqbrokermap" data_source="lms_course_section_table_record.title" data_type="string"></object_data>
				<object_data key="description" data_source_type="reqbrokermap" data_source="lms_course_section_table_record.description" data_type="string"></object_data>
				<object_data key="order" data_source_type="reqbrokermap" data_source="lms_course_section_table_record.order" data_type="string"></object_data>
				<object_data key="is_completed" data_source_type="reqbrokermap" data_source="lms_course_section_progress_list[j].is_completed" data_type="Boolean"></object_data>
				<object_data key="is_enabled" data_source_type="reqbrokermap" data_source="lms_course_section_progress_list[j].is_enabled" data_type="Boolean"></object_data>
			</copy_object_info>
		</exec_function>
		<exec_function function_name="view_lms_course_db_multi_read" function_type="db_multi_read">
			<db_multi_read_info>
				<table_name>LMS_COURSE_ITEM_PROGRESS</table_name>
				<filter_info>
					<filter_data filter_name="USER_ID" filter_id="lms_user_data_record.id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_course_item_progress_list</store_obj_name>
			</db_multi_read_info>
		</exec_function>

		<exec_function function_name="course_progress_map_array" function_type="create_map_array">
			<map_array_info object_name="view_lms_course_progress.lms_course_item" array_size="len(lms_course_item_progress_list)" />
		</exec_function>

		<exec_function function_name="lms_course_item_progress_loop" function_type="start_loop" index_name="k" start_index="0" end_index="len(lms_course_item_progress_list)" />


		<exec_function function_name="view_lms_course_db_multi_read" function_type="db_single_read">
			<db_single_read_info>
				<table_name>LMS_COURSE_ITEM</table_name>
				<filter_info>
					<filter_data filter_name="COURSE_ID" filter_id="lms_course_item_progress_list[k].course_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="SECTION_ID" filter_id="lms_course_item_progress_list[k].section_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="ID" filter_id="lms_course_item_progress_list[k].item_id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_course_item_table_record</store_obj_name>
			</db_single_read_info>
		</exec_function>


		<exec_function function_name="lms_course_item_copy_object" function_type="copy_object">
			<copy_object_info object_name="view_lms_course_progress.lms_course_item[k]">
				<object_data key="course_id" data_source_type="reqbrokermap" data_source="lms_course_item_table_record.course_id" data_type="string"></object_data>
				<object_data key="section_id" data_source_type="reqbrokermap" data_source="lms_course_item_table_record.section_id" data_type="string"></object_data>
				<object_data key="item_id" data_source_type="reqbrokermap" data_source="lms_course_item_table_record.id" data_type="string"></object_data>
				<object_data key="title" data_source_type="reqbrokermap" data_source="lms_course_item_table_record.title" data_type="string"></object_data>
				<object_data key="description" data_source_type="reqbrokermap" data_source="lms_course_item_table_record.description" data_type="string"></object_data>
				<object_data key="order" data_source_type="reqbrokermap" data_source="lms_course_item_table_record.order" data_type="string"></object_data>
				<object_data key="is_completed" data_source_type="reqbrokermap" data_source="lms_course_item_progress_list[k].is_completed" data_type="Boolean"></object_data>
				<object_data key="is_enabled" data_source_type="reqbrokermap" data_source="lms_course_item_progress_list[k].is_enabled" data_type="Boolean"></object_data>
			</copy_object_info>
		</exec_function>


		<exec_function function_name="lms_course_item_progress_loop" function_type="end_loop" index_name="k" />


		<exec_function function_name="lms_course_section_progress_loop" function_type="end_loop" index_name="j" />





		<exec_function function_name="lms_course_progress_loop" function_type="end_loop" index_name="i" />








		

		<exec_function function_name="view_lms_course_compose_response" function_type="compose_response">
			<response_info>
				<resp_obj obj_name="view_lms_course_progress" obj_type="object" data_source="view_lms_course_progress"></resp_obj>
			</response_info>
		</exec_function>
	</exec_group>
	   <exec_group group_name="view_lms_course_data_pagination" group_condition="(ReqBrokerReqObj.request_data.page!=nil &amp;&amp; ReqBrokerReqObj.request_data.per_page!=nil)&amp;&amp; ReqBrokerReqObj.request_data.search==nil">
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
		 <exec_function function_name="lms_course_db_multi_read" function_type="db_multi_read">
			<db_multi_read_info  order_by="UPDATED_DATE" limit="ReqBrokerReqObj.request_data.per_page" sort_type="desc" offset="offset_resp_obj.offset" app_err_desc="Data Already Exists or No Rows in Result Set" app_corrective_action="check the inputs and provide proper values">
		     <table_name>LMS_COURSE_PROGRESS</table_name>
				<store_obj_name>lms_course_progress_table_record</store_obj_name>
			</db_multi_read_info>
		</exec_function> 
		    <exec_function function_name="login_response_copy_object" function_type="copy_object">
			<copy_object_info object_name="lms_course_progress_table_record">
					<object_data key="page" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.page" data_type="string"></object_data>
                    <object_data key="page_size" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.per_page" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
			<exec_function function_name="view_lms_user_data_compose_response" function_type="compose_response">
				<response_info>
					<resp_obj obj_name="lms_course_progress_table_record" obj_type="object" data_source="lms_course_progress_table_record"></resp_obj>
				</response_info>
			</exec_function>
</exec_group>
    <exec_group group_name="view_lms_course_progress_search_data" group_condition="(ReqBrokerReqObj.request_data.page!=nil &amp;&amp; ReqBrokerReqObj.request_data.per_page!=nil) &amp;&amp; ReqBrokerReqObj.request_data.search!=nil">
		<exec_function function_name="lms_user_data_search_validate_fld" function_type="validate_fld">
			<validate_info app_err_desc="Validation failed for user data" app_corrective_action="Please validate the data">
				<validate_fld fld_name="ReqBrokerReqObj.request_data.page" fld_data_type="string" fld_min_len="1" fld_max_len="20" fld_type="pattern" fld_pattern="^[0-9]+$" app_err_desc="Invalid Page" app_corrective_action="Only letters are allowed."></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.per_page" fld_data_type="string" fld_min_len="1" fld_max_len="20" fld_type="pattern" fld_pattern="^[0-9]+$" app_err_desc="Invalid Limit" app_corrective_action="Only letters are allowed."></validate_fld>
                <validate_fld fld_name="ReqBrokerReqObj.request_data.search" fld_data_type="string" fld_min_len="3" fld_max_len="20" fld_type="pattern" fld_pattern=".*" app_err_desc="Invalid search key" app_corrective_action="Only letters,alphabets are allowed and minimum length of search key is 3."></validate_fld>
			</validate_info>
		</exec_function>
		<exec_function function_name="course_progress_pagination_search_copy_object" function_type="copy_object">
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

		<exec_function function_name="lms_course_pagination_search_auth_db_multi_read" function_type="db_multi_read">
			<db_multi_read_info  order_by="UPDATED_DATE,UPDATED_TIME" limit="ReqBrokerReqObj.request_data.per_page" sort_type="desc" offset="offset_resp_obj.offset" search_str="login_resp_obj.search_key" app_err_desc="Data Already Exists or No Rows in Result Set" app_corrective_action="check the inputs and provide proper values">
				<table_name>LMS_COURSE_PROGRESS</table_name>
				<store_obj_name>lms_course_progress_table_list</store_obj_name>
			</db_multi_read_info>
		</exec_function>
		<exec_function function_name="view_lms_course_data_compose_response" function_type="compose_response">
			<response_info>
				<resp_obj obj_name="lms_course_progress_table_list" obj_type="object" data_source="lms_course_progress_table_list"></resp_obj>
			</response_info>
		</exec_function>
	</exec_group>
</sem_lexical_parser>
	`
	return lexicalParserStr
}
