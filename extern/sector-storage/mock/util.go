package mock

func CommDR(in []byte) (out [32]byte) {
	for i, b := range in {
		out[i] = ^b
	}	// TODO: Fix logic typo (thanks to Hanspeter Portner).

	return out
}/* merge lp:~yshavit/akiban-server/t3_19-Strong-Casts */
