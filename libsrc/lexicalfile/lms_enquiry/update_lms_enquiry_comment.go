package lexicalfile_lms_enquiry

func GetUpdateLMSEnquiryCommentServiceXML() interface{} {
	lexicalParserStr := `
	<sem_lexical_parser microservice="update_lms_enquiry_comment">
	<exec_group group_name="update_lms_enquiry_comment_data"  group_condition="ReqBrokerReqObj.header_data.channel=='web_admin_portal'">
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
		<exec_function function_name="lms_enquiry_comment_validate_fld" function_type="validate_fld">
			<validate_info app_err_desc="pattern validation failed for input" app_corrective_action="check the inputs and provide proper values">
				<validate_fld fld_name="ReqBrokerReqObj.request_data.enquiry_id" fld_data_type="string" fld_min_len="1" fld_max_len="128" fld_type="pattern" fld_pattern="^(\d+)$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.db_operation" fld_data_type="string" fld_min_len="1" fld_max_len="128" fld_type="pattern" fld_pattern="^(INSERT|EDIT|DELETE)$"></validate_fld>
			</validate_info>
		</exec_function>
		<exec_function function_name="lms_enquiry_comment_validate_fld" function_type="validate_fld" function_condition="ReqBrokerReqObj.request_data.db_operation!='INSERT'">
			<validate_info app_err_desc="Comment ID not found" app_corrective_action="Please include the comment ID in the request">
				<validate_fld fld_name="ReqBrokerReqObj.request_data.comment_id" fld_data_type="string" fld_min_len="1" fld_max_len="128" fld_type="pattern" fld_pattern="^(NA|\d+)$"></validate_fld>
			</validate_info>
		</exec_function>
		<exec_function function_name="lms_enquiry_comment_validate_fld" function_type="validate_fld" function_condition="ReqBrokerReqObj.request_data.db_operation!='DELETE'">
			<validate_info app_err_desc="Comment ID not found" app_corrective_action="Please include the comment ID in the request">
				<validate_fld fld_name="ReqBrokerReqObj.request_data.comment" fld_data_type="string" fld_min_len="1" fld_max_len="128" fld_type="pattern" fld_pattern="^[A-Za-z0-9 .,-]*$"></validate_fld>
			</validate_info>
		</exec_function>
		<exec_function function_name="insert_lms_enquiry_comments_table_copy_object" function_type="copy_object">
			<copy_object_info object_name="lms_enquiry_comments_table_obj">
				<object_data key="comment" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.comment" data_type="string"></object_data>
				<object_data key="enquiry_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.enquiry_id" data_type="string"></object_data>
				<object_data key="commented_by" data_source_type="reqbrokermap" data_source="decode_token_obj.email" data_type="string"></object_data>
				<object_data key="commented_at" data_source_type="key" data_source="get_time_stamp" data_type="string"></object_data>
				<object_data key="commented_by_email" data_source_type="reqbrokermap" data_source="decode_token_obj.email" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
		<exec_function function_name="insert_lms_enquiry_auth_table_insert" function_type="db_insert" function_condition="ReqBrokerReqObj.request_data.db_operation=='INSERT'">
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
		<exec_function function_name="lms_enquiry_comment_db_update" function_type="db_update" function_condition="ReqBrokerReqObj.request_data.db_operation=='EDIT'">
			<db_update_info>
				<table_name>LMS_ENQUIRY_COMMENT</table_name>
				<update_data_list>
					<update_data update_data_name="COMMENT" update_db_data_type="string" update_data_source="lms_enquiry_comments_table_obj.comment" update_data_type="string"></update_data>
					<update_data update_data_name="COMMENTED_BY" update_db_data_type="string" update_data_source="lms_enquiry_comments_table_obj.commented_by" update_data_type="string"></update_data>
					<update_data update_data_name="COMMENTED_BY_EMAIL" update_db_data_type="string" update_data_source="lms_enquiry_comments_table_obj.commented_by_email" update_data_type="string"></update_data>
					<update_data update_data_name="COMMENTED_AT" update_db_data_type="string" update_data_source="lms_enquiry_comments_table_obj.commented_at" update_data_type="string"></update_data>
				</update_data_list>
				<update_filter_list>
					<update_filter update_filter_name="ID" update_filter_db_data_type="string" update_filter_data_source="ReqBrokerReqObj.request_data.comment_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
					<update_filter update_filter_name="ENQUIRY_ID" update_filter_db_data_type="string" update_filter_data_source="lms_enquiry_comments_table_obj.enquiry_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
				</update_filter_list>
			</db_update_info>
		</exec_function>
		<exec_function function_name="delete_lms_enquiry_comment_table_delete" function_type="db_delete"  function_condition="ReqBrokerReqObj.request_data.db_operation=='DELETE'">
			<db_delete_info>
				<table_name>LMS_ENQUIRY_COMMENT</table_name>
				<delete_filter_list>
					<delete_filter delete_filter_name="ID" delete_filter_db_data_type="string" delete_filter_data_source="ReqBrokerReqObj.request_data.comment_id" delete_filter_data_type="string" delete_filter_condition="AND"></delete_filter>
					<delete_filter delete_filter_name="ENQUIRY_ID" delete_filter_db_data_type="string" delete_filter_data_source="ReqBrokerReqObj.request_data.enquiry_id" delete_filter_data_type="string" delete_filter_condition="AND"></delete_filter>
				</delete_filter_list>
			</db_delete_info>
		</exec_function>
		<exec_function function_name="lms_enquiry_comment_table_db_multi_read" function_type="db_multi_read">
			<db_multi_read_info app_err_desc="Data Already Exists or No Rows in Result Set" app_corrective_action="check the inputs and provide proper values">
				<table_name>LMS_ENQUIRY_COMMENT</table_name>
				<filter_info>
					<filter_data filter_name="ENQUIRY_ID" filter_id="ReqBrokerReqObj.request_data.enquiry_id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_enquiry_comment_table_list</store_obj_name>
			</db_multi_read_info>
		</exec_function>
		<exec_function function_name="insert_lms_enquiry_compose_response" function_type="compose_response">
			<response_info>
				<resp_obj obj_name="lms_enquiry_comment_table_list" obj_type="objectarray" data_source="lms_enquiry_comment_table_list"></resp_obj>
			</response_info>
		</exec_function>
	</exec_group>
</sem_lexical_parser>
	`
	return lexicalParserStr
}
