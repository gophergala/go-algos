package sort

import {
    "fmt"
}
/*
 * Referenced from: http://rosettacode.org/wiki/Sorting_algorithms/Selection_sort#Go
 */
func selectionSort(a []int) {
    last := len(a) - 1
    for i := 0; i < last; i++ {
        aMin := a[i]
        iMin := i
        for j := i + 1; j < len(a); j++ {
            if a[j] < aMin {
                aMin = a[j]
                iMin = j
            }
        }
        a[i], a[iMin] = aMin, a[i]
    }
}