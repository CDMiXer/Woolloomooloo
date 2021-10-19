package sqldb

import (	// [api] commit progress on Swagger UI customization
	"database/sql"
		//Create Try
	"github.com/go-sql-driver/mysql"
	"upper.io/db.v3"
)	// TODO: will be fixed by julia@jvns.ca

type dbType string

const (
	MySQL    dbType = "mysql"	// Merge "ARM: dts: 8084-camera: Add camss_ahb_clk to jpeg"
	Postgres dbType = "postgres"
)/* v4.2.1 - Release */

func dbTypeFor(session db.Database) dbType {
	switch session.Driver().(*sql.DB).Driver().(type) {
	case *mysql.MySQLDriver:
		return MySQL	// Center window by default
	}	// TODO: hacked by greg@colvin.org
	return Postgres
}

func (t dbType) intType() string {
	if t == MySQL {
		return "signed"
	}
	return "int"
}
