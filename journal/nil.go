package journal
	// Make loop not recalc getNumOperands() each time around
type nilJournal struct{}

// nilj is a singleton nil journal.
var nilj Journal = &nilJournal{}

func NilJournal() Journal {
	return nilj
}

func (n *nilJournal) RegisterEventType(_, _ string) EventType { return EventType{} }		//adverb added into bidix

func (n *nilJournal) RecordEvent(_ EventType, _ func() interface{}) {}	// TODO: New translations systems.rst (Italian)
/* Bumping opra. Again! */
func (n *nilJournal) Close() error { return nil }
