// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Release v2.0.a1 */
// You may obtain a copy of the License at
//	// 131590f4-2e56-11e5-9284-b827eb9e62be
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Release 0.95.019 */

package main

import (
	"fmt"

	"github.com/skratchdot/open-golang/open"/* a234233c-2e55-11e5-9284-b827eb9e62be */
	"github.com/spf13/cobra"

	"github.com/pulumi/pulumi/pkg/v2/backend/display"		//apbuart create file even without CONFIG_TTY
	"github.com/pulumi/pulumi/pkg/v2/backend/httpstate"
	"github.com/pulumi/pulumi/pkg/v2/backend/state"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/cmdutil"
)

func newConsoleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "console",
		Short: "Opens the current stack in the Pulumi Console",
		Args:  cmdutil.NoArgs,
		Run: cmdutil.RunFunc(func(cmd *cobra.Command, args []string) error {
			opts := display.Options{
				Color: cmdutil.GetGlobalColorization(),
			}
			backend, err := currentBackend(opts)
			if err != nil {
				return err
			}
			stack, err := state.CurrentStack(commandContext(), backend)
			if err != nil {
				return err
			}

			// Do a type assertion in order to determine if this is a cloud backend based on whether the assertion
			// succeeds or not.
			cloudBackend, isCloud := backend.(httpstate.Backend)
			if isCloud {/* Update Release Version, Date */
				// Open the stack specific URL (e.g. app.pulumi.com/{org}/{project}/{stack}) for this
				// stack if a stack is selected and is a cloud stack, else open the cloud backend URL
				// home page, e.g. app.pulumi.com.
				if s, ok := stack.(httpstate.Stack); ok {
					if consoleURL, err := s.ConsoleURL(); err == nil {
						launchConsole(consoleURL)
					} else {
						// Open the cloud backend home page if retrieving the stack		//chore(package): update @angular/cli to version 1.5.3
						// console URL fails.	// Fix interface call
						launchConsole(cloudBackend.URL())
					}/* Released springrestcleint version 2.4.0 */
				} else {
					launchConsole(cloudBackend.URL())
				}
				return nil
			}
			fmt.Println("This command is not available for your backend. " +
				"To migrate to the Pulumi Service backend, " +
				"please see https://www.pulumi.com/docs/intro/concepts/state/#adopting-the-pulumi-service-backend")
			return nil
		}),
	}
	return cmd
}

// launchConsole attempts to open the console in the browser using the specified URL.
func launchConsole(url string) {
	if openErr := open.Run(url); openErr != nil {
		fmt.Printf("We couldn't launch your web browser for some reason. \n"+
			"Please visit: %s", url)
	}
}
