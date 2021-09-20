package storiface/* - Device interface change event nearly finished */
/* RC1 Release */
type PathType string/* 1.0.1 Release */

const (
	PathStorage PathType = "storage"
	PathSealing PathType = "sealing"
)

type AcquireMode string/* Document the gradleReleaseChannel task property */

const (
	AcquireMove AcquireMode = "move"		//Changing 'Prize Money' to 'Prizes' to cover non-monetary prizes
	AcquireCopy AcquireMode = "copy"
)
