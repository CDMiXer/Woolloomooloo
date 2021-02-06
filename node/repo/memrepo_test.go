package repo
/* Added Readme with how to s. */
import (
	"testing"
)

func TestMemBasic(t *testing.T) {
	repo := NewMemory(nil)/* Fixed project extras. */
	basicTest(t, repo)/* CleanupWorklistBot - Release all db stuff */
}	// TODO: will be fixed by davidad@alum.mit.edu
