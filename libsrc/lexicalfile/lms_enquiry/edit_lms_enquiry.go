package lexicalfile_lms_enquiry

func GetEditLMSEnquiryServiceXML() interface{} {
	lexicalParserStr := `
	<sem_lexical_parser microservice="edit_lms_enquiry_data">
	<exec_group group_name="edit_lms_enquiry_personal_info" group_condition="ReqBrokerReqObj.header_data.channel=='web_client_portal'">
		<exec_function function_name="edit_lms_enquiry_personal_info_validate_fld" function_type="validate_fld">
			<validate_info app_err_desc="pattern validation failed for input" app_corrective_action="check the inputs and provide proper values">
				<validate_fld fld_name="ReqBrokerReqObj.request_data.id" fld_data_type="string" fld_min_len="1" fld_max_len="128" fld_type="pattern" fld_pattern="^[a-zA-Z0-9.,|_* -]+$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.first_name" fld_data_type="string" fld_min_len="1" fld_max_len="128" fld_type="pattern" fld_pattern="^[a-zA-Z0-9.,|_* -]+$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.last_name" fld_data_type="string" fld_min_len="1" fld_max_len="128" fld_type="pattern" fld_pattern="^[a-zA-Z0-9.,|_* -]+$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.phone_number" fld_data_type="string" fld_min_len="1" fld_max_len="32" fld_type="pattern" fld_pattern="^\+\d{1,3}-\d{10}$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.email" fld_data_type="string" fld_min_len="1" fld_max_len="128" fld_type="pattern" fld_pattern="^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,3}$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.education" fld_data_type="string" fld_min_len="0" fld_max_len="255" fld_type="pattern" fld_pattern="^[A-Za-z0-9 .,-]*$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.college" fld_data_type="string" fld_min_len="0" fld_max_len="500" fld_type="pattern" fld_pattern="^[A-Za-z0-9 .,-]*$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.department" fld_data_type="string" fld_min_len="0" fld_max_len="128" fld_type="pattern" fld_pattern="^[A-Za-z0-9 ._-]*$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.address" fld_data_type="string" fld_min_len="0" fld_max_len="500" fld_type="pattern" fld_pattern="^[a-zA-Z0-9.,|_* /-]+$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.id_proof_type" fld_data_type="string" fld_min_len="0" fld_max_len="50" fld_type="pattern" fld_pattern="^[A-Za-z0-9 _-]*$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.id_proof_value" fld_data_type="string" fld_min_len="0" fld_max_len="128" fld_type="pattern" fld_pattern="^[A-Za-z0-9 -]*$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.message" fld_data_type="string" fld_min_len="0" fld_max_len="500" fld_type="pattern" fld_pattern="^[a-zA-Z0-9.,|_* -]+$"></validate_fld>
                <validate_fld fld_name="ReqBrokerReqObj.request_data.is_email_verified" fld_data_type="boolean" fld_min_len="1" fld_max_len="5" fld_type="pattern" fld_pattern="^(true|false)$"></validate_fld>
			</validate_info>
		</exec_function>
		<exec_function function_name="edit_lms_enquiry_personal_info_auth_db_single_read" function_type="db_single_read">
			<db_single_read_info app_err_desc="Data Already Exists or No Rows in Result Set" app_corrective_action="check the inputs and provide proper values">
				<table_name>LMS_ENQUIRY</table_name>
				<filter_info>
					<filter_data filter_name="PHONE_NUMBER" filter_id="ReqBrokerReqObj.request_data.phone_number" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_enquiry_table_record</store_obj_name>
			</db_single_read_info>
		</exec_function>
		<exec_function function_name="edit_lms_enquiry_personal_info_table_copy_object" function_type="copy_object">
			<copy_object_info object_name="lms_enquiry_table_obj">
				<object_data key="first_name" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.first_name" data_type="string"></object_data>
				<object_data key="last_name" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.last_name" data_type="string"></object_data>
				<object_data key="phone_number" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.phone_number" data_type="string"></object_data>
				<object_data key="email" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.email" data_type="string"></object_data>
				 <object_data key="is_email_verified" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.is_email_verified" data_type="Boolean"></object_data>
				<object_data key="email_verified_at" data_source_type="key" data_source="get_time_stamp" data_type="string"></object_data>
               <object_data key="onboarded" data_source_type="reqbrokermap" data_source="lms_enquiry_table_record.onboarded" data_type="Boolean"></object_data>
				<object_data key="onboarding_status" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.onboarding_status" data_type="string"></object_data>
				<object_data key="education" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.education" data_type="string"></object_data>
				<object_data key="college" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.college" data_type="string"></object_data>
				<object_data key="department" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.department" data_type="string"></object_data>
				<object_data key="address" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.address" data_type="string"></object_data>
				<object_data key="id_proof_type" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.id_proof_type" data_type="string"></object_data>
				<object_data key="id_proof_value" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.id_proof_value" data_type="string"></object_data>
				<object_data key="payment_status" data_source_type="reqbrokermap" data_source="lms_enquiry_table_record.payment_status" data_type="string"></object_data>
				<object_data key="payment_mode" data_source_type="reqbrokermap" data_source="lms_enquiry_table_record.payment_mode" data_type="string"></object_data>
				<object_data key="payment_id" data_source_type="reqbrokermap" data_source="lms_enquiry_table_record.payment_id" data_type="string"></object_data>
				<object_data key="payment_amount" data_source_type="reqbrokermap" data_source="lms_enquiry_table_record.payment_amount" data_type="string"></object_data>
				<object_data key="payment_date" data_source_type="reqbrokermap" data_source="lms_enquiry_table_record.payment_date" data_type="string"></object_data>
				<object_data key="message" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.message" data_type="string"></object_data>
				<object_data key="updated_by" data_source_type="reqbrokermap" data_source="decode_token_obj.email" data_type="string"></object_data>
				<object_data key="updated_date" data_source_type="key" data_source="get_date" data_type="string"></object_data>
				<object_data key="updated_time" data_source_type="key" data_source="get_time" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
		
		<exec_function function_name="lms_enquiry_personal_info_db_update" function_type="db_update">
			<db_update_info>
				<table_name>LMS_ENQUIRY</table_name>
				<update_data_list>
					<update_data update_data_name="FIRST_NAME" update_db_data_type="string" update_data_source="lms_enquiry_table_obj.first_name" update_data_type="string"></update_data>
					<update_data update_data_name="LAST_NAME" update_db_data_type="string" update_data_source="lms_enquiry_table_obj.last_name" update_data_type="string"></update_data>
					<update_data update_data_name="EMAIL" update_db_data_type="string" update_data_source="lms_enquiry_table_obj.email" update_data_type="string"></update_data>
                    <update_data update_data_name="IS_EMAIL_VERIFIED" update_db_data_type="boolean" update_data_source="lms_enquiry_table_obj.is_email_verified" update_data_type="boolean"></update_data>
                    <update_data update_data_name="EMAIL_VERIFIED_AT" update_db_data_type="string" update_data_source="lms_enquiry_table_obj.email_verified_at" update_data_type="string"></update_data>
					<update_data update_data_name="EDUCATION" update_db_data_type="string" update_data_source="lms_enquiry_table_obj.education" update_data_type="string"></update_data>
					<update_data update_data_name="COLLEGE" update_db_data_type="string" update_data_source="lms_enquiry_table_obj.college" update_data_type="string"></update_data>
					<update_data update_data_name="DEPARTMENT" update_db_data_type="string" update_data_source="lms_enquiry_table_obj.department" update_data_type="string"></update_data>
					<update_data update_data_name="ADDRESS" update_db_data_type="string" update_data_source="lms_enquiry_table_obj.address" update_data_type="string"></update_data>
					<update_data update_data_name="ID_PROOF_TYPE" update_db_data_type="string" update_data_source="lms_enquiry_table_obj.id_proof_type" update_data_type="string"></update_data>
					<update_data update_data_name="ID_PROOF_VALUE" update_db_data_type="string" update_data_source="lms_enquiry_table_obj.id_proof_value" update_data_type="string"></update_data>
					<update_data update_data_name="MESSAGE" update_db_data_type="string" update_data_source="lms_enquiry_table_obj.message" update_data_type="string"></update_data>
					<update_data update_data_name="UPDATED_BY" update_db_data_type="string" update_data_source="lms_enquiry_table_obj.updated_by" update_data_type="string"></update_data>
					<update_data update_data_name="UPDATED_DATE" update_db_data_type="date" update_data_source="lms_enquiry_table_obj.updated_date" update_data_type="string"></update_data>
					<update_data update_data_name="UPDATED_TIME" update_db_data_type="string" update_data_source="lms_enquiry_table_obj.updated_time" update_data_type="string"></update_data>
				</update_data_list>
				<update_filter_list>
					<update_filter update_filter_name="PHONE_NUMBER" update_filter_db_data_type="string" update_filter_data_source="ReqBrokerReqObj.request_data.phone_number" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
				</update_filter_list>
			</db_update_info>
		</exec_function>
	</exec_group>
	<exec_group group_name="edit_lms_enquiry_course_info" group_condition="ReqBrokerReqObj.header_data.channel=='web_client_portal'">
		<exec_function function_name="lms_enquiry_course_loop" function_type="start_loop" index_name="i" start_index="0" end_index="len(ReqBrokerReqObj.request_data.courses)" />
		<exec_function function_name="lms_enquiry_course_validate_fld" function_type="validate_fld">
			<validate_info app_err_desc="pattern validation failed for input" app_corrective_action="check the inputs and provide proper values">
				<validate_fld fld_name="ReqBrokerReqObj.request_data.courses[i].course_id" fld_data_type="string" fld_min_len="1" fld_max_len="128" fld_type="pattern" fld_pattern="^[a-zA-Z0-9.,|_* ]+$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.courses[i].db_operation" fld_data_type="string" fld_min_len="1" fld_max_len="128" fld_type="pattern" fld_pattern="^[INSERT|EDIT|DELETE]+$"></validate_fld>
			</validate_info>
		</exec_function>
		<exec_function function_name="insert_lms_enquiry_course_db_single_read" function_type="db_single_read" function_condition="ReqBrokerReqObj.request_data.courses[i].db_operation=='INSERT'">
			<db_single_read_info app_err_desc="Data Already Exists or No Rows in Result Set" app_corrective_action="check the inputs and provide proper values">
				<table_name>LMS_ENQUIRY_COURSE</table_name>
				<filter_info>
					<filter_data filter_name="ENQUIRY_ID" filter_id="ReqBrokerReqObj.request_data.id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="COURSE_ID" filter_id="ReqBrokerReqObj.request_data.courses[i].course_id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<result_info>
					<result_success>NoRows</result_success>
				</result_info>
			</db_single_read_info>
		</exec_function>
		<exec_function function_name="insert_lms_enquiry_course_table_copy_object" function_type="copy_object" function_condition="ReqBrokerReqObj.request_data.courses[i].db_operation=='INSERT'">
			<copy_object_info object_name="lms_enquiry_courses_table_obj">
				<object_data key="course_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.courses[i].course_id" data_type="string"></object_data>
				<object_data key="enquiry_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.id" data_type="string"></object_data>
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
		<exec_function function_name="lms_enquiry_course_table_insert" function_type="db_insert" function_condition="ReqBrokerReqObj.request_data.courses[i].db_operation=='INSERT'">
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
		<exec_function function_name="edit_lms_enquiry_course_db_single_read" function_type="db_single_read" function_condition="ReqBrokerReqObj.request_data.courses[i].db_operation=='EDIT'">
			<db_single_read_info app_err_desc="Data Already Exists or No Rows in Result Set" app_corrective_action="check the inputs and provide proper values">
				<table_name>LMS_ENQUIRY_COURSE</table_name>
				<filter_info>
					<filter_data filter_name="ENQUIRY_ID" filter_id="ReqBrokerReqObj.request_data.id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="COURSE_ID" filter_id="ReqBrokerReqObj.request_data.courses[i].course_id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_enquiry_course_table_record</store_obj_name>
			</db_single_read_info>
		</exec_function>
		<exec_function function_name="edit_lms_enquiry_course_table_copy_object" function_type="copy_object" function_condition="ReqBrokerReqObj.request_data.courses[i].db_operation=='EDIT'">
			<copy_object_info object_name="edit_lms_enquiry_course_table_obj">
				<object_data key="enquiry_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.id" data_type="string"></object_data>
				<object_data key="course_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.courses[i].course_id" data_type="string"></object_data>
				<object_data key="updated_by" data_source_type="reqbrokermap" data_source="decode_token_obj.email" data_type="string"></object_data>
				<object_data key="updated_date" data_source_type="key" data_source="get_date" data_type="string"></object_data>
				<object_data key="updated_time" data_source_type="key" data_source="get_time" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
		<exec_function function_name="edit_lms_enquiry_db_update" function_type="db_update" function_condition="ReqBrokerReqObj.request_data.courses[i].db_operation=='EDIT'">
			<db_update_info>
				<table_name>LMS_ENQUIRY_COURSE</table_name>
				<update_data_list>
					<update_data update_data_name="COURSE_ID" update_db_data_type="string" update_data_source="edit_lms_enquiry_course_table_obj.course_id" update_data_type="string"></update_data>
					<update_data update_data_name="UPDATED_BY" update_db_data_type="string" update_data_source="edit_lms_enquiry_course_table_obj.updated_by" update_data_type="string"></update_data>
					<update_data update_data_name="UPDATED_DATE" update_db_data_type="date" update_data_source="edit_lms_enquiry_course_table_obj.updated_date" update_data_type="string"></update_data>
					<update_data update_data_name="UPDATED_TIME" update_db_data_type="string" update_data_source="edit_lms_enquiry_course_table_obj.updated_time" update_data_type="string"></update_data>
				</update_data_list>
				<update_filter_list>
					<update_filter update_filter_name="ID" update_filter_db_data_type="string" update_filter_data_source="lms_enquiry_course_table_record.id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
					<update_filter update_filter_name="ENQUIRY_ID" update_filter_db_data_type="string" update_filter_data_source="edit_lms_enquiry_course_table_obj.enquiry_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
				</update_filter_list>
			</db_update_info>
		</exec_function>
		<exec_function function_name="delete_lms_enquiry_course_table_delete" function_type="db_delete"  function_condition="ReqBrokerReqObj.request_data.courses[i].db_operation=='DELETE'">
			<db_delete_info>
				<table_name>LMS_ENQUIRY_COURSE</table_name>
				<delete_filter_list>
					<delete_filter delete_filter_name="COURSE_ID" delete_filter_db_data_type="string" delete_filter_data_source="ReqBrokerReqObj.request_data.courses[i].course_id" delete_filter_data_type="string" delete_filter_condition="AND"></delete_filter>
					<delete_filter delete_filter_name="ENQUIRY_ID" delete_filter_db_data_type="string" delete_filter_data_source="ReqBrokerReqObj.request_data.id" delete_filter_data_type="string" delete_filter_condition="AND"></delete_filter>
				</delete_filter_list>
			</db_delete_info>
		</exec_function>
		<exec_function function_name="lms_enquiry_course_loop" function_type="end_loop" index_name="i" />
	</exec_group>
	<exec_group group_name="edit_lms_enquiry_comment_data"  group_condition="ReqBrokerReqObj.header_data.channel=='web_client_portal'">
		<exec_function function_name="lms_enquiry_comment_loop" function_type="start_loop" index_name="i" start_index="0" end_index="len(ReqBrokerReqObj.request_data.comments)" />
		<exec_function function_name="lms_enquiry_comment_validate_fld" function_type="validate_fld">
			<validate_info app_err_desc="pattern validation failed for input" app_corrective_action="check the inputs and provide proper values">
				<validate_fld fld_name="ReqBrokerReqObj.request_data.comments[i].comment_id" fld_data_type="string" fld_min_len="1" fld_max_len="128" fld_type="pattern" fld_pattern="^[a-zA-Z0-9.,|_* -]+$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.comments[i].comment" fld_data_type="string" fld_min_len="1" fld_max_len="128" fld_type="pattern" fld_pattern="^[a-zA-Z0-9.,|_* -]+$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.comments[i].db_operation" fld_data_type="string" fld_min_len="1" fld_max_len="128" fld_type="pattern" fld_pattern="^[EDIT|DELETE]+$"></validate_fld>
			</validate_info>
		</exec_function>
		<exec_function function_name="edit_lms_enquiry_comment_db_single_read" function_type="db_single_read" function_condition="ReqBrokerReqObj.request_data.comments[i].db_operation=='EDIT'">
			<db_single_read_info app_err_desc="Data Already Exists or No Rows in Result Set" app_corrective_action="check the inputs and provide proper values">
				<table_name>LMS_ENQUIRY_COMMENT</table_name>
				<filter_info>
					<filter_data filter_name="ID" filter_id="ReqBrokerReqObj.request_data.comments[i].comment_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="ENQUIRY_ID" filter_id="ReqBrokerReqObj.request_data.id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_enquiry_comment_table_record</store_obj_name>
			</db_single_read_info>
		</exec_function>
		<exec_function function_name="edit_lms_enquiry_comment_table_copy_object" function_type="copy_object" function_condition="ReqBrokerReqObj.request_data.comments[i].db_operation=='EDIT'">
			<copy_object_info object_name="lms_enquiry_comment_info_table_obj">
				<object_data key="comment_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.comments[i].comment_id" data_type="string"></object_data>
				<object_data key="enquiry_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.id" data_type="string"></object_data>
				<object_data key="comment" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.comments[i].comment" data_type="string"></object_data>
				<object_data key="commented_by" data_source_type="reqbrokermap" data_source="lms_enquiry_comment_table_record.commented_by" data_type="string"></object_data>
				<object_data key="commented_by_email" data_source_type="reqbrokermap" data_source="lms_enquiry_comment_table_record.commented_by_email" data_type="string"></object_data>
				<object_data key="commented_at" data_source_type="key" data_source="get_time" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
		<exec_function function_name="lms_enquiry_comment_db_update" function_type="db_update" function_condition="ReqBrokerReqObj.request_data.comments[i].db_operation=='EDIT'">
			<db_update_info>
				<table_name>LMS_ENQUIRY_COMMENT</table_name>
				<update_data_list>
					<update_data update_data_name="COMMENT" update_db_data_type="string" update_data_source="lms_enquiry_comment_info_table_obj.comment" update_data_type="string"></update_data>
					<update_data update_data_name="COMMENTED_BY" update_db_data_type="string" update_data_source="lms_enquiry_comment_info_table_obj.commented_by" update_data_type="string"></update_data>
					<update_data update_data_name="COMMENTED_BY_EMAIL" update_db_data_type="date" update_data_source="lms_enquiry_comment_info_table_obj.commented_by_email" update_data_type="string"></update_data>
					<update_data update_data_name="COMMENTED_AT" update_db_data_type="string" update_data_source="lms_enquiry_comment_info_table_obj.commented_at" update_data_type="string"></update_data>
				</update_data_list>
				<update_filter_list>
					<update_filter update_filter_name="ID" update_filter_db_data_type="string" update_filter_data_source="lms_enquiry_comment_info_table_obj.comment_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
					<update_filter update_filter_name="ENQUIRY_ID" update_filter_db_data_type="string" update_filter_data_source="lms_enquiry_comment_info_table_obj.enquiry_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
				</update_filter_list>
			</db_update_info>
		</exec_function>
		<exec_function function_name="delete_lms_enquiry_comment_table_delete" function_type="db_delete"  function_condition="ReqBrokerReqObj.request_data.comments[i].db_operation=='DELETE'">
			<db_delete_info>
				<table_name>LMS_ENQUIRY_COMMENT</table_name>
				<delete_filter_list>
					<delete_filter delete_filter_name="ID" delete_filter_db_data_type="string" delete_filter_data_source="ReqBrokerReqObj.request_data.comments[i].comment" delete_filter_data_type="string" delete_filter_condition="AND"></delete_filter>
					<delete_filter delete_filter_name="ENQUIRY_ID" delete_filter_db_data_type="string" delete_filter_data_source="ReqBrokerReqObj.request_data.id" delete_filter_data_type="string" delete_filter_condition="AND"></delete_filter>
				</delete_filter_list>
			</db_delete_info>
		</exec_function>
		<exec_function function_name="lms_enquiry_comment_loop" function_type="end_loop" index_name="i" />
	</exec_group>
	<exec_group group_name="read_lms_enquiry_db_multi_read"  group_condition="ReqBrokerReqObj.header_data.channel=='web_client_portal'">
		<exec_function function_name="read_view_lms_enquiry_db_multi_read" function_type="db_single_read">
			<db_single_read_info>
				<table_name>LMS_ENQUIRY</table_name>
				<filter_info>
					<filter_data filter_name="PHONE_NUMBER" filter_id="ReqBrokerReqObj.request_data.phone_number" filter_data_type="string" filter_condition="AND"></filter_data>
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
	<exec_group group_name="edit_lms_enquiry_personal_info" group_condition="ReqBrokerReqObj.header_data.channel=='web_admin_portal'">
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
		<exec_function function_name="edit_lms_enquiry_personal_info_validate_fld" function_type="validate_fld">
			<validate_info app_err_desc="pattern validation failed for input" app_corrective_action="check the inputs and provide proper values">
				<validate_fld fld_name="ReqBrokerReqObj.request_data.id" fld_data_type="string" fld_min_len="1" fld_max_len="128" fld_type="pattern" fld_pattern="^[a-zA-Z0-9.,|_* -]+$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.first_name" fld_data_type="string" fld_min_len="1" fld_max_len="128" fld_type="pattern" fld_pattern="^[a-zA-Z0-9.,|_* -]+$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.last_name" fld_data_type="string" fld_min_len="1" fld_max_len="128" fld_type="pattern" fld_pattern="^[a-zA-Z0-9.,|_* -]+$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.phone_number" fld_data_type="string" fld_min_len="1" fld_max_len="32" fld_type="pattern" fld_pattern="^\+\d{1,3}-\d{10}$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.email" fld_data_type="string" fld_min_len="1" fld_max_len="128" fld_type="pattern" fld_pattern="^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,3}$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.education" fld_data_type="string" fld_min_len="0" fld_max_len="255" fld_type="pattern" fld_pattern="^[A-Za-z0-9 .,-]*$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.college" fld_data_type="string" fld_min_len="0" fld_max_len="500" fld_type="pattern" fld_pattern="^[A-Za-z0-9 .,-]*$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.department" fld_data_type="string" fld_min_len="0" fld_max_len="128" fld_type="pattern" fld_pattern="^[A-Za-z0-9 ._-]*$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.address" fld_data_type="string" fld_min_len="0" fld_max_len="500" fld_type="pattern" fld_pattern="^[a-zA-Z0-9.,|_* /-]+$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.id_proof_type" fld_data_type="string" fld_min_len="0" fld_max_len="50" fld_type="pattern" fld_pattern="^[A-Za-z0-9 _-]*$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.id_proof_value" fld_data_type="string" fld_min_len="0" fld_max_len="128" fld_type="pattern" fld_pattern="^[A-Za-z0-9 -]*$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.message" fld_data_type="string" fld_min_len="0" fld_max_len="500" fld_type="pattern" fld_pattern="^[a-zA-Z0-9.,|_* -]+$"></validate_fld>
			    <validate_fld fld_name="ReqBrokerReqObj.request_data.onboarded" fld_data_type="boolean" fld_min_len="0" fld_max_len="5" fld_type="pattern" fld_pattern="^(true|false|0|1)$" />
				<validate_fld fld_name="ReqBrokerReqObj.request_data.onboarding_status" fld_data_type="string" fld_min_len="0" fld_max_len="50" fld_type="pattern" fld_pattern="^[A-Za-z0-9 _-]*$" />
				<validate_fld fld_name="ReqBrokerReqObj.request_data.payment_status" fld_data_type="string" fld_min_len="0" fld_max_len="25" fld_type="pattern" fld_pattern="^[A-Za-z0-9 _-]*$" />
				<validate_fld fld_name="ReqBrokerReqObj.request_data.payment_mode" fld_data_type="string" fld_min_len="0" fld_max_len="25" fld_type="pattern" fld_pattern="^[A-Za-z0-9 _-]*$" />
				<validate_fld fld_name="ReqBrokerReqObj.request_data.payment_id" fld_data_type="string" fld_min_len="0" fld_max_len="255" fld_type="pattern" fld_pattern="^[A-Za-z0-9 _-]*$" />
				<validate_fld fld_name="ReqBrokerReqObj.request_data.payment_amount" fld_data_type="string" fld_min_len="0" fld_max_len="15" fld_type="pattern" fld_pattern="^[0-9]+(\.[0-9]{1,2})?$" />
				<validate_fld fld_name="ReqBrokerReqObj.request_data.payment_date" fld_data_type="string" fld_min_len="0" fld_max_len="30" fld_type="pattern" fld_pattern="^[0-9T:\-+ ]*$" />
				</validate_info>
		</exec_function>
		<exec_function function_name="edit_lms_enquiry_personal_info_auth_db_single_read" function_type="db_single_read">
			<db_single_read_info app_err_desc="Data Already Exists or No Rows in Result Set" app_corrective_action="check the inputs and provide proper values">
				<table_name>LMS_ENQUIRY</table_name>
				<filter_info>
					<filter_data filter_name="PHONE_NUMBER" filter_id="ReqBrokerReqObj.request_data.phone_number" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_enquiry_table_record</store_obj_name>
			</db_single_read_info>
		</exec_function>
		<exec_function function_name="is_onboard_lms_enquiry" function_type="validate_condition" function_condition="lms_enquiry_table_record.onboarded==true">
			<validate_condition_info app_err_desc="Once user onboarded, status can't be changed!" app_corrective_action="status can't be changed.">
				<validate_condition_expression>ReqBrokerReqObj.request_data.onboarded==true</validate_condition_expression>
			</validate_condition_info>
		</exec_function>
		<exec_function function_name="edit_lms_enquiry_personal_info_table_copy_object" function_type="copy_object">
			<copy_object_info object_name="lms_enquiry_table_obj">
				<object_data key="first_name" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.first_name" data_type="string"></object_data>
				<object_data key="last_name" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.last_name" data_type="string"></object_data>
				<object_data key="phone_number" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.phone_number" data_type="string"></object_data>
				<object_data key="email" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.email" data_type="string"></object_data>
				<object_data key="onboarded" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.onboarded" data_type="Boolean"></object_data>
				<object_data key="onboarding_status" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.onboarding_status" data_type="string"></object_data>
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
				<object_data key="updated_by" data_source_type="reqbrokermap" data_source="decode_token_obj.email" data_type="string"></object_data>
				<object_data key="updated_date" data_source_type="key" data_source="get_date" data_type="string"></object_data>
				<object_data key="updated_time" data_source_type="key" data_source="get_time" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
		<exec_function function_name="edit_lms_enquiry_personal_info_email_true_table_copy_object" function_type="copy_object" function_condition="ReqBrokerReqObj.request_data.is_email_verified==true">
			<copy_object_info object_name="lms_enquiry_email_verify_obj">
                <object_data key="is_email_verified" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.is_email_verified" data_type="Boolean"></object_data>
				<object_data key="email_verified_at" data_source_type="key" data_source="get_time_stamp" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
		<exec_function function_name="edit_lms_enquiry_personal_info_email_table_copy_object" function_type="copy_object" function_condition="ReqBrokerReqObj.request_data.is_email_verified!=true">
			<copy_object_info object_name="lms_enquiry_email_verify_obj">
                <object_data key="is_email_verified" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.is_email_verified" data_type="Boolean"></object_data>
				<object_data key="email_verified_at" data_source_type="raw_value" data_source="1970-01-01 00:00:00.000" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
		<exec_function function_name="lms_enquiry_personal_info_db_update" function_type="db_update">
			<db_update_info>
				<table_name>LMS_ENQUIRY</table_name>
				<update_data_list>
					<update_data update_data_name="FIRST_NAME" update_db_data_type="string" update_data_source="lms_enquiry_table_obj.first_name" update_data_type="string"></update_data>
					<update_data update_data_name="LAST_NAME" update_db_data_type="string" update_data_source="lms_enquiry_table_obj.last_name" update_data_type="string"></update_data>
					  <update_data update_data_name="IS_EMAIL_VERIFIED" update_db_data_type="boolean" update_data_source="lms_enquiry_email_verify_obj.is_email_verified" update_data_type="boolean"></update_data>
                    <update_data update_data_name="EMAIL_VERIFIED_AT" update_db_data_type="string" update_data_source="lms_enquiry_email_verify_obj.email_verified_at" update_data_type="string"></update_data>
                    <update_data update_data_name="EMAIL" update_db_data_type="string" update_data_source="lms_enquiry_table_obj.email" update_data_type="string"></update_data>
					<update_data update_data_name="EDUCATION" update_db_data_type="string" update_data_source="lms_enquiry_table_obj.education" update_data_type="string"></update_data>
					<update_data update_data_name="COLLEGE" update_db_data_type="string" update_data_source="lms_enquiry_table_obj.college" update_data_type="string"></update_data>
					<update_data update_data_name="DEPARTMENT" update_db_data_type="string" update_data_source="lms_enquiry_table_obj.department" update_data_type="string"></update_data>
					<update_data update_data_name="ADDRESS" update_db_data_type="string" update_data_source="lms_enquiry_table_obj.address" update_data_type="string"></update_data>
					<update_data update_data_name="ID_PROOF_TYPE" update_db_data_type="string" update_data_source="lms_enquiry_table_obj.id_proof_type" update_data_type="string"></update_data>
					<update_data update_data_name="ID_PROOF_VALUE" update_db_data_type="string" update_data_source="lms_enquiry_table_obj.id_proof_value" update_data_type="string"></update_data>
					<update_data update_data_name="MESSAGE" update_db_data_type="string" update_data_source="lms_enquiry_table_obj.message" update_data_type="string"></update_data>
					<update_data update_data_name="UPDATED_BY" update_db_data_type="string" update_data_source="lms_enquiry_table_obj.updated_by" update_data_type="string"></update_data>
					<update_data update_data_name="UPDATED_DATE" update_db_data_type="date" update_data_source="lms_enquiry_table_obj.updated_date" update_data_type="string"></update_data>
					<update_data update_data_name="UPDATED_TIME" update_db_data_type="string" update_data_source="lms_enquiry_table_obj.updated_time" update_data_type="string"></update_data>
				    <update_data update_data_name="ONBOARDED" update_db_data_type="boolean" update_data_source="ReqBrokerReqObj.request_data.onboarded" update_data_type="boolean"></update_data>
                    <update_data update_data_name="ONBOARDING_STATUS" update_db_data_type="string" update_data_source="lms_enquiry_table_obj.onboarding_status" update_data_type="string"></update_data>
                    <update_data update_data_name="PAYMENT_STATUS" update_db_data_type="string" update_data_source="lms_enquiry_table_obj.payment_status" update_data_type="string"></update_data>
					<update_data update_data_name="PAYMENT_MODE" update_db_data_type="string" update_data_source="lms_enquiry_table_obj.payment_mode" update_data_type="string"></update_data>
                    <update_data update_data_name="PAYMENT_DATE" update_db_data_type="string" update_data_source="lms_enquiry_table_obj.payment_date" update_data_type="string"></update_data>
					<update_data update_data_name="PAYMENT_ID" update_db_data_type="string" update_data_source="lms_enquiry_table_obj.payment_id" update_data_type="string"></update_data>
					<update_data update_data_name="PAYMENT_AMOUNT" update_db_data_type="string" update_data_source="lms_enquiry_table_obj.payment_amount" update_data_type="string"></update_data>
					</update_data_list>
				<update_filter_list>
					<update_filter update_filter_name="PHONE_NUMBER" update_filter_db_data_type="string" update_filter_data_source="ReqBrokerReqObj.request_data.phone_number" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
				</update_filter_list>
			</db_update_info>
		</exec_function>
	</exec_group>
	<exec_group group_name="edit_lms_enquiry_course_info" group_condition="ReqBrokerReqObj.header_data.channel=='web_admin_portal'">
		<exec_function function_name="lms_enquiry_course_loop" function_type="start_loop" index_name="i" start_index="0" end_index="len(ReqBrokerReqObj.request_data.courses)" />
		<exec_function function_name="lms_enquiry_course_validate_fld" function_type="validate_fld">
			<validate_info app_err_desc="pattern validation failed for input" app_corrective_action="check the inputs and provide proper values">
				<validate_fld fld_name="ReqBrokerReqObj.request_data.courses[i].course_id" fld_data_type="string" fld_min_len="1" fld_max_len="128" fld_type="pattern" fld_pattern="^[a-zA-Z0-9.,|_* ]+$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.courses[i].db_operation" fld_data_type="string" fld_min_len="1" fld_max_len="128" fld_type="pattern" fld_pattern="^[INSERT|EDIT|DELETE]+$"></validate_fld>
			</validate_info>
		</exec_function>
		<exec_function function_name="insert_lms_enquiry_course_db_single_read" function_type="db_single_read" function_condition="ReqBrokerReqObj.request_data.courses[i].db_operation=='INSERT'">
			<db_single_read_info app_err_desc="Data Already Exists or No Rows in Result Set" app_corrective_action="check the inputs and provide proper values">
				<table_name>LMS_ENQUIRY_COURSE</table_name>
				<filter_info>
					<filter_data filter_name="ENQUIRY_ID" filter_id="ReqBrokerReqObj.request_data.id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="COURSE_ID" filter_id="ReqBrokerReqObj.request_data.courses[i].course_id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<result_info>
					<result_success>NoRows</result_success>
				</result_info>
			</db_single_read_info>
		</exec_function>
		<exec_function function_name="insert_lms_enquiry_course_table_copy_object" function_type="copy_object" function_condition="ReqBrokerReqObj.request_data.courses[i].db_operation=='INSERT'">
			<copy_object_info object_name="lms_enquiry_courses_table_obj">
				<object_data key="course_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.courses[i].course_id" data_type="string"></object_data>
				<object_data key="enquiry_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.id" data_type="string"></object_data>
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
		<exec_function function_name="lms_enquiry_course_table_insert" function_type="db_insert" function_condition="ReqBrokerReqObj.request_data.courses[i].db_operation=='INSERT'">
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
		<exec_function function_name="edit_lms_enquiry_course_db_single_read" function_type="db_single_read" function_condition="ReqBrokerReqObj.request_data.courses[i].db_operation=='EDIT'">
			<db_single_read_info app_err_desc="Data Already Exists or No Rows in Result Set" app_corrective_action="check the inputs and provide proper values">
				<table_name>LMS_ENQUIRY_COURSE</table_name>
				<filter_info>
					<filter_data filter_name="ENQUIRY_ID" filter_id="ReqBrokerReqObj.request_data.id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="COURSE_ID" filter_id="ReqBrokerReqObj.request_data.courses[i].course_id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_enquiry_course_table_record</store_obj_name>
			</db_single_read_info>
		</exec_function>
		<exec_function function_name="edit_lms_enquiry_course_table_copy_object" function_type="copy_object" function_condition="ReqBrokerReqObj.request_data.courses[i].db_operation=='EDIT'">
			<copy_object_info object_name="edit_lms_enquiry_course_table_obj">
				<object_data key="enquiry_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.id" data_type="string"></object_data>
				<object_data key="course_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.courses[i].course_id" data_type="string"></object_data>
				<object_data key="updated_by" data_source_type="reqbrokermap" data_source="decode_token_obj.email" data_type="string"></object_data>
				<object_data key="updated_date" data_source_type="key" data_source="get_date" data_type="string"></object_data>
				<object_data key="updated_time" data_source_type="key" data_source="get_time" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
		<exec_function function_name="edit_lms_enquiry_db_update" function_type="db_update" function_condition="ReqBrokerReqObj.request_data.courses[i].db_operation=='EDIT'">
			<db_update_info>
				<table_name>LMS_ENQUIRY_COURSE</table_name>
				<update_data_list>
					<update_data update_data_name="COURSE_ID" update_db_data_type="string" update_data_source="edit_lms_enquiry_course_table_obj.course_id" update_data_type="string"></update_data>
					<update_data update_data_name="UPDATED_BY" update_db_data_type="string" update_data_source="edit_lms_enquiry_course_table_obj.updated_by" update_data_type="string"></update_data>
					<update_data update_data_name="UPDATED_DATE" update_db_data_type="date" update_data_source="edit_lms_enquiry_course_table_obj.updated_date" update_data_type="string"></update_data>
					<update_data update_data_name="UPDATED_TIME" update_db_data_type="string" update_data_source="edit_lms_enquiry_course_table_obj.updated_time" update_data_type="string"></update_data>
				</update_data_list>
				<update_filter_list>
					<update_filter update_filter_name="ID" update_filter_db_data_type="string" update_filter_data_source="lms_enquiry_course_table_record.id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
					<update_filter update_filter_name="ENQUIRY_ID" update_filter_db_data_type="string" update_filter_data_source="edit_lms_enquiry_course_table_obj.enquiry_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
				</update_filter_list>
			</db_update_info>
		</exec_function>
		<exec_function function_name="delete_lms_enquiry_course_table_delete" function_type="db_delete"  function_condition="ReqBrokerReqObj.request_data.courses[i].db_operation=='DELETE'">
			<db_delete_info>
				<table_name>LMS_ENQUIRY_COURSE</table_name>
				<delete_filter_list>
					<delete_filter delete_filter_name="COURSE_ID" delete_filter_db_data_type="string" delete_filter_data_source="ReqBrokerReqObj.request_data.courses[i].course_id" delete_filter_data_type="string" delete_filter_condition="AND"></delete_filter>
					<delete_filter delete_filter_name="ENQUIRY_ID" delete_filter_db_data_type="string" delete_filter_data_source="ReqBrokerReqObj.request_data.id" delete_filter_data_type="string" delete_filter_condition="AND"></delete_filter>
				</delete_filter_list>
			</db_delete_info>
		</exec_function>
		<exec_function function_name="lms_enquiry_course_loop" function_type="end_loop" index_name="i" />
	</exec_group>
	<exec_group group_name="edit_lms_enquiry_comment_data"  group_condition="ReqBrokerReqObj.header_data.channel=='web_admin_portal'">
		<exec_function function_name="lms_enquiry_comment_loop" function_type="start_loop" index_name="i" start_index="0" end_index="len(ReqBrokerReqObj.request_data.comments)" />
		<exec_function function_name="lms_enquiry_comment_validate_fld" function_type="validate_fld">
			<validate_info app_err_desc="pattern validation failed for input" app_corrective_action="check the inputs and provide proper values">
				<validate_fld fld_name="ReqBrokerReqObj.request_data.comments[i].comment" fld_data_type="string" fld_min_len="1" fld_max_len="128" fld_type="pattern" fld_pattern="^[a-zA-Z0-9.,|_* -]+$"></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.comments[i].db_operation" fld_data_type="string" fld_min_len="1" fld_max_len="128" fld_type="pattern" fld_pattern="^[INSERT|EDIT|DELETE]+$"></validate_fld>
			</validate_info>
		</exec_function>
		<exec_function function_name="insert_lms_enquiry_comments_table_copy_object" function_type="copy_object" unction_condition="ReqBrokerReqObj.request_data.comments[i].db_operation=='INSERT'">
			<copy_object_info object_name="lms_enquiry_comments_table_obj">
				<object_data key="comment" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.comments[i].comment" data_type="string"></object_data>
				<object_data key="enquiry_id" data_source_type="reqbrokermap" data_source="lms_enquiry_table_record.id" data_type="string"></object_data>
				<object_data key="commented_by" data_source_type="reqbrokermap" data_source="decode_token_obj.email" data_type="string"></object_data>
				<object_data key="commented_at" data_source_type="key" data_source="get_time_stamp" data_type="string"></object_data>
				<object_data key="commented_by_email" data_source_type="reqbrokermap" data_source="decode_token_obj.email" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
		<exec_function function_name="insert_lms_enquiry_auth_table_insert" function_type="db_insert" function_condition="ReqBrokerReqObj.request_data.comments[i].db_operation=='INSERT'">
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
		<exec_function function_name="edit_lms_enquiry_comment_db_single_read" function_type="db_single_read" function_condition="ReqBrokerReqObj.request_data.comments[i].db_operation=='EDIT'">
			<db_single_read_info app_err_desc="Data Already Exists or No Rows in Result Set" app_corrective_action="check the inputs and provide proper values">
				<table_name>LMS_ENQUIRY_COMMENT</table_name>
				<filter_info>
					<filter_data filter_name="ID" filter_id="ReqBrokerReqObj.request_data.comments[i].comment_id" filter_data_type="string" filter_condition="AND"></filter_data>
					<filter_data filter_name="ENQUIRY_ID" filter_id="ReqBrokerReqObj.request_data.id" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_enquiry_comment_table_record</store_obj_name>
			</db_single_read_info>
		</exec_function>
		<exec_function function_name="edit_lms_enquiry_comment_table_copy_object" function_type="copy_object" function_condition="ReqBrokerReqObj.request_data.comments[i].db_operation=='EDIT'">
			<copy_object_info object_name="lms_enquiry_comment_info_table_obj">
				<object_data key="comment_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.comments[i].comment_id" data_type="string"></object_data>
				<object_data key="enquiry_id" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.id" data_type="string"></object_data>
				<object_data key="comment" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.comments[i].comment" data_type="string"></object_data>
				<object_data key="commented_by" data_source_type="reqbrokermap" data_source="lms_enquiry_comment_table_record.commented_by" data_type="string"></object_data>
				<object_data key="commented_by_email" data_source_type="reqbrokermap" data_source="lms_enquiry_comment_table_record.commented_by_email" data_type="string"></object_data>
				<object_data key="commented_at" data_source_type="key" data_source="get_time_stamp" data_type="string"></object_data>
			</copy_object_info>
		</exec_function>
		<exec_function function_name="lms_enquiry_comment_db_update" function_type="db_update" function_condition="ReqBrokerReqObj.request_data.comments[i].db_operation=='EDIT'">
			<db_update_info>
				<table_name>LMS_ENQUIRY_COMMENT</table_name>
				<update_data_list>
					<update_data update_data_name="COMMENT" update_db_data_type="string" update_data_source="lms_enquiry_comment_info_table_obj.comment" update_data_type="string"></update_data>
					<update_data update_data_name="COMMENTED_BY" update_db_data_type="string" update_data_source="lms_enquiry_comment_info_table_obj.commented_by" update_data_type="string"></update_data>
					<update_data update_data_name="COMMENTED_BY_EMAIL" update_db_data_type="string" update_data_source="lms_enquiry_comment_info_table_obj.commented_by_email" update_data_type="string"></update_data>
					<update_data update_data_name="COMMENTED_AT" update_db_data_type="string" update_data_source="lms_enquiry_comment_info_table_obj.commented_at" update_data_type="string"></update_data>
				</update_data_list>
				<update_filter_list>
					<update_filter update_filter_name="ID" update_filter_db_data_type="string" update_filter_data_source="lms_enquiry_comment_info_table_obj.comment_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
					<update_filter update_filter_name="ENQUIRY_ID" update_filter_db_data_type="string" update_filter_data_source="lms_enquiry_comment_info_table_obj.enquiry_id" update_filter_data_type="string" update_filter_condition="AND"></update_filter>
				</update_filter_list>
			</db_update_info>
		</exec_function>
		<exec_function function_name="delete_lms_enquiry_comment_table_delete" function_type="db_delete"  function_condition="ReqBrokerReqObj.request_data.comments[i].db_operation=='DELETE'">
			<db_delete_info>
				<table_name>LMS_ENQUIRY_COMMENT</table_name>
				<delete_filter_list>
					<delete_filter delete_filter_name="ID" delete_filter_db_data_type="string" delete_filter_data_source="ReqBrokerReqObj.request_data.comments[i].comment_id" delete_filter_data_type="string" delete_filter_condition="AND"></delete_filter>
					<delete_filter delete_filter_name="ENQUIRY_ID" delete_filter_db_data_type="string" delete_filter_data_source="ReqBrokerReqObj.request_data.id" delete_filter_data_type="string" delete_filter_condition="AND"></delete_filter>
				</delete_filter_list>
			</db_delete_info>
		</exec_function>
		<exec_function function_name="lms_enquiry_comment_loop" function_type="end_loop" index_name="i" />
	</exec_group>
	<exec_group group_name="insert_lms_user_data" group_condition="ReqBrokerReqObj.header_data.channel=='web_admin_portal' &amp;&amp; ReqBrokerReqObj.request_data.onboarded==true&amp;&amp;lms_enquiry_table_record.onboarded!=true" >
		<exec_function function_name="lms_user_data_validate_fld" function_type="validate_fld">
			<validate_info app_err_desc="Validation failed for user data" app_corrective_action="Please validate the data">
				<validate_fld fld_name="ReqBrokerReqObj.request_data.user_info.password" fld_data_type="string" fld_min_len="8" fld_max_len="255" fld_type="pattern" fld_pattern=".*" app_err_desc="Pattern validation failed for password" app_corrective_action="Please enter a valid password with a minimum of 6 characters, including at least one uppercase letter, one numeric digit, and one special character."></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.user_info.role" fld_data_type="string" fld_min_len="5" fld_max_len="20" fld_type="pattern" fld_pattern="^(STUDENT)$" app_err_desc="Invalid ROLE" app_corrective_action="Allowed values are ADMIN, STUDENT, TRAINER, INSTRUCTOR."></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.user_info.status" fld_data_type="boolean" fld_min_len="1" fld_max_len="5" fld_type="pattern" fld_pattern="^(true|false|0|1)$" app_err_desc="Invalid status" app_corrective_action="Only Boolean value is allowed."></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.user_info.profile_picture_url" fld_data_type="string" fld_min_len="0" fld_max_len="500" fld_type="pattern" fld_pattern=".*" app_err_desc="Invalid URL" app_corrective_action="Provide a valid URL."></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.user_info.date_of_birth" fld_data_type="string" fld_min_len="1" fld_max_len="20" fld_type="pattern" fld_pattern=".*" app_err_desc="Invalid DATE_OF_BIRTH" app_corrective_action="Only digits and special characters allowed."></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.user_info.gender" fld_data_type="string" fld_min_len="1" fld_max_len="10" fld_type="pattern" fld_pattern=".*" app_err_desc="Invalid GENDER" app_corrective_action="Allowed: male, female, others."></validate_fld>
				<validate_fld fld_name="ReqBrokerReqObj.request_data.user_info.bio" fld_data_type="string" fld_min_len="0" fld_max_len="500" fld_type="pattern" fld_pattern=".*" app_err_desc="Invalid BIO" app_corrective_action="Only letters, digits, spaces and special characters allowed."></validate_fld>
			</validate_info>
		</exec_function>
			<exec_function function_name="lms_user_data_password_validate_condition" function_type="validate_condition">
			<validate_condition_info app_err_desc="User onboarded status should be false" app_corrective_action="Make User Onboarded status false.">
				<validate_condition_expression>ReqBrokerReqObj.request_data.onboarding_status!=false</validate_condition_expression>
			</validate_condition_info>
		</exec_function>
			<exec_function function_name="insert_lms_user_data_db_single_read" function_type="db_single_read">
				<db_single_read_info app_err_desc="lms user data data table is not found" app_corrective_action="User data cannot be configured as the table does not exist">
					<table_name>LMS_USER_DATA</table_name>
					<filter_info>
						<filter_data filter_name="MOBILE_NUMBER" filter_id="ReqBrokerReqObj.request_data.phone_number" filter_data_type="string" filter_condition="AND"></filter_data>
					</filter_info>
					<result_info app_err_desc="The phone_number:{{.ReqBrokerReqObj.request_data.phone_number}} already exists." app_corrective_action="Please configure new mobile number and proceed">
						<result_success>NoRows</result_success>
					</result_info>
				</db_single_read_info>
				<err_info>
					<err_desc err_code="Duplicate" app_err_desc="Duplicate records found for the phone_number:{{.ReqBrokerReqObj.request_data.phone_number}} provided in the input." app_corrective_action="Make sure each Mobile Number shows up only once." />   
				</err_info>
			</exec_function>
    
			<exec_function function_name="lms_user_data_password_transform" function_type="transform_object">
			<transform_object_info object_name="password_transform_obj">
				<object_data algo="hash" key="hash_password" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.user_info.password" data_type="string" />
			</transform_object_info>
		</exec_function>
            <exec_function function_name="full_name_copy_object" function_type="copy_object">
				<copy_object_info object_name="full_name_copy_data_obj">
					<object_data key="full_name" data_source_type="compute_str_expr" data_source="ReqBrokerReqObj.request_data.first_name+ReqBrokerReqObj.request_data.last_name" data_type="string"></object_data>
				</copy_object_info>
			</exec_function>

			<exec_function function_name="insert_lms_user_data_create_object" function_type="copy_object">
				<copy_object_info object_name="insert_lms_user_data_table_obj">
					<object_data key="full_name" data_source_type="reqbrokermap" data_source="full_name_copy_data_obj.full_name" data_type="string"></object_data>
					<object_data key="email" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.email" data_type="string"></object_data>
					<object_data key="mobile_number" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.phone_number" data_type="string"></object_data>
					<object_data key="password_hash" data_source_type="reqbrokermap" data_source="password_transform_obj.hash_password" data_type="string"></object_data>
					<object_data key="role" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.user_info.role" data_type="string"></object_data>
					<object_data key="status" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.user_info.status" data_type="Boolean"></object_data>
					<object_data key="profile_picture_url" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.user_info.profile_picture_url" data_type="string"></object_data>
					<object_data key="date_of_birth" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.user_info.date_of_birth" data_type="string"></object_data>
					<object_data key="gender" data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.user_info.gender" data_type="string"></object_data>
					<object_data key="address" data_source_type="reqbrokermap"  data_source="ReqBrokerReqObj.request_data.user_info.address" data_type="string"></object_data>
					<object_data key="bio" data_source_type="reqbrokermap"  data_source="ReqBrokerReqObj.request_data.user_info.bio" data_type="string"></object_data>
                    <object_data key="created_by" data_source_type="reqbrokermap" data_source="decode_token_obj.email" data_type="string"></object_data>
					<object_data key="created_date" data_source_type="key" data_source="get_date" data_type="string"></object_data>
					<object_data key="created_time" data_source_type="key" data_source="get_time" data_type="string"></object_data>
					<object_data key="updated_by" data_source_type="key" data_source="get_na" data_type="string"></object_data>
					<object_data key="updated_date" data_source_type="key" data_source="get_date" data_type="string"></object_data>
					<object_data key="updated_time" data_source_type="key" data_source="get_time" data_type="string"></object_data>
					<object_data key="deleted_by" data_source_type="key" data_source="get_na" data_type="string"></object_data>
					<object_data key="deleted_date" data_source_type="key" data_source="get_date" data_type="string"></object_data>
					<object_data key="deleted_time" data_source_type="key" data_source="get_time" data_type="string"></object_data>
					<object_data key="passwd_retry_count"  data_source_type="raw_value" data_source="0" data_type="string"></object_data>
					<object_data key="passwd_modified_date" data_source_type="key" data_source="get_date" data_type="string"></object_data>
					<object_data key="passwd_modified_time"  data_source_type="key" data_source="get_time" data_type="string"></object_data>
					<object_data key="last_login_date" data_source_type="key" data_source="get_date" data_type="string"></object_data>
					<object_data key="last_login_time"  data_source_type="key" data_source="get_time" data_type="string"></object_data>
                    <object_data key="passwd_status"  data_source_type="raw_value" data_source="2" data_type="string"></object_data>
				</copy_object_info>
			</exec_function>

			<exec_function function_name="insert_lms_user_data_db_insert" function_type="db_insert">
				<db_insert_info>
					<table_name>LMS_USER_DATA</table_name>
					<insert_data_list>
					    <insert_data insert_data_name="FULL_NAME" insert_db_data_type="string" insert_data_source="full_name_copy_data_obj.full_name" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="EMAIL" insert_db_data_type="string" insert_data_source="insert_lms_user_data_table_obj.email" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="MOBILE_NUMBER" insert_db_data_type="string" insert_data_source="insert_lms_user_data_table_obj.mobile_number" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="PASSWORD_HASH" insert_db_data_type="string" insert_data_source="insert_lms_user_data_table_obj.password_hash" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="ROLE" insert_db_data_type="string" insert_data_source="insert_lms_user_data_table_obj.role" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="STATUS" insert_db_data_type="boolean" insert_data_source="ReqBrokerReqObj.request_data.user_info.status" insert_data_type="boolean"></insert_data>
						<insert_data insert_data_name="PROFILE_PICTURE_URL" insert_db_data_type="string" insert_data_source="insert_lms_user_data_table_obj.profile_picture_url" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="DATE_OF_BIRTH" insert_db_data_type="date" insert_data_source="insert_lms_user_data_table_obj.date_of_birth" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="GENDER" insert_db_data_type="string" insert_data_source="insert_lms_user_data_table_obj.gender" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="ADDRESS" insert_db_data_type="string" insert_data_source="insert_lms_user_data_table_obj.address" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="BIO" insert_db_data_type="string" insert_data_source="insert_lms_user_data_table_obj.bio" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="CREATED_BY" insert_db_data_type="string" insert_data_source="insert_lms_user_data_table_obj.created_by" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="CREATED_DATE" insert_db_data_type="date" insert_data_source="insert_lms_user_data_table_obj.created_date" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="CREATED_TIME" insert_db_data_type="string" insert_data_source="insert_lms_user_data_table_obj.created_time" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="UPDATED_BY" insert_db_data_type="string" insert_data_source="insert_lms_user_data_table_obj.updated_by" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="UPDATED_DATE" insert_db_data_type="date" insert_data_source="insert_lms_user_data_table_obj.updated_date" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="UPDATED_TIME" insert_db_data_type="string" insert_data_source="insert_lms_user_data_table_obj.updated_time" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="DELETED_BY" insert_db_data_type="string" insert_data_source="insert_lms_user_data_table_obj.deleted_by" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="DELETED_DATE" insert_db_data_type="date" insert_data_source="insert_lms_user_data_table_obj.deleted_date" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="DELETED_TIME" insert_db_data_type="string" insert_data_source="insert_lms_user_data_table_obj.deleted_time" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="PASSWD_RETRY_COUNT" insert_db_data_type="string" insert_data_source="insert_lms_user_data_table_obj.passwd_retry_count" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="PASSWD_MODIFIED_DATE" insert_db_data_type="date" insert_data_source="insert_lms_user_data_table_obj.passwd_modified_date" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="PASSWD_MODIFIED_TIME" insert_db_data_type="string" insert_data_source="insert_lms_user_data_table_obj.passwd_modified_time" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="LAST_LOGIN_DATE" insert_db_data_type="date" insert_data_source="insert_lms_user_data_table_obj.last_login_date" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="LAST_LOGIN_TIME" insert_db_data_type="string" insert_data_source="insert_lms_user_data_table_obj.last_login_time" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="PASSWD_STATUS" insert_db_data_type="string" insert_data_source="insert_lms_user_data_table_obj.passwd_status" insert_data_type="string"></insert_data>
						</insert_data_list>
				</db_insert_info>
			</exec_function>
			<exec_function function_name="edit_lms_enquiry_personal_info_db_single_read" function_type="db_single_read">
			    <db_single_read_info app_err_desc="Data Already Exists or No Rows in Result Set" app_corrective_action="check the inputs and provide proper values">
			     	<table_name>LMS_USER_DATA</table_name>
				<filter_info>
					<filter_data filter_name="EMAIL" filter_id="ReqBrokerReqObj.request_data.email" filter_data_type="string" filter_condition="AND"></filter_data>
				</filter_info>
				<store_obj_name>lms_user_data_table_record</store_obj_name>
			</db_single_read_info>
				<err_info>
					<err_desc err_code="NoRows" app_err_desc="The Email:{{.ReqBrokerReqObj.request_data.Email}} not found." app_corrective_action="Please create the user and proceed"/>   
					<err_desc err_code="Duplicate" app_err_desc="The Email:{{.ReqBrokerReqObj.request_data.Email}} already exists." app_corrective_action="Please Enter new email and proceed" />   
				</err_info>
			</exec_function>
			 <exec_function function_name="lms_paid_course_loop" function_type="start_loop" index_name="i" start_index="0" end_index="len(ReqBrokerReqObj.request_data.paid_courses)" />
		<exec_function function_name="lms_enquiry_courses_validate_fld" function_type="validate_fld">
			<validate_info app_err_desc="pattern validation failed for input" app_corrective_action="check the inputs and provide proper values">
				<validate_fld fld_name="ReqBrokerReqObj.request_data.paid_courses[i]" fld_data_type="string" fld_min_len="1" fld_max_len="50" fld_type="pattern" fld_pattern="^[0-9]+$"></validate_fld>
			</validate_info>
		</exec_function>
         
			
            <exec_function function_name="insert_lms_course_progress_copy_object" function_type="copy_object">
			<copy_object_info object_name="insert_lms_course_progress_additional_obj">
				<object_data key="user_id" data_source_type="reqbrokermap" data_source="lms_user_data_table_record.id" data_type="string"></object_data>
				<object_data key="course_id"  data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.paid_courses[i]" data_type="string"></object_data>
				<object_data key="percentage" data_source_type="raw_value" data_source="0" data_type="string"></object_data>
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
					<filter_data filter_name="COURSE_ID" filter_id="ReqBrokerReqObj.request_data.paid_courses[i]" filter_data_type="string" filter_condition="AND"></filter_data>
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
				<object_data key="course_id"  data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.paid_courses[i]" data_type="string"></object_data>
				<object_data key="section_id"  data_source_type="reqbrokermap" data_source="lms_course_section_table_list[j].id" data_type="string"></object_data>
				 <object_data key="is_enabled"  data_source_type="raw_value" data_source="false" data_type="Boolean"></object_data>
				<object_data key="percentage" data_source_type="raw_value" data_source="0" data_type="string"></object_data>
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
		 <exec_function function_name="insert_lms_course_progress_copy_object" function_type="copy_object" function_condition="lms_course_section_table_list[j].order==1">
			<copy_object_info object_name="insert_lms_course_progress_additional_obj"><object_data key="user_id" data_source_type="reqbrokermap" data_source="lms_user_data_table_record.id" data_type="string"></object_data>
				<object_data key="course_id"  data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.paid_courses[i]" data_type="string"></object_data>
				<object_data key="section_id"  data_source_type="reqbrokermap" data_source="lms_course_section_table_list[j].id" data_type="string"></object_data>
				 <object_data key="is_enabled"  data_source_type="raw_value" data_source="true" data_type="Boolean"></object_data>
				<object_data key="percentage" data_source_type="raw_value" data_source="0" data_type="string"></object_data>
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
					<filter_data filter_name="COURSE_ID" filter_id="ReqBrokerReqObj.request_data.paid_courses[i]" filter_data_type="string" filter_condition="AND"></filter_data>
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
				<object_data key="course_id"  data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.paid_courses[i]" data_type="string"></object_data>
				<object_data key="section_id"  data_source_type="reqbrokermap" data_source="lms_course_section_table_list[j].id" data_type="string"></object_data>
                <object_data key="item_id"  data_source_type="reqbrokermap" data_source="lms_course_item_table_list[k].id" data_type="string"></object_data>
				<object_data key="is_enabled"  data_source_type="raw_value" data_source="false" data_type="Boolean"></object_data>
				<object_data key="percentage" data_source_type="raw_value" data_source="0" data_type="string"></object_data>
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
		<exec_function function_name="insert_lms_course_progress_copy_object" function_type="copy_object" function_condition="lms_course_item_table_list[k].order==1">
			<copy_object_info object_name="insert_lms_course_progress_additional_obj">
				<object_data key="user_id" data_source_type="reqbrokermap" data_source="lms_user_data_table_record.id" data_type="string"></object_data>
				<object_data key="course_id"  data_source_type="reqbrokermap" data_source="ReqBrokerReqObj.request_data.paid_courses[i]" data_type="string"></object_data>
				<object_data key="section_id"  data_source_type="reqbrokermap" data_source="lms_course_section_table_list[j].id" data_type="string"></object_data>
                <object_data key="item_id"  data_source_type="reqbrokermap" data_source="lms_course_item_table_list[k].id" data_type="string"></object_data>
				 <object_data key="is_enabled"  data_source_type="raw_value" data_source="true" data_type="Boolean"></object_data>
				<object_data key="percentage" data_source_type="raw_value" data_source="0" data_type="string"></object_data>
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

				<exec_function function_name="lms_paid_course_loop" function_type="end_loop" index_name="i" />

			 <exec_function function_name="insert_lms_user_token_info_copy_object" function_type="copy_object">
				<copy_object_info object_name="insert_lms_user_token_info_table_obj">
					<object_data key="access_token" data_source_type="raw_value" data_source="NA" data_type="string"></object_data>
					<object_data key="refresh_token" data_source_type="raw_value" data_source="NA" data_type="string"></object_data>
				</copy_object_info>
			</exec_function>
			<exec_function function_name="insert_lms_token_info_db_insert" function_type="db_insert">
				<db_insert_info>
					<table_name>LMS_TOKEN_INFO</table_name>
					<insert_data_list>
					    <insert_data insert_data_name="USER_ID" insert_db_data_type="string" insert_data_source="ReqBrokerReqObj.request_data.email" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="ACCESS_TOKEN" insert_db_data_type="string" insert_data_source="insert_lms_user_token_info_table_obj.access_token" insert_data_type="string"></insert_data>
						<insert_data insert_data_name="REFRESH_TOKEN" insert_db_data_type="string" insert_data_source="insert_lms_user_token_info_table_obj.refresh_token" insert_data_type="string"></insert_data>
						</insert_data_list>
				</db_insert_info>
			</exec_function>
	</exec_group>
	<exec_group group_name="read_lms_enquiry_db_multi_read"  group_condition="ReqBrokerReqObj.header_data.channel=='web_admin_portal'">
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
