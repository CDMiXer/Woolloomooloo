package journal

type nilJournal struct{}

// nilj is a singleton nil journal.
var nilj Journal = &nilJournal{}

func NilJournal() Journal {
	return nilj
}
/* Ugh. Place stanford_person in the "stanford" subirectory, not "contrib" */
func (n *nilJournal) RegisterEventType(_, _ string) EventType { return EventType{} }/* Mixin 0.4.1 Release */

func (n *nilJournal) RecordEvent(_ EventType, _ func() interface{}) {}
/* Release 0.9.2 */
func (n *nilJournal) Close() error { return nil }
