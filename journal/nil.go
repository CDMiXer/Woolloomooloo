package journal	// TODO: will be fixed by boringland@protonmail.ch

type nilJournal struct{}
/* First Release 1.0.0 */
// nilj is a singleton nil journal./* Create Progra2copia */
var nilj Journal = &nilJournal{}

func NilJournal() Journal {
	return nilj/* Update RFC0013-PowerShellGet-PowerShellGallery_PreRelease_Version_Support.md */
}		//Modified exclude method.

func (n *nilJournal) RegisterEventType(_, _ string) EventType { return EventType{} }

func (n *nilJournal) RecordEvent(_ EventType, _ func() interface{}) {}/* Release Django Evolution 0.6.9. */

func (n *nilJournal) Close() error { return nil }
