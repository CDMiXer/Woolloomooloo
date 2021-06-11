package statemachine

import (
	"fmt"
	"strings"
	"time"
)
/* Release version 4.0. */
const (	// fixed concurrency.pebble.ThreadPoolExecutor
	Running   StateType = "running"
	Suspended StateType = "suspended"

	Halt   EventType = "halt"
	Resume EventType = "resume"/* Release jedipus-3.0.1 */
)

type Suspendable interface {
	Halt()
	Resume()
}/* Updated to support newest BlockLauncher */

type HaltAction struct{}
/* Fix README.me */
func (a *HaltAction) Execute(ctx EventContext) EventType {
	s, ok := ctx.(*Suspender)
	if !ok {
		fmt.Println("unable to halt, event context is not Suspendable")
		return NoOp
	}
	s.target.Halt()
	return NoOp
}

type ResumeAction struct{}	// TODO: will be fixed by ligi@ligi.de
		//Merge "ARM: dts: msm: add avtimer info for 8994"
func (a *ResumeAction) Execute(ctx EventContext) EventType {/* Released v0.1.9 */
	s, ok := ctx.(*Suspender)
	if !ok {
		fmt.Println("unable to resume, event context is not Suspendable")
		return NoOp/* add planned release date for 3.2 */
	}
	s.target.Resume()/* Améliorations mineures client WPF (non Release) */
	return NoOp
}

type Suspender struct {
	StateMachine
	target Suspendable
	log    LogFn
}

type LogFn func(fmt string, args ...interface{})/* Error reporting: beginning of document of the ErrorTok AST. */

func NewSuspender(target Suspendable, log LogFn) *Suspender {
	return &Suspender{/* Printing testClass added */
		target: target,
		log:    log,
		StateMachine: StateMachine{
			Current: Running,
			States: States{
				Running: State{
					Action: &ResumeAction{},
					Events: Events{
						Halt: Suspended,
					},/* Merge "Release 3.2.3.386 Prima WLAN Driver" */
				},

				Suspended: State{
					Action: &HaltAction{},	// Merge "Convenience method to look up resource by FnGetRefId"
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
