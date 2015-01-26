
// source: https://github.com/dcadenas/pagerank/blob/master/pagerank_test.go

package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"testing"
    "github.com/gophergala/go-algos/search"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func round(f float64) float64 {
	return math.Floor(f*10+0.5) / 10
}

func toPercentage(f float64) float64 {
	tenPow3 := math.Pow(10, 3)
	return round(100 * (f * tenPow3) / tenPow3)
}

func assertRank(t *testing.T, pageRank Interface, expected map[int]float64) {
	const tolerance = 0.0001
	pageRank.Rank(0.85, tolerance, func(label int, rank float64) {
		rankAsPercentage := toPercentage(rank)
		if math.Abs(rankAsPercentage - expected[label]) > tolerance {
			t.Error("Rank for", label, "should be", expected[label], "but was", rankAsPercentage)
		}
	})
}

func assertEqual(t *testing.T, actual, expected interface{}) {
	if actual != expected {
		t.Error("Should be", expected, "but was", actual)
	}
}

func assert(t *testing.T, actual bool) {
	if !actual {
		t.Error("Should be true")
	}
}

func TestRound(t *testing.T) {
	assertEqual(t, round(0.6666666), 0.7)
}

func TestRankToPercentage(t *testing.T) {
	assertEqual(t, toPercentage(0.6666666), 66.7)
}

func TestShouldEnterTheBlock(t *testing.T) {
	pageRank := New()
	pageRank.Link(0, 1)

	entered := false
	pageRank.Rank(0.85, 0.0001, func(_ int, _ float64) {
		entered = true
	})

	assert(t, entered)
}

func TestShouldBePossibleToRecalculateTheRanksAfterANewLinkIsAdded(t *testing.T) {
	pageRank := New()
	pageRank.Link(0, 1)
	assertRank(t, pageRank, map[int]float64{0: 35.1, 1: 64.9})
	pageRank.Link(1, 2)
	assertRank(t, pageRank, map[int]float64{0: 18.4, 1: 34.1, 2: 47.4})
}

func TestShouldBePossibleToClearTheGraph(t *testing.T) {
	pageRank := New()
	pageRank.Link(0, 1)
	pageRank.Link(1, 2)
	pageRank.Clear()
	pageRank.Link(0, 1)
	assertRank(t, pageRank, map[int]float64{0: 35.1, 1: 64.9})
}

func TestShouldNotFailWhenCalculatingTheRankOfAnEmptyGraph(t *testing.T) {
	pageRank := New()
	pageRank.Rank(0.85, 0.0001, func(label int, rank float64) {
		t.Error("This should not be seen")
	})
}

func TestShouldReturnCorrectResultsWhenHavingADanglingNode(t *testing.T) {
	pageRank := New()
	//node 2 is a dangling node because it has no outbound links
	pageRank.Link(0, 2)
	pageRank.Link(1, 2)

	expectedRank := map[int]float64{
		0: 21.3,
		1: 21.3,
		2: 57.4,
	}

	assertRank(t, pageRank, expectedRank)
}

func TestShouldNotChangeTheGraphWhenAddingTheSameLinkManyTimes(t *testing.T) {
	pageRank := New()
	pageRank.Link(0, 2)
	pageRank.Link(0, 2)
	pageRank.Link(0, 2)
	pageRank.Link(1, 2)
	pageRank.Link(1, 2)

	expectedRank := map[int]float64{
		0: 21.3,
		1: 21.3,
		2: 57.4,
	}

	assertRank(t, pageRank, expectedRank)
}

func TestShouldReturnCorrectResultsForAStarGraph(t *testing.T) {
	pageRank := New()
	pageRank.Link(0, 2)
	pageRank.Link(1, 2)
	pageRank.Link(2, 2)

	expectedRank := map[int]float64{
		0: 5,
		1: 5,
		2: 90,
	}

	assertRank(t, pageRank, expectedRank)
}

func TestShouldBeUniformForACircularGraph(t *testing.T) {
	pageRank := New()
	pageRank.Link(0, 1)
	pageRank.Link(1, 2)
	pageRank.Link(2, 3)
	pageRank.Link(3, 4)
	pageRank.Link(4, 0)

	expectedRank := map[int]float64{
		0: 20,
		1: 20,
		2: 20,
		3: 20,
		4: 20,
	}

	assertRank(t, pageRank, expectedRank)
}

func TestShouldReturnCorrectResultsForAConvergingGraph(t *testing.T) {
	pageRank := New()
	pageRank.Link(0, 1)
	pageRank.Link(0, 2)
	pageRank.Link(1, 2)
	pageRank.Link(2, 2)

	expectedRank := map[int]float64{
		0: 5,
		1: 7.1,
		2: 87.9,
	}

	assertRank(t, pageRank, expectedRank)
}

func TestShouldCorrectlyReproduceTheWikipediaExample(t *testing.T) {
	//http://en.wikipedia.org/wiki/File:PageRanks-Example.svg
	pageRank := New()
	pageRank.Link(1, 2)
	pageRank.Link(2, 1)
	pageRank.Link(3, 0)
	pageRank.Link(3, 1)
	pageRank.Link(4, 3)
	pageRank.Link(4, 1)
	pageRank.Link(4, 5)
	pageRank.Link(5, 4)
	pageRank.Link(5, 1)
	pageRank.Link(6, 1)
	pageRank.Link(6, 4)
	pageRank.Link(7, 1)
	pageRank.Link(7, 4)
	pageRank.Link(8, 1)
	pageRank.Link(8, 4)
	pageRank.Link(9, 4)
	pageRank.Link(10, 4)

	expectedRank := map[int]float64{
		0:  3.3,  //a
		1:  38.4, //b
		2:  34.3, //c
		3:  3.9,  //d
		4:  8.1,  //e
		5:  3.9,  //f
		6:  1.6,  //g
		7:  1.6,  //h
		8:  1.6,  //i
		9:  1.6,  //j
		10: 1.6,  //k
	}

	assertRank(t, pageRank, expectedRank)
}

func BenchmarkOneMillion(b *testing.B) {
	n := 1000000

	pageRank := New()

	rand.Seed(5)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for from := 0; from < n; from++ {
			for j := 0; j < rand.Intn(400); j++ {
				too := rand.Intn(n)

				to := too
				if too > 800000 {
					to = rand.Intn(3)
				}

				pageRank.Link(from, to)
			}
		}
	}

	result := make([]float64, n)
	pageRank.Rank(0.85, 0.001, func(key int, val float64) {
		result[key] = val
	})

	fmt.Println("5 first values are", result[0], ",", result[1], ",", result[2], ",", result[3], ",", result[4])
	pageRank.Clear()
}