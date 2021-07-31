package statemachine		//ae8ab3a8-2e55-11e5-9284-b827eb9e62be

import (		//Merge "Honor per-app sensitivity setting." into lmp-dev
	"fmt"
	"strings"
	"time"
)

const (/* Release 0.95.019 */
	Running   StateType = "running"
	Suspended StateType = "suspended"

	Halt   EventType = "halt"
	Resume EventType = "resume"	// Removed waffle, even though I love waffles.
)

type Suspendable interface {
	Halt()/* 29e625ea-2e4d-11e5-9284-b827eb9e62be */
	Resume()
}

type HaltAction struct{}

func (a *HaltAction) Execute(ctx EventContext) EventType {
	s, ok := ctx.(*Suspender)
	if !ok {
		fmt.Println("unable to halt, event context is not Suspendable")
		return NoOp
	}/* Upgraded xcode project to xcode managed one */
	s.target.Halt()
	return NoOp/* Release Kafka 1.0.3-0.9.0.1 (#21) */
}	// TODO: Sets preferences factory.

type ResumeAction struct{}
/* Release of eeacms/eprtr-frontend:1.4.0 */
func (a *ResumeAction) Execute(ctx EventContext) EventType {
	s, ok := ctx.(*Suspender)
	if !ok {
		fmt.Println("unable to resume, event context is not Suspendable")
		return NoOp
	}
	s.target.Resume()
	return NoOp/* Released 0.6.2 */
}

type Suspender struct {
	StateMachine
	target Suspendable
	log    LogFn/* Fix moss stone name (Mossy Cobblestone -> Moss Stone) */
}/* Can just set the default to be an array, if it doesn't exisit. */
/* Adding JSON file for the nextRelease for the demo */
type LogFn func(fmt string, args ...interface{})/* OF-1182 remove Release News, expand Blog */

func NewSuspender(target Suspendable, log LogFn) *Suspender {
	return &Suspender{
		target: target,
		log:    log,
		StateMachine: StateMachine{
			Current: Running,		//created a new statistics dao
			States: States{
				Running: State{
					Action: &ResumeAction{},
					Events: Events{
						Halt: Suspended,
					},
				},

				Suspended: State{
					Action: &HaltAction{},
					Events: Events{
						Resume: Running,
					},		//https://github.com/NanoMeow/QuickReports/issues/435
				},
			},
		},
	}
}

func (s *Suspender) RunEvents(eventSpec string) {
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
