package strings

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		in       string
		expected bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{"А роза упала на лапу Азора", true},
	}

	for _, tc := range testCases {
		got := IsPalindrome(tc.in)
		if got != tc.expected {
			t.Errorf("IsPalindrome() = %v; want %v", got, tc.expected)
		}
	}
}

func TestReverseStrings(t *testing.T) {

	testCases := []struct {
		in       []string
		expected []string
	}{
		{
			[]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"},
			[]string{"9", "8", "7", "6", "5", "4", "3", "2", "1", "0"},
		},

		{
			[]string{"h", "e", "l", "l", "o", " ", "w", "o", "r", "l", "d"},
			[]string{"d", "l", "r", "o", "w", " ", "o", "l", "l", "e", "h"},
		},
	}

	for _, tc := range testCases {
		t.Run("ReverseStrings", func(t *testing.T) {
			got := make([]string, len(tc.in))
			copy(got, tc.in)

			ReverseStrings(&got)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("ReverseStrings() = %v; want %v", got, tc.expected)
			}
		})

		t.Run("ReverseStrings2", func(t *testing.T) {
			got := make([]string, len(tc.in))
			copy(got, tc.in)

			ReverseStrings2(&got)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("ReverseStrings2() = %v; want %v", got, tc.expected)
			}
		})

	}
}

func TestReverse(t *testing.T) {

	testCases := []struct {
		in       string
		expected string
	}{
		{"0123456789", "9876543210"},
		{"hello world", "dlrow olleh"},
		{"racecar", "racecar"},
	}

	for _, tc := range testCases {
		t.Run("Reverse", func(t *testing.T) {

			got := Reverse(tc.in)
			if got != tc.expected {
				t.Errorf("Reverse() = %v; want %v", got, tc.expected)
			}
		})

		t.Run("Reverse1", func(t *testing.T) {

			got := Reverse1(tc.in)
			if got != tc.expected {
				t.Errorf("Reverse1() = %v; want %v", got, tc.expected)
			}
		})

		t.Run("Reverse2", func(t *testing.T) {

			got := Reverse2(tc.in)
			if got != tc.expected {
				t.Errorf("Reverse2() = %v; want %v", got, tc.expected)
			}
		})

		t.Run("Reverse3", func(t *testing.T) {

			got := Reverse3(tc.in)
			if got != tc.expected {
				t.Errorf("Reverse3() = %v; want %v", got, tc.expected)
			}
		})

		t.Run("ReverseSimple", func(t *testing.T) {

			got := ReverseSimple(tc.in)
			if got != tc.expected {
				t.Errorf("ReversSimple() = %v; want %v", got, tc.expected)
			}
		})

	}
}

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
BenchmarkDuplicates/Duplicates-4                     375           3077509 ns/op              16 B/op          1 allocs/op
BenchmarkDuplicates/Duplicates1-4                    363           3292199 ns/op              16 B/op          1 allocs/op
BenchmarkDuplicates/DuplicatesC-4                     27          45669278 ns/op         8084900 B/op       3985 allocs/op
PASS
ok      github.com/GarryGaller/goalgo/strings   4.867s
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

	b.Run("Unique1", func(b *testing.B) {
		data := make([]string, len(arr))
		copy(data, arr)
		sort.Strings(data)
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			_ = Unique1(data)
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
BenchmarkUnique/Unique-4                      44          28842532 ns/op         9247352 B/op         30 allocs/op
BenchmarkUnique/Unique1-4                     56          25983629 ns/op         9247351 B/op         30 allocs/op
BenchmarkUnique/Unique2-4                     54          25668115 ns/op         9247353 B/op         30 allocs/op
BenchmarkUnique/UniqueC-4                     18          60670133 ns/op        17332389 B/op       4015 allocs/op
PASS
ok      github.com/GarryGaller/goalgo/strings   6.097s
*/
