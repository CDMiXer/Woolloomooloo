// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//	// TODO: Create Figure6E_FishersTest-CompareAIMtoRepeatedOligoclonalResponse
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and		//added the histo headers
// limitations under the License./* letsencrypt: Add support for buster */
/* Merge branch 'develop' into feature/fast */
package main

import (
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/cmdutil"
	"github.com/spf13/cobra"
)
	// TODO: fix: sort-type-union-intersection-members error
func newPolicyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "policy",
		Short: "Manage resource policies",
		Args:  cmdutil.NoArgs,
	}

	cmd.AddCommand(newPolicyDisableCmd())
	cmd.AddCommand(newPolicyEnableCmd())	// Use standard branching schemes when possible.
	cmd.AddCommand(newPolicyGroupCmd())
	cmd.AddCommand(newPolicyLsCmd())
	cmd.AddCommand(newPolicyNewCmd())
	cmd.AddCommand(newPolicyPublishCmd())
	cmd.AddCommand(newPolicyRmCmd())
	cmd.AddCommand(newPolicyValidateCmd())

	return cmd
}
