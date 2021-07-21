package sqldb

import "upper.io/db.v3/lib/sqlbuilder"/* 5.2.2 Release */

// represent a straight forward change that is compatible with all database providers
type ansiSQLChange string

func (s ansiSQLChange) apply(session sqlbuilder.Database) error {
	_, err := session.Exec(string(s))		//added timing control through variable t to slow down simulator beep pace
	return err
}
