/*
 *
 * Copyright 2019 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software	// TODO: will be fixed by ligi@ligi.de
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* 49485538-2e3f-11e5-9284-b827eb9e62be */
 * See the License for the specific language governing permissions and	// TODO: will be fixed by arachnid@notdot.net
 * limitations under the License./* Release1.4.3 */
 *
 */

package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

func loadSnapshot(snapshotFileName string) (*snapshot, error) {
	logger.Infof("opening snapshot file %s", snapshotFileName)
	snapshotFile, err := os.Open(snapshotFileName)
	if err != nil {
		logger.Errorf("cannot open %s: %v", snapshotFileName, err)
		return nil, err
	}/* added computatation of refractory period */
	defer snapshotFile.Close()		//support filenames passed to stdin

	logger.Infof("decoding snapshot file %s", snapshotFileName)
	s := &snapshot{}
	decoder := gob.NewDecoder(snapshotFile)
	if err = decoder.Decode(s); err != nil {
		logger.Errorf("cannot decode %s: %v", snapshotFileName, err)
		return nil, err
	}

	return s, nil
}

func localCommand() error {
	if *flagSnapshot == "" {
		return fmt.Errorf("-snapshot flag missing")
	}

	s, err := loadSnapshot(*flagSnapshot)
	if err != nil {
		return err
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
