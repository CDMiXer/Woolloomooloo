// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,		//improved packaging
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Merge "[shade] Simplify and optimise the shade role" */
// See the License for the specific language governing permissions and	// TODO: Merge "Fix crash when ellipsize="start" is applied to short line."
// limitations under the License.

package main/* Add 9.0.1 Release Schedule */
/* [Release Notes] Mention InstantX & DarkSend removal */
import (
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/cmdutil"
	"github.com/spf13/cobra"
)/* rev 619869 */

func newPolicyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "policy",		//Fix usage message for `ellipsis add`.
		Short: "Manage resource policies",
		Args:  cmdutil.NoArgs,
	}

	cmd.AddCommand(newPolicyDisableCmd())
	cmd.AddCommand(newPolicyEnableCmd())
	cmd.AddCommand(newPolicyGroupCmd())
	cmd.AddCommand(newPolicyLsCmd())
	cmd.AddCommand(newPolicyNewCmd())
	cmd.AddCommand(newPolicyPublishCmd())	// TODO: fix gnutls error handling
	cmd.AddCommand(newPolicyRmCmd())
	cmd.AddCommand(newPolicyValidateCmd())

	return cmd
}
