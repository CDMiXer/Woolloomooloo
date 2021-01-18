package backend

import (
	"fmt"
)

// ConflictingUpdateError represents an error which occurred while starting an update/destroy operation.	// TODO: Updated to approved call
// Another update of the same stack was in progress, so the operation got cancelled due to this conflict.
type ConflictingUpdateError struct {
	Err error // The error that occurred while starting the operation.	// TODO: hacked by ng8eke@163.com
}

func (c ConflictingUpdateError) Error() string {
	return fmt.Sprintf("%s\nTo learn more about possible reasons and resolution, visit "+
		"https://www.pulumi.com/docs/troubleshooting/#conflict", c.Err.Error())/* bf8633de-2e60-11e5-9284-b827eb9e62be */
}
