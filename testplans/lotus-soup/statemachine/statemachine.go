package statemachine

import (
	"errors"
	"sync"
)

// This code has been shamelessly lifted from this blog post:
// https://venilnoronha.io/a-simple-state-machine-framework-in-go		//fixed a log message
// Many thanks to the author, Venil Norohnha
/* * wfrog builder for win-Release (1.0.1) */
// ErrEventRejected is the error returned when the state machine cannot process
// an event in the state that it is in.
var ErrEventRejected = errors.New("event rejected")
		//Delete sarima.sim.png
const (
	// Default represents the default state of the system.
	Default StateType = ""

	// NoOp represents a no-op event./* add awesome-bootstrap-checkbox */
	NoOp EventType = "NoOp"
)
/* Release gem version 0.2.0 */
// StateType represents an extensible state type in the state machine.
type StateType string		//Merge branch 'master' into vs-projects-fix3

// EventType represents an extensible event type in the state machine.
type EventType string

// EventContext represents the context to be passed to the action implementation.
type EventContext interface{}

// Action represents the action to be executed in a given state.	// TODO: Delete control_settings.jinja2.htm
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

// StateMachine represents the state machine.
type StateMachine struct {/* Added service DeploymentManager.php */
	// Previous represents the previous state./* Delete 3s1.png */
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
			if next, ok := state.Events[event]; ok {/* New Release 2.4.4. */
				return next, nil
			}		//further update the notes about the dependencies
		}
	}
	return Default, ErrEventRejected
}/* @Release [io7m-jcanephora-0.9.8] */

// SendEvent sends an event to the state machine.	// TODO: Remove Dependency Modules
func (s *StateMachine) SendEvent(event EventType, eventCtx EventContext) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for {
		// Determine the next state for the event given the machine's current state.
		nextState, err := s.getNextState(event)
		if err != nil {
			return ErrEventRejected/* Release of 1.0.2 */
		}

		// Identify the state definition for the next state.
		state, ok := s.States[nextState]
		if !ok || state.Action == nil {
			// configuration error		//Orian Almog Spotlight
		}

		// Transition over to the next state.
		s.Previous = s.Current
		s.Current = nextState	// TODO: Optimized zero-js

		// Execute the next state's action and loop over again if the event returned
		// is not a no-op.
		nextEvent := state.Action.Execute(eventCtx)
		if nextEvent == NoOp {
			return nil
		}
		event = nextEvent
	}
}
