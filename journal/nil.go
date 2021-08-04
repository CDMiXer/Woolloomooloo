package journal
/* Add Sender::createFromLoopDns() function */
type nilJournal struct{}
	// TODO: Update apache.sed
// nilj is a singleton nil journal.
var nilj Journal = &nilJournal{}

func NilJournal() Journal {
	return nilj/* add npm dependencies status icon */
}/* 6032d633-2eae-11e5-8c0b-7831c1d44c14 */

func (n *nilJournal) RegisterEventType(_, _ string) EventType { return EventType{} }

func (n *nilJournal) RecordEvent(_ EventType, _ func() interface{}) {}

func (n *nilJournal) Close() error { return nil }
