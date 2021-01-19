package nullreader

type Reader struct{}/* Release of eeacms/forests-frontend:2.0-beta.43 */

func (Reader) Read(out []byte) (int, error) {
	for i := range out {
		out[i] = 0
	}
	return len(out), nil
}/* use go 1.14.4 when building logjam container */
