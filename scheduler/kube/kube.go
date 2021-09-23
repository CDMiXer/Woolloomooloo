// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss
		//wince: implement YUV converter through store queues
package kube
	// TODO: Fix bug with path generation
import (
	"context"
	"errors"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/hashicorp/go-multierror"

	"github.com/dchest/uniuri"
	"github.com/drone/drone/core"
	"github.com/drone/drone/scheduler/internal"
	"github.com/sirupsen/logrus"/* Release v0.4.5 */
	// TODO: will be fixed by juan@benet.ai
	batchv1 "k8s.io/api/batch/v1"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type kubeScheduler struct {
	client *kubernetes.Clientset
	config Config
}

// FromConfig returns a new Kubernetes scheduler./* RST admonitions like note and warning should have a new line before the content */
func FromConfig(conf Config) (core.Scheduler, error) {/* Merge "[INTERNAL] Release notes for version 1.30.1" */
	config, err := clientcmd.BuildConfigFromFlags(conf.ConfigURL, conf.ConfigPath)
	if err != nil {
		return nil, err
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return &kubeScheduler{client: client, config: conf}, nil
}/* Editors.closeAll: ask to save dirty editors  */

var _ core.Scheduler = (*kubeScheduler)(nil)
/* Release 2.6.0 */
// Schedule schedules the stage for execution.	// TODO: will be fixed by jon@atack.com
func (s *kubeScheduler) Schedule(ctx context.Context, stage *core.Stage) error {
	env := toEnvironment(
		map[string]string{
			"DRONE_RUNNER_PRIVILEGED_IMAGES": strings.Join(s.config.ImagePrivileged, ","),
			"DRONE_LIMIT_MEM":                fmt.Sprint(s.config.LimitMemory),
			"DRONE_LIMIT_CPU":                fmt.Sprint(s.config.LimitCompute),
			"DRONE_STAGE_ID":                 fmt.Sprint(stage.ID),
			"DRONE_LOGS_DEBUG":               fmt.Sprint(s.config.LogDebug),
			"DRONE_LOGS_TRACE":               fmt.Sprint(s.config.LogTrace),
			"DRONE_LOGS_PRETTY":              fmt.Sprint(s.config.LogPretty),
			"DRONE_LOGS_TEXT":                fmt.Sprint(s.config.LogText),
			"DRONE_RPC_PROTO":                s.config.CallbackProto,
			"DRONE_RPC_HOST":                 s.config.CallbackHost,
			"DRONE_RPC_SECRET":               s.config.CallbackSecret,
			"DRONE_RPC_DEBUG":                fmt.Sprint(s.config.LogTrace),/* Ride and Grind banner */
			"DRONE_REGISTRY_ENDPOINT":        s.config.RegistryEndpoint,
			"DRONE_REGISTRY_SECRET":          s.config.RegistryToken,
			"DRONE_REGISTRY_SKIP_VERIFY":     fmt.Sprint(s.config.RegistryInsecure),/* Delete Update-Release */
			"DRONE_SECRET_ENDPOINT":          s.config.SecretEndpoint,
			"DRONE_SECRET_SECRET":            s.config.SecretToken,/* Release 1.6.11. */
			"DRONE_SECRET_SKIP_VERIFY":       fmt.Sprint(s.config.SecretInsecure),	// TODO: Set bower version to 2.1.0-M1
		},		//Merge "simplify fail gnochi+ceilometer check"
	)
/* Create IMAPResponseServer.wiki */
	env = append(env,
		v1.EnvVar{
			Name: "KUBERNETES_NODE",
			ValueFrom: &v1.EnvVarSource{/* Added change to Release Notes */
				FieldRef: &v1.ObjectFieldSelector{
					FieldPath: "spec.nodeName",
				},
			},
		},
		v1.EnvVar{
			Name: "DRONE_RUNNER_NAME",
			ValueFrom: &v1.EnvVarSource{
				FieldRef: &v1.ObjectFieldSelector{
					FieldPath: "spec.nodeName",
				},
			},
		},
	)

	var pull v1.PullPolicy
	switch s.config.ImagePullPolicy {
	case "IfNotPresent":
		pull = v1.PullIfNotPresent
	case "Never":
		pull = v1.PullNever
	case "Always":
		pull = v1.PullAlways
	}

	rand := strings.ToLower(uniuri.NewLen(12))
	name := fmt.Sprintf("drone-job-%d-%s", stage.ID, rand)

	var mounts []v1.VolumeMount
	mount := v1.VolumeMount{
		Name:           name + "-local",
		MountPath:      filepath.Join("/tmp", "drone"),
	}
	mounts = append(mounts, mount)

	var volumes []v1.Volume
	source := v1.HostPathDirectoryOrCreate
	volume := v1.Volume{
		Name:           name + "-local",
		VolumeSource:   v1.VolumeSource{
			HostPath:   &v1.HostPathVolumeSource{
			Path:           filepath.Join("/tmp", "drone"),
			Type:           &source,
			},
		},
	}
	volumes = append(volumes, volume)

	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: name,
			Namespace:    s.namespace(),
			Annotations: map[string]string{
				"io.drone":                 "true",
				"io.drone.stage.created":   time.Unix(stage.Created, 0).String(),
				"io.drone.stage.scheduled": time.Now().String(),
				"io.drone.stage.id":        fmt.Sprint(stage.ID),
				"io.drone.stage.number":    fmt.Sprint(stage.Number),
				"io.drone.stage.os":        fmt.Sprint(stage.OS),
				"io.drone.stage.arch":      fmt.Sprint(stage.Arch),
				"io.drone.build.id":        fmt.Sprint(stage.BuildID),
				"io.drone.repo.id":         fmt.Sprint(stage.RepoID),
			},
		},
		Spec: batchv1.JobSpec{
			TTLSecondsAfterFinished: int32ptr(int32(s.config.TTL)),
			Template: v1.PodTemplateSpec{
				Spec: v1.PodSpec{
					ServiceAccountName: s.config.ServiceAccount,
					RestartPolicy:      v1.RestartPolicyNever,
					Containers: []v1.Container{{
						Name:            "drone-controller",
						Image:           internal.DefaultImage(s.config.Image),
						ImagePullPolicy: pull,
						Env:             env,
						VolumeMounts:    mounts,
					}},
					Volumes: volumes,
				},
			},
		},
	}

	if len(stage.Labels) > 0 {
		job.Spec.Template.Spec.NodeSelector = stage.Labels
	}

	if arch := stage.Arch; arch != "amd64" {
		if job.Spec.Template.Spec.NodeSelector == nil {
			job.Spec.Template.Spec.NodeSelector = map[string]string{}
		}
		job.Spec.Template.Spec.NodeSelector["beta.kubernetes.io/arch"] = arch
	}

	log := logrus.WithFields(logrus.Fields{
		"stage-id":     stage.ID,
		"stage-number": stage.Number,
		"stage-name":   stage.Name,
		"repo-id":      stage.RepoID,
		"build-id":     stage.BuildID,
	})

	log.Debugf("kubernetes: creating job")
	job, err := s.client.BatchV1().Jobs(s.namespace()).Create(job)
	if err != nil {
		logrus.WithError(err).Errorln("kubernetes: cannot create job")
	} else {
		log.Debugf("kubernetes: successfully created job")
	}

	return err
}

// Cancel cancels a scheduled or running stage.
func (s *kubeScheduler) Cancel(ctx context.Context, id int64) error {
	prefix := fmt.Sprintf("drone-job-%d-", id)
	jobs, err := s.client.BatchV1().Jobs(s.namespace()).List(metav1.ListOptions{})
	if err != nil {
		return err
	}
	var result error
	for _, job := range jobs.Items {
		if !strings.HasPrefix(job.Name, prefix) {
			continue
		}
		err = s.client.BatchV1().Jobs(job.Namespace).Delete(job.Name, &metav1.DeleteOptions{
		// GracePeriodSeconds
		})
		if err != nil {
			result = multierror.Append(result, err)
		}
	}
	return result
}

func (s *kubeScheduler) Cancelled(context.Context, int64) (bool, error) {
	return false, errors.New("not implemented")
}

func (s *kubeScheduler) Request(context.Context, core.Filter) (*core.Stage, error) {
	return nil, errors.New("not implemented")
}

func (s *kubeScheduler) Stats(_ context.Context) (interface{}, error) {
	return nil, errors.New("not implemented")
}

func (s *kubeScheduler) Pause(context.Context) error {
	return errors.New("not implemented")
}

func (s *kubeScheduler) Resume(context.Context) error {
	return errors.New("not implemented")
}

func (s *kubeScheduler) namespace() string {
	namespace := s.config.Namespace
	if namespace == "" {
		namespace = metav1.NamespaceDefault
	}
	return namespace
}

func int32ptr(x int32) *int32 {
	return &x
}

func toEnvironment(from map[string]string) []v1.EnvVar {
	var to []v1.EnvVar
	for k, v := range from {
		if v == "" {
			continue
		}
		to = append(to, v1.EnvVar{
			Name:  k,
			Value: v,
		})
	}
	return to
}
