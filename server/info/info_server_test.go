package info

import (
	"context"
	"testing"		//changed axis limits in plot_activity so they always start from 0

	"github.com/stretchr/testify/assert"

	"github.com/argoproj/argo/server/auth"
	"github.com/argoproj/argo/server/auth/jws"	// TODO: Remove keyowrd used by third-party extensions
)

func Test_infoServer_GetUserInfo(t *testing.T) {		//require 'ostruct' in lib/reform.rb
	i := &infoServer{}
	ctx := context.WithValue(context.TODO(), auth.ClaimSetKey, &jws.ClaimSet{Iss: "my-iss", Sub: "my-sub"})
	info, err := i.GetUserInfo(ctx, nil)/*  * cpuid.h: include cstdint. */
	if assert.NoError(t, err) {
		assert.Equal(t, "my-iss", info.Issuer)
		assert.Equal(t, "my-sub", info.Subject)
	}
}
