package strings

func Deduplicate(in *[]string) {
	/* Classic duplicate removal algorithm
	   on sorted data
	   Data varies in-place
	*/

	j := 0

	for i := 1; i < len(*in); i++ {
		if (*in)[j] != (*in)[i] {
			j++
			(*in)[j] = (*in)[i]
        }
	}
	*in = (*in)[:j+1]
	return
}

func Counter(in []string) (out map[string]int) {
	/*The dictionary for counting occurrences */
	out = make(map[string]int)

	for i := 0; i < len(in); i++ {
		out[in[i]] += 1
	}

	return
}

func Repeated(in []string) (out []string) {
	/* Search for repeated strings*/

	counter := Counter(in)
	for k, v := range counter {
		if v > 1 {
			out = append(out, k)
		}
	}

	return
}

func Unique(in []string) (out []string) {
	/* Search for unique strings */

	counter := Counter(in)
	for k, v := range counter {
		if v == 1 {
			out = append(out, k)
		}
	}

	return
}

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

const (
	PatternSize int = 100
)

func KMP(haystack string, needle string) []int {
	next := preKMP(needle)
	i := 0
	j := 0
	m := len(needle)
	n := len(haystack)

	x := []byte(needle)
	y := []byte(haystack)
	var ret []int

	//got zero target or want, just return empty result
	if m == 0 || n == 0 {
		return ret
	}

	//want string bigger than target string
	if n < m {
		return ret
	}

	for j < n {
		for i > -1 && x[i] != y[j] {
			i = next[i]
		}
		i++
		j++

		//fmt.Println(i, j)
		if i >= m {
			ret = append(ret, j-i)
			//fmt.Println("find:", j, i)
			i = next[i]
		}
	}

	return ret
}

func preKMP(x string) [PatternSize]int {
	var i, j int
	length := len(x) - 1
	var kmpNext [PatternSize]int
	i = 0
	j = -1
	kmpNext[0] = -1

	for i < length {
		for j > -1 && x[i] != x[j] {
			j = kmpNext[j]
		}

		i++
		j++

		if x[i] == x[j] {
			kmpNext[i] = kmpNext[j]
		} else {
			kmpNext[i] = j
		}
	}
	return kmpNext
}

func preMP(x string) [PatternSize]int {
	var i, j int
	length := len(x) - 1
	var mpNext [PatternSize]int
	i = 0
	j = -1
	mpNext[0] = -1

	for i < length {
		for j > -1 && x[i] != x[j] {
			j = mpNext[j]
		}
		i++
		j++
		mpNext[i] = j
	}
	return mpNext
}
