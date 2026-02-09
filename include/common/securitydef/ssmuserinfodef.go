package securitydef

const (
	ADMIN1_USER_ID = "admin1"
	ADMIN2_USER_ID = "admin2"
)

const (
	MAX_PIN_RETRY = 3
)

const (
	SSM_USER_STATUS_FIRST_TIME_LOGIN    = "FirstTimeLogin"
	SSM_USER_STATUS_TEMP_BLOCK          = "TempBlock"
	SSM_USER_STATUS_FIRST_ACTIVE        = "Active"
	SSM_USER_STATUS_PIN_RETRY_EXECEEDED = "PINRetryExceeded"
)

const (
	SSM_RESP_MSG_ADM1_FIRST_TIME_LOGIN = "Adm1FirstTimeLogin"
	SSM_RESP_MSG_ADM2_FIRST_TIME_LOGIN = "Adm2FirstTimeLogin"
)
