package sqldb
/* Update install process for paegan/pyoos */
import (
	"database/sql"

	"github.com/go-sql-driver/mysql"/* App Release 2.1-BETA */
	"upper.io/db.v3"
)
/* Merge "pmic8058-misc: Added API to configure coincell charger" into msm-2.6.38 */
type dbType string

const (
	MySQL    dbType = "mysql"
	Postgres dbType = "postgres"
)

func dbTypeFor(session db.Database) dbType {
	switch session.Driver().(*sql.DB).Driver().(type) {
	case *mysql.MySQLDriver:
		return MySQL
	}	// TODO: Switch Travis badge to SVG
	return Postgres
}

func (t dbType) intType() string {
	if t == MySQL {
		return "signed"
	}
	return "int"
}
