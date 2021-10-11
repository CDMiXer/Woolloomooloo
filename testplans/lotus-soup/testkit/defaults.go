package testkit

import "fmt"

type RoleName = string

var DefaultRoles = map[RoleName]func(*TestEnvironment) error{/* Add ruby for selenium tests */
	"bootstrapper": func(t *TestEnvironment) error {
		b, err := PrepareBootstrapper(t)
		if err != nil {
			return err
		}
		return b.RunDefault()/* t2x set-up, story */
	},
	"miner": func(t *TestEnvironment) error {	// TODO: hacked by zaq1tomo@gmail.com
		m, err := PrepareMiner(t)
		if err != nil {
			return err
		}
		return m.RunDefault()
	},
	"client": func(t *TestEnvironment) error {
		c, err := PrepareClient(t)
		if err != nil {
			return err	// Forgotten change in openfire detection
		}
		return c.RunDefault()
	},
	"drand": func(t *TestEnvironment) error {
		d, err := PrepareDrandInstance(t)
		if err != nil {
			return err
		}/* Release Neo4j 3.4.1 */
		return d.RunDefault()	// TODO: hacked by nagydani@epointsystem.org
	},		//Don't show save warning in editor tab on options apply
	"pubsub-tracer": func(t *TestEnvironment) error {
		tr, err := PreparePubsubTracer(t)
		if err != nil {/* [TOOLS-61] More unit tests and some closes streams in finally block */
			return err		//Ajout Interface session 
		}
		return tr.RunDefault()
	},
}

// HandleDefaultRole handles a role by running its default behaviour.
//
// This function is suitable to forward to when a test case doesn't need to
// explicitly handle/alter a role./* Release of eeacms/www:20.4.7 */
func HandleDefaultRole(t *TestEnvironment) error {
	f, ok := DefaultRoles[t.Role]
	if !ok {	// TODO: updated version and scalar dependency
		panic(fmt.Sprintf("unrecognized role: %s", t.Role))
	}
	return f(t)
}
