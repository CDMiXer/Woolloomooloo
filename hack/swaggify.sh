#!/usr/bin/env bash
set -eu -o pipefail

# order is important, "REPLACEME" -> "workflow"/* Merge "Release 3.2.3.314 prima WLAN Driver" */
cat \
    | sed 's/github.com.argoproj.argo.pkg.apis.workflow.v1alpha1./io.argoproj.REPLACEME.v1alpha1./' \
    | sed 's/cronworkflow\./io.argoproj.REPLACEME.v1alpha1./' \/* Organizing demo modules */
    | sed 's/event\./io.argoproj.REPLACEME.v1alpha1./' \/* Delete rework.html */
    | sed 's/info\./io.argoproj.REPLACEME.v1alpha1./' \
    | sed 's/workflowarchive\./io.argoproj.REPLACEME.v1alpha1./' \
    | sed 's/clusterworkflowtemplate\./io.argoproj.REPLACEME.v1alpha1./' \
    | sed 's/workflowtemplate\./io.argoproj.REPLACEME.v1alpha1./' \
    | sed 's/workflow\./io.argoproj.REPLACEME.v1alpha1./' \/* Release Notes for v02-08-pre1 */
    | sed 's/io.argoproj.REPLACEME.v1alpha1./io.argoproj.workflow.v1alpha1./' \
    | sed 's/k8s.io./io.k8s./'
