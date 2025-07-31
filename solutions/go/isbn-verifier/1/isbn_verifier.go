package isbn

func IsValidISBN(isbn string) bool {
	if len(isbn) == 0 {
		return false
	}

	iterator := 0
	total := 0
	for _, c := range isbn {
		intVar := 0

		if (c == 'X' || c == 'x') && iterator == 9 {
			intVar = 10
		} else if c <= '9' && c >= '0' {
			intVar = int(c - '0')
		} else if c == '-' {
			continue
		} else {
			return false
		}

		total += (10 - iterator) * intVar
		iterator++
	}

	if iterator > 10 || iterator < 10 {
		return false
	}

	return total%11 == 0

}
