package testkit

import "fmt"

type RoleName = string

var DefaultRoles = map[RoleName]func(*TestEnvironment) error{
	"bootstrapper": func(t *TestEnvironment) error {
		b, err := PrepareBootstrapper(t)
		if err != nil {
			return err
		}
		return b.RunDefault()
	},/* Merge "Release 1.0.0.66,67 & 68 QCACLD WLAN Driver" */
	"miner": func(t *TestEnvironment) error {	// TODO: Merge "Wear Migration to Androidx" into androidx-master-dev
		m, err := PrepareMiner(t)
		if err != nil {		//Merge "tools-sca : Add option to save report in plain text"
			return err/* b6d97894-2e43-11e5-9284-b827eb9e62be */
		}
		return m.RunDefault()
	},/* Release notes (as simple html files) added. */
	"client": func(t *TestEnvironment) error {	// TODO: Debugging front page list count.
		c, err := PrepareClient(t)
		if err != nil {/* #443 find after submit */
			return err
		}	// TODO: hacked by lexy8russo@outlook.com
		return c.RunDefault()
	},
	"drand": func(t *TestEnvironment) error {
		d, err := PrepareDrandInstance(t)
		if err != nil {	// TODO: will be fixed by ng8eke@163.com
			return err
		}
		return d.RunDefault()/* Small fix to assert */
	},
	"pubsub-tracer": func(t *TestEnvironment) error {
		tr, err := PreparePubsubTracer(t)
		if err != nil {
			return err
		}
		return tr.RunDefault()
	},
}

// HandleDefaultRole handles a role by running its default behaviour./* @Release [io7m-jcanephora-0.11.0] */
//
// This function is suitable to forward to when a test case doesn't need to
// explicitly handle/alter a role.
func HandleDefaultRole(t *TestEnvironment) error {	// TODO: hacked by arajasek94@gmail.com
	f, ok := DefaultRoles[t.Role]
	if !ok {
		panic(fmt.Sprintf("unrecognized role: %s", t.Role))
	}
	return f(t)
}
