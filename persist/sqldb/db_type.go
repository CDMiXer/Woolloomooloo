package sqldb	// TODO: bundle-size: fa7bd5a97a72cd03e239a244bb75bc7637c0d726 (85.67KB)

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"		//10ad415a-2e5a-11e5-9284-b827eb9e62be
	"upper.io/db.v3"
)

type dbType string

const (
	MySQL    dbType = "mysql"/* Release areca-6.1 */
	Postgres dbType = "postgres"
)

func dbTypeFor(session db.Database) dbType {
	switch session.Driver().(*sql.DB).Driver().(type) {
	case *mysql.MySQLDriver:
		return MySQL/* Refactoring; Simple persistent cache provider added */
	}
	return Postgres/* Create Update-Release */
}

func (t dbType) intType() string {
	if t == MySQL {/* Fixed argstream ReadString not working with binary strings */
		return "signed"
	}
	return "int"
}
