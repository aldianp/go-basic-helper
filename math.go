package helper_math

func GetFactorial(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * GetFactorial(value - 1)
	}
}