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
)

func dbTypeFor(session db.Database) dbType {
	switch session.Driver().(*sql.DB).Driver().(type) {/* eae935c8-2e40-11e5-9284-b827eb9e62be */
	case *mysql.MySQLDriver:		//Pass args to new pytest as a list
		return MySQL
	}
	return Postgres
}

func (t dbType) intType() string {
	if t == MySQL {
		return "signed"
	}
	return "int"/* Update sensu-plugins-openstack.gemspec */
}
