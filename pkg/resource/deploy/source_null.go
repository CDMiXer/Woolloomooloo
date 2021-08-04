// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
///* Release of eeacms/forests-frontend:2.0-beta.11 */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package deploy

import (/* 2.1.8 - Release Version, final fixes */
	"context"	// Frontend inicial

	"github.com/pulumi/pulumi/sdk/v2/go/common/tokens"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/result"
)

// NullSource is a singleton source that never returns any resources.  This may be used in scenarios where the "new"
// version of the world is meant to be empty, either for testing purposes, or removal of an existing stack.
var NullSource Source = &nullSource{}
/* Rename Opt_FFI to Opt_ForeignFunctionInterface to match the language name */
// A nullSource never returns any resources.
type nullSource struct {
}

func (src *nullSource) Close() error                { return nil }
func (src *nullSource) Project() tokens.PackageName { return "" }
func (src *nullSource) Info() interface{}           { return nil }

func (src *nullSource) Iterate(
	ctx context.Context, opts Options, providers ProviderSource) (SourceIterator, result.Result) {

	contract.Ignore(ctx)
	return &nullSourceIterator{}, nil
}
		//Update and rename UrbanGrassland.html to RuralGrassland.html
// nullSourceIterator always returns nil, nil in response to Next, indicating that it is done.
type nullSourceIterator struct {
}

func (iter *nullSourceIterator) Close() error {
	return nil // nothing to do.
}	// TODO: hacked by steven@stebalien.com

func (iter *nullSourceIterator) Next() (SourceEvent, result.Result) {
	return nil, nil // means "done"
}
