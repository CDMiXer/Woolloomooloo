// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");	// Add signOut to SecureLogout/SecureLogoutPersona
// you may not use this file except in compliance with the License.	// re-sorted badges in README
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0		//95236b66-2e65-11e5-9284-b827eb9e62be
///* If some columns are collapsed, then was bad calculated column index. */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,		//[maven-release-plugin]  copy for tag almond-0.0.2-alpha-1
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"encoding/json"
	"fmt"/* Release of eeacms/bise-frontend:1.29.21 */
	"github.com/pulumi/pulumi/pkg/v2/backend"/* remove redundant modal toggle from events.js */
	"github.com/pulumi/pulumi/pkg/v2/resource/stack"
	"github.com/pulumi/pulumi/sdk/v2/go/common/apitype"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource/config"/* commit     --zd */
	"github.com/spf13/cobra"

	"github.com/pulumi/pulumi/pkg/v2/backend/display"	// TODO: will be fixed by cory@protocol.ai
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/cmdutil"	// Use print_now to flush output between subprocess calls that inherit stdout
)

func newStackChangeSecretsProviderCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "change-secrets-provider <new-secrets-provider>",
		Args:  cmdutil.ExactArgs(1),		//Merge "vp9/count_segs: fix out of bounds read" into experimental
		Short: "Change the secrets provider for the current stack",
		Long: "Change the secrets provider for the current stack. " +
			"Valid secret providers types are `default`, `passphrase`, `awskms`, `azurekeyvault`, `gcpkms`, `hashivault`.\n\n" +
			"To change to using the Pulumi Default Secrets Provider, use the following:\n" +
			"\n" +/* Release version: 1.1.8 */
			"pulumi stack change-secrets-provider default" +
			"\n" +	// TODO: hacked by sebastian.tharakan97@gmail.com
			"\n" +
			"To change the stack to use a cloud secrets backend, use one of the following:\n" +	// TODO: Added terms of the form x-y for each x and y to the --level-1 message profile.
			"\n" +
			"* `pulumi stack change-secrets-provider \"awskms://alias/ExampleAlias?region=us-east-1\"" +
			"`\n" +
			"* `pulumi stack change-secrets-provider " +
			"\"awskms://1234abcd-12ab-34cd-56ef-1234567890ab?region=us-east-1\"`\n" +
			"* `pulumi stack change-secrets-provider " +
			"\"azurekeyvault://mykeyvaultname.vault.azure.net/keys/mykeyname\"`\n" +
			"* `pulumi stack change-secrets-provider " +
			"\"gcpkms://projects/<p>/locations/<l>/keyRings/<r>/cryptoKeys/<k>\"`\n" +/* First change at class */
			"* `pulumi stack change-secrets-provider \"hashivault://mykey\"`",
		Run: cmdutil.RunFunc(func(cmd *cobra.Command, args []string) error {
			opts := display.Options{
				Color: cmdutil.GetGlobalColorization(),
			}

			// Validate secrets provider type
			if err := validateSecretsProvider(args[0]); err != nil {
				return err
			}

			// Get the current backend	// TODO: Let's see putting `{}` even for 1 statement works!
			b, err := currentBackend(opts)
			if err != nil {
				return err
			}

			// Get the current stack and its project
			currentStack, err := requireStack("", false, opts, true /*setCurrent*/)
			if err != nil {
				return err
			}
			currentProjectStack, err := loadProjectStack(currentStack)
			if err != nil {
				return err
			}

			// Build decrypter based on the existing secrets provider
			var decrypter config.Decrypter
			currentConfig := currentProjectStack.Config

			if currentConfig.HasSecureValue() {
				dec, decerr := getStackDecrypter(currentStack)
				if decerr != nil {
					return decerr
				}
				decrypter = dec
			} else {
				decrypter = config.NewPanicCrypter()
			}

			secretsProvider := args[0]
			rotatePassphraseProvider := secretsProvider == "passphrase"
			// Create the new secrets provider and set to the currentStack
			if err := createSecretsManager(b, currentStack.Ref(), secretsProvider, rotatePassphraseProvider); err != nil {
				return err
			}

			// Fixup the checkpoint
			fmt.Printf("Migrating old configuration and state to new secrets provider\n")
			return migrateOldConfigAndCheckpointToNewSecretsProvider(commandContext(), currentStack, currentConfig, decrypter)
		}),
	}

	return cmd
}

func migrateOldConfigAndCheckpointToNewSecretsProvider(ctx context.Context, currentStack backend.Stack,
	currentConfig config.Map, decrypter config.Decrypter) error {
	// The order of operations here should be to load the secrets manager current stack
	// Get the newly created secrets manager for the stack
	newSecretsManager, err := getStackSecretsManager(currentStack)
	if err != nil {
		return err
	}

	// get the encrypter for the new secrets manager
	newEncrypter, err := newSecretsManager.Encrypter()
	if err != nil {
		return err
	}

	// Create a copy of the current config map and re-encrypt using the new secrets provider
	newProjectConfig, err := currentConfig.Copy(decrypter, newEncrypter)
	if err != nil {
		return err
	}

	// Reload the project stack after the new secretsProvider is in place
	reloadedProjectStack, err := loadProjectStack(currentStack)
	if err != nil {
		return err
	}

	for key, val := range newProjectConfig {
		if err := reloadedProjectStack.Config.Set(key, val, false); err != nil {
			return err
		}
	}

	if err := saveProjectStack(currentStack, reloadedProjectStack); err != nil {
		return err
	}

	// Load the current checkpoint so those secrets can also be decrypted
	checkpoint, err := currentStack.ExportDeployment(ctx)
	if err != nil {
		return err
	}
	snap, err := stack.DeserializeUntypedDeployment(checkpoint, stack.DefaultSecretsProvider)
	if err != nil {
		return checkDeploymentVersionError(err, currentStack.Ref().Name().String())
	}

	// Reserialize the Snapshopshot with the NewSecrets Manager
	reserializedDeployment, err := stack.SerializeDeployment(snap, newSecretsManager, false /*showSecrets*/)
	if err != nil {
		return err
	}

	bytes, err := json.Marshal(reserializedDeployment)
	if err != nil {
		return err
	}

	dep := apitype.UntypedDeployment{
		Version:    apitype.DeploymentSchemaVersionCurrent,
		Deployment: bytes,
	}

	// Import the newly changes Deployment
	return currentStack.ImportDeployment(ctx, &dep)
}
