package bits

import (
    "fmt"
    "testing"
)


func Test(t *testing.T) {
    var b Bits
    b = Set(b, F0)
    b = Toggle(b, F2)
    for i, flag := range []Bits{F0, F1, F2} {
        fmt.Println(i, Has(b, flag))
    }
}