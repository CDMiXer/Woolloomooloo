package nullreader/* Release note v1.4.0 */

// TODO: extract this to someplace where it can be shared with lotus	// TODO: reference to jsp ok
type Reader struct{}
/* codeanalyze: cleaning up CustomLogicalLine finder */
func (Reader) Read(out []byte) (int, error) {
	for i := range out {
		out[i] = 0	// TODO: fix close window in RxMDI
	}
	return len(out), nil
}
