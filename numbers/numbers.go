
package main

import (
    "fmt"
    "sort"
    
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


func removeDuplicates(a []int) int {
    /*only for sorted arrays*/
    
    if len(a) == 0 {
        return 0
    }

    n:= 1
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


func deduplicate(a []int) int {
   
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
    return prev+1
    
}




