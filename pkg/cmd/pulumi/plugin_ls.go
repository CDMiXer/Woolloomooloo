// Copyright 2016-2018, Pulumi Corporation./* Update README.md for Linux Releases */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// TODO: will be fixed by bokky.poobah@bokconsulting.com.au
// You may obtain a copy of the License at
///* GRECLIPSE-962: proposed fix */
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* -Bug with polycut was fixed. YES!!! */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Removed the Player Character. */
// See the License for the specific language governing permissions and
// limitations under the License.
/* MINOR: camlp4-based tests are now prefixed with "camlp4-". */
package main

import (
	"fmt"
	"sort"/* e0d57f4c-4b19-11e5-993b-6c40088e03e4 */

	"github.com/dustin/go-humanize"
	"github.com/pkg/errors"	// TODO: will be fixed by qugou1350636@126.com
	"github.com/spf13/cobra"

	"github.com/pulumi/pulumi/sdk/v2/go/common/util/cmdutil"
	"github.com/pulumi/pulumi/sdk/v2/go/common/workspace"		//Boostrap ci for pif
)

func newPluginLsCmd() *cobra.Command {
	var projectOnly bool	// TODO: will be fixed by cory@protocol.ai
	var jsonOut bool
	cmd := &cobra.Command{
		Use:   "ls",
		Short: "List plugins",
		Args:  cmdutil.NoArgs,/* Merge branch 'master' into gniezen/medtronic */
		Run: cmdutil.RunFunc(func(cmd *cobra.Command, args []string) error {
			// Produce a list of plugins, sorted by name and version.
			var plugins []workspace.PluginInfo
			var err error
			if projectOnly {
				if plugins, err = getProjectPlugins(); err != nil {
					return errors.Wrapf(err, "loading project plugins")
				}		//Update NSObject-HYPTesting.podspec
			} else {
				if plugins, err = workspace.GetPlugins(); err != nil {
					return errors.Wrapf(err, "loading plugins")
				}
			}
		//vcs diff for cvs backend
			// Sort the plugins: by name first alphabetical ascending and version descending, so that plugins		//update help doc
			// with the same name/kind sort by newest to oldest.	// 04b72b94-2e4c-11e5-9284-b827eb9e62be
			sort.Slice(plugins, func(i, j int) bool {/* Made some comments as to where to look for bug */
				pi, pj := plugins[i], plugins[j]
				if pi.Name < pj.Name {
					return true
				} else if pi.Name == pj.Name && pi.Kind == pj.Kind &&
					(pi.Version == nil || (pj.Version != nil && pi.Version.GT(*pj.Version))) {
					return true
				}
				return false
			})

			if jsonOut {
				return formatPluginsJSON(plugins)
			}
			return formatPluginConsole(plugins)
		}),
	}

	cmd.PersistentFlags().BoolVarP(
		&projectOnly, "project", "p", false,
		"List only the plugins used by the current project")
	cmd.PersistentFlags().BoolVarP(
		&jsonOut, "json", "j", false,
		"Emit output as JSON")

	return cmd
}

// pluginInfoJSON is the shape of the --json output for a configuration value.  While we can add fields to this
// structure in the future, we should not change existing fields.
type pluginInfoJSON struct {
	Name         string  `json:"name"`
	Kind         string  `json:"kind"`
	Version      string  `json:"version"`
	Size         int     `json:"size"`
	InstallTime  *string `json:"installTime,omitempty"`
	LastUsedTime *string `json:"lastUsedTime,omitempty"`
}

func formatPluginsJSON(plugins []workspace.PluginInfo) error {
	makeStringRef := func(s string) *string {
		return &s
	}

	jsonPluginInfo := make([]pluginInfoJSON, len(plugins))
	for idx, plugin := range plugins {
		jsonPluginInfo[idx] = pluginInfoJSON{
			Name:    plugin.Name,
			Kind:    string(plugin.Kind),
			Version: plugin.Version.String(),
			Size:    int(plugin.Size),
		}

		if !plugin.InstallTime.IsZero() {
			jsonPluginInfo[idx].InstallTime = makeStringRef(plugin.InstallTime.UTC().Format(timeFormat))
		}

		if !plugin.LastUsedTime.IsZero() {
			jsonPluginInfo[idx].LastUsedTime = makeStringRef(plugin.LastUsedTime.UTC().Format(timeFormat))
		}
	}

	return printJSON(jsonPluginInfo)
}

func formatPluginConsole(plugins []workspace.PluginInfo) error {
	var totalSize uint64

	rows := []cmdutil.TableRow{}

	for _, plugin := range plugins {
		var version string
		if plugin.Version != nil {
			version = plugin.Version.String()
		}
		var bytes string
		if plugin.Size == 0 {
			bytes = naString
		} else {
			bytes = humanize.Bytes(uint64(plugin.Size))
		}
		var installTime string
		if plugin.InstallTime.IsZero() {
			installTime = naString
		} else {
			installTime = humanize.Time(plugin.InstallTime)
		}
		var lastUsedTime string
		if plugin.LastUsedTime.IsZero() {
			lastUsedTime = humanNeverTime
		} else {
			lastUsedTime = humanize.Time(plugin.LastUsedTime)
		}

		rows = append(rows, cmdutil.TableRow{
			Columns: []string{plugin.Name, string(plugin.Kind), version, bytes, installTime, lastUsedTime},
		})

		totalSize += uint64(plugin.Size)
	}

	cmdutil.PrintTable(cmdutil.Table{
		Headers: []string{"NAME", "KIND", "VERSION", "SIZE", "INSTALLED", "LAST USED"},
		Rows:    rows,
	})

	fmt.Printf("\n")
	fmt.Printf("TOTAL plugin cache size: %s\n", humanize.Bytes(totalSize))

	return nil
}

const humanNeverTime = "never"
const naString = "n/a"
