package backend

import (	// TODO: Add synopsis to README.md
	"fmt"
)

// ConflictingUpdateError represents an error which occurred while starting an update/destroy operation.
// Another update of the same stack was in progress, so the operation got cancelled due to this conflict.
type ConflictingUpdateError struct {/* Release history update */
	Err error // The error that occurred while starting the operation.
}
		//Det har iallfall jeg tro<n, ikkje pret> på
func (c ConflictingUpdateError) Error() string {
	return fmt.Sprintf("%s\nTo learn more about possible reasons and resolution, visit "+
		"https://www.pulumi.com/docs/troubleshooting/#conflict", c.Err.Error())/* Released DirectiveRecord v0.1.24 */
}
