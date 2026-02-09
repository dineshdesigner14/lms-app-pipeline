package lexicalfile_lms_course_item

func GetViewLMSCourseItemProdServiceXML() interface{} {
	lexicalParserStr := `
	<sem_lexical_parser microservice="view_lms_course_item">
	<exec_group group_name="view_lms_course_section_data" group_condition="ReqBrokerReqObj.header_data.channel=='web_admin_portal' &amp;&amp; ReqBrokerReqObj.request_data.page!=nil &amp;&amp; ReqBrokerReqObj.request_data.per_page!=nil">
		<exec_function function_name="lms_course_section_validate_fld" function_type="validate_fld">
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
			<db_multi_read_info  order_by="UPDATED_DATE" limit="ReqBrokerReqObj.request_data.per_page" sort_type="desc" offset="offset_resp_obj.offset" app_err_desc="Data Already Exists or No Rows in Result Set" app_corrective_action="check the inputs and provide proper values">
				<table_name>LMS_COURSE_ITEM</table_name>
				<store_obj_name>lms_course_item_table_record</store_obj_name>
			</db_multi_read_info>
		</exec_function>
		<exec_function function_name="login_response_copy_object" function_type="copy_object">
			<copy_object_info object_name="lms_course_item_table_record">
				<object_data key="page" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.page" data_type="string"></object_data>
				<object_data key="page_size" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.per_page" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
		<exec_function function_name="view_lms_user_data_compose_response" function_type="compose_response">
			<response_info>
				<resp_obj obj_name="lms_course_item_table_record" obj_type="object" data_source="lms_course_item_table_record"></resp_obj>
			</response_info>
		</exec_function>
	</exec_group>
	<exec_group group_name="validate_view_lms_course_item"  group_condition="ReqBrokerReqObj.header_data.channel=='web_admin_portal' &amp;&amp; (ReqBrokerReqObj.request_data.course_id==nil &amp;&amp; ReqBrokerReqObj.request_data.section_id==nil &amp;&amp; ReqBrokerReqObj.request_data.id==nil) &amp;&amp; (ReqBrokerReqObj.request_data.page==nil &amp;&amp; ReqBrokerReqObj.request_data.per_page==nil)" >
		<exec_function function_name="read_view_lms_course_item_db_multi_read" function_type="db_multi_read">
			<db_multi_read_info>
				<table_name>LMS_COURSE_SECTION</table_name>
				<store_obj_name>lms_course_section_table_list</store_obj_name>
			</db_multi_read_info>
		</exec_function>
		<exec_function function_name="lms_course_item_loop" function_type="start_loop" index_name="i" start_index="0" end_index="len(lms_course_section_table_list)" />
		<exec_function function_name="read_view_lms_course_item_db_multi_read" function_type="db_multi_read">
			<db_multi_read_info>
				<table_name>LMS_COURSE_ITEM</table_name>
				<filter_info>
					<filter_data filter_name="COURSE_ID" filter_id="lms_course_section_table_list[i].course_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="SECTION_ID" filter_id="lms_course_section_table_list[i].id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_course_section_table_list[i].courses_item</store_obj_name>
			</db_multi_read_info>
		</exec_function>
		<exec_function function_name="lms_course_item_loop" function_type="end_loop" index_name="i" />
		<exec_function function_name="view_lms_course_item_compose_response" function_type="compose_response">
			<response_info>
				<resp_obj obj_name="lms_course_section_table_list" obj_type="objectarray" data_source="lms_course_section_table_list"></resp_obj>
			</response_info>
		</exec_function>
	</exec_group>
	<exec_group group_name="validate_view_lms_course_item"  group_condition="ReqBrokerReqObj.header_data.channel=='web_admin_portal' &amp;&amp; ReqBrokerReqObj.request_data.course_id!=nil &amp;&amp; ReqBrokerReqObj.request_data.section_id!=nil" >
		<exec_function function_name="read_view_lms_course_item_db_multi_read" function_type="db_multi_read">
			<db_multi_read_info>
				<table_name>LMS_COURSE_ITEM</table_name>
				<filter_info>
					<filter_data filter_name="COURSE_ID" filter_id="ReqBrokerReqObj.request_data.course_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="SECTION_ID" filter_id="ReqBrokerReqObj.request_data.section_id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_course_section_item_table_list</store_obj_name>
			</db_multi_read_info>
		</exec_function>
		<exec_function function_name="view_lms_course_item_compose_response" function_type="compose_response">
			<response_info>
				<resp_obj obj_name="lms_course_section_item_table_list" obj_type="objectarray" data_source="lms_course_section_item_table_list"></resp_obj>
			</response_info>
		</exec_function>
	</exec_group>
	<exec_group group_name="validate_view_lms_course_item"  group_condition="ReqBrokerReqObj.header_data.channel=='web_admin_portal' &amp;&amp; ReqBrokerReqObj.request_data.course_id!=nil &amp;&amp; ReqBrokerReqObj.request_data.section_id!=nil &amp;&amp; ReqBrokerReqObj.request_data.id!=nil" >
		<exec_function function_name="read_view_lms_course_item_db_multi_read" function_type="db_multi_read">
			<db_multi_read_info>
				<table_name>LMS_COURSE_ITEM</table_name>
				<filter_info>
					<filter_data filter_name="COURSE_ID" filter_id="ReqBrokerReqObj.request_data.course_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="SECTION_ID" filter_id="ReqBrokerReqObj.request_data.section_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="ID" filter_id="ReqBrokerReqObj.request_data.id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_course_section_item_table_list</store_obj_name>
			</db_multi_read_info>
		</exec_function>
		<exec_function function_name="view_lms_course_item_compose_response" function_type="compose_response">
			<response_info>
				<resp_obj obj_name="lms_course_section_item_table_list" obj_type="objectarray" data_source="lms_course_section_item_table_list"></resp_obj>
			</response_info>
		</exec_function>
	</exec_group>
</sem_lexical_parser>
	`
	return lexicalParserStr
}
