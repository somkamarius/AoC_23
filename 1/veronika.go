package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	file, err := os.Open("veronika_data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	callibrationValue := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		re := regexp.MustCompile("oneight|eightwo|sevenine|eighthree|fiveight|one|two|three|four|five|six|seven|eight|nine|\\d+")
		// re := regexp.MustCompile("[0-9]+")
		numbers := re.FindAllString(scanner.Text(), -1)
		fmt.Println(numbers)
		// fmt.Println(reflect.TypeOf(numbers))

		firstNumber := numbers[0]
		lastNumber := numbers[len(numbers)-1]
		fmt.Println("firstNumber")
		// fmt.Println(reflect.TypeOf(firstNumber))
		fmt.Println(firstNumber)
		fmt.Println("lastNumber")
		// fmt.Println(reflects.TypeOf(lastNumber))
		fmt.Println(lastNumber)

		firstElement := convertStringNumbersToString(firstNumber, true)
		lastElement := convertStringNumbersToString(lastNumber, false)
	
		fmt.Println("firstElement")
		fmt.Println(firstElement)
		fmt.Println("lastElement")
		fmt.Println(lastElement)

		callibrationValue = callibrationValue + (firstElement*10 + lastElement)
	}

	fmt.Println(callibrationValue)

	elapsed := time.Since(start)
	fmt.Println("this took ", elapsed)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func convertStringNumbersToString(number string, firstNumber bool) int {
	switch number {
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "oneight":
		if (firstNumber) {
			return 1
		} else {
			return 8
		}
	case "eightwo":
		if (firstNumber) {
			return 8
		} else {
			return 2
		}
	case "sevenine":
		if (firstNumber) {
			return 7
		} else {
			return 9
		}
	case "eighthree":
		if (firstNumber) {
			return 8
		} else {
			return 3
		}
	case "fiveight":
		if (firstNumber) {
			return 5
		} else {
			return 8
		}
	default:
		firstValue, _ := strconv.Atoi(number[0:1])
		lastValue, _ := strconv.Atoi(number[len(number)-1 : len(number)])
		if (firstNumber) {
			return firstValue
		} else {
			return lastValue
		}
	}
}
