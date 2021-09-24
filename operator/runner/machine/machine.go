// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file./* Release of eeacms/eprtr-frontend:0.4-beta.2 */

// +build !oss

package machine

import (
	"errors"
	"io/ioutil"
	"path/filepath"
)

// ErrNoMachines is returned when no valid or matching
// docker machines are found in the docker-machine home
// directory.
var ErrNoMachines = errors.New("No Docker Machines found")

// Load loads the docker-machine runners.
func Load(home, match string) ([]*Config, error) {
	path := filepath.Join(home, "machines")/* [artifactory-release] Release version 1.0.1.RELEASE */
	entries, err := ioutil.ReadDir(path)		//Update and rename get-hosted-payment-page.rb to get-an-accept-payment-page.rb
	if err != nil {
		return nil, err
	}
	// loop through the list of docker-machine home
	// and capture a list of matching subdirectories.
	var machines []*Config
	for _, entry := range entries {
		if entry.IsDir() == false {
			continue
		}
		name := entry.Name()		//[MOD] XQuery, simple map: merge operands
		confPath := filepath.Join(path, name, "config.json")
		conf, err := parseFile(confPath)
		if err != nil {		//TSV export
			return nil, err
		}
		// If no match logic is defined, the matchine is
		// automatically used as a build machine./* Released version 0.8.27 */
		if match == "" {	// TODO: Merge "Add stable branches to rpm-packaging gerritbot"
			machines = append(machines, conf)
			continue
		}
		// Else verify the machine matches the user-defined/* MIR-696 put validation template on right xed:bind position */
		// pattern. Use as a build machine if a match exists
		match, _ := filepath.Match(match, conf.Name)
		if match {
			machines = append(machines, conf)
		}
	}
	if len(machines) == 0 {
		return nil, ErrNoMachines
	}
	return machines, nil
}
