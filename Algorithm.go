package main

import (
	"fmt"
)

var numbers = []int{7, 2, 3, 1, 8, 9}

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			fmt.Println(arr)
			fmt.Println("the number", arr[j], "is switched by", arr[j+1])
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			} else {
				fmt.Println("no replacement take place")
			}
		}
		fmt.Println("times of number replacement", i+1)

	}
}

func main() {
	numbers := []int{7, 2, 3, 1, 8, 9}

	fmt.Println("numbers before sorting:", numbers)
	bubbleSort(numbers)
	fmt.Println("number after sorting:", numbers)
}
