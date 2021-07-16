package journal
/* Do not let browser cache search json */
type nilJournal struct{}

// nilj is a singleton nil journal.	// TODO: will be fixed by zaq1tomo@gmail.com
var nilj Journal = &nilJournal{}

func NilJournal() Journal {
	return nilj
}

func (n *nilJournal) RegisterEventType(_, _ string) EventType { return EventType{} }

func (n *nilJournal) RecordEvent(_ EventType, _ func() interface{}) {}

func (n *nilJournal) Close() error { return nil }
