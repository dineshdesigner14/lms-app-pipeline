package lexicalfile_lms_question_bank

func GetInsertLMSQuestionBankServiceXML() interface{} {
	lexicalParserStr := `
	<sem_lexical_parser microservice="insert_lms_course_question_bank">
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
	<exec_group group_name="insert_lms_course_question_bank" group_condition="ReqBrokerReqObj.header_data.channel=='web_admin_portal'">
		<exec_function function_name="insert_lms_course_question_bank_validate_fld" function_type="validate_fld">
			<validate_info app_err_desc="Validation failed for course section data" app_corrective_action="Please validate the data">
				<validate_fld fld_name="ReqBrokerReqObj.request_data.course_id" fld_data_type="string" fld_min_len="1" fld_max_len="255" fld_type="pattern" fld_pattern="^[0-9.]+$" app_err_desc="Invalid course id" app_corrective_action="Only numbers are allowed."></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.section_id" fld_data_type="string" fld_min_len="1" fld_max_len="255" fld_type="pattern" fld_pattern="^[0-9.]+$" app_err_desc="Invalid section id" app_corrective_action="Only numbers are allowed."></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.item_id" fld_data_type="string" fld_min_len="1" fld_max_len="255" fld_type="pattern" fld_pattern="^[0-9.]+$" app_err_desc="Invalid item id" app_corrective_action="Only numbers are allowed."></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.title" fld_data_type="string" fld_min_len="1" fld_max_len="255" fld_type="pattern" fld_pattern=".*" app_err_desc="Invalid title" app_corrective_action="Only letters, numbers and spaces allowed."></validate_fld>
                <validate_fld fld_name="ReqBrokerReqObj.request_data.description" fld_data_type="string" fld_min_len="1" fld_max_len="255" fld_type="pattern" fld_pattern=".*" app_err_desc="Invalid description" app_corrective_action="Only letters, numbers and spaces allowed."></validate_fld>
                <validate_fld fld_name="ReqBrokerReqObj.request_data.question" fld_data_type="string" fld_min_len="1" fld_max_len="4000" fld_type="pattern" fld_pattern=".*" app_err_desc="Invalid question" app_corrective_action="Only letters, numbers and spaces allowed."></validate_fld>
			</validate_info>
		</exec_function>

		<exec_function function_name="insert_lms_course_question_bank_db_single_read" function_type="db_single_read">
			<db_single_read_info app_err_desc="lms course table is not found" app_corrective_action="course data cannot be configured as the table does not exist">
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
		<exec_function function_name="insert_lms_course_section_db_single_read" function_type="db_single_read">
			<db_single_read_info app_err_desc="lms question bank table is not found" app_corrective_action="question bank data cannot be configured as the table does not exist">
				<table_name>LMS_QUESTION_BANK</table_name>
				<filter_info>
					<filter_data filter_name="COURSE_ID" filter_id="ReqBrokerReqObj.request_data.course_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="SECTION_ID" filter_id="ReqBrokerReqObj.request_data.section_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="ITEM_ID" filter_id="ReqBrokerReqObj.request_data.item_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="QUESTION" filter_id="ReqBrokerReqObj.request_data.question" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<result_info app_err_desc="The question:{{.ReqBrokerReqObj.request_data.question}} already exists." app_corrective_action="Please configure new question and proceed">
						<result_success>NoRows</result_success>
					</result_info>
				</db_single_read_info>
				<err_info>
				  <err_desc err_code="Duplicate" app_err_desc="Duplicate records found for the question:{{.ReqBrokerReqObj.request_data.question}} provided in the input." app_corrective_action="Make sure each question shows up only once." />   
				</err_info>
		</exec_function>
		<exec_function function_name="lms_question_bank_copy_object" function_type="copy_object">
			<copy_object_info object_name="lms_question_bank_table_obj">
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
		<exec_function function_name="insert_lms_question_bank_db_insert" function_type="db_insert">
			<db_insert_info>
				<table_name>LMS_QUESTION_BANK</table_name>
				<insert_data_list>
					<insert_data insert_data_name="COURSE_ID" insert_db_data_type="string" insert_data_source="lms_question_bank_table_obj.course_id" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="SECTION_ID" insert_db_data_type="string" insert_data_source="lms_question_bank_table_obj.section_id" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="ITEM_ID" insert_db_data_type="string" insert_data_source="lms_question_bank_table_obj.item_id" insert_data_type="string"></insert_data>
                    <insert_data insert_data_name="TITLE" insert_db_data_type="string" insert_data_source="lms_question_bank_table_obj.title" insert_data_type="string"></insert_data>
                    <insert_data insert_data_name="DESCRIPTION" insert_db_data_type="string" insert_data_source="lms_question_bank_table_obj.description" insert_data_type="string"></insert_data>
                    <insert_data insert_data_name="QUESTION" insert_db_data_type="string" insert_data_source="lms_question_bank_table_obj.question" insert_data_type="string"></insert_data>
                    <insert_data insert_data_name="DIFFICULTY" insert_db_data_type="string" insert_data_source="lms_question_bank_table_obj.difficulty" insert_data_type="string"></insert_data>
                    <insert_data insert_data_name="POINTS" insert_db_data_type="string" insert_data_source="lms_question_bank_table_obj.points" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="EXPLANATION" insert_db_data_type="string" insert_data_source="lms_question_bank_table_obj.explanation" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="TIME_LIMIT_IN_SEC" insert_db_data_type="string" insert_data_source="lms_question_bank_table_obj.time_limit_in_sec" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="CREATED_BY" insert_db_data_type="string" insert_data_source="lms_question_bank_table_obj.created_by" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="CREATED_DATE" insert_db_data_type="date" insert_data_source="lms_question_bank_table_obj.created_date" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="CREATED_TIME" insert_db_data_type="string" insert_data_source="lms_question_bank_table_obj.created_time" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="UPDATED_BY" insert_db_data_type="string" insert_data_source="lms_question_bank_table_obj.updated_by" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="UPDATED_DATE" insert_db_data_type="date" insert_data_source="lms_question_bank_table_obj.updated_date" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="UPDATED_TIME" insert_db_data_type="string" insert_data_source="lms_question_bank_table_obj.updated_time" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="DELETED_BY" insert_db_data_type="string" insert_data_source="lms_question_bank_table_obj.deleted_by" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="DELETED_DATE" insert_db_data_type="date" insert_data_source="lms_question_bank_table_obj.deleted_date" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="DELETED_TIME" insert_db_data_type="string" insert_data_source="lms_question_bank_table_obj.deleted_time" insert_data_type="string"></insert_data>
				</insert_data_list>
			</db_insert_info>
		</exec_function>
<exec_function function_name="insert_lms_question_bank_db_single_read" function_type="db_single_read">
			<db_single_read_info app_err_desc="lms question bank table is not found" app_corrective_action="question bank data cannot be configured as the table does not exist">
				<table_name>LMS_QUESTION_BANK</table_name>
				<filter_info>
					<filter_data filter_name="COURSE_ID" filter_id="ReqBrokerReqObj.request_data.course_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="SECTION_ID" filter_id="ReqBrokerReqObj.request_data.section_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="ITEM_ID" filter_id="ReqBrokerReqObj.request_data.item_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="QUESTION" filter_id="ReqBrokerReqObj.request_data.question" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_question_bank_table_record</store_obj_name>
				</db_single_read_info>
				<err_info>
				  <err_desc err_code="Duplicate" app_err_desc="Duplicate records found for the question:{{.ReqBrokerReqObj.request_data.question}} provided in the input." app_corrective_action="Make sure each question shows up only once." />   
				</err_info>
		</exec_function>
	<exec_function function_name="lms_question_options_loop" function_type="start_loop" index_name="i" start_index="0" end_index="len(ReqBrokerReqObj.request_data.options)" />
		<exec_function function_name="lms_enquiry_comments_validate_fld" function_type="validate_fld">
			<validate_info app_err_desc="pattern validation failed for input" app_corrective_action="check the inputs and provide proper values">
				<validate_fld fld_name="ReqBrokerReqObj.request_data.options[i].option_text" fld_data_type="string" fld_min_len="1" fld_max_len="500" fld_type="pattern" fld_pattern=".*"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.options[i].order" fld_data_type="string" fld_min_len="1" fld_max_len="500" fld_type="pattern" fld_pattern="^[1-4]+$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.options[i].is_correct" fld_data_type="boolean" fld_min_len="1" fld_max_len="500" fld_type="pattern" fld_pattern="^[a-zA-Z0-9.,|_* -]+$"></validate_fld>
			</validate_info>
		</exec_function>

		<exec_function function_name="insert_lms_course_db_single_read" function_type="db_multi_read">
			<db_multi_read_info app_err_desc="lms question options table is not found" app_corrective_action="lms course data cannot be configured as the table does not exist">
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

		<exec_function function_name="lms_question_options_loop" function_type="start_loop" index_name="j" start_index="0" end_index="len(lms_question_options_table_list)" />

		<exec_function function_name="lms_question_options_auth_db_single_read" function_type="db_single_read">
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
				<result_info app_err_desc="The option order already exists." app_corrective_action="Please configure new option order and proceed">
						<result_success>NoRows</result_success>
					</result_info>			
			</db_single_read_info>
		</exec_function>

		<exec_function function_name="lms_question_options_loop" function_type="end_loop" index_name="j" />

		<exec_function function_name="lms_question_options_table_copy_object" function_type="copy_object">
			<copy_object_info object_name="lms_question_options_table_obj">
				<object_data key="course_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.course_id" data_type="string"></object_data>
				<object_data key="section_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.section_id" data_type="string"></object_data>
				<object_data key="item_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.item_id" data_type="string"></object_data>
				<object_data key="question_id" data_source_type="reqbrokermap" data_source="lms_question_bank_table_record.id" data_type="string"></object_data>
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
		<exec_function function_name="lms_question_options_table_insert" function_type="db_insert">
			<db_insert_info app_err_desc="pattern validation failed for input" app_corrective_action="check the inputs and provide proper values">
				<table_name>LMS_QUESTION_OPTIONS</table_name>
				<insert_data_list>
					<insert_data insert_data_name="COURSE_ID" insert_db_data_type="string" insert_data_source="lms_question_options_table_obj.course_id" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="SECTION_ID" insert_db_data_type="string" insert_data_source="lms_question_options_table_obj.section_id" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="ITEM_ID" insert_db_data_type="string" insert_data_source="lms_question_options_table_obj.item_id" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="QUESTION_ID" insert_db_data_type="string" insert_data_source="lms_question_options_table_obj.question_id" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="OPTION_TEXT" insert_db_data_type="string" insert_data_source="lms_question_options_table_obj.option_text" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="ORDER" insert_db_data_type="string" insert_data_source="lms_question_options_table_obj.order" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="IS_CORRECT" insert_db_data_type="boolean" insert_data_source="ReqBrokerReqObj.request_data.options[i].is_correct" insert_data_type="Boolean"></insert_data>
					<insert_data insert_data_name="CREATED_BY" insert_db_data_type="string" insert_data_source="lms_question_options_table_obj.created_by" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="CREATED_DATE" insert_db_data_type="date" insert_data_source="lms_question_options_table_obj.created_date" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="CREATED_TIME" insert_db_data_type="string" insert_data_source="lms_question_options_table_obj.created_time" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="UPDATED_BY" insert_db_data_type="string" insert_data_source="lms_question_options_table_obj.updated_by" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="UPDATED_DATE" insert_db_data_type="date" insert_data_source="lms_question_options_table_obj.updated_date" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="UPDATED_TIME" insert_db_data_type="string" insert_data_source="lms_question_options_table_obj.updated_time" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="DELETED_BY" insert_db_data_type="string" insert_data_source="lms_question_options_table_obj.deleted_by" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="DELETED_DATE" insert_db_data_type="date" insert_data_source="lms_question_options_table_obj.deleted_date" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="DELETED_TIME" insert_db_data_type="string" insert_data_source="lms_question_options_table_obj.deleted_time" insert_data_type="string"></insert_data>
				</insert_data_list>
			</db_insert_info>
		</exec_function>
		<exec_function function_name="lms_question_options_loop" function_type="end_loop" index_name="i" />

		<exec_function function_name="lms_question_options_auth_db_single_read" function_type="db_single_read">
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

		<exec_function function_name="insert_lms_question_bank_db_multi_read" function_type="db_multi_read">
			<db_multi_read_info app_err_desc="lms question bank table is not found" app_corrective_action="question bank data cannot be configured as the table does not exist">
				<table_name>LMS_QUESTION_OPTIONS</table_name>
				<filter_info>
					<filter_data filter_name="COURSE_ID" filter_id="ReqBrokerReqObj.request_data.course_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="SECTION_ID" filter_id="ReqBrokerReqObj.request_data.section_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="ITEM_ID" filter_id="ReqBrokerReqObj.request_data.item_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="QUESTION_ID" filter_id="lms_question_bank_table_record.id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_question_options_table_list</store_obj_name>
				</db_multi_read_info>
				<err_info>
				  <err_desc err_code="Duplicate" app_err_desc="Duplicate records found for the question:{{.ReqBrokerReqObj.request_data.question}} provided in the input." app_corrective_action="Make sure each question shows up only once." />   
				</err_info>
		</exec_function>

		<exec_function function_name="lms_question_tags_loop1" function_type="start_loop" index_name="i" start_index="0" end_index="len(lms_question_options_table_list)" />


		<exec_function function_name="insert_lms_question_bank_db_multi_read" function_type="db_single_read">
			<db_single_read_info app_err_desc="lms question bank table is not found" app_corrective_action="question bank data cannot be configured as the table does not exist">
				<table_name>LMS_QUESTION_OPTIONS</table_name>
				<filter_info>
					<filter_data filter_name="COURSE_ID" filter_id="ReqBrokerReqObj.request_data.course_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="SECTION_ID" filter_id="ReqBrokerReqObj.request_data.section_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="QUESTION_ID" filter_id="lms_question_bank_table_record.id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="ID" filter_id="lms_question_options_table_list[i].id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_question_options_table_record</store_obj_name>
				</db_single_read_info>
				<err_info>
				  <err_desc err_code="Duplicate" app_err_desc="Duplicate records found for the question:{{.ReqBrokerReqObj.request_data.question}} provided in the input." app_corrective_action="Make sure each question shows up only once." />   
				</err_info>
		</exec_function>

		<exec_function function_name="lms_question_tags_loop1" function_type="end_loop" index_name="i" />


<exec_function function_name="lms_question_tags_loop" function_type="start_loop" index_name="i" start_index="0" end_index="len(ReqBrokerReqObj.request_data.tags)" />
		<exec_function function_name="lms_enquiry_comments_validate_fld" function_type="validate_fld">
			<validate_info app_err_desc="pattern validation failed for input" app_corrective_action="check the inputs and provide proper values">
				<validate_fld fld_name="ReqBrokerReqObj.request_data.tags[i]" fld_data_type="string" fld_min_len="1" fld_max_len="500" fld_type="pattern" fld_pattern="^[a-zA-Z0-9.,|_* -]+$"></validate_fld>
			</validate_info>
		</exec_function>

		<exec_function function_name="insert_lms_course_section_db_single_read" function_type="db_single_read">
			<db_single_read_info app_err_desc="lms question bank table is not found" app_corrective_action="question bank data cannot be configured as the table does not exist">
				<table_name>LMS_QUESTION_BANK</table_name>
				<filter_info>
					<filter_data filter_name="COURSE_ID" filter_id="ReqBrokerReqObj.request_data.course_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="SECTION_ID" filter_id="ReqBrokerReqObj.request_data.section_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="ITEM_ID" filter_id="ReqBrokerReqObj.request_data.item_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="QUESTION" filter_id="ReqBrokerReqObj.request_data.question" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_question_bank_table_record</store_obj_name>
				</db_single_read_info>
				<err_info>
				  <err_desc err_code="Duplicate" app_err_desc="Duplicate records found for the question:{{.ReqBrokerReqObj.request_data.question}} provided in the input." app_corrective_action="Make sure each question shows up only once." />   
				</err_info>
		</exec_function>
		<exec_function function_name="lms_question_tags_table_copy_object" function_type="copy_object">
			<copy_object_info object_name="lms_question_tags_table_obj">
				<object_data key="course_id" data_source_type="reqbrokermap" data_source="lms_question_options_table_record.course_id" data_type="string"></object_data>
				<object_data key="section_id" data_source_type="reqbrokermap" data_source="lms_question_options_table_record.section_id" data_type="string"></object_data>
				<object_data key="item_id" data_source_type="reqbrokermap" data_source="lms_question_options_table_record.item_id" data_type="string"></object_data>
				<object_data key="question_id" data_source_type="reqbrokermap" data_source="lms_question_options_table_record.question_id" data_type="string"></object_data>
				<object_data key="option_id" data_source_type="reqbrokermap" data_source="lms_question_options_table_record.id" data_type="string"></object_data>
				<object_data key="tags" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.tags[i]" data_type="string"></object_data>
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
		<exec_function function_name="lms_question_options_table_insert" function_type="db_insert">
			<db_insert_info app_err_desc="pattern validation failed for input" app_corrective_action="check the inputs and provide proper values">
				<table_name>LMS_QUESTION_TAGS</table_name>
				<insert_data_list>
					<insert_data insert_data_name="COURSE_ID" insert_db_data_type="string" insert_data_source="lms_question_tags_table_obj.course_id" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="SECTION_ID" insert_db_data_type="string" insert_data_source="lms_question_tags_table_obj.section_id" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="ITEM_ID" insert_db_data_type="string" insert_data_source="lms_question_tags_table_obj.item_id" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="QUESTION_ID" insert_db_data_type="string" insert_data_source="lms_question_tags_table_obj.question_id" insert_data_type="string"></insert_data>
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
		<exec_function function_name="lms_question_tags_loop" function_type="end_loop" index_name="i" />


		<exec_function function_name="insert_lms_question_bank_db_single_read" function_type="db_single_read">
			<db_single_read_info app_err_desc="lms question bank table is not found" app_corrective_action="question bank data cannot be configured as the table does not exist">
				<table_name>LMS_QUESTION_BANK</table_name>
				<filter_info>
					<filter_data filter_name="COURSE_ID" filter_id="ReqBrokerReqObj.request_data.course_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="SECTION_ID" filter_id="ReqBrokerReqObj.request_data.section_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="ITEM_ID" filter_id="ReqBrokerReqObj.request_data.item_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="QUESTION" filter_id="ReqBrokerReqObj.request_data.question" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_question_bank_table_record</store_obj_name>
				</db_single_read_info>
				<err_info>
				  <err_desc err_code="Duplicate" app_err_desc="Duplicate records found for the question:{{.ReqBrokerReqObj.request_data.question}} provided in the input." app_corrective_action="Make sure each question shows up only once." />   
				</err_info>
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
