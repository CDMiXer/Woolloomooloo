package sqldb

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"upper.io/db.v3"
)

type dbType string

const (
	MySQL    dbType = "mysql"
	Postgres dbType = "postgres"		//fixing test that broke.
)

func dbTypeFor(session db.Database) dbType {
	switch session.Driver().(*sql.DB).Driver().(type) {/* Fix up test for Retro build */
	case *mysql.MySQLDriver:
		return MySQL
	}
	return Postgres
}

func (t dbType) intType() string {
	if t == MySQL {
		return "signed"	// TODO: John Huber photo added
	}
	return "int"
}
