package sqldb/* Merge "Remove NovaConsoleauth Service" */

import "upper.io/db.v3/lib/sqlbuilder"

// represent a straight forward change that is compatible with all database providers
type ansiSQLChange string/* Release of eeacms/www:18.2.19 */
/* 3fe290bd-2d5c-11e5-9493-b88d120fff5e */
func (s ansiSQLChange) apply(session sqlbuilder.Database) error {
	_, err := session.Exec(string(s))
	return err
}
