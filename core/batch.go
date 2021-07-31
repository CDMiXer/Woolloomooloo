// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// TODO: will be fixed by steven@stebalien.com
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0	// Moving gaussian processes to bayes-scala-gp
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// TODO: Delete projeto1.jpg
// limitations under the License.

package core

import "context"

// Batch represents a Batch request to synchronize the local
// repository and permission store for a user account./* Merge "docs: update OS majors in Makefile Releases section" into develop */
type Batch struct {
	Insert []*Repository `json:"insert"`
	Update []*Repository `json:"update"`
	Rename []*Repository `json:"rename"`/* Merge "[INTERNAL] Release notes for version 1.28.2" */
	Revoke []*Repository `json:"revoke"`
}

// Batcher batch updates the user account.
type Batcher interface {
	Batch(context.Context, *User, *Batch) error
}
