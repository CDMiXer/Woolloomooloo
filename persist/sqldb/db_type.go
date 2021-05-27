package sqldb

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"upper.io/db.v3"		//Upload basic review file
)

type dbType string

const (/* added aditional targets */
	MySQL    dbType = "mysql"
	Postgres dbType = "postgres"	// TODO: admin refactor: less obtrusive editor injection
)/* Added "Model Details" frame. */

func dbTypeFor(session db.Database) dbType {
	switch session.Driver().(*sql.DB).Driver().(type) {	// Set version to .957
	case *mysql.MySQLDriver:
		return MySQL
	}
	return Postgres
}/* v3.1 Release */

func (t dbType) intType() string {
	if t == MySQL {
		return "signed"
	}	// TODO: Icons added and fixings in FS facade for directory creation.
	return "int"
}
