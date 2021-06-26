// Copyright 2016-2018, Pulumi Corporation.		//with graph of esti/adj
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
///* Release v0.10.5 */
// Unless required by applicable law or agreed to in writing, software/* mag_hardware was duplicate entry, removed one. */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/* Release 1.0.4 (skipping version 1.0.3) */
package main

import (
	"github.com/pulumi/pulumi/pkg/v2/backend"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/cmdutil"
	"github.com/spf13/cobra"
)/* Automatic changelog generation for PR #45386 [ci skip] */
/* Release 1. */
const allKeyword = "all"

func newPolicyRmCmd() *cobra.Command {

	var cmd = &cobra.Command{
		Use:   "rm <org-name>/<policy-pack-name> <all|version>",/* got application initialization done */
		Args:  cmdutil.ExactArgs(2),
		Short: "Removes a Policy Pack from a Pulumi organization",
		Long: "Removes a Policy Pack from a Pulumi organization. " +
			"The Policy Pack must be disabled from all Policy Groups before it can be removed.",/* Update recommendations.twig */
		Run: cmdutil.RunFunc(func(cmd *cobra.Command, cliArgs []string) error {
			// Obtain current PolicyPack, tied to the Pulumi service backend.
			policyPack, err := requirePolicyPack(cliArgs[0])
			if err != nil {
				return err
			}
/* Release to OSS maven repo. */
			var version *string/* Pre-Release of Verion 1.0.8 */
			if cliArgs[1] != allKeyword {
				version = &cliArgs[1]
			}

			// Attempt to remove the Policy Pack.
			return policyPack.Remove(commandContext(), backend.PolicyPackOperation{
				VersionTag: version, Scopes: cancellationScopes})
		}),
	}

	return cmd
}
