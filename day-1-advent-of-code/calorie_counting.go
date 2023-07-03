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

  return elvesCalories, nil
}

func getHigherElvesCalories(elvesFilePath string) int {
  elvesCalories, err := parseElvesFile(elvesFilePath)
  if err != nil {
    log.Fatal("Can't parse file")
  }
  var greatest_cal int

  for _, elfCalories := range elvesCalories {
    elfCalorie := 0
    for _, cal := range elfCalories {
      elfCalorie += cal
    }
    if elfCalorie > greatest_cal {
      greatest_cal = elfCalorie
    }
  }
  return greatest_cal
}

func main() {
  fmt.Println(getHigherElvesCalories("test-cases/base.txt"))
  fmt.Println(getHigherElvesCalories("test-cases/my-input.txt"))
}
