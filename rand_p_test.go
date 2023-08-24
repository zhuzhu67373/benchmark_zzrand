package benchmark_zzrand

import (
	"github.com/valyala/fastrand"
	"github.com/zhuzhu67373/zzrand"
	"math/rand"
	"testing"
)

func Benchmark_ZzrandUint32P(b *testing.B) {
	b.SetParallelism(10)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			zzrand.Uint32()
		}
	})
}
func Benchmark_FastrandUint32P(b *testing.B) {
	b.SetParallelism(10)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			zzrand.Uint32()
		}
	})
}
func Benchmark_StdrandUint32P(b *testing.B) {
	b.SetParallelism(10)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rand.Uint32()
		}
	})
}

func Benchmark_ZzrandIntP(b *testing.B) {
	b.SetParallelism(10)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			zzrand.Int()
		}
	})
}
func Benchmark_FastrandIntP(b *testing.B) {
	b.SetParallelism(10)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = int(fastrand.Uint32())
		}
	})
}
func Benchmark_StdrandIntP(b *testing.B) {
	b.SetParallelism(10)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rand.Int()
		}
	})
}

func Benchmark_ZzrandIntnP(b *testing.B) {
	b.SetParallelism(10)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			zzrand.Intn(100)
		}
	})
}
func Benchmark_FastrandIntnP(b *testing.B) {
	b.SetParallelism(10)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = int(fastrand.Uint32n(100))
		}
	})
}
func Benchmark_StdrandIntnP(b *testing.B) {
	b.SetParallelism(10)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rand.Intn(100)
		}
	})
}
