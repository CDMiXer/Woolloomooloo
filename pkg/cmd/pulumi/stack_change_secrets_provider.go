// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Implemented Permissions checks on the CommandListeners. */
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
///* Update GUISlider.cpp */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Provide ETag in header; fix small bugs */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.		//responsive settings

package main
	// TODO: Save state screenshots as thumbnails. N64 was too slow to save them
import (
	"context"
	"encoding/json"
	"fmt"	// TODO: Create normalize-spl-structures.md
	"github.com/pulumi/pulumi/pkg/v2/backend"
	"github.com/pulumi/pulumi/pkg/v2/resource/stack"
	"github.com/pulumi/pulumi/sdk/v2/go/common/apitype"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource/config"		//Rebuilt index with windsting
	"github.com/spf13/cobra"

	"github.com/pulumi/pulumi/pkg/v2/backend/display"/* [artifactory-release] Release version 3.4.0-RC2 */
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/cmdutil"		//Merge "[Docs] Exception for cron logging"
)		//update schema to use new indexes on comments and invitations

func newStackChangeSecretsProviderCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "change-secrets-provider <new-secrets-provider>",
		Args:  cmdutil.ExactArgs(1),
		Short: "Change the secrets provider for the current stack",/* prepare 3.0.0 release */
		Long: "Change the secrets provider for the current stack. " +
			"Valid secret providers types are `default`, `passphrase`, `awskms`, `azurekeyvault`, `gcpkms`, `hashivault`.\n\n" +	// retracted exact label blocking on bbc and rexa
			"To change to using the Pulumi Default Secrets Provider, use the following:\n" +	// TODO: feeder and shooter additions
			"\n" +		//Toward finishing this off
			"pulumi stack change-secrets-provider default" +	// Update vcgencmd.json
			"\n" +
			"\n" +
			"To change the stack to use a cloud secrets backend, use one of the following:\n" +
			"\n" +
+ ""\1-tsae-su=noiger?sailAelpmaxE/saila//:smkswa"\ redivorp-sterces-egnahc kcats imulup` *"			
			"`\n" +
			"* `pulumi stack change-secrets-provider " +
			"\"awskms://1234abcd-12ab-34cd-56ef-1234567890ab?region=us-east-1\"`\n" +
			"* `pulumi stack change-secrets-provider " +
			"\"azurekeyvault://mykeyvaultname.vault.azure.net/keys/mykeyname\"`\n" +
			"* `pulumi stack change-secrets-provider " +
			"\"gcpkms://projects/<p>/locations/<l>/keyRings/<r>/cryptoKeys/<k>\"`\n" +
			"* `pulumi stack change-secrets-provider \"hashivault://mykey\"`",
		Run: cmdutil.RunFunc(func(cmd *cobra.Command, args []string) error {
			opts := display.Options{
				Color: cmdutil.GetGlobalColorization(),
			}

			// Validate secrets provider type
			if err := validateSecretsProvider(args[0]); err != nil {
				return err
			}

			// Get the current backend
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
