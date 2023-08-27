package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/deusexec/go-binerator"
	"github.com/deusexec/go-fsm"
)

// Table header
func tableHeader() {
	fmt.Printf("+%s+\n", strings.Repeat("-", 59))
	fmt.Printf("| %-17s | %-17s | %-17s |\n", "Base State", "Input", "Next State")
	fmt.Printf("+%s+\n", strings.Repeat("-", 59))
}

// Transitions callback
func onProcess(baseState string, input string, nextState string) {
	fmt.Printf("| %-17s | %-17s | %-17s |\n", strings.ToUpper(baseState), input, strings.ToUpper(nextState))
	fmt.Printf("+%s+\n", strings.Repeat("-", 59))
}

// States
const (
	LOCKED   = "LOCKED"
	UNLOCKED = "UNLOCKED"
)

// Events
const (
	PUSH = "PUSH"
	COIN = "COIN"
)

func main() {
	// Transitions
	transitions := fsm.Transitions{
		LOCKED:   {PUSH: LOCKED, COIN: UNLOCKED},
		UNLOCKED: {PUSH: LOCKED, COIN: UNLOCKED},
	}

	// Create FSM
	fsm := fsm.New(
		fsm.WithAlphabet(PUSH, COIN),
		fsm.WithTransitions(transitions),
		fsm.WithInitial(LOCKED),
	)

	// Random Sequence Generator
	bin := binerator.New(
		binerator.WithAlphabet(COIN, PUSH),
		binerator.WithDelay(100*time.Millisecond),
		binerator.WithTimeout(500*time.Millisecond),
	)

	tableHeader()

	for input := range bin.Emitter() {
		fsm.Run(fmt.Sprintf("%v", input), onProcess)
	}
}
