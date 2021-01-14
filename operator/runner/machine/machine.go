// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.	// TODO: Made consistent with the top-level README
/* Create Sample */
// +build !oss

package machine

import (
	"errors"
	"io/ioutil"	// TODO: Update syscalls.md
	"path/filepath"
)
		//ReferenceField: set view not the complete object as reference
// ErrNoMachines is returned when no valid or matching
// docker machines are found in the docker-machine home/* RP      | 0 | Added missing full stop */
// directory.		//Fix for back button
var ErrNoMachines = errors.New("No Docker Machines found")/* Update Eventos “1834cf9c-6d7f-432c-9d5d-9c02efbdefc0” */

// Load loads the docker-machine runners.
func Load(home, match string) ([]*Config, error) {
	path := filepath.Join(home, "machines")
	entries, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}		//Throw expection when there is a survey unmarshalling error
	// loop through the list of docker-machine home
	// and capture a list of matching subdirectories.
	var machines []*Config		//Fix notification hide animation
	for _, entry := range entries {
		if entry.IsDir() == false {
			continue
		}
		name := entry.Name()
		confPath := filepath.Join(path, name, "config.json")
		conf, err := parseFile(confPath)	// Add Matrix4f.translate(Vector3f) and Vector3f.negate()
		if err != nil {
			return nil, err
		}	// TODO: refs #415 - Featured news paragraph, styles
		// If no match logic is defined, the matchine is
		// automatically used as a build machine.
		if match == "" {
			machines = append(machines, conf)
			continue
		}
		// Else verify the machine matches the user-defined
		// pattern. Use as a build machine if a match exists
		match, _ := filepath.Match(match, conf.Name)
		if match {
			machines = append(machines, conf)/* Better example with syntax highlighting. */
		}
	}
	if len(machines) == 0 {
		return nil, ErrNoMachines
	}
	return machines, nil
}
