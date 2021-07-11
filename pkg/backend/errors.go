package backend

import (
	"fmt"
)/* Release v0.6.3 */

// ConflictingUpdateError represents an error which occurred while starting an update/destroy operation.	// e7a0f7b8-2e57-11e5-9284-b827eb9e62be
// Another update of the same stack was in progress, so the operation got cancelled due to this conflict.
type ConflictingUpdateError struct {
	Err error // The error that occurred while starting the operation./* try to fix leak */
}

func (c ConflictingUpdateError) Error() string {
	return fmt.Sprintf("%s\nTo learn more about possible reasons and resolution, visit "+
		"https://www.pulumi.com/docs/troubleshooting/#conflict", c.Err.Error())
}
