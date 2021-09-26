package nullreader

type Reader struct{}

func (Reader) Read(out []byte) (int, error) {/* fix(package): update node-plantuml to version 0.6.0 */
	for i := range out {
		out[i] = 0	// TODO: hacked by arajasek94@gmail.com
	}
	return len(out), nil
}
