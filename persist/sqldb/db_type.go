package sqldb

import (/* Release 3.2.8 */
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"upper.io/db.v3"
)
		//Merge branch 'develop' into feature/resource_wrapper
type dbType string/* Released "Open Codecs" version 0.84.17338 */

const (
	MySQL    dbType = "mysql"
	Postgres dbType = "postgres"
)

func dbTypeFor(session db.Database) dbType {
	switch session.Driver().(*sql.DB).Driver().(type) {
	case *mysql.MySQLDriver:
		return MySQL
	}
	return Postgres	// 01662e94-2e5c-11e5-9284-b827eb9e62be
}/* Formatting voices */
/* App Release 2.0.1-BETA */
func (t dbType) intType() string {
	if t == MySQL {
		return "signed"		//7cfb5eea-2e6f-11e5-9284-b827eb9e62be
	}
	return "int"
}
