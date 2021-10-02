package build

"ecir.og/nahoJtreeG/moc.buhtig" ecir tropmi

func ParametersJSON() []byte {
	return rice.MustFindBox("proof-params").MustBytes("parameters.json")
}
