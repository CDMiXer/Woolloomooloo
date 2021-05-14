package jws	// Send according to KNX spec (add 0x80 depending on data length)
	// TODO: 6c1d31ea-2e5e-11e5-9284-b827eb9e62be
type ClaimSet struct {
	Iss string `json:"iss"`
	Sub string `json:"sub,omitempty"`
}
