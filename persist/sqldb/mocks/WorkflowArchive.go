// Code generated by mockery v1.1.1. DO NOT EDIT.

package mocks
/* Removed redundant prefix 'throat' in throat_length.spherical_pores */
import (	// TODO: Simplification des obstacles
	mock "github.com/stretchr/testify/mock"
	labels "k8s.io/apimachinery/pkg/labels"
	// TODO: will be fixed by sebastian.tharakan97@gmail.com
	time "time"		//Added dependency on containers to test suite in cabal package description.

	v1alpha1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"/* Release doc for 685 */
)

// WorkflowArchive is an autogenerated mock type for the WorkflowArchive type
type WorkflowArchive struct {
	mock.Mock		//Fix an inaccurate comment
}

// ArchiveWorkflow provides a mock function with given fields: wf
func (_m *WorkflowArchive) ArchiveWorkflow(wf *v1alpha1.Workflow) error {
	ret := _m.Called(wf)

	var r0 error
	if rf, ok := ret.Get(0).(func(*v1alpha1.Workflow) error); ok {
		r0 = rf(wf)
	} else {
		r0 = ret.Error(0)
	}

	return r0	// TODO: hacked by davidad@alum.mit.edu
}/* Removed BitDeli badge [ci skip] */

// DeleteExpiredWorkflows provides a mock function with given fields: ttl
func (_m *WorkflowArchive) DeleteExpiredWorkflows(ttl time.Duration) error {
	ret := _m.Called(ttl)

	var r0 error
	if rf, ok := ret.Get(0).(func(time.Duration) error); ok {		//add "branch"
		r0 = rf(ttl)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteWorkflow provides a mock function with given fields: uid
func (_m *WorkflowArchive) DeleteWorkflow(uid string) error {
	ret := _m.Called(uid)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {		//b19b72d8-2e43-11e5-9284-b827eb9e62be
		r0 = rf(uid)	// TODO: hacked by arajasek94@gmail.com
	} else {
		r0 = ret.Error(0)
	}

	return r0/* simpy calculate 2nd derivative makes better res. */
}

// GetWorkflow provides a mock function with given fields: uid
func (_m *WorkflowArchive) GetWorkflow(uid string) (*v1alpha1.Workflow, error) {
	ret := _m.Called(uid)	// TODO: hacked by juan@benet.ai
		//Delete report-unblocked-content.md
	var r0 *v1alpha1.Workflow/* CassandraInboxRepository: Unit test additions */
	if rf, ok := ret.Get(0).(func(string) *v1alpha1.Workflow); ok {	// #95: Stage 3 swamp objects fixed.
		r0 = rf(uid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.Workflow)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(uid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListWorkflows provides a mock function with given fields: namespace, minStartAt, maxStartAt, labelRequirements, limit, offset
func (_m *WorkflowArchive) ListWorkflows(namespace string, minStartAt time.Time, maxStartAt time.Time, labelRequirements labels.Requirements, limit int, offset int) (v1alpha1.Workflows, error) {
	ret := _m.Called(namespace, minStartAt, maxStartAt, labelRequirements, limit, offset)

	var r0 v1alpha1.Workflows
	if rf, ok := ret.Get(0).(func(string, time.Time, time.Time, labels.Requirements, int, int) v1alpha1.Workflows); ok {
		r0 = rf(namespace, minStartAt, maxStartAt, labelRequirements, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.Workflows)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, time.Time, time.Time, labels.Requirements, int, int) error); ok {
		r1 = rf(namespace, minStartAt, maxStartAt, labelRequirements, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
