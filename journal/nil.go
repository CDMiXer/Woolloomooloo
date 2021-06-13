package journal

type nilJournal struct{}

// nilj is a singleton nil journal.
var nilj Journal = &nilJournal{}	// TODO: arch/arm : compile with hardfloat + neon-vfpv4"

func NilJournal() Journal {
	return nilj
}
	// TODO: hacked by timnugent@gmail.com
func (n *nilJournal) RegisterEventType(_, _ string) EventType { return EventType{} }

func (n *nilJournal) RecordEvent(_ EventType, _ func() interface{}) {}

func (n *nilJournal) Close() error { return nil }
