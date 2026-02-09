package lexicalfile_lms_question_bank

func GetHandleLMSQuizAttemptServiceXML() interface{} {
	lexicalParserStr := `
	<sem_lexical_parser microservice="handle_lms_quiz_attempt">
	<exec_group group_name="user_token_validation">
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
	<exec_group group_name="insert_lms_course_question_bank">
		<exec_function function_name="lms_quiz_attempt_validate_fld" function_type="validate_fld">
			<validate_info app_err_desc="pattern validation failed for input" app_corrective_action="check the inputs and provide proper values">
				<validate_fld fld_name="ReqBrokerReqObj.request_data.course_id" fld_data_type="string" fld_min_len="1" fld_max_len="500" fld_type="pattern" fld_pattern="^[0-9]+$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.section_id" fld_data_type="string" fld_min_len="1" fld_max_len="500" fld_type="pattern" fld_pattern="^[0-9]+$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.item_id" fld_data_type="string" fld_min_len="1" fld_max_len="500" fld_type="pattern" fld_pattern="^[0-9]+$"></validate_fld>
			</validate_info>
		</exec_function>
			<exec_function function_name="lms_quiz_attempt_validate_filed" function_type="db_multi_read">
				<db_multi_read_info app_err_desc="lms quiz attempt table is not found" app_corrective_action="lms quiz attempt cannot be configured as the table does not exist">
					<table_name>LMS_QUIZ_ATTEMPT</table_name>
					<filter_info>
						<filter_data filter_name="USER_ID" filter_id="decode_token_obj.user_id" filter_data_type="string" filter_condition="AND"></filter_data>
						<filter_data filter_name="COURSE_ID" filter_id="ReqBrokerReqObj.request_data.course_id" filter_data_type="string" filter_condition="AND"></filter_data>
						<filter_data filter_name="SECTION_ID" filter_id="ReqBrokerReqObj.request_data.section_id" filter_data_type="string" filter_condition="AND"></filter_data>
                        <filter_data filter_name="ITEM_ID" filter_id="ReqBrokerReqObj.request_data.item_id" filter_data_type="string" filter_condition="AND"></filter_data>
					</filter_info>
				        <store_obj_name>lms_quiz_attempt_table_list</store_obj_name>
				</db_multi_read_info>
				<err_info>
					<err_desc err_code="Duplicate" app_err_desc="Duplicate records found for the mobile_number:{{.ReqBrokerReqObj.request_data.mobile_number}} provided in the input." app_corrective_action="Make sure each Mobile Number shows up only once." />   
				</err_info>
			</exec_function>
				
			<exec_function function_name="lms_quiz_attempt_loop" function_type="start_loop" index_name="i" start_index="0" end_index="len(lms_quiz_attempt_table_list)" />
           	 <exec_function function_name="role_validate_condition" function_type="validate_condition" function_condition="lms_quiz_attempt_table_list!=nil">
			<validate_condition_info app_err_desc="Student has already passed the quiz and is not allowed to attempt it again." app_corrective_action="No action required. The quiz has already been successfully completed.">
				<validate_condition_expression>lms_quiz_attempt_table_list[i].score&gt;=lms_quiz_attempt_table_list[i].passing_score</validate_condition_expression>
			</validate_condition_info>
		</exec_function>
            
				<exec_function function_name="lms_quiz_attempt_loop" function_type="end_loop" index_name="i" />

			<exec_function function_name="insert_lms_quiz_attempt_copy_object" function_type="copy_object" function_condition="lms_quiz_attempt_table_list==nil">
				<copy_object_info object_name="insert_lms_quiz_attempt_data_table_obj">
					<object_data key="user_id" data_source_type="reqbrokermap" data_source="decode_token_obj.user_id" data_type="string"></object_data>
					<object_data key="course_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.course_id" data_type="string"></object_data>
					<object_data key="section_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.section_id" data_type="string"></object_data>
					<object_data key="item_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.item_id" data_type="string"></object_data>
                    <object_data key="passing_score" data_source_type="raw_value" data_source="80" data_type="string"></object_data>
                    <object_data key="score" data_source_type="raw_value" data_source="0" data_type="string"></object_data>
					<object_data key="submitted_at" data_source_type="key" data_source="get_time_stamp" data_type="string"></object_data>
                    <object_data key="created_by" data_source_type="reqbrokermap" data_source="decode_token_obj.email" data_type="string"></object_data>
					<object_data key="created_date" data_source_type="key" data_source="get_date" data_type="string"></object_data>
					<object_data key="created_time" data_source_type="key" data_source="get_time" data_type="string"></object_data>
					<object_data key="updated_by" data_source_type="reqbrokermap" data_source="decode_token_obj.email" data_type="string"></object_data>
					<object_data key="updated_date" data_source_type="key" data_source="get_date" data_type="string"></object_data>
					<object_data key="updated_time" data_source_type="key" data_source="get_time" data_type="string"></object_data>
					<object_data key="deleted_by" data_source_type="key" data_source="get_na" data_type="string"></object_data>
					<object_data key="deleted_date" data_source_type="key" data_source="get_date" data_type="string"></object_data>
					<object_data key="deleted_time" data_source_type="key" data_source="get_time" data_type="string"></object_data>
				</copy_object_info>
			</exec_function>
			<exec_function function_name="insert_lms_quiz_attempt_data_db_insert" function_type="db_insert" function_condition="lms_quiz_attempt_table_list==nil">
				<db_insert_info>
					<table_name>LMS_QUIZ_ATTEMPT</table_name>
					<insert_data_list>
					   	<insert_data insert_data_name="USER_ID" insert_db_data_type="string" insert_data_source="insert_lms_quiz_attempt_data_table_obj.user_id" insert_data_type="string"></insert_data>
                        <insert_data insert_data_name="COURSE_ID" insert_db_data_type="string" insert_data_source="insert_lms_quiz_attempt_data_table_obj.course_id" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="SECTION_ID" insert_db_data_type="string" insert_data_source="insert_lms_quiz_attempt_data_table_obj.section_id" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="ITEM_ID" insert_db_data_type="string" insert_data_source="insert_lms_quiz_attempt_data_table_obj.item_id" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="PASSING_SCORE" insert_db_data_type="string" insert_data_source="insert_lms_quiz_attempt_data_table_obj.passing_score" insert_data_type="string"></insert_data>
                        <insert_data insert_data_name="SCORE" insert_db_data_type="string" insert_data_source="insert_lms_quiz_attempt_data_table_obj.score" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="SUBMITTED_AT" insert_db_data_type="string" insert_data_source="insert_lms_quiz_attempt_data_table_obj.submitted_at" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="CREATED_BY" insert_db_data_type="string" insert_data_source="insert_lms_quiz_attempt_data_table_obj.created_by" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="CREATED_DATE" insert_db_data_type="date" insert_data_source="insert_lms_quiz_attempt_data_table_obj.created_date" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="CREATED_TIME" insert_db_data_type="string" insert_data_source="insert_lms_quiz_attempt_data_table_obj.created_time" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="UPDATED_BY" insert_db_data_type="string" insert_data_source="insert_lms_quiz_attempt_data_table_obj.updated_by" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="UPDATED_DATE" insert_db_data_type="date" insert_data_source="insert_lms_quiz_attempt_data_table_obj.updated_date" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="UPDATED_TIME" insert_db_data_type="string" insert_data_source="insert_lms_quiz_attempt_data_table_obj.updated_time" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="DELETED_BY" insert_db_data_type="string" insert_data_source="insert_lms_quiz_attempt_data_table_obj.deleted_by" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="DELETED_DATE" insert_db_data_type="date" insert_data_source="insert_lms_quiz_attempt_data_table_obj.deleted_date" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="DELETED_TIME" insert_db_data_type="string" insert_data_source="insert_lms_quiz_attempt_data_table_obj.deleted_time" insert_data_type="string"></insert_data>
						</insert_data_list>
				</db_insert_info>
			</exec_function>
			<exec_function function_name="edit_lms_quiz_attempt_copy_object" function_type="copy_object" function_condition="lms_quiz_attempt_table_list!=nil">
				<copy_object_info object_name="lms_quiz_attempt_additional_obj">
			        <object_data key="user_id" data_source_type="reqbrokermap" data_source="decode_token_obj.user_id" data_type="string"></object_data>
					<object_data key="course_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.course_id" data_type="string"></object_data>
					<object_data key="section_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.section_id" data_type="string"></object_data>
					<object_data key="item_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.item_id" data_type="string"></object_data>
                    <object_data key="passing_score" data_source_type="raw_value" data_source="80" data_type="string"></object_data>
                    <object_data key="score" data_source_type="raw_value" data_source="0" data_type="string"></object_data>
					<object_data key="submitted_at" data_source_type="key" data_source="get_time_stamp" data_type="string"></object_data>
					<object_data key="updated_by" data_source_type="reqbrokermap" data_source="decode_token_obj.email" data_type="string"></object_data>
					<object_data key="updated_date" data_source_type="key" data_source="get_date" data_type="string"></object_data>
					<object_data key="updated_time" data_source_type="key" data_source="get_time" data_type="string"></object_data>
				</copy_object_info>
			</exec_function>
			<exec_function function_name="edit_lms_quiz_attempt_db_update" function_type="db_update" function_condition="lms_quiz_attempt_table_list!=nil">
				<db_update_info>
					<table_name>LMS_QUIZ_ATTEMPT</table_name>
					<update_data_list>
                      <update_data update_data_name="PASSING_SCORE" update_db_data_type="string" update_data_source="lms_quiz_attempt_additional_obj.passing_score" update_data_type="string"></update_data>
                       <update_data update_data_name="SCORE" update_db_data_type="string" update_data_source="lms_quiz_attempt_additional_obj.score" update_data_type="string"></update_data>
						<update_data update_data_name="UPDATED_BY" update_db_data_type="string" update_data_source="lms_quiz_attempt_additional_obj.updated_by" update_data_type="string"></update_data>
						<update_data update_data_name="UPDATED_DATE" update_db_data_type="date" update_data_source="lms_quiz_attempt_additional_obj.updated_date" update_data_type="string"></update_data>
						<update_data update_data_name="UPDATED_TIME" update_db_data_type="string" update_data_source="lms_quiz_attempt_additional_obj.updated_time" update_data_type="string"></update_data>
					</update_data_list>
					<update_filter_list>
					<update_filter update_filter_name="USER_ID" update_filter_db_data_type="string" update_filter_data_source="lms_quiz_attempt_additional_obj.user_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
					<update_filter update_filter_name="COURSE_ID" update_filter_db_data_type="string" update_filter_data_source="lms_quiz_attempt_additional_obj.course_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
					<update_filter update_filter_name="SECTION_ID" update_filter_db_data_type="string" update_filter_data_source="lms_quiz_attempt_additional_obj.section_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
					<update_filter update_filter_name="ITEM_ID" update_filter_db_data_type="string" update_filter_data_source="lms_quiz_attempt_additional_obj.item_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
				</update_filter_list>
				</db_update_info>
			</exec_function>
	</exec_group>
	<exec_group group_name="quiz_attempt_score_calculation">
		<exec_function function_name="quiz_option_map_array" function_type="create_map_array">
			<map_array_info object_name="quiz_obj" array_size="len(ReqBrokerReqObj.request_data.quiz)" />
		</exec_function>
	<exec_function function_name="lms_quiz_question_loop" function_type="start_loop" index_name="i" start_index="0" end_index="len(ReqBrokerReqObj.request_data.quiz)" />
		<exec_function function_name="lms_enquiry_comments_validate_fld" function_type="validate_fld">
			<validate_info app_err_desc="pattern validation failed for input" app_corrective_action="check the inputs and provide proper values">
				<validate_fld fld_name="ReqBrokerReqObj.request_data.quiz[i].question_id" fld_data_type="string" fld_min_len="1" fld_max_len="500" fld_type="pattern" fld_pattern="^[a-zA-Z0-9.,|_*=\-; ]+$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.quiz[i].option_id" fld_data_type="string" fld_min_len="1" fld_max_len="500" fld_type="pattern" fld_pattern="^[a-zA-Z0-9.,|_* -]+$"></validate_fld>
			</validate_info>
		</exec_function>

		<exec_function function_name="lms_question_bank_auth_db_single_read" function_type="db_single_read">
			<db_single_read_info app_err_desc="Data Already Exists or No Rows in Result Set" app_corrective_action="check the inputs and provide proper values">
				<table_name>LMS_QUESTION_BANK</table_name>
				<filter_info>
					<filter_data filter_name="ID" filter_id="ReqBrokerReqObj.request_data.quiz[i].question_id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_question_bank_table_record</store_obj_name>
			</db_single_read_info>
			<err_info>
				<err_desc err_code="Duplicate" app_err_desc="Duplicate record found for the Question ID." app_corrective_action="Make sure question is configured in question bank table."/>
				<err_desc err_code="NoRows" app_err_desc="Duplicate Record Found for question id" app_corrective_action="Make sure question id should be present only once in table."/>
			</err_info>
		</exec_function>
            
		<exec_function function_name="lms_question_options_db_single_read" function_type="db_single_read">
			<db_single_read_info app_err_desc="lms question options table is not found" app_corrective_action="lms question options cannot be configured as the table does not exist">
				<table_name>LMS_QUESTION_OPTIONS</table_name>
				<filter_info>
					<filter_data filter_name="COURSE_ID" filter_id="lms_question_bank_table_record.course_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="SECTION_ID" filter_id="lms_question_bank_table_record.section_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="ITEM_ID" filter_id="lms_question_bank_table_record.item_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="QUESTION_ID" filter_id="lms_question_bank_table_record.id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="ID" filter_id="ReqBrokerReqObj.request_data.quiz[i].option_id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_question_options_table_record</store_obj_name>
			</db_single_read_info>
			<err_info>
				<err_desc err_code="Duplicate" app_err_desc="no record found for the Option ID ." app_corrective_action="Make sure option is configured in question bank table."/>
				<err_desc err_code="NoRows" app_err_desc="Duplicate Record Found for option id" app_corrective_action="Make sure option id should be present only once in table."/>
			</err_info>
		</exec_function>
		
         <exec_function function_name="lms_quiz_attempt_db_single_read" function_type="db_single_read">
				<db_single_read_info app_err_desc="lms quiz attempt data data table is not found" app_corrective_action="lms quiz attempt table  cannot be configured as the table does not exist">
					<table_name>LMS_QUIZ_ATTEMPT</table_name>
					<filter_info>
						<filter_data filter_name="USER_ID" filter_id="decode_token_obj.user_id" filter_data_type="string" filter_condition="AND"></filter_data>
						<filter_data filter_name="COURSE_ID" filter_id="ReqBrokerReqObj.request_data.course_id" filter_data_type="string" filter_condition="AND"></filter_data>
						<filter_data filter_name="SECTION_ID" filter_id="ReqBrokerReqObj.request_data.section_id" filter_data_type="string" filter_condition="AND"></filter_data>
                        <filter_data filter_name="ITEM_ID" filter_id="ReqBrokerReqObj.request_data.item_id" filter_data_type="string" filter_condition="AND"></filter_data>
					</filter_info>
				        <store_obj_name>lms_quiz_attempt_table_record</store_obj_name>
				</db_single_read_info>
				<err_info>
					<err_desc err_code="Duplicate" app_err_desc="Duplicate records found for the mobile_number:{{.ReqBrokerReqObj.request_data.mobile_number}} provided in the input." app_corrective_action="Make sure each Mobile Number shows up only once." />   
				</err_info>
			</exec_function>
             
        <exec_function function_name="lms_quiz_attempt_right_choice" function_type="copy_object" function_condition="lms_question_options_table_record.is_correct==true">
			<copy_object_info object_name="quiz_obj[i]">
			   		<object_data key="points" data_source_type="reqbrokermap" data_source="lms_question_bank_table_record.points" data_type="string"></object_data>
                    <object_data key="total_points" data_source_type="reqbrokermap" data_source="lms_question_bank_table_record.points" data_type="string"></object_data>
                    <object_data key="passing_score" data_source_type="reqbrokermap" data_source="lms_quiz_attempt_table_record.passing_score" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
		<exec_function function_name="lms_quiz_attempt_right_choice" function_type="copy_object" function_condition="lms_question_options_table_record.is_correct!=true">
			<copy_object_info object_name="quiz_obj[i]">
			   		<object_data key="points" data_source_type="raw_value" data_source="0" data_type="string"></object_data>
                    <object_data key="total_points" data_source_type="reqbrokermap" data_source="lms_question_bank_table_record.points" data_type="string"></object_data>
                    <object_data key="passing_score" data_source_type="reqbrokermap" data_source="lms_quiz_attempt_table_record.passing_score" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
        <exec_function function_name="HandleQuizScoreCalc" function_type="call_method"/>
		<exec_function function_name="lms_quiz_question_loop" function_type="end_loop" index_name="i" />
		
		 <exec_function function_name="lms_quiz_attempt" function_type="copy_object">
			<copy_object_info object_name="lms_quiz_attempt_table_obj">
				<object_data key="obtained_points" data_source_type="reqbrokermap" data_source="quiz_obj_summary.obtained_points" data_type="string"></object_data>
				<object_data key="total_points" data_source_type="reqbrokermap" data_source="quiz_obj_summary.total_points" data_type="string"></object_data>
				<object_data key="question_count" data_source_type="reqbrokermap" data_source="quiz_obj_summary.question_count" data_type="string"></object_data>
                <object_data key="average" data_source_type="reqbrokermap" data_source="quiz_obj_summary.average" data_type="string"></object_data>
                <object_data key="is_completed" data_source_type="reqbrokermap" data_source="quiz_obj_summary.is_completed" data_type="Boolean"></object_data>
			   </copy_object_info>
		</exec_function>

		<exec_function function_name="edit_lms_quiz_attempt_copy_object" function_type="copy_object">
				<copy_object_info object_name="lms_quiz_attempt_copy_object">
					<object_data key="updated_by" data_source_type="reqbrokermap" data_source="decode_token_obj.email" data_type="string"></object_data>
					<object_data key="updated_date" data_source_type="key" data_source="get_date" data_type="string"></object_data>
					<object_data key="updated_time" data_source_type="key" data_source="get_time" data_type="string"></object_data>
				</copy_object_info>
			</exec_function>
			
			 <exec_function function_name="edit_lms_quiz_attempt_db_update" function_type="db_update">
				<db_update_info>
					<table_name>LMS_QUIZ_ATTEMPT</table_name>
					<update_data_list>
                       <update_data update_data_name="SCORE" update_db_data_type="string" update_data_source="lms_quiz_attempt_table_obj.average" update_data_type="string"></update_data>
						<update_data update_data_name="UPDATED_BY" update_db_data_type="string" update_data_source="lms_quiz_attempt_copy_object.updated_by" update_data_type="string"></update_data>
						<update_data update_data_name="UPDATED_DATE" update_db_data_type="date" update_data_source="lms_quiz_attempt_copy_object.updated_date" update_data_type="string"></update_data>
						<update_data update_data_name="UPDATED_TIME" update_db_data_type="string" update_data_source="lms_quiz_attempt_copy_object.updated_time" update_data_type="string"></update_data>
					</update_data_list>
					<update_filter_list>
					<update_filter update_filter_name="USER_ID" update_filter_db_data_type="string" update_filter_data_source="decode_token_obj.user_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
					<update_filter update_filter_name="COURSE_ID" update_filter_db_data_type="string" update_filter_data_source="ReqBrokerReqObj.request_data.course_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
					<update_filter update_filter_name="SECTION_ID" update_filter_db_data_type="string" update_filter_data_source="ReqBrokerReqObj.request_data.section_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
					<update_filter update_filter_name="ITEM_ID" update_filter_db_data_type="string" update_filter_data_source="ReqBrokerReqObj.request_data.item_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
				</update_filter_list>
				</db_update_info>
			</exec_function>


			<exec_function function_name="insert_lms_course_section_compose_response" function_type="compose_response" function_condition="lms_quiz_attempt_table_obj.is_completed!=true">
			<response_info>
				<resp_obj obj_name="lms_quiz_attempt_table_obj" obj_type="object" data_source="lms_quiz_attempt_table_obj"></resp_obj>
			</response_info>
		</exec_function>
	</exec_group>
	
	<exec_group group_name="course_item_progress_update_status" group_condition="lms_quiz_attempt_table_obj.is_completed==true">
			<exec_function function_name="lms_course_item_db_single_read" function_type="db_single_read">
			<db_single_read_info app_err_desc="lms course item data table is not found" app_corrective_action="course item data cannot be configured as the table does not exist">
				<table_name>LMS_COURSE_ITEM</table_name>
				<filter_info>
					<filter_data filter_name="COURSE_ID" filter_id="ReqBrokerReqObj.request_data.course_id" filter_data_type="string" filter_condition="AND"></filter_data>
                    <filter_data filter_name="SECTION_ID" filter_id="ReqBrokerReqObj.request_data.section_id" filter_data_type="string" filter_condition="AND"></filter_data>
                    <filter_data filter_name="ID" filter_id="ReqBrokerReqObj.request_data.item_id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_course_item_table_record</store_obj_name>
			</db_single_read_info>
			<err_info>
				<err_desc err_code="NoRows" app_err_desc="Provided course,section,item id not found" app_corrective_action="Please Link the course,section,item and try again" />
                <err_desc err_code="Duplicate" app_err_desc="Duplicate records found for course and section provided in the input." app_corrective_action="Make sure each course item should be linked with one section" />
			</err_info>
		</exec_function>
			<exec_function function_name="update_lms_course_progress_copy_object" function_type="copy_object">
			<copy_object_info object_name="update_lms_course_progress_additional_obj">
				<object_data key="next_order" data_source_type="raw_value" data_source="1" data_type="string"></object_data>
                 <object_data key="user_id" data_source_type="reqbrokermap" data_source="decode_token_obj.user_id" data_type="string"></object_data>
			    <object_data key="course_id" data_source_type="reqbrokermap" data_source="lms_course_item_table_record.course_id" data_type="string"></object_data>
				<object_data key="section_id" data_source_type="reqbrokermap" data_source="lms_course_item_table_record.section_id" data_type="string"></object_data>
				<object_data key="item_id" data_source_type="reqbrokermap" data_source="lms_course_item_table_record.id" data_type="string"></object_data>
                <object_data key="is_enabled" data_source_type="raw_value" data_source="true" data_type="Boolean"></object_data>
                <object_data key="is_completed" data_source_type="raw_value" data_source="true" data_type="Boolean"></object_data>
                <object_data key="is_course_content_completed" data_source_type="raw_value" data_source="true" data_type="Boolean"></object_data>
                <object_data key="completed_at"  data_source_type="key" data_source="get_time_stamp" data_type="string"></object_data>
				<object_data key="updated_by" data_source_type="reqbrokermap" data_source="decode_token_obj.email" data_type="string"></object_data>
				<object_data key="updated_date" data_source_type="key" data_source="get_date" data_type="string"></object_data>
				<object_data key="updated_time" data_source_type="key" data_source="get_time" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
		 <exec_function function_name="edit_lms_course_item_progress_db_update" function_type="db_update">
				<db_update_info>
					<table_name>LMS_COURSE_ITEM_PROGRESS</table_name>
					<update_data_list>
                        <update_data update_data_name="IS_ENABLED" update_db_data_type="string" update_data_source="update_lms_course_progress_additional_obj.is_enabled" update_data_type="boolean"></update_data>
                        <update_data update_data_name="IS_COMPLETED" update_db_data_type="string" update_data_source="update_lms_course_progress_additional_obj.is_completed" update_data_type="boolean"></update_data>
                        <update_data update_data_name="IS_COURSE_CONTENT_COMPLETED" update_db_data_type="string" update_data_source="update_lms_course_progress_additional_obj.is_course_content_completed" update_data_type="boolean"></update_data>
                        <update_data update_data_name="COMPLETED_AT" update_db_data_type="string" update_data_source="update_lms_course_progress_additional_obj.completed_at" update_data_type="string"></update_data>
						<update_data update_data_name="UPDATED_BY" update_db_data_type="string" update_data_source="update_lms_course_progress_additional_obj.updated_by" update_data_type="string"></update_data>
						<update_data update_data_name="UPDATED_DATE" update_db_data_type="date" update_data_source="update_lms_course_progress_additional_obj.updated_date" update_data_type="string"></update_data>
						<update_data update_data_name="UPDATED_TIME" update_db_data_type="string" update_data_source="update_lms_course_progress_additional_obj.updated_time" update_data_type="string"></update_data>
					</update_data_list>
					<update_filter_list>
					<update_filter update_filter_name="USER_ID" update_filter_db_data_type="string" update_filter_data_source="update_lms_course_progress_additional_obj.user_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
					<update_filter update_filter_name="COURSE_ID" update_filter_db_data_type="string" update_filter_data_source="update_lms_course_progress_additional_obj.course_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
					<update_filter update_filter_name="SECTION_ID" update_filter_db_data_type="string" update_filter_data_source="update_lms_course_progress_additional_obj.section_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
					<update_filter update_filter_name="ITEM_ID" update_filter_db_data_type="string" update_filter_data_source="update_lms_course_progress_additional_obj.item_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
				</update_filter_list>
				</db_update_info>
			</exec_function>
		<exec_function function_name="calculate_math" function_type="arith_operation">
	<arith_operation_info operation="sum" left_operand="lms_course_item_table_record.order" right_operand="update_lms_course_progress_additional_obj.next_order" dest_object="order" />
       </exec_function>
            <exec_function function_name="arith_order_copy_object_item_progress" function_type="copy_object">
			<copy_object_info object_name="order_obj">
				<object_data key="order" data_source_type="reqbrokermap" data_source="order" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
        <exec_function function_name="lms_course_item_db_multi_read" function_type="db_multi_read">
			<db_multi_read_info app_err_desc="lms course item data table is not found" app_corrective_action="course item data cannot be configured as the table does not exist">
				<table_name>LMS_COURSE_ITEM</table_name>
				<filter_info>
					<filter_data filter_name="COURSE_ID" filter_id="ReqBrokerReqObj.request_data.course_id" filter_data_type="string" filter_condition="AND"></filter_data>
                    <filter_data filter_name="SECTION_ID" filter_id="ReqBrokerReqObj.request_data.section_id" filter_data_type="string" filter_condition="AND"></filter_data>
                    <filter_data filter_name="ORDER" filter_id="order_obj.order" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_course_item_table_list</store_obj_name>
			</db_multi_read_info>
			<err_info>
                <err_desc err_code="Duplicate" app_err_desc="Duplicate records found for course,section,item,order {{.order}} provided in the input." app_corrective_action="Make sure each course,section,item should be linked with unique order" />
			</err_info>
		</exec_function>
        	<exec_function function_name="lms_course_item_progress_loop" function_type="start_loop" index_name="i" start_index="0" end_index="len(lms_course_item_table_list)" />
			<exec_function function_name="lms_course_item_progress_db_single_read" function_type="db_single_read" function_condition="lms_course_item_table_list!=nil">
			<db_single_read_info app_err_desc="lms course item progress data table is not found" app_corrective_action="course item progress data cannot be configured as the table does not exist">
				<table_name>LMS_COURSE_ITEM_PROGRESS</table_name>
				<filter_info>
				    <filter_data filter_name="USER_ID" filter_id="decode_token_obj.user_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="COURSE_ID" filter_id="ReqBrokerReqObj.request_data.course_id" filter_data_type="string" filter_condition="AND"></filter_data>
                    <filter_data filter_name="SECTION_ID" filter_id="ReqBrokerReqObj.request_data.section_id" filter_data_type="string" filter_condition="AND"></filter_data>
                    <filter_data filter_name="ITEM_ID" filter_id="lms_course_item_table_list[i].id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_course_item_progress_table_record</store_obj_name>
			</db_single_read_info>
			<err_info>
			    <err_desc err_code="NoRows" app_err_desc="course,section,item provided input was not found" app_corrective_action="Make sure given course,section is linked to item" />
                <err_desc err_code="Duplicate" app_err_desc="Duplicate records found for course,section,item provided in the input." app_corrective_action="Make sure each course,section,item should be linked only once." />
			</err_info>
		</exec_function>
         <exec_function function_name="update_lms_course_progress_copy_object" function_type="copy_object" function_condition="lms_course_item_table_list!=nil">
			<copy_object_info object_name="update_lms_course_progress_additional_obj">
				 <object_data key="user_id" data_source_type="reqbrokermap" data_source="decode_token_obj.user_id" data_type="string"></object_data>
			    <object_data key="course_id" data_source_type="reqbrokermap" data_source="lms_course_item_table_list[i].course_id" data_type="string"></object_data>
				<object_data key="section_id" data_source_type="reqbrokermap" data_source="lms_course_item_table_list[i].section_id" data_type="string"></object_data>
				<object_data key="item_id" data_source_type="reqbrokermap" data_source="lms_course_item_table_list[i].id" data_type="string"></object_data>
                <object_data key="is_enabled" data_source_type="raw_value" data_source="true" data_type="Boolean"></object_data>
				<object_data key="updated_by" data_source_type="reqbrokermap" data_source="decode_token_obj.email" data_type="string"></object_data>
				<object_data key="updated_date" data_source_type="key" data_source="get_date" data_type="string"></object_data>
				<object_data key="updated_time" data_source_type="key" data_source="get_time" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
		 <exec_function function_name="edit_lms_course_item_progress_db_update" function_type="db_update" function_condition="lms_course_item_table_list!=nil">
				<db_update_info>
					<table_name>LMS_COURSE_ITEM_PROGRESS</table_name>
					<update_data_list>
                        <update_data update_data_name="IS_ENABLED" update_db_data_type="string" update_data_source="update_lms_course_progress_additional_obj.is_enabled" update_data_type="boolean"></update_data>
						<update_data update_data_name="UPDATED_BY" update_db_data_type="string" update_data_source="update_lms_course_progress_additional_obj.updated_by" update_data_type="string"></update_data>
						<update_data update_data_name="UPDATED_DATE" update_db_data_type="date" update_data_source="update_lms_course_progress_additional_obj.updated_date" update_data_type="string"></update_data>
						<update_data update_data_name="UPDATED_TIME" update_db_data_type="string" update_data_source="update_lms_course_progress_additional_obj.updated_time" update_data_type="string"></update_data>
					</update_data_list>
					<update_filter_list>
					<update_filter update_filter_name="USER_ID" update_filter_db_data_type="string" update_filter_data_source="update_lms_course_progress_additional_obj.user_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
					<update_filter update_filter_name="COURSE_ID" update_filter_db_data_type="string" update_filter_data_source="update_lms_course_progress_additional_obj.course_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
					<update_filter update_filter_name="SECTION_ID" update_filter_db_data_type="string" update_filter_data_source="update_lms_course_progress_additional_obj.section_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
					<update_filter update_filter_name="ITEM_ID" update_filter_db_data_type="string" update_filter_data_source="update_lms_course_progress_additional_obj.item_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
				</update_filter_list>
				</db_update_info>
			</exec_function>
        <exec_function function_name="lms_course_item_progress_loop" function_type="end_loop" index_name="i" />

        <exec_function function_name="lms_course_section_progress_db_single_read" function_type="db_single_read" function_condition="lms_course_item_table_list==nil">
			<db_single_read_info app_err_desc="lms course section progress data table is not found" app_corrective_action="course section progress data cannot be configured as the table does not exist">
				<table_name>LMS_COURSE_SECTION_PROGRESS</table_name>
				<filter_info>
				    <filter_data filter_name="USER_ID" filter_id="decode_token_obj.user_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="COURSE_ID" filter_id="ReqBrokerReqObj.request_data.course_id" filter_data_type="string" filter_condition="AND"></filter_data>
                    <filter_data filter_name="SECTION_ID" filter_id="ReqBrokerReqObj.request_data.section_id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_course_section_progress_table_record</store_obj_name>
			</db_single_read_info>
			<err_info>
			    <err_desc err_code="NoRows" app_err_desc="course,section provided input was not found" app_corrective_action="Make sure given course is linked to section" />
                <err_desc err_code="Duplicate" app_err_desc="Duplicate records found for course,section provided in the input." app_corrective_action="Make sure each course,section should be linked only once." />
			</err_info>
		</exec_function>
		<exec_function function_name="update_lms_course_progress_copy_object" function_type="copy_object" function_condition="lms_course_item_table_list==nil">
			<copy_object_info object_name="update_lms_course_section_progress_additional_obj">
				<object_data key="next_order" data_source_type="raw_value" data_source="1" data_type="string"></object_data>
                 <object_data key="user_id" data_source_type="reqbrokermap" data_source="decode_token_obj.user_id" data_type="string"></object_data>
			    <object_data key="course_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.course_id" data_type="string"></object_data>
				<object_data key="section_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.section_id" data_type="string"></object_data>
                <object_data key="is_enabled" data_source_type="raw_value" data_source="true" data_type="Boolean"></object_data>
                <object_data key="is_completed" data_source_type="raw_value" data_source="true" data_type="Boolean"></object_data>
                <object_data key="completed_at"  data_source_type="key" data_source="get_time_stamp" data_type="string"></object_data>
				<object_data key="updated_by" data_source_type="reqbrokermap" data_source="decode_token_obj.email" data_type="string"></object_data>
				<object_data key="updated_date" data_source_type="key" data_source="get_date" data_type="string"></object_data>
				<object_data key="updated_time" data_source_type="key" data_source="get_time" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
          <exec_function function_name="edit_lms_course_section_progress_db_update" function_type="db_update" function_condition="lms_course_item_table_list==nil">
				<db_update_info>
					<table_name>LMS_COURSE_SECTION_PROGRESS</table_name>
					<update_data_list>
                        <update_data update_data_name="IS_ENABLED" update_db_data_type="string" update_data_source="update_lms_course_section_progress_additional_obj.is_enabled" update_data_type="boolean"></update_data>
                        <update_data update_data_name="IS_COMPLETED" update_db_data_type="string" update_data_source="update_lms_course_section_progress_additional_obj.is_completed" update_data_type="boolean"></update_data>
                       <update_data update_data_name="COMPLETED_AT" update_db_data_type="string" update_data_source="update_lms_course_section_progress_additional_obj.completed_at" update_data_type="string"></update_data>
						<update_data update_data_name="UPDATED_BY" update_db_data_type="string" update_data_source="update_lms_course_section_progress_additional_obj.updated_by" update_data_type="string"></update_data>
						<update_data update_data_name="UPDATED_DATE" update_db_data_type="date" update_data_source="update_lms_course_section_progress_additional_obj.updated_date" update_data_type="string"></update_data>
						<update_data update_data_name="UPDATED_TIME" update_db_data_type="string" update_data_source="update_lms_course_section_progress_additional_obj.updated_time" update_data_type="string"></update_data>
					</update_data_list>
					<update_filter_list>
					<update_filter update_filter_name="USER_ID" update_filter_db_data_type="string" update_filter_data_source="update_lms_course_section_progress_additional_obj.user_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
					<update_filter update_filter_name="COURSE_ID" update_filter_db_data_type="string" update_filter_data_source="update_lms_course_section_progress_additional_obj.course_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
					<update_filter update_filter_name="SECTION_ID" update_filter_db_data_type="string" update_filter_data_source="update_lms_course_section_progress_additional_obj.section_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
				</update_filter_list>
				</db_update_info>
			</exec_function>
        <exec_function function_name="lms_course_section_db_single_read" function_type="db_single_read" function_condition="lms_course_item_table_list==nil">
			<db_single_read_info app_err_desc="lms course item progress data table is not found" app_corrective_action="course item progress data cannot be configured as the table does not exist">
				<table_name>LMS_COURSE_SECTION</table_name>
				<filter_info>
					<filter_data filter_name="COURSE_ID" filter_id="ReqBrokerReqObj.request_data.course_id" filter_data_type="string" filter_condition="AND"></filter_data>
                   <filter_data filter_name="ID" filter_id="ReqBrokerReqObj.request_data.section_id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_course_section_table_record</store_obj_name>
			</db_single_read_info>
			<err_info>
			    <err_desc err_code="NoRows" app_err_desc="course,section provided input was not found" app_corrective_action="Make sure given course is linked to section" />
                <err_desc err_code="Duplicate" app_err_desc="Duplicate records found for course,section provided in the input." app_corrective_action="Make sure each course,section should be linked only once." />
			</err_info>
		</exec_function>

		<exec_function function_name="calculate_math" function_type="arith_operation" function_condition="lms_course_item_table_list==nil">
	<arith_operation_info operation="sum" left_operand="lms_course_section_table_record.order" right_operand="update_lms_course_section_progress_additional_obj.next_order" dest_object="order" />
       </exec_function>
	     <exec_function function_name="arith_order_copy_object_course_section_progress" function_type="copy_object" function_condition="lms_course_item_table_list==nil">
			<copy_object_info object_name="order_obj">
				<object_data key="order" data_source_type="reqbrokermap" data_source="order" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
        <exec_function function_name="lms_course_section_db_multi_read" function_type="db_multi_read" function_condition="lms_course_item_table_list==nil">
			<db_multi_read_info app_err_desc="lms course item progress data table is not found" app_corrective_action="course item progress data cannot be configured as the table does not exist">
				<table_name>LMS_COURSE_SECTION</table_name>
				<filter_info>
					<filter_data filter_name="COURSE_ID" filter_id="ReqBrokerReqObj.request_data.course_id" filter_data_type="string" filter_condition="AND"></filter_data>
                    <filter_data filter_name="ORDER" filter_id="order_obj.order" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_course_section_table_list</store_obj_name>
			</db_multi_read_info>
			<err_info>
			    <err_desc err_code="NoRows" app_err_desc="course,section provided input was not found" app_corrective_action="Make sure given course is linked to section" />
                <err_desc err_code="Duplicate" app_err_desc="Duplicate records found for course,section provided in the input." app_corrective_action="Make sure each course,section should be linked only once." />
			</err_info>
		</exec_function>
				<exec_function function_name="lms_course_section_loop" function_type="start_loop" index_name="i" start_index="0" end_index="len(lms_course_section_table_list)" />

		  <exec_function function_name="section_copy_object" function_type="copy_object" function_condition="lms_course_item_table_list==nil &amp;&amp; lms_course_section_table_list!=nil">
			<copy_object_info object_name="section_next_obj">
				<object_data key="section_id" data_source_type="reqbrokermap" data_source="lms_course_section_table_list[i].id" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
		<exec_function function_name="edit_lms_course_section_progress_db_update" function_type="db_update" function_condition="lms_course_item_table_list==nil &amp;&amp; lms_course_section_table_list!=nil">
				<db_update_info>
					<table_name>LMS_COURSE_SECTION_PROGRESS</table_name>
					<update_data_list>
                        <update_data update_data_name="IS_ENABLED" update_db_data_type="string" update_data_source="update_lms_course_section_progress_additional_obj.is_enabled" update_data_type="boolean"></update_data>
						<update_data update_data_name="UPDATED_BY" update_db_data_type="string" update_data_source="update_lms_course_section_progress_additional_obj.updated_by" update_data_type="string"></update_data>
						<update_data update_data_name="UPDATED_DATE" update_db_data_type="date" update_data_source="update_lms_course_section_progress_additional_obj.updated_date" update_data_type="string"></update_data>
						<update_data update_data_name="UPDATED_TIME" update_db_data_type="string" update_data_source="update_lms_course_section_progress_additional_obj.updated_time" update_data_type="string"></update_data>
					</update_data_list>
					<update_filter_list>
					<update_filter update_filter_name="USER_ID" update_filter_db_data_type="string" update_filter_data_source="update_lms_course_section_progress_additional_obj.user_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
					<update_filter update_filter_name="COURSE_ID" update_filter_db_data_type="string" update_filter_data_source="update_lms_course_section_progress_additional_obj.course_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
					<update_filter update_filter_name="SECTION_ID" update_filter_db_data_type="string" update_filter_data_source="section_next_obj.section_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
				</update_filter_list>
				</db_update_info>
			</exec_function>
			 <exec_function function_name="lms_course_item_next_section_item_db_single_read" function_type="db_single_read" function_condition="lms_course_item_table_list==nil &amp;&amp; lms_course_section_table_list!=nil">
			<db_single_read_info app_err_desc="lms course item data table is not found" app_corrective_action="course item data cannot be configured as the table does not exist">
				<table_name>LMS_COURSE_ITEM</table_name>
				<filter_info>
					<filter_data filter_name="COURSE_ID" filter_id="ReqBrokerReqObj.request_data.course_id" filter_data_type="string" filter_condition="AND"></filter_data>
                    <filter_data filter_name="SECTION_ID" filter_id="section_next_obj.section_id" filter_data_type="string" filter_condition="AND"></filter_data>
                    <filter_data filter_name="ORDER"   filter_source_type="raw_value" filter_id="1" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_course_item_table_record</store_obj_name>
			</db_single_read_info>
			<err_info>
                <err_desc err_code="Duplicate" app_err_desc="Duplicate records found for course,section,item,order {{.order}} provided in the input." app_corrective_action="Make sure each course,section,item should be linked with unique order" />
			</err_info>
		</exec_function>
		    <exec_function function_name="item_next_section_copy_object" function_type="copy_object" function_condition="lms_course_item_table_list==nil &amp;&amp; lms_course_section_table_list!=nil">
			<copy_object_info object_name="item_next_section_obj">
                    <object_data key="item_id" data_source_type="reqbrokermap" data_source="lms_course_item_table_record.id" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
			 <exec_function function_name="edit_lms_course_item_progress_next_section_db_update" function_type="db_update" function_condition="lms_course_item_table_list==nil &amp;&amp; lms_course_section_table_list!=nil">
				<db_update_info>
					<table_name>LMS_COURSE_ITEM_PROGRESS</table_name>
					<update_data_list>
                        <update_data update_data_name="IS_ENABLED" update_db_data_type="string" update_data_source="update_lms_course_progress_additional_obj.is_enabled" update_data_type="boolean"></update_data>
						<update_data update_data_name="UPDATED_BY" update_db_data_type="string" update_data_source="update_lms_course_progress_additional_obj.updated_by" update_data_type="string"></update_data>
						<update_data update_data_name="UPDATED_DATE" update_db_data_type="date" update_data_source="update_lms_course_progress_additional_obj.updated_date" update_data_type="string"></update_data>
						<update_data update_data_name="UPDATED_TIME" update_db_data_type="string" update_data_source="update_lms_course_progress_additional_obj.updated_time" update_data_type="string"></update_data>
					</update_data_list>
					<update_filter_list>
					<update_filter update_filter_name="USER_ID" update_filter_db_data_type="string" update_filter_data_source="update_lms_course_progress_additional_obj.user_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
					<update_filter update_filter_name="COURSE_ID" update_filter_db_data_type="string" update_filter_data_source="update_lms_course_progress_additional_obj.course_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
					<update_filter update_filter_name="SECTION_ID" update_filter_db_data_type="string" update_filter_data_source="section_next_obj.section_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
					<update_filter update_filter_name="ITEM_ID" update_filter_db_data_type="string" update_filter_data_source="item_next_section_obj.item_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
				</update_filter_list>
				</db_update_info>
			</exec_function>
                   <exec_function function_name="lms_course_section_loop" function_type="end_loop" index_name="i" />
				   <exec_function function_name="edit_lms_course_progress_db_update" function_type="db_update" function_condition="lms_course_item_table_list==nil &amp;&amp;lms_course_section_table_list==nil">
				<db_update_info>
					<table_name>LMS_COURSE_PROGRESS</table_name>
					<update_data_list>
                        <update_data update_data_name="IS_COMPLETED" update_db_data_type="string" update_data_source="update_lms_course_section_progress_additional_obj.is_completed" update_data_type="boolean"></update_data>
                        <update_data update_data_name="COMPLETED_AT" update_db_data_type="string" update_data_source="update_lms_course_section_progress_additional_obj.completed_at" update_data_type="string"></update_data>
						<update_data update_data_name="UPDATED_BY" update_db_data_type="string" update_data_source="update_lms_course_section_progress_additional_obj.updated_by" update_data_type="string"></update_data>
						<update_data update_data_name="UPDATED_DATE" update_db_data_type="date" update_data_source="update_lms_course_section_progress_additional_obj.updated_date" update_data_type="string"></update_data>
						<update_data update_data_name="UPDATED_TIME" update_db_data_type="string" update_data_source="update_lms_course_section_progress_additional_obj.updated_time" update_data_type="string"></update_data>
					</update_data_list>
					<update_filter_list>
					<update_filter update_filter_name="USER_ID" update_filter_db_data_type="string" update_filter_data_source="update_lms_course_section_progress_additional_obj.user_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
					<update_filter update_filter_name="COURSE_ID" update_filter_db_data_type="string" update_filter_data_source="update_lms_course_section_progress_additional_obj.course_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
				</update_filter_list>
				</db_update_info>
			</exec_function>
		<exec_function function_name="insert_lms_course_section_compose_response" function_type="compose_response">
			<response_info>
				<resp_obj obj_name="lms_quiz_attempt_table_obj" obj_type="object" data_source="lms_quiz_attempt_table_obj"></resp_obj>
			</response_info>
				</exec_function>
	</exec_group>
</sem_lexical_parser>
	`
	return lexicalParserStr
}
