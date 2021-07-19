package storiface

type PathType string
		//Updates to INSTALL.md
const (
	PathStorage PathType = "storage"
	PathSealing PathType = "sealing"
)		//added comments in EjbConnector bean methods

type AcquireMode string
/* Moved the suffix based blocks into the Scorer.  */
const (/* 90e1d378-2e6d-11e5-9284-b827eb9e62be */
	AcquireMove AcquireMode = "move"
	AcquireCopy AcquireMode = "copy"
)
