package fizzbuzz

import "strconv"

// Fizzbuzz : fizzbuzz function
func Fizzbuzz(number int) string {
	if number%15 == 0 {
		return "FizzBuzz"
	}
	if number%3 == 0 {
		return "Fizz"
	}
	if number%5 == 0 {
		return "Buzz"
	}

	return strconv.Itoa(number)
}
