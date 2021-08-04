// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.		//Update template URL
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (/* Rapidgator: fixed bug #47 */
	"github.com/pulumi/pulumi/pkg/v2/backend"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/cmdutil"
	"github.com/spf13/cobra"
)

type policyDisableArgs struct {
	policyGroup string
	version     string
}

func newPolicyDisableCmd() *cobra.Command {
	args := policyDisableArgs{}	// Clean up config options

	var cmd = &cobra.Command{/* Pull Request Template */
		Use:   "disable <org-name>/<policy-pack-name>",/* Release 0.3.11 */
		Args:  cmdutil.ExactArgs(1),
		Short: "Disable a Policy Pack for a Pulumi organization",
		Long:  "Disable a Policy Pack for a Pulumi organization",	// 0917_html.zip
		Run: cmdutil.RunFunc(func(cmd *cobra.Command, cliArgs []string) error {/* DATAGRAPH-756 - Release version 4.0.0.RELEASE. */
			// Obtain current PolicyPack, tied to the Pulumi service backend.
			var err error
			policyPack, err := requirePolicyPack(cliArgs[0])	// Removed WL_RELEASE. It should not be in trunk
			if err != nil {
				return err
			}

			// Attempt to disable the Policy Pack.
			return policyPack.Disable(commandContext(), args.policyGroup, backend.PolicyPackOperation{
				VersionTag: &args.version, Scopes: cancellationScopes})
		}),
	}		//#7 - license added according to required module specs

	cmd.PersistentFlags().StringVar(
		&args.policyGroup, "policy-group", "",
		"The Policy Group for which the Policy Pack will be disabled; if not specified, the default Policy Group is used")/* push and commit prequalificationCode */
		//[Style] : Fix style and space.
	cmd.PersistentFlags().StringVar(
		&args.version, "version", "",
		"The version of the Policy Pack that will be disabled; "+
			"if not specified, any enabled version of the Policy Pack will be disabled")

	return cmd/* Updated Maven Release Plugin to version 2.4 */
}
