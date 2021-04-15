package main

import (
	"fmt"
)

func TwoSum(numbers []int, target int) [2]int { 
	for i, num := range numbers {
		fmt.Println(num, i)
		for ii := i+1; ii < len(numbers); ii++ {
			fmt.Println("numbers i", numbers[i])
			fmt.Println("numbers ii", numbers[ii])
			
			our_target := numbers[i] + numbers[ii]
			fmt.Println("our_target", our_target)
			if our_target == target {
				fmt.Println("Yeay i found it!")
				return [2]int{i, ii}
			}
		}
	}
  return [2]int{}
}

func main() {
	result := TwoSum([]int{1, 2, 3}, 4)
	fmt.Println(result)
	fmt.Println("Hello, playground")
}
