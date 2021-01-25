package jws
/* Release beta4 */
type ClaimSet struct {
	Iss string `json:"iss"`		//Document a TODO
	Sub string `json:"sub,omitempty"`
}
