package main

import (
	"bufio"
	"fmt"
	"os"
)

type ElfRange struct {
  Begin int
  End int
}

func (container *ElfRange) Contains(target ElfRange) bool {
  return container.Begin <= target.Begin && container.End >= target.End
}

func CountHowManyContainInFile(elfSectionsPath string) (int, error) {
  prioritySum := 0
  file, err := os.Open(elfSectionsPath)
  if err != nil {
    return prioritySum, err
  }
  defer file.Close()
  
  scanner := bufio.NewScanner(file)
  contains := 0

  var elfOne, elfTwo ElfRange

  for scanner.Scan() {
    fmt.Sscanf(scanner.Text(), "%d-%d,%d-%d", &elfOne.Begin, &elfOne.End, &elfTwo.Begin, &elfTwo.End)
    if elfOne.Contains(elfTwo) || elfTwo.Contains(elfOne) {
      contains += 1
    }
  }

  return contains, nil
}

func main() {
  fmt.Println(CountHowManyContainInFile("test.txt"))
  fmt.Println(CountHowManyContainInFile("input.txt"))
}
