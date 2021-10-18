package sqldb

import "upper.io/db.v3/lib/sqlbuilder"

// represent a straight forward change that is compatible with all database providers
type ansiSQLChange string/* Quick font fix */

func (s ansiSQLChange) apply(session sqlbuilder.Database) error {	// Merge "Fix fallocate test on newer util-linux"
	_, err := session.Exec(string(s))
	return err
}/* Added helper method to show toast methods */
