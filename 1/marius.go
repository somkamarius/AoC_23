package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var sum = 0
	for _, line := range readFile() {
		sum += getSingleLineRes(line)
	}
	fmt.Println("answer is ", sum)
}

func getSingleLineRes(str string) int {
	var first, last = 0, 0
	for _, ch := range str {
		val, err := strconv.Atoi(string(ch))
		if err == nil {
			if first == 0 {
				first = val
			} else {
				last = val
			}
		}
	}
	if last == 0 {
		last = first
	}
	fmt.Println(str, "|", first, last)
	return first*10 + last
}

func readFile() []string {
	array1 := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	array2 := []string{"one1one", "two2two", "three3three", "four4four", "five5five", "six6six", "seven7seven", "eight8eight", "nine9nine"}
	readFile, err := os.Open("testdata.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		line := fileScanner.Text()
		line = strings.NewReplacer(zip(array1, array2)...).Replace(line)
		fileLines = append(fileLines, line)
	}
	readFile.Close()

	return fileLines
}

func zip(a1, a2 []string) []string {
	r := make([]string, 2*len(a1))
	for i, e := range a1 {
		r[i*2] = e
		r[i*2+1] = a2[i]
	}
	return r
}
