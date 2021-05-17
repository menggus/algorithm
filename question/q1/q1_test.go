package main

import "testing"

func Benchmark_SearchFirstStrPosition(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SearchFirstStrPosition("diuahsdiushaiulh你好呀i而sa恩爱恩爱哈爱仕达书丢啊是啊上帝啊u上帝啊是你12213haihaidahsdiahsdiqhweihqwiehqiwhiahdannad阿的江啊还是", "d阿的")
	}
}

func Benchmark_SearchFirstStrPosition2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SearchFirstStrPosition("diuahsdiushaiulh你好呀i而sa恩爱恩爱哈爱仕达书丢啊是啊上帝啊u上帝啊是你12213haihaidahsdiahsdiqhweihqwiehqiwhiahdannad阿的江啊还是", "d阿的")
	}
}
