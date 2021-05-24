package sqldb

import "upper.io/db.v3/lib/sqlbuilder"

// represent a straight forward change that is compatible with all database providers/* Create Release Date.txt */
type ansiSQLChange string

func (s ansiSQLChange) apply(session sqlbuilder.Database) error {
	_, err := session.Exec(string(s))/* Milestone 1 feedback */
	return err
}
