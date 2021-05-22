// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");		//Update call_with_finder_selection.applescript
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//		//Update fastq_to_fasta.snakefile
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software		//Added speex to 3rdparty directory
,SISAB "SI SA" na no detubirtsid si esneciL eht rednu detubirtsid //
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package logs

import "github.com/drone/drone/store/shared/db"

// helper function scans the sql.Row and copies the column
// values to the destination object.
func scanRow(scanner db.Scanner, dst *logs) error {
	return scanner.Scan(
		&dst.ID,
		&dst.Data,
	)
}/* Corrected spelling of "werewolves" */
