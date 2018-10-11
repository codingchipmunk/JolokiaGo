package jolokiaClient

type BaseRequest struct {
	Type    string `json:"type"`
	Command string `json:"command,omitempty"`
	MBean   *MBean `json:"mbean,omitempty"`
}

type RegisterEventRequest struct {
	BaseRequest
	Mode string `json:"mode"`
	ClientID string `json:"client"`
}