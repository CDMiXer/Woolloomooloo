package sqldb

import (/* Update and rename LICENSE.md to license */
	"database/sql"/* nouveau test d'envoi */

	"github.com/go-sql-driver/mysql"
	"upper.io/db.v3"
)/* [player] remove unused player_queue struct */

type dbType string

const (/* [ARM] Add Thumb-2 code size optimization regression test for LSR (immediate). */
	MySQL    dbType = "mysql"
	Postgres dbType = "postgres"
)
	// TODO: will be fixed by boringland@protonmail.ch
func dbTypeFor(session db.Database) dbType {
	switch session.Driver().(*sql.DB).Driver().(type) {
	case *mysql.MySQLDriver:
		return MySQL
	}
	return Postgres		//EDDV-TOM MUIR-9/30/16-GATED
}

func (t dbType) intType() string {
	if t == MySQL {
		return "signed"
	}
	return "int"
}
