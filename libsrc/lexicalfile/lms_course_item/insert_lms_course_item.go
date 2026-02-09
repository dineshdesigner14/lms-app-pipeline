package lexicalfile_lms_course_item

func GetInsertLMSCourseItemServiceXML() interface{} {
	lexicalParserStr := `
	<sem_lexical_parser microservice="insert_lms_course_item">
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
	<exec_group group_name="insert_lms_course_item" group_condition="ReqBrokerReqObj.header_data.channel=='web_admin_portal'">
		<exec_function function_name="insert_lms_course_item_validate_fld" function_type="validate_fld">
			<validate_info app_err_desc="Validation failed for course item data" app_corrective_action="Please validate the data">
				<validate_fld fld_name="ReqBrokerReqObj.request_data.course_id" fld_data_type="string" fld_min_len="1" fld_max_len="255" fld_type="pattern" fld_pattern="^[0-9.]+$" app_err_desc="Invalid course id" app_corrective_action="Only numbers are allowed."></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.section_id" fld_data_type="string" fld_min_len="1" fld_max_len="255" fld_type="pattern" fld_pattern="^[0-9.]+$" app_err_desc="Invalid section id" app_corrective_action="Only numbers are allowed."></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.title" fld_data_type="string" fld_min_len="1" fld_max_len="255" fld_type="pattern" fld_pattern=".*" app_err_desc="Invalid title" app_corrective_action="Only letters, numbers and spaces allowed."></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.description" fld_data_type="string" fld_min_len="1" fld_max_len="155" fld_type="pattern" fld_pattern=".*" app_err_desc="Invalid description" app_corrective_action="Only letters, numbers and spaces allowed."></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.order" fld_data_type="string" fld_min_len="1" fld_max_len="155" fld_type="pattern" fld_pattern="^[0-9.]+$" app_err_desc="Invalid order" app_corrective_action="Only numbers are allowed."></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.is_locked" fld_data_type="boolean" fld_min_len="1" fld_max_len="10" fld_type="pattern" fld_pattern="^(true|false|0|1)$" app_err_desc="Invalid is locked" app_corrective_action="Only Boolean value is allowed."></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.unlock_condition" fld_data_type="string" fld_min_len="1" fld_max_len="155" fld_type="pattern" fld_pattern="^(PREVIOUS_SECTION_COMPLETE|MANUAL)$" app_err_desc="Invalid unlock condition" app_corrective_action="Only PREVIOUS_SECTION_COMPLETE,MANUAL allowed"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.item_type" fld_data_type="string" fld_min_len="1" fld_max_len="155" fld_type="pattern" fld_pattern="[a-zA-Z0-9.,|_* ]+$" app_err_desc="Invalid item type" app_corrective_action="Only letters, numbers and spaces allowed."></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.is_encrypt" fld_data_type="boolean" fld_min_len="1" fld_max_len="10" fld_type="pattern" fld_pattern="^(true|false|0|1)$" app_err_desc="Invalid is encrypt" app_corrective_action="Only Boolean value is allowed."></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.is_required" fld_data_type="boolean" fld_min_len="1" fld_max_len="10" fld_type="pattern" fld_pattern="^(true|false|0|1)$" app_err_desc="Invalid is required" app_corrective_action="Only Boolean value is allowed."></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.estimate_time_min" fld_data_type="string" fld_min_len="1" fld_max_len="255" fld_type="pattern" fld_pattern="[a-zA-Z0-9.,|_* ]+$" app_err_desc="Invalid estimate time min" app_corrective_action="Only letters, numbers and spaces allowed."></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.req_complete_percentage" fld_data_type="string" fld_min_len="1" fld_max_len="255" fld_type="pattern" fld_pattern="[a-zA-Z0-9.,|_* ]+$" app_err_desc="Invalid required completion percentage" app_corrective_action="Only letters, numbers and spaces allowed."></validate_fld>
			</validate_info>
		</exec_function>
		<exec_function function_name="insert_lms_course_item_db_single_read" function_type="db_single_read">
			<db_single_read_info app_err_desc="lms course item table is not found" app_corrective_action="course item cannot be configured as the table does not exist">
				<table_name>LMS_COURSE</table_name>
				<filter_info>
					<filter_data filter_name="ID" filter_id="ReqBrokerReqObj.request_data.course_id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_course_table_record</store_obj_name>
			</db_single_read_info>
			<err_info>
				<err_desc err_code="Duplicate" app_err_desc="Duplicate records found for the title:{{.ReqBrokerReqObj.request_data.title}} provided in the input." app_corrective_action="Make sure each title shows up only once." />
			</err_info>
		</exec_function>
		<exec_function function_name="insert_lms_course_item_db_single_read" function_type="db_single_read">
			<db_single_read_info app_err_desc="lms course item table is not found" app_corrective_action="course item cannot be configured as the table does not exist">
				<table_name>LMS_COURSE_SECTION</table_name>
				<filter_info>
					<filter_data filter_name="COURSE_ID" filter_id="ReqBrokerReqObj.request_data.course_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="ID" filter_id="ReqBrokerReqObj.request_data.section_id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_course_section_table_record</store_obj_name>
			</db_single_read_info>
			<err_info>
				<err_desc err_code="Duplicate" app_err_desc="Duplicate records found for the title:{{.ReqBrokerReqObj.request_data.title}} provided in the input." app_corrective_action="Make sure each title shows up only once." />
			</err_info>
		</exec_function>
		<exec_function function_name="insert_lms_course_item_db_single_read" function_type="db_multi_read">
			<db_multi_read_info app_err_desc="lms course data table is not found" app_corrective_action="lms course data cannot be configured as the table does not exist">
				<table_name>LMS_COURSE_ITEM</table_name>
				<filter_info>
					<filter_data filter_name="COURSE_ID" filter_id="ReqBrokerReqObj.request_data.course_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="SECTION_ID" filter_id="ReqBrokerReqObj.request_data.section_id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_course_item_table_list</store_obj_name>
			</db_multi_read_info>
		</exec_function>
		<exec_function function_name="lms_course_item_loop" function_type="start_loop" index_name="i" start_index="0" end_index="len(lms_course_item_table_list)" />
		<exec_function function_name="insert_lms_course_item_db_single_read" function_type="db_single_read">
			<db_single_read_info app_err_desc="Course Item Order Already Exists" app_corrective_action="Please Configure Different Order and Proceed">
				<table_name>LMS_COURSE_ITEM</table_name>
				<filter_info>
					<filter_data filter_name="COURSE_ID" filter_id="ReqBrokerReqObj.request_data.course_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="SECTION_ID" filter_id="lms_course_item_table_list[i].section_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="ID" filter_id="lms_course_item_table_list[i].id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="ORDER" filter_id="ReqBrokerReqObj.request_data.order" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<result_info app_err_desc="The Course Item Order:{{.ReqBrokerReqObj.request_data.order}} which is already exists." app_corrective_action="Please configure the new course item order and proceed">
					<result_success>NoRows</result_success>
				</result_info>
			</db_single_read_info>
		</exec_function>
		<exec_function function_name="lms_course_item_loop" function_type="end_loop" index_name="i" />
		<exec_function function_name="insert_lms_course_item_db_single_read" function_type="db_single_read">
			<db_single_read_info app_err_desc="lms course item data table is not found" app_corrective_action="course item data cannot be configured as the table does not exist">
				<table_name>LMS_COURSE_ITEM</table_name>
				<filter_info>
					<filter_data filter_name="COURSE_ID" filter_id="ReqBrokerReqObj.request_data.course_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="SECTION_ID" filter_id="ReqBrokerReqObj.request_data.section_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="TITLE" filter_id="ReqBrokerReqObj.request_data.title" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<result_info app_err_desc="The Course Item Title:{{.ReqBrokerReqObj.request_data.title}} already exists." app_corrective_action="Please configure new course item title and proceed">
					<result_success>NoRows</result_success>
				</result_info>
			</db_single_read_info>
			<err_info>
				<err_desc err_code="Duplicate" app_err_desc="Duplicate records found for the course item provided in the input." app_corrective_action="Make sure each course item shows up only once." />
			</err_info>
		</exec_function>
		<exec_function function_name="lms_course_item_copy_object" function_type="copy_object">
			<copy_object_info object_name="lms_course_item_table_obj">
				<object_data key="course_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.course_id" data_type="string"></object_data>
				<object_data key="section_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.section_id" data_type="string"></object_data>
				<object_data key="title" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.title" data_type="string"></object_data>
				<object_data key="description" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.description" data_type="string"></object_data>
				<object_data key="order" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.order" data_type="string"></object_data>
				<object_data key="is_locked" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.is_locked" data_type="Boolean"></object_data>
				<object_data key="unlock_condition" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.unlock_condition" data_type="string"></object_data>
				<object_data key="is_encrypt" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.is_encrypt" data_type="Boolean"></object_data>
				<object_data key="is_required" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.is_required" data_type="Boolean"></object_data>
				<object_data key="drm_type" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.drm_type" data_type="string"></object_data>
				<object_data key="estimate_time_min" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.estimate_time_min" data_type="string"></object_data>
				<object_data key="duration_in_min" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.duration_in_min" data_type="string"></object_data>
				<object_data key="item_type" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.item_type" data_type="string"></object_data>
				<object_data key="video_url" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.video_url" data_type="string"></object_data>
				<object_data key="hls_1080_url" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.hls_1080_url" data_type="string"></object_data>
				<object_data key="hls_720_url" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.hls_720_url" data_type="string"></object_data>
				<object_data key="hls_480_url" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.hls_480_url" data_type="string"></object_data>
				<object_data key="pdf_url" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.pdf_url" data_type="string"></object_data>
				<object_data key="ppt_url" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.ppt_url" data_type="string"></object_data>
				<object_data key="html_content" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.html_content" data_type="string"></object_data>
				<object_data key="req_complete_percentage" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.req_complete_percentage" data_type="string"></object_data>
				<object_data key="created_by" data_source_type="reqbrokermap" data_source="decode_token_obj.email" data_type="string"></object_data>
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
		<exec_function function_name="insert_lms_course_item_db_insert" function_type="db_insert">
			<db_insert_info>
				<table_name>LMS_COURSE_ITEM</table_name>
				<insert_data_list>
					<insert_data insert_data_name="COURSE_ID" insert_db_data_type="string" insert_data_source="lms_course_item_table_obj.course_id" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="SECTION_ID" insert_db_data_type="string" insert_data_source="lms_course_item_table_obj.section_id" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="ITEM_TYPE" insert_db_data_type="string" insert_data_source="lms_course_item_table_obj.item_type" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="TITLE" insert_db_data_type="string" insert_data_source="lms_course_item_table_obj.title" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="DESCRIPTION" insert_db_data_type="string" insert_data_source="lms_course_item_table_obj.description" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="ORDER" insert_db_data_type="string" insert_data_source="lms_course_item_table_obj.order" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="VIDEO_URL" insert_db_data_type="string" insert_data_source="lms_course_item_table_obj.video_url" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="DURATION_IN_MIN" insert_db_data_type="string" insert_data_source="lms_course_item_table_obj.duration_in_min" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="HLS_1080_URL" insert_db_data_type="string" insert_data_source="lms_course_item_table_obj.hls_1080_url" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="HLS_720_URL" insert_db_data_type="string" insert_data_source="lms_course_item_table_obj.hls_720_url" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="HLS_480_URL" insert_db_data_type="string" insert_data_source="lms_course_item_table_obj.hls_480_url" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="IS_ENCRYPT" insert_db_data_type="boolean" insert_data_source="ReqBrokerReqObj.request_data.is_encrypt" insert_data_type="Boolean"></insert_data>
					<insert_data insert_data_name="DRM_TYPE" insert_db_data_type="string" insert_data_source="lms_course_item_table_obj.drm_type" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="IS_LOCKED" insert_db_data_type="boolean" insert_data_source="ReqBrokerReqObj.request_data.is_locked" insert_data_type="Boolean"></insert_data>
					<insert_data insert_data_name="UNLOCK_CONDITION" insert_db_data_type="string" insert_data_source="lms_course_item_table_obj.unlock_condition" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="IS_REQUIRED" insert_db_data_type="boolean" insert_data_source="ReqBrokerReqObj.request_data.is_required" insert_data_type="Boolean"></insert_data>
					<insert_data insert_data_name="ESTIMATE_TIME_MIN" insert_db_data_type="string" insert_data_source="lms_course_item_table_obj.estimate_time_min" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="PDF_URL" insert_db_data_type="string" insert_data_source="lms_course_item_table_obj.pdf_url" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="PPT_URL" insert_db_data_type="string" insert_data_source="lms_course_item_table_obj.ppt_url" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="HTML_CONTENT" insert_db_data_type="string" insert_data_source="lms_course_item_table_obj.html_content" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="REQ_COMPLETE_PERCENTAGE" insert_db_data_type="string" insert_data_source="lms_course_item_table_obj.req_complete_percentage" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="CREATED_BY" insert_db_data_type="string" insert_data_source="lms_course_item_table_obj.created_by" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="CREATED_DATE" insert_db_data_type="date" insert_data_source="lms_course_item_table_obj.created_date" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="CREATED_TIME" insert_db_data_type="string" insert_data_source="lms_course_item_table_obj.created_time" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="UPDATED_BY" insert_db_data_type="string" insert_data_source="lms_course_item_table_obj.updated_by" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="UPDATED_DATE" insert_db_data_type="date" insert_data_source="lms_course_item_table_obj.updated_date" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="UPDATED_TIME" insert_db_data_type="string" insert_data_source="lms_course_item_table_obj.updated_time" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="DELETED_BY" insert_db_data_type="string" insert_data_source="lms_course_item_table_obj.deleted_by" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="DELETED_DATE" insert_db_data_type="date" insert_data_source="lms_course_item_table_obj.deleted_date" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="DELETED_TIME" insert_db_data_type="string" insert_data_source="lms_course_item_table_obj.deleted_time" insert_data_type="string"></insert_data>
				</insert_data_list>
			</db_insert_info>
		</exec_function>
		<exec_function function_name="insert_lms_course_item_db_single_read" function_type="db_single_read">
			<db_single_read_info app_err_desc="lms course item table is not found" app_corrective_action="lms course item cannot be configured as the table does not exist">
				<table_name>LMS_COURSE_SECTION</table_name>
				<filter_info>
					<filter_data filter_name="COURSE_ID" filter_id="ReqBrokerReqObj.request_data.course_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="ID" filter_id="ReqBrokerReqObj.request_data.section_id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_course_section_table_record</store_obj_name>
			</db_single_read_info>
		</exec_function>
		<exec_function function_name="insert_lms_course_item_db_single_read" function_type="db_multi_read">
			<db_multi_read_info app_err_desc="lms course item table is not found" app_corrective_action="lms course item cannot be configured as the table does not exist">
				<table_name>LMS_COURSE_ITEM</table_name>
				<filter_info>
					<filter_data filter_name="COURSE_ID" filter_id="ReqBrokerReqObj.request_data.course_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="SECTION_ID" filter_id="ReqBrokerReqObj.request_data.section_id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_course_section_table_record.course_item</store_obj_name>
			</db_multi_read_info>
		</exec_function>
		<exec_function function_name="insert_lms_course_item_compose_response" function_type="compose_response">
			<response_info>
				<resp_obj obj_name="lms_course_section_table_record" obj_type="object" data_source="lms_course_section_table_record"></resp_obj>
			</response_info>
		</exec_function>
	</exec_group>
</sem_lexical_parser>
	`
	return lexicalParserStr
}
