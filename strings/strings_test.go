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

func BenchmarkReverse(b *testing.B) {

	data := RandomString(100000)

	b.ResetTimer()

	b.Run("Reverse", func(b *testing.B) {

		for i := 0; i < b.N; i++ {
			Reverse(data)
		}

	})

	b.Run("Reverse1", func(b *testing.B) {

		for i := 0; i < b.N; i++ {
			Reverse1(data)
		}

	})

	b.Run("Reverse2", func(b *testing.B) {

		for i := 0; i < b.N; i++ {
			Reverse2(data)
		}

	})

	b.Run("Reverse3", func(b *testing.B) {

		for i := 0; i < b.N; i++ {
			Reverse3(data)
		}

	})

	b.Run("ReverseSimple", func(b *testing.B) {

		for i := 0; i < b.N; i++ {
			ReverseSimple(data)
		}

	})

}  

// benchmarks
//go test -bench=. -benchmem

/*
BenchmarkReverseStrings-4            685           1791343 ns/op               0 B/op          0 allocs/op
BenchmarkReverseStrings2-4           560           1905460 ns/op               0 B/op          0 allocs/op
BenchmarkReverse/Reverse-4           547           2188425 ns/op          507904 B/op          2 allocs/op
BenchmarkReverse/Reverse1-4          542           2215994 ns/op          507905 B/op          2 allocs/op
BenchmarkReverse/Reverse2-4          535           2246857 ns/op          507904 B/op          2 allocs/op
BenchmarkReverse/Reverse3-4          334           3596014 ns/op          212992 B/op          2 allocs/op
BenchmarkReverse/ReverseSimple-4     310           3871185 ns/op         2801410 B/op         30 allocs/op
PASS
ok      github.com/GarryGaller/goalgo/strings   10.592s
*/
