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

func EvalPrioritySum(elfPriorityFilePath string) (int, error) {
  prioritySum := 0
  file, err := os.Open(elfPriorityFilePath)
  if err != nil {
    return prioritySum, err
  }
  defer file.Close()
  
  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    half := len(scanner.Text()) / 2
    compartments := []string{scanner.Text()[0:half], scanner.Text()[half:len(scanner.Text())]}

    most_prioritized := 0

    for _, item := range compartments[0] {
      if strings.Contains(compartments[1], string(int(item))) {
        most_prioritized = int(math.Max(float64(asPriority(item)), float64(most_prioritized)))
      }
    }

    prioritySum += most_prioritized
  }
  return prioritySum, nil
}

func main() {
  fmt.Println(EvalPrioritySum("example.txt"))
  fmt.Println(EvalPrioritySum("input"))
}
