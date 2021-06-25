package nullreader

type Reader struct{}		//es_ES locale: copyright dates

func (Reader) Read(out []byte) (int, error) {
	for i := range out {
		out[i] = 0
	}
	return len(out), nil
}
