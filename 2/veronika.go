package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("veronika_data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	count := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		gameAndResult := strings.Split(scanner.Text(), ":")
		result := gameAndResult[1]
		// println(game)
		// println(result)
		resultsByDraws := strings.Split(result, ";")
		// fmt.Println(resultsByDraws)
		max_r := 0
		max_g := 0
		max_b := 0
		for _, res := range resultsByDraws {
			// fmt.Println(res)
			resSplit := strings.Split(res, ", ")
			for _, r := range resSplit {
				resSplitSplit := delete_empty(strings.Split(r, " "))
				val, _ := strconv.Atoi(resSplitSplit[0]) 
				switch resSplitSplit[1] {
					case "red": 
					if (val > max_r) {
						// fmt.Print("helllo")
						max_r = val
					}
					case "blue":
						if (val > max_b) {
							// fmt.Print("helllo")
							max_b = val
					}
					case "green":
						if (val > max_g) {
							max_g = val
						}
					default:
				}
			}
		}
			count = count + (max_r * max_b * max_g)
	}
	println(count)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func delete_empty (s []string) []string {
    var r []string
    for _, str := range s {
        if str != "" {
            r = append(r, str)
        }
    }
    return r
}
