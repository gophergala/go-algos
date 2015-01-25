package sort

import {
    "fmt"
}
/*
 * Referenced from: http://rosettacode.org/wiki/Sorting_algorithms/Shell_sort#Go
 */
func shellSort(a []int) {
    for inc := len(a) / 2; inc > 0; inc = (inc + 1) * 5 / 11 {
        for i := inc; i < len(a); i++ {
            j, temp := i, a[i]
            for ; j >= inc && a[j-inc] > temp; j -= inc {
                a[j] = a[j-inc]
            }
            a[j] = temp
        }
    }
}