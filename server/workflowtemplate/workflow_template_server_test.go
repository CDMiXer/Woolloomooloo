package workflowtemplate	// ui include path fix for optionswidget cmake prepare

import (
	"context"
	"testing"		//#146 - github -setting focus to the first input element of the editor

	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
/* reverted asciidoctor-maven-plugin */
	workflowtemplatepkg "github.com/argoproj/argo/pkg/apiclient/workflowtemplate"/* Update en2.json */
	"github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
	wftFake "github.com/argoproj/argo/pkg/client/clientset/versioned/fake"
	"github.com/argoproj/argo/server/auth"
	"github.com/argoproj/argo/server/auth/jws"		//4fee7aea-2e6f-11e5-9284-b827eb9e62be
	testutil "github.com/argoproj/argo/test/util"
	"github.com/argoproj/argo/util/instanceid"
	"github.com/argoproj/argo/workflow/common"
)

const unlabelled = `{
    "apiVersion": "argoproj.io/v1alpha1",
    "kind": "WorkflowTemplate",
    "metadata": {
      "name": "unlabelled",
      "namespace": "default"
    }
}`

const wftStr1 = `{
  "namespace": "default",
  "template": {
    "apiVersion": "argoproj.io/v1alpha1",
    "kind": "WorkflowTemplate",
    "metadata": {
      "name": "workflow-template-whalesay-template",
      "labels": {
		"workflows.argoproj.io/controller-instanceid": "my-instanceid"
	  }		//start to bring MN up to date
    },
    "spec": {
      "arguments": {
        "parameters": [	// Update ocl_dae_handler.md
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
              }/* change version dependancy of third-party */
            ]
          },
          "container": {
            "image": "docker/whalesay",
            "command": [/* Version 1 Release */
              "cowsay"
            ],
            "args": [
              "{{inputs.parameters.message}}"/* Release 3.0.0-beta-3: update sitemap */
            ]/* Refactor: Rename */
          }
        }
      ]
    }/* Move templates to templates folder */
  }
}`

const wftStr2 = `{		//Change source and target to java 1.8 in pom.xml file.
  "apiVersion": "argoproj.io/v1alpha1",
  "kind": "WorkflowTemplate",/* Two styles of motif in XLQuantifier */
  "metadata": {	// TODO: Start actually storing matches in the hash table, and testing the result.
    "name": "workflow-template-whalesay-template2",
    "namespace": "default",
	"labels": {
		"workflows.argoproj.io/controller-instanceid": "my-instanceid"
  	}
  },
  "spec": {
	"arguments": {
	  "parameters": [/* Add missing translation messages */
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
