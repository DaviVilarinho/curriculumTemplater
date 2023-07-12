package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

/*
input example:
[P]     [L]         [T]
[L]     [M] [G]     [G]     [S]
[M]     [Q] [W]     [H] [R] [G]
[N]     [F] [M]     [D] [V] [R] [N]
[W]     [G] [Q] [P] [J] [F] [M] [C]
[V] [H] [B] [F] [H] [M] [B] [H] [B]
[B] [Q] [D] [T] [T] [B] [N] [L] [D]
[H] [M] [N] [Z] [M] [C] [M] [P] [P]
 1   2   3   4   5   6   7   8   9

move 8 from 3 to 2
move 1 from 9 to 5
move 5 from 4 to 7
move 6 from 1 to 4
move 8 from 6 to 8
move 8 from 4 to 5
move 4 from 9 to 5
move 4 from 7 to 9
move 7 from 7 to 2
*/

/*
output
what ends up in each stack?
*/

type FakeStack []rune

func (s *FakeStack) Peek() (rune, error) {
  if s.IsEmpty() {
    return ' ', errors.New("Empty Stack")
  }
  return (*s)[len(*s)-1], nil
}

func (s *FakeStack) IsEmpty() bool {
  return len(*s) == 0
}

type CratesStacks struct {
  Stacks map[int]FakeStack
}

func NewCratesStacks() *CratesStacks {
  return &CratesStacks{Stacks: make(map[int]FakeStack)}
}

func ParseInputFileStackPartIntoCrates(lines []string) *CratesStacks {
  lastLineIndex := len(lines) - 1
  lastLineSize := len(lines[lastLineIndex])
  crates := NewCratesStacks()

  for column := 0; column < lastLineSize; column++ {
    if unicode.IsDigit(rune(lines[lastLineIndex][column])) {
      stackNumber, err := strconv.Atoi(string(lines[lastLineIndex][column]))
      if err != nil {
        panic("can't parse stackNumber")
      }

      stack := FakeStack{}
      for stackColumnElementIndex := lastLineIndex - 1; stackColumnElementIndex >= 0; stackColumnElementIndex-- {
        runeElement := rune(lines[stackColumnElementIndex][column])
        if unicode.IsLetter(runeElement) {
          stack = append(stack, runeElement)
        }
      }
      crates.Stacks[stackNumber] = stack
    }
  }
  return crates
}

type Log struct {
  HowMany int
  FromStack int
  ToStack int
}

func ParseInputAndProcessLogs(pathToFile string) (*CratesStacks, error) {
  file, err := os.Open(pathToFile) 
  if err != nil {
    return nil, err
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  var cratesStacks *CratesStacks
  var cratesText []string

  hasCreatedCratesYet := false
  for scanner.Scan() {
    if scanner.Text() == "" {
      cratesStacks = ParseInputFileStackPartIntoCrates(cratesText)
      hasCreatedCratesYet = true
    } else if !hasCreatedCratesYet {
      cratesText = append(cratesText, scanner.Text())
    } else {
      var log Log
      fmt.Sscanf(scanner.Text(), "move %d from %d to %d", &log.HowMany, &log.FromStack, &log.ToStack)
      toRemoveStack := cratesStacks.Stacks[log.FromStack]
      toAddStack := cratesStacks.Stacks[log.ToStack]
      if toRemoveStack == nil || toAddStack == nil {
        return nil, errors.New("can't find pop or add stack")
      }

      startingIndexRemoved := len(toRemoveStack) - log.HowMany
      for i := startingIndexRemoved; i < len(toRemoveStack); i++ {
        toAddStack = append(toAddStack, toRemoveStack[i])
      }
      toRemoveStack = toRemoveStack[0:startingIndexRemoved]
      cratesStacks.Stacks[log.FromStack] = toRemoveStack
      cratesStacks.Stacks[log.ToStack] = toAddStack
    }
  }
  return cratesStacks, nil
}

func GetToppersFromFile(pathFile string) string {
  crates, err := ParseInputAndProcessLogs(pathFile)
  if err != nil {
    panic(err)
  }
  toppers := []rune{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '}
  for i, stack := range crates.Stacks {
    newTopper, _ := stack.Peek()
    toppers[i] = (newTopper)
  }
  return string(toppers)
}

func main() {
  fmt.Println(GetToppersFromFile("test.txt"))
  fmt.Println(GetToppersFromFile("input.txt"))
}
