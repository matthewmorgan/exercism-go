package allergies

var allergens = []string{
	"eggs",
	"peanuts",
	"shellfish",
	"strawberries",
	"tomatoes",
	"chocolate",
	"pollen",
	"cats",
}

func Allergies(score uint) []string {
	if score == 0 {
		return []string{}
	}
	var intScore = int(score)
	var report = []string{}
	for i, allergen := range allergens {
		if 1 << uint(i) & intScore != 0 {
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