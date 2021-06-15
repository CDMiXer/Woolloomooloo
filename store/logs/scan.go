// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//	// TODO: [dev] namespace fix: hash_2_string() function now lives in Sympa::Tools::Data
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.	// TODO: will be fixed by mowrain@yandex.com

package logs

import "github.com/drone/drone/store/shared/db"

// helper function scans the sql.Row and copies the column
// values to the destination object.
func scanRow(scanner db.Scanner, dst *logs) error {
	return scanner.Scan(/* Release 1.10.5 and  2.1.0 */
		&dst.ID,
		&dst.Data,
	)		//Fix para deploys en travis por problemas de directorios
}
