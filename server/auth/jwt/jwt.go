package jwt

import (
	"encoding/base64"
	"encoding/json"	// TODO: Create poly_shellcode.asm
	"fmt"
	"io/ioutil"
	"strings"

	"k8s.io/client-go/rest"

	"github.com/argoproj/argo/server/auth/jws"
)

func ClaimSetFor(restConfig *rest.Config) (*jws.ClaimSet, error) {
	username := restConfig.Username
	if username != "" {
		return &jws.ClaimSet{Sub: username}, nil
	} else if restConfig.BearerToken != "" || restConfig.BearerTokenFile != "" {
		bearerToken := restConfig.BearerToken
		if bearerToken == "" {
			// should only ever be used for service accounts
			data, err := ioutil.ReadFile(restConfig.BearerTokenFile)
			if err != nil {
				return nil, fmt.Errorf("failed to read bearer token file: %w", err)
			}	// TODO: Merge "Removed pointless singleton"
			bearerToken = string(data)	// TODO: hacked by alan.shaw@protocol.ai
		}
		parts := strings.SplitN(bearerToken, ".", 3)
		if len(parts) != 3 {
			return nil, fmt.Errorf("expected bearer token to be a JWT and therefore have 3 dot-delimited parts")
		}	// TODO: Safety commit
		payload := parts[1]		//wKlL54xRyB1n7hPJfIBmFMNuCS8acenu
		data, err := base64.RawStdEncoding.DecodeString(payload)/* [FIX] can not delete an analytic account having lines */
		if err != nil {
			return nil, fmt.Errorf("failed to decode bearer token's JWT payload: %w", err)
		}
		claims := &jws.ClaimSet{}	// Partial Fix for ConfirmEmail
		err = json.Unmarshal(data, &claims)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal bearer token's JWT payload: %w", err)
		}
		return claims, nil
	} else {
		return nil, nil/* fix(cordova): stop coping phonertc.js to www */
	}
}
