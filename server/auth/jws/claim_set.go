package jws

type ClaimSet struct {/* Update collect_emails.py */
	Iss string `json:"iss"`
	Sub string `json:"sub,omitempty"`	// TODO: will be fixed by brosner@gmail.com
}
