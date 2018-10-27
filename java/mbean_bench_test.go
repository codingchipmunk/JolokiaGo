package java

import (
	"encoding/json"
	"testing"
)

//	BenchmarkMBean_MarshalText benchmarks the MarshalText function with the mbeanTextMarshalingSets testsets
func BenchmarkMBean_MarshalText(b *testing.B) {
	testSets := mbeanTextMarshalingSets
	b.ResetTimer()
	for _, testSet := range testSets {
		b.Run(testSet.Name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				testSet.MBean.MarshalText()
			}

		})
	}
}

//	BenchmarkMBean_UnMarshalText benchmarks the UnMarshalText function with the mbeanTextMarshalingSets testsets
func BenchmarkMBean_UnMarshalText(b *testing.B) {
	testSets := mbeanTextMarshalingSets
	b.ResetTimer()
	for _, testSet := range testSets {
		bt := []byte(testSet.Text)
		b.Run(testSet.Name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				testSet.MBean.UnmarshalText(bt)
			}

		})
	}
}

//	BenchmarkMBean_MarshalJSON benchmarks the marshaling of an mbean with json.Marshal
//	Uses the mbeanTextMarshalingSets testsets
func BenchmarkMBean_MarshalJSON(b *testing.B) {
	testSets := mbeanTextMarshalingSets
	b.ResetTimer()
	for _, testSet := range testSets {
		b.Run(testSet.Name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				json.Marshal(&testSet.MBean)
			}
		})
	}
}
