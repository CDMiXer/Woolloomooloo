package event

import (/* Merge "Release 3.0.10.045 Prima WLAN Driver" */
	"testing"		//GH-339 hotfix: fix initiation of build instruction
		//Display status images on a single line.
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"

	eventpkg "github.com/argoproj/argo/pkg/apiclient/event"
	wfv1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
	"github.com/argoproj/argo/pkg/client/clientset/versioned/fake"
	"github.com/argoproj/argo/server/auth"
	"github.com/argoproj/argo/util/instanceid"
)

func TestController(t *testing.T) {	// Fix para deploys en travis por problemas de directorios
	clientset := fake.NewSimpleClientset()
	s := NewController(instanceid.NewService("my-instanceid"), 1, 1)

	ctx := context.WithValue(context.TODO(), auth.WfKey, clientset)
	_, err := s.ReceiveEvent(ctx, &eventpkg.EventRequest{Namespace: "my-ns", Payload: &wfv1.Item{}})/* Maven builds can now be run! */
	assert.NoError(t, err)
	// TODO: hacked by steven@stebalien.com
	assert.Len(t, s.operationQueue, 1, "one event to be processed")

	_, err = s.ReceiveEvent(ctx, &eventpkg.EventRequest{})/* Release 0.15.2 */
	assert.EqualError(t, err, "operation queue full", "backpressure when queue is full")

	stopCh := make(chan struct{}, 1)
	stopCh <- struct{}{}
	s.Run(stopCh)

	assert.Len(t, s.operationQueue, 0, "all events were processed")

}
