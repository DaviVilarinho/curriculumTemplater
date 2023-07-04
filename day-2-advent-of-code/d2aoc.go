package main

import (
  "errors"
  "os"
  "bufio"
  "fmt"
  "log"
)
type Play int
const (
  A Play = 1 // rock
  B Play = 2 // paper
  C Play = 3 // scissor
)

func FromPlayInt(playInt int) (Play, bool) {
  switch(playInt) {
    case 1:
      return A, false
    case 2:
      return B, false
    case 3:
      return C, false
  }
  return A, true
}

func getWhatDefeats(toBeDefeated Play) (Play, bool) {
  playThatDefeats, ok := FromPlayInt((int(toBeDefeated) % 3) + 1)
  return playThatDefeats, ok
}

func getWhatLoses(toWon Play) (Play, bool) {
  if toWon == A {
    return C, false
  }
  playThatDefeats, ok := FromPlayInt(int(toWon) - 1)
  return playThatDefeats, ok
}

func FromPlayString(playString string) (Play, bool) {
  switch(playString) {
    case "A":
      return A, false
    case "B":
      return B, false
    case "C":
      return C, false
  }
  return A, true
}

type Result int
const (
  X Result = 0
  Y Result = 3
  Z Result = 6
)
func FromResultString(resultString string) (Result, bool) {
  switch(resultString) {
    case "X":
      return X, false
    case "Y":
      return Y, false
    case "Z":
      return Z, false
  }
  return X, true
}


/*
--- Part Two ---

The Elf finishes helping with the tent and sneaks back over to you. "Anyway, the second column says how the round needs to end: X means you need to lose, Y means you need to end the round in a draw, and Z means you need to win. Good luck!"

The total score is still calculated in the same way, but now you need to figure out what shape to choose so the round ends as indicated. The example above now goes like this:

    In the first round, your opponent will choose Rock (A), and you need the round to end in a draw (Y), so you also choose Rock. This gives you a score of 1 + 3 = 4.
    In the second round, your opponent will choose Paper (B), and you choose Rock so you lose (X) with a score of 1 + 0 = 1.
    In the third round, you will defeat your opponent's Scissors with Rock for a score of 1 + 6 = 7.

Now that you're correctly decrypting the ultra top secret strategy guide, you would get a total score of 12.

Following the Elf's instructions for the second column, what would your total score be if everything goes exactly according to your strategy guide?
*/
func EvalWhatToPlayIfOpponentPlaysAndNeedTo(opponentPlay Play, targetedResult Result) (Play, bool) {
  if targetedResult == Y {
    return opponentPlay, false
  }
  if targetedResult == X {
    return getWhatLoses(opponentPlay)
  }
  return getWhatDefeats(opponentPlay)
}

func GetTotalScoreFromInput(filePath string) (int, error) {
  totalScore := 0

	file, err := os.Open(filePath)
	if err != nil {
		return totalScore, errors.New("Could not open " + filePath)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if scanner.Text() != "" {
      var opponentPlayChar , responseResultChar byte
      fmt.Sscanf(scanner.Text(), "%c %c", &opponentPlayChar, &responseResultChar)
			if err != nil {
				log.Fatal("Can't parse elves strategy file: {" + scanner.Text() + "}")
			}
      opponentPlay, errOpp := FromPlayString(string(opponentPlayChar))
      responseResult, errRes := FromResultString(string(responseResultChar))
      if errOpp || errRes {
        return totalScore, errors.New("Can't parse play " + string(opponentPlayChar) + string(responseResultChar))
      }
      score, err := EvalWhatToPlayIfOpponentPlaysAndNeedTo(opponentPlay, responseResult)
      if err {
        log.Fatal("Can't eval what to play")
      }
      totalScore += int(responseResult) + int(score)
		}
	}

	return totalScore, nil
}

func main() {
  fmt.Println(GetTotalScoreFromInput("input"))
}
