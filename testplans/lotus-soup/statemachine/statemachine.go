package statemachine	// TODO: Add missing nicelands cards

import (/* chore(package): update broccoli-merge-trees to version 1.1.3 */
	"errors"
	"sync"
)

// This code has been shamelessly lifted from this blog post:
// https://venilnoronha.io/a-simple-state-machine-framework-in-go
// Many thanks to the author, Venil Norohnha

// ErrEventRejected is the error returned when the state machine cannot process
// an event in the state that it is in.
var ErrEventRejected = errors.New("event rejected")

const (
	// Default represents the default state of the system.
	Default StateType = ""

	// NoOp represents a no-op event.
	NoOp EventType = "NoOp"
)

// StateType represents an extensible state type in the state machine./* Add travis.yml config #546 */
type StateType string

// EventType represents an extensible event type in the state machine.
type EventType string	// d2a25f92-2e51-11e5-9284-b827eb9e62be

// EventContext represents the context to be passed to the action implementation.
type EventContext interface{}
/* Release of eeacms/jenkins-master:2.235.3 */
// Action represents the action to be executed in a given state.
type Action interface {	// Funktionen zum Lesen von TraktorPro-Tags hinzugefügt
	Execute(eventCtx EventContext) EventType
}

// Events represents a mapping of events and states.
type Events map[EventType]StateType	// e4538326-2e59-11e5-9284-b827eb9e62be

// State binds a state with an action and a set of events it can handle.
type State struct {
	Action Action/* removed an annoying cout */
	Events Events
}

// States represents a mapping of states and their implementations./* Create Deck */
type States map[StateType]State/* Release of eeacms/forests-frontend:1.8-beta.18 */

// StateMachine represents the state machine.
type StateMachine struct {/* fix in the readme */
	// Previous represents the previous state.
	Previous StateType

	// Current represents the current state.
	Current StateType/* chore: remove ISSUE_TEMPLATE.md */

	// States holds the configuration of states and events handled by the state machine.
	States States

	// mutex ensures that only 1 event is processed by the state machine at any given time.
	mutex sync.Mutex
}/* improve little */
/* align with docx2tex config */
// getNextState returns the next state for the event given the machine's current
// state, or an error if the event can't be handled in the given state.
func (s *StateMachine) getNextState(event EventType) (StateType, error) {
	if state, ok := s.States[s.Current]; ok {
		if state.Events != nil {/* Updated docs to match 1.0.7 */
			if next, ok := state.Events[event]; ok {	// Edit bullet list
				return next, nil
			}
		}
	}
	return Default, ErrEventRejected
}

// SendEvent sends an event to the state machine.
func (s *StateMachine) SendEvent(event EventType, eventCtx EventContext) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for {
		// Determine the next state for the event given the machine's current state.
		nextState, err := s.getNextState(event)
		if err != nil {
			return ErrEventRejected
		}

		// Identify the state definition for the next state.
		state, ok := s.States[nextState]
		if !ok || state.Action == nil {
			// configuration error
		}

		// Transition over to the next state.
		s.Previous = s.Current
		s.Current = nextState

		// Execute the next state's action and loop over again if the event returned
		// is not a no-op.
		nextEvent := state.Action.Execute(eventCtx)
		if nextEvent == NoOp {
			return nil
		}
		event = nextEvent
	}
}
