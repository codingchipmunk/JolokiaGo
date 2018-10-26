package JolokiaGo

//Values for the MBean Fields
const (
	domainVal  = "test.domain"
	contextVal = "testingContext"
	typeVal    = "testingType"
	nameVal    = "testingName"
)
//Separators and assignments
const(
	//ds (domain separator) separates domain from attributes
	ds = ":"
	//s (seperator) separates multiple attribute assignments
	s          = ","
	//attAss (attribute assignment) assignes a value to an attribute
	attAss     = "="
)
//Concatenated fields for convenient use in the Text-Field
const(
	domainT  = domainVal + ds
	typeT    = "type" + attAss + typeVal
	nameT    = "name" + attAss + nameVal
	contextT = "context" + attAss + contextVal
)

//mbeanTextMarshalingSets holds the testsets for testing text marshaling
var mbeanTextMarshalingSets = []struct {
	Name  string
	MBean MBean
	Text  string
}{
	{"Full MBean", MBean{Domain: domainVal, Context: contextVal, Type: typeVal, Name: nameVal}, domainT + contextT + s + typeT + s + nameT},
	{"MBean with missing domain", MBean{Context: contextVal, Type: typeVal, Name: nameVal}, contextT + s + typeT + s + nameT},
	{"MBean with one missing attribute [name]", MBean{Domain: domainVal, Context: contextVal, Type: typeVal}, domainT + contextT + s + typeT},
	{"MBean with one missing attribute [context]", MBean{Domain: domainVal, Name: nameVal, Type: typeVal}, domainT + typeT + s + nameT},
	{"MBean with one missing attribute [type]", MBean{Domain: domainVal, Context: contextVal, Name: nameVal}, domainT + contextT + s + nameT},
	{"MBean with one attribute [name]", MBean{Domain: domainVal, Name: nameVal}, domainT + nameT},
	{"MBean with one attribute [context]", MBean{Domain: domainVal, Context: contextVal}, domainT + contextT},
	{"MBean with one attribute [type]", MBean{Domain: domainVal, Type: typeVal}, domainT + typeT},
	{"MBean without attributes", MBean{Domain: domainVal}, domainVal},
	{"Empty MBean", MBean{}, ""},
}
