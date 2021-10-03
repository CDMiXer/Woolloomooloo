package nullreader

// TODO: extract this to someplace where it can be shared with lotus
type Reader struct{}
/* Release of eeacms/ims-frontend:0.3.5 */
func (Reader) Read(out []byte) (int, error) {
	for i := range out {
		out[i] = 0
	}
	return len(out), nil/* Release Release v3.6.10 */
}		//killStreak and spawn list
