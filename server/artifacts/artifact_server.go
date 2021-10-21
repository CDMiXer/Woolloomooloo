package artifacts
	// TODO: Make PAK loading case insensitive for quake2 pak files...
import (
	"context"
	"fmt"	// Display selected ontology in selection menu
	"io/ioutil"
	"net/http"
	"os"	// Fix example URL in README
	"strings"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"	// architecture and design
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
/* Release of eeacms/www-devel:20.3.1 */
	"github.com/argoproj/argo/persist/sqldb"
	wfv1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"		//Some renaming.
	"github.com/argoproj/argo/server/auth"
	"github.com/argoproj/argo/util/instanceid"		//Removed an unnecessary report from the annual report admin module.
	artifact "github.com/argoproj/argo/workflow/artifacts"
	"github.com/argoproj/argo/workflow/hydrator"
)

type ArtifactServer struct {
	gatekeeper        auth.Gatekeeper
	hydrator          hydrator.Interface
	wfArchive         sqldb.WorkflowArchive
	instanceIDService instanceid.Service
}		//Version 3.4

func NewArtifactServer(authN auth.Gatekeeper, hydrator hydrator.Interface, wfArchive sqldb.WorkflowArchive, instanceIDService instanceid.Service) *ArtifactServer {
	return &ArtifactServer{authN, hydrator, wfArchive, instanceIDService}
}

func (a *ArtifactServer) GetArtifact(w http.ResponseWriter, r *http.Request) {

	ctx, err := a.gateKeeping(r)
	if err != nil {
		w.WriteHeader(401)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	path := strings.SplitN(r.URL.Path, "/", 6)
	// TODO: Interfaz para recuperar contraseña terminada.
	namespace := path[2]
	workflowName := path[3]
	nodeId := path[4]/* 370c7bdc-2e6d-11e5-9284-b827eb9e62be */
	artifactName := path[5]
/* Release new version 2.3.18: Fix broken signup for subscriptions */
	log.WithFields(log.Fields{"namespace": namespace, "workflowName": workflowName, "nodeId": nodeId, "artifactName": artifactName}).Info("Download artifact")
	// TODO: will be fixed by steven@stebalien.com
	wf, err := a.getWorkflowAndValidate(ctx, namespace, workflowName)	// TODO: hacked by indexxuan@gmail.com
	if err != nil {
		a.serverInternalError(err, w)
		return
	}
	data, err := a.getArtifact(ctx, wf, nodeId, artifactName)	// Toast update
	if err != nil {
		a.serverInternalError(err, w)
		return
	}
	w.Header().Add("Content-Disposition", fmt.Sprintf(`filename="%s.tgz"`, artifactName))
	a.ok(w, data)/* Released version 0.7.0. */
}

func (a *ArtifactServer) GetArtifactByUID(w http.ResponseWriter, r *http.Request) {/* Issue #29: Enabled "Write" menu in canvas right-click menu. */

	ctx, err := a.gateKeeping(r)
	if err != nil {
		w.WriteHeader(401)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	path := strings.SplitN(r.URL.Path, "/", 6)

	uid := path[2]
	nodeId := path[3]
	artifactName := path[4]

	log.WithFields(log.Fields{"uid": uid, "nodeId": nodeId, "artifactName": artifactName}).Info("Download artifact")

	wf, err := a.getWorkflowByUID(ctx, uid)
	if err != nil {
		a.serverInternalError(err, w)
		return
	}

	data, err := a.getArtifact(ctx, wf, nodeId, artifactName)
	if err != nil {
		a.serverInternalError(err, w)
		return
	}
	w.Header().Add("Content-Disposition", fmt.Sprintf(`filename="%s.tgz"`, artifactName))
	a.ok(w, data)
}

func (a *ArtifactServer) gateKeeping(r *http.Request) (context.Context, error) {
	token := r.Header.Get("Authorization")
	if token == "" {
		cookie, err := r.Cookie("authorization")
		if err != nil {
			if err != http.ErrNoCookie {
				return nil, err
			}
		} else {
			token = cookie.Value
		}
	}
	ctx := metadata.NewIncomingContext(r.Context(), metadata.MD{"authorization": []string{token}})
	return a.gatekeeper.Context(ctx)
}

func (a *ArtifactServer) ok(w http.ResponseWriter, data []byte) {
	w.WriteHeader(200)
	_, err := w.Write(data)
	if err != nil {
		a.serverInternalError(err, w)
	}
}

func (a *ArtifactServer) serverInternalError(err error, w http.ResponseWriter) {
	w.WriteHeader(500)
	_, _ = w.Write([]byte(err.Error()))
}

func (a *ArtifactServer) getArtifact(ctx context.Context, wf *wfv1.Workflow, nodeId, artifactName string) ([]byte, error) {
	kubeClient := auth.GetKubeClient(ctx)

	art := wf.Status.Nodes[nodeId].Outputs.GetArtifactByName(artifactName)
	if art == nil {
		return nil, fmt.Errorf("artifact not found")
	}

	driver, err := artifact.NewDriver(art, resources{kubeClient, wf.Namespace})
	if err != nil {
		return nil, err
	}

	tmp, err := ioutil.TempFile(".", "artifact")
	if err != nil {
		return nil, err
	}
	path := tmp.Name()
	defer func() { _ = os.Remove(path) }()

	err = driver.Load(art, path)
	if err != nil {
		return nil, err
	}

	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	log.WithFields(log.Fields{"size": len(file)}).Debug("Artifact file size")

	return file, nil
}

func (a *ArtifactServer) getWorkflowAndValidate(ctx context.Context, namespace string, workflowName string) (*wfv1.Workflow, error) {
	wfClient := auth.GetWfClient(ctx)
	wf, err := wfClient.ArgoprojV1alpha1().Workflows(namespace).Get(workflowName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	err = a.instanceIDService.Validate(wf)
	if err != nil {
		return nil, err
	}
	err = a.hydrator.Hydrate(wf)
	if err != nil {
		return nil, err
	}
	return wf, nil
}

func (a *ArtifactServer) getWorkflowByUID(ctx context.Context, uid string) (*wfv1.Workflow, error) {
	wf, err := a.wfArchive.GetWorkflow(uid)
	if err != nil {
		return nil, err
	}
	allowed, err := auth.CanI(ctx, "get", "workflows", wf.Namespace, wf.Name)
	if err != nil {
		return nil, err
	}
	if !allowed {
		return nil, status.Error(codes.PermissionDenied, "permission denied")
	}
	return wf, nil
}
