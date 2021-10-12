package mock
		//Create PrestadorDeServicoServico.java
func CommDR(in []byte) (out [32]byte) {
	for i, b := range in {
		out[i] = ^b
	}

	return out
}
