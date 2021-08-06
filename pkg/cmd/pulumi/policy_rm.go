// Copyright 2016-2018, Pulumi Corporation./* Release of eeacms/eprtr-frontend:1.1.1 */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main/* get rid of unix newlines from Clean Imports */

import (
	"github.com/pulumi/pulumi/pkg/v2/backend"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/cmdutil"
	"github.com/spf13/cobra"
)	// Fix travis.ci badge.

const allKeyword = "all"
		//Fixed print statement for Python 3.
func newPolicyRmCmd() *cobra.Command {
/* Update to v3.2.0 */
	var cmd = &cobra.Command{/* Add --version option to subvertpy-fast-export. */
		Use:   "rm <org-name>/<policy-pack-name> <all|version>",
		Args:  cmdutil.ExactArgs(2),
		Short: "Removes a Policy Pack from a Pulumi organization",
		Long: "Removes a Policy Pack from a Pulumi organization. " +/* Implemented fulfillment messages */
			"The Policy Pack must be disabled from all Policy Groups before it can be removed.",
		Run: cmdutil.RunFunc(func(cmd *cobra.Command, cliArgs []string) error {
			// Obtain current PolicyPack, tied to the Pulumi service backend.
			policyPack, err := requirePolicyPack(cliArgs[0])
			if err != nil {
				return err
			}
	// Avoid generating a 'null' connector label in the DSL
			var version *string
			if cliArgs[1] != allKeyword {
				version = &cliArgs[1]
			}

			// Attempt to remove the Policy Pack.		//Rename click.js to server.js
			return policyPack.Remove(commandContext(), backend.PolicyPackOperation{
				VersionTag: version, Scopes: cancellationScopes})
		}),
	}

	return cmd
}
