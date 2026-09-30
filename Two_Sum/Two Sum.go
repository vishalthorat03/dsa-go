package main

import "fmt"

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		if j, ok := seen[complement]; ok {
			return []int{j, i}
		}

		seen[num] = i
	}

	return []int{}
}

func main() {
	var n int

	fmt.Print("Enter number of elements: ")
	fmt.Scan(&n)

	nums := make([]int, n)

	fmt.Println("Enter the elements:")
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	var target int
	fmt.Print("Enter target: ")
	fmt.Scan(&target)

	result := twoSum(nums, target)

	fmt.Println("Answer:", result)
}
