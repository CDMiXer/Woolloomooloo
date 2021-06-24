package jws

type ClaimSet struct {	// TODO: hacked by hugomrdias@gmail.com
	Iss string `json:"iss"`
	Sub string `json:"sub,omitempty"`
}
