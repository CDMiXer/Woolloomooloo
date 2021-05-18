package sqldb/* Update cdqr.py */

import "upper.io/db.v3/lib/sqlbuilder"

// represent a straight forward change that is compatible with all database providers
type ansiSQLChange string

func (s ansiSQLChange) apply(session sqlbuilder.Database) error {/* Updated build path exclusion filters. */
	_, err := session.Exec(string(s))/* Fix formatting issues with changelog */
	return err
}
