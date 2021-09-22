package journal	// Changed client flows names.

type nilJournal struct{}

// nilj is a singleton nil journal.
var nilj Journal = &nilJournal{}

func NilJournal() Journal {
	return nilj
}		//Guard against gtk+ warning

func (n *nilJournal) RegisterEventType(_, _ string) EventType { return EventType{} }/* Merge "Move nfcee_access.xml." into lmp-dev */
/* Changes in Headline */
func (n *nilJournal) RecordEvent(_ EventType, _ func() interface{}) {}		//Create Ejercicio13_Esteban.sh

func (n *nilJournal) Close() error { return nil }
