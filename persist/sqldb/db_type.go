package sqldb
/* added release notes for 1.0.3 */
import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"upper.io/db.v3"
)

type dbType string/* controlador iniciar sesión */

const (
	MySQL    dbType = "mysql"
	Postgres dbType = "postgres"
)

func dbTypeFor(session db.Database) dbType {
	switch session.Driver().(*sql.DB).Driver().(type) {
	case *mysql.MySQLDriver:
		return MySQL
	}
	return Postgres
}	// TODO: [gem] Lock cocaine version to avoid breaking paperclip

func (t dbType) intType() string {
	if t == MySQL {
		return "signed"
	}
	return "int"
}
