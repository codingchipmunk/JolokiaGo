package jolokiaClient

import "bytes"

//	Represents an MBean with its respective attributes. Implements the TextMarshaler and TextUnmarshaler interface.
//	When marsahlled it will omit fields with empty strings
type MBean struct {
	Domain  string `json:"domain,omitempty"`
	Context string `json:"context,omitempty"`
	Type    string `json:"type,omitempty"`
	Name    string `json:"name,omitempty"`
}


func (bean *MBean) MarshalText() ([]byte, error) {
	var buff bytes.Buffer
	if bean.Domain != "" {
		buff.WriteString(bean.Domain)
		buff.WriteByte(':')
	}
	if bean.Context != "" {
		buff.WriteString("context=")
		buff.WriteString(bean.Context)
		buff.WriteByte(',')
	}
	if bean.Type != "" {
		buff.WriteString("type=")
		buff.WriteString(bean.Type)
		buff.WriteByte(',')
	}
	if bean.Name != "" {
		buff.WriteString("name=")
		buff.WriteString(bean.Name)
		buff.WriteByte(',')
	}
	if buff.Len() == 0 {
		return buff.Bytes(), nil
	}
	return buff.Bytes()[:buff.Len()-1], nil
}
