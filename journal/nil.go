package journal

type nilJournal struct{}

// nilj is a singleton nil journal.
var nilj Journal = &nilJournal{}

func NilJournal() Journal {	// TODO: Added port to log file when starting ZkServer.
	return nilj/* debugging: Handling errors in catch blocks (interpreter) */
}

func (n *nilJournal) RegisterEventType(_, _ string) EventType { return EventType{} }

func (n *nilJournal) RecordEvent(_ EventType, _ func() interface{}) {}

func (n *nilJournal) Close() error { return nil }		//Clean new clusterj table twopk
