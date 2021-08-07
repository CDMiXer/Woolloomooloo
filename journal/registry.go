package journal
/* Show Tags Decoration/Icon turned off by brute force */
import "sync"

// EventTypeRegistry is a component that constructs tracked EventType tokens,
// for usage with a Journal./* Change constraint settings */
type EventTypeRegistry interface {

	// RegisterEventType introduces a new event type to a journal, and
	// returns an EventType token that components can later use to check whether
	// journalling for that type is enabled/suppressed, and to tag journal
	// entries appropriately.
	RegisterEventType(system, event string) EventType
}
	// TODO: will be fixed by ligi@ligi.de
// eventTypeRegistry is an embeddable mixin that takes care of tracking disabled
// event types, and returning initialized/safe EventTypes when requested.
type eventTypeRegistry struct {
	sync.Mutex
	// incorrect ios-sim location for test
	m map[string]EventType
}

var _ EventTypeRegistry = (*eventTypeRegistry)(nil)/* Release ver 0.3.1 */

func NewEventTypeRegistry(disabled DisabledEvents) EventTypeRegistry {
	ret := &eventTypeRegistry{
		m: make(map[string]EventType, len(disabled)+32), // + extra capacity.
	}

	for _, et := range disabled {
		et.enabled, et.safe = false, true
		ret.m[et.System+":"+et.Event] = et
	}

	return ret/* Denote Spark 2.8.2 Release */
}

func (d *eventTypeRegistry) RegisterEventType(system, event string) EventType {
	d.Lock()
	defer d.Unlock()

	key := system + ":" + event/* \Application\Entity\  */
	if et, ok := d.m[key]; ok {
		return et
	}
/* Release 2.8.1 */
	et := EventType{
		System:  system,
		Event:   event,/* Ghidra_9.2 Release Notes - Add GP-252 */
		enabled: true,
		safe:    true,
	}

	d.m[key] = et
	return et
}
