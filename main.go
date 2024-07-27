package main

import (
	"container/list"
	"fmt"
	"math"
	"sort"
	"sync"
)

func main() {
	// Create channels for communication
	chIntersection := make(chan *list.List)
	chUnion := make(chan *list.List)
	chDifference := make(chan *list.List)
	chSymmetricDifference := make(chan *list.List)
	chIsSubset := make(chan bool)
	chMean := make(chan float64)
	chMedian := make(chan float64)
	chMode := make(chan *list.List)
	chVariance := make(chan float64)
	chStandardDeviation := make(chan float64)

	var calcWg sync.WaitGroup
	calcWg.Add(10)

	var displayWg sync.WaitGroup
	displayWg.Add(10)

	results := list.New()

	// Start goroutines for calculations
	go CalculateIntersection(createList(1, 2, 3), createList(2, 3, 4), chIntersection, &calcWg)
	go CalculateUnion(createList(1, 2, 3), createList(2, 3, 4), chUnion, &calcWg)
	go CalculateDifference(createList(1, 2, 3), createList(2, 3, 4), chDifference, &calcWg)
	go CalculateSymmetricDifference(createList(1, 2, 3), createList(2, 3, 4), chSymmetricDifference, &calcWg)
	go CalculateIsSubset(createList(1, 2), createList(1, 2, 3), chIsSubset, &calcWg)
	go CalculateMean(createList(1, 2, 3, 4, 5), chMean, &calcWg)
	go CalculateMedian(createList(1, 2, 3, 4, 5), chMedian, &calcWg)
	go CalculateMode(createList(1, 2, 2, 3, 3, 3, 4), chMode, &calcWg)
	go CalculateVariance(createList(1, 2, 3, 4, 5), chVariance, &calcWg)
	go CalculateStandardDeviation(createList(1, 2, 3, 4, 5), chStandardDeviation, &calcWg)

	// Start goroutines to collect results
	go collectListResult(chIntersection, "Intersection", results, &displayWg)
	go collectListResult(chUnion, "Union", results, &displayWg)
	go collectListResult(chDifference, "Difference", results, &displayWg)
	go collectListResult(chSymmetricDifference, "Symmetric Difference", results, &displayWg)
	go collectBoolResult(chIsSubset, "Subset", results, &displayWg)
	go collectFloatResult(chMean, "Mean", results, &displayWg)
	go collectFloatResult(chMedian, "Median", results, &displayWg)
	go collectListResult(chMode, "Mode", results, &displayWg)
	go collectFloatResult(chVariance, "Variance", results, &displayWg)
	go collectFloatResult(chStandardDeviation, "Standard Deviation", results, &displayWg)

	// Wait for all calculation goroutines to finish
	calcWg.Wait()

	// Close all channels
	close(chIntersection)
	close(chUnion)
	close(chDifference)
	close(chSymmetricDifference)
	close(chIsSubset)
	close(chMean)
	close(chMedian)
	close(chMode)
	close(chVariance)
	close(chStandardDeviation)

	// Wait for all result collection goroutines to finish
	displayWg.Wait()

	// Print all results
	for e := results.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func createList(elements ...int) *list.List {
	l := list.New()
	for _, e := range elements {
		l.PushBack(e)
	}
	return l
}

func collectListResult(ch chan *list.List, name string, results *list.List, wg *sync.WaitGroup) {
	defer wg.Done()
	for result := range ch {
		var resStr string
		for e := result.Front(); e != nil; e = e.Next() {
			resStr += fmt.Sprintf("%v ", e.Value)
		}
		results.PushBack(fmt.Sprintf("%s: %s", name, resStr))
	}
}

func collectBoolResult(ch chan bool, name string, results *list.List, wg *sync.WaitGroup) {
	defer wg.Done()
	for result := range ch {
		results.PushBack(fmt.Sprintf("%s: %v", name, result))
	}
}

func collectFloatResult(ch chan float64, name string, results *list.List, wg *sync.WaitGroup) {
	defer wg.Done()
	for result := range ch {
		results.PushBack(fmt.Sprintf("%s: %f", name, result))
	}
}

// Collection calculations

// Intersection
func CalculateIntersection(set1, set2 *list.List, ch chan *list.List, wg *sync.WaitGroup) {
	defer wg.Done()
	result := Intersection(set1, set2)
	ch <- result
}

func Intersection(set1, set2 *list.List) *list.List {
	m := make(map[interface{}]bool)
	for e := set1.Front(); e != nil; e = e.Next() {
		m[e.Value] = true
	}
	result := list.New()
	for e := set2.Front(); e != nil; e = e.Next() {
		if m[e.Value] {
			result.PushBack(e.Value)
		}
	}
	return result
}

// Union
func CalculateUnion(set1, set2 *list.List, ch chan *list.List, wg *sync.WaitGroup) {
	defer wg.Done()
	result := Union(set1, set2)
	ch <- result
}

func Union(set1, set2 *list.List) *list.List {
	m := make(map[interface{}]bool)
	result := list.New()
	for e := set1.Front(); e != nil; e = e.Next() {
		if !m[e.Value] {
			m[e.Value] = true
			result.PushBack(e.Value)
		}
	}
	for e := set2.Front(); e != nil; e = e.Next() {
		if !m[e.Value] {
			m[e.Value] = true
			result.PushBack(e.Value)
		}
	}
	return result
}

// Difference
func CalculateDifference(set1, set2 *list.List, ch chan *list.List, wg *sync.WaitGroup) {
	defer wg.Done()
	result := Difference(set1, set2)
	ch <- result
}

func Difference(set1, set2 *list.List) *list.List {
	m := make(map[interface{}]bool)
	for e := set2.Front(); e != nil; e = e.Next() {
		m[e.Value] = true
	}
	result := list.New()
	for e := set1.Front(); e != nil; e = e.Next() {
		if !m[e.Value] {
			result.PushBack(e.Value)
		}
	}
	return result
}

// Symmetric Difference
func CalculateSymmetricDifference(set1, set2 *list.List, ch chan *list.List, wg *sync.WaitGroup) {
	defer wg.Done()
	result := SymmetricDifference(set1, set2)
	ch <- result
}

func SymmetricDifference(set1, set2 *list.List) *list.List {
	m1 := make(map[interface{}]bool)
	m2 := make(map[interface{}]bool)
	for e := set1.Front(); e != nil; e = e.Next() {
		m1[e.Value] = true
	}
	for e := set2.Front(); e != nil; e = e.Next() {
		m2[e.Value] = true
	}
	result := list.New()
	for e := set1.Front(); e != nil; e = e.Next() {
		if !m2[e.Value] {
			result.PushBack(e.Value)
		}
	}
	for e := set2.Front(); e != nil; e = e.Next() {
		if !m1[e.Value] {
			result.PushBack(e.Value)
		}
	}
	return result
}

// Subset
func CalculateIsSubset(set1, set2 *list.List, ch chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	result := IsSubset(set1, set2)
	ch <- result
}

func IsSubset(set1, set2 *list.List) bool {
	m := make(map[interface{}]bool)
	for e := set2.Front(); e != nil; e = e.Next() {
		m[e.Value] = true
	}
	for e := set1.Front(); e != nil; e = e.Next() {
		if !m[e.Value] {
			return false
		}
	}
	return true
}

// Mean
func CalculateMean(collection *list.List, ch chan float64, wg *sync.WaitGroup) {
	defer wg.Done()
	result := Mean(collection)
	ch <- result
}

func Mean(collection *list.List) float64 {
	sum := 0
	count := 0
	for e := collection.Front(); e != nil; e = e.Next() {
		sum += e.Value.(int)
		count++
	}
	return float64(sum) / float64(count)
}

// Median
func CalculateMedian(collection *list.List, ch chan float64, wg *sync.WaitGroup) {
	defer wg.Done()
	result := Median(collection)
	ch <- result
}

func Median(collection *list.List) float64 {
	var values []int
	for e := collection.Front(); e != nil; e = e.Next() {
		values = append(values, e.Value.(int))
	}
	sort.Ints(values)
	n := len(values)
	if n%2 == 0 {
		return float64(values[n/2-1]+values[n/2]) / 2
	}
	return float64(values[n/2])
}

// Mode
func CalculateMode(collection *list.List, ch chan *list.List, wg *sync.WaitGroup) {
	defer wg.Done()
	result := Mode(collection)
	ch <- result
}

func Mode(collection *list.List) *list.List {
	frequency := make(map[int]int)
	maxFreq := 0
	for e := collection.Front(); e != nil; e = e.Next() {
		frequency[e.Value.(int)]++
		if frequency[e.Value.(int)] > maxFreq {
			maxFreq = frequency[e.Value.(int)]
		}
	}
	result := list.New()
	for value, freq := range frequency {
		if freq == maxFreq {
			result.PushBack(value)
		}
	}
	return result
}

// Variance
func CalculateVariance(collection *list.List, ch chan float64, wg *sync.WaitGroup) {
	defer wg.Done()
	result := Variance(collection)
	ch <- result
}

func Variance(collection *list.List) float64 {
	mean := Mean(collection)
	var sum float64
	count := 0
	for e := collection.Front(); e != nil; e = e.Next() {
		sum += math.Pow(float64(e.Value.(int))-mean, 2)
		count++
	}
	return sum / float64(count)
}

// Standard Deviation
func CalculateStandardDeviation(collection *list.List, ch chan float64, wg *sync.WaitGroup) {
	defer wg.Done()
	result := StandardDeviation(collection)
	ch <- result
}

func StandardDeviation(collection *list.List) float64 {
	return math.Sqrt(Variance(collection))
}
