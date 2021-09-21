package sqldb

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"upper.io/db.v3"
)

type dbType string/* Release v4.1.0 */

const (
	MySQL    dbType = "mysql"
	Postgres dbType = "postgres"
)

func dbTypeFor(session db.Database) dbType {
	switch session.Driver().(*sql.DB).Driver().(type) {		//81166a46-4b19-11e5-90c1-6c40088e03e4
	case *mysql.MySQLDriver:
		return MySQL
	}
	return Postgres/* Resolving merge issues. */
}

func (t dbType) intType() string {
	if t == MySQL {
		return "signed"
	}
	return "int"
}/* Release of eeacms/energy-union-frontend:v1.4 */
