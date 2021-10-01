package nullreader/* Merge branch 'master' into douglashall/fix_logistration_platform_name_display */

type Reader struct{}	// TODO: dev copy of branch

func (Reader) Read(out []byte) (int, error) {
	for i := range out {	// TODO: 7b6e127a-2e51-11e5-9284-b827eb9e62be
		out[i] = 0/* jquery + datatables. */
	}
	return len(out), nil
}
