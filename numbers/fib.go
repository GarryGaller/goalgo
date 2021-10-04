package numbers

import (
	"errors"
	"math/big"
)

type BigIntOperators struct{}

func (b BigIntOperators) Mul(x, y *big.Int) *big.Int {
	return big.NewInt(0).Mul(x, y)
}
func (b BigIntOperators) Sub(x, y *big.Int) *big.Int {
	return big.NewInt(0).Sub(x, y)
}
func (b BigIntOperators) Add(x, y *big.Int) *big.Int {
	return big.NewInt(0).Add(x, y)
}

type Fib struct {
	BigIntOperators
}

//  (Public) Returns F(n).
func (f *Fib) Calc(n uint, fast bool) (*big.Int, error) {
	var err error
	var fst *big.Int
	if n < 0 {
		err = errors.New("Negative arguments not implemented")
	} else {
		switch fast {
		case true:
			fst, _ = f.fib_fast(n)
		default:
			fst, _ = f.fib(n)
		}
	}

	return fst, err
}

// (Private) Returns the tuple (F(n), F(n+1)).
func (f *Fib) fib_fast(n uint) (*big.Int, *big.Int) {
	/* n-е по индексу (от 0) число Фибоначчи
	   >> fib(10)
	   55
	   0s
	   >> fib(10000)
	   ...
	   0s
	   >> fib(100000)
	   ...
	   2.0001ms
	   >> fib(1000000)
	   ...
	   59.0034ms

	*/

	if n == 0 {
		return big.NewInt(0), big.NewInt(1)
	}

	a, b := f.fib_fast(n / 2)
	c := f.Mul(a, f.Sub(f.Mul(b, big.NewInt(2)), a))
	d := f.Add(f.Mul(a, a), f.Mul(b, b))
	if n%2 == 0 {
		return c, d
	} else {
		return d, f.Add(c, d)
	}
}

func (f *Fib) fib(n uint) (*big.Int, *big.Int) {
	/* n-е по индексу (от 0) число Фибоначчи
	   >> fib(10)
	   55
	   0s
	   >> fib(10000)
	   ...
	   1.0001ms
	   >> fib(100000)
	   ...
	   72.0042ms
	   >> fib(1000000)
	   ...
	   8.1724674s

	*/

	f0 := big.NewInt(0)
	f1 := big.NewInt(1)

	var i uint
	for i = 2; i <= n; i++ {
		f0.Add(f0, f1)
		f0, f1 = f1, f0

	}
	return f1, f0
}

func fibn(n int) []*big.Int {
	/* n первых чисел Фибоначчи
	   >> fibn(10)
	   [0 1 1 2 3 5 8 13 21 34]
	*/

	fn := make([]*big.Int, n, n)

	fn[0] = big.NewInt(0)
	fn[1] = big.NewInt(1)

	for i := 2; i < n; i++ {
		// bad
		var f = big.NewInt(0)
		fn[i] = f.Add(fn[i-1], fn[i-2])

	}

	return fn
}
