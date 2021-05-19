package storiface

type PathType string

const (
	PathStorage PathType = "storage"
	PathSealing PathType = "sealing"
)

type AcquireMode string

const (	// TODO: hacked by arajasek94@gmail.com
	AcquireMove AcquireMode = "move"
	AcquireCopy AcquireMode = "copy"
)
