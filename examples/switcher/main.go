package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/deusexec/go-binerator"
	"github.com/deusexec/go-fsm"
)

// Table Header
func tableHeader() {
	fmt.Printf("+%s+\n", strings.Repeat("-", 59))
	fmt.Printf("| %-17s | %-17s | %-17s |\n", "Base State", "Input", "Next State")
	fmt.Printf("+%s+\n", strings.Repeat("-", 59))
}

// Transition Callback
func onTransition(baseState string, input string, nextState string) {
	fmt.Printf("| %-17s | %-17s | %-17s |\n", strings.ToUpper(baseState), input, strings.ToUpper(nextState))
	fmt.Printf("+%s+\n", strings.Repeat("-", 59))
}

// States
const (
	ON  = "ON"
	OFF = "OFF"
)

// Events
const (
	PUSH = "PUSH"
)

func main() {
	// Transitions
	transitions := fsm.Transitions{
		ON:  {PUSH: OFF},
		OFF: {PUSH: ON},
	}

	// Finite State Machine
	fsm := fsm.New(
		fsm.WithAlphabet(PUSH),
		fsm.WithTransitions(transitions),
		fsm.WithInitial(OFF),
	)

	// Random Sequence Generator
	bin := binerator.New(
		binerator.WithAlphabet(PUSH),                // Events: {PUSH}
		binerator.WithDelay(100*time.Millisecond),   // Event Delay: 100 milliseconds
		binerator.WithTimeout(500*time.Millisecond), // Timeout: 500 milliseconds
	)

	tableHeader()

	// Read input from a binerator and run fsm processing
	for input := range bin.Emitter() {
		input := fmt.Sprintf("%v", input)
		fsm.Run(input, onTransition)
	}
}
