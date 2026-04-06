package arraysandslices

func Sum(nums []int) int{
	sum := 0
	//standard for
	// for i := 0; i < 5; i++{
	// 	sum += nums[i]
	// }
	for _, number := range nums{
		sum += number
	}
	return sum;
}

//SumAll is a variadic function. It can accept a variable size input
// that ultimately resolves to a slice within the function.
func SumAll(numbersToSum...[]int) (sums []int) {
	// lengthOfNumbers := len(numbersToSum)
	// sums := make([]int, lengthOfNumbers)//generates an empty slice with a start capacity of length

	
	//*** will still err if we try to say slice[10] = 1 because its initialized for a fixed capacity (if our capacity was 2 for instance.)
	// for i, numbers := range numbersToSum{
	// 	sums[i] = Sum(numbers)// reusing our prev function to get a sum for each slice.
	// }

	for _, numbers := range numbersToSum{
		//creates a new slice with all of the items
		// in this way, the slice is dynamic.
		sums = append(sums, Sum(numbers))
	} 
	return 
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums[]int
	for _, numbers := range numbersToSum{
		if len(numbers) == 0{
			sums = append(sums,0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}