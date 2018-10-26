package JolokiaGo

import (
	"testing"
)

//	TestMBean_MarshalText tests if the MarshalText function works correctly
//	Uses the MBeans from mbeanTextMarshalingSets and compares the result to the Text field
func TestMBean_MarshalText(t *testing.T) {
	for _, testset := range mbeanTextMarshalingSets {
		t.Run(testset.Name, func(t *testing.T) {
			got, err := testset.MBean.MarshalText()
			if err != nil {
				t.Error("Unexpected error: " + err.Error())
				t.FailNow()
			}
			gotStr := string(got)
			if gotStr != testset.Text {
				t.Error("Got unexpected string: " + gotStr)
				t.FailNow()
			}
		})
	}
}

//	TestMBean_UnmarshalText tests if the UnmarshalText function works correctly.
//	Uses the Text - Value from mbeanTextMarshalingSets and compares the result to the MBean stored there
func TestMBean_UnmarshalText(t *testing.T) {
	for _, testset := range mbeanTextMarshalingSets {
		t.Run(testset.Name, func(t *testing.T) {
			bean := MBean{}
			err := bean.UnmarshalText([]byte(testset.Text))
			if err != nil {
				t.Error("Unexpected error: " + err.Error())
				t.FailNow()
			}

			if bean.Type != testset.MBean.Type {
				t.Error("Got unexpected Type: " + bean.Type)
				t.Fail()
			}
			if bean.Context != testset.MBean.Context {
				t.Error("Got unexpected Context: " + bean.Context)
				t.Fail()
			}
			if bean.Domain != testset.MBean.Domain {
				t.Error("Got unexpected Domain: " + bean.Domain)
				t.Fail()
			}
			if bean.Name != testset.MBean.Name {
				t.Error("Got unexpected Name: " + bean.Name)
				t.Fail()
			}
		})
	}
}
