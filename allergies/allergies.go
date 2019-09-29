package allergies

var allergens = map[int]string{
	1: "eggs",
	2: "peanuts",
	4: "shellfish",
	8: "strawberries",
	16: "tomatoes",
	32: "chocolate",
	64: "pollen",
	128: "cats",
}

const MAXIMUM_INDEX = 128

func Allergies(score uint) []string {
	if score == 0 {
		return []string{}
	}
	var report = []string{}
	for allergenIndex := 1; allergenIndex <= MAXIMUM_INDEX; allergenIndex *= 2 {
		var allergen = allergens[allergenIndex]
		if uint(allergenIndex) & score != 0 {
			report = append(report, allergen)
		}
	}
	return report
}

func AllergicTo(score uint, substance string) bool {
	for _, allergen := range Allergies(score) {
		if allergen == substance {
			return true
		}
	}
	return false
}