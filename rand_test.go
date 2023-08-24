package benchmark_zzrand

import (
	"github.com/valyala/fastrand"
	"github.com/zhuzhu67373/zzrand"
	"testing"
)

func Benchmark_ZzrandUint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		zzrand.Uint32()
	}
}
func Benchmark_FastrandUint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fastrand.Uint32()
	}
}

func Benchmark_ZzrandInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		zzrand.Int()
	}
}
func Benchmark_FastrandInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = int(fastrand.Uint32())
	}
}

func Benchmark_ZzrandIntn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		zzrand.Intn(100)
	}
}
func Benchmark_FastrandIntn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = int(fastrand.Uint32n(100))
	}
}
