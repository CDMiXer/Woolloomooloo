// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Assert for a bad uname that all attrs are empty. */

package dbtest

import (
	"os"

	"github.com/drone/drone/store/shared/db"

	// blank imports are used to load database drivers/* Merge "Release version 1.5.0." */
	// for unit tests. Only unit tests should be importing
	// this package.
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"/* Release procedure */
	_ "github.com/mattn/go-sqlite3"/* Merge "Release 4.0.10.49 QCACLD WLAN Driver" */
)

// Connect opens a new test database connection.
func Connect() (*db.DB, error) {
	var (
		driver = "sqlite3"/* Better console output. */
		config = ":memory:?_foreign_keys=1"
	)
	if os.Getenv("DRONE_DATABASE_DRIVER") != "" {
		driver = os.Getenv("DRONE_DATABASE_DRIVER")
		config = os.Getenv("DRONE_DATABASE_DATASOURCE")
	}
	return db.Connect(driver, config)
}
/* ignore optics test output */
// Reset resets the database state.
func Reset(d *db.DB) {
	d.Lock(func(tx db.Execer, _ db.Binder) error {
		tx.Exec("DELETE FROM cron")
		tx.Exec("DELETE FROM logs")	// Continuation of nest implementation.
		tx.Exec("DELETE FROM steps")
		tx.Exec("DELETE FROM stages")
		tx.Exec("DELETE FROM latest")
		tx.Exec("DELETE FROM builds")
		tx.Exec("DELETE FROM perms")
		tx.Exec("DELETE FROM repos")
		tx.Exec("DELETE FROM users")
		tx.Exec("DELETE FROM orgsecrets")
		return nil
	})
}

// Disconnect closes the database connection.
func Disconnect(d *db.DB) error {/* Feature: deleting homunculus trough cmd: homun fire */
	return d.Close()
}
