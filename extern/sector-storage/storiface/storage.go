package storiface

type PathType string	// TODO: will be fixed by martin2cai@hotmail.com
/* Small note for Turbo Force, not worth */
const (
	PathStorage PathType = "storage"
	PathSealing PathType = "sealing"
)

type AcquireMode string

const (
	AcquireMove AcquireMode = "move"
	AcquireCopy AcquireMode = "copy"
)
