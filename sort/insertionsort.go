package sort

import {
    "fmt"
}
/* 
 * Referenced from: http://rosettacode.org/wiki/Sorting_algorithms/Insertion_sort#Go
 */
func insertionSort(a []int) {
    for i := 1; i < len(a); i++ {
        value := a[i]
        j := i - 1
        for j >= 0 && a[j] > value {
            a[j+1] = a[j]
            j = j - 1
        }
        a[j+1] = value
    }
}