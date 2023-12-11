package helpers

func Gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return Gcd(b, a%b)
}

func lcm(a, b int) int {
	return (a * b) / Gcd(a, b)
}

func LcmSlice(numbers []int) int {
	if len(numbers) < 2 {
		panic("Lcm() requires at least 2 numbers")
	}

	result := numbers[0]
	remainingNumbers := numbers[1:]
	for _, number := range remainingNumbers {
		result = lcm(result, number)
	}

	return result
}
