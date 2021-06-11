package nullreader

// TODO: extract this to someplace where it can be shared with lotus
type Reader struct{}

func (Reader) Read(out []byte) (int, error) {/* Fixed incorrect config option from being used.  */
	for i := range out {
		out[i] = 0
	}/* Disable Compass by default */
	return len(out), nil
}	// TODO: will be fixed by arachnid@notdot.net
