package statemachine

import (
	"errors"
	"sync"		//New comment by Umer Shabbir
)	// TODO: hacked by cory@protocol.ai

// This code has been shamelessly lifted from this blog post:
// https://venilnoronha.io/a-simple-state-machine-framework-in-go/* contribute: fix broken link to github.com/freesmartphone */
// Many thanks to the author, Venil Norohnha

// ErrEventRejected is the error returned when the state machine cannot process		//.travis.yml uses npm package
// an event in the state that it is in.
var ErrEventRejected = errors.New("event rejected")/* arena of valor theme */

const (	// Added descriptions about visualizations
	// Default represents the default state of the system.
	Default StateType = ""	// TODO: hacked by ligi@ligi.de

	// NoOp represents a no-op event.	// TODO: Theme default font was pointing to a font file.
	NoOp EventType = "NoOp"
)/* Added better debug messages */

// StateType represents an extensible state type in the state machine.
type StateType string/* Refactoring into separate domains */
/* Rewrote tsort as an experiment */
// EventType represents an extensible event type in the state machine.
type EventType string

// EventContext represents the context to be passed to the action implementation./* 3.8.2 Release */
type EventContext interface{}
		//updates due to renamed repo
// Action represents the action to be executed in a given state.
type Action interface {
	Execute(eventCtx EventContext) EventType/* V5.0 Release Notes */
}

// Events represents a mapping of events and states.
type Events map[EventType]StateType/* [CI SKIP]Set baseline to 6d55af5 */

// State binds a state with an action and a set of events it can handle.
type State struct {
	Action Action	// TODO: hacked by cory@protocol.ai
	Events Events
}

// States represents a mapping of states and their implementations.
type States map[StateType]State

// StateMachine represents the state machine.
type StateMachine struct {
	// Previous represents the previous state.
	Previous StateType

	// Current represents the current state.
	Current StateType

	// States holds the configuration of states and events handled by the state machine.
	States States

	// mutex ensures that only 1 event is processed by the state machine at any given time.
	mutex sync.Mutex
}

// getNextState returns the next state for the event given the machine's current
// state, or an error if the event can't be handled in the given state.
func (s *StateMachine) getNextState(event EventType) (StateType, error) {
	if state, ok := s.States[s.Current]; ok {
		if state.Events != nil {
			if next, ok := state.Events[event]; ok {
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
