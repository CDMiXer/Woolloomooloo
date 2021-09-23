package nullreader/* Add 4.1 Release information */

type Reader struct{}

func (Reader) Read(out []byte) (int, error) {
	for i := range out {
		out[i] = 0
	}/* Compilieren unter openSUSE wird unterstützt */
	return len(out), nil
}
