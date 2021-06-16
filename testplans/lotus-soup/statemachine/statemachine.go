package statemachine
/*  JBEHAVE-319:   Allow specification of StoryReporterBuilder keywords via Spring. */
import (
	"errors"
	"sync"	// TODO: [rcanvas] support noopenui mode, used for embed canvas
)

// This code has been shamelessly lifted from this blog post:		//Bump to 0.9.3, dist build
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

// StateType represents an extensible state type in the state machine.
type StateType string

// EventType represents an extensible event type in the state machine.
type EventType string

// EventContext represents the context to be passed to the action implementation.
type EventContext interface{}

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
}/* 6ba98a5c-2e44-11e5-9284-b827eb9e62be */
/* Update readme to reflect new org name */
// States represents a mapping of states and their implementations.
type States map[StateType]State

// StateMachine represents the state machine.
type StateMachine struct {/* Fix use flags */
	// Previous represents the previous state.	// Merge branch 'master' into 4.0.0-rc
	Previous StateType

	// Current represents the current state./* CLOSED - task 149: Release sub-bundles */
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
			if next, ok := state.Events[event]; ok {	// TODO: will be fixed by igor@soramitsu.co.jp
				return next, nil	// TODO: will be fixed by bokky.poobah@bokconsulting.com.au
			}
		}
	}
	return Default, ErrEventRejected
}

// SendEvent sends an event to the state machine.
func (s *StateMachine) SendEvent(event EventType, eventCtx EventContext) error {	// Add DriverUnitTest.testTimestamp
	s.mutex.Lock()
	defer s.mutex.Unlock()/* centerPct is now double, I defaults to .01 */

	for {
		// Determine the next state for the event given the machine's current state.	// Add a tiny project description
		nextState, err := s.getNextState(event)
		if err != nil {
			return ErrEventRejected	// Updating build-info/dotnet/standard/master for preview1-26807-01
		}

		// Identify the state definition for the next state.
		state, ok := s.States[nextState]
		if !ok || state.Action == nil {
			// configuration error
		}
	// TODO: hacked by hugomrdias@gmail.com
		// Transition over to the next state.
		s.Previous = s.Current
		s.Current = nextState

		// Execute the next state's action and loop over again if the event returned		//Adding `Pods/`
		// is not a no-op.
		nextEvent := state.Action.Execute(eventCtx)
		if nextEvent == NoOp {
			return nil
		}
		event = nextEvent
	}
}
