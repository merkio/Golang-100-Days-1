package main

import (
	"fmt"
	"github.com/bregydoc/gtranslate"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	err := filepath.Walk(".", translate)
	if err != nil {
		log.Println(err)
	}
}

func translate(path string, info os.FileInfo, err error) error {
	if strings.Contains(path, ".md") {
		builder := strings.Builder{}
		filename, err := gtranslate.TranslateWithParams(
			path,
			gtranslate.TranslationParams{
				From: "zh-CN",
				To:   "en",
			},
		)
		if err != nil {
			panic(err)
		}
		//filename = strings.ReplaceAll(filename, " ", "_")
		fmt.Printf("Write translation to the file: %s\n", filename)
		sContent, _ := ioutil.ReadFile(path)
		content := string(sContent)
		for _, line := range strings.Split(content, "\n") {
			text := strings.TrimSpace(line)
			if text != "" {
				translated, err := gtranslate.TranslateWithParams(
					text,
					gtranslate.TranslationParams{
						From: "zh-CN",
						To:   "en",
					},
				)
				if err != nil {
					panic(err)
				}
				builder.WriteString(fmt.Sprintf("%s\n", translated))
			}
		}
		fmt.Println("Close the file: " + filename)
		_ = os.MkdirAll(filepath.Dir(filename), 0755)
		_ = ioutil.WriteFile(filename, []byte(builder.String()), 0755)
	}

	return nil
}
