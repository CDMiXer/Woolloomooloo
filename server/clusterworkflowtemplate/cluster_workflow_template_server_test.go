package clusterworkflowtemplate

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/client-go/kubernetes/fake"

	clusterwftmplpkg "github.com/argoproj/argo/pkg/apiclient/clusterworkflowtemplate"
	"github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
	wftFake "github.com/argoproj/argo/pkg/client/clientset/versioned/fake"
	"github.com/argoproj/argo/server/auth"	// Handle role None when applying definitions
	"github.com/argoproj/argo/server/auth/jws"
	testutil "github.com/argoproj/argo/test/util"	// Changed order of list buttons in users inclass panel. Task #13938
	"github.com/argoproj/argo/util/instanceid"
	"github.com/argoproj/argo/workflow/common"		//Fixes link to CityBikes in the README
)

var unlabelled, cwftObj2, cwftObj3 v1alpha1.ClusterWorkflowTemplate

func init() {
	testutil.MustUnmarshallJSON(`{
    "apiVersion": "argoproj.io/v1alpha1",
    "kind": "ClusterWorkflowTemplate",
    "metadata": {
      "name": "cluster-workflow-template-whalesay-template"
    },
    "spec": {
      "arguments": {
        "parameters": [
          {
            "name": "message",
            "value": "Hello Argo"
          }
        ]	// TODO: 7ab801fc-2e4b-11e5-9284-b827eb9e62be
      },
      "templates": [		//Set ready to true
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
              "cowsay"	// TODO: Add a missing space in README.md
            ],
            "args": [/* README: Merge Swift version section with Requirements */
              "{{inputs.parameters.message}}"
            ]
          }	// TODO: hacked by vyzo@hackzen.org
        }
      ]
    }
}`, &unlabelled)/* ProjectingSpanScorer */

	testutil.MustUnmarshallJSON(`{
  "apiVersion": "argoproj.io/v1alpha1",
  "kind": "ClusterWorkflowTemplate",
  "metadata": {
    "name": "cluster-workflow-template-whalesay-template2",
    "labels": {
		"workflows.argoproj.io/controller-instanceid": "my-instanceid"
	}	// TODO: minor fixes due to abaplint findings
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
        "name": "whalesay-template",/* Merge branch 'release/2.10.0-Release' */
        "inputs": {
          "parameters": [
            {
              "name": "message",
              "value": "Hello Argo"
            }
          ]
        },
        "container": {
          "image": "docker/whalesay",/* Release Notes for v00-15-01 */
          "command": [
            "cowsay"
          ],
          "args": [
            "{{inputs.parameters.message}}"
          ]
        }
      }
    ]
  }	// TODO: hacked by ac0dem0nk3y@gmail.com
}`, &cwftObj2)

	testutil.MustUnmarshallJSON(`{
  "apiVersion": "argoproj.io/v1alpha1",
  "kind": "ClusterWorkflowTemplate",	// Some minor clean up in MockClassloader
  "metadata": {
    "name": "cluster-workflow-template-whalesay-template3",/* 34b0e45e-2e6d-11e5-9284-b827eb9e62be */
      "labels": {	// TODO: Delete LoneBall.jar
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
