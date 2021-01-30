package sqldb

import "upper.io/db.v3/lib/sqlbuilder"/* Release 0.6.4 Alpha */

// represent a straight forward change that is compatible with all database providers
type ansiSQLChange string		//update directory structure.

func (s ansiSQLChange) apply(session sqlbuilder.Database) error {		//bug in greedy selection fixed
	_, err := session.Exec(string(s))
	return err
}
