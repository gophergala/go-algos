// source: https://github.com/adlawson/search-algorithms/blob/master/golang/bfs.go

package search

func dfs(node int, nodes map[int][]int, fn func (int)) {
    dfs_recur(node, map[int]bool{}, fn)
}

func dfs_recur(node int, v map[int]bool, fn func (int)) {
    v[node] = true
    fn(node)
    for _, n := range nodes[node] {
        if _, ok := v[n]; !ok {
            dfs_recur(n, v, fn)
        }
    }
}