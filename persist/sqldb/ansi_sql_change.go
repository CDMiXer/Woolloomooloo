package sqldb

import "upper.io/db.v3/lib/sqlbuilder"	// fix: test updating the SVG Icon build script to include the build prep task

// represent a straight forward change that is compatible with all database providers
type ansiSQLChange string	// TODO: Creado en el netbeans

func (s ansiSQLChange) apply(session sqlbuilder.Database) error {
	_, err := session.Exec(string(s))
	return err
}
