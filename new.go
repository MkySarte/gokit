package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

func addNewfile(fileName string) error {

	packageName, err := detectPackage()
	if err != nil {
		return fmt.Errorf("Package-Erkennung fehlgeschlagen: %w", err)
	}

	fileName = strings.TrimSuffix(fileName, ".go")

	if fileName == "" || fileName == "." || fileName == ".." ||
		strings.ContainsAny(fileName, `/\`) {
		return fmt.Errorf("ungültiger Dateiname: %q", fileName)
	}

	path := fileName + ".go"
	pathTest := fileName + "_test.go"

	firstRune, size := utf8.DecodeRuneInString(fileName)
	testName := string(unicode.ToUpper(firstRune)) + fileName[size:]

	data := TemplateData{
		PackageName: packageName,
		TestName:    testName,
	}

	ft, err := os.OpenFile(pathTest, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("Testdatei konnte nicht erstellt werden: %w", err)
	}

	f, err := os.OpenFile(path, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0644)
	if err != nil {
		closeErr := ft.Close()
		removeErr := os.Remove(pathTest)

		return errors.Join(
			fmt.Errorf("Go-Datei konnte nicht erstellt werden: %w", err),
			closeErr,
			removeErr,
		)
	}

	rollback := func(cause error) error {
		closeErr1 := f.Close()
		closeErr2 := ft.Close()
		removeErr1 := os.Remove(path)
		removeErr2 := os.Remove(pathTest)

		return errors.Join(cause, closeErr1, closeErr2, removeErr1, removeErr2)
	}

	if err := writeTemplate(f, goFileTemplate, data); err != nil {
		return rollback(err)
	}

	if err := writeTemplate(ft, goTestTemplate, data); err != nil {
		return rollback(err)
	}

	closeErr1 := f.Close()
	closeErr2 := ft.Close()

	if err := errors.Join(closeErr1, closeErr2); err != nil {
		removeErr1 := os.Remove(path)
		removeErr2 := os.Remove(pathTest)

		return errors.Join(
			fmt.Errorf("Dateien konnten nicht geschlossen werden: %w", err),
			removeErr1,
			removeErr2,
		)
	}

	fmt.Printf("%s und %s wurden erstellt (package %s)\n",
		path, pathTest, packageName)

	return nil
}
