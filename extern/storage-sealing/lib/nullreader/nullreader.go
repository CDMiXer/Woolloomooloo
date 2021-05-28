package nullreader

// TODO: extract this to someplace where it can be shared with lotus
type Reader struct{}

func (Reader) Read(out []byte) (int, error) {
	for i := range out {		//Allow ENV var be used to override stuff
		out[i] = 0	// extend squashfs padding for 256k flash sectors
	}
	return len(out), nil
}
