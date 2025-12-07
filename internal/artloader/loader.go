package artloader

import (
	"embed"
	"fmt"
	"io/fs"
	"math/rand"
	"strings"
	"time"
)

//go:embed ascii/*
var asciiFS embed.FS

// GetRandomArt returns random Pokémon ASCII art and its name
func GetRandomArt() (string, string, error) {
	files, err := listArtFiles()
	if err != nil {
		return "", "", err
	}

	if len(files) == 0 {
		return "", "", fmt.Errorf("no art files found")
	}

	// Seed random number generator
	rand.Seed(time.Now().UnixNano())
	randomFile := files[rand.Intn(len(files))]

	art, err := readArtFile(randomFile)
	if err != nil {
		return "", "", err
	}

	// Extract name from filename (e.g., "ascii/pikachu.ans" -> "pikachu")
	name := strings.TrimPrefix(randomFile, "ascii/")
	name = strings.TrimSuffix(name, ".ans")

	return art, name, nil
}

// GetArt returns the ASCII art for a specific Pokémon
func GetArt(name string) (string, error) {
	filename := fmt.Sprintf("ascii/%s.ans", strings.ToLower(name))
	return readArtFile(filename)
}

// ListArt returns a list of available Pokémon names
func ListArt() ([]string, error) {
	files, err := listArtFiles()
	if err != nil {
		return nil, err
	}

	var names []string
	for _, file := range files {
		// Remove "ascii/" prefix and ".ans" suffix
		name := strings.TrimPrefix(file, "ascii/")
		name = strings.TrimSuffix(name, ".ans")
		names = append(names, name)
	}

	return names, nil
}

// listArtFiles returns all .ans files in the ascii directory
func listArtFiles() ([]string, error) {
	var files []string

	err := fs.WalkDir(asciiFS, "ascii", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && strings.HasSuffix(path, ".ans") {
			files = append(files, path)
		}

		return nil
	})

	return files, err
}

// readArtFile reads and returns the content of an art file
func readArtFile(filename string) (string, error) {
	data, err := asciiFS.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("art file '%s' not found", filename)
	}

	return string(data), nil
}
