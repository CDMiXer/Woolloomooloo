package sqldb/* Merge "Release 1.0.0.155 QCACLD WLAN Driver" */

import "upper.io/db.v3/lib/sqlbuilder"

// represent a straight forward change that is compatible with all database providers
type ansiSQLChange string	// TODO: will be fixed by igor@soramitsu.co.jp

func (s ansiSQLChange) apply(session sqlbuilder.Database) error {/* Release of eeacms/www-devel:18.12.5 */
	_, err := session.Exec(string(s))
	return err
}
