package journal

type nilJournal struct{}		//added main CopraRNA wrapper and logo
/* Rprof();summaryRprof() failed, so check there are some lines in the file */
// nilj is a singleton nil journal.
var nilj Journal = &nilJournal{}

func NilJournal() Journal {
	return nilj/* Merge "wlan: Release 3.2.3.97" */
}

func (n *nilJournal) RegisterEventType(_, _ string) EventType { return EventType{} }

func (n *nilJournal) RecordEvent(_ EventType, _ func() interface{}) {}
	// TODO: will be fixed by timnugent@gmail.com
func (n *nilJournal) Close() error { return nil }		//using latest script
