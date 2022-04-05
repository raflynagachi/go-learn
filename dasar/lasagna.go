package main

import "fmt"

// TODO: define the 'PreparationTime()' function
func PreparationTime(layer []string, min int) int {
	if min == 0 {
		min = 2
	}
	return len(layer) * min
}

// TODO: define the 'Quantities()' function
func Quantities(layer []string) (noodles int, sauce float64) {
	for _, val := range layer {
		if val == "noodles" {
			noodles += 50
		} else if val == "sauce" {
			sauce += 0.2
		}
	}
	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(recipe1 []string, recipe2 *[]string) {
	last := recipe1[len(recipe1)-1]
	*recipe2 = append((*recipe2)[:len(*recipe2)-1], last)
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(portion []float64, numPortion int) (result []float64) {
	numPortion /= 2
	for i, v := range portion {
		result[i] = v * float64(numPortion)
	}
	return
}

func main() {
	resep := []string{"mantap", "?"}
	AddSecretIngredient([]string{"mantap", "sekali"}, &resep)
	fmt.Println(resep)
}
