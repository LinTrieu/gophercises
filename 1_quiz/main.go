package main

import (
	"encoding/csv"
	"flag"
	"os"
	"fmt"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()

	// Open function needs the filename string; *csvFileName is a pointer to a string
	file, err := os.Open(*csvFileName)

	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFileName))
	}
	r := csv.NewReader(file)

	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file,")
	}
	fmt.Println(lines)
}


func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}