package strings


func Reverse(value string) string {
    // Convert string to rune slice.
    // ... This method works on the level of runes, not bytes.
    data := []rune(value)
    result := []rune{}

    // Add runes in reverse order.
    for i := len(data) - 1; i >= 0; i-- {
        result = append(result, data[i])
    }

    // Return new string.
    return string(result)
}

func Reverse2(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func Reverse3(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func Reverse4(s string) string {
    runes := []rune(s)
    n := len(runes)
    for i := 0; n / 2; i++ {
        runes[n-i-1], runes[i] = runes[i], runes[n-i-1]
    }
    return string(runes)
}

func Reverse5(str string) string {
    var size int

    tail := len(str)
    buf := make([]byte, tail)
    s := buf

    for len(str) > 0 {
        _, size = utf8.DecodeRuneInString(str)
        tail -= size
        s = append(s[:tail], []byte(str[:size])...)
        str = str[size:]
    }

    return string(buf)
}