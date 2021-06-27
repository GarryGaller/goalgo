package numbers_test

import (
    "fmt"
    "math/big"
    "testing"
    "github.com/GarryGaller/goalgo/numbers"
)

func TestFib(t *testing.T) {
    fib := &numbers.Fib{}
    //-----------------
    got, _ := fib.Calc(uint(5), true)
    fmt.Println(big.NewInt(5) == big.NewInt(5))
    if got.Cmp(big.NewInt(5)) != 0 {
        t.Errorf("fib.Calc(uint(5), true) = %d; want 5", got)
    }
    
    //-----------------
    got, _ = fib.Calc(uint(10), true)
    if got.Cmp(big.NewInt(55)) != 0 {
        t.Errorf("fib.Calc(uint(10), true) = %d; want 55", got)
    }
    //-----------------
    got, _ = fib.Calc(uint(25), true)
    if got.Cmp(big.NewInt(75025)) != 0 {
        t.Errorf("fib.Calc(uint(25), true) = %d; want 75025", got)
    }
    
    //-----------------
    got, _ = fib.Calc(uint(100), true)
    b:= new(big.Int)
    b.SetString("354224848179261915075",10)
    if got.Cmp(b) != 0 {
        t.Errorf("fib.Calc(uint(100), true) = %d; want 354224848179261915075", got)
    }
    
    
}


func BenchmarkFib10th(b *testing.B) {
    
    for i := 0; i < b.N; i++ {
        fib := &numbers.Fib{}
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
    fib := &numbers.Fib{}
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
    fib := &numbers.Fib{}
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
    fib := &numbers.Fib{}
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
    fib := &numbers.Fib{}
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

