package mBean

import "github.com/codingchipmunk/JolokiaGo"

const DomainVal = "test.domain"
const ContextVal = "testingContext"
const TypeVal = "testingType"
const NameVal = "testingName"


var MBean_text_marshaling_sets = []struct {
	Name  string
	MBean jolokiaClient.MBean
	Text  string
}{
	{"Full MBean", jolokiaClient.MBean{Domain: DomainVal, Context: ContextVal, Type: TypeVal, Name: NameVal}, "test.domain:context=testingContext,type=testingType,name=testingName"},
	{"MBean with missing domain", jolokiaClient.MBean{Context: "testingContext", Type: "testingType", Name: "testingName"}, "context=testingContext,type=testingType,name=testingName"},
	{"MBean with one missing attribute [name]", jolokiaClient.MBean{Domain: "test.domain", Context: "testingContext", Type: "testingType"}, "test.domain:context=testingContext,type=testingType"},
	{"MBean with one missing attribute [context]", jolokiaClient.MBean{Domain: "test.domain", Name: NameVal, Type: "testingType"}, "test.domain:type=testingType,name=testingName"},
	{"MBean with one missing attribute [type]", jolokiaClient.MBean{Domain: "test.domain", Context: "testingContext", Name: NameVal}, "test.domain:context=testingContext,name=testingName"},
	{"MBean with one attribute [name]", jolokiaClient.MBean{Domain: "test.domain", Name:NameVal}, "test.domain:name=testingName"},
	{"MBean with one attribute [context]", jolokiaClient.MBean{Domain: "test.domain", Context:ContextVal}, "test.domain:context=testingContext"},
	{"MBean with one attribute [type]", jolokiaClient.MBean{Domain: "test.domain", Type: "testingType"}, "test.domain:type=testingType"},
	{"MBean without attributes", jolokiaClient.MBean{Domain: "test.domain"}, "test.domain"},
	{"Empty MBean", jolokiaClient.MBean{}, ""},
}
