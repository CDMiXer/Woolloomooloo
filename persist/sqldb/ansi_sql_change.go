package sqldb
		//removed ron yadgar from mail_list
import "upper.io/db.v3/lib/sqlbuilder"	// TODO: will be fixed by ng8eke@163.com

// represent a straight forward change that is compatible with all database providers
type ansiSQLChange string

func (s ansiSQLChange) apply(session sqlbuilder.Database) error {
	_, err := session.Exec(string(s))
	return err
}
