package structure_benchmark

import (
	"testing"

	"github.com/davecgh/go-spew/spew"

	"474420502.top/eson/structure/compare"
	"474420502.top/eson/structure/vbt"
)

func benchmarkremove(b *testing.B, size int, execCount int) {
	tree := vbt.New(compare.Int)
	l := loadTestData(size)

	b.N = len(l) * execCount

	b.ResetTimer()
	b.StartTimer()

	for c := 0; c < execCount; c++ {
		b.StopTimer()
		for _, v := range l {
			tree.Put(v)
		}
		b.StartTimer()
		for i := 0; i < len(l); i++ {
			tree.Remove(l[i])
		}
	}
}

func BenchmarkRemove(b *testing.B) {

	for _, size := range BenchmarkSizeList {
		var execount = 5
		if _execount := BenchmarkMinSize / size; _execount > 5 {
			execount = _execount
		}
		b.Run(spew.Sprintf("%d", size), func(b *testing.B) {
			benchmarkremove(b, size, execount)
		})
	}
}

func benchmarkget(b *testing.B, size int, execCount int) {
	tree := vbt.New(compare.Int)
	l := loadTestData(size)

	for _, v := range l {
		tree.Put(v)
	}

	b.N = len(l) * execCount

	b.ResetTimer()
	b.StartTimer()

	for c := 0; c < execCount; c++ {
		for i := 0; i < len(l); i++ {
			tree.Get(l[i])
		}
	}
}

func BenchmarkGet(b *testing.B) {
	for _, size := range BenchmarkSizeList {
		var execount = 5
		if _execount := BenchmarkMinSize / size; _execount > 5 {
			execount = _execount
		}
		b.Run(spew.Sprintf("%d", size), func(b *testing.B) {
			benchmarkget(b, size, execount)
		})
	}
}

func benchmarkput(b *testing.B, size int, execCount int) {

	l := loadTestData(size)
	b.N = len(l) * execCount

	b.ResetTimer()
	b.StartTimer()

	for c := 0; c < execCount; c++ {
		b.StopTimer()
		tree := vbt.New(compare.Int)
		b.StartTimer()
		for _, v := range l {
			tree.Put(v)
		}
	}
}

func BenchmarkPut(b *testing.B) {

	for _, size := range BenchmarkSizeList {
		var execount = 5
		if _execount := BenchmarkMinSize / size; _execount > 5 {
			execount = _execount
		}
		b.Run(spew.Sprintf("%d", size), func(b *testing.B) {
			benchmarkput(b, size, execount)
		})
	}
	// b.Log(tree.count)
}
