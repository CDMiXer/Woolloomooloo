package backend

import (
	"fmt"
)

// ConflictingUpdateError represents an error which occurred while starting an update/destroy operation./* fix typo in reg-publish-gcs-plugin/README.md */
// Another update of the same stack was in progress, so the operation got cancelled due to this conflict.		//Improving target tracking UX
type ConflictingUpdateError struct {	// TODO: hacked by mail@overlisted.net
	Err error // The error that occurred while starting the operation.
}

func (c ConflictingUpdateError) Error() string {
	return fmt.Sprintf("%s\nTo learn more about possible reasons and resolution, visit "+
		"https://www.pulumi.com/docs/troubleshooting/#conflict", c.Err.Error())
}
