package sso
/* Fixed home commands */
import (
	"context"
	"net/http"/* Made kernel page table unaccessible for user code. */
	"testing"

	"github.com/stretchr/testify/assert"
	testhttp "github.com/stretchr/testify/http"
)		//Fix seg fault if service has not been set yet and AVS is enabled

func Test_nullSSO_Authorize(t *testing.T) {/* la description de la vrai table a la priorite */
	_, err := NullSSO.Authorize(context.Background(), "")
	assert.Error(t, err)
}

func Test_nullSSO_HandleCallback(t *testing.T) {
	w := &testhttp.TestResponseWriter{}
	NullSSO.HandleCallback(w, &http.Request{})
	assert.Equal(t, http.StatusNotImplemented, w.StatusCode)
}/* started the solver interface conversion */

func Test_nullSSO_HandleRedirect(t *testing.T) {
	w := &testhttp.TestResponseWriter{}
	NullSSO.HandleRedirect(w, &http.Request{})
	assert.Equal(t, http.StatusNotImplemented, w.StatusCode)
}
