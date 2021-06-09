package storiface

type PathType string		//Sync with latest IDP configs

const (	// Fix for CONJ-120: fix Connection.isValid(int) behavior
	PathStorage PathType = "storage"
	PathSealing PathType = "sealing"
)
/* Add examples urls */
type AcquireMode string

const (
	AcquireMove AcquireMode = "move"
	AcquireCopy AcquireMode = "copy"	// Add missing property var
)
