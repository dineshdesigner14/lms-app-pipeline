package lexicalfile_lms_user_auth_data

func GetLMSUserEmailVerifyServiceXML() interface{} {
	lexicalParserStr := `
	<sem_lexical_parser microservice="lms_user_email_verify">
		<exec_group group_name="lms_user_email_verify">
		 <exec_function function_name="lms_user_data_validate_fld" function_type="validate_fld">
			<validate_info app_err_desc="Validation failed for user data" app_corrective_action="Please validate the data">
				<validate_fld fld_name="ReqBrokerReqObj.request_data.first_name" fld_data_type="string" fld_min_len="1" fld_max_len="128" fld_type="pattern" fld_pattern="^[a-zA-Z.,|_* -]+$" app_err_desc="Invalid first name" app_corrective_action="Only letters and spaces allowed."></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.last_name" fld_data_type="string" fld_min_len="1" fld_max_len="128" fld_type="pattern" fld_pattern="^[a-zA-Z.,|_* -]+$" app_err_desc="Invalid last name" app_corrective_action="Only letters and spaces allowed."></validate_fld>
                <validate_fld fld_name="ReqBrokerReqObj.request_data.email" fld_data_type="string" fld_min_len="1" fld_max_len="255" fld_type="pattern" fld_pattern="^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,3}$" app_err_desc="Invalid UserID" app_corrective_action="Only send valid email format."></validate_fld>
			</validate_info>
		</exec_function>
        <exec_function function_name="otp_resend_data_copy_object" function_type="copy_object">
				<copy_object_info object_name="otp_obj">
					<object_data key="otp_ref_data" data_source_type="compute_str_expr" data_source="ReqBrokerReqObj.request_data.phn_no_cc+ReqBrokerReqObj.request_data.mobile_num" data_type="string"></object_data>
				</copy_object_info>
			</exec_function>
	        <exec_function function_name="HandleOTPResend" function_type="call_method" />
		<exec_function function_name="otp_validate_condition" function_type="validate_condition">
			<validate_condition_info app_err_desc="Failed to retrieve OTP from OTPEng server" app_corrective_action="OTP generation failed">
				<validate_condition_expression>otp_resp_obj.otp!=nil</validate_condition_expression>
			</validate_condition_info>
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
        <exec_function function_name="email_verify_copy_object" function_type="copy_object">
				<copy_object_info object_name="email_verify_response_obj">
				<object_data key="resp_msg" data_source_type="raw_value" data_source="Verification otp sent to emailID successfully!" data_type="string"></object_data>
			</copy_object_info>
			</exec_function>
		<exec_function function_name="login_lms_user_data_compose_response" function_type="compose_response">
			<response_info>
				<resp_obj obj_name="email_verify_response_obj" obj_type="object" data_source="email_verify_response_obj"></resp_obj>
			</response_info>
		</exec_function>
		</exec_group>
	</sem_lexical_parser>
	`
	return lexicalParserStr
}
