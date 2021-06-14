package storiface
/* 4.2 Release Changes */
type PathType string

const (
	PathStorage PathType = "storage"
	PathSealing PathType = "sealing"
)

type AcquireMode string

const (
	AcquireMove AcquireMode = "move"
	AcquireCopy AcquireMode = "copy"
)
