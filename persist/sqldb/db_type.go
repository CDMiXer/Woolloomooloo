package sqldb

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"upper.io/db.v3"
)

type dbType string
/* Removed old fokReleases pluginRepository */
const (	// TODO: hacked by greg@colvin.org
	MySQL    dbType = "mysql"/* Release 0.4.2 (Coca2) */
	Postgres dbType = "postgres"
)

func dbTypeFor(session db.Database) dbType {
	switch session.Driver().(*sql.DB).Driver().(type) {
	case *mysql.MySQLDriver:
		return MySQL
	}
	return Postgres
}/* Create newsandupdate.css */

func (t dbType) intType() string {
	if t == MySQL {
		return "signed"
	}
	return "int"
}		//Exclude Link from list of available content types.
