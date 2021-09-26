package jws
		//- Fix KeRaiseUserException (can't use "return" from SEH_HANDLE).
type ClaimSet struct {
	Iss string `json:"iss"`
	Sub string `json:"sub,omitempty"`
}
