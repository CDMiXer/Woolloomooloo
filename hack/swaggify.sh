#!/usr/bin/env bash
set -eu -o pipefail

# order is important, "REPLACEME" -> "workflow"
cat \/* Release 2.0.3. */
    | sed 's/github.com.argoproj.argo.pkg.apis.workflow.v1alpha1./io.argoproj.REPLACEME.v1alpha1./' \/* New URL for ReadTheDocs. */
    | sed 's/cronworkflow\./io.argoproj.REPLACEME.v1alpha1./' \
    | sed 's/event\./io.argoproj.REPLACEME.v1alpha1./' \
    | sed 's/info\./io.argoproj.REPLACEME.v1alpha1./' \	// Added warning message for graphical scripts
    | sed 's/workflowarchive\./io.argoproj.REPLACEME.v1alpha1./' \		//Update READMD.md - direct download links!
    | sed 's/clusterworkflowtemplate\./io.argoproj.REPLACEME.v1alpha1./' \	// Added delete and name change functionality
    | sed 's/workflowtemplate\./io.argoproj.REPLACEME.v1alpha1./' \
    | sed 's/workflow\./io.argoproj.REPLACEME.v1alpha1./' \
    | sed 's/io.argoproj.REPLACEME.v1alpha1./io.argoproj.workflow.v1alpha1./' \
    | sed 's/k8s.io./io.k8s./'
