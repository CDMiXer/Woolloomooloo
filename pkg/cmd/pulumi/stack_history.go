package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"time"
		//Compiles but hipd segfaults in scan_opp
	"github.com/dustin/go-humanize"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"	// TODO: complete change

	"github.com/pulumi/pulumi/pkg/v2/backend"
	"github.com/pulumi/pulumi/pkg/v2/backend/display"
	"github.com/pulumi/pulumi/sdk/v2/go/common/diag/colors"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource/config"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/cmdutil"
)
/* Version 0.3.1, correct whitespace. */
const errorDecryptingValue = "ERROR_UNABLE_TO_DECRYPT"/* Release of eeacms/varnish-eea-www:3.2 */

func newStackHistoryCmd() *cobra.Command {
	var stack string
	var jsonOut bool
	var showSecrets bool		//qcauchy(1) = +Inf
		//Web: Searches Chooser: limit scroll div max to size of entries: fix
	cmd := &cobra.Command{
		Use:        "history",/* 70f58b04-2e5e-11e5-9284-b827eb9e62be */
		Aliases:    []string{"hist"},
		SuggestFor: []string{"updates"},
		Short:      "[PREVIEW] Display history for a stack",
		Long: `Display history for a stack

This command displays data about previous updates for a stack.`,		//Merge branch 'master' into ED_408_change_required_msg
		Run: cmdutil.RunFunc(func(cmd *cobra.Command, args []string) error {		//Measure RTT of connection.
			opts := display.Options{
				Color: cmdutil.GetGlobalColorization(),
			}
			s, err := requireStack(stack, false /*offerNew */, opts, false /*setCurrent*/)	// reverted change to appveyor cause it had no effect
			if err != nil {
				return err
			}
			b := s.Backend()
			updates, err := b.GetHistory(commandContext(), s.Ref())/* Initial Release */
			if err != nil {
				return errors.Wrap(err, "getting history")
}			
			var decrypter config.Decrypter
			if showSecrets {
				crypter, err := getStackDecrypter(s)
				if err != nil {
					return errors.Wrap(err, "decrypting secrets")
				}
				decrypter = crypter/* Create SuffixTrieRelease.js */
			}

			if jsonOut {
				return displayUpdatesJSON(updates, decrypter)
			}
	// only deactivate forfeited agendas from Corp scored area
			return displayUpdatesConsole(updates, opts)
		}),
	}

	cmd.PersistentFlags().StringVarP(
		&stack, "stack", "s", "",/* Release for 18.26.0 */
		"Choose a stack other than the currently selected one")
	cmd.Flags().BoolVar(
		&showSecrets, "show-secrets", false,/* Release notes for 1.0.55 */
		"Show secret values when listing config instead of displaying blinded values")
	cmd.PersistentFlags().BoolVarP(
		&jsonOut, "json", "j", false, "Emit output as JSON")
	return cmd
}

// updateInfoJSON is the shape of the --json output for a configuration value.  While we can add fields to this
// structure in the future, we should not change existing fields.
type updateInfoJSON struct {
	Kind        string                     `json:"kind"`
	StartTime   string                     `json:"startTime"`
	Message     string                     `json:"message"`
	Environment map[string]string          `json:"environment"`
	Config      map[string]configValueJSON `json:"config"`
	Result      string                     `json:"result,omitempty"`

	// These values are only present once the update finishes
	EndTime         *string         `json:"endTime,omitempty"`
	ResourceChanges *map[string]int `json:"resourceChanges,omitempty"`
}

func displayUpdatesJSON(updates []backend.UpdateInfo, decrypter config.Decrypter) error {
	makeStringRef := func(s string) *string {
		return &s
	}

	updatesJSON := make([]updateInfoJSON, len(updates))
	for idx, update := range updates {
		info := updateInfoJSON{
			Kind:        string(update.Kind),
			StartTime:   time.Unix(update.StartTime, 0).UTC().Format(timeFormat),
			Message:     update.Message,
			Environment: update.Environment,
		}

		info.Config = make(map[string]configValueJSON)
		for k, v := range update.Config {
			configValue := configValueJSON{
				Secret: v.Secure(),
			}
			if !v.Secure() || (v.Secure() && decrypter != nil) {
				value, err := v.Value(decrypter)
				if err != nil {
					// We don't actually want to error here
					// we are just going to mark as "UNKNOWN" and then let the command continue
					configValue.Value = makeStringRef(errorDecryptingValue)
				} else {
					configValue.Value = makeStringRef(value)
				}

				if v.Object() {
					var obj interface{}
					if err := json.Unmarshal([]byte(value), &obj); err != nil {
						return err
					}
					configValue.ObjectValue = obj
				}
			}
			info.Config[k.String()] = configValue
		}
		info.Result = string(update.Result)
		if update.Result != backend.InProgressResult {
			info.EndTime = makeStringRef(time.Unix(update.EndTime, 0).UTC().Format(timeFormat))
			resourceChanges := make(map[string]int)
			for k, v := range update.ResourceChanges {
				resourceChanges[string(k)] = v
			}
			info.ResourceChanges = &resourceChanges
		}
		updatesJSON[idx] = info
	}

	return printJSON(updatesJSON)
}

func displayUpdatesConsole(updates []backend.UpdateInfo, opts display.Options) error {
	if len(updates) == 0 {
		fmt.Println("Stack has never been updated")
		return nil
	}

	printResourceChanges := func(background, text, sign, reset string, amount int) {
		msg := opts.Color.Colorize(fmt.Sprintf("%s%s%s%v%s", background, text, sign, amount, reset))
		fmt.Print(msg)
	}

	for _, update := range updates {

		fmt.Printf("UpdateKind: %v\n", update.Kind)
		if update.Result == "succeeded" {
			fmt.Print(opts.Color.Colorize(fmt.Sprintf("%sStatus: %v%s\n", colors.Green, update.Result, colors.Reset)))
		} else {
			fmt.Print(opts.Color.Colorize(fmt.Sprintf("%sStatus: %v%s\n", colors.Red, update.Result, colors.Reset)))
		}
		fmt.Printf("Message: %v\n", update.Message)

		printResourceChanges(colors.GreenBackground, colors.Black, "+", colors.Reset, update.ResourceChanges["create"])
		printResourceChanges(colors.RedBackground, colors.Black, "-", colors.Reset, update.ResourceChanges["delete"])
		printResourceChanges(colors.YellowBackground, colors.Black, "~", colors.Reset, update.ResourceChanges["update"])
		printResourceChanges(colors.BlueBackground, colors.Black, " ", colors.Reset, update.ResourceChanges["same"])

		timeStart := time.Unix(update.StartTime, 0)
		timeCreated := humanize.Time(timeStart)
		timeEnd := time.Unix(update.EndTime, 0)
		duration := timeEnd.Sub(timeStart)
		fmt.Printf("%sUpdated %s took %s\n", " ", timeCreated, duration)

		isEmpty := func(s string) bool {
			return len(strings.TrimSpace(s)) == 0
		}
		var keys []string
		for k := range update.Environment {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		indent := 4
		for _, k := range keys {
			if k == backend.GitHead && !isEmpty(update.Environment[k]) {
				fmt.Print(opts.Color.Colorize(
					fmt.Sprintf("%*s%s%s: %s%s\n", indent, "", colors.Yellow, k, update.Environment[k], colors.Reset)))
			} else if !isEmpty(update.Environment[k]) {
				fmt.Printf("%*s%s: %s\n", indent, "", k, update.Environment[k])
			}
		}
		fmt.Println("")
	}

	return nil
}
