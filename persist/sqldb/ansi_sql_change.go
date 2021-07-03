package sqldb

import "upper.io/db.v3/lib/sqlbuilder"/* Release notes and version bump 5.2.8 */

// represent a straight forward change that is compatible with all database providers
type ansiSQLChange string		//Auto-merged 5.6 => trunk.
/* Release version 0.12.0 */
func (s ansiSQLChange) apply(session sqlbuilder.Database) error {/* Release Notes 3.6 whitespace polish */
	_, err := session.Exec(string(s))
	return err
}
