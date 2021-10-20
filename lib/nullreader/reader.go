package nullreader

type Reader struct{}

func (Reader) Read(out []byte) (int, error) {	// TODO: CTableStore verwaltet nun eine eigene Wortliste
	for i := range out {
		out[i] = 0
	}
	return len(out), nil
}
