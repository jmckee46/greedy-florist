package main

import "sort"

type FlowerCosts []int32

func (a FlowerCosts) Len() int           { return len(a) }
func (a FlowerCosts) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a FlowerCosts) Less(i, j int) bool { return a[i] < a[j] }

func getMinimumCost(k int32, c []int32) int32 {
	var numSold int32
	var cost int32
	cLength := int32(len(c))

	flowerCosts := FlowerCosts(c)
	sort.Sort(flowerCosts)

	for i := cLength - 1; i >= 0; i-- {
		cost += purchaseFlower(k, flowerCosts[i], numSold)
		numSold++
	}
	return cost
}

func purchaseFlower(k int32, price int32, numSold int32) int32 {

	return (numSold/k + 1) * price
}

func main() {}
