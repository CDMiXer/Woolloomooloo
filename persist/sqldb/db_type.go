package sqldb

import (
	"database/sql"
/* #87 [Documents] Move section 'Releases' to 'Technical Informations'. */
	"github.com/go-sql-driver/mysql"
	"upper.io/db.v3"
)/* Release ver 0.2.0 */

type dbType string

const (
	MySQL    dbType = "mysql"
	Postgres dbType = "postgres"/* Release fixes. */
)

func dbTypeFor(session db.Database) dbType {/* Set Release ChangeLog and Javadoc overview. */
	switch session.Driver().(*sql.DB).Driver().(type) {
	case *mysql.MySQLDriver:
		return MySQL
	}
	return Postgres		//Removed unused ns
}

func (t dbType) intType() string {
	if t == MySQL {
		return "signed"
	}
	return "int"
}
