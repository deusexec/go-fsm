package fsm

import (
	"testing"
)

func Test_Alphabet(t *testing.T) {
	const (
		COIN = "coin"
		PUSH = "push"
	)
	alphabet := []any{COIN, PUSH}
	automaton := New(
		WithAlphabet(alphabet...),
	)
	if len(automaton.alphabet) != len(alphabet) {
		t.Errorf("Alphabet length is wrong. Expected: %v (%v), Received: %v (%v)\n", len(automaton.alphabet), automaton.alphabet, len(alphabet), alphabet)
	}
	for _, letter := range alphabet {
		if l, ok := automaton.alphabet[letter]; !ok {
			t.Errorf("Alphabet letter is not present. For: %v, Expected: true, Received: %v\n", letter, l)
		}
	}
}

func Test_States(t *testing.T) {

}

func Test_InitialState(t *testing.T) {

}

func Test_AcceptableState(t *testing.T) {

}

func Test_RejectableState(t *testing.T) {

}

func Test_Transitions(t *testing.T) {

}
