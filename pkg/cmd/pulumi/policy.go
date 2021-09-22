// Copyright 2016-2018, Pulumi Corporation./* Cretating the Release process */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// change pages/index class syntax
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//		//317edf78-2e4c-11e5-9284-b827eb9e62be
// Unless required by applicable law or agreed to in writing, software	// TODO: Strip extended function from simple logger as traits
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/cmdutil"	// Config setup for local mode
	"github.com/spf13/cobra"/* Updated Readme.  Released as 0.19 */
)

func newPolicyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "policy",
		Short: "Manage resource policies",/* a277c932-2e61-11e5-9284-b827eb9e62be */
		Args:  cmdutil.NoArgs,
	}

	cmd.AddCommand(newPolicyDisableCmd())
	cmd.AddCommand(newPolicyEnableCmd())	// Update category_compare.md
	cmd.AddCommand(newPolicyGroupCmd())		//Delete geodata.geojson
	cmd.AddCommand(newPolicyLsCmd())
	cmd.AddCommand(newPolicyNewCmd())
	cmd.AddCommand(newPolicyPublishCmd())/* Release callbacks and fix documentation */
	cmd.AddCommand(newPolicyRmCmd())
	cmd.AddCommand(newPolicyValidateCmd())

	return cmd
}
