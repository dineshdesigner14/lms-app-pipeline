package lexicalfile_lms_enquiry

func GetHandleLMSEnquiryCourseServiceXML() interface{} {
	lexicalParserStr := `
	<sem_lexical_parser microservice="insert_lms_enquiry_course">
	<exec_group group_name="insert_lms_enquiry_course">
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
		<exec_function function_name="lms_enquiry_course_validate_fld" function_type="validate_fld">
			<validate_info app_err_desc="pattern validation failed for input" app_corrective_action="check the inputs and provide proper values">
				<validate_fld fld_name="ReqBrokerReqObj.request_data.enquiry_id" fld_data_type="string" fld_min_len="1" fld_max_len="128" fld_type="pattern" fld_pattern="^(\d+)$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.course_id" fld_data_type="string" fld_min_len="1" fld_max_len="128" fld_type="pattern" fld_pattern="^(\d+)$"></validate_fld>
			</validate_info>
		</exec_function>
		<exec_function function_name="lms_user_table_db_single_read" function_type="db_single_read">
			<db_single_read_info app_err_desc="No user has been made for the portal" app_corrective_action="User has to be onboarded">
				<table_name>LMS_USER_DATA</table_name>
				<filter_info>
					<filter_data filter_name="EMAIL" filter_id="decode_token_obj.email" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_user_data_table_record</store_obj_name>
			</db_single_read_info>
		</exec_function>
		<exec_function function_name="lms_course_table_db_single_read" function_type="db_single_read">
			<db_single_read_info app_err_desc="No course has been made for the portal" app_corrective_action="Create and assign a course in the portal before proceeding">
				<table_name>LMS_COURSE</table_name>
				<filter_info>
					<filter_data filter_name="ID" filter_id="ReqBrokerReqObj.request_data.course_id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_course_table_record</store_obj_name>
			</db_single_read_info>
		</exec_function>
		<exec_function function_name="lms_enquiry_course_table_db_single_read" function_type="db_single_read">
			<db_single_read_info app_err_desc="No enquiry has been made for the student" app_corrective_action="Kindly complete the enquiry before proceeding">
				<table_name>LMS_ENQUIRY</table_name>
				<filter_info>
					<filter_data filter_name="ID" filter_id="ReqBrokerReqObj.request_data.enquiry_id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_enquiry_table_record</store_obj_name>
			</db_single_read_info>
		</exec_function>
		<exec_function function_name="lms_enquiry_course_table_db_multi_read" function_type="db_multi_read">
			<db_multi_read_info app_err_desc="Enquiry Course Table Doesn't Exists" app_corrective_action="Please Check with the Dev Team and Procees">
				<table_name>LMS_ENQUIRY_COURSE</table_name>
				<filter_info>
					<filter_data filter_name="ENQUIRY_ID" filter_id="ReqBrokerReqObj.request_data.enquiry_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="COURSE_ID" filter_id="ReqBrokerReqObj.request_data.course_id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_enquiry_course_table_list</store_obj_name>
			</db_multi_read_info>
		</exec_function>
		<exec_function function_name="lms_enquiry_courses_table_copy_object" function_type="copy_object" function_condition="len(lms_enquiry_course_table_list)==0">
			<copy_object_info object_name="lms_enquiry_courses_table_obj">
				<object_data key="course_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.course_id" data_type="string"></object_data>
				<object_data key="enquiry_id" data_source_type="reqbrokermap" data_source="lms_enquiry_table_record.id" data_type="string"></object_data>
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
		<exec_function function_name="lms_enquiry_courses_insert" function_type="db_insert" function_condition="len(lms_enquiry_course_table_list)==0">
			<db_insert_info>
				<table_name>LMS_ENQUIRY_COURSE</table_name>
				<insert_data_list>
					<insert_data insert_data_name="COURSE_ID" insert_db_data_type="string" insert_data_source="lms_enquiry_courses_table_obj.course_id" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="ENQUIRY_ID" insert_db_data_type="string" insert_data_source="lms_enquiry_courses_table_obj.enquiry_id" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="CREATED_BY" insert_db_data_type="string" insert_data_source="lms_enquiry_courses_table_obj.created_by" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="CREATED_DATE" insert_db_data_type="date" insert_data_source="lms_enquiry_courses_table_obj.created_date" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="CREATED_TIME" insert_db_data_type="string" insert_data_source="lms_enquiry_courses_table_obj.created_time" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="UPDATED_BY" insert_db_data_type="string" insert_data_source="lms_enquiry_courses_table_obj.updated_by" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="UPDATED_DATE" insert_db_data_type="date" insert_data_source="lms_enquiry_courses_table_obj.updated_date" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="UPDATED_TIME" insert_db_data_type="string" insert_data_source="lms_enquiry_courses_table_obj.updated_time" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="DELETED_BY" insert_db_data_type="string" insert_data_source="lms_enquiry_courses_table_obj.deleted_by" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="DELETED_DATE" insert_db_data_type="date" insert_data_source="lms_enquiry_courses_table_obj.deleted_date" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="DELETED_TIME" insert_db_data_type="string" insert_data_source="lms_enquiry_courses_table_obj.deleted_time" insert_data_type="string"></insert_data>
				</insert_data_list>
			</db_insert_info>
		</exec_function>
		<exec_function function_name="lms_enquiry_courses_table_copy_object" function_type="copy_object" function_condition="len(lms_enquiry_course_table_list)==1">
			<copy_object_info object_name="lms_enquiry_courses_table_obj">
				<object_data key="course_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.course_id" data_type="string"></object_data>
				<object_data key="enquiry_id" data_source_type="reqbrokermap" data_source="lms_enquiry_table_record.id" data_type="string"></object_data>
				<object_data key="updated_by" data_source_type="key" data_source="get_na" data_type="string"></object_data>
				<object_data key="updated_date" data_source_type="key" data_source="get_date" data_type="string"></object_data>
				<object_data key="updated_time" data_source_type="key" data_source="get_time" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
		<exec_function function_name="lms_enquiry_course_db_update" function_type="db_update" function_condition="len(lms_enquiry_course_table_list)==1">
			<db_update_info>
				<table_name>LMS_ENQUIRY_COURSE</table_name>
				<update_data_list>
					<update_data update_data_name="UPDATED_BY" update_db_data_type="string" update_data_source="lms_enquiry_courses_table_obj.updated_by" update_data_type="string"></update_data>
					<update_data update_data_name="UPDATED_DATE" update_db_data_type="date" update_data_source="lms_enquiry_courses_table_obj.updated_date" update_data_type="string"></update_data>
					<update_data update_data_name="UPDATED_TIME" update_db_data_type="string" update_data_source="lms_enquiry_courses_table_obj.updated_time" update_data_type="string"></update_data>
				</update_data_list>
				<update_filter_list>
					<update_filter update_filter_name="ENQUIRY_ID" update_filter_db_data_type="string" update_filter_data_source="lms_enquiry_courses_table_obj.enquiry_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
					<update_filter update_filter_name="COURSE_ID" update_filter_db_data_type="string" update_filter_data_source="lms_enquiry_courses_table_obj.course_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
				</update_filter_list>
			</db_update_info>
		</exec_function>
		<exec_function function_name="edit_lms_enquiry_course_db_multi_read" function_type="db_multi_read" >
			<db_multi_read_info app_err_desc="Data Already Exists or No Rows in Result Set" app_corrective_action="check the inputs and provide proper values">
				<table_name>LMS_ENQUIRY_COURSE</table_name>
				<filter_info>
					<filter_data filter_name="ENQUIRY_ID" filter_id="ReqBrokerReqObj.request_data.enquiry_id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_enquiry_course_table_list</store_obj_name>
			</db_multi_read_info>
		</exec_function>
		<exec_function function_name="lms_enquiry_course_loop" function_type="start_loop" index_name="i" start_index="0" end_index="len(lms_enquiry_course_table_list)" />
		<exec_function function_name="insert_lms_course_progress_copy_object" function_type="copy_object">
			<copy_object_info object_name="insert_lms_course_progress_additional_obj">
				<object_data key="user_id" data_source_type="reqbrokermap" data_source="lms_user_data_table_record.id" data_type="string"></object_data>
				<object_data key="course_id"  data_source_type="reqbrokermap" data_source="lms_enquiry_course_table_list[i].course_id" data_type="string"></object_data>
				<object_data key="percentage" data_source_type="raw_value" data_source="0" data_type="string"></object_data>
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
		<exec_function function_name="insert_lms_course_progress_db_insert" function_type="db_insert">
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
		<exec_function function_name="insert_lms_course_section_db_single_read" function_type="db_multi_read">
			<db_multi_read_info app_err_desc="lms course  section data table is not found" app_corrective_action="course section data cannot be configured as the table does not exist">
				<table_name>LMS_COURSE_SECTION</table_name>
				<filter_info>
					<filter_data filter_name="COURSE_ID" filter_id="lms_enquiry_course_table_list[i].course_id" filter_data_type="string" filter_condition="AND"></filter_data>
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
				<object_data key="user_id" data_source_type="reqbrokermap" data_source="lms_user_data_table_record.id" data_type="string"></object_data>
				<object_data key="course_id"  data_source_type="reqbrokermap" data_source="lms_enquiry_course_table_list[i].course_id" data_type="string"></object_data>
				<object_data key="section_id"  data_source_type="reqbrokermap" data_source="lms_course_section_table_list[j].id" data_type="string"></object_data>
				<object_data key="is_enabled"  data_source_type="raw_value" data_source="false" data_type="Boolean"></object_data>
				<object_data key="percentage" data_source_type="raw_value" data_source="0" data_type="string"></object_data>
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
		<exec_function function_name="insert_lms_course_progress_copy_object" function_type="copy_object" function_condition="lms_course_section_table_list[j].order==1">
			<copy_object_info object_name="insert_lms_course_progress_additional_obj">
				<object_data key="user_id" data_source_type="reqbrokermap" data_source="lms_user_data_table_record.id" data_type="string"></object_data>
				<object_data key="course_id"  data_source_type="reqbrokermap" data_source="lms_enquiry_course_table_list[i].course_id" data_type="string"></object_data>
				<object_data key="section_id"  data_source_type="reqbrokermap" data_source="lms_course_section_table_list[j].id" data_type="string"></object_data>
				<object_data key="is_enabled"  data_source_type="raw_value" data_source="true" data_type="Boolean"></object_data>
				<object_data key="percentage" data_source_type="raw_value" data_source="0" data_type="string"></object_data>
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
					<filter_data filter_name="COURSE_ID" filter_id="lms_enquiry_course_table_list[i].course_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="SECTION_ID" filter_id="lms_course_section_table_list[j].id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_course_item_table_list</store_obj_name>
			</db_multi_read_info>
			<err_info>
				<err_desc err_code="Duplicate" app_err_desc="Duplicate records found for course and section provided in the input." app_corrective_action="Make sure each course item should be linked with one section" />
			</err_info>
		</exec_function>
		<exec_function function_name="lms_enquiry_course_item_loop" function_type="start_loop" index_name="k" start_index="0" end_index="len(lms_course_item_table_list)" />
		<exec_function function_name="insert_lms_course_progress_copy_object" function_type="copy_object" function_condition="lms_course_item_table_list[k].order!=1">
			<copy_object_info object_name="insert_lms_course_progress_additional_obj">
				<object_data key="user_id" data_source_type="reqbrokermap" data_source="lms_user_data_table_record.id" data_type="string"></object_data>
				<object_data key="course_id"  data_source_type="reqbrokermap" data_source="lms_enquiry_course_table_list[i].course_id" data_type="string"></object_data>
				<object_data key="section_id"  data_source_type="reqbrokermap" data_source="lms_course_section_table_list[j].id" data_type="string"></object_data>
				<object_data key="item_id"  data_source_type="reqbrokermap" data_source="lms_course_item_table_list[k].id" data_type="string"></object_data>
				<object_data key="is_enabled"  data_source_type="raw_value" data_source="false" data_type="Boolean"></object_data>
				<object_data key="percentage" data_source_type="raw_value" data_source="0" data_type="string"></object_data>
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
		<exec_function function_name="insert_lms_course_progress_copy_object" function_type="copy_object" function_condition="lms_course_item_table_list[k].order==1">
			<copy_object_info object_name="insert_lms_course_progress_additional_obj">
				<object_data key="user_id" data_source_type="reqbrokermap" data_source="lms_user_data_table_record.id" data_type="string"></object_data>
				<object_data key="course_id"  data_source_type="reqbrokermap" data_source="lms_enquiry_course_table_list[i].course_id" data_type="string"></object_data>
				<object_data key="section_id"  data_source_type="reqbrokermap" data_source="lms_course_section_table_list[j].id" data_type="string"></object_data>
				<object_data key="item_id"  data_source_type="reqbrokermap" data_source="lms_course_item_table_list[k].id" data_type="string"></object_data>
				<object_data key="is_enabled"  data_source_type="raw_value" data_source="true" data_type="Boolean"></object_data>
				<object_data key="percentage" data_source_type="raw_value" data_source="0" data_type="string"></object_data>
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
		<exec_function function_name="insert_lms_course_item_progress_db_insert" function_type="db_insert">
			<db_insert_info>
				<table_name>LMS_COURSE_ITEM_PROGRESS</table_name>
				<insert_data_list>
					<insert_data insert_data_name="USER_ID" insert_db_data_type="string" insert_data_source="insert_lms_course_progress_additional_obj.user_id" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="COURSE_ID" insert_db_data_type="string" insert_data_source="insert_lms_course_progress_additional_obj.course_id" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="SECTION_ID" insert_db_data_type="string" insert_data_source="insert_lms_course_progress_additional_obj.section_id" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="ITEM_ID" insert_db_data_type="string" insert_data_source="insert_lms_course_progress_additional_obj.item_id" insert_data_type="string"></insert_data>
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
		<exec_function function_name="lms_enquiry_course_item_loop" function_type="end_loop" index_name="k" />
		<exec_function function_name="lms_enquiry_course_section_loop" function_type="end_loop" index_name="j" />
		<exec_function function_name="lms_enquiry_course_loop" function_type="end_loop" index_name="i" />
		<exec_function function_name="lms_enquiry_courses_table_copy_object" function_type="copy_object" >
			<copy_object_info object_name="lms_enquiry_courses_table_obj">
				<object_data key="status" data_source_type="raw_value" data_source="The course has been updated to the enquiry." data_type="string" condition="len(lms_enquiry_course_table_list)==1"></object_data>
				<object_data key="status" data_source_type="raw_value" data_source="The course has been added to the enquiry." data_type="string" condition="len(lms_enquiry_course_table_list)==0"></object_data>
			</copy_object_info>
		</exec_function>
		<exec_function function_name="insert_lms_enquiry_compose_response" function_type="compose_response">
			<response_info>
				<resp_obj obj_name="lms_enquiry_courses_table_obj" obj_type="object" data_source="lms_enquiry_courses_table_obj"></resp_obj>
			</response_info>
		</exec_function>
	</exec_group>
</sem_lexical_parser>
	`
	return lexicalParserStr
}
