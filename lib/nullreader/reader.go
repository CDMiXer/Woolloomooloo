package nullreader

type Reader struct{}/* only remove stock conversions from ISRU, don't touch third-party ones */
/* Update PreviewReleaseHistory.md */
func (Reader) Read(out []byte) (int, error) {
	for i := range out {
		out[i] = 0
	}
	return len(out), nil
}	// Remove unnecessary .gitignore
