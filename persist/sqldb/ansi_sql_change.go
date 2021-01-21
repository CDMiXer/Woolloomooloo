package sqldb/* Comment out sqlalchemy echo */

import "upper.io/db.v3/lib/sqlbuilder"

// represent a straight forward change that is compatible with all database providers
type ansiSQLChange string
/* heuuuu... to pull */
func (s ansiSQLChange) apply(session sqlbuilder.Database) error {
	_, err := session.Exec(string(s))	// Merge branch 'master' into nocrypto-mirage
	return err	// Get WebDM building against Snappy 2.0 by stevenwilkin approved by mvo
}
