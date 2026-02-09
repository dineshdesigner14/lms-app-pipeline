package lexicalfile_lms_enquiry

func GetInsertLMSEnquiryServiceXML() interface{} {
	lexicalParserStr := `
	<sem_lexical_parser microservice="insert_lms_enquiry">
	<exec_group group_name="validate_lms_enquiry"  group_condition="ReqBrokerReqObj.header_data.channel=='web_client_portal'" >
		<exec_function function_name="lms_enquiry_validate_fld" function_type="validate_fld">
			<validate_info app_err_desc="pattern validation failed for input" app_corrective_action="check the inputs and provide proper values">
				<validate_fld fld_name="ReqBrokerReqObj.request_data.first_name" fld_data_type="string" fld_min_len="1" fld_max_len="128" fld_type="pattern" fld_pattern="^[a-zA-Z0-9.,|_* -]+$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.last_name" fld_data_type="string" fld_min_len="1" fld_max_len="128" fld_type="pattern" fld_pattern="^[a-zA-Z0-9.,|_* -]+$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.phone_number" fld_data_type="string" fld_min_len="1" fld_max_len="128" fld_type="pattern" fld_pattern="^\+\d{1,3}-\d{10}$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.email" fld_data_type="string" fld_min_len="1" fld_max_len="128" fld_type="pattern" fld_pattern="^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,3}$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.message" fld_data_type="string" fld_min_len="0" fld_max_len="500" fld_type="pattern" fld_pattern="^[a-zA-Z0-9.,|_* -]+$"></validate_fld>
			</validate_info>
		</exec_function>
			<exec_function function_name="lms_enquiry_db_multi_read" function_type="db_multi_read">
			<db_multi_read_info app_err_desc="Data Already Exists or No Rows in Result Set" app_corrective_action="check the inputs and provide proper values">
				<table_name>LMS_ENQUIRY</table_name>
				<filter_info>
						<filter_data filter_name="EMAIL" filter_id="ReqBrokerReqObj.request_data.email" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_enquiry_table_list</store_obj_name>
				</db_multi_read_info>
		</exec_function>
		<exec_function function_name="lms_enquiry_validate_condition" function_type="validate_condition" function_condition="lms_enquiry_table_list!=nil">
			<validate_condition_info app_err_desc="Duplication Record Found For Email" app_corrective_action="Please Enter unqiue Email">
				<validate_condition_expression>len(lms_enquiry_table_list)==1</validate_condition_expression>
			</validate_condition_info>
		</exec_function>
        <exec_function function_name="lms_enquiry_loop" function_type="start_loop" index_name="i" start_index="0" end_index="len(lms_enquiry_table_list)" />
		<exec_function function_name="lms_is_email_verified_validate_condition" function_type="validate_condition" function_condition="lms_enquiry_table_list!=nil">
			<validate_condition_info app_err_desc="Email is already verified." app_corrective_action="The email address is already verified. Kindly reach out to the helpline number for further support">
				<validate_condition_expression>!lms_enquiry_table_list[i].is_email_verified</validate_condition_expression>
			</validate_condition_info>
		</exec_function>
		<exec_function function_name="edit_lms_enquiry_personal_info_table_copy_object" function_type="copy_object" function_condition="lms_enquiry_table_list!=nil">	
			<copy_object_info object_name="lms_enquiry_table_obj">
				<object_data key="first_name" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.first_name" data_type="string"></object_data>
				<object_data key="last_name" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.last_name" data_type="string"></object_data>
				<object_data key="phone_number" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.phone_number" data_type="string"></object_data>
				<object_data key="email" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.email" data_type="string"></object_data>
				<object_data key="email_verified_at" data_source_type="key" data_source="get_time_stamp" data_type="string"></object_data>
				<object_data key="updated_by" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.email" data_type="string"></object_data>
				<object_data key="updated_date" data_source_type="key" data_source="get_date" data_type="string"></object_data>
				<object_data key="updated_time" data_source_type="key" data_source="get_time" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
		
		<exec_function function_name="lms_enquiry_personal_info_db_update" function_type="db_update" function_condition="lms_enquiry_table_list!=nil">
			<db_update_info>
				<table_name>LMS_ENQUIRY</table_name>
				<update_data_list>
					<update_data update_data_name="FIRST_NAME" update_db_data_type="string" update_data_source="lms_enquiry_table_obj.first_name" update_data_type="string"></update_data>
					<update_data update_data_name="LAST_NAME" update_db_data_type="string" update_data_source="lms_enquiry_table_obj.last_name" update_data_type="string"></update_data>
					<update_data update_data_name="PHONE_NUMBER" update_db_data_type="string" update_data_source="lms_enquiry_table_obj.phone_number" update_data_type="string"></update_data>
                    <update_data update_data_name="IS_EMAIL_VERIFIED" update_db_data_type="boolean" update_data_source="ReqBrokerReqObj.request_data.is_email_verified" update_data_type="boolean"></update_data>
                    <update_data update_data_name="EMAIL_VERIFIED_AT" update_db_data_type="string" update_data_source="lms_enquiry_table_obj.email_verified_at" update_data_type="string"></update_data>
					<update_data update_data_name="UPDATED_BY" update_db_data_type="string" update_data_source="lms_enquiry_table_obj.updated_by" update_data_type="string"></update_data>
					<update_data update_data_name="UPDATED_DATE" update_db_data_type="date" update_data_source="lms_enquiry_table_obj.updated_date" update_data_type="string"></update_data>
					<update_data update_data_name="UPDATED_TIME" update_db_data_type="string" update_data_source="lms_enquiry_table_obj.updated_time" update_data_type="string"></update_data>
				</update_data_list>
				<update_filter_list>
					<update_filter update_filter_name="EMAIL" update_filter_db_data_type="string" update_filter_data_source="ReqBrokerReqObj.request_data.email" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
				</update_filter_list>
			</db_update_info>
		</exec_function>
		<exec_function function_name="lms_enquiry_loop" function_type="end_loop" index_name="i" />
         
		<exec_function function_name="lms_enquiry_db_single_read_validate" function_type="db_single_read" function_condition="lms_enquiry_table_list==nil">
			<db_single_read_info app_err_desc="Data Already Exists or No Rows in Result Set" app_corrective_action="check the inputs and provide proper values">
				<table_name>LMS_ENQUIRY</table_name>
				<filter_info>
						<filter_data filter_name="PHONE_NUMBER" filter_id="ReqBrokerReqObj.request_data.phone_number" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<result_info app_err_desc="The mobile_number:{{.ReqBrokerReqObj.request_data.phone_number}} already exists." app_corrective_action="Please enter new mobile number and  and proceed">
						<result_success>NoRows</result_success>
					</result_info>
				</db_single_read_info>
				<err_info>
					<err_desc err_code="Duplicate" app_err_desc="Duplicate records found for the mobile_number:{{.ReqBrokerReqObj.request_data.phone_number}} provided in the input." app_corrective_action="Make sure each Mobile Number and  shows up only once." />   
				</err_info>
		</exec_function>
	
		<exec_function function_name="lms_enquiry_table_copy_object" function_type="copy_object" function_condition="lms_enquiry_table_list==nil">
			<copy_object_info object_name="lms_enquiry_table_obj"> 
				<object_data key="first_name" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.first_name" data_type="string"></object_data>
				<object_data key="last_name" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.last_name" data_type="string"></object_data>
				<object_data key="phone_number" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.phone_number" data_type="string"></object_data>
				<object_data key="email" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.email" data_type="string"></object_data>
            	<object_data key="is_email_verified" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.is_email_verified" data_type="Boolean"></object_data>
				<object_data key="email_verified_at" data_source_type="key" data_source="get_time_stamp" data_type="string"></object_data>
				<object_data key="onboarded" data_source_type="raw_value" data_source="false" data_type="Boolean"></object_data>
				<object_data key="onboarding_status" data_source_type="raw_value" data_source="NOT STARTED" data_type="string"></object_data>
				<object_data key="education" data_source_type="raw_value" data_source="NA" data_type="string"></object_data>
				<object_data key="college" data_source_type="raw_value" data_source="NA" data_type="string"></object_data>
				<object_data key="department" data_source_type="raw_value" data_source="NA" data_type="string"></object_data>
				<object_data key="address" data_source_type="raw_value" data_source="NA" data_type="string"></object_data>
				<object_data key="id_proof_type" data_source_type="raw_value" data_source="NA" data_type="string"></object_data>
				<object_data key="id_proof_value" data_source_type="raw_value" data_source="NA" data_type="string"></object_data>
				<object_data key="payment_status" data_source_type="raw_value" data_source="NOT PAID" data_type="string"></object_data>
				<object_data key="payment_mode" data_source_type="raw_value" data_source="NA" data_type="string"></object_data>
				<object_data key="payment_id" data_source_type="raw_value" data_source="NA" data_type="string"></object_data>
				<object_data key="payment_amount" data_source_type="raw_value" data_source="0" data_type="string"></object_data>
				<object_data key="payment_date" data_source_type="key" data_source="get_time_stamp" data_type="string"></object_data>
				<object_data key="message" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.message" data_type="string"></object_data>
				<object_data key="created_by" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.email" data_type="string"></object_data>
				<object_data key="created_date" data_source_type="key" data_source="get_date" data_type="string"></object_data>
				<object_data key="created_time" data_source_type="key" data_source="get_time" data_type="string"></object_data>
				<object_data key="updated_by" data_source_type="key" data_source="get_na" data_type="string"></object_data>
				<object_data key="updated_date" data_source_type="key" data_source="get_date" data_type="string"></object_data>
				<object_data key="updated_time" data_source_type="key" data_source="get_time" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
		<exec_function function_name="lms_enquiry_table_insert" function_type="db_insert" function_condition="lms_enquiry_table_list==nil">
			<db_insert_info>
				<table_name>LMS_ENQUIRY</table_name>
				<insert_data_list>
					<insert_data insert_data_name="FIRST_NAME" insert_db_data_type="string" insert_data_source="lms_enquiry_table_obj.first_name" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="LAST_NAME" insert_db_data_type="string" insert_data_source="lms_enquiry_table_obj.last_name" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="EMAIL" insert_db_data_type="string" insert_data_source="lms_enquiry_table_obj.email" insert_data_type="string"></insert_data>
                    <insert_data insert_data_name="IS_EMAIL_VERIFIED" insert_db_data_type="boolean" insert_data_source="ReqBrokerReqObj.request_data.is_email_verified" insert_data_type="Boolean"></insert_data>
                    <insert_data insert_data_name="EMAIL_VERIFIED_AT" insert_db_data_type="string" insert_data_source="lms_enquiry_table_obj.email_verified_at" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="PHONE_NUMBER" insert_db_data_type="string" insert_data_source="lms_enquiry_table_obj.phone_number" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="ONBOARDED" insert_db_data_type="boolean" insert_data_source="lms_enquiry_table_obj.onboarded" insert_data_type="Boolean"></insert_data>
					<insert_data insert_data_name="ONBOARDING_STATUS" insert_db_data_type="string" insert_data_source="lms_enquiry_table_obj.onboarding_status" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="EDUCATION" insert_db_data_type="string" insert_data_source="lms_enquiry_table_obj.education" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="COLLEGE" insert_db_data_type="string" insert_data_source="lms_enquiry_table_obj.college" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="DEPARTMENT" insert_db_data_type="string" insert_data_source="lms_enquiry_table_obj.department" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="ADDRESS" insert_db_data_type="string" insert_data_source="lms_enquiry_table_obj.address" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="ID_PROOF_TYPE" insert_db_data_type="string" insert_data_source="lms_enquiry_table_obj.id_proof_type" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="ID_PROOF_VALUE" insert_db_data_type="string" insert_data_source="lms_enquiry_table_obj.id_proof_value" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="PAYMENT_STATUS" insert_db_data_type="string" insert_data_source="lms_enquiry_table_obj.payment_status" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="PAYMENT_MODE" insert_db_data_type="string" insert_data_source="lms_enquiry_table_obj.payment_mode" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="PAYMENT_ID" insert_db_data_type="string" insert_data_source="lms_enquiry_table_obj.payment_id" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="PAYMENT_AMOUNT" insert_db_data_type="string" insert_data_source="lms_enquiry_table_obj.payment_amount" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="MESSAGE" insert_db_data_type="string" insert_data_source="lms_enquiry_table_obj.message" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="CREATED_BY" insert_db_data_type="string" insert_data_source="lms_enquiry_table_obj.created_by" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="CREATED_DATE" insert_db_data_type="date" insert_data_source="lms_enquiry_table_obj.created_date" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="CREATED_TIME" insert_db_data_type="string" insert_data_source="lms_enquiry_table_obj.created_time" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="UPDATED_BY" insert_db_data_type="string" insert_data_source="lms_enquiry_table_obj.updated_by" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="UPDATED_DATE" insert_db_data_type="date" insert_data_source="lms_enquiry_table_obj.updated_date" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="UPDATED_TIME" insert_db_data_type="string" insert_data_source="lms_enquiry_table_obj.updated_time" insert_data_type="string"></insert_data>
				</insert_data_list>
			</db_insert_info>
		</exec_function>
		<exec_function function_name="lms_enquiry_comments_loop" function_type="start_loop" index_name="i" start_index="0" end_index="len(ReqBrokerReqObj.request_data.comments)" />
		<exec_function function_name="lms_enquiry_comments_validate_fld" function_type="validate_fld">
			<validate_info app_err_desc="pattern validation failed for input" app_corrective_action="check the inputs and provide proper values">
				<validate_fld fld_name="ReqBrokerReqObj.request_data.comments[i].comment" fld_data_type="string" fld_min_len="1" fld_max_len="500" fld_type="pattern" fld_pattern="^[a-zA-Z0-9.,|_* -]+$"></validate_fld>
			</validate_info>
		</exec_function>
		<exec_function function_name="edit_cms_customer_data_auth_db_single_read" function_type="db_single_read">
			<db_single_read_info app_err_desc="Customer data table error" app_corrective_action="Customer data table has not been configured properly">
				<table_name>LMS_ENQUIRY</table_name>
				<filter_info>
					<filter_data filter_name="PHONE_NUMBER" filter_id="ReqBrokerReqObj.request_data.phone_number" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_enquiry_table_record</store_obj_name>
			</db_single_read_info>
		</exec_function>
		<exec_function function_name="lms_enquiry_comments_table_copy_object" function_type="copy_object">
			<copy_object_info object_name="lms_enquiry_comments_table_obj">
				<object_data key="comment" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.comments[i].comment" data_type="string"></object_data>
				<object_data key="enquiry_id" data_source_type="reqbrokermap" data_source="lms_enquiry_table_record.id" data_type="string"></object_data>
				<object_data key="commented_by" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.email" data_type="string"></object_data>
				<object_data key="commented_at" data_source_type="key" data_source="get_time_stamp" data_type="string"></object_data>
				<object_data key="commented_by_email" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.email" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
		<exec_function function_name="lms_enquiry_auth_table_insert" function_type="db_insert">
			<db_insert_info app_err_desc="pattern validation failed for input" app_corrective_action="check the inputs and provide proper values">
				<table_name>LMS_ENQUIRY_COMMENT</table_name>
				<insert_data_list>
					<insert_data insert_data_name="ENQUIRY_ID" insert_db_data_type="string" insert_data_source="lms_enquiry_comments_table_obj.enquiry_id" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="COMMENT" insert_db_data_type="string" insert_data_source="lms_enquiry_comments_table_obj.comment" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="COMMENTED_BY" insert_db_data_type="string" insert_data_source="lms_enquiry_comments_table_obj.commented_by" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="COMMENTED_BY_EMAIL" insert_db_data_type="string" insert_data_source="lms_enquiry_comments_table_obj.commented_by_email" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="COMMENTED_AT" insert_db_data_type="string" insert_data_source="lms_enquiry_comments_table_obj.commented_at" insert_data_type="string"></insert_data>
				</insert_data_list>
			</db_insert_info>
		</exec_function>
		<exec_function function_name="lms_enquiry_comments_loop" function_type="end_loop" index_name="i" />
			<exec_function function_name="lms_enquiry_db_multi_read_course" function_type="db_multi_read">
			<db_multi_read_info app_err_desc="Data Already Exists or No Rows in Result Set" app_corrective_action="check the inputs and provide proper values">
				<table_name>LMS_ENQUIRY</table_name>
				<filter_info>
					<filter_data filter_name="EMAIL" filter_id="ReqBrokerReqObj.request_data.email" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_enquiry_table_list</store_obj_name>
			</db_multi_read_info>
			<err_info>
				<err_desc err_code="Duplicate" app_err_desc="Duplicate records found for the Email:{{.ReqBrokerReqObj.request_data.email}} provided in the input." app_corrective_action="Make sure each Email shows up only once." />
			</err_info>
		</exec_function>
		<exec_function function_name="lms_enquiry_loop" function_type="start_loop" index_name="i" start_index="0" end_index="len(lms_enquiry_table_list)" />
		
      <exec_function function_name="delete_lms_enquiry_course_table_delete" function_type="db_delete" function_condition="lms_enquiry_table_list!=nil">
			<db_delete_info>
				<table_name>LMS_ENQUIRY_COURSE</table_name>
				<delete_filter_list>
					<delete_filter delete_filter_name="ENQUIRY_ID" delete_filter_db_data_type="string" delete_filter_data_source="lms_enquiry_table_list[i].id" delete_filter_data_type="string" delete_filter_condition="AND"></delete_filter>
				</delete_filter_list>
			</db_delete_info>
		</exec_function>
		<exec_function function_name="lms_enquiry_loop" function_type="end_loop" index_name="i" />

		 <exec_function function_name="lms_course_loop" function_type="start_loop" index_name="i" start_index="0" end_index="len(ReqBrokerReqObj.request_data.courses)" />
		<exec_function function_name="lms_enquiry_courses_validate_fld" function_type="validate_fld">
			<validate_info app_err_desc="pattern validation failed for input" app_corrective_action="check the inputs and provide proper values">
				<validate_fld fld_name="ReqBrokerReqObj.request_data.courses[i]" fld_data_type="string" fld_min_len="1" fld_max_len="50" fld_type="pattern" fld_pattern="^[0-9]+$"></validate_fld>
			</validate_info>
		</exec_function>
		  <exec_function function_name="lms_enquiry_db_single_read" function_type="db_single_read">
			<db_single_read_info app_err_desc="Data Already Exists or No Rows in Result Set" app_corrective_action="check the inputs and provide proper values">
				<table_name>LMS_ENQUIRY</table_name>
				<filter_info>
						<filter_data filter_name="EMAIL" filter_id="ReqBrokerReqObj.request_data.email" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_enquiry_table_record</store_obj_name>
				</db_single_read_info>
		</exec_function>
		<exec_function function_name="lms_enquiry_courses_table_copy_object" function_type="copy_object" >
			<copy_object_info object_name="lms_enquiry_courses_table_obj">
				<object_data key="course_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.courses[i]" data_type="string"></object_data>
				<object_data key="enquiry_id" data_source_type="reqbrokermap" data_source="lms_enquiry_table_record.id" data_type="string"></object_data>
				<object_data key="created_by" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.email" data_type="string"></object_data>
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

		<exec_function function_name="lms_enquiry_courses_insert" function_type="db_insert">
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
		<exec_function function_name="lms_course_loop" function_type="end_loop" index_name="i" />
		  <exec_function function_name="otp_data_copy_object" function_type="copy_object">
				<copy_object_info object_name="otp_obj">
					<object_data key="otp_ref_data" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.email" data_type="string"></object_data>
                    <object_data key="otp_expiry_time" data_source_type="raw_value" data_source="300" data_type="string"></object_data>
				</copy_object_info>
			</exec_function>
			<exec_function function_name="check_otp_exist" function_type="db_multi_read">
				<db_multi_read_info>
					<schema_info module_name="SWITCH"/>
					<table_name>OTP_INFO</table_name>
					<filter_info>
						<filter_data filter_name="OTP_REF_DATA" filter_id="otp_obj.otp_ref_data" filter_data_type="string" filter_condition="AND"></filter_data>
					</filter_info>
					<store_obj_name>otp_info_list</store_obj_name>
				</db_multi_read_info>
			</exec_function>
	        <exec_function function_name="HandleOTPGen" function_type="call_method" function_condition="len(otp_info_list)==0"/>
	        <exec_function function_name="HandleOTPResend" function_type="call_method" function_condition="len(otp_info_list)!=0"/>
		<exec_function function_name="otp_validate_condition" function_type="validate_condition">
			<validate_condition_info app_err_desc="Failed to retrieve OTP from OTPEng server" app_corrective_action="OTP generation failed">
				<validate_condition_expression>otp_resp_obj.otp!=nil</validate_condition_expression>
			</validate_condition_info>
		</exec_function>
		  <exec_function function_name="full_name_copy_object" function_type="copy_object">
				<copy_object_info object_name="full_name_copy_data_obj">
					<object_data key="full_name" data_source_type="compute_str_expr" data_source="ReqBrokerReqObj.request_data.first_name+' '+ReqBrokerReqObj.request_data.last_name" data_type="string"></object_data>
				</copy_object_info>
			</exec_function>
        <exec_function function_name="am_email_template_info_db_single_read" function_type="db_single_read">
	<db_single_read_info>
		<table_name>LMS_EMAIL_TEMPLATE_INFO</table_name>
		<filter_info>
			<filter_data filter_name="EMAIL_TEMPLATE_ID"  filter_source_type="raw_value" filter_id="STUDENT" filter_data_type="string" filter_condition="AND"></filter_data>
		</filter_info>
		<store_obj_name>am_email_template_info_record</store_obj_name>
	</db_single_read_info>
</exec_function>
     <exec_function function_name="test_email_data_copy_object" function_type="copy_object">
	<copy_object_info object_name="email_obj">
		<object_data key="gateway_name" data_source_type="raw_value" data_source="DefaultGateway" data_type="string"></object_data>
		<object_data key="request_type" data_source_type="key" data_source="get_db_record_id" data_type="string"></object_data>
		<object_data key="request_num" data_source_type="key" data_source="get_db_record_id" data_type="string"></object_data>
		<object_data key="msg_id" data_source_type="key" data_source="get_db_record_id" data_type="string"></object_data>
		<object_data key="subject" data_source_type="reqbrokermap" data_source="am_email_template_info_record.subject" data_type="string"></object_data>
		<object_data key="body" data_source_type="reqbrokermap" data_source="am_email_template_info_record.body" data_type="string"></object_data>
	    <object_data key="from_address" data_source_type="raw_value" data_source="no-reply-alert@mindflix360.com" data_type="string"></object_data>
		</copy_object_info>
   </exec_function>
        <exec_function function_name="email_data_copy_object" function_type="copy_object">
     <copy_object_info object_name="email_obj">
       <object_data key="to_address_list" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.email" data_type="string"></object_data>
     </copy_object_info>
      </exec_function>
	 <exec_function function_name="HandleSendEmail" function_type="call_method"/> 
		<exec_function function_name="read_view_lms_enquiry_db_single_read" function_type="db_single_read">
			<db_single_read_info>
				<table_name>LMS_ENQUIRY</table_name>
				<filter_info>
					<filter_data filter_name="EMAIL" filter_id="ReqBrokerReqObj.request_data.email" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_enquiry_record</store_obj_name>
			</db_single_read_info>
		</exec_function>
		<exec_function function_name="view_lms_enquiry_id_db_single_read" function_type="db_multi_read">
			<db_multi_read_info>
				<table_name>LMS_ENQUIRY_COURSE</table_name>
				<filter_info>
					<filter_data filter_name="ENQUIRY_ID" filter_id="lms_enquiry_record.id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_enquiry_record.course_info</store_obj_name>
			</db_multi_read_info>
		</exec_function>
		<exec_function function_name="view_lms_enquiry_compose_response" function_type="compose_response">
			<response_info>
				<resp_obj obj_name="lms_enquiry_record" obj_type="object" data_source="lms_enquiry_record"></resp_obj>
			</response_info>
		</exec_function>
	</exec_group>
	<exec_group group_name="validate_lms_enquiry"  group_condition="ReqBrokerReqObj.header_data.channel=='web_admin_portal'" >
		<exec_function function_name="decode_token_function_type" function_type="decode_token">
			<decode_token_info token_value="ReqBrokerReqObj.header_data.auth_token" token_object="decode_token_obj" app_err_desc="Decode token failed!. Invalid access token." app_corrective_action="Please provide valid access token."/>
		</exec_function>
		 <exec_function function_name="role_validate_condition" function_type="validate_condition">
			<validate_condition_info app_err_desc="Invalid Role" app_corrective_action="Only ADMIN allowed to create User.">
				<validate_condition_expression>decode_token_obj.role=='ADMIN'||decode_token_obj.role=='SUPERADMIN'</validate_condition_expression>
			</validate_condition_info>
		</exec_function>
		<exec_function function_name="login_lms_token_guard_db_single_read" function_type="db_single_read">
			<db_single_read_info app_err_desc="Token guard data not found with Role:{{ReqBrokerReqObj.request_data.email}}" app_corrective_action="Please configure token guard for Role:{{decode_token_obj.role}}">
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
		<exec_function function_name="lms_enquiry_validate_fld" function_type="validate_fld">
			<validate_info app_err_desc="pattern validation failed for input" app_corrective_action="check the inputs and provide proper values">
				<validate_fld fld_name="ReqBrokerReqObj.request_data.first_name" fld_data_type="string" fld_min_len="1" fld_max_len="128" fld_type="pattern" fld_pattern="^[a-zA-Z0-9.,|_* -]+$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.last_name" fld_data_type="string" fld_min_len="1" fld_max_len="128" fld_type="pattern" fld_pattern="^[a-zA-Z0-9.,|_* -]+$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.phone_number" fld_data_type="string" fld_min_len="1" fld_max_len="128" fld_type="pattern" fld_pattern="^\+\d{1,3}-\d{10}$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.email" fld_data_type="string" fld_min_len="1" fld_max_len="128" fld_type="pattern" fld_pattern="^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,3}$"></validate_fld>
                <validate_fld fld_name="ReqBrokerReqObj.request_data.is_email_verified" fld_data_type="boolean" fld_min_len="1" fld_max_len="5" fld_type="pattern" fld_pattern="^(true|false)$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.education" fld_data_type="string" fld_min_len="0" fld_max_len="255" fld_type="pattern" fld_pattern="^[A-Za-z0-9 .,-]*$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.college" fld_data_type="string" fld_min_len="0" fld_max_len="500" fld_type="pattern" fld_pattern="^[A-Za-z0-9 .,-]*$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.department" fld_data_type="string" fld_min_len="0" fld_max_len="128" fld_type="pattern" fld_pattern="^[A-Za-z0-9 ._-]*$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.address" fld_data_type="string" fld_min_len="0" fld_max_len="500" fld_type="pattern" fld_pattern="^[a-zA-Z0-9.,|_* /-]+$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.id_proof_type" fld_data_type="string" fld_min_len="0" fld_max_len="50" fld_type="pattern" fld_pattern="^[A-Za-z0-9 _-]*$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.id_proof_value" fld_data_type="string" fld_min_len="0" fld_max_len="128" fld_type="pattern" fld_pattern="^[A-Za-z0-9 -]*$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.message" fld_data_type="string" fld_min_len="0" fld_max_len="500" fld_type="pattern" fld_pattern="^[a-zA-Z0-9.,|_* -]+$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.onboarded" fld_data_type="boolean" fld_min_len="1" fld_max_len="5" fld_type="pattern" fld_pattern="^(true|false)$" />
				<validate_fld fld_name="ReqBrokerReqObj.request_data.onboarding_status" fld_data_type="string" fld_min_len="0" fld_max_len="50" fld_type="pattern" fld_pattern="^[A-Za-z0-9 _-]*$" />
				<validate_fld fld_name="ReqBrokerReqObj.request_data.payment_status" fld_data_type="string" fld_min_len="0" fld_max_len="25" fld_type="pattern" fld_pattern="^[A-Za-z0-9 _-]*$" />
				<validate_fld fld_name="ReqBrokerReqObj.request_data.payment_mode" fld_data_type="string" fld_min_len="0" fld_max_len="25" fld_type="pattern" fld_pattern="^[A-Za-z0-9 _-]*$" />
				<validate_fld fld_name="ReqBrokerReqObj.request_data.payment_id" fld_data_type="string" fld_min_len="0" fld_max_len="255" fld_type="pattern" fld_pattern="^[A-Za-z0-9 _-]*$" />
				<validate_fld fld_name="ReqBrokerReqObj.request_data.payment_amount" fld_data_type="string" fld_min_len="0" fld_max_len="15" fld_type="pattern" fld_pattern="^[0-9]+(\.[0-9]{1,2})?$" />
				<validate_fld fld_name="ReqBrokerReqObj.request_data.message" fld_data_type="string" fld_min_len="0" fld_max_len="1000" fld_type="pattern" fld_pattern="^[a-zA-Z0-9.,|_* -]+$" />
			</validate_info>
		</exec_function>
		<exec_function function_name="lms_enquiry_db_single_read_validate_merchant_id" function_type="db_single_read">
			<db_single_read_info app_err_desc="Data Already Exists or No Rows in Result Set" app_corrective_action="check the inputs and provide proper values">
				<table_name>LMS_ENQUIRY</table_name>
				<filter_info>
						<filter_data filter_name="PHONE_NUMBER" filter_id="ReqBrokerReqObj.request_data.phone_number" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
					<result_info app_err_desc="The mobile_number:{{.ReqBrokerReqObj.request_data.phone_number}}  already exists." app_corrective_action="Please configure new mobile number  and proceed">
						<result_success>NoRows</result_success>
					</result_info>
				</db_single_read_info>
				<err_info>
					<err_desc err_code="Duplicate" app_err_desc="Duplicate records found for the mobile_number:{{.ReqBrokerReqObj.request_data.phone_number}}  provided in the input." app_corrective_action="Make sure each Mobile Number shows up only once." />   
				</err_info>
		</exec_function>
		<exec_function function_name="lms_enquiry_table_copy_object" function_type="copy_object">
			<copy_object_info object_name="lms_enquiry_table_obj">
				<object_data key="first_name" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.first_name" data_type="string"></object_data>
				<object_data key="last_name" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.last_name" data_type="string"></object_data>
				<object_data key="phone_number" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.phone_number" data_type="string"></object_data>
				<object_data key="email" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.email" data_type="string"></object_data>
                <object_data key="is_email_verified" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.is_email_verified" data_type="Boolean"></object_data>
				<object_data key="email_verified_at" data_source_type="raw_value" data_source="1970-01-01 00:00:00.000" data_type="string"></object_data>
				<object_data key="onboarded" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.onboarded" data_type="Boolean"></object_data>
				<object_data key="onboarding_status" data_source_type="raw_value" data_source="NOT STARTED" data_type="string"></object_data>
				<object_data key="education" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.education" data_type="string"></object_data>
				<object_data key="college" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.college" data_type="string"></object_data>
				<object_data key="department" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.department" data_type="string"></object_data>
				<object_data key="address" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.address" data_type="string"></object_data>
				<object_data key="id_proof_type" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.id_proof_type" data_type="string"></object_data>
				<object_data key="id_proof_value" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.id_proof_value" data_type="string"></object_data>
				<object_data key="payment_status" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.payment_status" data_type="string"></object_data>
				<object_data key="payment_mode" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.payment_mode" data_type="string"></object_data>
				<object_data key="payment_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.payment_id" data_type="string"></object_data>
				<object_data key="payment_amount" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.payment_amount" data_type="string"></object_data>
				<object_data key="payment_date" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.payment_date" data_type="string"></object_data>
				<object_data key="message" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.message" data_type="string"></object_data>
				<object_data key="created_by" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.email" data_type="string"></object_data>
				<object_data key="created_date" data_source_type="key" data_source="get_date" data_type="string"></object_data>
				<object_data key="created_time" data_source_type="key" data_source="get_time" data_type="string"></object_data>
				<object_data key="updated_by" data_source_type="key" data_source="get_na" data_type="string"></object_data>
				<object_data key="updated_date" data_source_type="key" data_source="get_date" data_type="string"></object_data>
				<object_data key="updated_time" data_source_type="key" data_source="get_time" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
		<exec_function function_name="lms_enquiry_table_insert" function_type="db_insert">
			<db_insert_info>
				<table_name>LMS_ENQUIRY</table_name>
				<insert_data_list>
					<insert_data insert_data_name="FIRST_NAME" insert_db_data_type="string" insert_data_source="lms_enquiry_table_obj.first_name" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="LAST_NAME" insert_db_data_type="string" insert_data_source="lms_enquiry_table_obj.last_name" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="EMAIL" insert_db_data_type="string" insert_data_source="lms_enquiry_table_obj.email" insert_data_type="string"></insert_data>
                	<insert_data insert_data_name="IS_EMAIL_VERIFIED" insert_db_data_type="boolean" insert_data_source="ReqBrokerReqObj.request_data.is_email_verified" insert_data_type="Boolean"></insert_data>
                    <insert_data insert_data_name="EMAIL_VERIFIED_AT" insert_db_data_type="string" insert_data_source="lms_enquiry_table_obj.email_verified_at" insert_data_type="string"></insert_data>
                    <insert_data insert_data_name="PHONE_NUMBER" insert_db_data_type="string" insert_data_source="lms_enquiry_table_obj.phone_number" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="ONBOARDED" insert_db_data_type="boolean" insert_data_source="ReqBrokerReqObj.request_data.onboarded" insert_data_type="Boolean"></insert_data>
					<insert_data insert_data_name="ONBOARDING_STATUS" insert_db_data_type="string" insert_data_source="lms_enquiry_table_obj.onboarding_status" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="EDUCATION" insert_db_data_type="string" insert_data_source="lms_enquiry_table_obj.education" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="COLLEGE" insert_db_data_type="string" insert_data_source="lms_enquiry_table_obj.college" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="DEPARTMENT" insert_db_data_type="string" insert_data_source="lms_enquiry_table_obj.department" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="ADDRESS" insert_db_data_type="string" insert_data_source="lms_enquiry_table_obj.address" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="ID_PROOF_TYPE" insert_db_data_type="string" insert_data_source="lms_enquiry_table_obj.id_proof_type" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="ID_PROOF_VALUE" insert_db_data_type="string" insert_data_source="lms_enquiry_table_obj.id_proof_value" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="PAYMENT_STATUS" insert_db_data_type="string" insert_data_source="lms_enquiry_table_obj.payment_status" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="PAYMENT_MODE" insert_db_data_type="string" insert_data_source="lms_enquiry_table_obj.payment_mode" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="PAYMENT_ID" insert_db_data_type="string" insert_data_source="lms_enquiry_table_obj.payment_id" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="PAYMENT_AMOUNT" insert_db_data_type="string" insert_data_source="lms_enquiry_table_obj.payment_amount" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="MESSAGE" insert_db_data_type="string" insert_data_source="lms_enquiry_table_obj.message" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="CREATED_BY" insert_db_data_type="string" insert_data_source="lms_enquiry_table_obj.created_by" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="CREATED_DATE" insert_db_data_type="date" insert_data_source="lms_enquiry_table_obj.created_date" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="CREATED_TIME" insert_db_data_type="string" insert_data_source="lms_enquiry_table_obj.created_time" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="UPDATED_BY" insert_db_data_type="string" insert_data_source="lms_enquiry_table_obj.updated_by" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="UPDATED_DATE" insert_db_data_type="date" insert_data_source="lms_enquiry_table_obj.updated_date" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="UPDATED_TIME" insert_db_data_type="string" insert_data_source="lms_enquiry_table_obj.updated_time" insert_data_type="string"></insert_data>
				</insert_data_list>
			</db_insert_info>
		</exec_function>
		<exec_function function_name="lms_enquiry_comments_loop" function_type="start_loop" index_name="i" start_index="0" end_index="len(ReqBrokerReqObj.request_data.comments)" />
		<exec_function function_name="lms_enquiry_comments_validate_fld" function_type="validate_fld">
			<validate_info app_err_desc="pattern validation failed for input" app_corrective_action="check the inputs and provide proper values">
				<validate_fld fld_name="ReqBrokerReqObj.request_data.comments[i].comment" fld_data_type="string" fld_min_len="1" fld_max_len="500" fld_type="pattern" fld_pattern="^[a-zA-Z0-9.,|_* -]+$"></validate_fld>
			</validate_info>
		</exec_function>
		<exec_function function_name="edit_cms_customer_data_auth_db_single_read" function_type="db_single_read">
			<db_single_read_info app_err_desc="Customer data table error" app_corrective_action="Customer data table has not been configured properly">
				<table_name>LMS_ENQUIRY</table_name>
				<filter_info>
					<filter_data filter_name="PHONE_NUMBER" filter_id="ReqBrokerReqObj.request_data.phone_number" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_enquiry_table_record</store_obj_name>
			</db_single_read_info>
		</exec_function>
		<exec_function function_name="lms_enquiry_comments_table_copy_object" function_type="copy_object">
			<copy_object_info object_name="lms_enquiry_comments_table_obj">
				<object_data key="comment" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.comments[i].comment" data_type="string"></object_data>
				<object_data key="enquiry_id" data_source_type="reqbrokermap" data_source="lms_enquiry_table_record.id" data_type="string"></object_data>
				<object_data key="commented_by" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.email" data_type="string"></object_data>
				<object_data key="commented_at" data_source_type="key" data_source="get_time_stamp" data_type="string"></object_data>
				<object_data key="commented_by_email" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.email" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
		<exec_function function_name="lms_enquiry_auth_table_insert" function_type="db_insert">
			<db_insert_info app_err_desc="pattern validation failed for input" app_corrective_action="check the inputs and provide proper values">
				<table_name>LMS_ENQUIRY_COMMENT</table_name>
				<insert_data_list>
					<insert_data insert_data_name="ENQUIRY_ID" insert_db_data_type="string" insert_data_source="lms_enquiry_comments_table_obj.enquiry_id" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="COMMENT" insert_db_data_type="string" insert_data_source="lms_enquiry_comments_table_obj.comment" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="COMMENTED_BY" insert_db_data_type="string" insert_data_source="lms_enquiry_comments_table_obj.commented_by" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="COMMENTED_BY_EMAIL" insert_db_data_type="string" insert_data_source="lms_enquiry_comments_table_obj.commented_by_email" insert_data_type="string"></insert_data>
					<insert_data insert_data_name="COMMENTED_AT" insert_db_data_type="string" insert_data_source="lms_enquiry_comments_table_obj.commented_at" insert_data_type="string"></insert_data>
				</insert_data_list>
			</db_insert_info>
		</exec_function>
		<exec_function function_name="lms_enquiry_comments_loop" function_type="end_loop" index_name="i" />
		<exec_function function_name="lms_enquiry_courses_loop" function_type="start_loop" index_name="i" start_index="0" end_index="len(ReqBrokerReqObj.request_data.courses)" />
		<exec_function function_name="lms_enquiry_courses_validate_fld" function_type="validate_fld">
			<validate_info app_err_desc="pattern validation failed for input" app_corrective_action="check the inputs and provide proper values">
				<validate_fld fld_name="ReqBrokerReqObj.request_data.courses[i]" fld_data_type="string" fld_min_len="1" fld_max_len="50" fld_type="pattern" fld_pattern="^[0-9]+$"></validate_fld>
			</validate_info>
		</exec_function>
		<exec_function function_name="edit_cms_customer_data_auth_db_single_read" function_type="db_single_read">
			<db_single_read_info app_err_desc="Customer data table error" app_corrective_action="Customer data table has not been configured properly">
				<table_name>LMS_ENQUIRY</table_name>
				<filter_info>
					<filter_data filter_name="PHONE_NUMBER" filter_id="ReqBrokerReqObj.request_data.phone_number" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_enquiry_table_record</store_obj_name>
			</db_single_read_info>
		</exec_function>
		<exec_function function_name="lms_enquiry_courses_table_copy_object" function_type="copy_object">
			<copy_object_info object_name="lms_enquiry_courses_table_obj">
				<object_data key="course_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.courses[i]" data_type="string"></object_data>
				<object_data key="enquiry_id" data_source_type="reqbrokermap" data_source="lms_enquiry_table_record.id" data_type="string"></object_data>
				<object_data key="created_by" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.email" data_type="string"></object_data>
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
		<exec_function function_name="lms_enquiry_courses_insert" function_type="db_insert">
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
		<exec_function function_name="lms_enquiry_courses_loop" function_type="end_loop" index_name="i" />
		<exec_function function_name="read_lms_enquiry_db_multi_read" function_type="db_multi_read">
			<db_multi_read_info app_err_desc="Data Already Exists or No Rows in Result Set" app_corrective_action="check the inputs and provide proper values">
				<table_name>LMS_ENQUIRY</table_name>
				<store_obj_name>lms_enquiry_list</store_obj_name>
			</db_multi_read_info>
		</exec_function>
		<exec_function function_name="lms_enquiry_course_loop" function_type="start_loop" index_name="i" start_index="0" end_index="len(lms_enquiry_list)" />
		<exec_function function_name="read_view_lms_enquiry_db_multi_read" function_type="db_multi_read">
			<db_multi_read_info>
				<table_name>LMS_ENQUIRY_COURSE</table_name>
				<filter_info>
					<filter_data filter_name="ENQUIRY_ID" filter_id="lms_enquiry_list[i].id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_enquiry_list[i].courses</store_obj_name>
			</db_multi_read_info>
		</exec_function>
		<exec_function function_name="read_view_lms_enquiry_comment_db_multi_read" function_type="db_multi_read">
			<db_multi_read_info>
				<table_name>LMS_ENQUIRY_COMMENT</table_name>
				<filter_info>
					<filter_data filter_name="ENQUIRY_ID" filter_id="lms_enquiry_list[i].id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_enquiry_list[i].comments</store_obj_name>
			</db_multi_read_info>
		</exec_function>
		<exec_function function_name="lms_enquiry_course_loop" function_type="end_loop" index_name="i" />
		<exec_function function_name="insert_lms_enquiry_compose_response" function_type="compose_response">
			<response_info>
				<resp_obj obj_name="lms_enquiry_list" obj_type="objectarray" data_source="lms_enquiry_list"></resp_obj>
			</response_info>
		</exec_function>
	</exec_group>
</sem_lexical_parser>
	`
	return lexicalParserStr
}
