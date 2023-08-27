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
	WINDOW1  = "WINDOW1"
	WINDOW2  = "WINDOW2"
	WINDOW3  = "WINDOW3"
	WINDOW4  = "WINDOW4"
	FINISH   = "FINISH"
	CANCELED = "CANCELED"
)

// Events
const (
	NEXT   = "NEXT"
	CANCEL = "CANCEL"
)

func main() {
	// Transitions
	transitions := fsm.Transitions{
		WINDOW1:  {NEXT: WINDOW2, CANCEL: CANCELED},
		WINDOW2:  {NEXT: WINDOW3, CANCEL: WINDOW1},
		WINDOW3:  {NEXT: WINDOW4, CANCEL: WINDOW1},
		WINDOW4:  {NEXT: FINISH, CANCEL: WINDOW1},
		FINISH:   {},
		CANCELED: {},
	}

	// Finite State Machine
	fsm := fsm.New(
		fsm.WithAlphabet(NEXT, CANCEL),
		fsm.WithTransitions(transitions),
		fsm.WithInitial(WINDOW1),
		fsm.WithAcceptable(FINISH),
		fsm.WithRejectable(CANCELED),
	)

	// Random Sequence Generator
	bin := binerator.New(
		binerator.WithAlphabet(NEXT, CANCEL),      // Events: {NEXT, CANCEL}
		binerator.WithDelay(100*time.Millisecond), // Event Delay: 100 milliseconds
	)

	tableHeader()

	// Read input from a binerator and run fsm processing
	for input := range bin.Emitter() {
		// Check if the fsm is finished
		if fsm.IsDone() {
			// Send `done` signal to the data emitter
			bin.Done()
		}
		input := fmt.Sprintf("%v", input)
		fsm.Run(input, onTransition)
	}
}
