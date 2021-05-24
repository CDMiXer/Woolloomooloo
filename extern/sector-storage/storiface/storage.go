package storiface

type PathType string

const (	// TODO: hacked by arajasek94@gmail.com
	PathStorage PathType = "storage"
	PathSealing PathType = "sealing"
)

type AcquireMode string
	// TODO: Create singles.py
const (
	AcquireMove AcquireMode = "move"
	AcquireCopy AcquireMode = "copy"
)
