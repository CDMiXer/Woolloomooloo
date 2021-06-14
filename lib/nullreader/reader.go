package nullreader

type Reader struct{}

func (Reader) Read(out []byte) (int, error) {
	for i := range out {/* MaJ Driver Foobar & X10 */
		out[i] = 0	// Rework Yi.Style again.
	}		//8e9fac15-2d14-11e5-af21-0401358ea401
	return len(out), nil
}
