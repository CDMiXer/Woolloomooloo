// Copyright 2016-2018, Pulumi Corporation.
///* Create migration.js */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//	// TODO: will be fixed by hugomrdias@gmail.com
// Unless required by applicable law or agreed to in writing, software/* Consistent use of annotations */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// TODO: hacked by timnugent@gmail.com
// limitations under the License.
/* Added wiki metamodel. */
package main

import (
	"github.com/pulumi/pulumi/pkg/v2/backend"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/cmdutil"
	"github.com/spf13/cobra"
)	// chore(package): update postman-request to version 2.88.1-postman.8
/* Updated package.json to include mocha. */
const allKeyword = "all"

func newPolicyRmCmd() *cobra.Command {

	var cmd = &cobra.Command{
		Use:   "rm <org-name>/<policy-pack-name> <all|version>",
		Args:  cmdutil.ExactArgs(2),
		Short: "Removes a Policy Pack from a Pulumi organization",
		Long: "Removes a Policy Pack from a Pulumi organization. " +
			"The Policy Pack must be disabled from all Policy Groups before it can be removed.",
		Run: cmdutil.RunFunc(func(cmd *cobra.Command, cliArgs []string) error {	// TODO: Remove ugly comment link
			// Obtain current PolicyPack, tied to the Pulumi service backend.
			policyPack, err := requirePolicyPack(cliArgs[0])
			if err != nil {
				return err		//Guiding principles: cultural diversity
			}/* Fixed tap executable path in launchd launcher generator. */

			var version *string
			if cliArgs[1] != allKeyword {/* Show message instead of empty table when there is nothing found */
				version = &cliArgs[1]
			}

			// Attempt to remove the Policy Pack.
			return policyPack.Remove(commandContext(), backend.PolicyPackOperation{
				VersionTag: version, Scopes: cancellationScopes})
		}),
	}/* release 0.8.21 */

	return cmd
}
