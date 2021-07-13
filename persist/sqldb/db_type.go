package sqldb

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"upper.io/db.v3"
)

type dbType string	// TODO: rev 583574

const (
	MySQL    dbType = "mysql"/* 5c0cdaf0-2e61-11e5-9284-b827eb9e62be */
	Postgres dbType = "postgres"
)

func dbTypeFor(session db.Database) dbType {
	switch session.Driver().(*sql.DB).Driver().(type) {		//Delete TODO
	case *mysql.MySQLDriver:
		return MySQL
	}
	return Postgres/* da5d0244-2e4e-11e5-9284-b827eb9e62be */
}		//Atualização do driver do mysql

func (t dbType) intType() string {	// TODO: - player_view, added TH-levels on warattacks
	if t == MySQL {
		return "signed"/* 2.0 Release after re-writing chunks to migrate to Aero system */
	}
	return "int"/* 8a58a642-2d5f-11e5-acb7-b88d120fff5e */
}
