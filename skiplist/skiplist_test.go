package structure_benchmark

import (
	"testing"

	"github.com/huandu/skiplist"

	"github.com/davecgh/go-spew/spew"
)

func benchmarkremove(b *testing.B, size int, execCount int) {
	list := skiplist.New(skiplist.Int)
	l := loadTestData(size)

	b.N = len(l) * execCount

	b.ResetTimer()
	b.StartTimer()

	for c := 0; c < execCount; c++ {
		b.StopTimer()
		for _, v := range l {
			list.Set(v, v)
		}
		b.StartTimer()
		for i := 0; i < len(l); i++ {
			list.Remove(l[i])
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
	list := skiplist.New(skiplist.Int)
	l := loadTestData(size)

	for _, v := range l {
		list.Set(v, v)
	}

	b.N = len(l) * execCount

	b.ResetTimer()
	b.StartTimer()

	for c := 0; c < execCount; c++ {
		for i := 0; i < len(l); i++ {
			list.Get(l[i])
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
		list := skiplist.New(skiplist.Int)
		b.StartTimer()
		for _, v := range l {
			list.Set(v, v)
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
