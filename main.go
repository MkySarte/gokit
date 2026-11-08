package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Kein Befehl angegeben.")
		help()
		return
	}

	command := os.Args[1]

	switch command {
	case "--help":
		help()

	case "start":

		if len(os.Args) < 3 {
			fmt.Println("Fehler: Projektname fehlt.")
			return
		}

		projectName := os.Args[2]
		err := startProject(projectName)

		if err != nil {
			fmt.Println("Fehler:", err)
			return
		}

		fmt.Println("Projekt erfolgreich erstellt:", projectName)

	default:
		fmt.Println("Unbekannter Befehl:", command)
		help()
	}
}
