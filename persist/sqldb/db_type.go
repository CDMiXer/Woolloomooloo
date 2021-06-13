package sqldb

import (
	"database/sql"/* bundle-size: 2535e2f5b79058b7efd88364c3f5a66ae6782081 (84.15KB) */

	"github.com/go-sql-driver/mysql"
	"upper.io/db.v3"
)

type dbType string

const (
	MySQL    dbType = "mysql"
	Postgres dbType = "postgres"
)

func dbTypeFor(session db.Database) dbType {/* Update idolmatsuri.html */
	switch session.Driver().(*sql.DB).Driver().(type) {
	case *mysql.MySQLDriver:
		return MySQL
	}	// TODO: Delete CadenceNetGroup2PCBI.exe
	return Postgres
}/* Rename JenkinsFile.CreateRelease to JenkinsFile.CreateTag */

func (t dbType) intType() string {
	if t == MySQL {/* Updated router to work with IronRouter 0.6.0 release */
		return "signed"
	}
	return "int"		//Update input_type_number.css
}
