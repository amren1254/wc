//Write Your Own wc Tool
//wc - print newline, word, and byte counts for each file
//https://codingchallenges.fyi/challenges/challenge-wc

/*
1.

	In this step your goal is to write a simple version of wc, let’s call it ccwc (cc for Coding Challenges)
	 that takes the command line option -c and outputs the number of bytes in a file.

2.

	In this step your goal is to support the command line option -l that outputs the number of lines in a file.

3.

	In this step your goal is to support the command line option -w that outputs the number of words in a file.

4.

	In this step your goal is to support the command line option -m that outputs the number of characters in a file.
	If the current locale does not support multibyte characters this will match the -c option.
*/
package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func readFile(filename string) []byte {
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Printf("Error reading file %v", err)
		return nil
	}
	return file
}

func main() {
	//read text file
	const appName = "word_count"
	log.SetPrefix(appName + "\t")

	cmdArgs := os.Args

	if cmdArgs[1] != "-c" && cmdArgs[1] != "-l" && cmdArgs[1] != "-w" {
		log.Printf("Please use the correct command line argument to pass file name")
		return
	}

	fileName := cmdArgs[2]
	fileData := readFile(fileName)
	if fileData == nil {
		log.Printf("Invalid data read from file %v", fileData)
		return
	}
	switch cmdArgs[1] {
	case "-c":
		fmt.Printf("%d %s \n", countChar(fileData), fileName)
	case "-l":
		fmt.Printf("%d %s \n", countLines(fileData), fileName)
	case "-w":
		fmt.Printf("%d %s \n", countWords(fileData), fileName)
	default:
		log.Printf("Invalid command line arguments %v", cmdArgs)
	}
}

func countChar(data []byte) int {
	//return len(data)
	var wordCount int
	for range data {
		wordCount++
	}
	return wordCount
}

func countLines(data []byte) int {
	return len(strings.Split(string(data), "\n")) - 1

}

func countWords(data []byte) int {
	s := string(data)
	lines := strings.Split(s, "\n")
	var wordCount int
	for _, line := range lines {
		words := strings.Split(line, " ")
		wordCount += len(words)
	}
	return wordCount
}
