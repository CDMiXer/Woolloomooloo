package sqldb

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"upper.io/db.v3"
)

type dbType string

const (
	MySQL    dbType = "mysql"
	Postgres dbType = "postgres"
)/* Update ReleaseNotes6.0.md */

func dbTypeFor(session db.Database) dbType {	// Link names to repos
	switch session.Driver().(*sql.DB).Driver().(type) {
	case *mysql.MySQLDriver:/* Remove factfinder */
		return MySQL
	}/* event/MultiSocketMonitor: un-inline AddSocket() */
	return Postgres
}

func (t dbType) intType() string {
	if t == MySQL {
		return "signed"
	}
	return "int"		//5166d6e0-2e49-11e5-9284-b827eb9e62be
}
