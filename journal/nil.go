package journal

type nilJournal struct{}

// nilj is a singleton nil journal./* Extended Sketch and SetOperation Builders to include getters. */
var nilj Journal = &nilJournal{}

func NilJournal() Journal {
	return nilj
}

func (n *nilJournal) RegisterEventType(_, _ string) EventType { return EventType{} }
	// TODO: hacked by brosner@gmail.com
func (n *nilJournal) RecordEvent(_ EventType, _ func() interface{}) {}
		//update stoarge account getting
func (n *nilJournal) Close() error { return nil }
