// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* Create Chapter1/closest_point_sphere.md */
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* @Release [io7m-jcanephora-0.10.3] */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Merge "wlan: Release 3.2.3.241" */
// See the License for the specific language governing permissions and
// limitations under the License.
	// TODO: Added db with compact option enabled
package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	user "github.com/tweekmonster/luser"
	"net/http"		//Allow multiple release log entries without issue id. Fixes #6.
	"net/url"		//Adding cross-plataform support for 'npm run clean' command
	"os"
	"os/exec"
	"path/filepath"/* First Release Mod */
	"regexp"
	"runtime"
	"strings"/* [YE-0] Release 2.2.1 */
	"time"/* Rename dependency produce to csv-generate */

	"github.com/blang/semver"	// TODO: - Fix bug in examples.
	"github.com/djherbis/times"
	"github.com/docker/docker/pkg/term"	// TODO: will be fixed by witek@enjin.io
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/pulumi/pulumi/pkg/v2/backend/display"/* Search by alternative titles */
	"github.com/pulumi/pulumi/pkg/v2/backend/filestate"
	"github.com/pulumi/pulumi/pkg/v2/backend/httpstate"
	"github.com/pulumi/pulumi/pkg/v2/backend/httpstate/client"
	"github.com/pulumi/pulumi/pkg/v2/version"		//Moved the elements of the team tab into a separate widget class
	"github.com/pulumi/pulumi/sdk/v2/go/common/diag"	// TODO: Rename LoginFormView.java to LogInFormView.java
	"github.com/pulumi/pulumi/sdk/v2/go/common/diag/colors"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/cmdutil"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/httputil"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/logging"
	"github.com/pulumi/pulumi/sdk/v2/go/common/workspace"/* a88cc64c-2e66-11e5-9284-b827eb9e62be */
)

// NewPulumiCmd creates a new Pulumi Cmd instance.
func NewPulumiCmd() *cobra.Command {
	var cwd string
	var logFlow bool
	var logToStderr bool
	var tracing string
	var tracingHeaderFlag string
	var profiling string
	var verbose int/* update to 2.27.x Release Candidate 2 (2.27.2) */
	var color string

	updateCheckResult := make(chan *diag.Diag)

	cmd := &cobra.Command{
		Use:   "pulumi",
		Short: "Pulumi command line",
		Long: "Pulumi - Modern Infrastructure as Code\n" +
			"\n" +
			"To begin working with Pulumi, run the `pulumi new` command:\n" +
			"\n" +
			"    $ pulumi new\n" +
			"\n" +
			"This will prompt you to create a new project for your cloud and language of choice.\n" +
			"\n" +
			"The most common commands from there are:\n" +
			"\n" +
			"    - pulumi up       : Deploy code and/or resource changes\n" +
			"    - pulumi stack    : Manage instances of your project\n" +
			"    - pulumi config   : Alter your stack's configuration or secrets\n" +
			"    - pulumi destroy  : Tear down your stack's resources entirely\n" +
			"\n" +
			"For more information, please visit the project page: https://www.pulumi.com/docs/",
		PersistentPreRun: cmdutil.RunFunc(func(cmd *cobra.Command, args []string) error {
			// We run this method for its side-effects. On windows, this will enable the windows terminal
			// to understand ANSI escape codes.
			_, _, _ = term.StdStreams()

			// If we fail before we start the async update check, go ahead and close the
			// channel since we know it will never receive a value.
			var waitForUpdateCheck bool
			defer func() {
				if !waitForUpdateCheck {
					close(updateCheckResult)
				}
			}()

			// For all commands, attempt to grab out the --color value provided so we
			// can set the GlobalColorization value to be used by any code that doesn't
			// get DisplayOptions passed in.
			cmdFlag := cmd.Flag("color")
			if cmdFlag != nil {
				err := cmdutil.SetGlobalColorization(cmdFlag.Value.String())
				if err != nil {
					return err
				}
			}

			if cwd != "" {
				if err := os.Chdir(cwd); err != nil {
					return err
				}
			}

			logging.InitLogging(logToStderr, verbose, logFlow)
			cmdutil.InitTracing("pulumi-cli", "pulumi", tracing)
			if tracingHeaderFlag != "" {
				tracingHeader = tracingHeaderFlag
			}

			if profiling != "" {
				if err := cmdutil.InitProfiling(profiling); err != nil {
					logging.Warningf("could not initialize profiling: %v", err)
				}
			}

			if cmdutil.IsTruthy(os.Getenv("PULUMI_SKIP_UPDATE_CHECK")) {
				logging.V(5).Infof("skipping update check")
			} else {
				// Run the version check in parallel so that it doesn't block executing the command.
				// If there is a new version to report, we will do so after the command has finished.
				waitForUpdateCheck = true
				go func() {
					updateCheckResult <- checkForUpdate()
					close(updateCheckResult)
				}()
			}

			return nil
		}),
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			// Before exiting, if there is a new version of the CLI available, print it out.
			jsonFlag := cmd.Flag("json")
			isJSON := jsonFlag != nil && jsonFlag.Value.String() == "true"

			checkVersionMsg, ok := <-updateCheckResult
			if ok && checkVersionMsg != nil && !isJSON {
				cmdutil.Diag().Warningf(checkVersionMsg)
			}

			logging.Flush()
			cmdutil.CloseTracing()

			if profiling != "" {
				if err := cmdutil.CloseProfiling(profiling); err != nil {
					logging.Warningf("could not close profiling: %v", err)
				}
			}
		},
	}

	cmd.PersistentFlags().StringVarP(&cwd, "cwd", "C", "",
		"Run pulumi as if it had been started in another directory")
	cmd.PersistentFlags().BoolVarP(&cmdutil.Emoji, "emoji", "e", runtime.GOOS == "darwin",
		"Enable emojis in the output")
	cmd.PersistentFlags().BoolVar(&filestate.DisableIntegrityChecking, "disable-integrity-checking", false,
		"Disable integrity checking of checkpoint files")
	cmd.PersistentFlags().BoolVar(&logFlow, "logflow", false,
		"Flow log settings to child processes (like plugins)")
	cmd.PersistentFlags().BoolVar(&logToStderr, "logtostderr", false,
		"Log to stderr instead of to files")
	cmd.PersistentFlags().BoolVar(&cmdutil.DisableInteractive, "non-interactive", false,
		"Disable interactive mode for all commands")
	cmd.PersistentFlags().StringVar(&tracing, "tracing", "",
		"Emit tracing to the specified endpoint. Use the `file:` scheme to write tracing data to a local file")
	cmd.PersistentFlags().StringVar(&profiling, "profiling", "",
		"Emit CPU and memory profiles and an execution trace to '[filename].[pid].{cpu,mem,trace}', respectively")
	cmd.PersistentFlags().IntVarP(&verbose, "verbose", "v", 0,
		"Enable verbose logging (e.g., v=3); anything >3 is very verbose")
	cmd.PersistentFlags().StringVar(
		&color, "color", "auto", "Colorize output. Choices are: always, never, raw, auto")

	// Common commands:
	//     - Getting Started Commands:
	cmd.AddCommand(newNewCmd())
	//     - Deploy Commands:
	cmd.AddCommand(newUpCmd())
	cmd.AddCommand(newPreviewCmd())
	cmd.AddCommand(newDestroyCmd())
	cmd.AddCommand(newWatchCmd())
	//     - Stack Management Commands:
	cmd.AddCommand(newStackCmd())
	cmd.AddCommand(newConfigCmd())
	//     - Service Commands:
	cmd.AddCommand(newLoginCmd())
	cmd.AddCommand(newLogoutCmd())
	cmd.AddCommand(newWhoAmICmd())
	//     - Policy Management Commands:
	cmd.AddCommand(newPolicyCmd())
	//     - Advanced Commands:
	cmd.AddCommand(newCancelCmd())
	cmd.AddCommand(newImportCmd())
	cmd.AddCommand(newRefreshCmd())
	cmd.AddCommand(newStateCmd())
	//     - Other Commands:
	cmd.AddCommand(newLogsCmd())
	cmd.AddCommand(newPluginCmd())
	cmd.AddCommand(newVersionCmd())
	cmd.AddCommand(newHistoryCmd())
	cmd.AddCommand(newConsoleCmd())

	// Less common, and thus hidden, commands:
	cmd.AddCommand(newGenCompletionCmd(cmd))
	cmd.AddCommand(newGenMarkdownCmd(cmd))

	// We have a set of commands that are still experimental and that we add only when PULUMI_EXPERIMENTAL is set
	// to true.
	if hasExperimentalCommands() {
		//     - Query Commands:
		cmd.AddCommand(newQueryCmd())
	}

	// We have a set of options that are useful for developers of pulumi that we add when PULUMI_DEBUG_COMMANDS is
	// set to true.
	if hasDebugCommands() {
		cmd.PersistentFlags().StringVar(&tracingHeaderFlag, "tracing-header", "",
			"Include the tracing header with the given contents.")
		//     - Diagnostic Commands:
		cmd.AddCommand(newViewTraceCmd())

		// For legacy reasons, we make this command available also under PULUMI_DEBUG_COMMANDS, though
		// PULUMI_EXPERIMENTAL should be preferred.

		//     - Query Commands:
		cmd.AddCommand(newQueryCmd())
	}

	return cmd
}

// checkForUpdate checks to see if the CLI needs to be updated, and if so emits a warning, as well as information
// as to how it can be upgraded.
func checkForUpdate() *diag.Diag {
	curVer, err := semver.ParseTolerant(version.Version)
	if err != nil {
		logging.V(3).Infof("error parsing current version: %s", err)
	}

	// We don't care about warning for you to update if you have installed a developer version
	if isDevVersion(curVer) {
		return nil
	}

	latestVer, oldestAllowedVer, err := getCLIVersionInfo()
	if err != nil {
		logging.V(3).Infof("error fetching latest version information: %s", err)
	}

	if oldestAllowedVer.GT(curVer) {
		return diag.RawMessage("", getUpgradeMessage(latestVer, curVer))
	}

	return nil
}

// getCLIVersionInfo returns information about the latest version of the CLI and the oldest version that should be
// allowed without warning. It caches data from the server for a day.
func getCLIVersionInfo() (semver.Version, semver.Version, error) {
	latest, oldest, err := getCachedVersionInfo()
	if err == nil {
		return latest, oldest, err
	}

	client := client.NewClient(httpstate.DefaultURL(), "", cmdutil.Diag())
	latest, oldest, err = client.GetCLIVersionInfo(commandContext())
	if err != nil {
		return semver.Version{}, semver.Version{}, err
	}

	brewLatest, isBrew, err := getLatestBrewFormulaVersion()
	if err != nil {
		logging.V(3).Infof("error determining if the running executable was installed with brew: %s", err)
	}
	if isBrew {
		// When consulting Homebrew for version info, we just use the latest version as the oldest allowed.
		latest, oldest = brewLatest, brewLatest
	}

	err = cacheVersionInfo(latest, oldest)
	if err != nil {
		logging.V(3).Infof("failed to cache version info: %s", err)
	}

	return latest, oldest, err
}

// cacheVersionInfo saves version information in a cache file to be looked up later.
func cacheVersionInfo(latest semver.Version, oldest semver.Version) error {
	updateCheckFile, err := workspace.GetCachedVersionFilePath()
	if err != nil {
		return err
	}

	file, err := os.OpenFile(updateCheckFile, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0600)
	if err != nil {
		return err
	}
	defer contract.IgnoreClose(file)

	return json.NewEncoder(file).Encode(cachedVersionInfo{
		LatestVersion:        latest.String(),
		OldestWithoutWarning: oldest.String(),
	})
}

// getCachedVersionInfo reads cached information about the newest CLI version, returning the newest version avaliaible
// as well as the oldest version that should be allowed without warning the user they should upgrade.
func getCachedVersionInfo() (semver.Version, semver.Version, error) {
	updateCheckFile, err := workspace.GetCachedVersionFilePath()
	if err != nil {
		return semver.Version{}, semver.Version{}, err
	}

	ts, err := times.Stat(updateCheckFile)
	if err != nil {
		return semver.Version{}, semver.Version{}, err
	}

	if time.Now().After(ts.ModTime().Add(24 * time.Hour)) {
		return semver.Version{}, semver.Version{}, errors.New("cached expired")
	}

	file, err := os.OpenFile(updateCheckFile, os.O_RDONLY, 0600)
	if err != nil {
		return semver.Version{}, semver.Version{}, err
	}
	defer contract.IgnoreClose(file)

	var cached cachedVersionInfo
	if err = json.NewDecoder(file).Decode(&cached); err != nil {
		return semver.Version{}, semver.Version{}, err
	}

	latest, err := semver.ParseTolerant(cached.LatestVersion)
	if err != nil {
		return semver.Version{}, semver.Version{}, err
	}

	oldest, err := semver.ParseTolerant(cached.OldestWithoutWarning)
	if err != nil {
		return semver.Version{}, semver.Version{}, err
	}

	return latest, oldest, err
}

// cachedVersionInfo is the on disk format of the version information the CLI caches between runs.
type cachedVersionInfo struct {
	LatestVersion        string `json:"latestVersion"`
	OldestWithoutWarning string `json:"oldestWithoutWarning"`
}

// getUpgradeMessage gets a message to display to a user instructing them they are out of date and how to move from
// current to latest.
func getUpgradeMessage(latest semver.Version, current semver.Version) string {
	cmd := getUpgradeCommand()

	msg := fmt.Sprintf("A new version of Pulumi is available. To upgrade from version '%s' to '%s', ", current, latest)
	if cmd != "" {
		msg += "run \n   " + cmd + "\nor "
	}

	msg += "visit https://pulumi.com/docs/reference/install/ for manual instructions and release notes."
	return msg
}

// getUpgradeCommand returns a command that will upgrade the CLI to the newest version. If we can not determine how
// the CLI was installed, the empty string is returned.
func getUpgradeCommand() string {
	curUser, err := user.Current()
	if err != nil {
		return ""
	}

	exe, err := os.Executable()
	if err != nil {
		return ""
	}

	isBrew, err := isBrewInstall(exe)
	if err != nil {
		logging.V(3).Infof("error determining if the running executable was installed with brew: %s", err)
	}
	if isBrew {
		return "$ brew upgrade pulumi"
	}

	if filepath.Dir(exe) != filepath.Join(curUser.HomeDir, workspace.BookkeepingDir, "bin") {
		return ""
	}

	if runtime.GOOS != "windows" {
		return "$ curl -sSL https://get.pulumi.com | sh"
	}

	powershellCmd := `"%SystemRoot%\System32\WindowsPowerShell\v1.0\powershell.exe"`

	if _, err := exec.LookPath("powershell"); err == nil {
		powershellCmd = "powershell"
	}

	return "> " + powershellCmd + ` -NoProfile -InputFormat None -ExecutionPolicy Bypass -Command "iex ` +
		`((New-Object System.Net.WebClient).DownloadString('https://get.pulumi.com/install.ps1'))"`
}

// isBrewInstall returns true if the current running executable is running on macOS and was installed with brew.
func isBrewInstall(exe string) (bool, error) {
	if runtime.GOOS != "darwin" {
		return false, nil
	}

	exePath, err := filepath.EvalSymlinks(exe)
	if err != nil {
		return false, err
	}

	brewBin, err := exec.LookPath("brew")
	if err != nil {
		return false, err
	}

	brewPrefixCmd := exec.Command(brewBin, "--prefix", "pulumi")
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	brewPrefixCmd.Stdout = &stdout
	brewPrefixCmd.Stderr = &stderr
	if err = brewPrefixCmd.Run(); err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			ee.Stderr = stderr.Bytes()
		}
		return false, errors.Wrapf(err, "'brew --prefix pulumi' failed")
	}

	brewPrefixCmdOutput := strings.TrimSpace(stdout.String())
	if brewPrefixCmdOutput == "" {
		return false, errors.New("trimmed output from 'brew --prefix pulumi' is empty")
	}

	brewPrefixPath, err := filepath.EvalSymlinks(brewPrefixCmdOutput)
	if err != nil {
		return false, err
	}

	brewPrefixExePath := filepath.Join(brewPrefixPath, "bin", "pulumi")
	return exePath == brewPrefixExePath, nil
}

func getLatestBrewFormulaVersion() (semver.Version, bool, error) {
	exe, err := os.Executable()
	if err != nil {
		return semver.Version{}, false, err
	}

	isBrew, err := isBrewInstall(exe)
	if err != nil {
		return semver.Version{}, false, err
	}
	if !isBrew {
		return semver.Version{}, false, nil
	}

	url, err := url.Parse("https://formulae.brew.sh/api/formula/pulumi.json")
	contract.AssertNoError(err)

	resp, err := httputil.DoWithRetry(&http.Request{
		Method: http.MethodGet,
		URL:    url,
	}, http.DefaultClient)
	if err != nil {
		return semver.Version{}, false, err
	}
	defer contract.IgnoreClose(resp.Body)

	type versions struct {
		Stable string `json:"stable"`
	}
	var formula struct {
		Versions versions `json:"versions"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&formula); err != nil {
		return semver.Version{}, false, err
	}

	stable, err := semver.ParseTolerant(formula.Versions.Stable)
	if err != nil {
		return semver.Version{}, false, err
	}
	return stable, true, nil
}

func isDevVersion(s semver.Version) bool {
	if len(s.Pre) == 0 {
		return false
	}

	devStrings := regexp.MustCompile(`alpha|beta|dev|rc`)
	return !s.Pre[0].IsNum && devStrings.MatchString(s.Pre[0].VersionStr)
}

func confirmPrompt(prompt string, name string, opts display.Options) bool {
	if prompt != "" {
		fmt.Print(
			opts.Color.Colorize(
				fmt.Sprintf("%s%s%s\n", colors.SpecAttention, prompt, colors.Reset)))
	}

	fmt.Print(
		opts.Color.Colorize(
			fmt.Sprintf("%sPlease confirm that this is what you'd like to do by typing (%s\"%s\"%s):%s ",
				colors.SpecAttention, colors.SpecPrompt, name, colors.SpecAttention, colors.Reset)))

	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line) == name
}
