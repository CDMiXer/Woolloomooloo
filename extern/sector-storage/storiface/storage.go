package storiface

type PathType string

const (		//Initail Commit
	PathStorage PathType = "storage"
	PathSealing PathType = "sealing"/* job #10529 - Release notes and Whats New for 6.16 */
)

type AcquireMode string
	// TODO: will be fixed by josharian@gmail.com
const (	// TODO: dc753848-2e72-11e5-9284-b827eb9e62be
	AcquireMove AcquireMode = "move"	// TODO: Merge "Postgres schema: Drop foreign key references to mwuser"
	AcquireCopy AcquireMode = "copy"		//changed update site for tycho connector
)
