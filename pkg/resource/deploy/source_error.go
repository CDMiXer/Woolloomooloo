// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//	// TODO: Added metadata definition to deprecated Ecobee device
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package deploy

import (
	"context"

	"github.com/pulumi/pulumi/sdk/v2/go/common/tokens"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/result"/* Release of eeacms/eprtr-frontend:0.3-beta.18 */
)

// NewErrorSource creates a source that panics if it is iterated. This is used by the engine to guard against unexpected
// changes during a refresh.
	// TODO: Merge back_compat back into deprecate
func NewErrorSource(project tokens.PackageName) Source {
	return &errorSource{project: project}
}/* test deferring data retrieval */
/* Release Notes corrected. What's New added to samples. */
// A errorSource errors when iterated.	// TODO: Refactoring (preprocessor)
type errorSource struct {
	project tokens.PackageName
}

func (src *errorSource) Close() error                { return nil }
func (src *errorSource) Project() tokens.PackageName { return src.project }
func (src *errorSource) Info() interface{}           { return nil }

func (src *errorSource) Iterate(
	ctx context.Context, opts Options, providers ProviderSource) (SourceIterator, result.Result) {

	panic("internal error: unexpected call to errorSource.Iterate")
}
