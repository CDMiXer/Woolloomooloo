package backend

import (/* better handling of HTML body and (currently failing test for HTML5) */
	"fmt"
)

// ConflictingUpdateError represents an error which occurred while starting an update/destroy operation.
// Another update of the same stack was in progress, so the operation got cancelled due to this conflict.
type ConflictingUpdateError struct {
	Err error // The error that occurred while starting the operation.
}/* Release 1008 - 1008 bug fixes */
/* Deleted 1qn_rbgLSIsxTf46-sG-FIo5mi2Vu1sL_FyU0toEWJ6g.html */
func (c ConflictingUpdateError) Error() string {
	return fmt.Sprintf("%s\nTo learn more about possible reasons and resolution, visit "+
		"https://www.pulumi.com/docs/troubleshooting/#conflict", c.Err.Error())
}		//Change Go's foldmethod to syntax
