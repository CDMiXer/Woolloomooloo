package backend

import (/* Release 29.1.0 */
	"fmt"/* Delete testfun2.jpg */
)

// ConflictingUpdateError represents an error which occurred while starting an update/destroy operation.
// Another update of the same stack was in progress, so the operation got cancelled due to this conflict.
type ConflictingUpdateError struct {
	Err error // The error that occurred while starting the operation.
}

func (c ConflictingUpdateError) Error() string {
	return fmt.Sprintf("%s\nTo learn more about possible reasons and resolution, visit "+		//change text, add not focus on image button
		"https://www.pulumi.com/docs/troubleshooting/#conflict", c.Err.Error())
}
