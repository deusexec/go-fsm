package fsm

// todo: change names to the lower-case
type Alphabet map[any]bool
type States map[string]*State

type TransitionsKeys map[string]string
type Transitions map[string]TransitionsKeys

type State struct {
	name       string
	isAccepted bool
	isRejected bool
}

type FiniteStateMachine struct {
	alphabet    Alphabet
	transitions Transitions
	states      States
	activeState *State
}

// Returns new instance of a fsm
func New(options ...func(*FiniteStateMachine)) *FiniteStateMachine {
	fsm := new(FiniteStateMachine)
	for _, option := range options {
		option(fsm)
	}
	return fsm
}

// Set alphabet for the fsm
func WithAlphabet(alphabet ...any) func(*FiniteStateMachine) {
	return func(fsm *FiniteStateMachine) {
		length := len(alphabet)
		if length > 0 {
			fsm.alphabet = make(Alphabet, length)
			for _, letter := range alphabet {
				fsm.alphabet[letter] = true
			}
		}
	}
}

// Set transitions
func WithTransitions(transitions Transitions) func(*FiniteStateMachine) {
	return func(fsm *FiniteStateMachine) {
		fsm.transitions = transitions
		fsm.states = make(States, len(fsm.transitions))
		for state := range fsm.transitions {
			fsm.states[state] = &State{name: state}
		}
	}
}

// Set initial state
func WithInitial(state string) func(*FiniteStateMachine) {
	return func(fsm *FiniteStateMachine) {
		if activeState, ok := fsm.states[state]; ok {
			fsm.activeState = activeState
		}
	}
}

// Set acceptable state
func WithAcceptable(state string) func(*FiniteStateMachine) {
	return func(fsm *FiniteStateMachine) {
		if state, ok := fsm.states[state]; ok {
			state.isAccepted = true
		}
	}
}

// Set rejectable state
func WithRejectable(state string) func(*FiniteStateMachine) {
	return func(fsm *FiniteStateMachine) {
		if state, ok := fsm.states[state]; ok {
			state.isRejected = true
		}
	}
}

// Check if fsm is done
func (fsm *FiniteStateMachine) IsDone() bool {
	return fsm.activeState.isAccepted || fsm.activeState.isRejected
}

// Check input for valid values
func (fsm *FiniteStateMachine) IsValid(input any) bool {
	if _, ok := fsm.alphabet[input]; ok {
		return true
	}
	return false
}

// Run fsm
func (fsm *FiniteStateMachine) Run(input string, onProcess func(base string, input string, next string)) {
	if fsm.IsDone() {
		return
	}
	nextState := fsm.transitions[fsm.activeState.name][input]
	onProcess(fsm.activeState.name, input, nextState)
	fsm.activeState = fsm.states[nextState]
}
