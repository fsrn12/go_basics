package lib

import (
	f "fmt"
)

func BasicDataStructures() {
	f.Println("==================Arrays, Slices & Maps==================")
	// numsOnInit := [...]int{10, 20,30}
	nums := [3]int{10, 20, 30}
	f.Printf("Length of Nums Array: %d\n", len(nums))
	f.Printf("First value of Nums Array: %d\n", nums[0])
	f.Printf("Last value of Nums Array: %d\n", nums[len(nums)-1])
	f.Printf("Value at index 1 of Nums Array: %d\n", nums[1])

	// multi-dimensional array
	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	f.Printf("multi-dimensional array: %v\n", matrix)
	f.Printf("Third value of the second element of matrix: %d\n", matrix[1][2]) // 6

	// Slices are dynamic arrays.
	slideColors := []string{"red", "purple", "pink", "magenta"}
	println("SlideColors Slice Length", len(slideColors))
	f.Printf("The theme colours for the slide: %v \n", slideColors)
	secondSlideColors := append(slideColors, "white", "crimson")
	println("SecondSlideColors Slice Length", len(secondSlideColors))
	f.Printf("The theme colours for the slide: %v \n", secondSlideColors)

	// appending two or more slices together
	fruits := []string{"mango", "passion fruit", "banana", "watermelon"}
	moreFruits := []string{"apple", "strawberries", "pears"}
	fruits = append(fruits, moreFruits...)
	f.Printf("Fruits in our fruit chat: %v \n", fruits)

	// pre allocated memory for a slices
	scoreSliceLength := 3
	scoreSliceCapacity := 5
	// creating a slice with length = 3 and capacity = 5
	scoreSlice := make([]int, scoreSliceLength, scoreSliceCapacity)
	f.Printf("scoreSlice allocated length in memory:%d\n", len(scoreSlice))
	f.Printf("scoreSlice allocated capacity in memory:%d\n", cap(scoreSlice))
	f.Printf("scoreSlice: %d\n", scoreSlice)
	f.Printf("First value of score Slice: %d\n", scoreSlice[0])

	// Looping over an array
	for index, value := range slideColors {
		f.Printf("No %d - colors: %s\n", index, value)
	}

	// HashMap or Map in Golang
	currencies := map[string]string{
		"USA":           "Dollar (US)",
		"Germany":       "Euro",
		"Great Britain": "Pound Sterling",
		"Norway":        "Krona",
	}

	for key, value := range currencies {
		f.Printf("Currency for %s is %s\n", key, value)
	}

	currency, exists := currencies["UAE"] // exist keyword is only available in map
	if exists {
		f.Printf("Currency for UAEis %s\n", currency)
	} else {
		f.Printf("Currency for UAE is not present in the map")
	}
	// to delete a key/value from map
	// delete(mapVariable, key)
}
