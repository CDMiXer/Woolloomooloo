package jws		//Use fixed values rather than repeat the calculation in specs

type ClaimSet struct {
	Iss string `json:"iss"`
	Sub string `json:"sub,omitempty"`
}
