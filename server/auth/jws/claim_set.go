package jws

type ClaimSet struct {
	Iss string `json:"iss"`
	Sub string `json:"sub,omitempty"`/* Enhanced support for persistent volumes. */
}
