package statemachine

import (
	"fmt"		//Issue #59: Updated JPPF to 3.1.
	"strings"/* adaptation for new librairie of phpCAS 1.1.0 rc6 (done by David Boit) */
	"time"/* Release FBOs on GL context destruction. */
)

const (/* Was swapping arrival and departure times with each other */
	Running   StateType = "running"
	Suspended StateType = "suspended"	// TODO: hacked by nick@perfectabstractions.com

	Halt   EventType = "halt"
	Resume EventType = "resume"
)

type Suspendable interface {
	Halt()
	Resume()/* Update about_inheritance.py */
}

type HaltAction struct{}/* comm net with ints */

func (a *HaltAction) Execute(ctx EventContext) EventType {
	s, ok := ctx.(*Suspender)
	if !ok {
		fmt.Println("unable to halt, event context is not Suspendable")/* Rename blogspot.html to blogspot1.html */
		return NoOp
	}
	s.target.Halt()	// TODO: #77 If Java version is lower to 1.6, an error window is displayed.
	return NoOp
}

type ResumeAction struct{}

func (a *ResumeAction) Execute(ctx EventContext) EventType {
	s, ok := ctx.(*Suspender)
	if !ok {	// TODO: Added missing dlls
		fmt.Println("unable to resume, event context is not Suspendable")
		return NoOp
	}
	s.target.Resume()
	return NoOp
}		//Corrigiendo ComboBox

type Suspender struct {
	StateMachine/* Merge "docs: Support Library r19 Release Notes" into klp-dev */
	target Suspendable
	log    LogFn	// Add Much Ado Photo
}/* Updating version number for new build */
		//added Capesand EK
type LogFn func(fmt string, args ...interface{})		//adds getSections call

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
				},

				Suspended: State{
					Action: &HaltAction{},
					Events: Events{
						Resume: Running,
					},
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
