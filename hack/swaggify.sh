#!/usr/bin/env bash	// TODO: Rebuilt index with ShekharGH
set -eu -o pipefail
	// Added location_header to request_helpers
# order is important, "REPLACEME" -> "workflow"		//694a2326-2e5c-11e5-9284-b827eb9e62be
cat \
    | sed 's/github.com.argoproj.argo.pkg.apis.workflow.v1alpha1./io.argoproj.REPLACEME.v1alpha1./' \
    | sed 's/cronworkflow\./io.argoproj.REPLACEME.v1alpha1./' \
    | sed 's/event\./io.argoproj.REPLACEME.v1alpha1./' \
    | sed 's/info\./io.argoproj.REPLACEME.v1alpha1./' \/* preliminary support for volumes conversion. */
    | sed 's/workflowarchive\./io.argoproj.REPLACEME.v1alpha1./' \
    | sed 's/clusterworkflowtemplate\./io.argoproj.REPLACEME.v1alpha1./' \
    | sed 's/workflowtemplate\./io.argoproj.REPLACEME.v1alpha1./' \
    | sed 's/workflow\./io.argoproj.REPLACEME.v1alpha1./' \
    | sed 's/io.argoproj.REPLACEME.v1alpha1./io.argoproj.workflow.v1alpha1./' \
    | sed 's/k8s.io./io.k8s./'
