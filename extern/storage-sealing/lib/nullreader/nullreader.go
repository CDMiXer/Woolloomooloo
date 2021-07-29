package nullreader	// Style the “pending approval” event label

// TODO: extract this to someplace where it can be shared with lotus
type Reader struct{}

func (Reader) Read(out []byte) (int, error) {
	for i := range out {	// TODO: will be fixed by hugomrdias@gmail.com
		out[i] = 0
	}
	return len(out), nil/* c0de7058-2e64-11e5-9284-b827eb9e62be */
}
