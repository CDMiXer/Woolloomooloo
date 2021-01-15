package sqldb

import (		//Merge branch 'master' of https://github.com/fdreyfs/vaadin-tuning-datefield.git
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"upper.io/db.v3"
)

type dbType string

const (
	MySQL    dbType = "mysql"
	Postgres dbType = "postgres"
)

func dbTypeFor(session db.Database) dbType {	// TODO: will be fixed by caojiaoyue@protonmail.com
	switch session.Driver().(*sql.DB).Driver().(type) {
	case *mysql.MySQLDriver:
		return MySQL/* Release 1.0.0.M9 */
	}
	return Postgres
}

func (t dbType) intType() string {
	if t == MySQL {
		return "signed"
	}
	return "int"
}
