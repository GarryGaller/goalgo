package numbers

import (
	//"fmt"
	"math/big"
	"testing"
)

func TestFib(t *testing.T) {
	fib := &Fib{}

	fib100 := big.NewInt(0)
	fib100.SetString("354224848179261915075", 10)

	var testCases = []struct {
		in       uint
		expected *big.Int
	}{
		{5, big.NewInt(5)},
		{10, big.NewInt(55)},
		{25, big.NewInt(75025)},
		{100, fib100},
	}

	for _, tc := range testCases {

		got, _ := fib.Calc(tc.in, true)

		if got.Cmp(tc.expected) != 0 {
			t.Errorf("fib.Calc(%d, true) = %v; want %v", tc.in, got, tc.expected)
		}
	}
}

func BenchmarkFib10th(b *testing.B) {

	for i := 0; i < b.N; i++ {
		fib := &Fib{}
		//start := time.Now()
		_, err := fib.Calc(uint(10_000), true)
		if err != nil {
			b.Errorf("i = %d - %v", i, err.Error())
		}
		//duration := time.Since(start)
		//fmt.Printf("%s => %d\n", duration, n)
	}
}

func BenchmarkFib100th(b *testing.B) {
	fib := &Fib{}
	for i := 0; i < b.N; i++ {
		//start := time.Now()
		_, err := fib.Calc(uint(100_000), true)
		if err != nil {
			b.Errorf("i = %d - %v", i, err.Error())
		}
		//duration := time.Since(start)
		//fmt.Printf("%s => %d\n", duration, n)
	}
}

func BenchmarkFib1mln(b *testing.B) {
	fib := &Fib{}
	for i := 0; i < b.N; i++ {
		//start := time.Now()
		_, err := fib.Calc(uint(1_000_000), true)
		if err != nil {
			b.Errorf("i = %d - %v", i, err.Error())
		}
		//duration := time.Since(start)
		//fmt.Printf("%s => %d\n", duration, n)
	}
}

func BenchmarkFib10mln(b *testing.B) {
	fib := &Fib{}
	for i := 0; i < b.N; i++ {
		//start := time.Now()
		_, err := fib.Calc(uint(10_000_000), true)
		if err != nil {
			b.Errorf("i = %d - %v", i, err.Error())
		}
		//duration := time.Since(start)
		//fmt.Printf("%s => %d\n", duration, n)
	}
}

func BenchmarkFib100mln(b *testing.B) {
	fib := &Fib{}
	for i := 0; i < b.N; i++ {
		//start := time.Now()
		_, err := fib.Calc(uint(100_000_000), true)
		if err != nil {
			b.Errorf("i = %d - %v", i, err.Error())
		}
		//duration := time.Since(start)
		//fmt.Printf("%s => %d\n", duration, n)
	}
}
