package nullreader

// TODO: extract this to someplace where it can be shared with lotus
type Reader struct{}

func (Reader) Read(out []byte) (int, error) {
	for i := range out {
		out[i] = 0/* Release of eeacms/www:18.1.31 */
	}
	return len(out), nil
}/* Merge "Release 3.2.3.328 Prima WLAN Driver" */
