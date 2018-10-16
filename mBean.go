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

const Context_identifier = "context"
const Type_identifier = "type"
const Name_identifier = "name"
const Domain_seperator = ':'
const Attribute_seperator = ','
const Attribute_assignment = '='

//	Marsahlls an MBean into a form of Domain:context=bean.context,type=bean.type,name=bean.name
//	Empty fields will be ignored when marshalled
func (bean *MBean) MarshalText() ([]byte, error) {
	var buff bytes.Buffer

	// Omit writing a domain if its empty
	if bean.Domain != "" {
		buff.WriteString(bean.Domain)
		buff.WriteByte(Domain_seperator)
	}

	// Omit empty attributes
	if bean.Context != "" {
		writeAttribute(&buff, Context_identifier, bean.Context)
	}
	if bean.Type != "" {
		writeAttribute(&buff, Type_identifier, bean.Type)
	}
	if bean.Name != "" {
		writeAttribute(&buff, Name_identifier, bean.Name)
	}

	// if the buffer is empty, return it without modifications
	if buff.Len() == 0 {
		return buff.Bytes(), nil
	}

	// return the bytes in the buffer without the trailing seperator
	return buff.Bytes()[:buff.Len()-1], nil
}

//	Calls a set of functions to write the attribute name and its value into the buffer
func writeAttribute(buff *bytes.Buffer, attribute_name string, attribute_value string) {
	buff.WriteString(attribute_name)
	buff.WriteByte(Attribute_assignment)
	buff.WriteString(attribute_value)
	buff.WriteByte(Attribute_seperator)
}