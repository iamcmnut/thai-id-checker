package thaichecker

import "strconv"

func Validate(thaiIDNumber string) bool {
	if len(thaiIDNumber) != 13 {
		return false
	}

	_, err := strconv.Atoi(thaiIDNumber)
	if err != nil {
		return false
	}

	lastDigit, _ := strconv.Atoi(thaiIDNumber[12:])

	sum := 0

	for i := 0; i < len(thaiIDNumber)-1; i++ {
		num, _ := strconv.Atoi(thaiIDNumber[i : i+1])
		sum += num * (13 - i)
	}

	sum = sum % 11
	checkDigit := 11 - sum
	checkDigit = checkDigit % 10

	return checkDigit == lastDigit
}
