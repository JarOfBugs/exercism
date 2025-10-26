package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgPrepTime int) int {
    if avgPrepTime == 0 {
        avgPrepTime = 2
    }
    return len(layers) * avgPrepTime
} 

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodle int, sauce float64) {
    noodle = 0
    sauce = 0.0
    for i := 0; i < len(layers); i++ {
        layer := layers[i]
        if layer == "noodles" {
            noodle += 50
        } else if layer == "sauce" {
            sauce += 0.2
        }
    }
    return
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList, myList []string) {
    myList[len(myList)-1] = friendList[len(friendList)-1]
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amounts []float64, portions int) []float64{
    newAmounts := make([]float64, len(amounts))
    for i := 0; i < len(amounts); i++ {
        newAmounts[i] = amounts[i] * (float64(portions) / 2.0)
    }
    return newAmounts
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
