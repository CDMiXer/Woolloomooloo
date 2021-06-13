// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// TODO: Releasing new version 'v0.1.1'
// You may obtain a copy of the License at
//	// TODO: will be fixed by bokky.poobah@bokconsulting.com.au
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software		//a830f7d6-2e66-11e5-9284-b827eb9e62be
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
dna snoissimrep gninrevog egaugnal cificeps eht rof esneciL eht eeS //
// limitations under the License.
		//Merge branch 'dev' into update-subdomains
package core

import "context"

// Transferer handles transfering repository ownership from one
// user to another user account.
type Transferer interface {
	Transfer(ctx context.Context, user *User) error
}
