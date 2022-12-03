package main

import (
	"fmt"
)

type Play int

const (
	ROCK Play = iota
	PAPER
	SCISSORS
)

type Strategy1 struct {
	opponent Play
	you      Play
}

type Goal int

const (
	LOSS Goal = 0
	DRAW Goal = 3
	WIN  Goal = 6
)

func determineTotalScore1(strategyGuide []Strategy1) int {
	total := 0
	for _, strategy := range strategyGuide {
		score := int(strategy.you) + 1
		if strategy.opponent == ROCK {
			if strategy.you == ROCK {
				score += int(DRAW)
			} else if strategy.you == PAPER {
				score += int(WIN)
			} else if strategy.you == SCISSORS {
				score += int(LOSS)
			}
		}
		if strategy.opponent == PAPER {
			if strategy.you == ROCK {
				score += int(LOSS)
			} else if strategy.you == PAPER {
				score += int(DRAW)
			} else if strategy.you == SCISSORS {
				score += int(WIN)
			}
		}
		if strategy.opponent == SCISSORS {
			if strategy.you == ROCK {
				score += int(WIN)
			} else if strategy.you == PAPER {
				score += int(LOSS)
			} else if strategy.you == SCISSORS {
				score += int(DRAW)
			}
		}
		total += score
	}
	return total
}

type Strategy2 struct {
	opponent Play
	goal     Goal
}

func determineTotalScore2(strategyGuide []Strategy2) int {
	total := 0
	for _, strategy := range strategyGuide {
        score := int(strategy.goal)
        var you Play
        if strategy.goal == DRAW {
            you = strategy.opponent;
        }
        if strategy.goal == LOSS {
            if strategy.opponent == ROCK {
                you = SCISSORS
            } else
            if strategy.opponent == PAPER {
                you = ROCK
            } else
            if strategy.opponent == SCISSORS {
                you = PAPER
            }
        }
        if strategy.goal == WIN {
            if strategy.opponent == ROCK {
                you = PAPER
            } else
            if strategy.opponent == PAPER {
                you = SCISSORS
            } else
            if strategy.opponent == SCISSORS {
                you = ROCK
            }
        }
        score += int(you) + 1
		total += score
	}
	return total
}

func main() {
	strategyGuide1 := readStrategyGuide1("inputs/challenge.txt")
	fmt.Printf("Challenge #1 Total score: %d\n", determineTotalScore1(strategyGuide1))

	strategyGuide2 := readStrategyGuide2("inputs/challenge.txt")
	fmt.Printf("Challenge #2 Total score: %d\n", determineTotalScore2(strategyGuide2))
}
