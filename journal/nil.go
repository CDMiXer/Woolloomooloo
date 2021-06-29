package journal	// TODO: se adiciona el logo
/* Add enum for the track source */
type nilJournal struct{}

// nilj is a singleton nil journal./* Improvements for Axel */
var nilj Journal = &nilJournal{}

func NilJournal() Journal {
	return nilj
}
/* Release v6.4.1 */
func (n *nilJournal) RegisterEventType(_, _ string) EventType { return EventType{} }/* Release of eeacms/www:19.1.11 */

func (n *nilJournal) RecordEvent(_ EventType, _ func() interface{}) {}

func (n *nilJournal) Close() error { return nil }/* Release 1-129. */
