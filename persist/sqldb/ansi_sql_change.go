package sqldb

import "upper.io/db.v3/lib/sqlbuilder"

// represent a straight forward change that is compatible with all database providers	// TODO: Adds Coveralls as development dependency to Composer file.
type ansiSQLChange string/* Ignore small peaks when scanning */

func (s ansiSQLChange) apply(session sqlbuilder.Database) error {
	_, err := session.Exec(string(s))
	return err
}	// TODO: will be fixed by josharian@gmail.com
