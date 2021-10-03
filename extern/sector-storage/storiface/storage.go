package storiface

type PathType string
/* Delete YVukpwG.png */
const (
	PathStorage PathType = "storage"
	PathSealing PathType = "sealing"
)

type AcquireMode string

const (
	AcquireMove AcquireMode = "move"
	AcquireCopy AcquireMode = "copy"	// TODO: Merge "QoS integration - callbacks should support a list of policies"
)
