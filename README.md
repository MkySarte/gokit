# GoKit

GoKit ist ein Kommandozeilenwerkzeug, das typische Aufgaben rund um Go-Projekte bündelt: Projekte anlegen, Dateien generieren, bauen, testen und prüfen.

Das Projekt ist zugleich ein Lernprojekt. Jeder Meilenstein führt neue Teile der Go-Standardbibliothek ein, die unten jeweils mit der offiziellen Dokumentation verlinkt sind.

## Installation

Voraussetzung ist eine installierte Go-Toolchain (siehe `go.mod` für die Version).

```sh
git clone <repository-url>
cd gokit
go build -o gokit .     # erzeugt die ausführbare Datei im Projektordner
# oder
go install .            # installiert gokit nach $GOPATH/bin bzw. $GOBIN
```

### `gokit` in den PATH aufnehmen

`go build` legt die Datei im Projektordner ab. Gestartet wird sie dort mit `./gokit` (Windows: `.\gokit.exe`).

Damit `gokit` nach `go install` in jedem Verzeichnis funktioniert, muss der Installationsordner im `PATH` stehen. Diesen Ordner zeigt:

```sh
go env GOBIN    # falls leer, gilt der nächste Befehl
go env GOPATH   # Programme liegen dann in <GOPATH>/bin
```

Standardordner: `%USERPROFILE%\go\bin` (Windows) bzw. `~/go/bin` (Linux/macOS).

**Windows (PowerShell),** Ordner dauerhaft für den eigenen Benutzer hinzufügen:

```powershell
$bin = "$(go env GOPATH)\bin"
[Environment]::SetEnvironmentVariable("Path", [Environment]::GetEnvironmentVariable("Path", "User") + ";$bin", "User")
```

**Linux/macOS,** in `~/.bashrc` bzw. `~/.zshrc` eintragen:

```sh
export PATH="$PATH:$(go env GOPATH)/bin"
```

Danach ein neues Terminal öffnen und prüfen:

```sh
gokit --help
```

## Verwendung

```sh
gokit start <projektname>   # Erstellt ein neues Go-Projekt
gokit --help                # Zeigt die Hilfe
```

### Beispiel

```sh
gokit start meinprojekt
```

Erzeugt:

```text
meinprojekt/
├── go.mod        # über "go mod init meinprojekt"
├── README.md
└── .gitignore
```

## Befehle

| Befehl                  | Status   | Beschreibung                                  |
| ----------------------- | -------- | --------------------------------------------- |
| `gokit start <name>`    | ✅ fertig | Neues Go-Modul mit README und `.gitignore`    |
| `gokit new <name>`      | geplant  | Erzeugt `<name>.go` und `<name>_test.go`      |
| `gokit run`             | geplant  | Führt das Projekt aus (`go run`)              |
| `gokit build`           | geplant  | Baut das Projekt (`go build`)                 |
| `gokit test`            | geplant  | Führt Tests aus (`go test`)                   |
| `gokit test --coverage` | geplant  | Tests mit Abdeckungsbericht                   |
| `gokit check`           | geplant  | `gofmt`, `go vet` und `go test` in einem Lauf |
| `gokit add <modul>`     | geplant  | Fügt eine Abhängigkeit hinzu (`go get`)       |
| `gokit clean`           | geplant  | Entfernt erzeugte Dateien (`go clean`)        |
| `gokit doctor`          | geplant  | Prüft die lokale Go-Umgebung                  |

## Projektstruktur

| Datei      | Inhalt                                                  |
| ---------- | ------------------------------------------------------- |
| `main.go`  | Einstiegspunkt, liest `os.Args` und verteilt die Befehle |
| `help.go`  | Hilfetext                                               |
| `start.go` | Implementierung von `gokit start`                       |

---

# Lernplan

## Meilenstein 1 – `gokit start` (Projekte erstellen)

**Status:** umgesetzt

| Thema                     | Go-Funktion                | Dokumentation                                                  |
| ------------------------- | -------------------------- | -------------------------------------------------------------- |
| CLI-Argumente             | `os.Args`                  | [pkg.go.dev/os](https://pkg.go.dev/os#pkg-variables)           |
| Bedingungen und `switch`  | `if`, `switch`             | [Go Tour](https://go.dev/tour/flowcontrol/1)                   |
| Verzeichnisse erstellen   | `os.Mkdir()`               | [os.Mkdir](https://pkg.go.dev/os#Mkdir)                        |
| Dateien schreiben         | `os.WriteFile()`           | [os.WriteFile](https://pkg.go.dev/os#WriteFile)                |
| Dateipfade zusammensetzen | `filepath.Join()`          | [filepath.Join](https://pkg.go.dev/path/filepath#Join)         |
| Externe Prozesse          | `exec.Command()`           | [exec.Command](https://pkg.go.dev/os/exec#Command)             |
| Prozess ausführen         | `cmd.Run()`                | [Cmd.Run](https://pkg.go.dev/os/exec#Cmd.Run)                  |
| Fehler weitergeben        | `fmt.Errorf()`, `%w`       | [Error Wrapping](https://go.dev/blog/go1.13-errors)            |

**Ergebnis:** `gokit start meinprojekt` erzeugt ein Go-Modul mit README und `.gitignore`.

## Meilenstein 2 – `gokit new` (Go-Dateien generieren)

| Thema                        | Go-Funktion                 | Dokumentation                                                 |
| ---------------------------- | --------------------------- | ------------------------------------------------------------- |
| Dateien sicher erstellen     | `os.OpenFile()` mit `O_EXCL` | [os.OpenFile](https://pkg.go.dev/os#OpenFile)                |
| Verzeichnisse ermitteln      | `os.Getwd()`                | [os.Getwd](https://pkg.go.dev/os#Getwd)                       |
| Dateien prüfen               | `os.Stat()`                 | [os.Stat](https://pkg.go.dev/os#Stat)                         |
| Dateinamen bearbeiten        | `strings.TrimSuffix()`      | [strings.TrimSuffix](https://pkg.go.dev/strings#TrimSuffix)   |
| Go-Quelldateien formatieren  | `go/format`                 | [go/format](https://pkg.go.dev/go/format)                     |
| Vorlagen erzeugen            | `text/template`             | [text/template](https://pkg.go.dev/text/template)             |

**Ergebnis:** `gokit new user` erstellt `user.go` und `user_test.go`.

## Meilenstein 3 – `run`, `build`, `test`

| Befehl                  | Technische Grundlage | Dokumentation                                                                       |
| ----------------------- | -------------------- | ----------------------------------------------------------------------------------- |
| `gokit run`             | `go run`             | [go run](https://pkg.go.dev/cmd/go#hdr-Compile_and_run_Go_program)                  |
| `gokit build`           | `go build`           | [go build](https://pkg.go.dev/cmd/go#hdr-Compile_packages_and_dependencies)         |
| `gokit test`            | `go test`            | [go test](https://pkg.go.dev/cmd/go#hdr-Test_packages)                              |
| `gokit test --coverage` | Testabdeckung        | [Coverage](https://go.dev/blog/cover)                                               |

Hier kommt erneut `os/exec` zum Einsatz, das bereits aus `start.go` bekannt ist.

## Meilenstein 4 – `check`, `add`, `clean`, `doctor`

| Befehl         | Technische Grundlage              | Dokumentation                                                                                                     |
| -------------- | --------------------------------- | ----------------------------------------------------------------------------------------------------------------- |
| `gokit check`  | `gofmt`, `go vet`, `go test`      | [gofmt](https://pkg.go.dev/cmd/gofmt), [go vet](https://pkg.go.dev/cmd/vet)                                       |
| `gokit add`    | `go get`                          | [go get](https://go.dev/ref/mod#go-get)                                                                           |
| `gokit clean`  | `go clean`                        | [go clean](https://pkg.go.dev/cmd/go#hdr-Remove_object_files_and_cached_files)                                    |
| `gokit doctor` | `exec.LookPath()`, `go env`       | [LookPath](https://pkg.go.dev/os/exec#LookPath), [go env](https://pkg.go.dev/cmd/go#hdr-Print_Go_environment_information) |

Bei `clean` ist später zu unterscheiden, welche erzeugten Dateien Go selbst bereinigen kann und welche GoKit verwaltet.

## Meilenstein 5 – Tests, Fehlerbehandlung und Qualität

Für ein öffentliches Entwicklerwerkzeug besonders wichtig.

| Thema                       | Was gelernt wird | Dokumentation                                       |
| --------------------------- | ---------------- | --------------------------------------------------- |
| Unit-Tests                  | `testing.T`      | [testing](https://pkg.go.dev/testing)               |
| Temporäre Testverzeichnisse | `t.TempDir()`    | [T.TempDir](https://pkg.go.dev/testing#T.TempDir)   |
| Fehler vergleichen          | `errors.Is()`    | [errors.Is](https://pkg.go.dev/errors#Is)           |
| Fehlerarten prüfen          | `errors.As()`    | [errors.As](https://pkg.go.dev/errors#As)           |
| Fehlercodes für CLI         | `os.Exit()`      | [os.Exit](https://pkg.go.dev/os#Exit)               |
| Sichere Dateierstellung     | `os.OpenFile()`  | [os.OpenFile](https://pkg.go.dev/os#OpenFile)       |

Sinnvolle Tests für `startProject()`:

- [ ] Projekt wird erfolgreich erstellt
- [ ] Verzeichnis existiert bereits
- [ ] Externer Prozess (`go mod init`) schlägt fehl
- [ ] `README.md` und `.gitignore` haben den erwarteten Inhalt

## Meilenstein 6 – Installation und CI/CD

| Thema                         | Werkzeug / Konzept   | Dokumentation                                                              |
| ----------------------------- | -------------------- | -------------------------------------------------------------------------- |
| GoKit kompilieren             | `go build`           | [go build](https://pkg.go.dev/cmd/go#hdr-Compile_packages_and_dependencies) |
| Installation                  | `go install`         | [go install](https://go.dev/ref/mod#go-install)                            |
| Plattformübergreifende Builds | `GOOS`, `GOARCH`     | [Build-Umgebung](https://go.dev/doc/install/source#environment)            |
| GitHub CI/CD                  | GitHub Actions       | [GitHub Actions](https://docs.github.com/actions)                          |
| Releases                      | Go-Release-Prozess   | [Release Workflow](https://go.dev/doc/modules/release-workflow)            |

---

## Lizenz

GoKit ist unter der [MIT-Lizenz](LICENSE) veröffentlicht. Weitere Informationen stehen in der Datei [LICENSE](LICENSE).
