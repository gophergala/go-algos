// source: http://rosettacode.org/wiki/Heapsort#Go
// adapted to our library

package main
 
import (
  "sort"
  "container/heap"
  "github.com/gophergala/go-algos/sort"
  "fmt"
)

func main() {
    a := []int{170, 45, 75, -90, -802, 24, 2, 66}

    fmt.Println("before:", a)

    heapSort(sort.IntSlice(a))

    fmt.Println("after: ", a)

}