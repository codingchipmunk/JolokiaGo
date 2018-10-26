package JolokiaGo


const domainVal = "test.domain"
const contextVal = "testingContext"
const typeVal = "testingType"
const nameVal = "testingName"
const d = ":"
const s = ","
const attAss = "="


const domainT = domainVal + d
const typeT = "type" + attAss + typeVal
const nameT = "name" + attAss + nameVal
const contT = "context" + attAss + contextVal

var mbeanTextMarshalingSets = []struct {
	Name  string
	MBean MBean
	Text  string
}{
	{"Full MBean", MBean{Domain: domainVal, Context: contextVal, Type: typeVal, Name: nameVal}, domainT + contT + s + typeT + s + nameT},
	{"MBean with missing domain", MBean{Context: contextVal, Type: typeVal, Name: nameVal}, contT + s + typeT + s + nameT},
	{"MBean with one missing attribute [name]", MBean{Domain: domainVal, Context: contextVal, Type: typeVal}, domainT+ contT + s + typeT},
	{"MBean with one missing attribute [context]", MBean{Domain: domainVal, Name: nameVal, Type: typeVal}, domainT + typeT + s + nameT},
	{"MBean with one missing attribute [type]", MBean{Domain: domainVal, Context: contextVal, Name: nameVal}, domainT + contT + s + nameT},
	{"MBean with one attribute [name]", MBean{Domain: domainVal, Name: nameVal}, domainT + nameT},
	{"MBean with one attribute [context]", MBean{Domain: domainVal, Context: contextVal}, domainT + contT},
	{"MBean with one attribute [type]", MBean{Domain: domainVal, Type: typeVal}, domainT + typeT},
	{"MBean without attributes", MBean{Domain: domainVal}, domainVal},
	{"Empty MBean", MBean{}, ""},
}
