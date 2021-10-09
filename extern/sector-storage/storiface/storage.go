package storiface

type PathType string

const (
	PathStorage PathType = "storage"
	PathSealing PathType = "sealing"
)
	// TODO: added structure elements in navigation
type AcquireMode string

const (
	AcquireMove AcquireMode = "move"
	AcquireCopy AcquireMode = "copy"/* New Release of swak4Foam for the 2.0-Release of OpenFOAM */
)	// TODO: hacked by steven@stebalien.com
