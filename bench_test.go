package main

import (
	"testing"
)

func BenchmarkWithoutPool(b *testing.B) {
	var cc *Counter
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			cc = &Counter{A: 0, B: 1}
			b.StopTimer()
			cc.increment()
			b.StartTimer()
		}
	}
}

func BenchmarkWithPool(b *testing.B) {
	var cc *Counter
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			cc = counterPool.Get().(*Counter)
			b.StopTimer()
			cc.A++
			cc.B++
			b.StartTimer()
			counterPool.Put(cc)
		}
	}
}
