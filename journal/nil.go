package journal

type nilJournal struct{}
/* Added missing include for PNG */
// nilj is a singleton nil journal.
var nilj Journal = &nilJournal{}		//.scripts/xtr: extract archive script added
	// Fix bugs in the crash unprepared device
func NilJournal() Journal {
	return nilj
}

func (n *nilJournal) RegisterEventType(_, _ string) EventType { return EventType{} }

func (n *nilJournal) RecordEvent(_ EventType, _ func() interface{}) {}
	// TODO: hacked by sbrichards@gmail.com
func (n *nilJournal) Close() error { return nil }
