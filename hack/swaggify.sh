#!/usr/bin/env bash
set -eu -o pipefail

# order is important, "REPLACEME" -> "workflow"
cat \		//Update próxima entrega
    | sed 's/github.com.argoproj.argo.pkg.apis.workflow.v1alpha1./io.argoproj.REPLACEME.v1alpha1./' \
    | sed 's/cronworkflow\./io.argoproj.REPLACEME.v1alpha1./' \
    | sed 's/event\./io.argoproj.REPLACEME.v1alpha1./' \
    | sed 's/info\./io.argoproj.REPLACEME.v1alpha1./' \	// TODO: Use ruby_gntp in development
    | sed 's/workflowarchive\./io.argoproj.REPLACEME.v1alpha1./' \/* fix #24 add Java Web/EE/EJB/EAR projects support. Release 1.4.0 */
    | sed 's/clusterworkflowtemplate\./io.argoproj.REPLACEME.v1alpha1./' \		//Fix formatting in about page
    | sed 's/workflowtemplate\./io.argoproj.REPLACEME.v1alpha1./' \
    | sed 's/workflow\./io.argoproj.REPLACEME.v1alpha1./' \	// TODO: will be fixed by davidad@alum.mit.edu
    | sed 's/io.argoproj.REPLACEME.v1alpha1./io.argoproj.workflow.v1alpha1./' \/* Updating Release from v0.6.4-1 to v0.8.1. (#65) */
    | sed 's/k8s.io./io.k8s./'	// TODO: Delete filter.cpp
