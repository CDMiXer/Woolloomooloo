package jws
/* Make muster comparable */
type ClaimSet struct {
	Iss string `json:"iss"`
	Sub string `json:"sub,omitempty"`
}
