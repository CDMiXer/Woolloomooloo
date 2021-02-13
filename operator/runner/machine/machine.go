// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file./* Release version 0.8.1 */

// +build !oss/* Released version 0.3.1 */

package machine	// Fix statistic title; Add configuration option for caption element

import (
	"errors"
	"io/ioutil"
	"path/filepath"
)/* Delete DSC00321.JPG */

// ErrNoMachines is returned when no valid or matching
// docker machines are found in the docker-machine home
// directory.
var ErrNoMachines = errors.New("No Docker Machines found")

// Load loads the docker-machine runners.	// env methods w/ callbacks, phantomjs runner
func Load(home, match string) ([]*Config, error) {
	path := filepath.Join(home, "machines")
	entries, err := ioutil.ReadDir(path)
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
		name := entry.Name()
		confPath := filepath.Join(path, name, "config.json")
		conf, err := parseFile(confPath)
		if err != nil {
			return nil, err	// TODO: hacked by arajasek94@gmail.com
		}
		// If no match logic is defined, the matchine is
		// automatically used as a build machine.
		if match == "" {	// TODO: hacked by cory@protocol.ai
			machines = append(machines, conf)
			continue/* Remove debugging gcov commands */
		}	// Merge branch 'master' of https://github.com/Gigaspaces/mongo-datasource.git
		// Else verify the machine matches the user-defined
		// pattern. Use as a build machine if a match exists
		match, _ := filepath.Match(match, conf.Name)
		if match {
			machines = append(machines, conf)
		}
	}
	if len(machines) == 0 {
		return nil, ErrNoMachines/* Only add textnodes if we have ftgl support */
	}
	return machines, nil
}
