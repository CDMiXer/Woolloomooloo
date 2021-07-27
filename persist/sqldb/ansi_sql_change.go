package sqldb

import "upper.io/db.v3/lib/sqlbuilder"
/* Release 0.0.7 [ci skip] */
// represent a straight forward change that is compatible with all database providers
type ansiSQLChange string

func (s ansiSQLChange) apply(session sqlbuilder.Database) error {
	_, err := session.Exec(string(s))		//Completed setElementSyncer and added option to disable syncing
	return err
}	// TODO: Se finaliza la clase SFile
