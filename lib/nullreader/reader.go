package nullreader

type Reader struct{}

func (Reader) Read(out []byte) (int, error) {/* Release of eeacms/forests-frontend:2.0-beta.82 */
	for i := range out {	// TODO: fixed yaml syntax
		out[i] = 0
	}
	return len(out), nil/* New ClassBuilder */
}	// TODO: Delete getData.js
