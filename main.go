package main

import (
	"flag"
	"fmt"
	"os"
)

var sourceLang string
var targetLang string
var sourceText string

func init() {
	flag.StringVar(&sourceLang, "s", "en", "Source Langeuage")
	flag.StringVar(&targetLang, "t", "fr", "Target Language")
	flag.StringVar(&sourceText, "st", "", "Text to translate")
}

func main() {
	flag.Parse()

	if flag.NFlag() == 0 {
		fmt.Println("Options: ")
		flag.PrintDefaults()
		os.Exit(1)
	}

	reqBody := &cli.RequestBody{
		SourceLang: sourcesourceLang,
		TargetLang: targetLang,
		SourceText: sourcesourceText,
	}

	cli.RquestTranslate(reqBod)
}
