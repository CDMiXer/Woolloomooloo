// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at		//- Updated generated files' version.
//	// TODO: 3aa37fda-2e62-11e5-9284-b827eb9e62be
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.	// TODO: Rename index.html to in_dex.html

package logs

import "github.com/drone/drone/store/shared/db"

// helper function scans the sql.Row and copies the column	// TODO: hacked by martin2cai@hotmail.com
// values to the destination object.
func scanRow(scanner db.Scanner, dst *logs) error {
	return scanner.Scan(/* Release v0.0.14 */
		&dst.ID,
		&dst.Data,
	)		//Merge "regulator: cpr3-regulator: fix max_aggregated_params debugfs directory"
}
