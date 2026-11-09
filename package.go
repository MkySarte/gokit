package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

func detectPackage() (string, error) {

	entries, err := os.ReadDir(".")
	if err != nil {
		return "", fmt.Errorf("Verzeichnis konnte nicht gelesen werden: %w", err)
	}

	var packageName string

	for _, entry := range entries {

		if entry.IsDir() {
			continue
		}

		name := entry.Name()

		if !strings.HasSuffix(name, ".go") ||
			strings.HasSuffix(name, "_test.go") {
			continue
		}

		filePath := filepath.Join(".", name)

		file, err := parser.ParseFile(
			token.NewFileSet(),
			filePath,
			nil,
			parser.PackageClauseOnly,
		)
		if err != nil {
			return "", fmt.Errorf("Package konnte nicht gelesen werden: %w", err)
		}

		currentPackage := file.Name.Name

		if packageName == "" {
			packageName = currentPackage
		} else if packageName != currentPackage {
			return "", fmt.Errorf("unterschiedliche Packages im Verzeichnis gefunden")
		}
	}

	if packageName == "" {
		return "", fmt.Errorf("kein bestehendes Go-Package gefunden")
	}

	return packageName, nil
}
