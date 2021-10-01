package strings

import (
	"fmt"
	"sort"
	"testing"
)

func BenchmarkReverseStrings(b *testing.B) {

	data := make([]string, 100000)
	for i := 0; i < 100000; i++ {
		data[i] = fmt.Sprint(i)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ReverseStrings(&data)
	}
}

func BenchmarkReverseStrings2(b *testing.B) {

	data := make([]string, 100000)
	for i := 0; i < 100000; i++ {
		data[i] = fmt.Sprint(i)
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
//go test -bench=Reverse -benchmem

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

func BenchmarkDeduplicate(b *testing.B) {

	N := 100000
	arr := make([]string, N)
	for i := 1; i < N-1; i++ {
		arr[i] = RandomString(10)
	}
	arr[0] = "aaaaaaaaaa"
	arr[N-1] = "aaaaaaaaaa"

	b.ResetTimer()

	b.Run("Deduplicate", func(b *testing.B) {
		data := make([]string, len(arr))
		copy(data, arr)
		sort.Strings(data)
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			Deduplicate(&data)
		}

	})

	b.Run("Deduplicate2", func(b *testing.B) {
		data := make([]string, len(arr))
		copy(data, arr)
		sort.Strings(data)
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			_ = Deduplicate2(data)
		}

	})

	b.Run("Deduplicate3", func(b *testing.B) {
		data := make([]string, len(arr))
		copy(data, arr)
		sort.Strings(data)
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			_ = Deduplicate3(data)
		}

	})

	b.Run("Deduplicate4", func(b *testing.B) {
		data := make([]string, len(arr))
		copy(data, arr)
		sort.Strings(data)
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			_ = Deduplicate4(data)
		}

	})

	b.Run("DeduplicateC", func(b *testing.B) {
		data := make([]string, len(arr))
		copy(data, arr)
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			_ = DeduplicateCounter(data)
		}

	})

}

// benchmarks
//go test -bench=Deduplicate -benchmem
/*

BenchmarkDeduplicate/Deduplicate-4                   340           3467846 ns/op               0 B/op          0 allocs/op
BenchmarkDeduplicate/Deduplicate2-4                  338           3470613 ns/op               0 B/op          0 allocs/op
BenchmarkDeduplicate/Deduplicate3-4                   54          28834983 ns/op         9247364 B/op         30 allocs/op
BenchmarkDeduplicate/Deduplicate4-4                   40          27551575 ns/op         9247361 B/op         30 allocs/op
BenchmarkDeduplicate/DeduplicateC-4                   19          60582416 ns/op        17332395 B/op       4015 allocs/op
PASS
ok      github.com/GarryGaller/goalgo/strings   7.872s


*/

func BenchmarkDuplicates(b *testing.B) {
	N := 100000
	arr := make([]string, N)
	for i := 1; i < N-1; i++ {
		arr[i] = RandomString(10)
	}
	arr[0] = "aaaaaaaaaa"
	arr[N-1] = "aaaaaaaaaa"

	b.ResetTimer()

	b.Run("Duplicates", func(b *testing.B) {
		data := make([]string, len(arr))
		copy(data, arr)
		sort.Strings(data)
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			_ = Duplicates(data)
		}

	})

	b.Run("Duplicates1", func(b *testing.B) {
		data := make([]string, len(arr))
		copy(data, arr)
		sort.Strings(data)
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			_ = Duplicates1(data)
		}

	})

	b.Run("Duplicates2", func(b *testing.B) {
		data := make([]string, len(arr))
		copy(data, arr)
		sort.Strings(data)
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			_ = Duplicates2(data)
		}

	})

	b.Run("DuplicatesC", func(b *testing.B) {
		data := make([]string, len(arr))
		copy(data, arr)
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			_ = DuplicatesCounter(data)
		}

	})
}

// benchmarks
//go test -bench=Duplicates -benchmem
/*
BenchmarkDuplicates/Duplicates-4                     372           3312017 ns/op              16 B/op          1 allocs/op
BenchmarkDuplicates/Duplicates1-4                    327           3446680 ns/op              16 B/op          1 allocs/op
BenchmarkDuplicates/Duplicates2-4                    354           3522801 ns/op              16 B/op          1 allocs/op
BenchmarkDuplicates/DuplicatesC-4                     24          46794342 ns/op         8086087 B/op       3991 allocs/op
PASS
ok      github.com/GarryGaller/goalgo/strings   7.515s
*/

func BenchmarkUnique(b *testing.B) {

	N := 100000
	arr := make([]string, N)
	for i := 1; i < N-1; i++ {
		arr[i] = RandomString(10)
	}
	arr[0] = "aaaaaaaaaa"
	arr[N-1] = "aaaaaaaaaa"
	b.ResetTimer()

	b.Run("Unique", func(b *testing.B) {
		data := make([]string, len(arr))
		copy(data, arr)
		sort.Strings(data)
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			_ = Unique(data)
		}

	})

	b.Run("Unique2", func(b *testing.B) {
		data := make([]string, len(arr))
		copy(data, arr)
		sort.Strings(data)
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			_ = Unique2(data)
		}

	})

	b.Run("UniqueC", func(b *testing.B) {
		data := make([]string, len(arr))
		copy(data, arr)
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			_ = UniqueCounter(data)
		}

	})
}

// benchmarks
//go test -bench=Unique -benchmem
/*
BenchmarkUnique/Unique-4                      42          30192202 ns/op         9247360 B/op         30 allocs/op
BenchmarkUnique/Unique2-4                     56          27430139 ns/op         9247352 B/op         30 allocs/op
BenchmarkUnique/UniqueC-4                     18          60559022 ns/op        17335775 B/op       4031 allocs/op
PASS
ok      github.com/GarryGaller/goalgo/strings   4.534s
*/
