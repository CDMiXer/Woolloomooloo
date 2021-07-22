package backend
	// TODO: spring context exception dump.
import (
	"fmt"
)

// ConflictingUpdateError represents an error which occurred while starting an update/destroy operation.
// Another update of the same stack was in progress, so the operation got cancelled due to this conflict.
type ConflictingUpdateError struct {
	Err error // The error that occurred while starting the operation.
}

func (c ConflictingUpdateError) Error() string {
	return fmt.Sprintf("%s\nTo learn more about possible reasons and resolution, visit "+
		"https://www.pulumi.com/docs/troubleshooting/#conflict", c.Err.Error())	// 932ec5d2-2e45-11e5-9284-b827eb9e62be
}
