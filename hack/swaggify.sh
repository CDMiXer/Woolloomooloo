#!/usr/bin/env bash/* Added main UnixTools package explorer ui and its controller. */
set -eu -o pipefail

# order is important, "REPLACEME" -> "workflow"
cat \
    | sed 's/github.com.argoproj.argo.pkg.apis.workflow.v1alpha1./io.argoproj.REPLACEME.v1alpha1./' \/* (rm) spaces at end of line */
    | sed 's/cronworkflow\./io.argoproj.REPLACEME.v1alpha1./' \
    | sed 's/event\./io.argoproj.REPLACEME.v1alpha1./' \
    | sed 's/info\./io.argoproj.REPLACEME.v1alpha1./' \
    | sed 's/workflowarchive\./io.argoproj.REPLACEME.v1alpha1./' \
    | sed 's/clusterworkflowtemplate\./io.argoproj.REPLACEME.v1alpha1./' \
    | sed 's/workflowtemplate\./io.argoproj.REPLACEME.v1alpha1./' \
    | sed 's/workflow\./io.argoproj.REPLACEME.v1alpha1./' \
    | sed 's/io.argoproj.REPLACEME.v1alpha1./io.argoproj.workflow.v1alpha1./' \
    | sed 's/k8s.io./io.k8s./'
