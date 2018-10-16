package mBean

import (
	"encoding/json"
	"testing"
)

//	Benchmarks the MarshalText function with the MBean_text_marshaling_sets testsets
func BenchmarkMBean_MarshalText(b *testing.B) {
	testSets := MBean_text_marshaling_sets
	b.ResetTimer()
	for _,testSet := range testSets{
		b.Run(testSet.Name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				testSet.MBean.MarshalText()
			}

		})
	}
}

//	Benchmarks the UnMarshalText function with the MBean_text_marshaling_sets testsets
func BenchmarkMBean_UnMarshalText(b *testing.B) {
	testSets := MBean_text_marshaling_sets
	b.ResetTimer()
	for _,testSet := range testSets{
		bt := []byte(testSet.Text)
		b.Run(testSet.Name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				testSet.MBean.UnmarshalText(bt)
			}

		})
	}
}

//	Benchmarks the marshaling of an MBean Object with json.Marshal
//	Uses the MBean_text_marshaling_sets testsets
func BenchmarkMBean_MarshalJSON(b *testing.B){
	testSets := MBean_text_marshaling_sets
	b.ResetTimer()
	for _,testSet := range testSets{
		b.Run(testSet.Name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				json.Marshal(&testSet.MBean)
			}
		})
	}
}
