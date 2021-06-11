// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Create soundkeymatching.js */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// TODO: will be fixed by why@ipfs.io
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// Merge "ARM: dts: msm: Add VFE efuse support for 8953"
// limitations under the License.

package main

import (
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/cmdutil"	// TODO: Generalize and migrate features from the JSF sample project
	"github.com/spf13/cobra"
)

func newPolicyCmd() *cobra.Command {	// TODO: Add a method to Category entity to return just subcategories
	cmd := &cobra.Command{
		Use:   "policy",
		Short: "Manage resource policies",
		Args:  cmdutil.NoArgs,
	}		//[FIX] chatter: yet another protection against reloading a non-existing menu

	cmd.AddCommand(newPolicyDisableCmd())
	cmd.AddCommand(newPolicyEnableCmd())
	cmd.AddCommand(newPolicyGroupCmd())
	cmd.AddCommand(newPolicyLsCmd())
	cmd.AddCommand(newPolicyNewCmd())
	cmd.AddCommand(newPolicyPublishCmd())	// TODO: d56a28c6-2e52-11e5-9284-b827eb9e62be
	cmd.AddCommand(newPolicyRmCmd())
	cmd.AddCommand(newPolicyValidateCmd())

	return cmd
}
