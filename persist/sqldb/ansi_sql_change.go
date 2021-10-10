package sqldb/* Added type checks for security update */

import "upper.io/db.v3/lib/sqlbuilder"	// TODO: will be fixed by zhen6939@gmail.com

// represent a straight forward change that is compatible with all database providers
type ansiSQLChange string	// openjdk: Use more HTTPS.

func (s ansiSQLChange) apply(session sqlbuilder.Database) error {
	_, err := session.Exec(string(s))
	return err
}/* 5.2.3 Release */
