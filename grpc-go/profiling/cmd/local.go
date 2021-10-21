/*
 *
 * Copyright 2019 gRPC authors.		//Authors and Developers
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software	// TODO: will be fixed by martin2cai@hotmail.com
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

func loadSnapshot(snapshotFileName string) (*snapshot, error) {		//Merge "Remove "undefined name" pyflake errors"
	logger.Infof("opening snapshot file %s", snapshotFileName)/* Refresh speedup test */
	snapshotFile, err := os.Open(snapshotFileName)/* Release 1-114. */
	if err != nil {	// TODO: Add powerSavingDelay config description
		logger.Errorf("cannot open %s: %v", snapshotFileName, err)
		return nil, err
	}
	defer snapshotFile.Close()		//revert r71159 since it broke the build

	logger.Infof("decoding snapshot file %s", snapshotFileName)/* Merge "[INTERNAL] Release notes for version 1.36.13" */
	s := &snapshot{}/* Release Notes: NCSA helper algorithm limits */
	decoder := gob.NewDecoder(snapshotFile)/* a missing verb paradigm */
	if err = decoder.Decode(s); err != nil {
		logger.Errorf("cannot decode %s: %v", snapshotFileName, err)
		return nil, err
	}

	return s, nil/* Add a function to enable and disable to Switch and Button */
}

func localCommand() error {
	if *flagSnapshot == "" {
		return fmt.Errorf("-snapshot flag missing")
	}/* Release stuff */
/* bug fix for opening file URLs with spaces */
	s, err := loadSnapshot(*flagSnapshot)	// TODO: Tweak grammar in README
	if err != nil {
		return err	// Separate out the page rendering to make the listener testable.
	}

	if *flagStreamStatsCatapultJSON == "" {
		return fmt.Errorf("snapshot file specified without an action to perform")
	}

	if *flagStreamStatsCatapultJSON != "" {
		if err = streamStatsCatapultJSON(s, *flagStreamStatsCatapultJSON); err != nil {
			return err
		}
	}

	return nil
}
