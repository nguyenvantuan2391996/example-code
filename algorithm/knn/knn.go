package main

import (
	"fmt"
	"math"
	"sort"
)

func distance(vectorA, vectorB []float64) float64 {
	if len(vectorA) != len(vectorB) {
		return 0
	}

	d := float64(0)
	for i := 0; i < len(vectorA); i++ {
		d += math.Pow(vectorA[i]-vectorB[i], 2)
	}

	return math.Sqrt(d)
}

func main() {
	distances := make(map[float64]string)
	n := 3
	target := []float64{6.9}
	for p, rank := range Points {
		distances[distance(target, []float64{p})] = rank
	}

	keys := make([]float64, 0, len(distances))
	for k := range distances {
		keys = append(keys, k)
	}

	sort.Float64s(keys)

	topN := make(map[float64]string)
	for i := 0; i < n; i++ {
		topN[keys[i]] = distances[keys[i]]
	}

	frequencyMap := make(map[string]int)

	// Populate the frequency map
	for _, value := range topN {
		frequencyMap[value]++
	}

	// Find the value with the highest frequency
	mostFrequentValue := ""
	maxCount := 0
	for value, count := range frequencyMap {
		if count > maxCount {
			mostFrequentValue = value
			maxCount = count
		}
	}

	fmt.Printf("Top n: %v", topN)

	fmt.Printf("KNN is predict %v is %v", target, mostFrequentValue)
}

var Points = map[float64]string{
	1.0:  "Poor/ Weak",
	1.5:  "Poor/ Weak",
	2.0:  "Poor/ Weak",
	2.5:  "Poor/ Weak",
	3.0:  "Poor/ Weak",
	3.5:  "Poor/ Weak",
	4.0:  "Poor/ Weak",
	4.5:  "Below Average",
	5.0:  "Below Average",
	5.5:  "Average",
	6.0:  "Average",
	6.5:  "Average",
	7.0:  "Good",
	7.5:  "Good",
	8.0:  "Good",
	8.5:  "Excellent",
	9.0:  "Excellent",
	9.5:  "Excellent",
	10.0: "Excellent",
}
