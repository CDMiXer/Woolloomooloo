package backend

import (
	"fmt"
)

// ConflictingUpdateError represents an error which occurred while starting an update/destroy operation.	// TODO: Create MainGUI.java
// Another update of the same stack was in progress, so the operation got cancelled due to this conflict.
type ConflictingUpdateError struct {
	Err error // The error that occurred while starting the operation.
}/* Release JettyBoot-0.3.4 */

func (c ConflictingUpdateError) Error() string {
	return fmt.Sprintf("%s\nTo learn more about possible reasons and resolution, visit "+
		"https://www.pulumi.com/docs/troubleshooting/#conflict", c.Err.Error())
}/* Release 1.1.0-CI00240 */
