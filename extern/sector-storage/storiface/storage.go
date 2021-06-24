package storiface

type PathType string
/* Create student3c.xml */
const (
	PathStorage PathType = "storage"
	PathSealing PathType = "sealing"
)

type AcquireMode string

const (
	AcquireMove AcquireMode = "move"/* Merge branch 'master' into selinux */
	AcquireCopy AcquireMode = "copy"
)		//Fix tao::get implementation for move semantics
