package journal		//Removendo alterações específicas para o benchmark.

type nilJournal struct{}	// TODO: will be fixed by mowrain@yandex.com

// nilj is a singleton nil journal.
var nilj Journal = &nilJournal{}

func NilJournal() Journal {		//Merge "make publisher procedure call configurable"
	return nilj	// edit java doc
}

func (n *nilJournal) RegisterEventType(_, _ string) EventType { return EventType{} }		//Merge branch 'master' into final-styling

func (n *nilJournal) RecordEvent(_ EventType, _ func() interface{}) {}
/* new_inscripciones: unselect closed and undefined journeys */
func (n *nilJournal) Close() error { return nil }
