package numbers

import (
	//"fmt"
	"math"
	"math/rand"
	//"sort"
)

// SwapInt two array values given their indices.
func SwapInt(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}

func mergeSortedArray(a1, a2 []int) []int {
	out := []int{}

	// keep two "pointers" at index 0 and move up accordingly as one get
	// merged in.
	i, j := 0, 0
	for i < len(a1) && j < len(a2) {
		if a1[i] < a2[j] {
			out = append(out, a1[i])
			i++
		} else {
			out = append(out, a2[j])
			j++
		}
	}

	// if we get here, one array must have bigger size than the other. could
	// figure out which one is it then copy the rest of its to our final one.
	if i < len(a1) {
		out = append(out, a1[i:]...)
	}

	if j < len(a2) {
		out = append(out, a2[j:]...)
	}

	return out
}

// Mimax returns min and max from a list of integers.
func Mimax(nums ...int) (int, int) {
	min, max := nums[0], nums[0]

	for _, num := range nums {
		if min > num {
			min = num
		}

		if max < num {
			max = num
		}
	}

	return min, max
}

// Min returns min from a list of integers.
func Min(nums ...int) int {
	min := nums[0]

	for _, num := range nums {
		if min > num {
			min = num
		}
	}

	return min
}

// Max returns max from a list of integers.
func Max(nums ...int) int {
	max := nums[0]

	for _, num := range nums {
		if max < num {
			max = num
		}
	}

	return max
}

// Random rertuns a random number over a range
func Random(min, max int) int {
	if min == max {
		return min
	}
	return rand.Intn(max-min) + min
}

// Permutations
func Perm(a []rune, f func([]rune)) {
	perm(a, f, 0)
}

func perm(a []rune, f func([]rune), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

// Contain checks if the target is in a slice.
func ContainsInt(s []int, target int) bool {
	for _, v := range s {
		if v == target {
			return true
		}
	}

	return false
}

func Deduplicate(a []int) int {
	/*only for sorted arrays*/

	if len(a) == 0 {
		return 0
	}

	n := 1
	i := 1
	for i < len(a) {
		// if we see a non-duplicate number, move it next to the last
		// non-duplicate one that we've seen.
		if a[n-1] != a[i] {
			a[n] = a[i]
			n++
		}
		i++

	}

	return n
}

func Deduplicate2(a []int) int {

	if len(a) == 0 {
		return 0
	}

	prev := 0
	for next := 1; next < len(a); next++ {
		if a[prev] == a[next] {
			continue
		}
		prev++
		a[prev] = a[next]
	}
	return prev + 1

}

func ReverseInteger(in int64) int64 {
	var out, remainder int64

	for in != 0 {
		remainder = in % 10

		// check for overflow/underflow before multiplying by 10.
		if out > math.MaxInt64/10 || (out == math.MaxInt64/10 && remainder > 7) {
			return 0
		}
		if out < math.MinInt64/10 || (out == math.MinInt64/10 && remainder < -8) {
			return 0
		}

		out = out*10 + remainder
		in /= 10
	}

	return out
}

func FindPairTargetSum(a []int, target int) []int {
	/*
	   Approach:
	   - Have one pointer start at the beginning and one at the end of the array.
	   - At each step, see if the two pointers add up to the target sum and move
	     them toward each other accordingly.

	   Cost:
	   - O(n) time, O(n) space.
	*/

	out := []int{}

	start := 0
	end := len(a) - 1

	for start < end {
		sum := a[start] + a[end]
		if sum > target {
			end--
		} else if sum < target {
			start++
		} else {
			out = append(out, start, end)
			break
		}
	}

	return out
}
