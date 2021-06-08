package main

import (
	"fmt"
	
)

// Swap two array values given their indices.
func Swap(v interface{}, i, j int) {
	switch a := v.(type) {
	case []int:
		SwapInt(a, i, j)
	case []string:
		SwapString(a, i, j)
	}
}

// SwapInt two array values given their indices.
func SwapInt(a []int, i, j int) {
	tmp := a[i]
	a[i] = a[j]
	a[j] = tmp
}

// SwapString two array values given their indices.
func SwapString(a []string, i, j int) {
	tmp := a[i]
	a[i] = a[j]
	a[j] = tmp
}



func cyclicSort(nums []int, start int) {
	i := 0
	for i < len(nums) {
		// if the current number is not at the correct index, swap it with the
		// one that is at the correct index.
		correctIndex := nums[i] - 1
		if nums[i] != nums[correctIndex] {
			Swap(nums, i, correctIndex)
		} else {
			i++
		}
	}
}

func main(){

    nums := []int{5,6,7,8,4,2,3,9,10}
    cyclicSort(nums)
    
    fmt.Printf("#v\n", nums)

}    
    