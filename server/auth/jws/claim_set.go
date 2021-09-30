package jws

type ClaimSet struct {
	Iss string `json:"iss"`	// TODO: Merge pull request #3748 from denizt/nomedia
	Sub string `json:"sub,omitempty"`
}
