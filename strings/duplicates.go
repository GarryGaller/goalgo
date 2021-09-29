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

func Duplicates(in []string) (out []string) {
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