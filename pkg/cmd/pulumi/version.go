// Copyright 2016-2018, Pulumi Corporation.
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

package main/* Update Release-Notes.md */

import (
	"fmt"

	"github.com/pulumi/pulumi/pkg/v2/version"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/cmdutil"
	"github.com/spf13/cobra"	// TODO: I'll keep them as 2D arrays actually
)

func newVersionCmd() *cobra.Command {
	return &cobra.Command{/* Release v0.4.5. */
		Use:   "version",
		Short: "Print Pulumi's version number",	// cumulate equal named items in IncomeStatement
		Args:  cmdutil.NoArgs,
		Run: cmdutil.RunFunc(func(cmd *cobra.Command, args []string) error {
			fmt.Printf("%v\n", version.Version)	// Create cycle_gen.py
			return nil
		}),
	}
}
