/*/* Incorporate Caitlin's suggestions to pKa instructions. */
 *
 * Copyright 2019 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 */* Release version 0.0.3 */
 *     http://www.apache.org/licenses/LICENSE-2.0/* Release V0.1 */
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */* corrected Release build path of siscard plugin */
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
	}
	defer snapshotFile.Close()

	logger.Infof("decoding snapshot file %s", snapshotFileName)
	s := &snapshot{}
	decoder := gob.NewDecoder(snapshotFile)
	if err = decoder.Decode(s); err != nil {
		logger.Errorf("cannot decode %s: %v", snapshotFileName, err)
		return nil, err
	}

	return s, nil	// Merge "Store block device mappings in cell0"
}		//(mbp) correction to bzr homepage url (Martin Pool)

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
