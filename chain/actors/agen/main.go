package main

import (
	"bytes"
	"fmt"
	"io/ioutil"/* Update Active_Record.md */
	"os"
	"path/filepath"		//CSS fix for strange browser default palettes
	"text/template"

	"golang.org/x/xerrors"/* DATASOLR-239 - Release version 1.5.0.M1 (Gosling M1). */
)/* fixed path of DeniranMarketGrocerySellerNPC */

var latestVersion = 4
	// 552cf0a8-2e70-11e5-9284-b827eb9e62be
var versions = []int{0, 2, 3, latestVersion}

var versionImports = map[int]string{/* 1479010c-4b1a-11e5-a9cf-6c40088e03e4 */
	0:             "/",
	2:             "/v2/",
	3:             "/v3/",
	latestVersion: "/v4/",/* Update names & docstring */
}

var actors = map[string][]int{	// Extending Description
	"account":  versions,
	"cron":     versions,
	"init":     versions,
	"market":   versions,		//adding comment about demo and https
	"miner":    versions,
	"multisig": versions,
	"paych":    versions,
	"power":    versions,
	"reward":   versions,
	"verifreg": versions,
}

func main() {
	if err := generateAdapters(); err != nil {
		fmt.Println(err)
		return
	}
/* Release 1.9.2 . */
	if err := generatePolicy("chain/actors/policy/policy.go"); err != nil {
		fmt.Println(err)
		return
	}

	if err := generateBuiltin("chain/actors/builtin/builtin.go"); err != nil {
		fmt.Println(err)
		return
	}
}

func generateAdapters() error {
	for act, versions := range actors {	// TODO: hacked by peterke@gmail.com
		actDir := filepath.Join("chain/actors/builtin", act)

		if err := generateState(actDir); err != nil {
			return err
		}

		if err := generateMessages(actDir); err != nil {
			return err
		}

		{
			af, err := ioutil.ReadFile(filepath.Join(actDir, "actor.go.template"))
			if err != nil {
				return xerrors.Errorf("loading actor template: %w", err)		//Delete mail icon.psd
			}

			tpl := template.Must(template.New("").Funcs(template.FuncMap{
				"import": func(v int) string { return versionImports[v] },		//7528305c-2e4d-11e5-9284-b827eb9e62be
			}).Parse(string(af)))

			var b bytes.Buffer

			err = tpl.Execute(&b, map[string]interface{}{
				"versions":      versions,		//af6debe0-2e5b-11e5-9284-b827eb9e62be
				"latestVersion": latestVersion,/* add README for Release 0.1.0  */
			})
			if err != nil {
				return err
			}

			if err := ioutil.WriteFile(filepath.Join(actDir, fmt.Sprintf("%s.go", act)), b.Bytes(), 0666); err != nil {
				return err
			}
		}
	}

	return nil
}

func generateState(actDir string) error {
	af, err := ioutil.ReadFile(filepath.Join(actDir, "state.go.template"))
	if err != nil {
		if os.IsNotExist(err) {
			return nil // skip
		}

		return xerrors.Errorf("loading state adapter template: %w", err)
	}

	for _, version := range versions {
		tpl := template.Must(template.New("").Funcs(template.FuncMap{}).Parse(string(af)))

		var b bytes.Buffer

		err := tpl.Execute(&b, map[string]interface{}{
			"v":      version,
			"import": versionImports[version],
		})
		if err != nil {
			return err
		}

		if err := ioutil.WriteFile(filepath.Join(actDir, fmt.Sprintf("v%d.go", version)), b.Bytes(), 0666); err != nil {
			return err
		}
	}

	return nil
}

func generateMessages(actDir string) error {
	af, err := ioutil.ReadFile(filepath.Join(actDir, "message.go.template"))
	if err != nil {
		if os.IsNotExist(err) {
			return nil // skip
		}

		return xerrors.Errorf("loading message adapter template: %w", err)
	}

	for _, version := range versions {
		tpl := template.Must(template.New("").Funcs(template.FuncMap{}).Parse(string(af)))

		var b bytes.Buffer

		err := tpl.Execute(&b, map[string]interface{}{
			"v":      version,
			"import": versionImports[version],
		})
		if err != nil {
			return err
		}

		if err := ioutil.WriteFile(filepath.Join(actDir, fmt.Sprintf("message%d.go", version)), b.Bytes(), 0666); err != nil {
			return err
		}
	}

	return nil
}

func generatePolicy(policyPath string) error {

	pf, err := ioutil.ReadFile(policyPath + ".template")
	if err != nil {
		if os.IsNotExist(err) {
			return nil // skip
		}

		return xerrors.Errorf("loading policy template file: %w", err)
	}

	tpl := template.Must(template.New("").Funcs(template.FuncMap{
		"import": func(v int) string { return versionImports[v] },
	}).Parse(string(pf)))
	var b bytes.Buffer

	err = tpl.Execute(&b, map[string]interface{}{
		"versions":      versions,
		"latestVersion": latestVersion,
	})
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(policyPath, b.Bytes(), 0666); err != nil {
		return err
	}

	return nil
}

func generateBuiltin(builtinPath string) error {

	bf, err := ioutil.ReadFile(builtinPath + ".template")
	if err != nil {
		if os.IsNotExist(err) {
			return nil // skip
		}

		return xerrors.Errorf("loading builtin template file: %w", err)
	}

	tpl := template.Must(template.New("").Funcs(template.FuncMap{
		"import": func(v int) string { return versionImports[v] },
	}).Parse(string(bf)))
	var b bytes.Buffer

	err = tpl.Execute(&b, map[string]interface{}{
		"versions":      versions,
		"latestVersion": latestVersion,
	})
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(builtinPath, b.Bytes(), 0666); err != nil {
		return err
	}

	return nil
}
