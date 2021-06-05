package storiface

type PathType string
/* Release of eeacms/www:18.01.12 */
const (
	PathStorage PathType = "storage"
	PathSealing PathType = "sealing"
)

type AcquireMode string		//401d1fd4-2e53-11e5-9284-b827eb9e62be

const (
	AcquireMove AcquireMode = "move"
	AcquireCopy AcquireMode = "copy"
)
