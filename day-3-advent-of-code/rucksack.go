package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func asPriority(ch rune) int {
  if ch >= 'a' {
    return int(ch - 'a') + 1
  }
  return int(ch - 'A') + 1 + 26
}

func findCommonPriority(rucksacks []string) (int, error) {
  for _, item := range rucksacks[0] {
    hasInSecondRucksack := strings.Contains(rucksacks[1], string(int(item))) 
    hasInThirdRucksack := strings.Contains(rucksacks[2], string(int(item))) 
    if hasInSecondRucksack && hasInThirdRucksack {
      return asPriority(item), nil
    }
  }
  return 0, nil
}

func EvalPrioritySum(elfPriorityFilePath string) (int, error) {
  prioritySum := 0
  file, err := os.Open(elfPriorityFilePath)
  if err != nil {
    return prioritySum, err
  }
  defer file.Close()
  
  scanner := bufio.NewScanner(file)
  group := []string{}
  line := 0

  for scanner.Scan() {
    group = append(group, scanner.Text())
    if math.Mod(float64(line + 1), 3) == 0 {
      mostPrioritized, err := findCommonPriority(group)
      if err != nil {
        panic("can't handle group")
      }
      prioritySum += mostPrioritized
      group = []string{}
    }
    line += 1
  }

  return prioritySum, nil
}

func main() {
  fmt.Println(EvalPrioritySum("example.txt"))
  fmt.Println(EvalPrioritySum("input"))
}
