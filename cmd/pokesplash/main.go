package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strings"
	"syscall"
	"unsafe"

	"github.com/gopik/pokesplash/internal/artloader"
	"github.com/gopik/pokesplash/internal/pokedex"
)

func init() {
	if runtime.GOOS == "windows" {
		kernel32 := syscall.NewLazyDLL("kernel32.dll")
		setConsoleMode := kernel32.NewProc("SetConsoleMode")
		getConsoleMode := kernel32.NewProc("GetConsoleMode")

		var mode uint32
		handle := syscall.Handle(os.Stdout.Fd())

		r1, _, _ := getConsoleMode.Call(uintptr(handle), uintptr(unsafe.Pointer(&mode)))
		if r1 != 0 {
			mode |= 0x0004 // ENABLE_VIRTUAL_TERMINAL_PROCESSING
			setConsoleMode.Call(uintptr(handle), uintptr(mode))
		}
	}
}

func main() {
	// Define command-line flags
	pokemonName := flag.String("pokemon", "", "Display specific Pokémon by name")
	listMode := flag.Bool("list", false, "List all available Pokémon")

	flag.Parse()

	// Handle list mode
	if *listMode {
		names, err := artloader.ListArt()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error listing Pokémon: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Available Pokémon:")
		for _, name := range names {
			fmt.Printf("  - %s\n", name)
		}
		return
	}

	var art string
	var selected string // Variable to store the name of the selected Pokémon
	var err error

	// Handle specific Pokémon mode
	if *pokemonName != "" {
		selected = *pokemonName
		art, err = artloader.GetArt(*pokemonName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			fmt.Fprintln(os.Stderr, "Use --list to see available Pokémon")
			os.Exit(1)
		}
	} else {
		// Default: random mode
		art, selected, err = artloader.GetRandomArt() // Get the name of the random Pokémon
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting random Pokémon: %v\n", err)
			os.Exit(1)
		}
	}

	// Print the art
	fmt.Print(art)

	// Display metadata
	pInfo, found := pokedex.GetPokemon(selected)

	// Fallback for names with dashes or different formatting (e.g. mr-mime vs Mr. Mime)
	if !found {
		// Simple heuristic: try removing dashes or replacing with spaces
		altName := strings.ReplaceAll(selected, "-", " ")
		pInfo, found = pokedex.GetPokemon(altName)
	}

	if found {
		gen := pokedex.GetGeneration(pInfo.ID)
		weaknesses := pokedex.GetWeaknesses(pInfo.Type)

		fmt.Printf("\n")
		fmt.Printf(" Name:      %s (#%d)\n", pInfo.Name.English, pInfo.ID)
		fmt.Printf(" Type:      %s\n", strings.Join(pInfo.Type, " / "))
		fmt.Printf(" Gen:       %d\n", gen)
		fmt.Printf(" Weakness:  %s\n", strings.Join(weaknesses, ", "))
		fmt.Printf(" Info:      %s\n", pInfo.Description)
		fmt.Printf("\n")
	}
}
