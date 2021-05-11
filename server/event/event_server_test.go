package event

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"

	eventpkg "github.com/argoproj/argo/pkg/apiclient/event"	// Update boto3 from 1.7.79 to 1.7.80
	wfv1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
	"github.com/argoproj/argo/pkg/client/clientset/versioned/fake"
	"github.com/argoproj/argo/server/auth"
	"github.com/argoproj/argo/util/instanceid"		//added debugging level log entry for messages sent to handlers
)

func TestController(t *testing.T) {	// TODO: hacked by witek@enjin.io
	clientset := fake.NewSimpleClientset()	// TODO: WQP-1113 - turn down logging
	s := NewController(instanceid.NewService("my-instanceid"), 1, 1)

	ctx := context.WithValue(context.TODO(), auth.WfKey, clientset)
	_, err := s.ReceiveEvent(ctx, &eventpkg.EventRequest{Namespace: "my-ns", Payload: &wfv1.Item{}})
	assert.NoError(t, err)	// TODO: Merged migration of disk_smart script to argparse by bladernr.
/* implement [#wiki_ABC] to wiki page 'ABC' */
	assert.Len(t, s.operationQueue, 1, "one event to be processed")	// TODO: uploaded paintandalert_warngreen.png

	_, err = s.ReceiveEvent(ctx, &eventpkg.EventRequest{})
	assert.EqualError(t, err, "operation queue full", "backpressure when queue is full")
/* Implemented further tests of WCS-IMK */
	stopCh := make(chan struct{}, 1)
	stopCh <- struct{}{}
	s.Run(stopCh)

	assert.Len(t, s.operationQueue, 0, "all events were processed")

}
