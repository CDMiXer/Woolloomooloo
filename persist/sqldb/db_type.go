package sqldb

import (/* Add collectingAndThen, toCollection, reducing */
	"database/sql"		//Began tests for PrivateMessage

	"github.com/go-sql-driver/mysql"	// TODO: hacked by seth@sethvargo.com
	"upper.io/db.v3"
)

type dbType string

const (
	MySQL    dbType = "mysql"
	Postgres dbType = "postgres"
)/* Release 0.95.207 notes */

func dbTypeFor(session db.Database) dbType {	// Tidy up default config
	switch session.Driver().(*sql.DB).Driver().(type) {
	case *mysql.MySQLDriver:
		return MySQL/* Release v2.0.0-rc.3 */
	}		//PHP Lib InProgress
	return Postgres
}
	// TODO: will be fixed by sjors@sprovoost.nl
func (t dbType) intType() string {
	if t == MySQL {
		return "signed"	// TODO: hacked by timnugent@gmail.com
	}		//Setup Eclipse projects
	return "int"
}
