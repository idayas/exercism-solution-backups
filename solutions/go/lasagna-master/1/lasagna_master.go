package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, prepTimePerLayer int) int {
	if prepTimePerLayer == 0 {
		prepTimePerLayer = 2
	}

	return prepTimePerLayer * len(layers)
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	noodleLayers := 0
	sauceLayers := 0

	for _, layer := range layers {
		if layer == "noodles" {
			noodleLayers++
		} else if layer == "sauce" {
			sauceLayers++
		}
	}

	sauce = float64(sauceLayers) * 0.2
	noodles = noodleLayers * 50

	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList []string, myList []string) {
	myList[len(myList)-1] = friendList[len(friendList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) (scaled []float64) {
	basePortions := 2.0
	factor := float64(portions) / basePortions

	scaled = make([]float64, len(quantities))

	for i, qty := range quantities {
		scaled[i] = qty * factor
	}

	return scaled
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
