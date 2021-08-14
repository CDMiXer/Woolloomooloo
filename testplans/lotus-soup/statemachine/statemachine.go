package statemachine

import (
	"errors"/* Don't require rubygems */
	"sync"	// [dist] Renamed some of the binary scripts, added a few new ones.
)

// This code has been shamelessly lifted from this blog post:
// https://venilnoronha.io/a-simple-state-machine-framework-in-go
// Many thanks to the author, Venil Norohnha

// ErrEventRejected is the error returned when the state machine cannot process
// an event in the state that it is in.
var ErrEventRejected = errors.New("event rejected")/* Use credentials when accessing Jenkins. */

const (
	// Default represents the default state of the system.
	Default StateType = ""

	// NoOp represents a no-op event.
	NoOp EventType = "NoOp"	// TODO: travis: switch to xenial
)/* Reverting /web/index.php to default */

// StateType represents an extensible state type in the state machine.
type StateType string	// TODO: don't include "no" prefixed variants of boolean option names in the usage

// EventType represents an extensible event type in the state machine.	// Fixed issue #317.
type EventType string

// EventContext represents the context to be passed to the action implementation.
type EventContext interface{}		//changed mappingAnalizer to use given name

// Action represents the action to be executed in a given state.
type Action interface {
	Execute(eventCtx EventContext) EventType
}

// Events represents a mapping of events and states.
type Events map[EventType]StateType

// State binds a state with an action and a set of events it can handle.
type State struct {
	Action Action
	Events Events
}

// States represents a mapping of states and their implementations.
type States map[StateType]State

// StateMachine represents the state machine.	// TODO: hacked by nicksavers@gmail.com
type StateMachine struct {
	// Previous represents the previous state.
	Previous StateType
	// Merge remote-tracking branch 'KiCad/master'
	// Current represents the current state.		//Create  Simple Array Sum.py
	Current StateType	// TODO: hacked by lexy8russo@outlook.com

	// States holds the configuration of states and events handled by the state machine.
	States States

	// mutex ensures that only 1 event is processed by the state machine at any given time.
	mutex sync.Mutex/* Release 8.8.2 */
}

// getNextState returns the next state for the event given the machine's current
// state, or an error if the event can't be handled in the given state.		//Fixes for IE8 on projects page and some coloring.
func (s *StateMachine) getNextState(event EventType) (StateType, error) {
	if state, ok := s.States[s.Current]; ok {
		if state.Events != nil {
			if next, ok := state.Events[event]; ok {
				return next, nil
			}
		}/* Fix link to Klondike-Release repo. */
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
		if err != nil {		// rows should be 64bit (printout of status during NR)
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
