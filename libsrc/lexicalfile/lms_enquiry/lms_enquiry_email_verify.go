package lexicalfile_lms_enquiry

func GetLMSEnquiryEmailVerifyServiceXML() interface{} {
	lexicalParserStr := `
	<sem_lexical_parser microservice="lms_enquiry_email_verify">
		<exec_group group_name="lms_enquiry_email_verify">
		 <exec_function function_name="lms_user_data_validate_fld" function_type="validate_fld">
			<validate_info app_err_desc="Validation failed for user data" app_corrective_action="Please validate the data">
				<validate_fld fld_name="ReqBrokerReqObj.request_data.email" fld_data_type="string" fld_min_len="1" fld_max_len="255" fld_type="pattern" fld_pattern="^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,3}$" app_err_desc="Invalid Email" app_corrective_action="Only send valid email format."></validate_fld>
                 <validate_fld fld_name="ReqBrokerReqObj.request_data.otp" fld_data_type="string" fld_min_len="1" fld_max_len="8" fld_type="pattern" fld_pattern="^[0-9]+$" app_err_desc="Invalid OTP" app_corrective_action="Only send valid OTP."></validate_fld>
                <validate_fld fld_name="ReqBrokerReqObj.request_data.is_email_verified" fld_data_type="boolean" fld_min_len="1" fld_max_len="5" fld_type="pattern" fld_pattern="^(true|false)$"></validate_fld>
			</validate_info>
		</exec_function>
        <exec_function function_name="verify_otp_service_db_single_read_1" function_type="db_single_read">
				<db_single_read_info app_err_desc="OTP Ref Data Not Found" app_corrective_action="Ensure that OTP Reference Data is Correct">
					<schema_info module_name='SWITCH'/>
                    <table_name>OTP_INFO</table_name>  
					<filter_info>
						<filter_data filter_name="OTP_REF_DATA" filter_id="ReqBrokerReqObj.request_data.email" filter_data_type="string" filter_condition="AND"></filter_data>
					</filter_info>
					<store_obj_name>otp_info_table_record</store_obj_name>
				</db_single_read_info>
				<err_info>
					<err_desc err_code="NoRows" app_err_desc="The entered OTP has expired " app_corrective_action="Click on Resend OTP to receive a fresh OTP and complete the verification." />
					<err_desc err_code="Duplicate" app_err_desc="Duplicate records found for the otp ref data:{{.ReqBrokerReqObj.request_data.email}}" app_corrective_action="Please ensure that otp ref data unique for customer." />
				</err_info>
			</exec_function>
			<exec_function function_name="verify_otp_data_copy_object" function_type="copy_object">
				<copy_object_info object_name="otp_obj">
					<object_data key="otp_ref_data" data_source_type="reqbrokermap" data_source="otp_info_table_record.otp_ref_data" data_type="string"></object_data>
					<object_data key="otp" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.otp" data_type="string"></object_data>
				</copy_object_info>
			</exec_function>
	         <exec_function function_name="HandleOTPVerify" function_type="call_method" />
              <exec_function function_name="otp_validate_condition" function_type="validate_condition">
			<validate_condition_info app_err_desc="Invalid OTP" app_corrective_action="Make Sure Entered OTP is correct">
				<validate_condition_expression>otp_resp_obj.otp_status=='1'</validate_condition_expression>
			</validate_condition_info>
		</exec_function>
          <exec_function function_name="edit_lms_enquiry_personal_info_auth_db_single_read" function_type="db_single_read">
			<db_single_read_info app_err_desc="Data Already Exists or No Rows in Result Set" app_corrective_action="check the inputs and provide proper values">
				<table_name>LMS_ENQUIRY</table_name>
				<filter_info>
					<filter_data filter_name="EMAIL" filter_id="ReqBrokerReqObj.request_data.email" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_enquiry_table_record</store_obj_name>
			</db_single_read_info>
		</exec_function>  
         
		 <exec_function function_name="full_name_copy_object" function_type="copy_object">
				<copy_object_info object_name="full_name_copy_data_obj">
					<object_data key="full_name" data_source_type="compute_str_expr" data_source="lms_enquiry_table_record.first_name+' '+lms_enquiry_table_record.last_name" data_type="string"></object_data>
				</copy_object_info>
			</exec_function>
            <exec_function function_name="am_email_template_info_db_single_read" function_type="db_single_read">
	<db_single_read_info>
		<table_name>LMS_EMAIL_TEMPLATE_INFO</table_name>
		<filter_info>
			<filter_data filter_name="EMAIL_TEMPLATE_ID"  filter_source_type="raw_value" filter_id="ENQUIRY" filter_data_type="string" filter_condition="AND"></filter_data>
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
       <object_data key="cc_address_list" data_source_type="raw_value" data_source="academy@mindflix360.com" data_type="string"></object_data>
	   </copy_object_info>
      </exec_function>
        <exec_function function_name="HandleSendEmail" function_type="call_method"/> 
		<exec_function function_name="edit_lms_enquiry_personal_info_table_copy_object" function_type="copy_object">
			<copy_object_info object_name="lms_enquiry_table_obj">
                <object_data key="is_email_verified" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.is_email_verified" data_type="Boolean"></object_data>
				<object_data key="email_verified_at" data_source_type="key" data_source="get_time_stamp" data_type="string"></object_data>
				<object_data key="updated_by" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.email" data_type="string"></object_data>
				<object_data key="updated_date" data_source_type="key" data_source="get_date" data_type="string"></object_data>
				<object_data key="updated_time" data_source_type="key" data_source="get_time" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
		<exec_function function_name="lms_enquiry_personal_info_db_update" function_type="db_update">
			<db_update_info>
				<table_name>LMS_ENQUIRY</table_name>
				<update_data_list>
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

        <exec_function function_name="otp_verify_copy_object" function_type="copy_object">
				<copy_object_info object_name="otp_verify_response_obj">
				<object_data key="resp_msg" data_source_type="raw_value" data_source="Email verified successfully!" data_type="string"></object_data>
			</copy_object_info>
			</exec_function>

		<exec_function function_name="login_lms_user_data_compose_response" function_type="compose_response">
			<response_info>
				<resp_obj obj_name="otp_verify_response_obj" obj_type="object" data_source="otp_verify_response_obj"></resp_obj>
			</response_info>
		</exec_function>
		</exec_group>
	</sem_lexical_parser>
	`
	return lexicalParserStr
}
