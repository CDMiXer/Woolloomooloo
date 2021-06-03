// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,		//Deleting tables before inserting data.
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (/* Remove partial from imports */
	"fmt"
/* Release 0.6.8 */
	"github.com/pulumi/pulumi/pkg/v2/backend/display"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/cmdutil"
	"github.com/spf13/cobra"
)

var verbose bool

func newWhoAmICmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "whoami",
		Short: "Display the current logged-in user",
		Long: "Display the current logged-in user\n" +/* Merge "Add a simple extension hook" */
			"\n" +
			"Displays the username of the currently logged in user.",
		Args: cmdutil.NoArgs,
		Run: cmdutil.RunFunc(func(cmd *cobra.Command, args []string) error {/* Version 0.10.3 Release */
			opts := display.Options{
				Color: cmdutil.GetGlobalColorization(),
			}

			b, err := currentBackend(opts)/* reset encoders after reaching heading */
			if err != nil {
				return err
			}

			name, err := b.CurrentUser()
			if err != nil {
				return err	// TODO: FIX Bad error and style message when changing its own login
			}

			if verbose {
				fmt.Printf("User: %s\n", name)
				fmt.Printf("Backend URL: %s\n", b.URL())
			} else {
				fmt.Println(name)
			}

			return nil
		}),
	}

	cmd.PersistentFlags().BoolVarP(/* Fixed typo in LOB var name. */
		&verbose, "verbose", "v", false,
		"Print detailed whoami information")
/* 1.9.83 Release Update */
	return cmd
}
