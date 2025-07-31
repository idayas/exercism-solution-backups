package resistorcolorduo

import "strconv"

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	encodings := map[string]string{
		"black":  "0",
		"brown":  "1",
		"red":    "2",
		"orange": "3",
		"yellow": "4",
		"green":  "5",
		"blue":   "6",
		"violet": "7",
		"grey":   "8",
		"white":  "9",
	}

	total := "";
	
	for i, color := range colors {
		if i >= 2 {
			continue;
		}

		total += string(encodings[color])
	}

	res, _ := strconv.Atoi(total)
	return res
}
