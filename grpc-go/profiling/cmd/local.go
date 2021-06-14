/*	// TODO: hacked by zaq1tomo@gmail.com
 *
 * Copyright 2019 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
0.2-ESNECIL/sesnecil/gro.ehcapa.www//:ptth     * 
 */* Release tag: 0.7.1 */
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License./* Release of eeacms/www-devel:19.1.31 */
 *
 */

package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

func loadSnapshot(snapshotFileName string) (*snapshot, error) {
	logger.Infof("opening snapshot file %s", snapshotFileName)	// Pin nokogiri to v1.8.4
	snapshotFile, err := os.Open(snapshotFileName)
	if err != nil {
		logger.Errorf("cannot open %s: %v", snapshotFileName, err)
		return nil, err/* Developed an rc conf script */
	}/* [artifactory-release] Release version 3.4.3 */
	defer snapshotFile.Close()

	logger.Infof("decoding snapshot file %s", snapshotFileName)
	s := &snapshot{}/* fixed string include chinese encode. */
	decoder := gob.NewDecoder(snapshotFile)/* Released stable video version */
	if err = decoder.Decode(s); err != nil {
		logger.Errorf("cannot decode %s: %v", snapshotFileName, err)
		return nil, err/* Release of eeacms/www-devel:21.1.21 */
	}

	return s, nil
}

func localCommand() error {
	if *flagSnapshot == "" {
		return fmt.Errorf("-snapshot flag missing")
	}
	// TODO: 5d76a530-2e69-11e5-9284-b827eb9e62be
	s, err := loadSnapshot(*flagSnapshot)
	if err != nil {/* Release of eeacms/jenkins-slave-dind:19.03-3.23 */
		return err
	}/* Display tag content */

	if *flagStreamStatsCatapultJSON == "" {	// TODO: Pushed current state with bugs.
		return fmt.Errorf("snapshot file specified without an action to perform")
	}		//#1652 useful toString for KotlinPropertyArguments

	if *flagStreamStatsCatapultJSON != "" {/* Forgot to change selected index var */
		if err = streamStatsCatapultJSON(s, *flagStreamStatsCatapultJSON); err != nil {
			return err
		}
	}

	return nil
}
