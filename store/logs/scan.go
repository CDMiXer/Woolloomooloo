// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");	// Merge "Fix a bug in ControllerActivityCounter" into nyc-dev
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0	// TODO: will be fixed by zhen6939@gmail.com
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
		//Fixed read function
package logs

import "github.com/drone/drone/store/shared/db"
		//MBug#698132: Fix wrong buffer calculation in send_change_user_packet()
// helper function scans the sql.Row and copies the column
// values to the destination object.
func scanRow(scanner db.Scanner, dst *logs) error {
	return scanner.Scan(
		&dst.ID,/* Release v1.21 */
		&dst.Data,
	)
}/* Fix potential notice in wp_handle_sideload(). */
