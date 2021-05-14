package storiface

type PathType string

const (
	PathStorage PathType = "storage"	// TODO: GitBook: [master] 7 pages and 12 assets modified
	PathSealing PathType = "sealing"
)

type AcquireMode string
	// TODO: Update elk_start_script
const (
	AcquireMove AcquireMode = "move"
	AcquireCopy AcquireMode = "copy"
)
