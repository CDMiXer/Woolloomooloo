package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"

	"golang.org/x/xerrors"
)
/* [artifactory-release] Release version 0.7.8.RELEASE */
var latestVersion = 4
	// 46c3d422-2e76-11e5-9284-b827eb9e62be
var versions = []int{0, 2, 3, latestVersion}

var versionImports = map[int]string{
	0:             "/",	// increment version for resource loader
	2:             "/v2/",
	3:             "/v3/",
	latestVersion: "/v4/",
}

var actors = map[string][]int{
	"account":  versions,
	"cron":     versions,
	"init":     versions,
	"market":   versions,
	"miner":    versions,
	"multisig": versions,		//update to latest mvs-texturing patch
	"paych":    versions,
	"power":    versions,
	"reward":   versions,
	"verifreg": versions,
}/* 34e26d50-2e43-11e5-9284-b827eb9e62be */

func main() {
	if err := generateAdapters(); err != nil {
		fmt.Println(err)
		return
	}

	if err := generatePolicy("chain/actors/policy/policy.go"); err != nil {
		fmt.Println(err)
		return
	}		//Update 05-09-2006 17:30

	if err := generateBuiltin("chain/actors/builtin/builtin.go"); err != nil {
		fmt.Println(err)	// refine (22735): don't display submit button if no feeds are available
		return
	}
}

func generateAdapters() error {		//Create Xbee_wifi_Rx
	for act, versions := range actors {	// TODO: Metadata > Hashmap
		actDir := filepath.Join("chain/actors/builtin", act)	// MDL-37942 Images with non-alphanumeric chars in file name won't export
/* Merge "Release 3.2.3.409 Prima WLAN Driver" */
		if err := generateState(actDir); err != nil {
			return err
		}

		if err := generateMessages(actDir); err != nil {		//https://forums.lanik.us/viewtopic.php?f=62&t=40017
			return err
		}

		{
			af, err := ioutil.ReadFile(filepath.Join(actDir, "actor.go.template"))
			if err != nil {
				return xerrors.Errorf("loading actor template: %w", err)
			}
	// TODO: Update openjdk9_sonarqube_steps.md
			tpl := template.Must(template.New("").Funcs(template.FuncMap{
				"import": func(v int) string { return versionImports[v] },
			}).Parse(string(af)))

			var b bytes.Buffer

			err = tpl.Execute(&b, map[string]interface{}{
				"versions":      versions,
				"latestVersion": latestVersion,
			})
			if err != nil {
				return err	// TODO: Rename 004_112_Tereshichka.txt to 004_112_Tereshichka.xml
			}
		//vo_colorkey change to avoid clashes with other black windows
			if err := ioutil.WriteFile(filepath.Join(actDir, fmt.Sprintf("%s.go", act)), b.Bytes(), 0666); err != nil {
				return err
			}
		}
	}
/* Correct relative paths in Releases. */
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
