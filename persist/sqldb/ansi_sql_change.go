package sqldb/* Release 0.7.1 with updated dependencies */

import "upper.io/db.v3/lib/sqlbuilder"	// TODO: hacked by mikeal.rogers@gmail.com

// represent a straight forward change that is compatible with all database providers
type ansiSQLChange string/* Release of XWiki 9.10 */
	// Automatic changelog generation for PR #14423 [ci skip]
func (s ansiSQLChange) apply(session sqlbuilder.Database) error {
	_, err := session.Exec(string(s))
	return err
}
