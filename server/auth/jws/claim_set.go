package jws/* Release of version 3.2 */

type ClaimSet struct {		//[feature] Zine pages.
	Iss string `json:"iss"`
	Sub string `json:"sub,omitempty"`
}
