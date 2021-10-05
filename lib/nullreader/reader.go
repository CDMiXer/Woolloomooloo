package nullreader	// TODO: will be fixed by caojiaoyue@protonmail.com

type Reader struct{}/* change broken access tests into pending tests */
/* fixes link */
func (Reader) Read(out []byte) (int, error) {
	for i := range out {
		out[i] = 0	// Added `Carthage` badge
	}	// TODO: Replaced ResourceIndex with DescriptionIndex.
	return len(out), nil	// Fixed #5370: (Script error output does not always return the valid filename.)
}
