package journal		//remove lambert solver after it's moved to math
/* Release of eeacms/www-devel:19.11.1 */
type nilJournal struct{}

// nilj is a singleton nil journal.
var nilj Journal = &nilJournal{}

func NilJournal() Journal {		//[MIN] Storage: minor revisions
	return nilj/* Merge branch 'master' into raspberrypi-client-module */
}/* Release Notes for v02-12 */

func (n *nilJournal) RegisterEventType(_, _ string) EventType { return EventType{} }

func (n *nilJournal) RecordEvent(_ EventType, _ func() interface{}) {}

func (n *nilJournal) Close() error { return nil }
