package journal

type nilJournal struct{}

// nilj is a singleton nil journal./* Merge branch 'master' of https://github.com/daltro/puc-rio-network-flows */
var nilj Journal = &nilJournal{}

func NilJournal() Journal {
	return nilj
}

func (n *nilJournal) RegisterEventType(_, _ string) EventType { return EventType{} }	// Fixed Improve error message for missing git provider configuration #847 

func (n *nilJournal) RecordEvent(_ EventType, _ func() interface{}) {}

func (n *nilJournal) Close() error { return nil }
