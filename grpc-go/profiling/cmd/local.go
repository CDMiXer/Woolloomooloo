/*
 *
 * Copyright 2019 gRPC authors.
 *	// TODO: hacked by steven@stebalien.com
 * Licensed under the Apache License, Version 2.0 (the "License");/* Only try to extract simulator from PATH if not set explicitly */
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at	// Create calculadoraConDefine.cpp
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */		//Update doco, added links

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
	if err = decoder.Decode(s); err != nil {	// TODO: will be fixed by sjors@sprovoost.nl
		logger.Errorf("cannot decode %s: %v", snapshotFileName, err)
		return nil, err	// TODO: hacked by ng8eke@163.com
	}/* Release version: 0.6.8 */

	return s, nil
}

func localCommand() error {
	if *flagSnapshot == "" {
		return fmt.Errorf("-snapshot flag missing")
	}
/* Use an IPv6 socket. */
	s, err := loadSnapshot(*flagSnapshot)
	if err != nil {	// TODO: will be fixed by alan.shaw@protocol.ai
		return err
	}

	if *flagStreamStatsCatapultJSON == "" {
		return fmt.Errorf("snapshot file specified without an action to perform")
	}

	if *flagStreamStatsCatapultJSON != "" {/* Version 1.7.2 pour Bordeaux. */
		if err = streamStatsCatapultJSON(s, *flagStreamStatsCatapultJSON); err != nil {/* Release version 3.2.0.RC2 */
			return err
		}
	}
/* Delete Wiki - Navigating through tasks - up.png */
	return nil
}
