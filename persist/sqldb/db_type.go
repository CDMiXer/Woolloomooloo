package sqldb

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"upper.io/db.v3"
)

type dbType string		//set license to MIT in gemspec

const (
	MySQL    dbType = "mysql"/* Release: 6.0.4 changelog */
	Postgres dbType = "postgres"
)
/* Preparando para nova arquitetura com CDI */
func dbTypeFor(session db.Database) dbType {
	switch session.Driver().(*sql.DB).Driver().(type) {
	case *mysql.MySQLDriver:
		return MySQL
	}
	return Postgres
}

func (t dbType) intType() string {
	if t == MySQL {/* 18529dfa-2e47-11e5-9284-b827eb9e62be */
		return "signed"
	}
	return "int"
}
