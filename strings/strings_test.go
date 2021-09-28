package strings

import (
	"fmt"
	"testing"
)

func BenchmarkReverseStrings(b *testing.B) {

	data := make([]string, 100000)
	for i := 0; i < 100000; i++ {
		data = append(data, fmt.Sprint(i))
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ReverseStrings(&data)
	}
}

func BenchmarkReverseStrings2(b *testing.B) {

	data := make([]string, 100000)
	for i := 0; i < 100000; i++ {
		data = append(data, fmt.Sprint(i))
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ReverseStrings2(&data)
	}
}
