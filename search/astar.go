// source: https://github.com/xarg/gopathfinding/blob/master/astar.go

//pathfinding package implements pathfinding algorithms such as Dijkstra and A*
package search

import (
	"fmt"
)

//Defining possible graph elements
const (
	UNKNOWN int = iota - 1
	LAND
	WALL
	START
	STOP
)

type MapData [][]int

//Return a new MapData by value given some dimensions
func NewMapData(rows, cols int) *MapData {
	result := make(MapData, rows)
	for i := 0; i < rows; i++ {
		result[i] = make([]int, cols)
	}
	return &result
}

//A node is just a set of x, y coordinates with a parent node and a
//heuristic value H
type Node struct {
	x, y   int //Using int for efficiency
	parent *Node
	H      int //Heuristic (aproximate distance)
	cost   int //Path cost for this node
}

//Create a new node
func NewNode(x, y int) *Node {
	node := &Node{
		x:      x,
		y:      y,
		parent: nil,
		H:      0,
		cost:   0,
	}
	return node
}

//Return string representation of the node
func (self *Node) String() string {
	return fmt.Sprintf("<Node x:%d y:%d addr:%d>", self.x, self.y, &self)
}

//Start, end nodes and a slice of nodes
type Graph struct {
	start, stop *Node
	nodes       []*Node
	data        *MapData
}

//Return a Graph from a map of coordinates (those that are passible)
func NewGraph(map_data *MapData) *Graph {
	var start, stop *Node
	var nodes []*Node
	for i, row := range *map_data {
		for j, _type := range row {
			if _type == START || _type == STOP {
				node := NewNode(i, j)
				nodes = append(nodes, node)
				if _type == START {
					start = node
				}
				if _type == STOP {
					stop = node
				}
			}
		}
	}
	g := &Graph{
		nodes: nodes,
		start: start,
		stop:  stop,
		data:  map_data,
	}
	return g
}

//Get *Node based on x, y coordinates.
func (self *Graph) Node(x, y int) *Node {
	//Check if node is not already in the graph and append that node
	for _, n := range self.nodes {
		if n.x == x && n.y == y {
			return n
		}
	}
	map_data := *self.data
	if map_data[x][y] == LAND || map_data[x][y] == STOP {
		//Create a new node and add it to the graph
		n := NewNode(x, y)
		self.nodes = append(self.nodes, n)
		return n
	}
	return nil
}

//Get the nodes near some node
func (self *Graph) adjacentNodes(node *Node) []*Node {
	var result []*Node
	map_data := *self.data
	rows := len(map_data)
	cols := len(map_data[0])

	//If the coordinates are passable then create a new node and add it
	if node.x <= rows && node.y+1 < cols {
		if new_node := self.Node(node.x, node.y+1); new_node != nil {
			result = append(result, new_node)
		}
	}
	if node.x <= rows && node.y-1 >= 0 {
		new_node := self.Node(node.x, node.y-1)
		if new_node != nil {
			result = append(result, new_node)
		}
	}
	if node.y <= cols && node.x+1 < rows {
		new_node := self.Node(node.x+1, node.y)
		if new_node != nil {
			result = append(result, new_node)
		}
	}
	if node.y <= cols && node.x-1 >= 0 {
		new_node := self.Node(node.x-1, node.y)
		if new_node != nil {
			result = append(result, new_node)
		}
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func removeNode(nodes []*Node, node *Node) []*Node {
	ith := -1
	for i, n := range nodes {
		if n == node {
			ith = i
			break
		}
	}
	if ith != -1 {
		copy(nodes[ith:], nodes[ith+1:])
		nodes = nodes[:len(nodes)-1]
	}
	return nodes
}

func hasNode(nodes []*Node, node *Node) bool {
	for _, n := range nodes {
		if n == node {
			return true
		}
	}
	return false
}

//Return the node with the minimum H
func minH(nodes []*Node) *Node {
	if len(nodes) == 0 {
		return nil
	}
	result_node := nodes[0]
	minH := result_node.H
	for _, node := range nodes {
		if node.H < minH {
			minH = node.H
			result_node = node
		}
	}
	return result_node
}

func retracePath(current_node *Node) []*Node {
	var path []*Node
	path = append(path, current_node)
	for current_node.parent != nil {
		path = append(path, current_node.parent)
		current_node = current_node.parent
	}
	//Reverse path
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}

// In our particular case: Manhatan distance
func Heuristic(graph *Graph, tile *Node) int {
	return abs(graph.stop.x-tile.x) + abs(graph.stop.y-tile.y)
}

//A* search algorithm. See http://en.wikipedia.org/wiki/A*_search_algorithm
func Astar(graph *Graph) []*Node {
	var path, openSet, closedSet []*Node

	openSet = append(openSet, graph.start)
	for len(openSet) != 0 {
		//Get the node with the min H
		current := minH(openSet)
		if current.parent != nil {
			current.cost = current.parent.cost + 1
		}
		if current == graph.stop {
			return retracePath(current)
		}
		openSet = removeNode(openSet, current)
		closedSet = append(closedSet, current)
		for _, tile := range graph.adjacentNodes(current) {
			if tile != nil && graph.stop != nil && !hasNode(closedSet, tile) {
				tile.H = Heuristic(graph, tile) + current.cost
				if !hasNode(openSet, tile) {
					openSet = append(openSet, tile)
				}
				tile.parent = current
			}
		}
	}
	return path
}
