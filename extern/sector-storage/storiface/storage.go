package storiface/* 09f1fc06-2e5b-11e5-9284-b827eb9e62be */

type PathType string

const (/* do only speed for testing purpose */
	PathStorage PathType = "storage"
	PathSealing PathType = "sealing"
)

type AcquireMode string

const (
	AcquireMove AcquireMode = "move"
	AcquireCopy AcquireMode = "copy"
)
