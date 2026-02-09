package lexicalfile_lms_question_bank

func GetEditLMSQuestionBankServiceXML() interface{} {
	lexicalParserStr := `
	<sem_lexical_parser microservice="edit_lms_question_bank">
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
	<exec_group group_name="edit_lms_question_bank" group_condition="ReqBrokerReqObj.header_data.channel=='web_admin_portal'">
		<exec_function function_name="edit_lms_question_bank_validate_fld" function_type="validate_fld">
			<validate_info app_err_desc="pattern validation failed for input" app_corrective_action="check the inputs and provide proper values">
				<validate_fld fld_name="ReqBrokerReqObj.request_data.course_id" fld_data_type="string" fld_min_len="1" fld_max_len="128" fld_type="pattern" fld_pattern="^[a-zA-Z0-9.,|_* -]+$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.section_id" fld_data_type="string" fld_min_len="1" fld_max_len="128" fld_type="pattern" fld_pattern="^[a-zA-Z0-9.,|_* -]+$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.item_id" fld_data_type="string" fld_min_len="1" fld_max_len="128" fld_type="pattern" fld_pattern="^[a-zA-Z0-9.,|_* -]+$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.question_id" fld_data_type="string" fld_min_len="1" fld_max_len="128" fld_type="pattern" fld_pattern="^[a-zA-Z0-9.,|_* -]+$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.title" fld_data_type="string" fld_min_len="1" fld_max_len="255" fld_type="pattern" fld_pattern=".*"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.description" fld_data_type="string" fld_min_len="1" fld_max_len="255" fld_type="pattern" fld_pattern=".*"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.question" fld_data_type="string" fld_min_len="0" fld_max_len="4000" fld_type="pattern" fld_pattern=".*"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.difficulty" fld_data_type="string" fld_min_len="0" fld_max_len="128" fld_type="pattern" fld_pattern="^[A-Za-z0-9 ._-]*$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.points" fld_data_type="string" fld_min_len="0" fld_max_len="500" fld_type="pattern" fld_pattern="^[a-zA-Z0-9.,|_* -]+$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.explanation" fld_data_type="string" fld_min_len="0" fld_max_len="4000" fld_type="pattern" fld_pattern=".*"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.time_limit_in_sec" fld_data_type="string" fld_min_len="0" fld_max_len="128" fld_type="pattern" fld_pattern="^[A-Za-z0-9 -]*$"></validate_fld>
			</validate_info>
		</exec_function>
		<exec_function function_name="edit_lms_question_bank_auth_db_single_read" function_type="db_single_read">
			<db_single_read_info app_err_desc="Data Already Exists or No Rows in Result Set" app_corrective_action="check the inputs and provide proper values">
				<table_name>LMS_QUESTION_BANK</table_name>
				<filter_info>
					<filter_data filter_name="COURSE_ID" filter_id="ReqBrokerReqObj.request_data.course_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="SECTION_ID" filter_id="ReqBrokerReqObj.request_data.section_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="ITEM_ID" filter_id="ReqBrokerReqObj.request_data.item_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="ID" filter_id="ReqBrokerReqObj.request_data.question_id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_question_bank_table_record</store_obj_name>
			</db_single_read_info>
		</exec_function>
		<exec_function function_name="edit_lms_question_bank_table_copy_object" function_type="copy_object">
			<copy_object_info object_name="lms_question_bank_table_obj">
				<object_data key="id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.question_id" data_type="string"></object_data>
				<object_data key="course_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.course_id" data_type="string"></object_data>
				<object_data key="section_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.section_id" data_type="string"></object_data>
				<object_data key="item_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.item_id" data_type="string"></object_data>
				<object_data key="title" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.title" data_type="string"></object_data>
				<object_data key="description" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.description" data_type="string"></object_data>
				<object_data key="question" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.question" data_type="string"></object_data>
				<object_data key="difficulty" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.difficulty" data_type="string"></object_data>
				<object_data key="points" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.points" data_type="string"></object_data>
				<object_data key="explanation" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.explanation" data_type="string"></object_data>
				<object_data key="time_limit_in_sec" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.time_limit_in_sec" data_type="string"></object_data>
				<object_data key="updated_by" data_source_type="key" data_source="get_na" data_type="string"></object_data>
				<object_data key="updated_date" data_source_type="key" data_source="get_date" data_type="string"></object_data>
				<object_data key="updated_time" data_source_type="key" data_source="get_time" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
		<exec_function function_name="lms_question_bank_db_update" function_type="db_update">
			<db_update_info>
				<table_name>LMS_QUESTION_BANK</table_name>
				<update_data_list>
					<update_data update_data_name="TITLE" update_db_data_type="string" update_data_source="lms_question_bank_table_obj.title" update_data_type="string"></update_data>
					<update_data update_data_name="DESCRIPTION" update_db_data_type="string" update_data_source="lms_question_bank_table_obj.description" update_data_type="string"></update_data>
					<update_data update_data_name="QUESTION" update_db_data_type="string" update_data_source="lms_question_bank_table_obj.question" update_data_type="string"></update_data>
					<update_data update_data_name="DIFFICULTY" update_db_data_type="string" update_data_source="lms_question_bank_table_obj.difficulty" update_data_type="string"></update_data>
					<update_data update_data_name="POINTS" update_db_data_type="string" update_data_source="lms_question_bank_table_obj.points" update_data_type="string"></update_data>
					<update_data update_data_name="EXPLANATION" update_db_data_type="string" update_data_source="lms_question_bank_table_obj.explanation" update_data_type="string"></update_data>
					<update_data update_data_name="TIME_LIMIT_IN_SEC" update_db_data_type="string" update_data_source="lms_question_bank_table_obj.time_limit_in_sec" update_data_type="string"></update_data>
					<update_data update_data_name="UPDATED_BY" update_db_data_type="string" update_data_source="lms_question_bank_table_obj.updated_by" update_data_type="string"></update_data>
					<update_data update_data_name="UPDATED_DATE" update_db_data_type="date" update_data_source="lms_question_bank_table_obj.updated_date" update_data_type="string"></update_data>
					<update_data update_data_name="UPDATED_TIME" update_db_data_type="string" update_data_source="lms_question_bank_table_obj.updated_time" update_data_type="string"></update_data>
				</update_data_list>
				<update_filter_list>
					<update_filter update_filter_name="COURSE_ID" update_filter_db_data_type="string" update_filter_data_source="ReqBrokerReqObj.request_data.course_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
					<update_filter update_filter_name="SECTION_ID" update_filter_db_data_type="string" update_filter_data_source="ReqBrokerReqObj.request_data.section_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
					<update_filter update_filter_name="ITEM_ID" update_filter_db_data_type="string" update_filter_data_source="ReqBrokerReqObj.request_data.item_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
					<update_filter update_filter_name="ID" update_filter_db_data_type="string" update_filter_data_source="ReqBrokerReqObj.request_data.question_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
				</update_filter_list>
			</db_update_info>
		</exec_function>
		<exec_function function_name="lms_question_options_loop" function_type="start_loop" index_name="i" start_index="0" end_index="len(ReqBrokerReqObj.request_data.options)" />

		<exec_function function_name="lms_question_options_validate_fld_edit_delete" function_type="validate_fld" function_condition="ReqBrokerReqObj.request_data.options[i].db_operation=='EDIT' || ReqBrokerReqObj.request_data.options[i].db_operation=='DELETE'">
			<validate_info app_err_desc="pattern validation failed for input" app_corrective_action="check the inputs and provide proper values">
				<validate_fld fld_name="ReqBrokerReqObj.request_data.options[i].id" fld_data_type="string" fld_min_len="1" fld_max_len="128" fld_type="pattern" fld_pattern="^[a-zA-Z0-9.,|_*=\-; ]+$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.options[i].option_text" fld_data_type="string" fld_min_len="1" fld_max_len="128" fld_type="pattern" fld_pattern=".*"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.options[i].order" fld_data_type="string" fld_min_len="1" fld_max_len="128" fld_type="pattern" fld_pattern="^[a-zA-Z0-9.,|_* ]+$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.options[i].is_correct" fld_data_type="Boolean" fld_min_len="1" fld_max_len="128" fld_type="pattern" fld_pattern="^[a-zA-Z0-9.,|_* ]+$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.options[i].db_operation" fld_data_type="string" fld_min_len="1" fld_max_len="128" fld_type="pattern" fld_pattern="^[INSERT|EDIT|DELETE]+$"></validate_fld>
			</validate_info>
		</exec_function>


		<exec_function function_name="lms_question_options_validate_fld" function_type="validate_fld" function_condition="ReqBrokerReqObj.request_data.options[i].db_operation=='INSERT'">
			<validate_info app_err_desc="pattern validation failed for input" app_corrective_action="check the inputs and provide proper values">
				<validate_fld fld_name="ReqBrokerReqObj.request_data.options[i].option_text" fld_data_type="string" fld_min_len="1" fld_max_len="128" fld_type="pattern" fld_pattern="^[a-zA-Z0-9.,|_*=\-; ]+$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.options[i].order" fld_data_type="string" fld_min_len="1" fld_max_len="128" fld_type="pattern" fld_pattern="^[a-zA-Z0-9.,|_* ]+$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.options[i].is_correct" fld_data_type="Boolean" fld_min_len="1" fld_max_len="128" fld_type="pattern" fld_pattern="^[a-zA-Z0-9.,|_* ]+$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.options[i].db_operation" fld_data_type="string" fld_min_len="1" fld_max_len="128" fld_type="pattern" fld_pattern="^[INSERT|EDIT|DELETE]+$"></validate_fld>
			</validate_info>
		</exec_function>

<exec_function function_name="insert_lms_course_db_single_read" function_type="db_multi_read" function_condition="ReqBrokerReqObj.request_data.options[i].db_operation=='INSERT'">
			<db_multi_read_info app_err_desc="lms course data table is not found" app_corrective_action="lms course data cannot be configured as the table does not exist">
				<table_name>LMS_QUESTION_OPTIONS</table_name>
				<filter_info>
					<filter_data filter_name="COURSE_ID" filter_id="ReqBrokerReqObj.request_data.course_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="SECTION_ID" filter_id="ReqBrokerReqObj.request_data.section_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="ITEM_ID" filter_id="ReqBrokerReqObj.request_data.item_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="QUESTION_ID" filter_id="lms_question_bank_table_record.id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_question_options_table_list</store_obj_name>
			</db_multi_read_info>
		</exec_function>

		<exec_function function_name="lms_question_options_loop" function_type="start_loop" index_name="j" start_index="0" end_index="len(lms_question_options_table_list)" function_condition="ReqBrokerReqObj.request_data.options[i].db_operation=='INSERT'"/>

		<exec_function function_name="lms_question_options_auth_db_single_read" function_type="db_single_read" function_condition="ReqBrokerReqObj.request_data.options[i].db_operation=='INSERT'">
			<db_single_read_info app_err_desc="lms question options table error" app_corrective_action="lms question options data table has not been configured properly">
				<table_name>LMS_QUESTION_OPTIONS</table_name>
				<filter_info>
					<filter_data filter_name="COURSE_ID" filter_id="ReqBrokerReqObj.request_data.course_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="SECTION_ID" filter_id="ReqBrokerReqObj.request_data.section_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="ITEM_ID" filter_id="ReqBrokerReqObj.request_data.item_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="QUESTION_ID" filter_id="lms_question_options_table_list[j].question_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="ID" filter_id="lms_question_options_table_list[j].id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="ORDER" filter_id="ReqBrokerReqObj.request_data.options[i].order" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<result_info app_err_desc="The Option Order which is already exists." app_corrective_action="Please configure new option order and proceed">
						<result_success>NoRows</result_success>
					</result_info>			
			</db_single_read_info>
		</exec_function>
		<exec_function function_name="lms_question_options_loop" function_type="end_loop" index_name="j" function_condition="ReqBrokerReqObj.request_data.options[i].db_operation=='INSERT'"/>

		<exec_function function_name="insert_lms_question_option_table_copy_object" function_type="copy_object" function_condition="ReqBrokerReqObj.request_data.options[i].db_operation=='INSERT'">
			<copy_object_info object_name="lms_question_option_table_obj">
				<object_data key="course_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.course_id" data_type="string"></object_data>
				<object_data key="section_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.section_id" data_type="string"></object_data>
				<object_data key="item_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.item_id" data_type="string"></object_data>
				<object_data key="question_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.question_id" data_type="string"></object_data>
				<object_data key="option_text" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.options[i].option_text" data_type="string"></object_data>
				<object_data key="order" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.options[i].order" data_type="string"></object_data>
				<object_data key="is_correct" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.options[i].is_correct" data_type="Boolean"></object_data>
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
		
		<exec_function function_name="lms_question_option_table_insert" function_type="db_insert" function_condition="ReqBrokerReqObj.request_data.options[i].db_operation=='INSERT'">
			<db_insert_info>
				<table_name>LMS_QUESTION_OPTIONS</table_name>
				<insert_data_list>
					<insert_data insert_data_name="COURSE_ID" insert_db_data_type="string" insert_data_source="lms_question_option_table_obj.course_id" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="SECTION_ID" insert_db_data_type="string" insert_data_source="lms_question_option_table_obj.section_id" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="ITEM_ID" insert_db_data_type="string" insert_data_source="lms_question_option_table_obj.item_id" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="QUESTION_ID" insert_db_data_type="string" insert_data_source="lms_question_option_table_obj.question_id" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="OPTION_TEXT" insert_db_data_type="string" insert_data_source="lms_question_option_table_obj.option_text" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="ORDER" insert_db_data_type="string" insert_data_source="lms_question_option_table_obj.order" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="IS_CORRECT" insert_db_data_type="boolean" insert_data_source="lms_question_option_table_obj.is_correct" insert_data_type="Boolean"></insert_data>
					<insert_data insert_data_name="CREATED_BY" insert_db_data_type="string" insert_data_source="lms_question_option_table_obj.created_by" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="CREATED_DATE" insert_db_data_type="date" insert_data_source="lms_question_option_table_obj.created_date" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="CREATED_TIME" insert_db_data_type="string" insert_data_source="lms_question_option_table_obj.created_time" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="UPDATED_BY" insert_db_data_type="string" insert_data_source="lms_question_option_table_obj.updated_by" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="UPDATED_DATE" insert_db_data_type="date" insert_data_source="lms_question_option_table_obj.updated_date" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="UPDATED_TIME" insert_db_data_type="string" insert_data_source="lms_question_option_table_obj.updated_time" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="DELETED_BY" insert_db_data_type="string" insert_data_source="lms_question_option_table_obj.deleted_by" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="DELETED_DATE" insert_db_data_type="date" insert_data_source="lms_question_option_table_obj.deleted_date" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="DELETED_TIME" insert_db_data_type="string" insert_data_source="lms_question_option_table_obj.deleted_time" insert_data_type="string"></insert_data>
				</insert_data_list>
			</db_insert_info>
		</exec_function>
		<exec_function function_name="lms_question_options_auth_db_single_read" function_type="db_single_read" function_condition="ReqBrokerReqObj.request_data.options[i].db_operation=='INSERT'">
			<db_single_read_info app_err_desc="lms question options table error" app_corrective_action="lms question options data table has not been configured properly">
				<table_name>LMS_QUESTION_OPTIONS</table_name>
				<filter_info>
					<filter_data filter_name="COURSE_ID" filter_id="ReqBrokerReqObj.request_data.course_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="SECTION_ID" filter_id="ReqBrokerReqObj.request_data.section_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="ITEM_ID" filter_id="ReqBrokerReqObj.request_data.item_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="QUESTION_ID" filter_id="lms_question_bank_table_record.id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="IS_CORRECT" filter_source_type="raw_value" filter_id="true" filter_data_type="Boolean" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_question_options_table_records</store_obj_name>
			</db_single_read_info>
		</exec_function>
		<exec_function function_name="edit_lms_question_options_db_single_read" function_type="db_single_read" function_condition="ReqBrokerReqObj.request_data.options[i].db_operation=='EDIT'">
			<db_single_read_info app_err_desc="lms question options table error" app_corrective_action="lms question options data table has not been configured properly">
				<table_name>LMS_QUESTION_OPTIONS</table_name>
				<filter_info>
					<filter_data filter_name="COURSE_ID" filter_id="ReqBrokerReqObj.request_data.course_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="SECTION_ID" filter_id="ReqBrokerReqObj.request_data.section_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="QUESTION_ID" filter_id="ReqBrokerReqObj.request_data.question_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="ID" filter_id="ReqBrokerReqObj.request_data.options[i].id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="ORDER" filter_id="ReqBrokerReqObj.request_data.options[i].order" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_question_options_table_record</store_obj_name>
			</db_single_read_info>
		</exec_function>
		<exec_function function_name="edit_lms_question_bank_table_copy_object" function_type="copy_object" function_condition="ReqBrokerReqObj.request_data.options[i].db_operation=='EDIT'">
			<copy_object_info object_name="lms_question_option_table_obj">
				<object_data key="id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.options[i].id" data_type="string"></object_data>
				<object_data key="course_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.course_id" data_type="string"></object_data>
				<object_data key="section_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.section_id" data_type="string"></object_data>
				<object_data key="question_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.question_id" data_type="string"></object_data>
				<object_data key="option_text" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.options[i].option_text" data_type="string"></object_data>
				<object_data key="order" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.options[i].order" data_type="string"></object_data>
				<object_data key="is_correct" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.options[i].is_correct" data_type="Boolean"></object_data>
				<object_data key="updated_by" data_source_type="key" data_source="get_na" data_type="string"></object_data>
				<object_data key="updated_date" data_source_type="key" data_source="get_date" data_type="string"></object_data>
				<object_data key="updated_time" data_source_type="key" data_source="get_time" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
		<exec_function function_name="BuildHasCorrectOptionFlag" function_type="call_method"/>
		<exec_function function_name="edit_lms_question_options_db_update" function_type="db_update" function_condition="ReqBrokerReqObj.request_data.options[i].db_operation=='EDIT'">
			<db_update_info>
				<table_name>LMS_QUESTION_OPTIONS</table_name>
				<update_data_list>
					<update_data update_data_name="OPTION_TEXT" update_db_data_type="string" update_data_source="lms_question_option_table_obj.option_text" update_data_type="string"></update_data>
					<update_data update_data_name="ORDER" update_db_data_type="string" update_data_source="lms_question_option_table_obj.order" update_data_type="string"></update_data>
					<update_data update_data_name="IS_CORRECT" update_db_data_type="boolean" update_data_source="ReqBrokerReqObj.request_data.options[i].is_correct" update_data_type="Boolean"></update_data>
					<update_data update_data_name="UPDATED_BY" update_db_data_type="string" update_data_source="lms_question_option_table_obj.updated_by" update_data_type="string"></update_data>
					<update_data update_data_name="UPDATED_DATE" update_db_data_type="date" update_data_source="lms_question_option_table_obj.updated_date" update_data_type="string"></update_data>
					<update_data update_data_name="UPDATED_TIME" update_db_data_type="string" update_data_source="lms_question_option_table_obj.updated_time" update_data_type="string"></update_data>
				</update_data_list>
				<update_filter_list>
					<update_filter update_filter_name="COURSE_ID" update_filter_db_data_type="string" update_filter_data_source="lms_question_option_table_obj.course_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
					<update_filter update_filter_name="SECTION_ID" update_filter_db_data_type="string" update_filter_data_source="lms_question_option_table_obj.section_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
					<update_filter update_filter_name="QUESTION_ID" update_filter_db_data_type="string" update_filter_data_source="lms_question_option_table_obj.question_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
					<update_filter update_filter_name="ID" update_filter_db_data_type="string" update_filter_data_source="lms_question_option_table_obj.id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
				</update_filter_list>
			</db_update_info>
		</exec_function>
		<exec_function function_name="edit_lms_question_options_db_single_read" function_type="db_single_read" function_condition="ReqBrokerReqObj.request_data.options[i].db_operation=='DELETE'">
			<db_single_read_info app_err_desc="lms question options table error" app_corrective_action="lms question options data table has not been configured properly">
				<table_name>LMS_QUESTION_OPTIONS</table_name>
				<filter_info>
					<filter_data filter_name="COURSE_ID" filter_id="ReqBrokerReqObj.request_data.course_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="SECTION_ID" filter_id="ReqBrokerReqObj.request_data.section_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="QUESTION_ID" filter_id="ReqBrokerReqObj.request_data.question_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="ID" filter_id="ReqBrokerReqObj.request_data.options[i].id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="ORDER" filter_id="ReqBrokerReqObj.request_data.options[i].order" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_question_options_table_record</store_obj_name>
			</db_single_read_info>
		</exec_function>
		<exec_function function_name="insert_lms_question_option_table_copy_object" function_type="copy_object" function_condition="ReqBrokerReqObj.request_data.options[i].db_operation=='DELETE'">
			<copy_object_info object_name="lms_question_option_table_obj">
				<object_data key="question_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.question_id" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
		<exec_function function_name="delete_lms_question_options_table_delete" function_type="db_delete"  function_condition="ReqBrokerReqObj.request_data.options[i].db_operation=='DELETE'">
			<db_delete_info>
				<table_name>LMS_QUESTION_OPTIONS</table_name>
				<delete_filter_list>
					<delete_filter delete_filter_name="COURSE_ID" delete_filter_db_data_type="string" delete_filter_data_source="ReqBrokerReqObj.request_data.course_id" delete_filter_data_type="string" delete_filter_condition="AND"></delete_filter>
					<delete_filter delete_filter_name="SECTION_ID" delete_filter_db_data_type="string" delete_filter_data_source="ReqBrokerReqObj.request_data.section_id" delete_filter_data_type="string" delete_filter_condition="AND"></delete_filter>
					<delete_filter delete_filter_name="QUESTION_ID" delete_filter_db_data_type="string" delete_filter_data_source="ReqBrokerReqObj.request_data.question_id" delete_filter_data_type="string" delete_filter_condition="AND"></delete_filter>
					<delete_filter delete_filter_name="ID" delete_filter_db_data_type="string" delete_filter_data_source="ReqBrokerReqObj.request_data.options[i].id" delete_filter_data_type="string" delete_filter_condition="AND"></delete_filter>
				</delete_filter_list>
			</db_delete_info>
		</exec_function>
		<exec_function function_name="lms_question_tags_loop" function_type="start_loop" index_name="j" start_index="0" end_index="len(ReqBrokerReqObj.request_data.tags)" />
		<exec_function function_name="lms_question_tags_validate_fld" function_type="validate_fld">
			<validate_info app_err_desc="pattern validation failed for input" app_corrective_action="check the inputs and provide proper values">
				<validate_fld fld_name="ReqBrokerReqObj.request_data.tags[j].tags" fld_data_type="string" fld_min_len="1" fld_max_len="128" fld_type="pattern" fld_pattern="^[a-zA-Z0-9.,|_* ]+$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.tags[j].db_operation" fld_data_type="string" fld_min_len="1" fld_max_len="128" fld_type="pattern" fld_pattern="^[INSERT|EDIT|DELETE]+$"></validate_fld>
			</validate_info>
		</exec_function>


		<exec_function function_name="edit_lms_question_bank_auth_db_single_read" function_type="db_single_read">
			<db_single_read_info app_err_desc="Data Already Exists or No Rows in Result Set" app_corrective_action="check the inputs and provide proper values">
				<table_name>LMS_QUESTION_BANK</table_name>
				<filter_info>
					<filter_data filter_name="COURSE_ID" filter_id="ReqBrokerReqObj.request_data.course_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="SECTION_ID" filter_id="ReqBrokerReqObj.request_data.section_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="ITEM_ID" filter_id="ReqBrokerReqObj.request_data.item_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="ID" filter_id="ReqBrokerReqObj.request_data.question_id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_question_bank_table_record</store_obj_name>
			</db_single_read_info>
		</exec_function>

		<exec_function function_name="edit_lms_question_bank_auth_db_single_read" function_type="db_single_read" function_condition="ReqBrokerReqObj.request_data.tags[j].db_operation=='INSERT'">
			<db_single_read_info app_err_desc="Data Already Exists or No Rows in Result Set" app_corrective_action="check the inputs and provide proper values">
				<table_name>LMS_QUESTION_OPTIONS</table_name>
				<filter_info>
					<filter_data filter_name="COURSE_ID" filter_id="ReqBrokerReqObj.request_data.course_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="SECTION_ID" filter_id="ReqBrokerReqObj.request_data.section_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="ITEM_ID" filter_id="ReqBrokerReqObj.request_data.item_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="QUESTION_ID" filter_id="lms_question_bank_table_record.id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="OPTION_TEXT" filter_id="ReqBrokerReqObj.request_data.options[i].option_text" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_question_options_table_record</store_obj_name>
			</db_single_read_info>
		</exec_function>

		<exec_function function_name="insert_lms_question_tags_table_copy_object" function_type="copy_object" function_condition="ReqBrokerReqObj.request_data.tags[j].db_operation=='INSERT'">
			<copy_object_info object_name="lms_question_tags_table_obj">
				<object_data key="course_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.course_id" data_type="string"></object_data>
				<object_data key="section_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.section_id" data_type="string"></object_data>
				<object_data key="item_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.item_id" data_type="string"></object_data>
				<object_data key="question_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.question_id" data_type="string"></object_data>
				<object_data key="option_id" data_source_type="reqbrokermap" data_source="lms_question_options_table_record.id" data_type="string"></object_data>
				<object_data key="tags" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.tags[j].tags" data_type="string"></object_data>
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
		<exec_function function_name="lms_question_tags_table_insert" function_type="db_insert" function_condition="ReqBrokerReqObj.request_data.tags[j].db_operation=='INSERT'">
			<db_insert_info>
				<table_name>LMS_QUESTION_TAGS</table_name>
				<insert_data_list>
					<insert_data insert_data_name="COURSE_ID" insert_db_data_type="string" insert_data_source="lms_question_option_table_obj.course_id" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="SECTION_ID" insert_db_data_type="string" insert_data_source="lms_question_option_table_obj.section_id" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="ITEM_ID" insert_db_data_type="string" insert_data_source="lms_question_option_table_obj.item_id" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="QUESTION_ID" insert_db_data_type="string" insert_data_source="lms_question_option_table_obj.question_id" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="OPTION_ID" insert_db_data_type="string" insert_data_source="lms_question_tags_table_obj.option_id" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="TAGS" insert_db_data_type="string" insert_data_source="lms_question_tags_table_obj.tags" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="CREATED_BY" insert_db_data_type="string" insert_data_source="lms_question_tags_table_obj.created_by" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="CREATED_DATE" insert_db_data_type="date" insert_data_source="lms_question_tags_table_obj.created_date" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="CREATED_TIME" insert_db_data_type="string" insert_data_source="lms_question_tags_table_obj.created_time" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="UPDATED_BY" insert_db_data_type="string" insert_data_source="lms_question_tags_table_obj.updated_by" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="UPDATED_DATE" insert_db_data_type="date" insert_data_source="lms_question_tags_table_obj.updated_date" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="UPDATED_TIME" insert_db_data_type="string" insert_data_source="lms_question_tags_table_obj.updated_time" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="DELETED_BY" insert_db_data_type="string" insert_data_source="lms_question_tags_table_obj.deleted_by" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="DELETED_DATE" insert_db_data_type="date" insert_data_source="lms_question_tags_table_obj.deleted_date" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="DELETED_TIME" insert_db_data_type="string" insert_data_source="lms_question_tags_table_obj.deleted_time" insert_data_type="string"></insert_data>
				</insert_data_list>
			</db_insert_info>
		</exec_function>
		<exec_function function_name="edit_lms_question_options_db_single_read" function_type="db_single_read" function_condition="ReqBrokerReqObj.request_data.tags[j].db_operation=='EDIT'">
			<db_single_read_info app_err_desc="Data Already Exists or No Rows in Result Set" app_corrective_action="check the inputs and provide proper values">
				<table_name>LMS_QUESTION_TAGS</table_name>
				<filter_info>
					<filter_data filter_name="COURSE_ID" filter_id="ReqBrokerReqObj.request_data.course_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="SECTION_ID" filter_id="ReqBrokerReqObj.request_data.section_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="ITEM_ID" filter_id="ReqBrokerReqObj.request_data.item_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="QUESTION_ID" filter_id="lms_question_option_table_obj.question_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="OPTION_ID" filter_id="lms_question_option_table_obj.id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="ID" filter_id="ReqBrokerReqObj.request_data.tags[j].id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_question_options_table_record</store_obj_name>
			</db_single_read_info>
		</exec_function>
		<exec_function function_name="edit_lms_question_tags_table_copy_object" function_type="copy_object" function_condition="ReqBrokerReqObj.request_data.tags[j].db_operation=='EDIT'">
			<copy_object_info object_name="lms_question_tags_table_obj">
				<object_data key="id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.tags[j].id" data_type="string"></object_data>
				<object_data key="course_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.course_id" data_type="string"></object_data>
				<object_data key="section_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.section_id" data_type="string"></object_data>
				<object_data key="item_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.item_id" data_type="string"></object_data>
				<object_data key="question_id" data_source_type="reqbrokermap" data_source="lms_question_option_table_obj.question_id" data_type="string"></object_data>
				<object_data key="option_id" data_source_type="reqbrokermap" data_source="lms_question_options_table_record.id" data_type="string"></object_data>
				<object_data key="tags" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.tags[j].tags" data_type="string"></object_data>
				<object_data key="updated_by" data_source_type="key" data_source="get_na" data_type="string"></object_data>
				<object_data key="updated_date" data_source_type="key" data_source="get_date" data_type="string"></object_data>
				<object_data key="updated_time" data_source_type="key" data_source="get_time" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
		<exec_function function_name="edit_lms_question_tags_table_copy_object" function_type="copy_object" function_condition="ReqBrokerReqObj.request_data.tags[j].db_operation=='DELETE'">
			<copy_object_info object_name="lms_question_tags_table_obj">
				<object_data key="id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.tags[j].id" data_type="string"></object_data>
				<object_data key="course_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.course_id" data_type="string"></object_data>
				<object_data key="section_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.section_id" data_type="string"></object_data>
				<object_data key="item_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.item_id" data_type="string"></object_data>
				<object_data key="question_id" data_source_type="reqbrokermap" data_source="lms_question_option_table_obj.question_id" data_type="string"></object_data>
				<object_data key="option_id" data_source_type="reqbrokermap" data_source="lms_question_options_table_record.id" data_type="string"></object_data>
        		<object_data key="tags" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.tags[j].tags" data_type="string"></object_data>
				<object_data key="updated_by" data_source_type="key" data_source="get_na" data_type="string"></object_data>
				<object_data key="updated_date" data_source_type="key" data_source="get_date" data_type="string"></object_data>
				<object_data key="updated_time" data_source_type="key" data_source="get_time" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
		<exec_function function_name="edit_lms_question_tags_db_update" function_type="db_update" function_condition="ReqBrokerReqObj.request_data.tags[j].db_operation=='EDIT'">
			<db_update_info>
				<table_name>LMS_QUESTION_TAGS</table_name>
				<update_data_list>
					<update_data update_data_name="TAGS" update_db_data_type="string" update_data_source="lms_question_tags_table_obj.tags" update_data_type="string"></update_data>
					<update_data update_data_name="UPDATED_BY" update_db_data_type="string" update_data_source="lms_question_tags_table_obj.updated_by" update_data_type="string"></update_data>
					<update_data update_data_name="UPDATED_DATE" update_db_data_type="date" update_data_source="lms_question_tags_table_obj.updated_date" update_data_type="string"></update_data>
					<update_data update_data_name="UPDATED_TIME" update_db_data_type="string" update_data_source="lms_question_tags_table_obj.updated_time" update_data_type="string"></update_data>
				</update_data_list>
				<update_filter_list>
					<update_filter update_filter_name="COURSE_ID" update_filter_db_data_type="string" update_filter_data_source="lms_question_tags_table_obj.course_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
					<update_filter update_filter_name="SECTION_ID" update_filter_db_data_type="string" update_filter_data_source="lms_question_tags_table_obj.section_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
					<update_filter update_filter_name="ITEM_ID" update_filter_db_data_type="string" update_filter_data_source="lms_question_tags_table_obj.item_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
					<update_filter update_filter_name="QUESTION_ID" update_filter_db_data_type="string" update_filter_data_source="lms_question_tags_table_obj.question_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
					<update_filter update_filter_name="OPTION_ID" update_filter_db_data_type="string" update_filter_data_source="lms_question_tags_table_obj.option_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
					<update_filter update_filter_name="ID" update_filter_db_data_type="string" update_filter_data_source="lms_question_tags_table_obj.id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
				</update_filter_list>
			</db_update_info>
		</exec_function>
		<exec_function function_name="delete_lms_question_options_table_delete" function_type="db_delete"  function_condition="ReqBrokerReqObj.request_data.tags[j].db_operation=='DELETE'">
			<db_delete_info>
				<table_name>LMS_QUESTION_TAGS</table_name>
				<delete_filter_list>
					<delete_filter delete_filter_name="COURSE_ID" delete_filter_db_data_type="string" delete_filter_data_source="lms_question_tags_table_obj.course_id" delete_filter_data_type="string" delete_filter_condition="AND"></delete_filter>
					<delete_filter delete_filter_name="SECTION_ID" delete_filter_db_data_type="string" delete_filter_data_source="lms_question_tags_table_obj.section_id" delete_filter_data_type="string" delete_filter_condition="AND"></delete_filter>
					<delete_filter delete_filter_name="ITEM_ID" delete_filter_db_data_type="string" delete_filter_data_source="lms_question_tags_table_obj.item_id" delete_filter_data_type="string" delete_filter_condition="AND"></delete_filter>
					<delete_filter delete_filter_name="QUESTION_ID" delete_filter_db_data_type="string" delete_filter_data_source="lms_question_tags_table_obj.question_id" delete_filter_data_type="string" delete_filter_condition="AND"></delete_filter>
					<delete_filter delete_filter_name="OPTION_ID" delete_filter_db_data_type="string" delete_filter_data_source="lms_question_tags_table_obj.option_id" delete_filter_data_type="string" delete_filter_condition="AND"></delete_filter>
					<delete_filter delete_filter_name="ID" delete_filter_db_data_type="string" delete_filter_data_source="lms_question_tags_table_obj.id" delete_filter_data_type="string" delete_filter_condition="AND"></delete_filter>
				</delete_filter_list>
			</db_delete_info>
		</exec_function>
		<exec_function function_name="lms_question_tags_loop" function_type="end_loop" index_name="j" />
		<exec_function function_name="lms_question_options_loop" function_type="end_loop" index_name="i" />
		 <exec_function function_name="role_validate_condition_has_correct_option" function_type="validate_condition">
			<validate_condition_info app_err_desc="There should be one correct option" app_corrective_action="Please give one correct option and proceed">
				<validate_condition_expression>has_correct_option == true</validate_condition_expression>
			</validate_condition_info>
		</exec_function>
		<exec_function function_name="edit_lms_question_bank_auth_db_single_read" function_type="db_single_read">
			<db_single_read_info app_err_desc="Question Bank configuration missing" app_corrective_action="Please configure the question bank and proceed">
				<table_name>LMS_QUESTION_BANK</table_name>
				<filter_info>
					<filter_data filter_name="COURSE_ID" filter_id="ReqBrokerReqObj.request_data.course_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="SECTION_ID" filter_id="ReqBrokerReqObj.request_data.section_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="ITEM_ID" filter_id="ReqBrokerReqObj.request_data.item_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="ID" filter_id="ReqBrokerReqObj.request_data.question_id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_question_bank_table_record</store_obj_name>
			</db_single_read_info>
		</exec_function>
		<exec_function function_name="edit_lms_question_bank_auth_db_single_read" function_type="db_multi_read">
			<db_multi_read_info app_err_desc="Question Options configuration missing" app_corrective_action="Please configure the question options and proceed">
				<table_name>LMS_QUESTION_OPTIONS</table_name>
				<filter_info>
					<filter_data filter_name="COURSE_ID" filter_id="ReqBrokerReqObj.request_data.course_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="SECTION_ID" filter_id="ReqBrokerReqObj.request_data.section_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="ITEM_ID" filter_id="ReqBrokerReqObj.request_data.item_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="QUESTION_ID" filter_id="ReqBrokerReqObj.request_data.question_id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_question_bank_table_record.lms_question_options_table_list</store_obj_name>
			</db_multi_read_info>
		</exec_function>
		<exec_function function_name="insert_lms_course_section_compose_response" function_type="compose_response">
			<response_info>
				<resp_obj obj_name="lms_question_bank_table_record" obj_type="object" data_source="lms_question_bank_table_record"></resp_obj>
			</response_info>
		</exec_function>

	</exec_group>

	
</sem_lexical_parser>
	`
	return lexicalParserStr
}
