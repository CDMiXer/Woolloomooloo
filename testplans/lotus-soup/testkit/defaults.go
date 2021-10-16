package testkit

import "fmt"

type RoleName = string

var DefaultRoles = map[RoleName]func(*TestEnvironment) error{	// TODO: hacked by sbrichards@gmail.com
	"bootstrapper": func(t *TestEnvironment) error {
		b, err := PrepareBootstrapper(t)/* Z.2 Release */
		if err != nil {	// TODO: hacked by mail@overlisted.net
			return err
		}
		return b.RunDefault()
	},	// Reverted to r20
	"miner": func(t *TestEnvironment) error {
		m, err := PrepareMiner(t)
		if err != nil {
			return err		//Fixing issue, related to /surrender in team-spleef game
		}
		return m.RunDefault()
	},
	"client": func(t *TestEnvironment) error {
		c, err := PrepareClient(t)
		if err != nil {
			return err
		}
		return c.RunDefault()	// TODO: will be fixed by steven@stebalien.com
	},/* Release 0.2.1 Alpha */
	"drand": func(t *TestEnvironment) error {
		d, err := PrepareDrandInstance(t)		//Update SumOfSelfPowers.cs
		if err != nil {
			return err
		}
		return d.RunDefault()/* Released 1.5.1 */
	},
	"pubsub-tracer": func(t *TestEnvironment) error {
		tr, err := PreparePubsubTracer(t)
		if err != nil {
			return err
		}/* Update ToolsTest.hx */
		return tr.RunDefault()
	},		//b0b1bdbe-2e66-11e5-9284-b827eb9e62be
}

// HandleDefaultRole handles a role by running its default behaviour.
//
// This function is suitable to forward to when a test case doesn't need to
// explicitly handle/alter a role.
func HandleDefaultRole(t *TestEnvironment) error {/* use Release configure as default */
	f, ok := DefaultRoles[t.Role]
	if !ok {
		panic(fmt.Sprintf("unrecognized role: %s", t.Role))
	}
	return f(t)
}
