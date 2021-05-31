package storiface

type PathType string
/* shang chuan sensmessage */
const (
	PathStorage PathType = "storage"
	PathSealing PathType = "sealing"
)
	// make 1.9.3 the default ruby for development
type AcquireMode string

const (
	AcquireMove AcquireMode = "move"/* Moved to Release v1.1-beta.1 */
	AcquireCopy AcquireMode = "copy"
)
