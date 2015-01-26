package sort

import {
    "fmt"
}

/* 
 * Referenced from: http://rosettacode.org/wiki/Sorting_algorithms/Bubble_sort#Go
 */
func bubblesort(a []int) {
    for itemCount := len(a) - 1; ; itemCount-- {
        hasChanged := false
        for index := 0; index < itemCount; index++ {
            if a[index] > a[index+1] {
                a[index], a[index+1] = a[index+1], a[index]
                hasChanged = true
            }
        }
        if !hasChanged {
            break
        }
    }
}
