package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type ElfRange struct {
  Begin float64
  End float64
}

func (elf *ElfRange) DoNotOverlap(anotherElf ElfRange) bool {
  return anotherElf.Begin > elf.End || elf.Begin > anotherElf.End
}


func (elf *ElfRange) Overlaps(anotherElf ElfRange) float64 {
  if elf.DoNotOverlap(anotherElf) {
    return 0
  }
  return math.Max(math.Min(elf.End, anotherElf.End) - math.Max(elf.Begin, anotherElf.Begin), 1)
}

func CountOverlapping(elfSectionsPath string) (int, error) {
  prioritySum := 0
  file, err := os.Open(elfSectionsPath)
  if err != nil {
    return prioritySum, err
  }
  defer file.Close()
  
  scanner := bufio.NewScanner(file)
  var overlaps float64 = 0

  var elfOne, elfTwo ElfRange

  for scanner.Scan() {
    fmt.Sscanf(scanner.Text(), "%f-%f,%f-%f", &elfOne.Begin, &elfOne.End, &elfTwo.Begin, &elfTwo.End)
    if elfOne.Overlaps(elfTwo) > 0 {
      overlaps += 1
    }
  }

  return int(overlaps), nil
}

func main() {
  fmt.Println(CountOverlapping("test.txt"))
  fmt.Println(CountOverlapping("input.txt"))
}
