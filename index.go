package main

import "fmt"

func main() {
	nums := []int{1, 3, 3, 5, 5, 6, 6, 5, 3, 3}
	fmt.Println(twoMostFrequent(nums))
}

func twoMostFrequent(nums []int) []int {

	// Hashmap to hold the count of the element
	hash := make(map[int]int)

	count_most_freq := 0

	count_second_freg := 0

	most := 0

	second := 0

	results := []int{}

	// Just return empty if provided array is empty
	if len(nums) == 0 {
		return results
	}

	for _, value := range nums {

		_, ok := hash[value]

		if !ok {
			hash[value] = 0
		}
		hash[value] = hash[value] + 1
	}

	for key, value := range hash {
		if value > count_most_freq {
			second = most
			count_second_freg = count_most_freq

			most = key
			count_most_freq = value
		} else if count_second_freg < count_most_freq && value > count_second_freg {
			second = key
			count_second_freg = value
		}
	}

	results = []int{most, second}

	return results
}
