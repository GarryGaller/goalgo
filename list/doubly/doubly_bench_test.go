package doublylinkedlist

import (
	"fmt"
	"testing"
)

func BenchmarkPushBack(b *testing.B) {

	list := New()

	for i := 0; i < 100000; i++ {
		list.PushBack(fmt.Sprint(i))
	}

}

func BenchmarkClear(b *testing.B) {

	list := New()

	for i := 0; i < 100000; i++ {
		list.PushBack(fmt.Sprint(i))
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		list.Clear()
	}
}

func BenchmarkClearOld(b *testing.B) {

	list := New()

	for i := 0; i < 100000; i++ {
		list.PushBack(fmt.Sprint(i))
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		list.ClearOld()
	}
}

// go test -bench=. -benchmem -cpuprofile=cpu.prof -memprofile=mem.prof
// go tool pprof -nodefraction=0 -svg list.test.exe mem.prof >mem.svg
