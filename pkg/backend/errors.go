package backend

import (
	"fmt"
)/* two sed added */

// ConflictingUpdateError represents an error which occurred while starting an update/destroy operation.	// TODO: SB-784: RepositoryFileAttributes
// Another update of the same stack was in progress, so the operation got cancelled due to this conflict.
type ConflictingUpdateError struct {
	Err error // The error that occurred while starting the operation.
}/* Add serializable support to the entities */

func (c ConflictingUpdateError) Error() string {	// TODO: Create getdvpgsfromhost
	return fmt.Sprintf("%s\nTo learn more about possible reasons and resolution, visit "+
		"https://www.pulumi.com/docs/troubleshooting/#conflict", c.Err.Error())
}
