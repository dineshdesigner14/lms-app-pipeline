package lexicalfile_lms_question_bank

func GetViewLMSQuestionBankProdServiceXML() interface{} {
	lexicalParserStr := `
	<sem_lexical_parser microservice="view_lms_question_bank">
	<exec_group group_name="validate_view_lms_question_bank"  group_condition="ReqBrokerReqObj.header_data.channel=='web_admin_portal' &amp;&amp; (ReqBrokerReqObj.request_data.question_id==nil &amp;&amp; ReqBrokerReqObj.request_data.option_id==nil &amp;&amp; ReqBrokerReqObj.request_data.id==nil) &amp;&amp; (ReqBrokerReqObj.request_data.page==nil &amp;&amp; ReqBrokerReqObj.request_data.per_page==nil)" >
		<exec_function function_name="read_view_lms_question_bank_db_multi_read" function_type="db_multi_read">
			<db_multi_read_info>
				<table_name>LMS_QUESTION_BANK</table_name>
				<store_obj_name>lms_question_bank_table_list</store_obj_name>
			</db_multi_read_info>
		</exec_function>
		<exec_function function_name="lms_question_bank_loop" function_type="start_loop" index_name="i" start_index="0" end_index="len(lms_question_bank_table_list)" />
		<exec_function function_name="read_view_lms_question_bank_db_multi_read" function_type="db_multi_read">
			<db_multi_read_info>
				<table_name>LMS_QUESTION_OPTIONS</table_name>
				<filter_info>
					<filter_data filter_name="COURSE_ID" filter_id="lms_question_bank_table_list[i].course_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="SECTION_ID" filter_id="lms_question_bank_table_list[i].section_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="ITEM_ID" filter_id="lms_question_bank_table_list[i].item_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="QUESTION_ID" filter_id="lms_question_bank_table_list[i].id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_question_bank_table_list[i].lms_question_options_table_list</store_obj_name>
			</db_multi_read_info>
		</exec_function>
		<exec_function function_name="lms_question_options_loop" function_type="start_loop" index_name="j" start_index="0" end_index="len(lms_question_bank_table_list[i].lms_question_options_table_list)" />
		<exec_function function_name="read_view_lms_question_bank_db_multi_read" function_type="db_multi_read">
			<db_multi_read_info>
				<table_name>LMS_QUESTION_TAGS</table_name>
				<filter_info>
					<filter_data filter_name="COURSE_ID" filter_id="lms_question_bank_table_list[i].course_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="SECTION_ID" filter_id="lms_question_bank_table_list[i].section_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="ITEM_ID" filter_id="lms_question_bank_table_list[i].item_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="QUESTION_ID" filter_id="lms_question_bank_table_list[i].id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="OPTION_ID" filter_id="lms_question_bank_table_list[i].lms_question_options_table_list[j].id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_question_bank_table_list[i].lms_question_options_table_list[j].lms_question_tags</store_obj_name>
			</db_multi_read_info>
		</exec_function>
		<exec_function function_name="lms_question_options_loop" function_type="end_loop" index_name="j" />
		<exec_function function_name="lms_question_bank_loop" function_type="end_loop" index_name="i" />
		<exec_function function_name="view_lms_course_section_compose_response" function_type="compose_response">
			<response_info>
				<resp_obj obj_name="lms_question_bank_table_list" obj_type="objectarray" data_source="lms_question_bank_table_list"></resp_obj>
			</response_info>
		</exec_function>
	</exec_group>
</sem_lexical_parser>
	`
	return lexicalParserStr
}
