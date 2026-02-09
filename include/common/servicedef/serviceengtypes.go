package servicedef

type ServiceEngReqMsg struct {
	Command     string `json:"command,omitempty"`
	ServiceName string `json:"service_name,omitempty"`
	ServiceID   string `json:"service_id,omitempty"`
	Status      string `json:"status,omitempty"`
	StartDate   string `json:"start_date,omitempty"`
	StartTime   string `json:"start_time,omitempty"`
	DebugLevel  string `json:"debug_level,omitempty"`
}

type ServiceEngTable struct {
	ServiceName string `json:"service_name,omitempty"`
	ServiceID   string `json:"service_id,omitempty"`
	Status      string `json:"status,omitempty"`
	StartDate   string `json:"start_date,omitempty"`
	StartTime   string `json:"start_time,omitempty"`
	DebugLevel  string `json:"debug_level,omitempty"`
	AddStatus   string `json:"add_status,omitempty"`
}
