package pokedex

// Simple Type Chart (Gen 6+)
// 0 = Normal, 0.5 = Resistant, 2 = Weak, 0 = Immune
// Map[Attacker][Defender] -> Multiplier
var typeChart = map[string]map[string]float64{
	"Normal":   {"Rock": 0.5, "Ghost": 0, "Steel": 0.5},
	"Fire":     {"Fire": 0.5, "Water": 0.5, "Grass": 2, "Ice": 2, "Bug": 2, "Rock": 0.5, "Dragon": 0.5, "Steel": 2},
	"Water":    {"Fire": 2, "Water": 0.5, "Grass": 0.5, "Ground": 2, "Rock": 2, "Dragon": 0.5},
	"Electric": {"Water": 2, "Electric": 0.5, "Grass": 0.5, "Ground": 0, "Flying": 2, "Dragon": 0.5},
	"Grass":    {"Fire": 0.5, "Water": 2, "Grass": 0.5, "Poison": 0.5, "Ground": 2, "Flying": 0.5, "Bug": 0.5, "Rock": 2, "Dragon": 0.5, "Steel": 0.5},
	"Ice":      {"Fire": 0.5, "Water": 0.5, "Grass": 2, "Ice": 0.5, "Ground": 2, "Flying": 2, "Dragon": 2, "Steel": 0.5},
	"Fighting": {"Normal": 2, "Ice": 2, "Poison": 0.5, "Flying": 0.5, "Psychic": 0.5, "Bug": 0.5, "Rock": 2, "Ghost": 0, "Dark": 2, "Steel": 2, "Fairy": 0.5},
	"Poison":   {"Grass": 2, "Poison": 0.5, "Ground": 0.5, "Rock": 0.5, "Ghost": 0.5, "Steel": 0, "Fairy": 2},
	"Ground":   {"Fire": 2, "Electric": 2, "Grass": 0.5, "Poison": 2, "Flying": 0, "Bug": 0.5, "Rock": 2, "Steel": 2},
	"Flying":   {"Electric": 0.5, "Grass": 2, "Fighting": 2, "Bug": 2, "Rock": 0.5, "Steel": 0.5},
	"Psychic":  {"Fighting": 2, "Poison": 2, "Psychic": 0.5, "Dark": 0, "Steel": 0.5},
	"Bug":      {"Fire": 0.5, "Grass": 2, "Fighting": 0.5, "Poison": 0.5, "Flying": 0.5, "Psychic": 2, "Ghost": 0.5, "Dark": 2, "Steel": 0.5, "Fairy": 0.5},
	"Rock":     {"Fire": 2, "Ice": 2, "Fighting": 0.5, "Ground": 0.5, "Flying": 2, "Bug": 2, "Steel": 0.5},
	"Ghost":    {"Normal": 0, "Psychic": 2, "Ghost": 2, "Dark": 0.5},
	"Dragon":   {"Dragon": 2, "Steel": 0.5, "Fairy": 0},
	"Dark":     {"Fighting": 0.5, "Psychic": 2, "Ghost": 2, "Dark": 0.5, "Fairy": 0.5},
	"Steel":    {"Fire": 0.5, "Water": 0.5, "Electric": 0.5, "Ice": 2, "Rock": 2, "Steel": 0.5, "Fairy": 2},
	"Fairy":    {"Fire": 0.5, "Fighting": 2, "Poison": 0.5, "Dragon": 2, "Dark": 2, "Steel": 0.5},
}

var allTypes = []string{"Normal", "Fire", "Water", "Electric", "Grass", "Ice", "Fighting", "Poison", "Ground", "Flying", "Psychic", "Bug", "Rock", "Ghost", "Dragon", "Dark", "Steel", "Fairy"}

// GetWeaknesses returns types that deal >1x damage to the given types
func GetWeaknesses(defendTypes []string) []string {
	weaknessMap := make(map[string]float64)

	// Initialize with 1.0 effectiveness
	for _, t := range allTypes {
		weaknessMap[t] = 1.0
	}

	// Calculate effectiveness for each defending type
	for _, dType := range defendTypes {
		for _, aType := range allTypes {
			// Get multiplier from chart (default 1.0 if not explicit)
			mult := 1.0
			if validMap, ok := typeChart[aType]; ok {
				if val, exists := validMap[dType]; exists {
					mult = val
				}
			}
			weaknessMap[aType] *= mult
		}
	}

	// Collect types with > 1.0 effectiveness
	var weakTypes []string
	for _, t := range allTypes {
		if weaknessMap[t] > 1.0 {
			weakTypes = append(weakTypes, t)
		}
	}
	return weakTypes
}
