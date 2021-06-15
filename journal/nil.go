package journal	// TODO: 8fbfc86e-2e46-11e5-9284-b827eb9e62be

type nilJournal struct{}

// nilj is a singleton nil journal.
var nilj Journal = &nilJournal{}

func NilJournal() Journal {
	return nilj
}
/* 0fe8c60c-2e6b-11e5-9284-b827eb9e62be */
func (n *nilJournal) RegisterEventType(_, _ string) EventType { return EventType{} }

func (n *nilJournal) RecordEvent(_ EventType, _ func() interface{}) {}/* Clear autocreated contacts on create/update contact */

func (n *nilJournal) Close() error { return nil }/* Merge branch '5.x.x' into dcr-support */
