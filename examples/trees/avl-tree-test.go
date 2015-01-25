// source: http://rosettacode.org/wiki/AVL_tree#Go
// adapted to our library

package main

import (
    "encoding/json"
    "fmt"
    "github.com/gophergala/go-algos/trees"
    "log"
)

type intKey int

// satisfy trees.Key
func (k intKey) Less(k2 trees.Key) bool { return k < k2.(intKey) }
func (k intKey) Eq(k2 trees.Key) bool   { return k == k2.(intKey) }

// use json for cheap tree visualization
func dump(tree *trees.Node) {
    b, err := json.MarshalIndent(tree, "", "   ")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(b))
}

func main() {
    var tree *trees.Node
    fmt.Println("Empty tree:")
    dump(tree)

    fmt.Println("\nInsert test:")
    trees.Insert(&tree, intKey(3))
    trees.Insert(&tree, intKey(1))
    trees.Insert(&tree, intKey(4))
    trees.Insert(&tree, intKey(1))
    trees.Insert(&tree, intKey(5))
    dump(tree)

    fmt.Println("\nRemove test:")
    trees.Remove(&tree, intKey(3))
    trees.Remove(&tree, intKey(1))
    dump(tree)
}

