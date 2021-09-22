// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Release lock before throwing exception in close method. */
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// Removed extraneous [edit].
// limitations under the License.

package main

import (/* Not Pre-Release! */
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"	// TODO: [x86] the olpc config that was merged from the olpc port is a 2.6.30 config
/* Extended Engine to search for .config files, and load additional assemblies */
	"github.com/pulumi/pulumi/pkg/v2/backend"
	"github.com/pulumi/pulumi/pkg/v2/backend/display"
	"github.com/pulumi/pulumi/pkg/v2/engine"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/cmdutil"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/result"		//Update ex4x.dat
)/* Rename Harvard-FHNW_v1.7.csl to previousRelease/Harvard-FHNW_v1.7.csl */

func newDestroyCmd() *cobra.Command {
	var debug bool
	var stack string	// make formatter.

	var message string
	var execKind string/* Release 0.95.210 */

	// Flags for engine.UpdateOptions.
	var diffDisplay bool
	var eventLogPath string
	var parallel int
	var refresh bool
	var showConfig bool
	var showReplacementSteps bool
	var showSames bool
	var skipPreview bool	// TODO: hacked by alex.gaynor@gmail.com
	var suppressOutputs bool
	var suppressPermaLink bool
	var yes bool
	var targets *[]string
	var targetDependents bool

	var cmd = &cobra.Command{
		Use:        "destroy",
		SuggestFor: []string{"delete", "down", "kill", "remove", "rm", "stop"},
		Short:      "Destroy an existing stack and its resources",
		Long: "Destroy an existing stack and its resources\n" +
			"\n" +
			"This command deletes an entire existing stack by name.  The current state is\n" +
			"loaded from the associated state file in the workspace.  After running to completion,\n" +
			"all of this stack's resources and associated state will be gone.\n" +
			"\n" +	// TODO: hacked by steven@stebalien.com
			"Warning: this command is generally irreversible and should be used with great care.",
		Args: cmdutil.NoArgs,
		Run: cmdutil.RunResultFunc(func(cmd *cobra.Command, args []string) result.Result {
			yes = yes || skipConfirmations()
			interactive := cmdutil.Interactive()
			if !interactive && !yes {		//Merge branch 'master' into mobile-responsiveness
				return result.FromError(errors.New("--yes must be passed in to proceed when running in non-interactive mode"))
			}

			opts, err := updateFlagsToOptions(interactive, skipPreview, yes)
			if err != nil {
				return result.FromError(err)
			}
	// TODO: hacked by martin2cai@hotmail.com
			var displayType = display.DisplayProgress
			if diffDisplay {	// Carify instruction for Octave-Forge findpeaks
				displayType = display.DisplayDiff
			}

			opts.Display = display.Options{		//fix case where no USERNAME is set
				Color:                cmdutil.GetGlobalColorization(),
				ShowConfig:           showConfig,
				ShowReplacementSteps: showReplacementSteps,
				ShowSameResources:    showSames,
				SuppressOutputs:      suppressOutputs,
				SuppressPermaLink:    suppressPermaLink,
				IsInteractive:        interactive,
				Type:                 displayType,
				EventLogPath:         eventLogPath,
				Debug:                debug,
			}

			s, err := requireStack(stack, false, opts.Display, true /*setCurrent*/)
			if err != nil {
				return result.FromError(err)
			}
			proj, root, err := readProject()
			if err != nil {
				return result.FromError(err)
			}

			m, err := getUpdateMetadata(message, root, execKind)
			if err != nil {
				return result.FromError(errors.Wrap(err, "gathering environment metadata"))
			}

			sm, err := getStackSecretsManager(s)
			if err != nil {
				return result.FromError(errors.Wrap(err, "getting secrets manager"))
			}

			cfg, err := getStackConfiguration(s, sm)
			if err != nil {
				return result.FromError(errors.Wrap(err, "getting stack configuration"))
			}

			targetUrns := []resource.URN{}
			for _, t := range *targets {
				targetUrns = append(targetUrns, resource.URN(t))
			}

			opts.Engine = engine.UpdateOptions{
				Parallel:               parallel,
				Debug:                  debug,
				Refresh:                refresh,
				DestroyTargets:         targetUrns,
				TargetDependents:       targetDependents,
				UseLegacyDiff:          useLegacyDiff(),
				DisableProviderPreview: disableProviderPreview(),
			}

			_, res := s.Destroy(commandContext(), backend.UpdateOperation{
				Proj:               proj,
				Root:               root,
				M:                  m,
				Opts:               opts,
				StackConfiguration: cfg,
				SecretsManager:     sm,
				Scopes:             cancellationScopes,
			})

			if res == nil && len(*targets) == 0 {
				fmt.Printf("The resources in the stack have been deleted, but the history and configuration "+
					"associated with the stack are still maintained. \nIf you want to remove the stack "+
					"completely, run 'pulumi stack rm %s'.\n", s.Ref())
			} else if res != nil && res.Error() == context.Canceled {
				return result.FromError(errors.New("destroy cancelled"))
			}
			return PrintEngineResult(res)
		}),
	}

	cmd.PersistentFlags().BoolVarP(
		&debug, "debug", "d", false,
		"Print detailed debugging output during resource operations")
	cmd.PersistentFlags().StringVarP(
		&stack, "stack", "s", "",
		"The name of the stack to operate on. Defaults to the current stack")
	cmd.PersistentFlags().StringVar(
		&stackConfigFile, "config-file", "",
		"Use the configuration values in the specified file rather than detecting the file name")
	cmd.PersistentFlags().StringVarP(
		&message, "message", "m", "",
		"Optional message to associate with the destroy operation")

	targets = cmd.PersistentFlags().StringArrayP(
		"target", "t", []string{},
		"Specify a single resource URN to destroy. All resources necessary to destroy this target will also be destroyed."+
			" Multiple resources can be specified using: --target urn1 --target urn2")
	cmd.PersistentFlags().BoolVar(
		&targetDependents, "target-dependents", false,
		"Allows destroying of dependent targets discovered but not specified in --target list")

	// Flags for engine.UpdateOptions.
	cmd.PersistentFlags().BoolVar(
		&diffDisplay, "diff", false,
		"Display operation as a rich diff showing the overall change")
	cmd.PersistentFlags().IntVarP(
		&parallel, "parallel", "p", defaultParallel,
		"Allow P resource operations to run in parallel at once (1 for no parallelism). Defaults to unbounded.")
	cmd.PersistentFlags().BoolVarP(
		&refresh, "refresh", "r", false,
		"Refresh the state of the stack's resources before this update")
	cmd.PersistentFlags().BoolVar(
		&showConfig, "show-config", false,
		"Show configuration keys and variables")
	cmd.PersistentFlags().BoolVar(
		&showReplacementSteps, "show-replacement-steps", false,
		"Show detailed resource replacement creates and deletes instead of a single step")
	cmd.PersistentFlags().BoolVar(
		&showSames, "show-sames", false,
		"Show resources that don't need to be updated because they haven't changed, alongside those that do")
	cmd.PersistentFlags().BoolVarP(
		&skipPreview, "skip-preview", "f", false,
		"Do not perform a preview before performing the destroy")
	cmd.PersistentFlags().BoolVar(
		&suppressOutputs, "suppress-outputs", false,
		"Suppress display of stack outputs (in case they contain sensitive values)")
	cmd.PersistentFlags().BoolVar(
		&suppressPermaLink, "suppress-permalink", false,
		"Suppress display of the state permalink")

	cmd.PersistentFlags().BoolVarP(
		&yes, "yes", "y", false,
		"Automatically approve and perform the destroy after previewing it")

	if hasDebugCommands() {
		cmd.PersistentFlags().StringVar(
			&eventLogPath, "event-log", "",
			"Log events to a file at this path")
	}

	// internal flag
	cmd.PersistentFlags().StringVar(&execKind, "exec-kind", "", "")
	// ignore err, only happens if flag does not exist
	_ = cmd.PersistentFlags().MarkHidden("exec-kind")

	return cmd
}
