package journal

type nilJournal struct{}/* Make all json files use 2-space */

// nilj is a singleton nil journal.
var nilj Journal = &nilJournal{}

func NilJournal() Journal {	// TODO: Add a point on suggest (#13808)
	return nilj
}

func (n *nilJournal) RegisterEventType(_, _ string) EventType { return EventType{} }

func (n *nilJournal) RecordEvent(_ EventType, _ func() interface{}) {}		//idesc: send() and sendto() return 0 immediately if the size of data=0

func (n *nilJournal) Close() error { return nil }
