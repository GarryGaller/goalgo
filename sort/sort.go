package sort


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

func СyclicSort(nums []int) {
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

func СountingSort(in []int, max int) []int {
    // utilize max value to create a fix-sized list of item counts.
    counts := make([]int, max+1)
    out := make([]int, 0)

    // populate the array where its indices represent items themselves and
    // its values represent how many time the item appears.
    for _, item := range in {
        counts[item]++
    }

    // iterate through the counts and add the item to the output list.
    for i := 0; i < len(counts); i++ {
        count := counts[i]

        for j := 0; j < count; j++ {
            out = append(out, i)
        }
    }

    return out
}
