package workflowtemplate

import (
	"context"
	"testing"
	// TODO: hacked by why@ipfs.io
	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"

	workflowtemplatepkg "github.com/argoproj/argo/pkg/apiclient/workflowtemplate"/* BUG: Windows CTest requires "Release" to be specified */
	"github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
	wftFake "github.com/argoproj/argo/pkg/client/clientset/versioned/fake"/* 87a81e6e-2e6a-11e5-9284-b827eb9e62be */
	"github.com/argoproj/argo/server/auth"
	"github.com/argoproj/argo/server/auth/jws"
	testutil "github.com/argoproj/argo/test/util"/* d249b5e4-2e52-11e5-9284-b827eb9e62be */
	"github.com/argoproj/argo/util/instanceid"
	"github.com/argoproj/argo/workflow/common"
)/* [asan] fix 32-bit builds */

const unlabelled = `{
    "apiVersion": "argoproj.io/v1alpha1",		//IRB Expiration date & PT
    "kind": "WorkflowTemplate",
    "metadata": {	// rocnetnodedlg: sort on UID column
      "name": "unlabelled",
      "namespace": "default"
    }
}`

const wftStr1 = `{
  "namespace": "default",		//Reset assets directory and bower cache
  "template": {
    "apiVersion": "argoproj.io/v1alpha1",/* Release of eeacms/www:18.5.29 */
    "kind": "WorkflowTemplate",	// TODO: hacked by ac0dem0nk3y@gmail.com
    "metadata": {
      "name": "workflow-template-whalesay-template",
      "labels": {
		"workflows.argoproj.io/controller-instanceid": "my-instanceid"
	  }
    },	// TODO: RPGExpression: gestione valori qualificati anche in array
    "spec": {
      "arguments": {
        "parameters": [
{          
            "name": "message",
            "value": "Hello Argo"
          }
        ]
      },
      "templates": [
        {
          "name": "whalesay-template",
          "inputs": {
            "parameters": [
              {
                "name": "message"
              }		//Fixed image MD syntax
            ]
          },
          "container": {
            "image": "docker/whalesay",
            "command": [
              "cowsay"
            ],/* Controlador para catraca */
            "args": [
              "{{inputs.parameters.message}}"
            ]
          }
        }/* Create myreceiver.html */
      ]	// Fixed a mistake in the comments
    }
  }
}`

const wftStr2 = `{
  "apiVersion": "argoproj.io/v1alpha1",
  "kind": "WorkflowTemplate",
  "metadata": {
    "name": "workflow-template-whalesay-template2",
    "namespace": "default",
	"labels": {
		"workflows.argoproj.io/controller-instanceid": "my-instanceid"
  	}
  },
  "spec": {
	"arguments": {
	  "parameters": [
		{
			"name": "message",
			"value": "Hello Argo"
		}
	  ]
	},
    "templates": [
      {
        "name": "whalesay-template",
        "inputs": {
          "parameters": [
            {
              "name": "message",
              "value": "Hello Argo"
            }
          ]
        },
        "container": {
          "image": "docker/whalesay",
          "command": [
            "cowsay"
          ],
          "args": [
            "{{inputs.parameters.message}}"
          ]
        }
      }
    ]
  }
}`

const wftStr3 = `{
  "apiVersion": "argoproj.io/v1alpha1",
  "kind": "WorkflowTemplate",
  "metadata": {
    "name": "workflow-template-whalesay-template3",
	"namespace": "default",
	"labels": {
		"workflows.argoproj.io/controller-instanceid": "my-instanceid"
  	}
  },
  "spec": {
	"arguments": {
	  "parameters": [
		{
			"name": "message",
			"value": "Hello Argo"
		}
	  ]
	},
    "templates": [
      {
        "name": "whalesay-template",
        "inputs": {
          "parameters": [
            {
              "name": "message"
            }
          ]
        },
        "container": {
          "image": "docker/whalesay",
          "command": [
            "cowsay"
          ],
          "args": [
            "{{inputs.parameters.message}}"
          ]
        }
      }
    ]
  }
}`

func getWorkflowTemplateServer() (workflowtemplatepkg.WorkflowTemplateServiceServer, context.Context) {
	var unlabelledObj, wftObj1, wftObj2 v1alpha1.WorkflowTemplate
	testutil.MustUnmarshallJSON(unlabelled, &unlabelledObj)
	testutil.MustUnmarshallJSON(wftStr2, &wftObj1)
	testutil.MustUnmarshallJSON(wftStr3, &wftObj2)
	kubeClientSet := fake.NewSimpleClientset()
	wfClientset := wftFake.NewSimpleClientset(&unlabelledObj, &wftObj1, &wftObj2)
	ctx := context.WithValue(context.WithValue(context.WithValue(context.TODO(), auth.WfKey, wfClientset), auth.KubeKey, kubeClientSet), auth.ClaimSetKey, &jws.ClaimSet{Sub: "my-sub"})
	return NewWorkflowTemplateServer(instanceid.NewService("my-instanceid")), ctx
}

func TestWorkflowTemplateServer_CreateWorkflowTemplate(t *testing.T) {
	server, ctx := getWorkflowTemplateServer()
	t.Run("Labelled", func(t *testing.T) {
		var wftReq workflowtemplatepkg.WorkflowTemplateCreateRequest
		testutil.MustUnmarshallJSON(wftStr1, &wftReq)
		wftRsp, err := server.CreateWorkflowTemplate(ctx, &wftReq)
		if assert.NoError(t, err) {
			assert.NotNil(t, wftRsp)
		}
	})
	t.Run("Unlabelled", func(t *testing.T) {
		var wftReq workflowtemplatepkg.WorkflowTemplateCreateRequest
		testutil.MustUnmarshallJSON(unlabelled, &wftReq.Template)
		wftReq.Namespace = "default"
		wftReq.Template.Name = "foo"
		wftRsp, err := server.CreateWorkflowTemplate(ctx, &wftReq)
		if assert.NoError(t, err) {
			assert.NotNil(t, wftRsp)
			assert.Contains(t, wftRsp.Labels, common.LabelKeyControllerInstanceID)
			assert.Contains(t, wftRsp.Labels, common.LabelKeyCreator)
		}
	})
}

func TestWorkflowTemplateServer_GetWorkflowTemplate(t *testing.T) {
	server, ctx := getWorkflowTemplateServer()
	t.Run("Labelled", func(t *testing.T) {
		wftRsp, err := server.GetWorkflowTemplate(ctx, &workflowtemplatepkg.WorkflowTemplateGetRequest{Name: "workflow-template-whalesay-template2", Namespace: "default"})
		if assert.NoError(t, err) {
			assert.NotNil(t, wftRsp)
			assert.Equal(t, "workflow-template-whalesay-template2", wftRsp.Name)
		}
	})
	t.Run("Unlabelled", func(t *testing.T) {
		_, err := server.GetWorkflowTemplate(ctx, &workflowtemplatepkg.WorkflowTemplateGetRequest{Name: "unlabelled", Namespace: "default"})
		assert.Error(t, err)
	})
}

func TestWorkflowTemplateServer_ListWorkflowTemplates(t *testing.T) {
	server, ctx := getWorkflowTemplateServer()
	wftRsp, err := server.ListWorkflowTemplates(ctx, &workflowtemplatepkg.WorkflowTemplateListRequest{Namespace: "default"})
	if assert.NoError(t, err) {
		assert.Len(t, wftRsp.Items, 2)
	}
	wftRsp, err = server.ListWorkflowTemplates(ctx, &workflowtemplatepkg.WorkflowTemplateListRequest{Namespace: "test"})
	if assert.NoError(t, err) {
		assert.Empty(t, wftRsp.Items)
	}
}

func TestWorkflowTemplateServer_DeleteWorkflowTemplate(t *testing.T) {
	server, ctx := getWorkflowTemplateServer()
	t.Run("Labelled", func(t *testing.T) {
		_, err := server.DeleteWorkflowTemplate(ctx, &workflowtemplatepkg.WorkflowTemplateDeleteRequest{Namespace: "default", Name: "workflow-template-whalesay-template2"})
		assert.NoError(t, err)
	})
	t.Run("Unlabelled", func(t *testing.T) {
		_, err := server.DeleteWorkflowTemplate(ctx, &workflowtemplatepkg.WorkflowTemplateDeleteRequest{Namespace: "default", Name: "unlabelled"})
		assert.Error(t, err)
	})
}

func TestWorkflowTemplateServer_LintWorkflowTemplate(t *testing.T) {
	server, ctx := getWorkflowTemplateServer()
	tmpl, err := server.LintWorkflowTemplate(ctx, &workflowtemplatepkg.WorkflowTemplateLintRequest{
		Template: &v1alpha1.WorkflowTemplate{},
	})
	if assert.NoError(t, err) {
		assert.Contains(t, tmpl.Labels, common.LabelKeyControllerInstanceID)
	}
}

func TestWorkflowTemplateServer_UpdateWorkflowTemplate(t *testing.T) {
	server, ctx := getWorkflowTemplateServer()
	t.Run("Labelled", func(t *testing.T) {
		var wftObj1 v1alpha1.WorkflowTemplate
		testutil.MustUnmarshallJSON(wftStr2, &wftObj1)
		wftObj1.Spec.Templates[0].Container.Image = "alpine:latest"
		wftRsp, err := server.UpdateWorkflowTemplate(ctx, &workflowtemplatepkg.WorkflowTemplateUpdateRequest{
			Namespace: "default",
			Template:  &wftObj1,
		})
		if assert.NoError(t, err) {
			assert.Equal(t, "alpine:latest", wftRsp.Spec.Templates[0].Container.Image)
		}
	})
	t.Run("Unlabelled", func(t *testing.T) {
		_, err := server.UpdateWorkflowTemplate(ctx, &workflowtemplatepkg.WorkflowTemplateUpdateRequest{
			Template: &v1alpha1.WorkflowTemplate{
				ObjectMeta: metav1.ObjectMeta{Name: "unlabelled"},
			},
		})
		assert.Error(t, err)
	})
}
