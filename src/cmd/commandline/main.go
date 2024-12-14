package main

import (
	"fmt"

	"github.com/bhvalent/adventOfCode2024/src/internal/app/day1"
)

func main() {
	fmt.Print("Enter AoC day you want to execute (1-25): ")

	var day string
	fmt.Scanln(&day)

	switch day {
	case "1":
		printDayHeader("1")
		resultPart1 := day1.GetTotalDistance("./src/internal/app/day1/input_data.txt")
		printProblemTitleAndAnswer("Part 1", resultPart1)
		resultPart2 := day1.GetSimilarityScore("./src/internal/app/day1/input_data.txt")
		printProblemTitleAndAnswer("Part 2", resultPart2)
		return
	}

	fmt.Printf("Invalid Day: %s", day)
}

func printDayHeader(day string) {
	fmt.Printf("Day %s: \n", day)
}

func printProblemTitleAndAnswer(title string, data int) {
	fmt.Printf("\t%s: %d\n", title, data)
}
