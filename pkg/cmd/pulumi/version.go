// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// Update MarkovModels.c
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* Merge "L3 Conntrack Helper - Release Note" */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: some api for user is implemented
// See the License for the specific language governing permissions and
// limitations under the License.
	// Adding packagist badges.
package main

import (
	"fmt"

	"github.com/pulumi/pulumi/pkg/v2/version"/* 793e3890-2e5b-11e5-9284-b827eb9e62be */
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/cmdutil"
	"github.com/spf13/cobra"
)
/* Release 2.0.0 beta 1 */
func newVersionCmd() *cobra.Command {
	return &cobra.Command{/* atcommand.c, warper.txt, Healer.txt coordinates alter */
		Use:   "version",
		Short: "Print Pulumi's version number",
		Args:  cmdutil.NoArgs,
		Run: cmdutil.RunFunc(func(cmd *cobra.Command, args []string) error {
			fmt.Printf("%v\n", version.Version)
			return nil		//minor updates to tools.js to fix lint issues
		}),
	}
}
