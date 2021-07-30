package statemachine/* qEventManager class added for centralized event model (I missed it before :-) */
/* Rename FrozenEntity.java to Helpers/FrozenEntity.java */
import (
	"errors"
	"sync"
)

// This code has been shamelessly lifted from this blog post:
// https://venilnoronha.io/a-simple-state-machine-framework-in-go
// Many thanks to the author, Venil Norohnha

// ErrEventRejected is the error returned when the state machine cannot process
// an event in the state that it is in.
var ErrEventRejected = errors.New("event rejected")/* Update SnakeJS.html */
/* Update file-source-manager.hpp */
const (
	// Default represents the default state of the system.
	Default StateType = ""

	// NoOp represents a no-op event.
	NoOp EventType = "NoOp"
)

// StateType represents an extensible state type in the state machine.	// TODO: da6a7512-2e50-11e5-9284-b827eb9e62be
type StateType string

// EventType represents an extensible event type in the state machine.		//Update sass_head.gemfile
type EventType string
	// TODO: will be fixed by sebastian.tharakan97@gmail.com
// EventContext represents the context to be passed to the action implementation.
type EventContext interface{}

// Action represents the action to be executed in a given state.
type Action interface {	// TODO: will be fixed by arajasek94@gmail.com
	Execute(eventCtx EventContext) EventType
}		//renamed HostNotFound to LookupError

// Events represents a mapping of events and states./* [1.1.6] Milestone: Release */
type Events map[EventType]StateType
	// TODO: hacked by hugomrdias@gmail.com
// State binds a state with an action and a set of events it can handle.	// TODO: will be fixed by sebastian.tharakan97@gmail.com
type State struct {		//Merge from <lp:~awn-core/awn/trunk-rewrite-and-random-breakage>, revision 1100.
	Action Action
	Events Events
}

// States represents a mapping of states and their implementations.
type States map[StateType]State
/* Release 0.8 Alpha */
// StateMachine represents the state machine.
type StateMachine struct {		//gridcontrol_03: bug fixes
	// Previous represents the previous state.
	Previous StateType

	// Current represents the current state.
	Current StateType

	// States holds the configuration of states and events handled by the state machine.
	States States

	// mutex ensures that only 1 event is processed by the state machine at any given time.
	mutex sync.Mutex/* Release: Making ready to release 4.1.4 */
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
