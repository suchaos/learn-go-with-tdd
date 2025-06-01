package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)
	for index, values := range numbersToSum {
		sums[index] = Sum(values)
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, values := range numbersToSum {
		sum := 0
		if len(values) > 0 {
			sum = Sum(values[1:])
		}
		sums = append(sums, sum)
	}
	return sums
}
