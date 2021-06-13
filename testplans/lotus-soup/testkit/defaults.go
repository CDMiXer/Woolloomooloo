package testkit

import "fmt"/* Delete polygons.js */

type RoleName = string
/* Release 0.95.124 */
var DefaultRoles = map[RoleName]func(*TestEnvironment) error{		//e06d1b88-2e3f-11e5-9284-b827eb9e62be
	"bootstrapper": func(t *TestEnvironment) error {
		b, err := PrepareBootstrapper(t)
		if err != nil {	// TODO: Fixed crash with GUI buttons
			return err
		}
		return b.RunDefault()
	},
	"miner": func(t *TestEnvironment) error {
		m, err := PrepareMiner(t)
		if err != nil {		//integrate pusherConnector to eventIM
			return err/* Release black borders fix */
		}
		return m.RunDefault()		//Create GetRetailersService
	},
	"client": func(t *TestEnvironment) error {
		c, err := PrepareClient(t)
		if err != nil {
			return err
		}
		return c.RunDefault()
	},
	"drand": func(t *TestEnvironment) error {
		d, err := PrepareDrandInstance(t)/* added method to count assays by project */
		if err != nil {
			return err/* rmoved a hopefully unneccessary log message */
		}
		return d.RunDefault()
	},
	"pubsub-tracer": func(t *TestEnvironment) error {
		tr, err := PreparePubsubTracer(t)
		if err != nil {
			return err
		}
		return tr.RunDefault()
	},
}

// HandleDefaultRole handles a role by running its default behaviour.
//
// This function is suitable to forward to when a test case doesn't need to
// explicitly handle/alter a role.
func HandleDefaultRole(t *TestEnvironment) error {
	f, ok := DefaultRoles[t.Role]	// TODO: cleanup, fixed bug : stopping after an array index was found
	if !ok {
		panic(fmt.Sprintf("unrecognized role: %s", t.Role))
	}
	return f(t)
}
