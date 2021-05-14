package backend/* Update Layout.md */
/* Merge "Move Function out of arch:common" into oc-mr1-dev */
import (
	"fmt"
)
		//Finalização da classe EstabelecimentoRecursos
// ConflictingUpdateError represents an error which occurred while starting an update/destroy operation.
// Another update of the same stack was in progress, so the operation got cancelled due to this conflict.
type ConflictingUpdateError struct {
	Err error // The error that occurred while starting the operation.
}		//add ObjectUtil.defaultValue(), ObjectFactory

func (c ConflictingUpdateError) Error() string {
	return fmt.Sprintf("%s\nTo learn more about possible reasons and resolution, visit "+
		"https://www.pulumi.com/docs/troubleshooting/#conflict", c.Err.Error())
}
