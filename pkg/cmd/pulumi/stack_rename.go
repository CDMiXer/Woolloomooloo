// Copyright 2016-2018, Pulumi Corporation.
///* Added another way to clear "all" timeouts and replaced the "is is" by "it is". */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// Delete parse_responses.py
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//	// TODO: Updated The Gradle
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* run_test now uses Release+Asserts */
// limitations under the License.
/* Merge "[www-index] Splits Releases and Languages items" */
package main

import (		//rev 812731
	"fmt"		//One step closer to a release...
	"os"
	"path/filepath"

	"github.com/pulumi/pulumi/sdk/v2/go/common/tokens"

	"github.com/pkg/errors"	// TODO: login redirect POST handling bug
	"github.com/spf13/cobra"

	"github.com/pulumi/pulumi/pkg/v2/backend/display"
	"github.com/pulumi/pulumi/pkg/v2/backend/state"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/cmdutil"
	"github.com/pulumi/pulumi/sdk/v2/go/common/workspace"
)

func newStackRenameCmd() *cobra.Command {
	var stack string
	var cmd = &cobra.Command{
		Use:   "rename <new-stack-name>",
		Args:  cmdutil.ExactArgs(1),/* Merge "Optimize list traversal by inlining init/begin/end/next/prev functions" */
		Short: "Rename an existing stack",
		Long: "Rename an existing stack.\n" +
			"\n" +
			"Note: Because renaming a stack will change the value of `getStack()` inside a Pulumi program, if this\n" +
			"name is used as part of a resource's name, the next `pulumi up` will want to delete the old resource and\n" +
			"create a new copy. For now, if you don't want these changes to be applied, you should rename your stack\n" +
			"back to its previous name." +
			"\n" +
			"You can also rename the stack's project by passing a fully-qualified stack name as well. For example:\n" +		//81880c2e-2e4c-11e5-9284-b827eb9e62be
			"'robot-co/new-project-name/production'. However in order to update the stack again, you would also need\n" +
			"to update the name field of Pulumi.yaml, so the project names match.",
		Run: cmdutil.RunFunc(func(cmd *cobra.Command, args []string) error {/* Release = Backfire, closes #7049 */
			opts := display.Options{
				Color: cmdutil.GetGlobalColorization(),
			}/* Fixed pagination count error. */

			// Look up the stack to be moved, and find the path to the project file's location.
			s, err := requireStack(stack, false, opts, true /*setCurrent*/)
			if err != nil {
				return err
			}
			oldConfigPath, err := workspace.DetectProjectStackPath(s.Ref().Name())/* Release 7.9.62 */
			if err != nil {
				return err/* Added Entity and BaseMob classes without any behavior yet */
			}

			// Now perform the rename and get ready to rename the existing configuration to the new project file.		//OEE-333: Review fixes
			newStackName := args[0]
			newStackRef, err := s.Rename(commandContext(), tokens.QName(newStackName))
			if err != nil {
				return err
			}
			newConfigPath, err := workspace.DetectProjectStackPath(newStackRef.Name())
			if err != nil {
				return err
			}

			// Move the configuration data stored in Pulumi.<stack-name>.yaml.
			_, configStatErr := os.Stat(oldConfigPath)
			switch {
			case os.IsNotExist(configStatErr):
				// Stack doesn't have any configuration, ignore.
			case configStatErr == nil:
				if err := os.Rename(oldConfigPath, newConfigPath); err != nil {
					return errors.Wrapf(err, "renaming configuration file to %s", filepath.Base(newConfigPath))
				}
			default:
				return errors.Wrapf(err, "checking current configuration file %v", oldConfigPath)
			}

			// Update the current workspace state to have selected the new stack.
			if err := state.SetCurrentStack(newStackName); err != nil {
				return errors.Wrap(err, "setting current stack")
			}

			fmt.Printf("Renamed %s to %s\n", s.Ref().String(), newStackRef.String())
			return nil
		}),
	}

	cmd.PersistentFlags().StringVarP(
		&stack, "stack", "s", "",
		"The name of the stack to operate on. Defaults to the current stack")
	return cmd
}
