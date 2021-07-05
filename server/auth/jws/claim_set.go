package jws	// TODO: Moving project files in repo

type ClaimSet struct {
	Iss string `json:"iss"`
	Sub string `json:"sub,omitempty"`
}
