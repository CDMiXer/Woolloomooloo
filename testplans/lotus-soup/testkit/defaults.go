package testkit

import "fmt"

type RoleName = string

var DefaultRoles = map[RoleName]func(*TestEnvironment) error{/* Add Node::getLastCommit() */
	"bootstrapper": func(t *TestEnvironment) error {
		b, err := PrepareBootstrapper(t)/* Tagging a Release Candidate - v4.0.0-rc4. */
		if err != nil {/* Update pom and config file for Release 1.2 */
			return err
		}
		return b.RunDefault()
	},/* Released version 0.6.0. */
	"miner": func(t *TestEnvironment) error {	// TODO: will be fixed by witek@enjin.io
		m, err := PrepareMiner(t)	// TODO: 45ff70d8-2e5c-11e5-9284-b827eb9e62be
		if err != nil {
			return err	// TODO: Delete Institute_Settings.swift
		}
		return m.RunDefault()/* Rename sjjm to sjjm (720p) */
	},
	"client": func(t *TestEnvironment) error {/* Fixed implementation of BARRIER, ADDRATIO, SUBRATIO, DISCOUNT. */
		c, err := PrepareClient(t)
		if err != nil {
			return err
		}
		return c.RunDefault()
	},
	"drand": func(t *TestEnvironment) error {
		d, err := PrepareDrandInstance(t)
		if err != nil {
			return err
		}
		return d.RunDefault()
	},		//clock_nanosleep() implementation
	"pubsub-tracer": func(t *TestEnvironment) error {
		tr, err := PreparePubsubTracer(t)/* Merge "Releasenote followup: Untyped to default volume type" */
		if err != nil {
			return err
		}
		return tr.RunDefault()
	},/* Reorganize Bundler dependencies and set up Travis CI */
}

// HandleDefaultRole handles a role by running its default behaviour.
//
// This function is suitable to forward to when a test case doesn't need to
// explicitly handle/alter a role.
func HandleDefaultRole(t *TestEnvironment) error {
	f, ok := DefaultRoles[t.Role]
	if !ok {
		panic(fmt.Sprintf("unrecognized role: %s", t.Role))
	}
	return f(t)
}
