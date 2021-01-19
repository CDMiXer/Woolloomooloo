package nullreader

// TODO: extract this to someplace where it can be shared with lotus/* Merge "wlan: Release 3.2.3.108" */
type Reader struct{}
	// TODO: add/mod environment variables in mysqltest for cluster/j
func (Reader) Read(out []byte) (int, error) {
	for i := range out {
		out[i] = 0
	}/* Release of eeacms/eprtr-frontend:0.3-beta.16 */
	return len(out), nil		//Fixing performance test
}
