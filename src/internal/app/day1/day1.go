package day1

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"

	"github.com/bhvalent/adventOfCode2024/src/internal/utilities"
)

func GetTotalDistance(filename string) int {
	// clean/format data into 2 slices
	lines, err := utilities.GetLinesFromFile(filename)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return 0
	}
	left, right := getLocationLists(lines)

	// sort slices
	sort.Ints(left)
	sort.Ints(right)

	// loop through slices to find distance
	totalDistance := 0
	for i := 0; i < len(left); i++ {
		totalDistance += utilities.IntAbs(left[i] - right[i])
	}

	return totalDistance
}

func GetSimilarityScore(filename string) int {
	// clean/format data into 2 slices
	lines, err := utilities.GetLinesFromFile(filename)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return 0
	}
	left, right := getLocationLists(lines)

	// get similarity score
	similarityScore := 0
	frequenciesInRight := countFreqencies(right)
	for _, num := range left {
		similarityScore += (num * frequenciesInRight[num])
	}

	return similarityScore
}

func getLocationLists(data []string) ([]int, []int) {
	left := []int{}
	right := []int{}

	pattern := `^(\d+)\s+(\d+)$`
	re := regexp.MustCompile(pattern)

	for i := 0; i < len(data); i++ {
		regexResults := re.FindStringSubmatch(data[i])

		if len(regexResults) == 3 {
			leftNum, _ := strconv.Atoi(regexResults[1])
			left = append(left, leftNum)

			rightNum, _ := strconv.Atoi((regexResults[2]))
			right = append(right, rightNum)

			continue
		}

		fmt.Printf("Unable to read values from line: %v", i+1)
	}

	return left, right
}

func countFreqencies(list []int) map[int]int {
	frequencies := make(map[int]int)
	for _, num := range list {
		frequencies[num]++
	}

	return frequencies
}
