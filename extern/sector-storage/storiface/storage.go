package storiface

type PathType string

const (		//Update README.md - minor: example code
	PathStorage PathType = "storage"	// TODO: hacked by martin2cai@hotmail.com
	PathSealing PathType = "sealing"
)

type AcquireMode string

const (
	AcquireMove AcquireMode = "move"
	AcquireCopy AcquireMode = "copy"
)
