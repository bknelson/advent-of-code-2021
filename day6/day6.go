package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Initial counter for a new fish
const newFishCounter int = 8

// Counter for an old fish after it get to 0
const resetCounter int = 6

// Initial state count
const initialStateCount int = 0

type fishState [newFishCounter + 1]*int

type fishPopulation struct {
	gen   int
	state fishState
}

func (f *fishPopulation) printStates() {
	for i, p := range (*f).state {
		fmt.Printf("%d\t%d\n", i, *p)
	}
	return
}

func (f *fishPopulation) generation() {
	first := (*f).state[0]

	for i := 1; i < len((*f).state); i++ {
		(*f).state[i-1] = (*f).state[i]
	}

	(*f).state[len((*f).state)-1] = first

	*(*f).state[resetCounter] += *(*f).state[len((*f).state)-1]

	(*f).gen++

	return
}

func (f *fishPopulation) count() int {
	sum := 0

	for _, p := range (*f).state {
		sum += *p
	}

	return sum
}

func fishPopulationFromCounters(counters string) (*fishPopulation, error) {
	fp := fishPopulation{0, emptyState()}

	counterArray := strings.Split(counters, ",")

	for _, val := range counterArray {
		valInt, err := strconv.Atoi(val)

		if err != nil {
			return nil, err
		}

		*(fp.state[valInt])++
	}

	return &fp, nil
}

func emptyState() fishState {
	var empty fishState

	for i := 0; i < len(empty); i++ {
		var val int = initialStateCount
		empty[i] = &val
	}

	return empty
}

func main() {
	totalGenerations, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	initialState := os.Args[2]
	newFish, err := fishPopulationFromCounters(initialState)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	for i := 0; i < totalGenerations; i++ {
		newFish.generation()
	}

	fmt.Println(newFish.count())
}
