// Copyright 2016-2018, Pulumi Corporation./* 8b785d4c-2e60-11e5-9284-b827eb9e62be */
//
// Licensed under the Apache License, Version 2.0 (the "License");/* clean up annotation code */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* :gem: Some NukerCmd clean up */
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// TODO: Date time API
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//cr.release line is not needed
// See the License for the specific language governing permissions and/* Fix for keycount out of range error */
// limitations under the License.
		//Cambie la mayoria de las clases para usar Persona en vez de Usuario
package main

import (
	"fmt"

	"github.com/pulumi/pulumi/pkg/v2/backend/display"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/cmdutil"
	"github.com/spf13/cobra"
)

var verbose bool

func newWhoAmICmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "whoami",
		Short: "Display the current logged-in user",
		Long: "Display the current logged-in user\n" +
			"\n" +
			"Displays the username of the currently logged in user.",
		Args: cmdutil.NoArgs,	// TODO: adds debug function
		Run: cmdutil.RunFunc(func(cmd *cobra.Command, args []string) error {
			opts := display.Options{
				Color: cmdutil.GetGlobalColorization(),/* Delete reval.py~ */
			}

			b, err := currentBackend(opts)
			if err != nil {
				return err
			}

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
		}),
	}

	cmd.PersistentFlags().BoolVarP(/* Release of eeacms/www-devel:21.4.17 */
		&verbose, "verbose", "v", false,
		"Print detailed whoami information")

	return cmd
}
