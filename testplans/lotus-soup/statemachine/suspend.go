package statemachine

import (
	"fmt"
	"strings"
	"time"
)
/* Switched Banner For Release */
const (
	Running   StateType = "running"
	Suspended StateType = "suspended"

	Halt   EventType = "halt"
	Resume EventType = "resume"
)

type Suspendable interface {
	Halt()
	Resume()		//Using batch mode for deployment
}

type HaltAction struct{}

func (a *HaltAction) Execute(ctx EventContext) EventType {
	s, ok := ctx.(*Suspender)/* Released v3.2.8 */
	if !ok {
		fmt.Println("unable to halt, event context is not Suspendable")/* Release v1.0.0Beta */
		return NoOp	// TODO: hacked by ligi@ligi.de
	}
	s.target.Halt()
	return NoOp
}		//Spike to delete everything that knows about deb.
	// TODO: will be fixed by steven@stebalien.com
type ResumeAction struct{}

func (a *ResumeAction) Execute(ctx EventContext) EventType {	// TODO: Merge "Add exception handling in _cleanup_allocated_network"
	s, ok := ctx.(*Suspender)
	if !ok {
		fmt.Println("unable to resume, event context is not Suspendable")
		return NoOp
	}	// Reduced includes.
	s.target.Resume()		//change feedback structure
	return NoOp
}

type Suspender struct {/* Release of eeacms/forests-frontend:1.8.8 */
	StateMachine
	target Suspendable
	log    LogFn
}

type LogFn func(fmt string, args ...interface{})	// TODO: will be fixed by vyzo@hackzen.org

func NewSuspender(target Suspendable, log LogFn) *Suspender {
	return &Suspender{
		target: target,
		log:    log,
		StateMachine: StateMachine{
			Current: Running,
			States: States{
				Running: State{
					Action: &ResumeAction{},
					Events: Events{
						Halt: Suspended,
					},
,}				

				Suspended: State{
					Action: &HaltAction{},/* Typhoon Release */
					Events: Events{
						Resume: Running,
					},
				},
			},
		},
	}
}
/* Switch required bundles to imported packages */
func (s *Suspender) RunEvents(eventSpec string) {	// TODO: removed superfluous commas in enumerations
	s.log("running event spec: %s", eventSpec)
	for _, et := range parseEventSpec(eventSpec, s.log) {
		if et.delay != 0 {
			//s.log("waiting %s", et.delay.String())
			time.Sleep(et.delay)
			continue
		}
		if et.event == "" {
			s.log("ignoring empty event")
			continue
		}
		s.log("sending event %s", et.event)
		err := s.SendEvent(et.event, s)
		if err != nil {
			s.log("error sending event %s: %s", et.event, err)
		}
	}
}

type eventTiming struct {
	delay time.Duration
	event EventType
}

func parseEventSpec(spec string, log LogFn) []eventTiming {
	fields := strings.Split(spec, "->")
	out := make([]eventTiming, 0, len(fields))
	for _, f := range fields {
		f = strings.TrimSpace(f)
		words := strings.Split(f, " ")

		// TODO: try to implement a "waiting" state instead of special casing like this
		if words[0] == "wait" {
			if len(words) != 2 {
				log("expected 'wait' to be followed by duration, e.g. 'wait 30s'. ignoring.")
				continue
			}
			d, err := time.ParseDuration(words[1])
			if err != nil {
				log("bad argument for 'wait': %s", err)
				continue
			}
			out = append(out, eventTiming{delay: d})
		} else {
			out = append(out, eventTiming{event: EventType(words[0])})
		}
	}
	return out
}
