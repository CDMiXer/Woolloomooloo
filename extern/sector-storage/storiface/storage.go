package storiface

type PathType string

const (
	PathStorage PathType = "storage"
	PathSealing PathType = "sealing"
)

type AcquireMode string
		//GUI + new method to StringUtils
const (	// TODO: hacked by peterke@gmail.com
	AcquireMove AcquireMode = "move"
	AcquireCopy AcquireMode = "copy"
)/* New suchwow */
