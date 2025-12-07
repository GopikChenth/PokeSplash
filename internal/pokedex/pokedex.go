package pokedex

import (
	_ "embed"
	"encoding/json"
	"strings"
)

//go:embed pokedex.json
var pokedexJSON []byte

type PokemonName struct {
	English string `json:"english"`
}

type PokemonProfile struct {
	Height string `json:"height"`
	Weight string `json:"weight"`
}

type PokemonEntry struct {
	ID          int            `json:"id"`
	Name        PokemonName    `json:"name"`
	Type        []string       `json:"type"`
	Description string         `json:"description"`
	Profile     PokemonProfile `json:"profile"`
}

var pokedex []PokemonEntry
var nameMap map[string]PokemonEntry

func init() {
	if err := json.Unmarshal(pokedexJSON, &pokedex); err != nil {
		// handle error or just ignore (better to panic in init for embedded data issues)
		panic(err)
	}
	nameMap = make(map[string]PokemonEntry)
	for _, p := range pokedex {
		nameMap[strings.ToLower(p.Name.English)] = p
	}
}

func GetPokemon(name string) (PokemonEntry, bool) {
	// Name from file might be "nidoran-m" or "mr-mime", need to normalize
	// For now simple lookup
	p, ok := nameMap[strings.ToLower(name)]
	return p, ok
}

func GetGeneration(id int) int {
	switch {
	case id <= 151:
		return 1
	case id <= 251:
		return 2
	case id <= 386:
		return 3
	case id <= 493:
		return 4
	case id <= 649:
		return 5
	case id <= 721:
		return 6
	case id <= 809:
		return 7
	case id <= 905:
		return 8
	default:
		return 9
	}
}
