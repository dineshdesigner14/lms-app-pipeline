package servicedef

const (
	NodeServiceSeparator = "."
	ServicePortSeparator = "_"
)

const (
	SERVICE_STATUS_RUNNING         = "running"
	SERVICE_STATUS_STOPPED         = "stopped"
	SERVICE_STATUS_INITIALIZED     = "initialized"
	SERVICE_STATUS_NOT_INITIALIZED = "not initialized"
)

const (
	ServiceEngCommand_StartService    = "StartService"
	ServiceEngCommand_StopService     = "StopService"
	ServiceEngCommand_Shutdown        = "Shutdown"
	ServiceEngCommand_StartServiceAck = "StartServiceAck"
	ServiceEngCommand_ListServices    = "ListServices"
	ServiceEngCommand_SetDebugLevel   = "SetDebugLevel"
)
