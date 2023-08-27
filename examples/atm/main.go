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
	READY            = "READY"
	PIN              = "PIN"
	MENU             = "MENU"
	DEPOSIT_ACCOUNT  = "DEPOSIT_ACCOUNT"
	DEPOSIT_AMOUNT   = "DEPOSIT_AMOUNT"
	DEPOSIT_CONFIRM  = "DEPOSIT_CONFIRM"
	CASH_COLLECT     = "CASH_COLLECT"
	WITHDRAW_ACCOUNT = "WITHDRAW_ACCOUNT"
	WITHDRAW_AMOUNT  = "WITHDRAW_AMOUNT"
	WITHDRAW_CONFIRM = "WITHDRAW_CONFIRM"
	CASH_DISPENSE    = "CASH_DISPENSE"
	CONTINUE         = "CONTINUE"
	RETURN           = "RETURN"
)

// Events
const (
	CARD     = "CARD"
	DEPOSIT  = "DEPOSIT"
	WITHDRAW = "WITHDRAW"
	PROVIDE  = "PROVIDE"
	CONFIRM  = "CONFIRM"
	CANCEL   = "CANCEL"
	REJECT   = "REJECT"
)

func main() {
	// Transitions
	transitions := fsm.Transitions{
		READY:  {CARD: PIN},
		PIN:    {CONFIRM: MENU, REJECT: RETURN},
		REJECT: {},
	}

	// Finite State Machine
	fsm := fsm.New(
		fsm.WithAlphabet(CARD, DEPOSIT, WITHDRAW, PROVIDE, CONFIRM, CANCEL, REJECT),
		fsm.WithTransitions(transitions),
		fsm.WithInitial(READY),
		fsm.WithAcceptable(RETURN),
		fsm.WithRejectable(RETURN),
	)

	// Random Sequence Generator
	bin := binerator.New(
		binerator.WithAlphabet(CARD, DEPOSIT, WITHDRAW, PROVIDE, CONFIRM, CANCEL, REJECT),
		binerator.WithDelay(100*time.Millisecond),
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
