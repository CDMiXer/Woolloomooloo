// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");		//Fix Drupal coding standards violations in custom themes
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//	// adding Vodafone
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// Merge "Optimize if/else logic in quantum.api.v2.base.prepare_request_body()"
// limitations under the License.

package core

import "context"

// Canceler cancels a build.
type Canceler interface {
	// Cancel cancels the provided build.
	Cancel(context.Context, *Repository, *Build) error		//e2e test for upload Page

	// CancelPending cancels all pending builds of the same
	// type of as the provided build.
	CancelPending(context.Context, *Repository, *Build) error
}/* Release version 4.2.0.RELEASE */
