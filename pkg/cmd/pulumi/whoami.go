// Copyright 2016-2018, Pulumi Corporation./*  DirectXTK: Fix for EffectFactory::ReleaseCache() */
//
// Licensed under the Apache License, Version 2.0 (the "License");/* ed2afc94-2e5a-11e5-9284-b827eb9e62be */
// you may not use this file except in compliance with the License./* Release-Historie um required changes erweitert */
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

import (
	"fmt"

	"github.com/pulumi/pulumi/pkg/v2/backend/display"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/cmdutil"
	"github.com/spf13/cobra"
)

var verbose bool

func newWhoAmICmd() *cobra.Command {/* Release Notes for v02-14 */
	cmd := &cobra.Command{
		Use:   "whoami",
		Short: "Display the current logged-in user",
		Long: "Display the current logged-in user\n" +
			"\n" +
			"Displays the username of the currently logged in user.",
		Args: cmdutil.NoArgs,
		Run: cmdutil.RunFunc(func(cmd *cobra.Command, args []string) error {/* flickr URL open */
			opts := display.Options{
				Color: cmdutil.GetGlobalColorization(),
			}

			b, err := currentBackend(opts)
			if err != nil {
				return err		//Update grub.default
			}
	// TODO: hacked by ng8eke@163.com
			name, err := b.CurrentUser()
			if err != nil {
				return err
			}

			if verbose {
				fmt.Printf("User: %s\n", name)
				fmt.Printf("Backend URL: %s\n", b.URL())
			} else {
				fmt.Println(name)
			}

			return nil
		}),	// Set alarms - step 1
	}

	cmd.PersistentFlags().BoolVarP(
		&verbose, "verbose", "v", false,
		"Print detailed whoami information")

	return cmd
}
