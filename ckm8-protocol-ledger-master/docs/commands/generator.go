package main

import (
	"log"
	"strings"

	"github.com/spf13/cobra/doc"
	ckm8 "https://github.com/fsmile2/ckm8/cmd/ckm8/cmd"
	ckm8cli "https://github.com/fsmile2/ckm8/cmd/ckm8cli/cmd"
)

func generateckm8CLIDoc(filePrepender, linkHandler func(string) string) {
	var all = ckm8cli.RootCmd
	err := doc.GenMarkdownTreeCustom(all, "./wallet/", filePrepender, linkHandler)
	if err != nil {
		log.Fatal(err)
	}
}

func generateckm8Doc(filePrepender, linkHandler func(string) string) {
	var all = ckm8.RootCmd
	err := doc.GenMarkdownTreeCustom(all, "./ledger/", filePrepender, linkHandler)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	filePrepender := func(filename string) string {
		return ""
	}

	linkHandler := func(name string) string {
		return strings.ToLower(name)
	}

	generateckm8CLIDoc(filePrepender, linkHandler)
	generateckm8Doc(filePrepender, linkHandler)
	Walk()
}
