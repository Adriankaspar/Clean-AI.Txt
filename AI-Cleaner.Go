package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide the file path and name of the file you want to clean.")
		fmt.Println("go run AI-Cleaner go <file_path>")
		return
	}
	inputPath := os.Args[1]

	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
		fmt.Println("Error, Either No File Was Provided or The File Type is not supported or to large.")
		return
	}
	rawText := string(data)
	replacer := strings.NewReplacer(
		"```", "",
		"---", "",
		"##", "",
		"#", "",
		"-", "",
		"`", "",
		"**", "",
		"|", "",
		"###", "",
		"####", "",
	)
	cleanText := replacer.Replace(rawText)
	dir := filepath.Dir(inputPath)
	fileName := filepath.Base(inputPath)
	outputPath := filepath.Join(dir, "clean_"+fileName)

	err = os.WriteFile(outputPath, []byte(cleanText), 0644)
	if err != nil {
		fmt.Println("Error While Rewriting the New File.")
		return
	}
	fmt.Println("File was Succesfully Rewritten! File was saved")
}
