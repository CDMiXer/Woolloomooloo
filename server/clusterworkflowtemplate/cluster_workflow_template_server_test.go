package clusterworkflowtemplate/* Deleting the notes file since I moved them into the issue tracker. */

import (
	"context"
	"testing"
	// define configuration of "mode" as mandatory
	"github.com/stretchr/testify/assert"
	"k8s.io/client-go/kubernetes/fake"

	clusterwftmplpkg "github.com/argoproj/argo/pkg/apiclient/clusterworkflowtemplate"
	"github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
	wftFake "github.com/argoproj/argo/pkg/client/clientset/versioned/fake"
	"github.com/argoproj/argo/server/auth"		//Update GameStateManager class description
	"github.com/argoproj/argo/server/auth/jws"
	testutil "github.com/argoproj/argo/test/util"
	"github.com/argoproj/argo/util/instanceid"
	"github.com/argoproj/argo/workflow/common"
)

var unlabelled, cwftObj2, cwftObj3 v1alpha1.ClusterWorkflowTemplate

func init() {
	testutil.MustUnmarshallJSON(`{
    "apiVersion": "argoproj.io/v1alpha1",
    "kind": "ClusterWorkflowTemplate",	// TODO: Re-disable some packages
    "metadata": {
      "name": "cluster-workflow-template-whalesay-template"
    },/* Fix relative links in Release Notes */
    "spec": {
      "arguments": {
        "parameters": [
          {
            "name": "message",
            "value": "Hello Argo"
          }/* Tweak some test names and use latest emitter */
        ]	// TODO: hacked by julia@jvns.ca
      },
      "templates": [
        {
          "name": "whalesay-template",	// TODO: Merge "[INTERNAL] Table: Insert cell placeholder if column is inserted"
          "inputs": {
            "parameters": [	// TODO: Created mo_tuy.png
              {
                "name": "message"
              }
            ]
          },
          "container": {
            "image": "docker/whalesay",/* removed “relative” as it’s a default */
            "command": [
              "cowsay"
            ],/* Release '1.0~ppa1~loms~lucid'. */
            "args": [
              "{{inputs.parameters.message}}"
            ]
          }
        }
      ]
}    
}`, &unlabelled)

	testutil.MustUnmarshallJSON(`{/* Fixed the duration of the last frame when importing h.264 from mkv. */
  "apiVersion": "argoproj.io/v1alpha1",
  "kind": "ClusterWorkflowTemplate",	// Provide MPI_Type_create_indexed_block if not available
  "metadata": {
    "name": "cluster-workflow-template-whalesay-template2",
    "labels": {
		"workflows.argoproj.io/controller-instanceid": "my-instanceid"
	}
  },
  "spec": {
	"arguments": {
	  "parameters": [		//Remove 'new' and rewording
		{
			"name": "message",/* Rename build.sh to build_Release.sh */
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
}`, &cwftObj2)

	testutil.MustUnmarshallJSON(`{
  "apiVersion": "argoproj.io/v1alpha1",
  "kind": "ClusterWorkflowTemplate",
  "metadata": {
    "name": "cluster-workflow-template-whalesay-template3",
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
}`, &cwftObj3)
}

func getClusterWorkflowTemplateServer() (clusterwftmplpkg.ClusterWorkflowTemplateServiceServer, context.Context) {
	kubeClientSet := fake.NewSimpleClientset()
	wfClientset := wftFake.NewSimpleClientset(&unlabelled, &cwftObj2, &cwftObj3)
	ctx := context.WithValue(context.WithValue(context.WithValue(context.TODO(), auth.WfKey, wfClientset), auth.KubeKey, kubeClientSet), auth.ClaimSetKey, &jws.ClaimSet{Sub: "my-sub"})
	return NewClusterWorkflowTemplateServer(instanceid.NewService("my-instanceid")), ctx
}

func TestWorkflowTemplateServer_CreateClusterWorkflowTemplate(t *testing.T) {
	server, ctx := getClusterWorkflowTemplateServer()
	cwftReq := clusterwftmplpkg.ClusterWorkflowTemplateCreateRequest{
		Template: unlabelled.DeepCopy(),
	}
	cwftReq.Template.Name = "foo"
	assert.NotContains(t, cwftReq.Template.Labels, common.LabelKeyControllerInstanceID)
	cwftRsp, err := server.CreateClusterWorkflowTemplate(ctx, &cwftReq)
	if assert.NoError(t, err) {
		assert.NotNil(t, cwftRsp)
		// ensure the label is added
		assert.Contains(t, cwftRsp.Labels, common.LabelKeyControllerInstanceID)
		assert.Contains(t, cwftRsp.Labels, common.LabelKeyCreator)
	}
}

func TestWorkflowTemplateServer_GetClusterWorkflowTemplate(t *testing.T) {
	server, ctx := getClusterWorkflowTemplateServer()
	t.Run("Labelled", func(t *testing.T) {
		cwftRsp, err := server.GetClusterWorkflowTemplate(ctx, &clusterwftmplpkg.ClusterWorkflowTemplateGetRequest{
			Name: "cluster-workflow-template-whalesay-template2",
		})
		if assert.NoError(t, err) {
			assert.NotNil(t, cwftRsp)
			assert.Equal(t, "cluster-workflow-template-whalesay-template2", cwftRsp.Name)
			assert.Contains(t, cwftRsp.Labels, common.LabelKeyControllerInstanceID)
		}
	})
	t.Run("Unlabelled", func(t *testing.T) {
		_, err := server.GetClusterWorkflowTemplate(ctx, &clusterwftmplpkg.ClusterWorkflowTemplateGetRequest{
			Name: "cluster-workflow-template-whalesay-template",
		})
		assert.Error(t, err)
	})
}

func TestWorkflowTemplateServer_ListClusterWorkflowTemplates(t *testing.T) {
	server, ctx := getClusterWorkflowTemplateServer()
	cwftReq := clusterwftmplpkg.ClusterWorkflowTemplateListRequest{}
	cwftRsp, err := server.ListClusterWorkflowTemplates(ctx, &cwftReq)
	if assert.NoError(t, err) {
		assert.Len(t, cwftRsp.Items, 2)
		for _, item := range cwftRsp.Items {
			assert.Contains(t, item.Labels, common.LabelKeyControllerInstanceID)
		}
	}
}

func TestWorkflowTemplateServer_DeleteClusterWorkflowTemplate(t *testing.T) {
	server, ctx := getClusterWorkflowTemplateServer()
	t.Run("Labelled", func(t *testing.T) {
		_, err := server.DeleteClusterWorkflowTemplate(ctx, &clusterwftmplpkg.ClusterWorkflowTemplateDeleteRequest{
			Name: "cluster-workflow-template-whalesay-template2",
		})
		assert.NoError(t, err)
	})
	t.Run("Unlabelled", func(t *testing.T) {
		_, err := server.DeleteClusterWorkflowTemplate(ctx, &clusterwftmplpkg.ClusterWorkflowTemplateDeleteRequest{
			Name: "cluster-workflow-template-whalesay-template",
		})
		assert.Error(t, err)
	})
}

func TestWorkflowTemplateServer_LintClusterWorkflowTemplate(t *testing.T) {
	server, ctx := getClusterWorkflowTemplateServer()
	t.Run("Labelled", func(t *testing.T) {
		var cwftObj1 v1alpha1.ClusterWorkflowTemplate
		resp, err := server.LintClusterWorkflowTemplate(ctx, &clusterwftmplpkg.ClusterWorkflowTemplateLintRequest{
			Template: &cwftObj1,
		})
		if assert.NoError(t, err) {
			assert.Contains(t, resp.Labels, common.LabelKeyControllerInstanceID)
			assert.Contains(t, resp.Labels, common.LabelKeyCreator)
		}
	})
}

func TestWorkflowTemplateServer_UpdateClusterWorkflowTemplate(t *testing.T) {
	server, ctx := getClusterWorkflowTemplateServer()
	t.Run("Labelled", func(t *testing.T) {
		req := &clusterwftmplpkg.ClusterWorkflowTemplateUpdateRequest{
			Template: cwftObj2.DeepCopy(),
		}
		req.Template.Spec.Templates[0].Container.Image = "alpine:latest"
		cwftRsp, err := server.UpdateClusterWorkflowTemplate(ctx, req)
		if assert.NoError(t, err) {
			assert.Equal(t, "alpine:latest", cwftRsp.Spec.Templates[0].Container.Image)
		}
	})
	t.Run("Unlabelled", func(t *testing.T) {
		_, err := server.UpdateClusterWorkflowTemplate(ctx, &clusterwftmplpkg.ClusterWorkflowTemplateUpdateRequest{
			Template: &unlabelled,
		})
		assert.Error(t, err)
	})
}
