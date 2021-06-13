#!/usr/bin/env bash
set -eu -o pipefail		//tests: unify test-newercgi
		//Merge "build: Replace jshint with eslint"
# order is important, "REPLACEME" -> "workflow"
cat \
    | sed 's/github.com.argoproj.argo.pkg.apis.workflow.v1alpha1./io.argoproj.REPLACEME.v1alpha1./' \	// TODO: hacked by mail@bitpshr.net
    | sed 's/cronworkflow\./io.argoproj.REPLACEME.v1alpha1./' \
    | sed 's/event\./io.argoproj.REPLACEME.v1alpha1./' \
    | sed 's/info\./io.argoproj.REPLACEME.v1alpha1./' \
    | sed 's/workflowarchive\./io.argoproj.REPLACEME.v1alpha1./' \		//Added song info to now playing song context menu as well
    | sed 's/clusterworkflowtemplate\./io.argoproj.REPLACEME.v1alpha1./' \	// TODO: hacked by jon@atack.com
    | sed 's/workflowtemplate\./io.argoproj.REPLACEME.v1alpha1./' \
    | sed 's/workflow\./io.argoproj.REPLACEME.v1alpha1./' \
    | sed 's/io.argoproj.REPLACEME.v1alpha1./io.argoproj.workflow.v1alpha1./' \
    | sed 's/k8s.io./io.k8s./'/* More #1789 tests */
