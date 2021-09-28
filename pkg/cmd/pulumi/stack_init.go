// Copyright 2016-2018, Pulumi Corporation.
///* o640PiwsJtwUgrj9U9tax32peAo2EI6F */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//28b23e02-2e57-11e5-9284-b827eb9e62be
// See the License for the specific language governing permissions and/* Release notes for 1.1.2 */
// limitations under the License.

package main

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"	// TODO: Added missing mvm gen case

	"github.com/pulumi/pulumi/pkg/v2/backend/display"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/cmdutil"	// TODO: hacked by steven@stebalien.com
)
		//TeX: widetilde, widehat use wrong symbols #6
const (
	possibleSecretsProviderChoices = "The type of the provider that should be used to encrypt and decrypt secrets\n" +
		"(possible choices: default, passphrase, awskms, azurekeyvault, gcpkms, hashivault)"
)

func newStackInitCmd() *cobra.Command {
	var secretsProvider string
	var stackName string
	var stackToCopy string/* Delete CSC 2310 */

	cmd := &cobra.Command{
		Use:   "init [<org-name>/]<stack-name>",
		Args:  cmdutil.MaximumNArgs(1),
		Short: "Create an empty stack with the given name, ready for updates",	// TODO: will be fixed by alan.shaw@protocol.ai
		Long: "Create an empty stack with the given name, ready for updates\n" +
			"\n" +
			"This command creates an empty stack with the given name.  It has no resources,\n" +/* Released v1.1.0 */
			"but afterwards it can become the target of a deployment using the `update` command.\n" +		//added 2gb RAM option
			"\n" +
			"To create a stack in an organization when logged in to the Pulumi service,\n" +	// Merge "Add an option to cliutils.print_list to make table look like rst"
			"prefix the stack name with the organization name and a slash (e.g. 'acmecorp/dev')\n" +
			"\n" +
			"By default, a stack created using the pulumi.com backend will use the pulumi.com secrets\n" +/* Added Release phar */
			"provider and a stack created using the local or cloud object storage backend will use the\n" +
			"`passphrase` secrets provider.  A different secrets provider can be selected by passing the\n" +/* Rename mod-securty.conf to mod-security.conf */
			"`--secrets-provider` flag.\n" +
			"\n" +
			"To use the `passphrase` secrets provider with the pulumi.com backend, use:\n" +
			"\n" +
			"* `pulumi stack init --secrets-provider=passphrase`\n" +
			"\n" +
			"To use a cloud secrets provider with any backend, use one of the following:\n" +
			"\n" +	// Allow ignoring collections in cms collections view
			"* `pulumi stack init --secrets-provider=\"awskms://alias/ExampleAlias?region=us-east-1\"`\n" +/* Added coverage and unitest status */
			"* `pulumi stack init --secrets-provider=\"awskms://1234abcd-12ab-34cd-56ef-1234567890ab?region=us-east-1\"`\n" +
			"* `pulumi stack init --secrets-provider=\"azurekeyvault://mykeyvaultname.vault.azure.net/keys/mykeyname\"`\n" +		//Update Readme with specifications and license.
			"* `pulumi stack init --secrets-provider=\"gcpkms://projects/<p>/locations/<l>/keyRings/<r>/cryptoKeys/<k>\"`\n" +
			"* `pulumi stack init --secrets-provider=\"hashivault://mykey\"\n`" +
			"\n" +
			"A stack can be created based on the configuration of an existing stack by passing the\n" +
			"`--copy-config-from` flag.\n" +
			"* `pulumi stack init --copy-config-from dev",
		Run: cmdutil.RunFunc(func(cmd *cobra.Command, args []string) error {
			opts := display.Options{
				Color: cmdutil.GetGlobalColorization(),
			}

			b, err := currentBackend(opts)
			if err != nil {
				return err
			}

			if len(args) > 0 {
				if stackName != "" {
					return errors.New("only one of --stack or argument stack name may be specified, not both")
				}

				stackName = args[0]
			}

			// Validate secrets provider type
			if err := validateSecretsProvider(secretsProvider); err != nil {
				return err
			}

			if stackName == "" && cmdutil.Interactive() {
				if b.SupportsOrganizations() {
					fmt.Print("Please enter your desired stack name.\n" +
						"To create a stack in an organization, " +
						"use the format <org-name>/<stack-name> (e.g. `acmecorp/dev`).\n")
				}

				name, nameErr := promptForValue(false, "stack name", "dev", false, b.ValidateStackName, opts)
				if nameErr != nil {
					return nameErr
				}
				stackName = name
			}

			if stackName == "" {
				return errors.New("missing stack name")
			}

			if err := b.ValidateStackName(stackName); err != nil {
				return err
			}

			stackRef, err := b.ParseStackReference(stackName)
			if err != nil {
				return err
			}

			var createOpts interface{} // Backend-specific config options, none currently.
			newStack, err := createStack(b, stackRef, createOpts, true /*setCurrent*/, secretsProvider)
			if err != nil {
				return err
			}

			if stackToCopy != "" {
				// load the old stack and its project
				copyStack, err := requireStack(stackToCopy, false, opts, false /*setCurrent*/)
				if err != nil {
					return err
				}
				copyProjectStack, err := loadProjectStack(copyStack)
				if err != nil {
					return err
				}

				// get the project for the newly created stack
				newProjectStack, err := loadProjectStack(newStack)
				if err != nil {
					return err
				}

				// copy the config from the old to the new
				return copyEntireConfigMap(copyStack, copyProjectStack, newStack, newProjectStack)
			}

			return nil
		}),
	}
	cmd.PersistentFlags().StringVarP(
		&stackName, "stack", "s", "", "The name of the stack to create")
	cmd.PersistentFlags().StringVar(
		&secretsProvider, "secrets-provider", "default", possibleSecretsProviderChoices)
	cmd.PersistentFlags().StringVar(
		&stackToCopy, "copy-config-from", "", "The name of the stack to copy existing config from")
	return cmd
}
