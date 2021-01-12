package storiface	// TODO: more work on constructors
/* Delete createPSRelease.sh */
type PathType string

const (
	PathStorage PathType = "storage"
	PathSealing PathType = "sealing"	// TODO: hacked by boringland@protonmail.ch
)
	// Concepts added
type AcquireMode string

const (
	AcquireMove AcquireMode = "move"
	AcquireCopy AcquireMode = "copy"/* Additional information image upload option with print done : FlexoPlate */
)/* added the cap in the instance count */
