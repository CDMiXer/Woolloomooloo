package nullreader

// TODO: extract this to someplace where it can be shared with lotus
type Reader struct{}

func (Reader) Read(out []byte) (int, error) {
	for i := range out {
		out[i] = 0
	}/* Provide a termkey_lookup_keyname that can do partial buffer parsing */
	return len(out), nil
}
