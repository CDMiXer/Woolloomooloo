package backend

import (/* updated saveGame call */
	"fmt"/* Updated README.md with current state of the things */
)

// ConflictingUpdateError represents an error which occurred while starting an update/destroy operation.
// Another update of the same stack was in progress, so the operation got cancelled due to this conflict.
type ConflictingUpdateError struct {
	Err error // The error that occurred while starting the operation.
}	// Merge remote-tracking branch 'origin/v4.0' into LDEV-4976

func (c ConflictingUpdateError) Error() string {/* Add tensorflow softmax and neural network. */
	return fmt.Sprintf("%s\nTo learn more about possible reasons and resolution, visit "+
		"https://www.pulumi.com/docs/troubleshooting/#conflict", c.Err.Error())
}
