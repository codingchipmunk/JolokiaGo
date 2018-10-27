package jolokiago

import "bytes"

//	Represents an MBean with its respective attributes. Implements the TextMarshaler and TextUnmarshaler interface.
//	When marsahlled it will omit fields with empty strings
type MBean struct {
	Domain  string `json:"domain,omitempty"`
	Context string `json:"context,omitempty"`
	Type    string `json:"type,omitempty"`
	Name    string `json:"name,omitempty"`
}

const (
	contextIdentifier   = "context"
	typeIdentifier      = "type"
	nameIdentifier      = "name"
	domainSeperator     = ':'
	attributeSeperator  = ','
	attributeAssignment = '='
)

//	Marsahlls an MBean into a form of Domain:context=bean.context,type=bean.type,name=bean.name
//	Empty fields will be ignored when marshalled
func (bean *MBean) MarshalText() ([]byte, error) {
	var buff bytes.Buffer

	// Omit writing a domain if its empty
	if bean.Domain != "" {
		buff.WriteString(bean.Domain)
		buff.WriteByte(domainSeperator)
	}

	// Omit empty attributes
	if bean.Context != "" {
		writeAttribute(&buff, contextIdentifier, bean.Context)
	}
	if bean.Type != "" {
		writeAttribute(&buff, typeIdentifier, bean.Type)
	}
	if bean.Name != "" {
		writeAttribute(&buff, nameIdentifier, bean.Name)
	}

	// if the buffer is empty, return it without modifications
	if buff.Len() == 0 {
		return buff.Bytes(), nil
	}

	// return the bytes in the buffer without the trailing seperator
	return buff.Bytes()[:buff.Len()-1], nil
}

//	Calls a set of functions to write the attribute name and its value into the buffer
func writeAttribute(buff *bytes.Buffer, attributeName string, attributeValue string) {
	buff.WriteString(attributeName)
	buff.WriteByte(attributeAssignment)
	buff.WriteString(attributeValue)
	buff.WriteByte(attributeSeperator)
}

//	Unmarshals a text into an MBean object
func (bean *MBean) UnmarshalText(text []byte) error {
	mainSplit := bytes.SplitN(text, []byte{domainSeperator}, 2)
	var attribSlice []byte
	bean.Domain = string(mainSplit[0])
	if len(mainSplit) > 1 {
		attribSlice = mainSplit[1]
	} else {
		return nil
	}

	attribSplit := bytes.Split(attribSlice, []byte{attributeSeperator})
	for i := range attribSplit {
		bean.extractAttribute(attribSplit[i])
	}

	return nil
}

func (bean *MBean) extractAttribute(text []byte) {
	split := bytes.SplitN(text, []byte{attributeAssignment}, 2)
	attrName := split[0]
	if bytes.Equal(attrName, []byte(contextIdentifier)) {
		bean.Context = string(split[1])
	} else if bytes.Equal(attrName, []byte(typeIdentifier)) {
		bean.Type = string(split[1])
	} else if bytes.Equal(attrName, []byte(nameIdentifier)) {
		bean.Name = string(split[1])
	}
}
