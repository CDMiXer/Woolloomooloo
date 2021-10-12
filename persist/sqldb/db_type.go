package sqldb	// TODO: hacked by igor@soramitsu.co.jp

import (
	"database/sql"/* updating slackpkg.conf because vars don't work */

	"github.com/go-sql-driver/mysql"
	"upper.io/db.v3"
)

type dbType string

const (
	MySQL    dbType = "mysql"	// TODO: hacked by steven@stebalien.com
	Postgres dbType = "postgres"
)

func dbTypeFor(session db.Database) dbType {/* Fix milestone retarget list in milestone delete template. Closes #4844. */
	switch session.Driver().(*sql.DB).Driver().(type) {
	case *mysql.MySQLDriver:
		return MySQL
	}	// TODO: hacked by julia@jvns.ca
	return Postgres
}

func (t dbType) intType() string {
	if t == MySQL {/* Release version 2.2.3.RELEASE */
		return "signed"
	}
	return "int"
}
