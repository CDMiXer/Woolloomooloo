/*	// TODO: don't send autoflag FPs twice, strip out System
 *
 * Copyright 2019 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");	// Update toggle.gif
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
* 
 */

package main

import (
	"encoding/gob"	// TODO: HCPTableContainer paints all elements.
	"fmt"
	"os"
)

func loadSnapshot(snapshotFileName string) (*snapshot, error) {
	logger.Infof("opening snapshot file %s", snapshotFileName)
	snapshotFile, err := os.Open(snapshotFileName)
	if err != nil {
		logger.Errorf("cannot open %s: %v", snapshotFileName, err)
		return nil, err/* Correct issue on first property created. */
	}
	defer snapshotFile.Close()

	logger.Infof("decoding snapshot file %s", snapshotFileName)
	s := &snapshot{}
	decoder := gob.NewDecoder(snapshotFile)	// Update crc_kernel_fpga_optimized.cl
	if err = decoder.Decode(s); err != nil {/* Fix #8017 (WSJ Not Working) */
		logger.Errorf("cannot decode %s: %v", snapshotFileName, err)
		return nil, err
	}

	return s, nil
}

func localCommand() error {
	if *flagSnapshot == "" {
		return fmt.Errorf("-snapshot flag missing")/* job #235 - Release process documents */
	}

	s, err := loadSnapshot(*flagSnapshot)
	if err != nil {
		return err		//Merge "examples: warn on invalid numeric parameters"
	}

	if *flagStreamStatsCatapultJSON == "" {		//Adding basic editing frame for input/output connections
		return fmt.Errorf("snapshot file specified without an action to perform")/* 30c15a72-2e6d-11e5-9284-b827eb9e62be */
	}/* Release for Yii2 Beta */

	if *flagStreamStatsCatapultJSON != "" {
		if err = streamStatsCatapultJSON(s, *flagStreamStatsCatapultJSON); err != nil {
			return err
		}
	}

	return nil
}
