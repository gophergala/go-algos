package search

// source: https://github.com/dcadenas/pagerank/blob/master/pagerank.go

import "math"

type Interface interface {
	Rank(followingProb, tolerance float64, resultFunc func(label int, rank float64))
	Link(from, to int)
}

type pageRank struct {
	inLinks               [][]int
	numberOutLinks        []int
	currentAvailableIndex int
	keyToIndex            map[int]int
	indexToKey            map[int]int
}

func New() *pageRank {
	pr := new(pageRank)
	pr.Clear()
	return pr
}

func (pr *pageRank) keyAsArrayIndex(key int) int {
	index, ok := pr.keyToIndex[key]

	if !ok {
		pr.currentAvailableIndex++
		index = pr.currentAvailableIndex
		pr.keyToIndex[key] = index
		pr.indexToKey[index] = key
	}

	return index
}

func (pr *pageRank) updateInLinks(fromAsIndex, toAsIndex int) {
	missingSlots := len(pr.keyToIndex) - len(pr.inLinks)

	if missingSlots > 0 {
		pr.inLinks = append(pr.inLinks, make([][]int, missingSlots)...)
	}

	pr.inLinks[toAsIndex] = append(pr.inLinks[toAsIndex], fromAsIndex)
}

func (pr *pageRank) updateNumberOutLinks(fromAsIndex int) {
	missingSlots := len(pr.keyToIndex) - len(pr.numberOutLinks)

	if missingSlots > 0 {
		pr.numberOutLinks = append(pr.numberOutLinks, make([]int, missingSlots)...)
	}

	pr.numberOutLinks[fromAsIndex] += 1
}

func (pr *pageRank) linkWithIndices(fromAsIndex, toAsIndex int) {
	pr.updateInLinks(fromAsIndex, toAsIndex)
	pr.updateNumberOutLinks(fromAsIndex)
}

func (pr *pageRank) Link(from, to int) {
	fromAsIndex := pr.keyAsArrayIndex(from)
	toAsIndex := pr.keyAsArrayIndex(to)

	pr.linkWithIndices(fromAsIndex, toAsIndex)
}

func (pr *pageRank) calculateDanglingNodes() []int {
	danglingNodes := make([]int, 0, len(pr.numberOutLinks))

	for i, numberOutLinksForI := range pr.numberOutLinks {
		if numberOutLinksForI == 0 {
			danglingNodes = append(danglingNodes, i)
		}
	}

	return danglingNodes
}

func (pr *pageRank) step(followingProb, tOverSize float64, p []float64, danglingNodes []int) []float64 {
	innerProduct := 0.0

	for _, danglingNode := range danglingNodes {
		innerProduct += p[danglingNode]
	}

	innerProductOverSize := innerProduct / float64(len(p))
	vsum := 0.0
	v := make([]float64, len(p))

	for i, inLinksForI := range pr.inLinks {
		ksum := 0.0

		for _, index := range inLinksForI {
			ksum += p[index] / float64(pr.numberOutLinks[index])
		}

		v[i] = followingProb*(ksum+innerProductOverSize) + tOverSize
		vsum += v[i]
	}

	inverseOfSum := 1.0 / vsum

	for i := range v {
		v[i] *= inverseOfSum
	}

	return v
}

func calculateChange(p, new_p []float64) float64 {
	acc := 0.0

	for i, pForI := range p {
		acc += math.Abs(pForI - new_p[i])
	}

	return acc
}

func (pr *pageRank) Rank(followingProb, tolerance float64, resultFunc func(label int, rank float64)) {
	size := len(pr.keyToIndex)
	inverseOfSize := 1.0 / float64(size)
	tOverSize := (1.0 - followingProb) / float64(size)
	danglingNodes := pr.calculateDanglingNodes()

	p := make([]float64, size)
	for i := range p {
		p[i] = inverseOfSize
	}

	change := 2.0

	for change > tolerance {
		new_p := pr.step(followingProb, tOverSize, p, danglingNodes)
		change = calculateChange(p, new_p)
		p = new_p
	}

	for i, pForI := range p {
		resultFunc(pr.indexToKey[i], pForI)
	}
}

func (pr *pageRank) Clear() {
	pr.inLinks = [][]int{}
	pr.numberOutLinks = []int{}
	pr.currentAvailableIndex = -1
	pr.keyToIndex = make(map[int]int)
	pr.indexToKey = make(map[int]int)
}