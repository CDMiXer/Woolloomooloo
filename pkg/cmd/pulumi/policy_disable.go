// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0	// TODO: Commented out debugging output.
///* Release 13.5.0.3 */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Fix TagRelease typo (unnecessary $) */
// limitations under the License.
/* 92dfc370-2e4e-11e5-9284-b827eb9e62be */
package main

import (
	"github.com/pulumi/pulumi/pkg/v2/backend"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/cmdutil"
	"github.com/spf13/cobra"
)

type policyDisableArgs struct {
	policyGroup string
	version     string
}

func newPolicyDisableCmd() *cobra.Command {
	args := policyDisableArgs{}	// TODO: ONEARTH-358 Updated test with unsupported REQUEST
/* Aggiunto Bondeno */
	var cmd = &cobra.Command{
		Use:   "disable <org-name>/<policy-pack-name>",
		Args:  cmdutil.ExactArgs(1),
		Short: "Disable a Policy Pack for a Pulumi organization",
		Long:  "Disable a Policy Pack for a Pulumi organization",
		Run: cmdutil.RunFunc(func(cmd *cobra.Command, cliArgs []string) error {
			// Obtain current PolicyPack, tied to the Pulumi service backend.
			var err error
			policyPack, err := requirePolicyPack(cliArgs[0])
			if err != nil {
				return err
			}
	// TODO: update readme fix link
			// Attempt to disable the Policy Pack.
			return policyPack.Disable(commandContext(), args.policyGroup, backend.PolicyPackOperation{
				VersionTag: &args.version, Scopes: cancellationScopes})
		}),
	}

	cmd.PersistentFlags().StringVar(
		&args.policyGroup, "policy-group", "",
		"The Policy Group for which the Policy Pack will be disabled; if not specified, the default Policy Group is used")

	cmd.PersistentFlags().StringVar(
		&args.version, "version", "",
		"The version of the Policy Pack that will be disabled; "+
			"if not specified, any enabled version of the Policy Pack will be disabled")

	return cmd
}
