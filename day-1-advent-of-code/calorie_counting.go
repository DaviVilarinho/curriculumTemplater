package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func parseElvesFile(filePath string) ([][]int, error) {
	var elvesCalories [][]int

	file, err := os.Open(filePath)
	if err != nil {
		return elvesCalories, errors.New("Could not open " + filePath)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var loopElfCalories []int
	for scanner.Scan() {
		if scanner.Text() != "" && scanner.Text() != "\n" {
			calories, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println(elvesCalories)
				fmt.Println(loopElfCalories)
				log.Fatal("Can't parse elves file: {" + scanner.Text() + "}")
			}
			loopElfCalories = append(loopElfCalories, calories)
		} else {
			elvesCalories = append(elvesCalories, loopElfCalories)
			loopElfCalories = make([]int, 0)
		}
	}
  elvesCalories = append(elvesCalories, loopElfCalories)

	return elvesCalories, nil
}

func sumSlice(sliceValues []int) int {
	sum := 0
	for _, v := range sliceValues {
		sum += v
	}
	return sum
}

func refreshTop3(top3 []int, newCal int) []int {
	for i := 0; i < 3; i++ {
		if newCal > top3[i] {
			for j := 3 - 1; j > i; j-- {
				top3[j] = top3[j-1]
			}
			top3[i] = newCal
			return top3
		}
	}
	return top3
}

func getHigherElvesCalories(elvesFilePath string) int {
	elvesCalories, err := parseElvesFile(elvesFilePath)
	if err != nil {
		log.Fatal("Can't parse file")
	}
	var greatestCal int

	for _, elfCalories := range elvesCalories {
		elfCalorie := 0
		for _, cal := range elfCalories {
			elfCalorie += cal
		}
		if elfCalorie > greatestCal {
			greatestCal = elfCalorie
		}
	}
	return greatestCal
}

func getTop3ElvesCalories(elvesFilePath string) int {
	elvesCalories, err := parseElvesFile(elvesFilePath)
	if err != nil {
		log.Fatal("Can't parse file")
	}
	var greatestCalsTop3 = []int{0,0,0}

	for _, elfCalories := range elvesCalories {
		refreshTop3(greatestCalsTop3, sumSlice(elfCalories))
	}
	return sumSlice(greatestCalsTop3)
}

func main() {
	fmt.Println(getHigherElvesCalories("test-cases/base.txt"))
	fmt.Println(getHigherElvesCalories("test-cases/my-input.txt"))
  fmt.Println("Top3")
  fmt.Println(getTop3ElvesCalories("test-cases/base.txt"))
	fmt.Println(getTop3ElvesCalories("test-cases/my-input.txt"))
}
