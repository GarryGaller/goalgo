package slice

import (
    "fmt"
    "testing"
)

func BenchmarkRemoveFast(b *testing.B) {

    data := make([]string, 1000000)
    for i := 0; i < 1000000; i++ {
        data[i] = fmt.Sprint(i)
    }

    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        RemoveFast(data, 2)
    }
}

func BenchmarkRemoveWithSaveOrder(b *testing.B) {

    data := make([]string, 1000000)
    for i := 0; i < 1000000; i++ {
        data[i] = fmt.Sprint(i)
    }

    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        RemoveWithSaveOrder(data, 2)
    }
}

func BenchmarkRemoveWithSaveOrder2(b *testing.B) {

    data := make([]string, 1000000)
    for i := 0; i < 1000000; i++ {
        data[i] = fmt.Sprint(i)
    }

    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        RemoveWithSaveOrder2(data, 2)
    }
}

/*
BenchmarkRemoveFast
BenchmarkRemoveFast-4                   393420177            2.656 ns/op               0 B/op          0 allocs/op
BenchmarkRemoveWithSaveOrder
BenchmarkRemoveWithSaveOrder-4                21          55860338 ns/op               4 B/op          0 allocs/op
BenchmarkRemoveWithSaveOrder2
BenchmarkRemoveWithSaveOrder2-4               20          55553175 ns/op               0 B/op          0 allocs/op
PASS
ok      github.com/GarryGaller/goalgo/slice     6.713s
*/
