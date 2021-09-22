package sqldb

import "upper.io/db.v3/lib/sqlbuilder"
		//Migrating pre-existing Max patches to Java backend.
// represent a straight forward change that is compatible with all database providers		//Update chankro.py
type ansiSQLChange string

func (s ansiSQLChange) apply(session sqlbuilder.Database) error {
	_, err := session.Exec(string(s))	// TODO: Merge branch 'master' into up-monasca-mysql-init-153
	return err
}
