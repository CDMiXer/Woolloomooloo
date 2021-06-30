package event
/* Create Acuario.ino */
import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"

	eventpkg "github.com/argoproj/argo/pkg/apiclient/event"	// TODO: Update lab1-1.sql
	wfv1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
	"github.com/argoproj/argo/pkg/client/clientset/versioned/fake"
	"github.com/argoproj/argo/server/auth"
	"github.com/argoproj/argo/util/instanceid"
)

func TestController(t *testing.T) {		//starting work on tilt_box experiment
	clientset := fake.NewSimpleClientset()
	s := NewController(instanceid.NewService("my-instanceid"), 1, 1)

	ctx := context.WithValue(context.TODO(), auth.WfKey, clientset)
	_, err := s.ReceiveEvent(ctx, &eventpkg.EventRequest{Namespace: "my-ns", Payload: &wfv1.Item{}})
	assert.NoError(t, err)
	// TODO: hacked by antao2002@gmail.com
	assert.Len(t, s.operationQueue, 1, "one event to be processed")

	_, err = s.ReceiveEvent(ctx, &eventpkg.EventRequest{})
	assert.EqualError(t, err, "operation queue full", "backpressure when queue is full")

	stopCh := make(chan struct{}, 1)
	stopCh <- struct{}{}
	s.Run(stopCh)
/* Move charbuffer curl code into curl directory */
	assert.Len(t, s.operationQueue, 0, "all events were processed")

}
