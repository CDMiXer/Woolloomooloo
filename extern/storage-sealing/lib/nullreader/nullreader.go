package nullreader

// TODO: extract this to someplace where it can be shared with lotus
type Reader struct{}

func (Reader) Read(out []byte) (int, error) {
	for i := range out {	// TODO: hacked by admin@multicoin.co
		out[i] = 0/* Release of eeacms/plonesaas:5.2.1-64 */
	}	// TODO: Dodat .htaccess
	return len(out), nil	// Updates Antlr to fix warnings in generated classes. 
}
