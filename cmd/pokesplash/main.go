package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"syscall"
	"unsafe"

	"github.com/gopik/pokesplash/internal/artloader"
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
	var err error

	// Handle specific Pokémon mode
	if *pokemonName != "" {
		art, err = artloader.GetArt(*pokemonName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			fmt.Fprintln(os.Stderr, "Use --list to see available Pokémon")
			os.Exit(1)
		}
	} else {
		// Default: random mode
		art, err = artloader.GetRandomArt()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting random Pokémon: %v\n", err)
			os.Exit(1)
		}
	}

	// Print the art
	fmt.Print(art)
}
