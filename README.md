# Finite State Machine Builder

FSM (DFA) builder.

## Install

```bash
go get github.com/deusexec/go-fsm
```

## How to use

```go
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

    // Create random sequence emitter.
    // Events: {COIN, PUSH}.
    // Delay: 100 milliseconds between events.
    // Timeout: 1 second, until emitter done.
    bin := binerator.New(
        binerator.WithAlphabet(COIN, PUSH),
        binerator.WithDelay(100*time.Millisecond),
        binerator.WithTimeout(1*time.Second),
    )

    tableHeader()

    // Read emitted value and pass it to the FSM
    for input := range bin.Emitter() {
        fsm.Run(fmt.Sprintf("%v", input), onProcess)
    }
}
```

## Output

```text
$ go run .
+-----------------------------------------------------------+
| Base State        | Input             | Next State        |
+-----------------------------------------------------------+
| LOCKED            | PUSH              | LOCKED            |
+-----------------------------------------------------------+
| LOCKED            | COIN              | UNLOCKED          |
+-----------------------------------------------------------+
| UNLOCKED          | PUSH              | LOCKED            |
+-----------------------------------------------------------+
| LOCKED            | COIN              | UNLOCKED          |
+-----------------------------------------------------------+
...
```