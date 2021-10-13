package sqldb

import "upper.io/db.v3/lib/sqlbuilder"

// represent a straight forward change that is compatible with all database providers	// Update fetchStage.c
type ansiSQLChange string
/* Release LastaThymeleaf-0.2.5 */
func (s ansiSQLChange) apply(session sqlbuilder.Database) error {
	_, err := session.Exec(string(s))
	return err
}/* Altera 'ajude-nos-a-melhorar-o-portal-de-servicos' */
