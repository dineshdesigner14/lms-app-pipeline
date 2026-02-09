package securitydef

type GenLMKReqInfo struct {
	SSMSecurityID string `json:"ssm_group_id"`
	Component1    string `json:"component_1"`
	Component2    string `json:"component_2"`
}

type GenTMKReqInfo struct {
	SSMSecurityID string `json:"ssm_group_id"`
	Component1    string `json:"component_1"`
	Component2    string `json:"component_2"`
}

type GenPVKReqInfo struct {
	SSMSecurityID string `json:"ssm_group_id"`
	Component1    string `json:"component_1"`
	Component2    string `json:"component_2"`
}

type GenCVKReqInfo struct {
	SSMSecurityID string `json:"ssm_group_id"`
	Component1    string `json:"component_1"`
	Component2    string `json:"component_2"`
}

type GenZMKReqInfo struct {
	SSMSecurityID string `json:"ssm_group_id"`
	Component1    string `json:"component_1"`
	Component2    string `json:"component_2"`
}

type GenKEKReqInfo struct {
	SSMSecurityID string `json:"ssm_group_id"`
	Component1    string `json:"component_1"`
	Component2    string `json:"component_2"`
}

type GenDEKReqInfo struct {
	SSMSecurityID string `json:"ssm_group_id"`
	Component1    string `json:"component_1"`
	Component2    string `json:"component_2"`
}

type ExportKeyReqInfo struct {
	SSMSecurityID string `json:"ssm_group_id"`
	Component1    string `json:"component_1"`
	Component2    string `json:"component_2"`
	KEK           string `json:"kek"`
}
