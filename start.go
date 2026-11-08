package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func startProject(projectName string) error {

	err := os.Mkdir(projectName, 0755)

	if err != nil {
		return fmt.Errorf("Projektordner konnte nicht erstellt werden: %w", err)
	}

	cmd := exec.Command("go", "mod", "init", projectName)

	cmd.Dir = projectName

	err = cmd.Run()

	if err != nil {
		return fmt.Errorf("go mod init fehlgeschlagen: %w", err)
	}

	// 3. Pfad zur README.md erstellen
	readmePath := filepath.Join(projectName, "README.md")

	// 4. Inhalt der README.md festlegen
	content := []byte("# " + projectName + "\n")

	// 5. README.md erstellen und Inhalt schreiben
	err = os.WriteFile(readmePath, content, 0644)
	if err != nil {
		return fmt.Errorf("README.md konnte nicht erstellt werden: %w", err)
	}

	gitignorePath := filepath.Join(projectName, ".gitignore")

	content = []byte(`# Kompilierte Binärdateien (ausgeführt mit go build)
*.exe
*.exe~
*.dll
*.so
*.dylib
*.test
*.prof

# Falls du deine ausführbare Datei standardmäßig benennst (z.B. wie das Verzeichnis)
# Füge hier spezifische Binärnamen hinzu, z. B.:
# /mein-go-projekt

# Lokale Umgebungsvariablen (Geheimnisse/Passwörter)
.env
*.env
.env.local

# Go-spezifische Tool-Verzeichnisse
/vendor/
# Hinweis: Wenn du Dependency-Vendoring nutzt und den Code einchecken willst,
# kommentiere die obige Zeile mit # aus.

# Lokale Go-Workspace-Konfigurationen
go.work
go.work.sum

# Test-Coverage-Berichte
*.out
*.coverprofile

# IDEs und Editoren
.idea/
.vscode/
*.suo
*.ntvs*
*.njsproj
*.sln
*.swp
*~`)

	err = os.WriteFile(gitignorePath, content, 0644)
	if err != nil {
		return fmt.Errorf(".gitignore konnte nicht erstellt werden: %w", err)
	}

	return nil
}
