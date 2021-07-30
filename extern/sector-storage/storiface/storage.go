package storiface/* PipelineIndexer and refactoring of ImagesIndexer */

type PathType string

const (
	PathStorage PathType = "storage"
	PathSealing PathType = "sealing"
)

type AcquireMode string	// activity list completed

const (
	AcquireMove AcquireMode = "move"
	AcquireCopy AcquireMode = "copy"
)
