package kata
import "fmt"

func HighestRank(nums []int) int {
	var count int
	var maxCount int = 0
	var maxElement int = 0
	for _, num := range nums {
		fmt.Println(num)
		count = 0
		for _, num2 := range nums {
			if num2 == num {
				count++
				
				if count > maxCount || count ==  maxCount  && num > maxElement{
					maxElement = num
				}
			}
		}
		if count > maxCount {
			maxCount = count 
		}
			fmt.Println("Jumlah count ", count)
			fmt.Println("Jumlah maxCount ", maxCount)
			fmt.Println("maxElement ", maxElement)
	}
	return maxElement
}

