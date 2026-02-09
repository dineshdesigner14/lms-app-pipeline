package lexicalfile_lms_question_bank

func GetFetchLMSQuestionBankServiceXML() interface{} {
	lexicalParserStr := `
	<sem_lexical_parser microservice="view_lms_question_bank">
	<exec_group group_name="user_token_validation" group_condition="ReqBrokerReqObj.header_data.channel=='web_admin_portal'">
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
	</exec_group>
	<exec_group group_name="validate_view_lms_question_bank"  group_condition="(ReqBrokerReqObj.request_data.course_id!=nil &amp;&amp; ReqBrokerReqObj.request_data.section_id!=nil &amp;&amp; ReqBrokerReqObj.request_data.item_id!=nil) &amp;&amp; (ReqBrokerReqObj.request_data.page==nil &amp;&amp; ReqBrokerReqObj.request_data.per_page==nil)" >
		<exec_function function_name="read_view_lms_question_bank_db_multi_read" function_type="db_multi_read">
			<db_multi_read_info>
				<table_name>LMS_QUESTION_BANK</table_name>
				<filter_info>
					<filter_data filter_name="COURSE_ID" filter_id="ReqBrokerReqObj.request_data.course_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="SECTION_ID" filter_id="ReqBrokerReqObj.request_data.section_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="ITEM_ID" filter_id="ReqBrokerReqObj.request_data.item_id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_question_bank_table_list</store_obj_name>
			</db_multi_read_info>
		</exec_function>
		<exec_function function_name="course_progress_map_array" function_type="create_map_array">
			<map_array_info object_name="questions" array_size="len(lms_question_bank_table_list)" />
		</exec_function>
		
		<exec_function function_name="lms_question_bank_loop" function_type="start_loop" index_name="i" start_index="0" end_index="len(lms_question_bank_table_list)" />
		<exec_function function_name="lms_course_item_copy_object1" function_type="copy_object">
			<copy_object_info object_name="questions[i]">
				<object_data key="question_id" data_source_type="reqbrokermap" data_source="lms_question_bank_table_list[i].id" data_type="string"></object_data>
				<object_data key="question" data_source_type="reqbrokermap" data_source="lms_question_bank_table_list[i].question" data_type="string"></object_data>
				<object_data key="question_type" data_source_type="raw_value" data_source="MCQ_SINGLE" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
		<exec_function function_name="read_view_lms_question_bank_db_multi_read" function_type="db_multi_read">
			<db_multi_read_info>
				<table_name>LMS_QUESTION_OPTIONS</table_name>
				<filter_info>
					<filter_data filter_name="COURSE_ID" filter_id="ReqBrokerReqObj.request_data.course_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="SECTION_ID" filter_id="ReqBrokerReqObj.request_data.section_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="ITEM_ID" filter_id="ReqBrokerReqObj.request_data.item_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="QUESTION_ID" filter_id="lms_question_bank_table_list[i].id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_questions_options_table_list</store_obj_name>
			</db_multi_read_info>
		</exec_function>
		<exec_function function_name="lms_question_options_loop" function_type="start_loop" index_name="j" start_index="0" end_index="len(lms_questions_options_table_list)" />

		<exec_function function_name="lms_course_item_copy_object2" function_type="copy_array">
			<copy_object_info object_name="questions[i].options[j]">
				<object_data key="id" data_source_type="reqbrokermap" data_source="lms_questions_options_table_list[j].id" data_type="string"></object_data>
				<object_data key="order" data_source_type="reqbrokermap" data_source="lms_questions_options_table_list[j].order" data_type="string"></object_data>
				<object_data key="text" data_source_type="reqbrokermap" data_source="lms_questions_options_table_list[j].option_text" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
         <exec_function function_name="HandleRandomQuestionShuffle" function_type="call_method"/>
		<exec_function function_name="lms_question_options_loop" function_type="end_loop" index_name="j" />

		<exec_function function_name="lms_question_bank_loop" function_type="end_loop" index_name="i" />
		<exec_function function_name="view_lms_course_section_compose_response" function_type="compose_response">
			<response_info>
				<resp_obj obj_name="question_bank_list" obj_type="objectarray" data_source="question_bank_list"></resp_obj>
			</response_info>
		</exec_function>
	</exec_group>
</sem_lexical_parser>
	`
	return lexicalParserStr
}
