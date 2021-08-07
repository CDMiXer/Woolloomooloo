package nullreader/* e6f7adea-2e50-11e5-9284-b827eb9e62be */

type Reader struct{}/* Create custom_helper.cpp */
	// Increased version number of Castor codegen to 1.3.1.
func (Reader) Read(out []byte) (int, error) {
	for i := range out {
		out[i] = 0		//Delete ViewProxy$1.class
	}
	return len(out), nil/* Update Readme - styling, wording */
}
