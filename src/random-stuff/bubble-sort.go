package main

import (
	"fmt"
	"os"
	"strings"
)

func main(){

	// declare a slice of integers
	var nums []int 

	// for storing the user input integer
	var input int

	// sentinel value to stop taking more inputs
	var quit string

	fmt.Println("\n\nEnter up to 10 integers that needs to be sorted from least to greatest...\n")

	for {
		
		fmt.Println("\nEnter an integer ")
		_, err := fmt.Scanf("%d",&input)
		if err !=nil {
			fmt.Println(err)
			os.Exit(1)
		}
		nums= append(nums, input)

		fmt.Println("Please enter X to sort or any other key to continue entering ")
		fmt.Scanf("%s", &quit)

		if len(nums)== 10 ||  strings.ToUpper(quit) == "X" {
			BubbleSort(nums)
			break
		}
		
	}
	
	// display the sorted integer
	fmt.Println("\n\n\nThe Sorted List of Integers - ")
	fmt.Println(nums)
}

// sort the sequence using bubble sort algo
func BubbleSort(nums []int){

	// sentinel value indicating if slice is sorted
    isSorted := false
	
	// to optimize the sweep of the slice 
	lastUnsorted := len(nums) -1

	for {
		
		// continue until the slice is sorted 
		if isSorted {
			break
		}
		
		isSorted = true
		
		// sweep the slice upto last but one 
		for i := 0; i< lastUnsorted; i++ {
			if nums[i] > nums[i+1] {
				Swap(nums, i)
				isSorted = false
			} 
		}
		
		// after every sweep of the slice the last element is at its correct position
		lastUnsorted = lastUnsorted - 1 
	}

}

// swap two adjacent elements
func Swap(nums []int, i int){

	temp := nums[i]
	nums[i] = nums[i+1]
	nums[i+1] = temp

}