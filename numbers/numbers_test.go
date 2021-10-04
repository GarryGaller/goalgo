package numbers

import (
	"reflect"
	"testing"
)

func TestDeduplicate(t *testing.T) {

	testCases := []struct {
		in       []int
		expected []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{1, 2, 2}, []int{1, 2}},
		{[]int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}, []int{1, 2, 3, 4}},
	}

	for _, tc := range testCases {
		t.Run("Deduplicate", func(t *testing.T) {
			tmp := make([]int, len(tc.in))
			copy(tmp, tc.in)
			got := Deduplicate(tmp)
			if !reflect.DeepEqual(tmp[:got], tc.expected) {
				t.Errorf("Deduplicate() = %v; want %v", tc.in[:got], tc.expected)
			}
		})

		t.Run("Deduplicate2", func(t *testing.T) {
			tmp := make([]int, len(tc.in))
			copy(tmp, tc.in)
			got := Deduplicate2(tmp)
			if !reflect.DeepEqual(tmp[:got], tc.expected) {
				t.Errorf("Deduplicate2() = %v; want %v", tc.in[:got], tc.expected)
			}
		})
	}
}
