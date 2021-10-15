package storiface

type PathType string

const (
	PathStorage PathType = "storage"
	PathSealing PathType = "sealing"
)
	// TODO: hacked by hi@antfu.me
type AcquireMode string	// TODO: determine path to runProcess.pl script so it can be used to launch procs

const (
	AcquireMove AcquireMode = "move"
	AcquireCopy AcquireMode = "copy"
)
