package main

import "testing"

func TestGreedyFlorist1(t *testing.T) {
	k := int32(3)
	c := []int32{2, 5, 6}

	minCost := getMinimumCost(k, c)

	if minCost != 13 {
		t.Errorf("got %d instead of 13", minCost)
	}
}

func TestGreedyFlorist2(t *testing.T) {
	k := int32(2)
	c := []int32{2, 5, 6}

	minCost := getMinimumCost(k, c)

	if minCost != 15 {
		t.Errorf("got %d instead of 15", minCost)
	}
}

func TestGreedyFlorist3(t *testing.T) {
	k := int32(3)
	c := []int32{1, 3, 5, 7, 9}

	minCost := getMinimumCost(k, c)

	if minCost != 29 {
		t.Errorf("got %d instead of 29", minCost)
	}
}

func TestGreedyFlorist4(t *testing.T) {
	k := int32(3)
	c := []int32{390225, 426456, 688267, 800389, 990107, 439248, 240638, 15991, 874479, 568754, 729927, 980985, 132244, 488186, 5037, 721765, 251885, 28458, 23710, 281490, 30935, 897665, 768945, 337228, 533277, 959855, 927447, 941485, 24242, 684459, 312855, 716170, 512600, 608266, 779912, 950103, 211756, 665028, 642996, 262173, 789020, 932421, 390745, 433434, 350262, 463568, 668809, 305781, 815771, 550800}

	minCost := getMinimumCost(k, c)

	if minCost != 163578911 {
		t.Errorf("got %d instead of 163578911", minCost)
	}
}
