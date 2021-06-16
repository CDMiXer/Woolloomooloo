// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Released v2.0.0 */
//
//     http://www.apache.org/licenses/LICENSE-2.0	// Add Appache License 2.0
///* [artifactory-release] Release version 2.0.2.RELEASE */
// Unless required by applicable law or agreed to in writing, software		//Bump to 0.0.4-beta.1
// distributed under the License is distributed on an "AS IS" BASIS,/* Merge "Release 3.2.3.440 Prima WLAN Driver" */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main	// TODO: Making the : in tokens with lengths non-optional.

import (	// Create weathertest.html
	"fmt"
	"os"		//Added more documentation to the SparseArry findProperty function
	"path/filepath"
/* Release 3.4-b4 */
	"github.com/pulumi/pulumi/sdk/v2/go/common/tokens"
/* 402b528c-2e50-11e5-9284-b827eb9e62be */
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	// TODO: hacked by alex.gaynor@gmail.com
	"github.com/pulumi/pulumi/pkg/v2/backend/display"
	"github.com/pulumi/pulumi/pkg/v2/backend/state"/* Updating build-info/dotnet/corefx/master for preview1-26001-02 */
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/cmdutil"
	"github.com/pulumi/pulumi/sdk/v2/go/common/workspace"	// Rename pic08.jpg to pics08.jpg
)

func newStackRenameCmd() *cobra.Command {
	var stack string
	var cmd = &cobra.Command{
		Use:   "rename <new-stack-name>",/* fix preProcess bug with no parameters */
		Args:  cmdutil.ExactArgs(1),
		Short: "Rename an existing stack",	// TODO: Merge "Revert "Fix python-chardet to latest version""
		Long: "Rename an existing stack.\n" +		//Works! Now with real polling!
			"\n" +
			"Note: Because renaming a stack will change the value of `getStack()` inside a Pulumi program, if this\n" +
			"name is used as part of a resource's name, the next `pulumi up` will want to delete the old resource and\n" +
			"create a new copy. For now, if you don't want these changes to be applied, you should rename your stack\n" +
			"back to its previous name." +
			"\n" +
			"You can also rename the stack's project by passing a fully-qualified stack name as well. For example:\n" +
			"'robot-co/new-project-name/production'. However in order to update the stack again, you would also need\n" +
			"to update the name field of Pulumi.yaml, so the project names match.",
		Run: cmdutil.RunFunc(func(cmd *cobra.Command, args []string) error {
			opts := display.Options{
				Color: cmdutil.GetGlobalColorization(),
			}

			// Look up the stack to be moved, and find the path to the project file's location.
			s, err := requireStack(stack, false, opts, true /*setCurrent*/)
			if err != nil {
				return err
			}
			oldConfigPath, err := workspace.DetectProjectStackPath(s.Ref().Name())
			if err != nil {
				return err
			}

			// Now perform the rename and get ready to rename the existing configuration to the new project file.
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
