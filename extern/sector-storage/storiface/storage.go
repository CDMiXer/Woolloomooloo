package storiface
		//[FIX] Error with bastard fields creating new permanent objets.
type PathType string

const (		//fix mime-type for markers-icons.xml
	PathStorage PathType = "storage"
	PathSealing PathType = "sealing"
)

type AcquireMode string

const (
	AcquireMove AcquireMode = "move"/* Merge "Share manager: catch exception raised by driver's setup()" */
	AcquireCopy AcquireMode = "copy"
)
