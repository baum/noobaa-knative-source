/*
Copyright 2019 The Knative Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	appsv1 "k8s.io/api/apps/v1"
	"knative.dev/eventing/pkg/apis/duck"
	"knative.dev/pkg/apis"
)

const (
	// NooBaaSourceConditionReady has status True when the NooBaaSource is ready to send events.
	// NooBaaSourceConditionReady = apis.ConditionReady

	// NooBaaSourceConditionSinkProvided has status True when the NooBaaSource has been configured with a sink target.
	NooBaaSourceConditionSinkProvided apis.ConditionType = "SinkProvided"

	// NooBaaSourceConditionDeployed has status True when the NooBaaSource has had it's deployment created.
	NooBaaSourceConditionDeployed apis.ConditionType = "Deployed"
)

var NooBaaSourceCondSet = apis.NewLivingConditionSet(
	NooBaaSourceConditionSinkProvided,
	NooBaaSourceConditionDeployed,
)

// GetCondition returns the condition currently associated with the given type, or nil.
func (s *NooBaaSourceStatus) GetCondition(t apis.ConditionType) *apis.Condition {
	return NooBaaSourceCondSet.Manage(s).GetCondition(t)
}

// InitializeConditions sets relevant unset conditions to Unknown state.
func (s *NooBaaSourceStatus) InitializeConditions() {
	NooBaaSourceCondSet.Manage(s).InitializeConditions()
}

// GetConditionSet returns NooBaaSource ConditionSet.
func (*NooBaaSource) GetConditionSet() apis.ConditionSet {
	return NooBaaSourceCondSet
}

// MarkSink sets the condition that the source has a sink configured.
func (s *NooBaaSourceStatus) MarkSink(uri *apis.URL) {
	s.SinkURI = uri
	if len(uri.String()) > 0 {
		NooBaaSourceCondSet.Manage(s).MarkTrue(NooBaaSourceConditionSinkProvided)
	} else {
		NooBaaSourceCondSet.Manage(s).MarkUnknown(NooBaaSourceConditionSinkProvided, "SinkEmpty", "Sink has resolved to empty.")
	}
}

// MarkNoSink sets the condition that the source does not have a sink configured.
func (s *NooBaaSourceStatus) MarkNoSink(reason, messageFormat string, messageA ...interface{}) {
	NooBaaSourceCondSet.Manage(s).MarkFalse(NooBaaSourceConditionSinkProvided, reason, messageFormat, messageA...)
}

// PropagateDeploymentAvailability uses the availability of the provided Deployment to determine if
// SampleConditionDeployed should be marked as true or false.
func (s *NooBaaSourceStatus) PropagateDeploymentAvailability(d *appsv1.Deployment) {
	if duck.DeploymentIsAvailable(&d.Status, false) {
		NooBaaSourceCondSet.Manage(s).MarkTrue(NooBaaSourceConditionDeployed)
	} else {
		// I don't know how to propagate the status well, so just give the name of the Deployment
		// for now.
		NooBaaSourceCondSet.Manage(s).MarkFalse(NooBaaSourceConditionDeployed, "DeploymentUnavailable", "The Deployment '%s' is unavailable.", d.Name)
	}
}

// IsReady returns true if the resource is ready overall.
func (s *NooBaaSourceStatus) IsReady() bool {
	return NooBaaSourceCondSet.Manage(s).IsHappy()
}
