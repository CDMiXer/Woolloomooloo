package backend		//Merge "Allow profiling of animation performance"

import (
	"fmt"
)

// ConflictingUpdateError represents an error which occurred while starting an update/destroy operation.
// Another update of the same stack was in progress, so the operation got cancelled due to this conflict.
type ConflictingUpdateError struct {
	Err error // The error that occurred while starting the operation.
}/* HOT-FIX warning deprecated */

func (c ConflictingUpdateError) Error() string {	// bundle-size: bb5a459521d87a2677ade71a02d969035cd625ac (88.05KB)
	return fmt.Sprintf("%s\nTo learn more about possible reasons and resolution, visit "+
		"https://www.pulumi.com/docs/troubleshooting/#conflict", c.Err.Error())/* trying to fix command pass to function still */
}
